# Factories -- ways of controlling how object is constructed

## Overview

All right so now let's talk about the factory design pattern.

Now notice here I'm seeing factories.

It's actually a different design patterns that will encounter which all somehow relate to the process of construction kind of like factories in the real world.

Because in the real world factories they they construct objects like computers or iPhones or whatever.

So it's a similar thing here.

So object creation sometimes becomes too complicated.

Sometimes you want to have additional things like for example let's suppose you have a struct which has lots of lists and maps in it.

Now you want to initialize those but you don't want to initialize them element by element every single time that somebody wants to use it.

You want to give them default values you want to call make on whatever types are actually required.

So you create some extra constructs to handle this.

So yeah a typical example is a struct with too many fields which needs to be which need to be initialized correctly and maybe there is some business logic involved there in terms of how things are initialized.

But the point here is that we're talking about wholesale object creation.

So when we talk about the builder or the builder does piecewise creation which means that the builder slowly builds up an object step by step it doesn't try to do everything at once whereas in the factory we are mainly talking about creating an object in one single invocation of let's say the factory function.

So when we want to do a wholesale object creation we can outsource it to different construct.

So first of all you can have a separate function a factory function and I'm also going to be referring to it as a constructor.

In actual fact I'm not the only one if you look inside the go Lang idea for for instance and if you open the Generate menu you will see a generate menu for the constructor.

So I will be referring to these I's either factory functions or constructors and factory functions are of course a common approach in the Go language.

It's basically what you what you create if you need some special function which acts as a constructor in other programming languages.

So this is one approach but you can also if you have a lot of these you can take all of them and you can put them into a separate truck just for the sake of organization.

Now you can also put them into a separate package but having a separate struct Well it's a little bit neater it's really up to how you want to handle it but you can just take if you have several factory functions you can group them together somehow and this becomes a factory so a factory can for example be passed into as a parameter into some method or other.

So that's an additional benefit here.

So a factory is just any component which is responsible solely for the wholesale not piece wise creation of objects.

## Factory Function

In most cases when working with go you can take a struct and you can just fully initialize struct using the curly brace sanitation.

So for example if I make a person struct which has for example and name which is a string and an age which is an end to it to initialize a person all you typically do is you'd say P is a person and then in between the curly braces you would provide the name and you would provide the person's age and that's pretty much all there is to it.

However there are situations when you want to have some sort of behavior happen as an object is created as destruct is created for example.

Let's suppose that you want a default value so let's say there is another property here another field which is called I count.

So that's the number of eyes a person has now a person typically has two eyes and in most cases like in ninety nine point nine recurring percent of cases you wouldn't really want to customize this field you wouldn't want to give it some other value like three for example it's always two and you don't want to keep typing too all over the place.

The question is how can you do this well you can do this by having some sort of factory function.

So a factory function is nothing more than a freestanding function which returns an instance of this truck that you want to create.

So here we would call this new person and here you would provide the arguments like a name which would be a string and age which would be an ant and then you would return a person like so and here you can return a person way you apply those settings so you have the name and The Age and then when it comes to I count you just put the number two here because that's a good default and that's a default that people can customize in their code.

So here we can create a person using this new person syntax so we can provide the name and The Age for example but if we do want to customize it like if if John gets gets shot in the eye or something then you can say Pete I count equals 1 or something equally sad now or let's say John doing something radioactive he gets four eyes instead.

So that's that's also a possibility.

So you can customize the object but you do have this default value that's being specified right here.

Now what we can also do is instead of returning by value what we can do is we can return the pointer.

So here you can specify the pointer type and here you would just put an ampersand so you would construct a person and you would return a pointer to that person which is somewhat more efficient than returning it by value.

So this is essentially the gist of factory functions U.S. see factory functions all over the place when there is some logic that needs to be applied when initializing a particular object.

So let's say for example that you want to check that when you're constructing a person the age value is actually valid or The Age value maybe maybe you want to check that the person is of legal age.

So you say if age is less than 16 and then you can you can perform some sort of validation here.

So the role of the story here is that if you have simple objects then you can initialize those objects just using the curly brace notation but sometimes you want additional logic additional stuff to happen as you are creating a particular struct and that's when you would use these factory functions.

### Factory Function code: creational.factories.factoryfunction.go

```go
package factories

import "fmt"

type Person struct {
  Name string
  Age int
}

// factory function
//func NewPerson(name string, age int) Person {
//  return Person{name, age}
//}

func NewPerson(name string, age int) *Person {
  return &Person{name, age}
}

func main_() {
  // initialize directly
  p := Person{"John", 22}
  fmt.Println(p)

  // use a constructor
  p2 := NewPerson("Jane", 21)
  p2.Age = 30
  fmt.Println(p2)
}
```

## Interface Factory

When you have a factory function you don't always have to return a struct.
Instead what you can do is you can return an interface that that struct conforms to.
So let me show you how this can work.

So essentially you will go with the same demo as before so we'll have a type called Person.

But this time round I'm going to do something different.
I'll have person with lower case as opposed to Africa.

So person is going to be a lowercase struct and we'll have a name as a string and age as an int but those things are going to be kind of hidden from the consumer.

So what we want a person to be able to do is we want the person to for example be able to say hello.

So what you can do is you can define an interface so we can define an upper case person interface.

And this is going to define the behaviors that are available on the person.

So the person's methods and one of those methods is going to be say hello like so.

So what we can do now is we can actually implement this interface.

So here I can have a function which takes the person points and the notice lowercase P here and it's going to be say hello like so and here I'll just do a print f quickly so we'll print f something like

Hi my name is whatever the name is.
I am however many years old followed by a line break.
And here we can specify Pete out name and Pete out h respectively.

So this is how you a person can actually say hello and say something about themselves.

But when it comes to making a factory function what you do is you don't return person instead you return the interface.

So you return a lowercase person as opposed to the uppercase person.

So here you have a function called new person which takes a name as a string and ages and ends and it returns a person uppercase.

So notice they use off like lowercase here for example in uppercase here.

So here you would return the uppercase notice that because it's an interface we don't have to put a star in front but we do have to put an ampersand before we actually construct a lowercase person with the name and the Age and return that.

So to use this is really no difference you make you say P is equal to a new person with the name James and H 34 for example.

And then you can say pedo say hello and if we run this then we should see the expected result.

Hi my name is James.
I am 34 years old.

So what is the difference between this approach and the approach where we return the actual struct.

Well the difference is obviously that now that you have just a an interface to work with you cannot really use that interface to for example modify the underlying type.

Because remember we made everything lowercase and we're not exposing the lowercase person at all we're just exposing the interface.

So you cannot say something like Pitot H plus plus for example you cannot do this because age and name are not accessible so this is a neat way of encapsulating some information and just having your factory expose just the interface that you can subsequently work with.

And that way you can you can for example have different underlying types because well you can imagine that we can have some other structure here.

So for example you can have you're going to have some I don't know tied person.

So a retired person would also have a name and an age like so.

But for example for the tired person their implementation of say hello can be completely different because they can say Sorry I'm too tired to talk to you or something.

So we can we can do it like this.
Sorry I'm too tired.

And then a tired person let's suppose that a person that is older than 100 is by definition tired.

So here inside the factory function we can say if age is greater than 100 then we return a tired person.

So a person that's seen far too much name and age otherwise we return this bus and then it becomes interesting because then what you get the actual underlying object is different depending on what in invocation here so.

So if we do one hundred and thirty four for example well obviously this person is too tired and we are we are outputting putting this extra thing here that we're not supposed to be doing all those extra variables.

But apart from that that's that's what you would have to fix.

So as you can see we're actually getting a different object but because we have an interface in front of it we just return the interface and that way we can have different types of objects in the background and that can also be useful in some situations.

### Interface Factory code: creational.factories.interfacefactory.go

```go
package main

import "fmt"

type Person interface {
  SayHello()
}

type person struct {
  name string
  age int
}

type tiredPerson struct {
  name string
  age int
}

func (p *person) SayHello() {
  fmt.Printf("Hi, my name is %s, I am %d years old.\n", p.name, p.age)
}

func (p *tiredPerson) SayHello() {
  fmt.Printf("Sorry, I'm too tired to talk to you.\n")
}

// note no * in front of Person, because it is an interface
// note & in front of person, we return a pointer

func NewPerson(name string, age int) Person {
  if age > 100 {
    return &tiredPerson{name, age}
  }
  return &person{name, age}
}

func main() {
  p1 := NewPerson("James", 34)
  p1.SayHello()

  p2 := NewPerson("Jill", 134)
  p2.SayHello()
}
```

## Factory Generator

OK so now we're beginning to take a look at is the idea of generating factories.
So it's a different approach to what we've seen before.
Basically we're going to work with the same scenario.

So we'll have an employee which is just going to be composed of a name and position as strings as well as annual income as an end.

And what we want to be able to do is we want to be able to create factories dependent upon the settings that we want employees to subsequently be manufactured.

Because remember in the previous example what we did is we basically created an employee.

So we use new employee to basically just just specify some sort of flag but then we customize it.

So the next thing we did is we say we said something like EDA name and we provided that missing information it would be nice if we could do it in one statement.

So we can say new employee where you specify you know the type of the employee that with a we want and also provide the name here as well and have those two things combined somehow.

So that's what we want to build here.

So we want to be able to create factories for specific roles within the company and there are two ways of doing it.

There is the functional approach and the structural approach.
So first of all let's take a look at the functional approach.

So the functional approach would be that you would make factory function that would actually not return an object but instead return a function.

That's why it's called the functional approach.
So we would call it something like new employee factory.

So notice we're not creating an employee we're creating an employee factory that you can subsequently use to fine tune those details of that object.

So here we can make a factory where you specify the position.

You also specify the annual income and those things get baked into the factory and what the factory function here returns is a function.

So it's a function that returns a function.

In other words a higher or the function it returns a function which takes a name and uses that information to fully initialize an employee and return that so the return type here is employee point.

And what we're doing is we're returning a function which takes a name as a string returns an employee pointer like so and it fully initialized as the employee is fully initialized as the employee including the name the position and the annual income.

So I just forgot to specify that this is an end in actual fact.
There we go.
OK.
So how does this work.
Well it works as follows.

Basically what you do is you make factories and you still those factories in variable so for example you can have a developer factory so a developer factory would be a new employee factory where the position is developer.

Sorry about that.
The position is developer and the annual income is let's say well 60000.

And similarly you can have a manager factory so a manager factory is a new employee factory where the position is manager and annual income is 80000.

And now that you have both of these these are functions.

So these are not just objects the door structures these are functions and you can invoke those functions to get the actual result.

So for example you can make a developer so you make a developer by saying developer is developer factory with the name Adam for example and similarly you can say manager he is a and you invoke the manager factory specifying the manager as Jane and then we can print those out.

So F.A. got a print line here so so I can print line the developer.
And similarly I can print out the manager as well.
Let's take a look at what we get here.
So as you can see we're getting Adam which is a developer in Jane which is a manager.
So what is the advantage of this approach.

Well one advantage is that now that you have these factories stored in parables You can also pass these variables into other functions and that is the core of functional programming passing functions into functions.

So here we have a function being written from a function and then you can subsequently pass it into other functions to be consumed and by consumed we mean that somebody invokes it.

So somebody takes this factory function and simply invokes it providing the last missing piece because we've provided all the pieces here with provide the position as well as the annual income.

All you have to do is you have to fill in the gap and provide the name and you have a fully initialized object and that's exactly what's happening here and here respectively.

So that's the functional approach.

The functional approach is pretty good but there is another approach of course and that is a more structural approach basically making a factory a struct now it's not strictly speaking necessary at least not in go but you can do this too.

So so for example if you want to somehow incorporate information about the fact that a particular factory initialize as an employee with a particular position an annual income you would do it as follows you'd make a type called let's say employee factory as a struct and that's where you would actually store the information so you would store position as a string and you would store annual income as an end.

These can be lowercase if you want them to be and then what you would do is you would make a factory function for actually returning instances of this particular factory and you would also give this factory some sort of method for actually creating the object.

So first of all the method would look something like the following.

So you would have an employee factory and you would give it a method called Create where you specify that missing piece which is the name so you specify the name and it returns an employee like so and here you just fully construct the employee and return that so you return employee with a name and you take the position and the annual income from the actual factory so you say after position and after annual income like so and then of course what you need to do is you need to have a factory function for creating this factory because well because you want to have a different predefined employee factories obviously and you can also use the curly brace notation if you want to do this but just for the sake of completeness you could have a function like new employee factory.

I'll put it to here because we've already got one of those.

So this is where you would specify a position as well has an annual income and you would return an employee factory like so and here you would just constructed and return it.

Employee factory with position lowercase position and annual income annual income like so.

OK.
So how would you use this.
Well it's it's rather simple.

So here somewhere let's suppose we have a new new position like CEO so we can make an boss factory.

So both factory in this case would be a new employee factory to notice the two we're using that new approach.

And here once again you would specify the position as well as the yearly earnings like a hundred thousand for example and then you would use the Create a method to actually create an object.

So the boss will be created using boss factory not create where you would specify that missing piece.

The name of the manager like Sam for example.
And once again we can print line what we actually got.

So let's take a look at that OK so you can see here we have Sam who is CEO earning 100000 now what is the advantage of this approach over this approach.

Well the only real advantage is that how after these factories up here are created the functional factories you can really customize them.

So if you set your developer to 60000 like we did here you cannot modify to 65000 later on.

Meaning you cannot say develop a factory dot annual income plus equals whatever you cannot do this but you can do this here you can change this factory because remember this factory actually just stores these gas fields.

So what you can do is you can say Boss factory does annual income equals and let's say one hundred and ten thousand.

So you had a salary increase and now if you run this you'll see a hundred and ten thousand here because we've actually been able to modify it.

But in terms of usability in third party code let's say providing a function and passing it into some other some other piece of API then the situation is different because obviously in the in these cases these are just ordinary functions in passing ordinary functions into something is easier than passing in a specialized object because for a specialized object whoever is consuming that object has to explicitly know that there is some sort of create method and then they have to call this create method.

So this is a situation where for example you might try to introduce some sort of interface which tells you explicitly that there is a create method there and here are the arguments and then you could also use this kind of interface approach to to pass an interface of the factory rather than the factory itself.

So that's a possibility but whichever option you go both of them are just fine the first one is probably more idiomatic more kind of more functional more idiomatic.

So I would recommend the first option but the second option is also there if you need it for some reason.

### Factory Generator code: creational.factories.factorygenerator.go

```go
package main

import "fmt"

type Employee struct {
  Name, Position string
  AnnualIncome int
}

// what if we want factories for specific roles?

// functional approach
func NewEmployeeFactory(position string,
  annualIncome int) func(name string) *Employee {
  return func(name string) *Employee {
    return &Employee{name, position, annualIncome}
  }
}

// structural approach
type EmployeeFactory struct {
  Position string
  AnnualIncome int
}

func NewEmployeeFactory2(position string,
  annualIncome int) *EmployeeFactory {
  return &EmployeeFactory{position, annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee {
  return &Employee{name, f.Position, f.AnnualIncome}
}

func main() {
  developerFactory := NewEmployeeFactory("Developer", 60000)
  managerFactory := NewEmployeeFactory("Manager", 80000)

  developer := developerFactory("Adam")
  fmt.Println(developer)

  manager := managerFactory("Jane")
  fmt.Println(manager)

  bossFactory := NewEmployeeFactory2("CEO", 100000)
  // can modify post-hoc
  bossFactory.AnnualIncome = 110000
  boss := bossFactory.Create("Sam")
  fmt.Println(boss)
}
```

## Prototype Factory

So sometimes you'll find yourself in a situation where you're creating lots of very similar objects like for example you have a company and you have a couple of fixed positions in the company and each of those positions pays a certain fixed amount and you want to create a developer struct or employee struct where you have the position and the annual income selected from some table or other.

So in this case what you can do is you can have what's effectively a prototype factory.

So this is also related to the prototype design pattern so you can take a look at that but basically you can have a freakin figured objects and then you can have a factory function which actually operates on and gives you a particular pre configured object.

So let me show you a very simple example let's suppose we have an employee which is a struct so an employee has a name and a position within the company which are both strings and in addition they also have annual income which isn't in.

Now what we can do is we can create a couple of constants representing the kind of predefined employee prototypes or predefined employee templates if you will.

So we'll have a few of those.
So let's suppose that we have a developer and we'll also have a manager like so.
So these are the two options that we can create.

Now what we can do is we can create a factory function which is going to actually take as a parameter and not like the name and position and whatever but they can actually take one of these as a parameter and depending on the value taken from this election it can give you an appropriate the employee which is already pre initialized with some predefined data relative to your company.

So we can have a factory function called new employee where you specify the role which is an integer and you return an employee pointer.

And here we look at the role.

So we switch roles and here we say well if it's a developer then we can return an employee where we don't specify the name.

So there's an empty string here but we do say this is a developer and we do say that their annual salary is let's say 60000.

And similarly if it's a manager for example we can similarly return an employee once again leaving an empty name where this is a manager and let's say the manager earns a bit more let's say 80000.

So you can also have the default case where if somebody provides some spurious in value you can just panic you can say unsupported row and then you can use this factory function as follows so you can make a manager for example you simply call new employee and you provide the manager constant as the one and only argument and then once this initialize is the structure for you you can customize it so you can say the name equals sound for example so you give the manager a name and then we can print f or just a print line this and take a look at what we get here.

Hopefully the right things so we get Sam who is a manager earning 80000 a year.

So what does this demonstrate.

Basically it demonstrates that there is yet another approach to this where you have these predefined pre configured objects and you have them sort of being returned on demand depending on some flag or rather and here I'm using just integers but you can you can probably see that this can be a string for example or something else.

So it doesn't really matter.

So this is yet another approach that can work for sort of specifying the kind of object that you want created and of course instead of having the actual struct here you can have an interface type instead and that that is also going to work just fine.

### Prototype Factory code: creational.factories.prototypefactory.go

```go
package factories

import "fmt"

type Employee struct {
  Name, Position string
  AnnualIncome int
}

const (
  Developer = iota
  Manager
)

// functional
func NewEmployee(role int) *Employee {
  switch role {
  case Developer:
    return &Employee{"", "Developer", 60000}
  case Manager:
    return &Employee{"", "Manager", 80000}
  default:
    panic("unsupported role")
  }
}

func main() {
  m := NewEmployee(Manager)
  m.Name = "Sam"
  fmt.Println(m)
}
```

## Summary

So let's discuss the things that we've learned about the factory design pattern.

So a factory function.

That's the most common thing you're likely to see and the idea can help you.

Generally those things is basically a helper function for making struct instances.

And this is where you can perform additional initialization or you can omit the initialization of some fields.

It gives you additional flexibility.
That's what I'm trying to say.
And a factory really.

Generally the word factory means that it's any entity that can take care of object creation so it can be a function but it can also be just just a separate struct that has methods for constructing objects.

So you have these two different options and I'm sure you can think of more options like for example you could have a function which takes a function so a higher order function and you can pass in a factory which is just a function that generates things.

So plenty of options here.
