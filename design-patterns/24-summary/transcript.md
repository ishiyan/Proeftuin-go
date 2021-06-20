# Course Summary

All right.
So we've now reached the end of the course.

So let's try to summarize all the patterns that we've learned so first of all let's go through the creation of patterns.

So we looked at the builder design pattern and that is essentially a component for when the construction of an object gets too complicated.

We also looked at using several builders to build an object so several mutually cooperating builders.

Builders often have a fluent interface.
It just makes it a bit easier to use.
That's all.
No magic there.
We also looked at factories.

So first of all we talked about factory functions also known as constructors they are a very common thing because God doesn't have ordinary constructors like many other programming languages but also we noted the fact that a factory doesn't have to be a function it can be a simple function or it can be a dedicated struct full of specific functions for different ways that you want to construct something.

Then we looked at the prototype design pattern the idea of creating objects from existing objects and we looked that looked at the fact that this requires some sort of deep copy support.

So whether it's through serialization or through writing some sort of copying code yourself you have to have support for copying everything including the pointers because that's really the major problem that when you copy a pointer you that you don't just want to copy the pointer value you want to create a brand new object which is a copy of what the pointer actually pointed to then we looked at the much maligned singleton pattern the idea that when you sometimes you just want to ensure that there is only a single instance of a particular object in the system.

We saw that it's very easy to make this thread safe and lazy so there is no problem in actually implementing the Singleton design pattern.

But realistically you want to adhere to the dependency inversion principle so for a given Singleton you might want to define an interface for that Singleton and then use that interface or of course you can use dependency injection which also makes things a bit easier and a bit safer as well.

So having looked at the creation of design power let's talk about the structure design pattern.

So we have adaptor which adapts the interface that you get to the interface that you need.

We looked at Bridge which tries to decouple abstraction from implementation.

We looked at the composite design pattern which allows you to treat scalar objects or individual objects and entire compositions of objects in a uniform manner meaning having the same API.

We also looked at decorator which is this idea of attaching additional responsibility to object while preserving the open closed principle so instead of modifying the object you make you make some sort of aggregate which includes the original and also includes additional behaviors and fields and so on and so forth.

And we saw that this can be done well we looked at the approach when we did it through embedding but you can also do three pointers there's really no major difference there OK so other structured patterns we looked at the facade design pattern basically hiding complicated systems entire set of systems behind a simple interface.

The flyway design pattern the idea of efficiently storing large numbers of similar objects by preventing any kind of duplication we looked at proxy which is a way of providing some sort of surrogate that forwards calls to the real object and also performs additional functions typically just controlling the way that the underlying object operates.

So all things the proxy isn't really just one pattern.

There are lots of approaches you you can use it for access control for communication across let's say process boundaries for it adding just just some basic functionality like logging for example so it's really up to you.

But the overall approach is that you try to replicate the interface of the original object and then we move onto the behavioral design patterns.

So we looked at chain of responsibility where all the.

So we allowed the components to process informations one after another in some sort of chain and this this is where you can take an approach where every single element in the chain simply refers to the next element with a pointer for example or you can just make a list somewhere and use the mediator design pattern just have everyone refer to the mediator and the mediator has a list so the mediator knows how to inform everyone of something happening.

Then we looked at the command which is this idea of encapsulating a request into a separate object which has lots of benefits like being able to audit things and being able to replay the events and undo and redo those.

It's also part of the U.S. command query separation as well as come back with responsibility segregation the kind of enterprise key approaches to software development.

Then we talked about the interpretive design pattern the idea that we transform textual inputs into structures like everywhere like all the compilers and interpreters and ideas do it all the time.

So this is done literally everywhere and we use it everywhere.

So in actual fact the interpreter design pattern is a separate branch of computer science typically called compiler theory.

Unfortunately nowadays not many universities actually teach this because they think that it's not relevant to all students so you wouldn't find this anywhere.

But if you're interested there are plenty of books available to figure out how to build the different passes and so on.

The examples I've provided are very simple but in actual fact that's a fairly complicated field.

Then we looked at iterator the idea that you have some special component which provides an interface for accessing the elements of some aggregate object basically a way of navigating through the elements of a particular object so the iterator is in a way.

And those sensitive to the visitor you can do one or the other.

OK.
What else did we look at.

We looked at mediator which is a some sort of central component which provides mediation services between several objects and those objects don't have to be necessarily aware of one another.

And the examples include just message passing and things like building a chat room for example.

Then we looked at the mental design pattern.

Basically when you have some sort of tokens that represent system states and then maybe you can revert the system back to the state typically those tokens do not allow any manipulation but they can be used in various APIs as just representations without without any mutable state of course.

We talked about the observed design pattern the idea that you have events and you allow notification of things happening in a component.

So one component can generate events another other component can listen to that events and they can unsubscribe from those events as well.

What else.
So we have the state pattern.

That's when you have basically a system modeled by a number of states and transitions between those days and this is typically called a state machine.

We there are special frameworks which actually orchestrate state machines but simple ones like the ones I've shown in the demos in this course you can just build yourself.

So if you have a very simple state machine it's probably better to just make it yourself but if you have something complicated then it might be worth finding some sort of library which has more advanced features shall we say.

Then we looked at strategy and template method.

They are both very similar both defined some sort of skeleton algorithm where the details are filled in by the implementer strategy uses composition and the template method doesn't although the distinction is blurred a little bit between those two in the Go language because there is no inheritance.

Then we looked at the visitor design pattern that allows a kind of non-intrusive additional functionality to hierarchies of types.

Of course it's not entirely non-intrusive because you have to have that except method propagate throughout the entire hierarchy.

Otherwise nothing will work but you only have to do this once and once you've done this you can build different visitors and have them visit your hierarchy of types.
