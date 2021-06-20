# State -- Fun with Finite State Machines

## Overview

All right.
Let's talk about these state design pattern.
So what is the motivation what is this all about.

Well if you consider like an ordinary telephone then what you do with the telephone really depends on the state of the phone and the State of the phone line.

So for example if the phone is ringing or if you want to make a call for example you can pick up the phone if well to to actually make a call and to talk to somebody on the phone must be off the hook.

You cannot talk to people while the phone is still on the hook.

And if you try calling someone then it's busy.

You put the handset down or for example you might leave a message if the answering machine comes on.

So the idea is that you have all these changes in state so the phone goes from the on home state to the off hook state for example and changes in state.

They can be explicit meaning you do something or they can be in response to an event and that takes us to the observer pattern.

So something happens and you change from one state to another.

So for example somebody calls you which is an event and the phone go goes from the on hook state to the ringing state and then you can respond to that new state.

So the state design pattern is basically a pattern in which the objects behavior is determined by its state and the object gets to transition from one state to another.

And something needs to trigger a transition that could be action from you the user or it could be an action from an external system.

Now we also have like a formal definition of this whole thing with with states and transitions and the formalized construct which actually manages all of this is called a state machine.

And we also use the term finite state machine when the state machine has a specific starting state and also a specific terminal state after which the execution of the state machine is finished.

## Classic Implementation

Before we start talking about state machines as they are really used in the real world I want to show you an academic example that you're likely to find in many books and online examples and others the idea of states being represented by structures and having some sort of replacement of one struct with another.

So it all sounds confusing and in actual fact the demo I'm going to do is also going to be a rather confusing so I'm not really hoping to explain it fully here but I will show you the demo or the classic implementation of a simple state machine with just two states.

So imagine you have a light and the lighthouse two states it's on or it's off.

That's pretty much it.

But even on the basis of this very simple definition what we can do is we can build a rather grotesque looking state machine not the kind of state machine you will do in the real world.

But once again this is a classic example is one of the classic examples so I want you to see it.

I want you to appreciate that it's possible and also to appreciate the problems with it.

So what we're going to do is we're going to have a couple of types.

We'll have a type called switch so a switch is going to be that thing which allows you to turn something on or off.

So typically what you would do is you'd have a function for switching the light on on like so.

And we'll fill that in in a moment and we have a function on this.

So sort of a method of the switch struct called off which is also going to be filled in in a while.

OK.

But in addition to this which just in case you thought this which is kind of a self-contained thing we also have this idea of a state now a state can also be represented as a as a bunch of struct but we need some sort of interface for the state.

So what we do is we define a state as an interface and for this interface we can define two members we can define a member for turning something on which takes a switch pointer and we can have and members a effectively a method for turning something off.

Now you're probably thinking Well hold on.

This is some really heavy over engineering here.

Why can't we just keep everything inside us which why do we have to have the state idea.

Well you see it in just a moment.

Basically the idea is that the switch kind of stays the same but you use implementers of this interface and you switch from one implementer to another.

What I mean by this is that switch has a member called State of type State State State like this.

And when we switch the light on or we switch the light off we basically replace the value of this particular variable but the way this is done is not particularly intuitive shall we say.

So moving on what I'm going to do is I'm going to have a struct called base state meaning you have a stroke called base state.

Now this struct isn't going to contain anything but what it is going to do is it's going to provide the Deep Field behaviors for the state interface meaning that these are going to be things that you can subsequently aggregate as part of some other struct.

So I'm going to go ahead and I'm going to implement the methods of a state interface.

And here we're going to do something weird.
So this is a base state.

And in the base state whenever somebody tries to switch the light on I'm going to tell them that the light is already on.

And similarly I will tell them here that the light is already off.

Now I know it's still probably very confusing as to why we're defining this base state which also makes an assumption that we haven't really switched the state from one to another.

Well the idea is that you really define on and off state subsequently as separate strikes.

Now this is for many reasons very inefficient but we can define an on state just gonna be a struct which just uses base state.

So it's also base state and.
Well let's have a constructor for it.

Let's just have a factory function for creating it so func you on state.

So in your on state is going to be initialized like the following.

Let's actually output a bit of debugging information here let's say lights turned on just so that we know so the light gets turned on and then we return having you on state where we have to specify the base state we have to initialize it as well.

So this is how you define an on state.

And the idea is that whenever you're working with the on state the only operation which is allowed on the on state is an operation to turn something off.

But remember on state is what you could say in a way that anon state as a base state meaning it has the same map methods but the method which which actually turns something off here says that the light is already off so we need to replace it basically.

So we go here and we make a function.

We'll make a method of the on state class called off.

Now this method takes a switch.
And here is what happens.

So first of all I'll once again say something.

This time around we are really turning the light off.

So turning the light off.
And then here comes the really crazy part.
So we have a pointer to a switch.

Now if we go to the switch up here you can see that the switch has a state.

And what we perform whenever we are really switching the light off is we replace that state completely which means we do as W does state equals new off state.

Now we don't have this constructor yet but we're going to do it in just a moment.

So this is how it goes.

Basically you have a pointed to a switch and you replace that switches state with another state.

And let's do the off state.

So we'll have a function called New off state which creates an off state and by the way we don't have off state yet.

It's the same definitions for the on state.
So type of state struct.

And that has a base state and then we have a function called New off state of state.

And here we print that the light is turned off and we return off state where the base state is also initialized.

OK.

And now for the off state we do the symmetrical thing that we did for the on state here we defined a method called off.

And so here we need to define a method called on so func o off state off on rather take a switch.

And once again I'll output some debugging information here as a turning lights on and then as w dot state equals new on state.

So this is really complicated way of switching the state.

So first of all you have these on and off invocations which kind of override what is given by the base state.

So the base state provides default implementations.

So if you call on on the light that's already on it's going to tell you light is already on.

Now we come back here to the switch implementations and we can actually replace these members so whenever you want to turn the light on you basically call the on method on the state so you say it is not state DOT on on but you provide the switch and similarly here you say as Dot state DOT off and you provide the switch.

So what's happening here is effectively double dispatches the kind of double this patch that happens on the visitor design pattern except that unlike the visitor design pattern what's happening here is completely unnecessary.

Meaning that this entire model can be simplified to to be much simpler.

But let's actually go ahead and finish the examples so the idea is that you make a switch.

OK.

We don't have a constructive for the switch yet.

That's that's something that we need to do and also provide the default initial state.

So fun can you switch news which which.

Like so and.
Okay.

So here we return the switch and we get to decide on the initial state.

For example.
Let's start with the off state.

So this is how you set up a switch and now coming back here so.

So we make a switch we can switch the light on we can switch the light off and then we can try switching the light off again.

So let's just duplicate this line.
OK.

So let me just run this and we'll take a look at what we're getting here no.

It's taking asylum.
OK.
So we're turning.

So so so initially the light is turned off.

That's the initial state.
Now then we're turning the light on.
And suddenly the light is on.
Then we're turning the light off.
And then the light is off.

And here's the interesting part when we try to turn the light off again.

Let's actually trace through the code what's going on here so as w dot off gets called now.

The thing about this is that the state is already an office state.

If you look at the off state right here you can see that the upstate doesn't really have an off method.

It only has an on method.
So the question is well how can we invoke this then.

And the answer is that because we're aggregating the base state we're actually calling the off method of base state right here.

So here is base state here is off on the base state.

So we're calling this instead of calling the actual implementation on the the type that aggregates this so where we're calling this method which tells us that the light is already off.

So this is a purely academic example that you're looking at.

It's not the kind of stuff that you are likely to be building in the real world.

But it's a good illustration that I recommend that you download the source code here and you take a look at this yourself and try tracing through all the different invocations and see what's actually going on.

Because essentially the way that state management happens here is by replacement.

So in order to set a new state what you do is you have some object and then the state of that object gets replaced by a new state.

But the interesting thing is that the replacement is done by the state itself meaning that if you're currently in the off state it is the off state.

You can see the off state here it is the off state that performs the replacement and the adjustment to the state.

So this is the classic definition of the state design pattern when you have states switching themselves to other states in some other external system but in the real world this typically doesn't get implemented this way.

So consider this as simply an academic demo.

### Classic Implementation code: behavioral.state.classic.go

```go
package state

import "fmt"

type Switch struct {
  State State
}

func NewSwitch() *Switch {
  return &Switch{NewOffState()}
}

func (s *Switch) On() {
  s.State.On(s)
}

func (s *Switch) Off() {
  s.State.Off(s)
}

type State interface {
  On(sw *Switch)
  Off(sw *Switch)
}

type BaseState struct {}

func (s *BaseState) On(sw *Switch) {
  fmt.Println("Light is already on")
}

func (s *BaseState) Off(sw *Switch) {
  fmt.Println("Light is already off")
}

type OnState struct {
  BaseState
}

func NewOnState() *OnState {
  fmt.Println("Light turned on")
  return &OnState{BaseState{}}
}

func (o *OnState) Off(sw *Switch) {
  fmt.Println("Turning light off...")
  sw.State = NewOffState()
}

type OffState struct {
  BaseState
}

func NewOffState() *OffState {
  fmt.Println("Light turned off")
  return &OffState{BaseState{}}
}

func (o *OffState) On(sw *Switch) {
  fmt.Println("Turning light on...")
  sw.State = NewOnState()
}

func main() {
  sw := NewSwitch()
  sw.On()
  sw.Off()
  sw.Off()
}
```

## Handmade State Machine

In the previous lesson we looked at a rather academic example of the classic implementation of the state design pattern and that's probably not the kind of state machine that you actually want to build.

One of the reasons is that in most cases the states and the transitions should not be defined by some heavyweight constructs like structure.

For example you can define states by just defining a bunch of constants.

So here I can define a state as simply an ant and then I can just have a bunch of integers.

Now let's imagine that we want to simulate a situation where we're simulating a phone call.

So you pick up the phone and you call somebody and there's lots of things that can happen for example you get placed on hold or you get to leave a message or you just get tired of waiting and you put the phone down that sort of thing.

So the idea is that we can represent each of the states over the phone with a constant with a an integer in this case so let's go ahead and let's make a bunch of constants.

So one constant will be when the phone is off the hook.

So you can see I'm using these state type here saw a defined a type called state the phone can be connecting to somebody.

It can be connected.

We can be placed on hold and we can just say that the phone is back on the hook which means that we're done using the phone.

OK.
So these are a bunch of states.

Now there's one thing one problem and go and that is the printing of these constants like if you want to print them a strings you would have to use a string or a generator or some other kind of generate.

I'm just going to paste the implementation here so this is basically the definition of the string method on a state.

And here I have a bunch of different just just strings being returned for every single case.

OK so these are the states of the system.

These are the states that the system can be in.

And then we can have the triggers.

So the triggers are explicit definitions of a what can cause us to go from one state to another.

So for example when you dial a call you can be connected.

So you transition from a state of off hook to a state connected for example.

So let's also have a type called Trigger and let's define a bunch of triggers.

So we'll have a call dialed.

We'll also have hung up call connected placed on hold taken off hold and left message.

So these are some of the things that you can do and as a result of leaving the message the phone gets placed back on the hook because when you're done using the phone effectively.

So once again what I would typically do behind the scenes and we're not going to do this here is I would generate a string implementation so the string your interface here for printing all of these constructs.

OK so now what we need to do is we need to define the rules which transition us from one state to another.

So for example when we dial a call we move from the Off Hook state to the connecting state.

So how can we actually define all these data structures.

Now unless you're using some extended library you can just define your own map.

So that's exactly what we're going to do.

So we're going to have a map of rules.

So rules is going to be a map where for any given state you're going to define a bunch of different trigger results.

Let me actually make a time for it.

So a trigger result is a combination of a trigger and the state you transition to when that trigger actually happens.

So we're going to have a trigger trigger and state state now we.

So then our rules map is a map from a state to an array of trigger results.

So it's not just a single trigger result because remember from any given state you might it might be possible to transition to more than one state depending on the trigger.

So we want to define all of these.
OK.

So let me just show you an example let's suppose the phone is off the hook the phone is off the hook.

And if you dial a number call dialed then you you are in the connecting state.

So that's an example of a transition.

So whenever you are connecting there can be lots of different things like for example you can you can just hang up.

So if this stay if the transition states that we hung up then the phone is what's on the hook or maybe it's off the hook.

But let's have on the hook if we're done for example.

In addition we can get connected.
So call might be connected.

And in this case we are connected and so on and so forth so as you can see we're filling in this map and so the map goes from state which is this part to an array of pairs effectively a pair consisting of a trigger as well as the state and let's not forget all the commas that are required here.

So I'm just going to add a couple of more definitions to this map.

So here are a couple of more definitions another we have all of this.

We can actually build our state machine and orchestrate this by orchestrate I mean that we run the state machine we we operate the state machine and see see how it goes.

So we're going to have an initial state and also the exit stage the exit state is the state that when you reach the state we are done effectively we're done using the state machine.

So the starting state is going to be off the hook but the terminal state is going to be on hook.

So when you put the phone back on the hook that means that you're done making a phone call and then let's just make a loop.

So for OK equals true Okay equals.
Well okay.

Is the precondition actually and OK equals the OK condition here is that state is not equal to the exit state.

So long as this holds we are still in the game we're still continuing to operate this whenever they become the same whenever the current state becomes the exit state.

That's time for us to leave basically.

So first of all well I'm going to do is I'm going to do a printout of all the possible states I'm print

printout the current state and I'll specify a bunch of triggers that you can fire at this particular point in time.

So first of all we'll print the current state of the phone so the phone is currently and then we print the state.

Now the reason why this works the reason why this part works is because up here we have defined the string your interface for the state.

So that's why I don't have to call it explicitly because of print line.

We just use it right here without any problem.

And then I'm going to say select a trigger and move print out all the triggers available for the particular state.

So for I.
Equal to zero.

I's less than the length of rules given the particular state.

So we're looking at this list of rules list of list of trigger results for a particular state.

And we're saying that well for every single one of them I plus plus what we're going to do is we're going to get the transition so transition is rules at State.

At index i.
And then we'll just print out a numbered list of all of these transitions.

So here I'm going to print line.

Well first of all we'll take the value I and convert it to a string so I too weigh the value of i.

And then I'm going to put a dot here and then I will put the trigger Nazi dot trigger because TR In this case is the transition end result.

So it contains both the trigger as well as the state that you transition to as a consequence of having that trigger in the system.

So now we're going to ask the user what what exactly they want to do.

So inputs comma underscore comma underscore I'm just going to ignore all the extra stuff.

So here we're going to have a new reader from both i o o s dot S2 the N and I'm going to read line.

So I'm going to read a line and then I'll convert that line into a number so I come on this score is equal to y.

So convert the inputs into a string.
Okay.

So now that we have this what we can do is we can take that value of i and we can find of the transition result for that value.

So TR is going to be rules at State at position i and then state is equal to TR that state that we get.

So this changes the state effectively and then we get to run the thing once again assuming that we haven't reached the exit state.

And finally when this is all said and done let's just print line we are done using the phone.

OK.

So now that we've built this state machine that's actually run this let's see if it does in fact work the way we want it to work.

So coming down here so the phone is currently off the hook.

We only have one trigger which is the trigger that the call gets dialed.

So I'm going to put a zero here and we continue.

So now the phone is currently connecting and there are two possibilities here we can hang up the phone or we can assume that the phone did in fact get connected.

So I'm going to choose call connected and then we have a bunch of other trigger.

So we are connected what we can do is we can leave leave a message we can hang up or maybe we get placed on hold.

So if we get placed on hold then we are currently on hold you can see the states being printed here once again thanks to the implementation of the stringer interface and the string method for the different collections for the different constants that we have.

So we're currently on hold.

We get taken off hold for some reason and then let's suppose that for example we leave a message.

So I am I'm going to leave a message here and then we're done.

So we leave a message with on using the phone we transition to the on hook state and the system knows that we are no longer running the state machine.

So this is how you implement a state machine by hand.

And this is how you do it in a more realistic setting so unlike the previous example where we had different struct modeling the different states and having invoke having the location of a method on the state itself which is somewhat weird nothing like that is happening here.

All we're doing is we just have a bunch of constants and then we have a map which basically defines all the transition rules that can happen inside a system.

### Handmade State Machine code: behavioral.state.handmade.go

```go
package state

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

type State int

const (
  OffHook State = iota
  Connecting
  Connected
  OnHold
  OnHook
)

func (s State) String() string {
  switch s {
  case OffHook: return "OffHook"
  case Connecting: return "Connecting"
  case Connected: return "Connected"
  case OnHold: return "OnHold"
  case OnHook: return "OnHook"
  }
  return "Unknown"
}

type Trigger int

const (
  CallDialed Trigger = iota
  HungUp
  CallConnected
  PlacedOnHold
  TakenOffHold
  LeftMessage
)

func (t Trigger) String() string {
  switch t {
  case CallDialed: return "CallDialed"
  case HungUp: return "HungUp"
  case CallConnected: return "CallConnected"
  case PlacedOnHold: return "PlacedOnHold"
  case TakenOffHold: return "TakenOffHold"
  case LeftMessage: return "LeftMessage"
  }
  return "Unknown"
}

type TriggerResult struct {
  Trigger Trigger
  State State
}

var rules = map[State][]TriggerResult {
  OffHook: {
    {CallDialed, Connecting},
  },
  Connecting: {
    {HungUp, OffHook},
    {CallConnected, Connected},
  },
  Connected: {
    {LeftMessage, OnHook},
    {HungUp, OnHook},
    {PlacedOnHold, OnHold},
  },
  OnHold: {
    {TakenOffHold, Connected},
    {HungUp, OnHook},
  },
}

func main() {
  state, exitState := OffHook, OnHook
  for ok := true; ok; ok = state != exitState {
    fmt.Println("The phone is currently", state)
    fmt.Println("Select a trigger:")

    for i := 0; i < len(rules[state]); i++ {
      tr := rules[state][i]
      fmt.Println(strconv.Itoa(i), ".", tr.Trigger)
    }

    input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
    i, _ := strconv.Atoi(string(input))

    tr := rules[state][i]
    state = tr.State
  }
  fmt.Println("We are done using the phone")
}
```

## Switch-Based State Machine

There is yet another type of state machine that I want to show you and this state machine is very special because instead of having a map of the different transitions what happens is you encode that information somewhere else.

And in this particular case what you do is you encoded inside a switch statement.

So we're going to take a look at how that works.
Now first of all let's set up a scenario.

So we're going to try and model a combination lock.

So a combination law consists of basically you have four digits for the lock and somebody makes up the combination and you have to enter the right combination to unlock the thing if you enter the wrong combination then the thing remains locked.

So what we're going to do is we're going to have a state once again as just a bunch of integers and what I'm going to do is I'll have a state locked.

I'll also have the state failed.

That's when you enter an incorrect code and a state unlocked.

If you do in fact entered the correct code so we can model this and you'll notice that I'm not making any kind of maps here because I'm going to specify all the transitions from one state to another within a switch statement.

So basically first of all let's define the code.

So this is the combination code that you have to enter in order to unlock this thing.

Initially the state of the system is locked because you haven't entered the code yet.

And we're also going to collect the numbers that some of the entries as they are manipulating the lock in a string builder.

So entry is going to be a String start builder.

Now we go and from here on out what we're going to do is we're going to have an infinite loop just a four.

And inside that four we'll have a switch so we switch on the state and depending on the state we do different things.

And this is all happening in an infinite loop.

So let's imagine the first situation the most common situation is when the lock is actually locked.

So if it's locked.

What we're going to do is we're going to read a character so effectively we're going to read the run from the standard input and then A will added to the string builder that we made.

And then we'll check whether or not the string builder has the right stuff in it.

So first of all are on the score on the score equals new reader.

So that's again from Buffalo New reader from standard input.

We're going to read runes.

So we're going to read just a single rune or a single character and then we'll take that character and

we'll write it to the strings builder.
So I'll say entry dots right rune.
So we write the rune too.
Ah there we go.
OK.

So now that we have this there are a couple of situations here that we need to handle.

If the right code has been entered fully then the state of our system is unlocked and we simply break from the from the for loop.

So if entry dot string is equal to code then we are done so we said the state unlocked and we break.

And that's probably the simple case another case is that we are still entering the digits.

Now as soon as you enter an incorrect ID what you can do is you can detect that fact and if somebody enters an incorrect digit.

So if the code actually doesn't start with the entry as it is right now that means something is wrong somebody has entered an incorrect code then we have to set the state to failed.

So if strings index so I'm going to use the index function here.

So if the index of entry dots string inside code is not equal to zero.

That means that effectively the code doesn't start with the entry.

So there is a mismatch.

In this case the code is wrong and we set the state to failed.

We don't do any kind of breaks here because we're inside an infinite loop and we know that this which will eventually be executed once again and when it's executed once again then everything will go to the failed case so let's define the failed case.

So the fail case is where you get to prints that somebody has failed because they entered an incorrect code then we reset the entry so they can try again.

So I will just reset the string is builder and the state goes back to locked and we continue executing this because we're still inside an infinite loop.

Now in the case where the state is unlocked what happens here is we print line unlocked and then we return.

So we terminate the program effectively.

So let's actually take a look at how it works.

We're going to do a few incorrect entries and also some correct entries.

So we are running this already and we can put the cursor here and start typing things so if I do 1 2 5 for example 5 is obviously incorrect because our code is 1 2 3 4 I do 1 2 5 and it tells me that I have failed in the whole thing resets if I do 1 2 3 6.

Once again I'm going to get a failed but if I do one two three four then I get unlocked and the process terminates because we do that return.

So this is an example of a different kind of machine in the sense of the way that it's written so we do have the states but we don't specify the transitions explicitly instead what happens is we have an infinite loop inside us with a switch in it and inside that switch we handle every single case and we perform the orchestration of the state machine inside these cases.

Meaning that to transition from one state to another we simply change some variable and that's pretty much it.

And then as this thing runs forever you encounter this variable and one of these variables one when you have to find the terminating condition you simply do a return of something to that effect.

That way you know when the state machine needs to terminate.

### Switch-Based State Machine code: behavioral.state.switchbased.go

```go
package state

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type State int

const (
  Locked State = iota
  Failed
  Unlocked
)

func main() {
  code := "1234"
  state := Locked
  entry := strings.Builder{}

  for {
    switch state {
    case Locked:
      // only reads input when you press Return
      r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
      entry.WriteRune(r)

      if entry.String() == code {
        state = Unlocked
        break
      }

      if strings.Index(code, entry.String()) != 0 {
        // code is wrong
        state = Failed
      }
    case Failed:
      fmt.Println("FAILED")
      entry.Reset()
      state = Locked
    case Unlocked:
      fmt.Println("UNLOCKED")
      return
    }
  }
}
```

## Summary

All right so let's talk about the things that we learned about these state design pattern so given sufficient complexity.

It's actually worth formally defining all the different states and all the different events or triggers that take you from one state to another.

But in addition to that you can define lots of other things that we haven't covered and that make the whole finite state machine construction a bit more complicated.

So for example you can define the behaviors that happen as you enter a particular state or you exit a particular state.

You can also have actions when a particular event causes a transition and that goes to the Observer design pattern obviously you can also have guard conditions which actually enable or disable transitions.

So transitions don't have to be like turned on all the time.

Sometimes that transition is impossible and you want to be able to turn those on or off and you can also have for example default actions when no transitions are found for an event so some event happens but there is no transition you can also handle the situation maybe maybe a raise and there may be panic or maybe do something else depending on your scenario.

So state machines can be symbol but they can also be really complicated at times.

It all depends on the scenario that you're working with.
