# Observer -- I am watching you

## Overview

We're now going to take a look at the observed design pattern.

So what is this pattern all about.

Well sometimes we need to be informed when certain things happen like for example let's say we have an object and the object has a field.

We want to maybe be informed whenever these field changes value when the value of the field changes from one to another or maybe when the object does something or maybe when there is some external event that occurs and we want our system to be able to handle it somehow.

Now there are different ways of actually handling these situations.

So for example suppose you want to detect that an object's field changes.

Now what you can do is you can start polling you can check the objects field every 100 milliseconds but obviously that's not practical.

Realistically instead of doing that what you want is you want the object to tell you when one of its fields actually change.

So we want to be able to generate events and then what we want to do is we want to listen to these events and be notified when they actually occur.

Now there are two persistence in this story.

Something called an observable and something called an observer.

So the observer design pattern involves both of these.

So we have something called an observer which is the object that wishes to be informed about events happening in the system and then we have somebody who is actually generating those events and that thing is typically called an observable.

## Observer and Observable

The whole idea behind the observer design pattern is that you have some components which generate certain events and other components which want to be notified about these events happening.

So essentially some component does something or some information comes into the system and you want notifications in other components.

So the question is how can you do this and furthermore how can you do this in the most general way possible.

Well let's take a look at how we can implement a general purpose kind of framework generally just a bunch of structure and interfaces to support this idea of notification.

So there's gonna be two participants to this story one participant is typically called observable and the other is called the observer.

So let me give you an example.

Let's suppose that we have a situation where you have a person that becomes ill.

So when a person becomes ill they send a notification that they are ill and then somebody let's say a doctor can get this notification and can maybe come visit the ill patient.

So here the observable would be the patient and the observer the one that's monitoring for changes is the doctor.

So there's always these two counterparts these the observable and the observer.

Now we're going to implement these differently because for the observable I'm going to make it a struct.

So I'm going to have a drug called observable and an observable is basically going to keep a list of subscriptions is basically going to keep a list of observers which are connected to it which are subscribed to the observable as events.

So we're going to have subs which is going to be a list as just have it as a list.

You can use some other data structure if you want but the list works just fine in this particular case.

So what do you want the observable to be able to do what you want to be able to subscribe to the events that happen on the observable.

So that's the first thing that we'll do so the observable is going to have a method called subscribe.

And here is the question What do you provide into this method.

You have to provide some observers some component which actually wants the notifications.

So here I'll call it just x.

Unfortunately I cannot call it oh because we're already using the variable here.

So we're going to have an observer.

Now before I finish this let's actually implement the observer so in my implementation the observer is going to be an interface not a struct just an interface.

So observe is going to be an interface.

And the idea here is that the observer has a method which gets notified effectively this method gets called whenever there is something happening so when something happens the observer has some method let's call it notify.

So the observer gets notified that something happened.

Quite often when you have the observer pattern you also want to send some data about what exactly happened.

You want to send some service data like for example if a person's state changes you might want to encode that state as a string and send it as as a kind of message to the observer.

So here we'll have some data and I didn't know what kind of data this is so I'll just define it as an interface.

So I have no idea if it's a string or some composite struct or something like that.

We don't know.

So now it makes sense for us to come back to subscribe and this is where we're using that observer interface.

So we simply added to the list.

So to subscribe you simply take the subscriptions.

So Oh Dot subs and I should correct that curly brace here.

Oh Dot subs dot push back and I add this observer to the set of subscriptions and similarly we can have unsubscribe so somewhat more complicated though so we'll have unsubscribe.

That's the method you would use when some observer is no longer interested about being notified of something happening so here will once again take the observer not the observable obviously.

And here we would go through every single subscription and if the subscription matches the observer then we remove that subscription.

So for some Z starting at 0 adult subs stop frauds.

So we'll start at the front end while Z is not equal to nil.

Z equals Z next.
So here is the basic iterator for you.

So here we say that if z value cast to an observer is in fact equal to the Observer we're getting which is X.

This X right here then we remove it we say 0 adult subs don't remove z so we remove this observer from the set of observers because sometimes you want to be notified about some event happening but only to a point after which you no longer really care and you don't want any notifications to happen.

Okay.
So that's another method.

And the final method that you would have on the observable is some method for actually firing the event some method of notifying the Observer that something happens.

It's once again that's really up to you how you want to call it.

But I typically call it fire so you fire some data off and every single subscriber gets this data.

So once again this data we don't really know what kind of data it is.

So I don't specify explicitly here that it's like an end or something else.

And once again we go through every single subscription.

So for Z starting at 0 dot subs dot France Z not equal to nil.

Z equals Z next.
And what we do here is we take the value.

So we take Z the value then we cast it to an observer because remember they are store there's just interface nothing no information there but we know are observers and we call notify on them.

So we notify and we pass in the data so that data the data that is being sent from this location is going to be accepted here.

And here we can process it somehow.
OK.

So we've set everything up but now what we can do is we can actually use the observer interface as well as the observable struct to build some sort of a scenario.

So let's suppose that we do in fact have a person that maybe the person catches a cold and we just want to have some sort of doctor's service to be informed that a person has become ill and maybe it's time to visit them to see what's going on.

So we'll have type person struck.

Now this is going to include everything from observable but in addition will have the person's name as a string.

So I'll make a factory function here initializing the name.

Now what I also want to do is I want to initialize not just the name I also want to initialize that observable and here is going to be.

And while it may be a bit tedious because let me just fill in all the fields.

So here are the observable would also have to be initialized with a new list.

So here you would make a new list the list for the subscriptions and of course you would specify the name you would take the name from the argument that gets passed into the factory function.

So this is how you would initialize a person and then we can simulate the person catching a cold.

So here we have a function method of person called Catch a cold.

So this is where you want to perform the notification.

So what you do is you use the information that you have inside the observable.

Remember we are including everything from observable.

So you use the observable to basically fire off the events so you can take P the person and you can say pedo fire.

And here you can provide whatever information about the event you want to pass like for example you want to tell the doctor the name of the patient.

So here you can pass in pedo name like so.
OK.
So this is our patient.

This is our person and now we want some sort of Doctor Service.

I'm not going to add any members here we'll just have Doctor Service has a struct so it's not going to store any data but we are going to implement the the observer interface I'm going to go ahead here and implement methods from the observer interface as you can see we're implementing they notify method and this is the method that gets called whenever a person falls ill because it will connect the two together in just a moment.

So here I will simply print off certain things so I will say that a doctor has been called 4 percent s and that would be the name of the patient.

And here to get the name of the patient we have data but data has to be cast to a string because we don't really know what type it is.

So you would cast it to a string like so.
Okay.

So now that we have this entire scenario let's actually take a look at how to use it.

So first of all I'd make a person a new person that's called him Boris for example and then let's make the Doctor Service.

So Doctor Service like so there's no data to see into the doctor's service and then what we do is we connect the two together.

So we take the person and we subscribe so we get to the doctor's service to subscribe to events which happen on the person.

So now if something happens to the person the doctor service gets notified so we can take the person and we could call the method catch a cold.

Which should inform the doctor service that a doctor has been called.

So let's take a look at how this works.

Let's just run the whole thing and as you can see it's working just fine.

So as soon as we get ill a doctor gets called for this particular person.

As you can imagine you can have more than one observer for a given observable So for example we can have multiple doctor's offices being attached to a person.

And similarly you could have.

You could have different handles for different objects so it's a very flexible framework.

And this is the gist of the observed design pattern as it's implemented in the Go programming language.

### Observer and Observable code: behavioral.observer.observer.go

```go
package observer

import (
  "container/list"
  "fmt"
)

type Observable struct {
  subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
  o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    if z.Value.(Observer) == x {
      o.subs.Remove(z)
    }
  }
}

func (o *Observable) Fire(data interface{}) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    z.Value.(Observer).Notify(data)
  }
}

type Observer interface {
  Notify(data interface{})
}

// whenever a person catches a cold,
// a doctor must be called
type Person struct {
  Observable
  Name string
}

func NewPerson(name string) *Person {
  return &Person {
    Observable: Observable{new(list.List)},
    Name: name,
  }
}

func (p *Person) CatchACold() {
  p.Fire(p.Name)
}

type DoctorService struct {}

func (d *DoctorService) Notify(data interface{}) {
  fmt.Printf("A doctor has been called for %s",
    data.(string))
}

func main() {
  p := NewPerson("Boris")
  ds := &DoctorService{}
  p.Subscribe(ds)

  // let's test it!
  p.CatchACold()
}
```

## Property Observers

So one very common use of the observer design pattern is to implement notifications related to property changes on a object.

So in this case a struct might change somehow and you want to be informed that there is in fact a change.

So let me show you how this can work.
I'll get rid of the demo that we currently have.

I'll try to reuse person as much as I can but we're going to make such modifications.

So here instead of specifying the person's name let's suppose that we want to be informed when a person's age changes.

So we'll have h here as an integer I will once again make a factory function which initialize the persons age and also initialize as the actual observable so here.

Let's initialize the observable to a new list.

So when you list the list and you could have a factory function for the observable as well if you wanted to that would suddenly simplify things let's set the age to the age parameter right here.

So this is how you would make a person.

Now imagine that you want to get notifications about a person getting older or indeed getting younger if they find some way of reversing the aging process.

So how does this whole thing work.

Well you have to send information about the person's property being changed now what do I mean by property.

You'll notice that I made h here a lowercase the idea of properties is that a property is a combination of a getter in the center.

So for the getter you would have a method called age and for the setter you would have a method called set H.

Now properties are not particularly idiomatic to go unless of course you want to do the kind of thing that we're doing here we just change notifications because the idea is that when you change the age when you call set age.

Not only do you want to change the actual value that would be this value but you also want to use this as an observable and get the observable firing some sort of event telling others that the age has changed and age might only be one of a set of properties that you want to notify about.

So when we have a change notification when we want to tell somebody that something has changed we can have a separate struct for encoding this information so we can have a type called let's say property change property change.

So a property change is going to include two things is going to have the name of the property.

So for example if you change the H property it would have the string age for if you change let's say the Hyde property it would have the string height.

It's up to you whether to make it upper case a lowercase doesn't really matter here.

Okay.

And then the other field that will have as part of property changed is going to be the new value.

The value of the element as it's been changed.

So when you change the age let's say from 10 to 11 named would have the value of age being a string and then we would have the actual value.

So once again the value depends on the type of the field.

So all I can do here is define it.
Like so there is no other way of doing it.

So now that we have this property change struct what we can do is we can use this struct to actually perform notifications so to get the getter in the centre for the age I would type it as follows.

So first of all let's have the getter.

So when somebody wants the age we just return Pete out age.

Small a here.

You'll notice there is nothing to do here for the getter there is no additional things to do it's the setter that becomes interesting when you do the change notifications because here we would have a method called set age where you specify the ages and end but you only want to perform this change if the value of age that's being set is actually different from the current value.

So we say that if age is already equal to out age there is really nothing for us to do we don't want to generate a notification for that.

So we simply return from the method altogether.

Then of course we perform the modification we say Pitot H equals age.

So this is good.
And then we want to send off this event.

We basically want to notify every single observer that this observable has changed so we say Pete fire and we need to provide this data.

Now remember we made this special property change struct.

So that's what we're going to be using we're going to be using the property change struct the first parameter here is going to be the name of the property that has in fact changed.

And the second is going to be its value which in our case is Pitot age which is an end.

So this is how you actually notify somebody.
Now let me show you how you can use this.

Let's suppose that we have some let's say we have a traffic management company or system or whatever.

Now this traffic management wants to be informed about a person's age and if the person is too young to drive for example they would they would keep monitoring their age but as soon as a person turns let's say 16.

The traffic management let me fix it here.

The traffic management no longer cares about the person's age.

The traffic management congratulates the person on being able to drive and then they unsubscribe from the observable so how would you implement this.

Well first of all traffic management has to have some sort of pointer to the observable.

So here I would have the observable that we are actually interested in and then of course the the real sort of interesting part is what happens when traffic management gets notified.

So here we go ahead and we implement the observer interface like so and in the notify method we do the following.

So first of all we have a data which is defined as an interface.

Now we know for a fact that this interface is actually going to be property change.

But let's just be on the safe side.

Let's try a cast.

So if P.C. come up OK is equal to data being cast to property changed.

And if we do in fact get an OK.

Then obviously the cast succeeded and we can use.

P.S. The P.C. variable right here to unpack information about the property so we only care when the person's age is greater than or equal to 16.

So we say if P.C. that value cast to an end is greater than or equal to 16 then we can tell the person that they are now old enough to drive so we can say congrats.

You can drive now and then we no longer care about this person because now they are old enough to drive.

We no longer need to monitor them at least an hour scenario.

So what we do is we take the traffic management pointer we try to grab the observable which is oh I'm talking about this observable right here we grab the observable and then only observable we call unsubscribe passing in the original object so teed up oh the unsubscribe t I know it looks a bit weird but it's it's the correct way of on subscribing so we grab the observable and on the observable we unsubscribe the current object the object on in whose method we are actually operating right now.

So we no longer care about the person getting older and now what we can do is we can try this whole scenario out.

So here I'll make a person which is originally let's say they're 50 so they're 15 will make traffic management which is going to be just traffic management which monitors it monitors the person.

But here is another interesting thing the variable p is a person.

But remember traffic management if you look at it it's actually interested in the observable so we cannot just pass P here.

We have to say Pete out observable we grab the observable part of that person and we pass that in there and then what we do is we perform the actual subscription so we say Pete out subscribe t so this is how you subscribe one element to another.

And then let's try cycling the age for four for example from 16 to 20.

So for IE is equal to 16 I is less than or equal to 20 I plus plus.

So let's just say that we are setting the age to the value ie.

And then we try to set the age piece at age I.

All right let's run this let's take a look at what we get.

All right.

So as you can see that as soon as we set the heat to 16 we get that congratulations and notice that this congratulation message doesn't repeat itself it doesn't repeat itself at 17 at 18 and so on because we are on subscribing here.

If I were to prevent this if I were to prevent on subscription we would just get it on every single age after 16.

As you can see here we're getting it all over the place.

And also let's try changing it to some other value.

Let's say I change it to 18 for example in some countries you can only drive from age 18.

So in this case you would said the H2 16 17 18 and then you would get the congratulation message.

So everything is working correctly here.

So hopefully I've been able to demonstrate that there is a case for using properties meaning a combination of a getter and a setter as opposed to just ordinary fields.

That case is when you want notifications of property changes.

So here as soon as somebody tries to change the age there is a notification there is an event that's being fired and you can subscribe to this event and you can get this event and process this in one way or another.

However this approach does have certain problems with dependencies and that's what we're going to take a look at in the next lecture.

### Property Observers code: behavioral.observer.propertychanges.go

```go
package observer

import (
  "container/list"
  "fmt"
)

type Observable struct {
  subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
  o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    if z.Value.(Observer) == x {
      o.subs.Remove(z)
    }
  }
}

func (o *Observable) Fire(data interface{}) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    z.Value.(Observer).Notify(data)
  }
}

type Observer interface {
  Notify(data interface{})
}

type Person struct {
  Observable
  age int
}

func NewPerson(age int) *Person {
  return &Person{Observable{new(list.List)}, age}
}

type PropertyChanged struct {
  Name string
  Value interface{}
}

func (p *Person) Age() int { return p.age }
func (p *Person) SetAge(age int) {
  if age == p.age { return } // no change
  p.age = age
  p.Fire(PropertyChanged{"Age", p.age})
}

type TrafficManagement struct {
  o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
  if pc, ok := data.(PropertyChanged); ok {
    if pc.Value.(int) >= 16 {
      fmt.Println("Congrats, you can drive now!")
      // we no longer care
      t.o.Unsubscribe(t)
    }
  }
}

func main() {
  p := NewPerson(15)
  t := &TrafficManagement{p.Observable}
  p.Subscribe(t)

  for i := 16; i <= 20; i++ {
    fmt.Println("Setting age to", i)
    p.SetAge(i)
  }
}
```

## Property Dependencies

In the previous lecture I told you that there are certain problems with using change notifications on properties by property I mean a combination of a getter in the center.

So what are those problems.
Well let me show you.

I'm going to show only one type of problem.

There might be other problems as well.

So let's imagine that for a given person we decide to make a read only property.

So a property which is computed from some other value as opposed to just being used directly.

So for example let's suppose we want a boolean method telling us whether a person can vote.

So we'll have a function P person.

So we'll have a method called count votes which is returns of boolean value here and we return whether or not P that age is greater than or equal to 18 for example.

So what is the problem with having things like this.

Well the problem is that you want to change notification for this as well.

It would be really nice.

But remember change notification can only happen in setters.

It cannot really happen in getters right here.

So the question is where exactly would you send the notification for the change in the voting status.

The answer of course is that you would do this in the center for the H because we are changing the age right here and age affects the result of Ken Vote.

It makes sense that not only do we set the notification for age but right here somewhere we also set the notification for can vote.

But first of all let me actually complete the scenario so you can see what's going on.

Let's imagine once again that we have some sort of structure for an electoral roll.

That's gonna be just just an empty struct like so now what I want to be able to do is to once again implement the observer interface so let's implement observer on the electoral roll.

And let's suppose that we want to notify a person or to notify the system when a person's voting status has changed to true when they can finally vote.

So once again we perform the same casting stuff as before we're going to assume here that we are reusing the same property change struck that we have up here so we'll keep reusing the struct.

We'll say if P.S. come out OK.

Is equal to data being cast to a property change.

And if everything is OK then what we can say is if the property that we are interested in is called

Can the vote and the Boolean value is true so if P.S. the value dot cast it to a Boolean if this is true then we can congratulate the person that they can actually vote.

So here we can say congratulations.
You can vote like so.
So this is our handler.

This is the handler for the changes in a person's can vote but the question is where exactly would we generate a UN event where would we fire an event where the argument here is can vote and the answer is we would have to do this inset age.

There is no other place for us to do it but in order to be able to do this we need to make sure that can vote has actually changed that the voting status changed from false to true or from true to false because otherwise we don't want to fire any change notification.

And this is where things get really difficult because remember the changing of the age changes the result off can vote.

And if we want to make sure that can vote has in fact changed.

We need to cache its previous value and then compare it to the new value.

Annoying I know.

So here you would catch the previous value old can vote.

He can vote like so and after setting the age you would also check whether or not the value has changed.

So if old can vote is not equal to PDA can vote then the voting status has in fact changed.

And here as well we would say PDA fire and we would fire the property change event with can vote as the name of the property as well as out can vote as the new value.

So this is how you would implement the dependent property because essentially what we have is we have can vote which is a property which depends on the property or indeed the age field.

You can treat it one way or another it doesn't really matter.

So going back to our scenario let's connect everything together.

So we have the person already.
May as well put a zero in here.

We'll change it later on we'll have the electoral roll.

So now we do the subscription as before you don't subscribe get the electoral roll subscribing to the events that happen on the person.

And then once again let's do a loop starting at the value of 10 going to the value of 19 let's say I plus plus and here first of all let's inform that we are setting the age to a particular value ie.

And then we call Pitot said age and actually perform the modification right here.

OK so let's run this let's just see the scenario working.

So as you can see everything is working is expect so we're changing the age to 10 Levin and all the way up to 18.

And finally when it's 18 and what happens is the voting status changes and we get the congratulations message.

OK.
So how does everything work.

Right here what exactly is the problem here.

Well the problem here is the problem of dependencies.

The problem is that the voting status depends on age and age gets modified as part of the set age set or so the set age setter becomes very large IT STARTS caching previous values of all the properties it affects and then compares the previous values of the affected properties to the current values and sends the notifications here.

As you can probably imagine this situation doesn't really scale.

If you have lots of properties dependent upon the age property or if you have one property such as can vote that depends on let's say five other different properties then you end up with a nightmare you end up with a very complicated scenario.

And typically we don't do things this way.

We don't define dependency properties inside the setters of the properties.

Instead we try to build some sort of higher level framework some sort of map where all the dependencies between all the different properties are catalogued and then subsequently you iterate through this map and you perform the notifications in a more regularized way.

But this example I've shown right here is just a very simple illustration of how you can get complexity virtually out of nowhere and suddenly there is no go language mechanism that would help us with this.

So if you want to have change notifications on dependent properties you would have to be building some sort of complex infrastructure of your own.

I'm afraid.

### Property Dependencies code: behavioral.observer.propertydependencies.go

```go
package observer

import (
  "container/list"
  "fmt"
)

type Observable struct {
  subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
  o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    if z.Value.(Observer) == x {
      o.subs.Remove(z)
    }
  }
}

func (o *Observable) Fire(data interface{}) {
  for z := o.subs.Front(); z != nil; z = z.Next() {
    z.Value.(Observer).Notify(data)
  }
}

type Observer interface {
  Notify(data interface{})
}

type Person struct {
  Observable
  age int
}

func NewPerson(age int) *Person {
  return &Person{Observable{new(list.List)}, age}
}

type PropertyChanged struct {
  Name string
  Value interface{}
}

func (p *Person) Age() int { return p.age }
func (p *Person) SetAge(age int) {
  if age == p.age { return } // no change

  oldCanVote := p.CanVote()

  p.age = age
  p.Fire(PropertyChanged{"Age", p.age})

  if oldCanVote != p.CanVote() {
    p.Fire(PropertyChanged{"CanVote", p.CanVote()})
  }
}

func (p *Person) CanVote() bool {
  return p.age >= 18
}

type ElectrocalRoll struct {
}

func (e *ElectrocalRoll) Notify(data interface{}) {
  if pc, ok := data.(PropertyChanged); ok {
    if pc.Name == "CanVote" && pc.Value.(bool) {
      fmt.Println("Congratulations, you can vote!")
    }
  }
}

func main() {
  p := NewPerson(0)
  er := &ElectrocalRoll{}
  p.Subscribe(er)

  for i := 10; i < 20; i++ {
    fmt.Println("Setting age to", i)
    p.SetAge(i)
  }
}
```

## Summary

All right.

So let's summarize the things that we learnt about the observer design pattern.

So the observer is by definition an intrusive approach which means that to make an object observable you have to change that object.

Unfortunately in that why do you need to do this while you need this because you have to have a way for clients to subscribe to the events generated by a particular object and then you take the event data and you send it from the observable to all the observers.

Now in our case since we try to make things as generic as possible we use dynamic typing we use that interface definition instead of saying that our events specify a specific kind of data structure for the event and said we let the different components define what kind of data they send.

And of course you also have to have additional functionality like for example on subscriptions so sometimes you want to listen to events only up to a point and then you're no longer interested.

So you want to cut that connection as well.

And we've looked at how to implement all of this.
