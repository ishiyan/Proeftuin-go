# Chain of Responsibility -- Sequence of handlers processing an event one after another

## Overview

The chain of responsibility is the kind of design pattern that is mirrored quite a lot.

In the real world.

So let's imagine a situation just a hypothetical situation.

Let's imagine that you have some unethical behavior by an employee in the company.

Like for example maybe there is a company that trades on the markets and there is some insider trading the question is who actually takes the blame.

So an employee gets caught for insider trading but who actually takes the blame.

So one approach and typical approach taken by most companies is that they just blame the employee so they blame the employee they fire them they make them the fall guy or the fall girl.

And that's pretty much it.

But if the manager of that employee knew about the insider trading and did nothing then it could be said that the manager also bears some responsibility.

And of course if the practice of insider trading was institutional then you would typically reach all the way up to the CEO and blame the CEO for basically setting up a company which violates the law.

So another example of the chain of responsibility is more to do with software like let's say you click a graphical element on a form like a button for example.

So the question is who gets to handle the event.

So one approach is that when you click a button the button actually handles it and no further processing is done.

So you detect a button press and that's pretty much it.

And in actual fact that's how most systems operate.

But let's suppose the button itself does nothing there is also the underlying Group box.

So the button is inside a group box and the group box also gets a notification that somebody's collected and maybe the group box I don't know turns a different color or something.

And it can go all the way up to the underlying window or the underlying web page if we're talking about web development.

So here you could once again handle that event that click event and also do something in relation to that.

So you have different objects who get to handle a particular event like like a click for example or one after another.

So another example is like if we talk about card games which are nowadays that they used to be just paper but now you get these collectible card games which are on the computer in those cards you typically have some sort of creatures fighting each other who typically have attack and defense values but then again you can boost these creatures by other cards and the question is how do you apply those rules.

How do you apply all of that all of those modifications and so we come to the chain responsibility design pattern.

So basically the idea is that you've got a chain of components who all get a chance to process some command or query or event whatever it happens to be and they can process it one after another but they also have the option of having some default processing implementations and they also have an ability to terminate the entire processing chain so a particular element of the chain of responsibility has an ability to prevent subsequent elements from actually processing a particular message.

## Method Chain

So we're going to take a look at the chain of responsibility design pattern.

And we're going to consider a same scenario into different implementations.

So the first of these implementations is what I typically call a method chain because well there's gonna be a linked list of method invocations.

So how does all of this work.

Well let's imagine you're working on some sort of game.

So the game has a bunch of creatures and these creatures are participating in the game so a creature might have a name for example and the creature might also have attack and defense value so that it can engage in combat with other creatures.

So we can ad attack and defense integer attributes here and then we can even implement to the string your interface.

So let's do that.

So this trigger and like so so for the creature I'm just going to print half let's print some information about the creature so we'll print the name followed by the attack slash defense values like so.

So there's going to be sealed up names.
You don't attack and see the defense.

And we'll also have a constructor so let's generate the constructor which basically just initialize is every single thing.

So here we go.
OK.

So with this setup what we want to be able to do in our game is we want to be able to apply modifiers to a creature.

So for example the creature might be a roaming the grounds and the creature picks up a magic sword.

So the creature tries to eat a mushroom and gets poison.

We want to modify the aspects of the creature and we want to build a stack or a list of these modifications

so that they can be applied to a creature one after another.

So what we really want is we want to be able to have some sort of modifier so modifier is going to be an interface and a modifier is going to do two things.

So first of all we'll have a method called add which adds an additional modifier to the chain of responsibility.

So the idea is that you can apply more than one modify to a creature and whenever you apply additional modifiers you stick them on the end of an already existing modifier.

And the other part is that we want to actually apply the modifier.

So we'll have some sort of handle method you can also call it apply or something else up to you really and that is where the modifier actually gets applied.

So now that we have the interface we also need some sort of concrete type.

So we'll have creature modifier as a struct in this truck is going to specify what creature the modification is applied to.

So we'll have a creature creature so that is going to be a pointer.

And in addition we'll have a linked list of these modifiers.

So this modifier will have a pointer to the next modifier.

So we'll have a pointer to the next modifier thereby making this a effectively a singly linked list.

So now let's go ahead and let's implement the modifier interface on creature modifier so modifier high interface like so so we're going to implement this on point the receiver is now and when it comes to adding an additional element to this linked list you want to add it to the end.

So you want to say that if the next element in the list is not equal to nil then you obviously need to apply it to the next element.

And this is a recursive operation that will find the end of the list and add that modifier to the very end.

So if we do have a next element then we say see the next dot add.

And so we add the modify recursively otherwise what we do is we apply it here.

So we say see the next equals M.

That's pretty much all there is to it really.

OK.

So this is the set up for adding elements to the linked list.

And the other part is handling the application of those creature modifiers.

So the idea is that if you have some subsequent modifier then you want to apply that modifier.

But because this is a struct which doesn't specify explicitly how we're changing the creature there isn't much that we can do here.

I mean by default a creature modifier doesn't do anything and just gives instructions to whoever is aggregating it that you want to do something about it.

So here all we have to do is we check whether or not there is a next element in the list and if it's not nil then we also call handle on that element.

However you'll notice that at the moment handle doesn't really do anything.

So you call it on next and it calls handle but handle itself doesn't really do anything he just calls every single element one after another.

So you might be wondering well why are we doing this.

And the answer is that we're going to aggregate this type and we're going to actually make use of all of this stuff so now that we have a creature modified that's actually make a constructive for it where we initialize it with a creature.

So this is how it's going to go we don't specify the next modify or we don't specify this value because that can be added later on.

So we have the creature modify but we want something concrete.

We want a modifier which actually does something to a creature.

So for example let's suppose there is an item that the creature can pick up which doubles the creatures attack.

So we can have a double attack modifier.

There's going to be a struct and what we can do here is we can simply aggregate creature modify there is really nothing else to add here there's no additional storage we know what we're doing we're doubling the attack value but now we can have a constructor which actually takes a pointer to the creature and returns a modifier where this part is fully initialized.

So that's right.

This we're going to have a function called new double attack modifier where you specify the creature and it returns a double attack modify a pointer.

And here we return a double attack modifier where you specify the creature modifier and inside you would define the creature a C and you don't define next.

So we'd get rid of this completely and we just leave it like this.

Okay.

So once again what's happening here is we're initializing the aggregated part of the double attack modifiers so we initialize the creature modifier with the creature pointer and we stick it in here so we have some sort of initialization of this part of the double attack modifier of course there's really not much else around the double attack modifier yet but we are going to add additional functionality because what we need to be able to do is we need to be able to handle the application of this modifier and to handle it you have to you have to basically implement the the handle method as defined up here.

So that's what we're going to do.

OK so we're going to have a function which takes double attack modifier so it's gonna be a method of the double attack modifier is going to be called handle.

OK.

So I'm going to have a bit of diagnostic output here.

I'm just going to print line the fact that we're doubling the decrease her name and then we're doubling this creature's attack like so.

OK.

So we take the creature we take its attack value and we multiply assign it by two.

And then here is the critical part.

So you might be thinking that we're done that there is nothing else for us to do.

But there is one critical thing for us to do.

If you go up to a creature modify and you look at the handle method you'll see that the handle method is not empty.

The handle method actually has this implementation that it goes to the next element and calls the next element handle function.

So that's what we need to do right here so we need to call the handle but not the handle of the double attack on the fire but rather the handle that we got from the creature modifier.

So the way this is done is we say the DART creature modifier door handle and that way we call that method and we propagate the application of every single one of these modifies in the hierarchy because after we apply this modifier it's gone and called handle and handle it and so on is going to look at the next element in the linked list and if it's not nil then it's going to call handle on that which in turn can call this again.

And so on and so forth to infinity.

So let's take a look at how all of this can work.

So here I'm going to make a goblin goblin is gonna be a new creature.

So it's gonna be a new creature called goblin.

It's going to have one attack and one defense and then let me actually print it out.

That's print print line.

Goblin string and then what we're going to do is we're going to create a route modify.

Now remember if you're working with a linked list that list has to start somewhere.

So we'll make a route which is going to be a new creature modifier applied to this goblin that we made.

Now what we can do is we can call route dot ad and we can add additional modifiers to the goblin.

So for example I can double its attack.

So here I can call a double attack modifier.

Or rather we have a constructor so let's do that new double attack modifier where once again I can specify the goblin for example.

Now eventually what we want to be able to do at some point is to apply this list of modifiers to the creature.

And this is why you say routes dot handle and remember handler was a special method is a method that goes through every single element and actually checks whether there is a subsequent element as well so it's a traversal of the linked list where every single modifiers apply to the creature now.

And then after all is said and done we can print line the creature.

So once again goblin dot string that will give us hopefully modified value.

So that's actually run this let's take a look at what we get here.

So you can see that I have 1 1 goblin and then I'm doubling the goblins attack and I'm getting into one goblin and we can try applying the modifier twice we can replicate this line once again and we can go ahead and double the goblins attack twice.

So after doubling it twice we get a goblin with an attack value of four.

So this is how you can apply different modifiers.

Let me show you an example of yet another modifier.

Let's suppose you want to have some sort of increased defense modifier.

So let me just drop it in here so an increased defense modifier also aggregates creature modifier but it just has different behaviors.

So when you call handle on it you check the creature's attack value and if it's less than or equal to 2 then that defense is increased by 1.

So we can stick it here in the middle just to see that it works I can say root it out and I can make a new increased defense modifier once again on the goblin.

And let's actually run this let's see what we get.

Okay so as you can see we're doubling the goblins attack and then we're also increasing goblins defense which is why we have a number two here.

Now let's imagine that at some point the goblin gets hit with a spell and that spell disables any of the other modifiers.

How can you actually disable this entire list of modifiers.

Let's suppose I want to stick something here I want to say root dot ad and I want to add some sort of modifier which basically cancels everything else.

This is also possible we can make a kind of no bonus as modifier if you will let me show you how this works.

So we can have a type.

No bonuses modifier struct which also aggregates creature modifier and we can have a constructor so func.

Well you know bonuses modifier which takes a creature and it returns a no bonuses modify a pointer and we return simply no bonus as modifier creature modifier fill in the fields we get rid of next we get rid of this thing here and we add the creature.

Okay.
So that's it.

This is how you manufactures a modifier which doesn't allow any bonuses.

But of course one thing that's missing is the handle method.

So let's implement that so func and no bonuses modify a handle.

Okay.

Now what I'm going to do here is I'm going to leave this handle empty.

Now you'll notice that whenever we actually perform the modifications like when we do here we always call the underlying creature modify handle.

This allows us to propagate the entire list of the chain of responsibility we do it here and simply we do it here in the increased defense modify.

But what if I forget to do this.

What if in the no bonus is modifier I decided I'm not going to propagate this list.

That means that every single modify which comes up to this modifier is actually not going to be applied which is what we wanted in the first place because remember we wanted the no bonuses modifier to disable any other modifiers.

So let me show you how this can work.

I can say you know bonuses modifier on the goblin.

And if I now run this then even though I have plenty of modifications here we start with a 1 1 goblin and we end with a 1 1 goblin because essentially this modifier has prevented the traversal of the entire length list thereby preventing the application of any other modifier.

So this implementation of the chain responsibility is called a method chain really is called that because you are following a linked list of these modifiers and you're calling this magical handle method on every single one of them.

And you can also prevent the invocation of this handle method if you decide that you don't want to invoke it from whatever implementer you have constructed.

### Method Chain code: behavioral.chainofresponsibility.methodchain.go

```go
package chainofresponsibility

import "fmt"

type Creature struct {
  Name string
  Attack, Defense int
}

func (c *Creature) String() string {
  return fmt.Sprintf("%s (%d/%d)",
    c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack int, defense int) *Creature {
  return &Creature{Name: name, Attack: attack, Defense: defense}
}

type Modifier interface {
  Add(m Modifier)
  Handle()
}

type CreatureModifier struct {
  creature *Creature
  next Modifier // singly linked list
}

func (c *CreatureModifier) Add(m Modifier) {
  if c.next != nil {
    c.next.Add(m)
  } else { c.next = m }
}

func (c *CreatureModifier) Handle() {
  if c.next != nil {
    c.next.Handle()
  }
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
  return &CreatureModifier{creature: creature}
}

type DoubleAttackModifier struct {
  CreatureModifier
}

func NewDoubleAttackModifier(
  c *Creature) *DoubleAttackModifier {
  return &DoubleAttackModifier{CreatureModifier{
    creature: c }}
}

type IncreasedDefenseModifier struct {
  CreatureModifier
}

func NewIncreasedDefenseModifier(
  c *Creature) *IncreasedDefenseModifier {
  return &IncreasedDefenseModifier{CreatureModifier{
    creature: c }}
}

func (i *IncreasedDefenseModifier) Handle() {
  if i.creature.Attack <= 2 {
    fmt.Println("Increasing",
      i.creature.Name, "\b's defense")
    i.creature.Defense++
  }
  i.CreatureModifier.Handle()
}

func (d *DoubleAttackModifier) Handle() {
  fmt.Println("Doubling", d.creature.Name,
    "attack...")
  d.creature.Attack *= 2
  d.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
  CreatureModifier
}

func NewNoBonusesModifier(
  c *Creature) *NoBonusesModifier {
  return &NoBonusesModifier{CreatureModifier{
    creature: c }}
}

func (n *NoBonusesModifier) Handle() {
  // nothing here!
}

func main() {
  goblin := NewCreature("Goblin", 1, 1)
  fmt.Println(goblin.String())

  root := NewCreatureModifier(goblin)

  //root.Add(NewNoBonusesModifier(goblin))

  root.Add(NewDoubleAttackModifier(goblin))
  root.Add(NewIncreasedDefenseModifier(goblin))
  root.Add(NewDoubleAttackModifier(goblin))



  // eventually process the entire chain
  root.Handle()
  fmt.Println(goblin.String())
}
```

## Command Query Separation

So let's talk about something called Command query separation that's an important topic when discussing the whole business of commands and queries and chain chains of those as well.

So we have this idea of a command.

We talked about this as part of the command design pattern.

So basically that's when you wrap some action as a separate structure and you pass that in and this gives you lots of different abilities.

So for example you can specify an action like please set your attack value to two and wrap it inside a command.

And we also have queries so queries is when you're asking for information instead of telling a component to do something or to change itself.

You are asking for information.

So a command is a mutable operation an operation that will change state whereas a query is typically an immutable operation an operation that does not change state like you're asking for.

Let's say some attack value.

So the idea of command query separation is when you have separate means of sending commands and queries.

So you have let's say you have one event buzz for processing commands in another event.

Buzz for sending the queries or you might be sharing a buzz but you're still separating those two ideas.

So you have commands and commands behave in a particular way.

They perform actual modifications for example and then you have queries which simply return a bunch of values.

## Broker Chain

We are now going to take a look at a much more advanced implementation of the chain of responsibility design pattern.

Now this design pattern is going to combine multiple approaches.

It's going to combine multiple patterns.

So we're that going to have the chain of responsibility design pattern.

But we also have the mediator pattern by having a centralized component that everyone talks to will have the observe a design pattern.

Because this mediator is going to be observable and well there's gonna be a lots of other enterprise approaches shall we say including the idea of command query separation.

Now I'm throwing everything in the kitchen sink into this demo so I'll try to go slowly and explain what's going on.

So we are replicating the example that we had in the previous lesson.

The example of the method chain.

But this time round things are going to be different this time round of the creature is going to also have a name attack and defence values.

But the way that those values are queried is going to be different.

So let's define a creature.

So here is a creature the creature is going to have a name which is a string and it will also have attack and defense values.

But notice I'm doing lowercase here so I'm keeping them lowercase why is why am I doing that.

What's the point.

Well in the previous example we only applied the modifiers explicitly by calling the handle method.

What I want to be able to do is I want to be able to apply these modifiers automatically so as soon as you make modifier passing in a creature it automatically gets applied.

And so when you queried the creatures attack and defence values you automatically get the final calculated value in order to do this the attack and defence values here only store the initial values not the calculated values.

And we're going to grow a whole system of strokes on top of this in order to make this possible.

So first of all we're going to introduce a mediator or a central component which every creature refers to.

Now this component is going to be a game because most creatures participate in a game so we're going to have a game pointer here and game is going to be a centralized component that everyone kind of attaches to.

But before we get to implementing the game there's lots and lots of stuff that we need to implement.

So first of all I'm going to introduce this idea of command and query separation.

Now we're only going to deal with queries and the idea is that when you want to get a creature's attack or defense value you make a query which is a separate struct and you send this truck to the creature as opposed to just calling a method on it.

So how can this work.
Well we'll have a type called query.

There's gonna be a struct and you can specify in the struct what exactly you want to query so you can specify the name of the creature that you're interested in.

You can also specify what you want to query.

And we typically want to query one of two things.

We want to query either the attack value or the defense value.

So what I'm going to do up here is I'll define argument as an integer and I'll have two constants.

So I'll have a constant for Attack of type argument and I'll have similarly a defense constant so these are the two constants as they're going to be used as part of what to query.

So the type here is going to be argument and finally will have value as an int.

Now value is interesting because obviously that's the value you expect to read.

So you send off a query you look at the query after it's been processed and you look at its value but you can also specify the initial value so that whoever handles this query has a possibility of actually taking an existing value and modifying it like doubling it for example or increasing it by one and that way you can apply those bonuses that we talked about in the previous lesson.

So we have our query and now what we want to do is we wanted to find a bunch of interfaces for implementing the observed design pattern.

Now you can look for more information about this when we talk about the observer but here I'm just going to do this very quickly so we'll have an observer interface and this observer interface is going to have a method for handling a query like so we'll also have an observable interface.

So that is the interface implemented by whatever type wants to notify other users about something happening and here you can subscribe to events on this observer.

In our case the event is of course a query we don't have any explicit struct called event but queries here are treated as events.

We also have an unsubscribe.
Just duplicate this unsubscribe.

So that is when you want to stop listening to events coming from the observer.

And we have fire.

So fire is when somebody actually fires off a query and that query gets processed by whoever is interested in actually processing it.

So now we can finally build our game.

So the game is a centralized component so a game is going to be a struct and there's going to be just a single field in it and that field is going to be called observers and there's going to be a sync map so a sync map is going to allow us to basically keep a map of every single subscriber and to iterate this map to go through every single subscriber and notify on that subscribe.

So basically tell the subscriber that something has happened.

Okay so now that we have this what we need to be able to do is we need to implement the the observable

interface on the game because game is what every single participant in the game is going to be subscribed to.

They're going to be watching as basically so let's implement all the methods of the observable interface on the game like so.

So there's three methods here.
Subscribe unsubscribe and fire.
So let's begin with subscribe.

So here the idea is that you simply adds this observer to the list of observers specified here.

So we say oh adult observers or GDR themselves rather GDR observers they'll store and we store the observer.

But remember it's it's a key value kind of thing and we don't really care about the value we're just storing keys in this map so I'll just have an empty honest struck here like so.

And that's pretty much it.

When it comes to on subscribing you can simply remove it by value so to speak.

So you say gee all observers not delete and you delete this particular observer from the map.

And finally here's the fun part.
This is where you get to fire the events.

So here we have to go through every single observer.

So we say observe as the range and for every single observer we actually have to specify this as should be a capital range by the way we have to specify a function which actually does something for every single element.

So here the idea is that the key is going to represent the observer whose handle method we have to call.

But of course at some point this is going to be nil which means that we're out of value.

So if Ki is equal to nil then we just return false means that we are done here.

Otherwise we take key.
But of course key here is defined as interface.
So we have to cast it to an observer.

So we say ki dot observe a dot handle and we handle that query that we got passed in and then we return true.

Now we go.
So that's really all there is to it.

So we've now implemented the observable interface on our game so we have a creature here we can make a constructor which makes it easier to initialize the creature.

So here we go.

We have initialization of a creature given a game pointer here as well as the name attack and defense value.

So these are the initial attack and defence values and subsequently we can query more fool that is shall we say.

So the idea is that you don't address attack and defence directly instead you have gathers.

So you do need getters if you want to follow this approach.

So we'll have a function which is defined on the creature pointer called attack.

So this is going to give you the actual attack value.

Now in order to get this value what we do is we make a query object we get the game to fire the query.

Now whoever's listening gets to process the query and modify the final attack value and then we return that value.

So here we go we make a query we specify the parameter is the name of the creature of the value that we want to query in this case it's the attack value we're interested in and we specify the initial value which is seed out attack.

So this query now can be fired.

So we say see Dot Game dot fire and we pass a pointer like so and then we return cued up values.

So after every single element that wanted to process this event has in fact processed it and maybe a modified value we return the final value.

So this is the implementation of attack and I can copy this and have a similar implementation of defence.

So defence here would query for defence the initial value would be CDO defence like so and I'm just doing something on screen and then you simply return the value as before.

So that gives you the Defence value.
OK.
So that is the implementation.

Let's also have a stringer interface for the creature.

Did we do a stringer interface I can recall now.

But let me just do it quickly so we'll implement stringer on the creature here.

I'll just once again as print f percent s and the slash percent D.

So we'll have the creature's name and then look at the attack value.

So here we call attack method.
And here we call the defence method.

We don't just get the initial values we get the values as they are calculated.

OK.

So now that we have all of this the question is Will how can you implement those modifiers remember we had the modifiers that we had to apply explicitly well now they're going to be a flight implicitly.

So how is this going to work.

So we'll have once again we'll have a TIBCO creature modifier.

There's gonna be a struct is going to have a pointer to the game and there's going to have a pointer to the creature.

Okay so what happens inside the creature modifier.

Well a creature modifier is quite simply a template which means that even though you can theoretically give it a handle method make it effectively an observer which I can do right now let's let's implement observer here even though you can have a handle method there's really nothing to put in here because by default this thing only exists so that you can compose it as part of actual modifiers like the double attack modifier for example.

So let's do exactly that.

So if you want a type double attack modifier then you can compose creature modify.

You can simply stick it in here and that gets you all the field.

So now double attack modifier has game as well as creature.

You can also have a constructor and the constructor is actually required.

So it's not optional for us to have a construct that we absolutely need a double attack modify a constructor here and the reason why we need one is that before you actually return this thing.

So the result is the like.

So before you return the result you also have to make a subscription.

So whenever you actually make a double attack modifier you want that attack modifier to participate in the calculation of any values being queried from a creature.

So here I can say D dot subscribe Oh.

It's actually G don't subscribe.
Right.
One moment.
Oh of course.
Sorry.

That was a slight lapse so when you make a double attack modifier you have to specify a couple of things and I have neglected those so you have to specify the game that's being played as well as the creature that this modifier is being applied to.

Okay so now let me get rid of it.
Let's start again.

So you make a double attack modifier and you pass and increase your modifier where you specify the game as well as the creature.

Then you take the game and you subscribe on that creature.

So so you basically take the game and you subscribe the double attack modifier that you've just made so that whenever the game generates events the double attack modifier gets to process those events in our cases just the queries and then you return.

Okay so now we have a constructive for the double attack modifier we also want the handle method because double attack modifier has to be an observer.

So let's actually implement the observer interface or let's let me just verify this one moment.

Okay.
So.

So we do have the observer interface but it is just not letting me do it automatically.

So let me type that.
So we'll have fun.
The double attack modifier handle Q query.

Now we go.
Okay.
So how can we handle the query.

Basically if somebody is asking for a creature where the name matches the name of our creature remember we're using creature modify.

So we have a creature pointer.

So if the name matches and some of these querying attack because remember we're in attack value modifier then we can double that value so we say if you don't create your name is equal to D dot creature that name and Q What to query is in fact an attack value then we take the value of the query so cute up value and we multiply assign it by two.

So we double the value effectively.

Now another thing we can do is we can implement the closer interface on the double attack modifier so that's gonna be fun.

So let's take a closer interface.

Now the idea here is that you now have a close method which can be used to unsubscribe this particular modifier from from the game events.

So here we say d out game deduct game.

Unsubscribe on Subscribe not subscribe on subscribe.

And here we pass in ourselves we pass in a pointer to ourselves and then I'm just going to return nil here.

Okay.

So this is our closer interface and now we can put everything together and we can see how the whole thing works and now we have a central mediator which is the game.

So the game gets initialized with a sync map and then we make a goblin so a goblin is a new creature we first you specify the game pointer but then you give it a name like here we'll have a strong goblin and this golden is going to be a two to goblin.

Let's actually print it out so goblin does string like so.

And then what we can do is we can apply the double attack on to fire temporarily.

So for example what I can do is I can make a modifier and you double double attack modifier.

I could specify the game as well as the Goblin.

I can print out the golden once again so let's take a look of what happens after the modifier has been applied.

And then I can also call MDA close on the modifier there by on subscribing from the events and I can print the goblin once again after the end of this entire scope.

Okay.
So let's run all of us and let's take a look at what we get here.

Okay.
So as you can see the goblin is starting out with 2 2 values.
So that is the starting value.

Then we apply the double attack modifier the goblin becomes a 4 to notice we don't have to call any handler methods here because everything is kind of handled automatically through this central mediator.

So the modifier gets applied the goblin is now a 4 2.

But then of course we call close on the double attack modifier which unsubscribe is it from any handlers of the attack value queries.

And so by the time we get to query the attack values again the goblin is back to a 2 2.

So this has been much more sophisticated example of how you would build a how would you build a mediator with a chain of responsibility on top of it.

Because remember you can apply a lot of these modifiers one after another that they can come into the system they can go out of the system using the closed method and the State of the Goblin is always going to be consistent because effectively every time you're asking for the goblins attack or defense values you're recalculating it on the basis of a system.

So this is a maybe a more flexible implementation of the chain of responsibility design pattern.

### Broker Chain code: behavioral.chainofresponsibility.brokerchain.go

```go
package main

import (
  "fmt"
  "sync"
)

// cqs, mediator, cor

type Argument int

const (
  Attack Argument = iota
  Defense
)

type Query struct {
  CreatureName string
  WhatToQuery Argument
  Value int
}

type Observer interface {
  Handle(*Query)
}

type Observable interface {
  Subscribe(o Observer)
  Unsubscribe(o Observer)
  Fire(q *Query)
}

type Game struct {
  observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
  g.observers.Store(o, struct{}{})
  //                   ↑↑↑ empty anon struct
}

func (g *Game) Unsubscribe(o Observer) {
  g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
  g.observers.Range(func(key, value interface{}) bool {
    if key == nil {
      return false
    }
    key.(Observer).Handle(q)
    return true
  })
}

type Creature struct {
  game *Game
  Name string
  attack, defense int // ← private!
}

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
  return &Creature{game: game, Name: name, attack: attack, defense: defense}
}

func (c *Creature) Attack() int {
  q := Query{c.Name, Attack, c.attack}
  c.game.Fire(&q)
  return q.Value
}

func (c *Creature) Defense() int {
  q := Query{c.Name, Defense, c.defense}
  c.game.Fire(&q)
  return q.Value
}

func (c *Creature) String() string {
  return fmt.Sprintf("%s (%d/%d)",
    c.Name, c.Attack(), c.Defense())
}

// data common to all modifiers
type CreatureModifier struct {
  game *Game
  creature *Creature
}

func (c *CreatureModifier) Handle(*Query) {
  // nothing here!
}

type DoubleAttackModifier struct {
  CreatureModifier
}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
  d := &DoubleAttackModifier{CreatureModifier{g, c}}
  g.Subscribe(d)
  return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
  if q.CreatureName == d.creature.Name &&
    q.WhatToQuery == Attack {
    q.Value *= 2
  }
}

func (d *DoubleAttackModifier) Close() error {
  d.game.Unsubscribe(d)
  return nil
}

func main() {
  game := &Game{sync.Map{}}
  goblin := NewCreature(game, "Strong Goblin", 2, 2)
  fmt.Println(goblin.String())

  {
    m := NewDoubleAttackModifier(game, goblin)
    fmt.Println(goblin.String())
    m.Close()
  }

  fmt.Println(goblin.String())
}
```

## Summary

So let's try to summarize what we learned about the chain of responsibility design pattern so the chain of responsibility can be implemented as any kind of list you can use a linked list of pointers like we did or you can have some sort of centralized construct where you have a list of subscriptions for example and then it's fairly obvious how to work with it.

So you take the objects which are part of the chain you enlist them into the chain.

You can also control their orders.

So if some component has to be the first component to process a particular message you can force its insertion at the beginning of the list or at the beginning of the chain.

So it gets to process events before other elements do so.

That's another thing that you can do.

And then of course you also need to control other things like the removal of the objects from the chain.

That's also important because sometimes you no longer want a component to process particular messages.

And so it should be possible for them to to stop processing those.
