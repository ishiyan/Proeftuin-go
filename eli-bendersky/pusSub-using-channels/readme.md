# PubSub using channels in Go

[Article](https://eli.thegreenplace.net/2020/pubsub-using-channels-in-go/)
[github](https://github.com/eliben/code-for-blog/tree/master/2020/go-pubsub)

The idiomatic way of writing concurrent code in Go is as a collection of goroutines communicating over channels.
In my experience, the Publish-subscibe pattern (PubSub) comes up often as a way to structure code.
The pattern presented here has topic-based subscriptions, but publish-subscribe can appear in other disguises as well.
In its most simple form, it could be a goroutine that produces data and wants to notify a group of other goroutines of that data,
with each downstream goroutine having access to the data separately (rather than on a first-come-first-serve basis as in a work queue).
If "PubSub" doesn't ring a bell, you might be familiar with its alter egos "message broker" and "event bus".

In this post I'll present a brief overview of some design decisions that arise when implementing PubSub for a Go application.
To be clear: this is PubSub for in-process communication between multiple goroutines over channels.
It does not attempt to solve a distributed PubSub problem, which requires sophisticated mechanisms for fault-tolerance.
Within a single Go process we assume goroutines don't just fail and all data sent into channels can be reliably read from the other end.

Let's start with a simple and incomplete implementation.
We'll have the type Pubsub with some methods, which clients can use to subscribe to topics and publish on topics:

```go
type Pubsub struct {
  mu   sync.RWMutex
  subs map[string][]chan string
}
```

The key data structure here is `subs`, which maps topic names into a slice of channels.
Each channel represents a subscription to the topic. I'll talk more about the lock later.

The struct fields aren't exported.
Clients interact with Pubsub solely using its methods. Let's start with a constructor:

```go
func NewPubsub() *Pubsub {
  ps := &Pubsub{}
  ps.subs = make(map[string][]chan string)
  return ps
}
```

Now, a `Subscribe` method through which clients can subscribe to new topics. To subscribe, the client will provide:

- The topic it's interested in.
- A channel on which Pubsub will send it new messages for this topic from now on.

```go
func (ps *Pubsub) Subscribe(topic string, ch chan string) {
  ps.mu.Lock()
  defer ps.mu.Unlock()

  ps.subs[topic] = append(ps.subs[topic], ch)
}
```

The code is very concise thanks to Go's default value semantics.
If `ps.subs` has no topic key, it returns a default value for its value type, or an empty slice of chan string.
This can be appended to and the result is what we expect regardless of the initial contents of `ps.subs`.

Publishing on the `Pubsub` is done with the `Publish` method, which takes a topic and the message:

```go
func (ps *Pubsub) Publish(topic string, msg string) {
  ps.mu.RLock()
  defer ps.mu.RUnlock()

  for _, ch := range ps.subs[topic] {
    ch <- msg
  }
}
```

Once again the default value semantics in Go are useful.
If there are no subscribers to topic, `ps.subs[topic]` is an empty slice so the loop doesn't run.

This is the place to mention the lock.
One of Go's most famous philosophies is "share memory by communicating", but Go is also a pragmatic language.
When we have a shared data structure accessed by multiple goroutines, it's OK to use a lock to protect access to it if this results in the clearest code.
In our case, each Pubsub method starts with a lock + defer unlock sequence, so the code is really simple.
We do have to be very careful about blocking inside Pubsub methods though; more on this shortly.

Note that we don't have an Unsubscribe method. This is left as an exercise to the reader.

## Closing the subscription channels

The code shown so far has a serious issue.
The channels on which messages are sent aren't closed; this is not great, because there's no way for subscribers to be notified that no more messages are going to be sent.
In Go, closing channels is important once we're done sending on them, because closing a channel is a signal that some job is done and resources can be cleaned up.

Here is a version of the code with a Close method:

```go
type Pubsub struct {
  mu     sync.RWMutex
  subs   map[string][]chan string
  closed bool
}
```

We're adding a closed flag to the Pubsub struct. It's initialized to false in the constructor. Publish is modified to:

```go
func (ps *Pubsub) Publish(topic string, msg string) {
  ps.mu.RLock()
  defer ps.mu.RUnlock()

  if ps.closed {
    return
  }

  for _, ch := range ps.subs[topic] {
    ch <- msg
  }
}
```

And we add a new Close method:

```go
func (ps *Pubsub) Close() {
  ps.mu.Lock()
  defer ps.mu.Unlock()

  if !ps.closed {
    ps.closed = true
    for _, subs := range ps.subs {
      for _, ch := range subs {
        close(ch)
      }
    }
  }
}
```

When a Pubsub is done, Close ought to be called to signal on all the subscription channels that no more data will be sent.

Note that these channels weren't created by Pubsub; they are provided in calls to Subscribe.
Is Pubsub.Close the right place to close them? This is a good question.
In general, it is idiomatic for the sending side to close a channel, because this is its way to signal to the receiving side that no more data is going to be sent.
Moreover, since sending on a closed channel panics, it's dangerous to close channels on the receiving side because then the sending side doesn't know that the channel it is sending into may be closed.

This brings us to the more important topic of where should these channels be created in the first place.
Is creating them outside Pubsub and passing them in the right design, or should Pubsub create them?

## Buffering in pubsub channels

The critical issue here is blocking. Recall the sending loop in Publish:

```go
for _, ch := range ps.subs[topic] {
  ch <- msg
}
```

If `ch` is unbuffered, then `ch <- msg` will block until the message is consumed by a receiver.
This prevents Pubsub from notifying other subscribers on the same channel.
Is this the desired behavior? Not likely.
Unless you can guarantee that receivers consume messages from subscriptions very quickly, it may be a good idea to buffer the channels.
A buffer of size 1 would make it much more robust, wherein the publishing loop could finish notifying all topic subscribers quickly (unless a receiver is badly backed up and hasn't even consumed the previous message yet).

In our current design, channels are created outside Pubsub, so their buffering is determined by clients.
This has both positives and negatives:

- Positive: Pubsub doesn't know how clients consume the channels, so it doesn't have to guess what buffer size is appropriate when creating a channel. The client passes it a channel that's already created with the right buffer size.
- Negative: the correctness of Pubsub becomes dependent on its clients. A slow client that passed in an unbuffered channel can block all other clients from consuming their messages.

## Creating the subscription channels in Pubsub

An alternative design is to create subscription channels in Pubsub. Only the Subscribe method would have to change. Here it is:

```go
func (ps *Pubsub) Subscribe(topic string) <-chan string {
  ps.mu.Lock()
  defer ps.mu.Unlock()

  ch := make(chan string, 1)
  ps.subs[topic] = append(ps.subs[topic], ch)
  return ch
}
```

Note that the buffer size is hardcoded to 1.
While this is a good default, we may want to let the client configure the buffer size with an argument.
This can either be done in the constructor for all subscriptions, or in Subscribe with a different buffer size per subscription.

This version of Pubsub has the nice property that it both creates and closes the channels, so the separation of responsibilities is cleaner.
Subscribers just get a channel and listen on it until it's closed.

One slight inconvenience with this approach is that clients may want to subscribe the same channel to multiple topics.
In the previous version of Pubsub they could do so by passing in the same channel to multiple Subscribe calls; in this version they cannot.

However, subscribing the same channel to multiple topics is problematic in other ways.
For example, Pubsub may attempt to close the same channel multiple times when done - this panics.
We'd have to add special provisions to Close to avoid that (such as keep a set of all channels already closed).

In general, I would recommend avoiding this and sticking to a cleaner one-channel-per-subscription approach.
In case the client wants to use the same range loop to receive from multiple topics, it's easy to use some kind of channel fan-in solution instead.

## Doing each send in a goroutine

When we discussed the danger of ch <- msg blocking all clients, you may have wondered why we don't just perform each send in its own goroutine.
Here is a version of Publish that does this:

```go
func (ps *Pubsub) Publish(topic string, msg string) {
  ps.mu.RLock()
  defer ps.mu.RUnlock()

  if ps.closed {
    return
  }

  for _, ch := range ps.subs[topic] {
    go func(ch chan string) {
      ch <- msg
    }(ch)
  }
}
```

Now it doesn't matter how much buffering each channel has; the send will not block any other sends because it runs in its own goroutine.

There may be performance implications, of course.
Even though starting and tearing down goroutines is very quick, do you really want a new one to run for every message?
The answer depends on your particular application. When in doubt, benchmark it.

But performance implications are not the most serious potential issue with this code.
It decouples the places where data is sent on subscription channels and where these channels are closed, which always leaves me a bit uneasy.

Consider a slow client that causes its subscription channel to block for a long while.
Meanwhile, Pubsub may be closed and attempt to close the channel.
But closing channels that have writes pending on them is bad - it's a race condition, which is one of the worst kinds of bug to have.
In the original code this can't happen because Publish holds a lock that prevents Close from running at all.

## Conclusion

The goal of this post was to demonstrate some design choices for a simple yet functional piece of code.
Channels in Go are powerful, but they're not magic.
Difficult questions of ownership and ordering still arise, and it's instructive to think through a single problem from multiple angles.

Of the approaches presented here, I personally prefer the one where Subscribe creates new channels and returns them.
This approach is the most conceptually simple, IMHO, because the ownership of these channels is the most centralized.
Pubsub creates them, sends on them, and closes them.
For a client, the life cycle of a subscription channel is very clear: a new channel is created by Subscribe and can be read from until it's closed.
Calling Pubsub.Close will close all outstanding subscription channel and is useful for cleanup.
If we need configurable buffering, this is easy to add.
