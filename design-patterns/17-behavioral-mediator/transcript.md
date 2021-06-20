# Mediator -- Facilitates communication between components

## Overview

Let's talk about the mediator design pattern.
What is the motivation for using the mediator.

Well if you have a system where components can go in and out of the system at any time meaning they can be created or destroyed like for example you could have a chat room.

So in the chat room you have participants they go into the chat room but they can leave the chat room at any one time or for example players in some sort of online game they can go in and then can go out.

Then it doesn't really make sense for them to have like direct links to one another.

So if you have several participants in a game it would make sense for you to have pointers from one to another.

Because remember they are leaving the game at any time.

We're leaving the chat room and then your point is are effectively dead.

So you don't want any kind of incorrect operation.

So the references may go dead and the solution to this problem is that you instead of having them refer to one another you have all of these constructs refer to some central component that actually takes care of the communication and that component happens to be the mediator or so the mediator is quite simply a component that facilitates the communication between different components without them necessarily being aware of each other or having any kind of direct pointers for example to each other so they can certainly be aware of each other in the sense of being notified by the mediator that they exist but they don't have direct links to one another.

## Chat Room

The best and most succinct example of the media design pattern is the simulation of a chatroom.

So that's exactly what we're going to build.

So we'll begin by defining a participant of the chat room.

So we're gonna have a time called person.

So this is going to be somebody that is participating in a group chat.

A person has a name which is a string and the person has a pointer to the mediator.

The mediator in our case is the chat room the chat room is the thing that allows different people to communicate with one another without necessarily being aware of one another's presence.

So we'll have a chat room like this and we'll also have a chat log.

So this is just going to be a log of all the messages that we have sent or received.

That's going to be stored inside the person.
We don't actually need it.

But I want to illustrate how the thing is done in the real world.

So let's make a construct of a person we'll just initialize the name.

That's the only thing that we want to initialize.

And then what we're going to do is we're going to have a bunch of methods on the person.

So first of all what you need to be able to do is you need to be able to receive a message from somewhere.

You can receive a message from another person or you can receive a message from the chatroom itself.

For example let's suppose you get kicked out the chatroom can tell you you were kicked out or maybe the room just got disconnected or deleted or something.

So we're gonna have our first method.

So this is going to be a method called receive.

So a person can receive a message and we can specify both the sender as well as the message as strings like so.

So the idea here is that I'm going to format the message and I'm also going to tell you who's chat session we're actually in.

So first of all I'm going to do the formatting so I'm going to say ask is and I'm just going to printf here.

So let's s print f the fact that we have a sender.

So we're going to have a sender has a string and then we have a message.

So here we specify the sender.
And here we specify the message like so.
OK.

So now that we have this what I'm going to do is I'll just print this out to the console so we'll print out whose chat session it actually is.

So I guess instead of print line we'll do a print f here.

So I'm going to say that I have the person's chat log more chat session and then I'm going to specify the actual message the full message.

So here the name of the person is of course pedo name and the the actual message is s the thing that we've prepared in the previous line right here and then we'll we're going to do is we're going to add that to the chat logs.

So our chat log let's just append it to the chat log like so.

So that's our way of receiving a message from somewhere and then we'll also have a way of actually saying something in the room.

So when you say something in the room it's broadcast to every single person so we'll have a method called say where you simply specify the message and then we take the room.

So Peter room that's our mediator by the way.
That's the central component through which everything goes.

And we do a broadcast we broadcast the name our name as well as the message that we want to actually send to everyone.

So as you can see this uses a broadcast method though we haven't defined yet.

But before we do let's add another method of person so a person is going to have a method called private message.

So unlike a broadcast the private message gets sent to a particular individual.

So we specify who we're sending the message to and the message itself and both of them are going to be strings and then all we do is we say Pitot room.

Once again we're using that mediator and we say message.

So we send the message we specify who it's from who it's for as well as the message itself.

So as you can see what's happening here is every single person every single participant in the chatroom has a pointed to the room they're in.

So they can join the room and when they join the room they can receive messages using the receive method and then they can also send messages whether it's messages being sent to the entire room or messages being sent to specific individuals.

OK.

So now let's define the actual room so the room is just a collection of people being together and chatting together.

So we'll have type chat room defined as a struct.

So here I'll just have on the Ray of person point.

So we'll have an array of person pointers and what we can do is we can first of all support the idea of broadcasting so you can specify a message as well as the source of that message and you can send it to every participant in the chat room chatroom so here we'll have a method on the chat room.

This method is going to be called broadcast.
OK.

So we specify the source of the broadcasts who is actually doing the broadcasting.

And we also specify the message and they're both strings and then what we do is we call receive on every single participant in the chat room.

So we say for on this car call my p in range of CDO people so we take all the people and we can only send that message if the source of the message isn't the target.

So you don't want to send the message to yourself you only want to send it to every other participant beside yourself.

So we say if we don't name is not equal to the source then you can actually send it because if you don't name is equal to the source then we are sending a message from person next to that person which makes absolutely no sense because they're already typed the message they don't want to now see the message once again.

So we say P don't receive.
So the participant receives a message from the source.
And here is the message itself.

So this is how you can broadcast to every single participant in the chatroom but you can also do other things like you can send private messages.

So here you can have a chat room a method called message and here things are more interesting because you specify the source of the message.

Mm hmm.

You also specify the destination who you're sending the message to and the message itself.

Let's just use very short variables as RC DSD and as are going to be of type String like so.

Okay.

So once again the idea is very similar except that this time round it's kind of like what we did in broadcast except we need to find the person that we're sending the message to.

So if B The name is in fact equal to the destination then P receive and then here we specify the source and the message.

So this allows us to send a targeted message from one participant to another and of course if you imagine this message being sent.

Imagine what happens if a person has left by the time you send a message.

So you send the message but that person you send it to has actually left.

Well there is absolutely no problem.

We will simply not be able to find it when we iterate over the people.

So there's gonna be no message send there's nobody there to receive it.

Okay.

So the final thing I want to add is a method for actually getting a person to join the chat room.

So chat room we're gonna have a method called Join where you specify a pointer to the person that's joining the room and let's suppose that we want to do a broadcast so whenever somebody joins the room we're going to do a broadcast will inform everyone that somebody has joined the chat.

So the join message in our case is going to be a PDA name plus let's say joins the chat.

Okay.
And then we're going to broadcast this message.

So we'll say see adult broadcast and in this particular case the first argument is the source of the message.

Now in this case the source isn't the person joining the room.

So we'll just say that the room is broadcasting the message and here is the joint message itself.

And then what we can do is we can take the person that's joining and we can assign the person's room pointer so we can say P that room equals c.

That way we make sure that the person knows the room they're in and we can also add this person to the list of people so we can say see the people and we can append to that list we can append this person.

Okay.
So now we have everything ready and we can try out our chat rooms so here is the chat room.

So I'll just make a chat room like so and we can make to participants to begin with.

We'll have a participant John.
So that's going to be a person called John.
We'll have another one.
Let's say we have Jane.

So John and Jane are going to both join the chat room and we'll take a look at the kind of messages that are generated so first of all both of them have joined so we say room door join John.

And that has to be a capital J here and then room door join Jane.

Okay.
So let's get them actually saying things.

So John is going to say hi room and Jane is going to say Oh hey John.

Okay now let's make it more interesting let's make another participant Simon.

So Simon is a new person called Simon.
Simon Well the room Simon joins this particular room.

So let's put Simon here and then Simon says something so Simon does say Hi everyone and let's try that private messaging functionality let's suppose Jane actually no Simon so Jane.

Go ahead goes ahead and she sends a private message to Simon saying Glad you could join us.

Like so.
So let's try to simulate this entire scenario.

Let's actually run this and see what kind of output we get.

Well one problem is we forgot a line break.
So that's disappointing.

And that is something that we need to add right here when you actually receive a message.

So let's try this once again.
Let's see if it's better this time.
OK.

So first of all let's take a look at that chart session.

So when John joins the chat session nobody receives a message but when Jane joins the chat then John receives the message as the only other participant.

And then they start chatting so John says hi room and that only appears in Jane's chat session.

And then Jane says Oh hey John.

And this only appears in John's chat session so you'll notice that Jane is saying something but she doesn't get it as part of her chat session because well it's not necessary.

And then Simon joins the chat and you'll notice that both John's chat session and Jane's chat session informs us that the room is broadcasting a message that Simon joins the chat.

So Simon says Hi everyone and once again both John and Jane receives this information.

And finally here's the private message.

So Jane sends a private message saying Glad you could join us and it's only Simon that receives this private message because he's the one that the message is targeted to.

So this entire scenario illustrates the mediator.
So what is the mediator here.
Well the mediator here is the chat room.

Basically we have a bunch of people who are participants in the chat room but they're not directly aware of one another.

What do I mean by not directly what I mean is that they don't half point us to one another.

So for example the person here doesn't really know about other people's presence it doesn't have to track people entering the system and exiting the system because all it has is it has a reference to the mediator in this case it's a pointer and it uses this pointer so so that it can actually say things like say things to the room or send private messages and also the mediator on the other hand is able to inform every participant of something happening and that is done using the received method.

So the mediator here is just a central component that everyone knows and everyone can connect to and people can safely communicate with one another without being afraid that for example they will call something on a nil pointer.

So if you had a private message where you had to have a pointer to the person you're sending the private message to you could run into a situation where that person no longer exists because maybe they left the system altogether.

So the mediator here provides an additional layer of safety as well.

### Chat Room code: behavioral.mediator.chatroom.go

```go
package mediator

import "fmt"

type Person struct {
  Name string
  Room *ChatRoom
  chatLog []string
}

func NewPerson(name string) *Person {
  return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
  s := fmt.Sprintf("%s: '%s'", sender, message)
  fmt.Printf("[%s's chat session] %s\n", p.Name, s)
  p.chatLog = append(p.chatLog, s)
}

func (p *Person) Say(message string) {
  p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
  p.Room.Message(p.Name, who, message)
}

type ChatRoom struct {
  people []*Person
}

func (c *ChatRoom) Broadcast(source, message string) {
  for _, p := range c.people {
    if p.Name != source {
      p.Receive(source, message)
    }
  }
}

func (c *ChatRoom) Join(p *Person) {
  joinMsg := p.Name + " joins the chat"
  c.Broadcast("Room", joinMsg)

  p.Room = c
  c.people = append(c.people, p)
}

func (c *ChatRoom) Message(src, dst, msg string) {
  for _, p := range c.people {
    if p.Name == dst {
      p.Receive(src, msg)
    }
  }
}

func main() {
  room := ChatRoom{}

  john := NewPerson("John")
  jane := NewPerson("Jane")

  room.Join(john)
  room.Join(jane)

  john.Say("hi room")
  jane.Say("oh, hey john")

  simon := NewPerson("Simon")
  room.Join(simon)
  simon.Say("hi everyone!")

  jane.PrivateMessage("Simon", "glad you could join us!")
}
```

## Summary

Let's try to summarize what we've learned about the mediator design pattern.

So the idea is that you create a mediator as a kind of central component and then each object in the system gets to point to the mediator.

So for example you if your objects have factory functions inside those factory functions you can provide that mediator as a point of assignment.

And then the mediator engages in bi directional communication with the connected components.

So on the one hand the mediator has methods that the components can call because they though refer to the mediator and on the other hand the components have methods that the mediator can call so for example in the case of a broadcast the mediator can send one message to every single component that is actually connected to it.

And we have plenty of different libraries like reactor extensions for example that make this kind of communication a lot easier to implement.
