# Decorator -- adding behaviour without altering the type itself

## Overview

We're now going to take a look at the decorator design pattern.
So what is this all about.

Well the motivation for using decorator is as follows We want to take an existing object and we want to augment this object with additional functionality so we don't want to rewrite or alter the existing object.

We don't want to go into an existing struct and start modifying it because we want to stick by the open closed principle remember open for extension but close from modification.

So essentially we treat code that's already been written and tested as close for modification you cannot modify it unless there's an emergency or something and sure in some situations you can modify existing code if if you're in total control of it and you don't have let's say external clients using it already but in some cases you just want to extend things.

So another motivation for having a decorator is that you want to keep the new functionality separate.

So you want to keep enhancements to a particular object separate from that object.

And this goes to the separation of concerns and this idea of the Single Responsibility Principle.

So we're adhering to two solid design principles at once.
In this particular design pattern.

So we want to be able to interact with existing structures meaning that when you make a decorator over some structure X you also want access to that original structures data and its methods as well.

So we want to take care of all of this.
And the solution here is to simply use embedding.

So you embed the decorated object into your new object your decorator and then you provide the additional functionality.

So the decorator is a pattern which facilitates the addition of behaviors to individual objects through the use of embedding.

## Multiple Inheritance

One issue that developers face in different programming languages is how to combine the functionality of several structures inside a single structure.

Now unfortunately with languages that do not support multiple inheritance that becomes difficult.

And of course with go there is no inheritance.

There is only aggregation and unfortunately there are situations where aggregation doesn't give you the desired result.

So I want to show you where this can go wrong as well as how this problem can be fixed so let's imagine that we were going to do a very classic demo of a dragon which is both a bud and a lizard.

At the same time so we're going to try to construct a dragon from separate bud and lizard structures.

So first of all I'll define the structure called Bud.

It's going to have an age which is an integer and a bird is going to be able to fly assuming it's old enough.

So we'll have a function for birds.
I'll call it fly.

So if the bird is let's say all the or equal to 10 years old then it can fly.

So we can just print line that the bird is flying flying.
Now we go.
Okay.
So a dragon is also lizard.

So we can define another type called lizard which also has an age which is an integer and we'll have a function for crawling.

So a lizard can crawl assuming it's younger than 10 years old.

So if l that age is less than less than 10 then we can just print line that it's crawling.

Okay.
So these are these separate struct.
So we have a struct for the bird and another struct for the lizard.

Then we wanted to find that dragon which is both a bird and the lizard at the same time.

And in canonical go what you would do is you would simply aggregate both of these.

So you would say type dragons shrugged and you would simply put both bird and lizard in here.

Unfortunately this can cause certain problems you can see that in my case both the bird and the lizard have an age field so there is an age field here and there is an age field here and that's going to be a problem as we try to sort of combine the operations of these two.

So we can certainly do some things with the dragon so let's make a dragon so I can make a dragon like this.

And now the problem is that I want to initialize the Dragon and I want to set the dragon's age and this is where things get difficult because you cannot say detailed age equals 10 for example that's not allowed in.

In fact if you try to compile it you'll see that there is an ambiguous selector here.

So an ambiguous selector means that you basically have to specify which age specifically you're referring to are you referring to the age stored in the in the bird or in the lizard.

So you can do things like lead up bird that age and the adult lizard adult age just set them both to some value and then yes if you do this then everything is going to work fine so that you can do the fly detail crawl and let's get rid of extra import here and if we now run this you should be able to see that we are flying just fine.

Now the problem with this scenario is that you can introduce a really nasty inconsistency into the behavior of the dragon if you said the different ages to different values and after all you don't need to separate fields.

It's a single age you want to keep it in a single field.
So how do you do this.
Well there is really no obvious way of doing it.

If you go ahead with this approach so you can suddenly have utility methods for actually setting the age and getting the age as well.

So for example you can have a function for getting the age

age where you would simply return dead.

But that age for example and similarly you can have a function for getting the age or rather setting it so you would have something like set age where age is an end.

And here you would set the properties for both the bird part as well as the lizard part so you would say that Bud dot age equals age.

And similarly you would say these dots lizard dot age equals age.

So this is something that's going to work to some degree in the sense that we can now go ahead and we can do the dots set age 5 for example.

This is going to work and it's going to keep the dragon in a consistent state.

Unfortunately it's still very possible to make it inconsistent by going into the dragon explicitly and saying that but that age equals 55 for example.

And in this case you've kind of broken the operations because they're inconsistent across the two parts that you've aggregated.

So there is no real solution to this at least not in go.

There is no language feature that will allow you to kind of regularize this whole situation.

But what we can try to do is we can try to design the entire set of structure differently so that this is avoided.

And so that instead of simply aggregating we'll build a decorator we'll build a proper decorator around the bird and lizard types.

Okay so I'm going to get rid of all of this and we're going to try again and this time round what I'm going to do is I'll just be pasting lots of code because there is a lot more typing involved here.

So first of all what you want is you want some sort of interface.

Like here I have an HD interface and this defines a contract for having the H.

Get her and set her and I getters and setters are not particularly idiomatic go.

In most cases you want to avoid having getters and setters but in this particular situation unfortunately there is no way around it.

So what you would do now is you would construct a bug type which would conform to this interface let me paste it down here so as before but is a struct it has an age but notice the age is lowercase here and then we have a getter and a setter for the H field.

We also have an ability to fly aware that field is being used.

Similarly you can have a definition for a lizard.

So let me paste it down here so the lizard also has an age it has the getter the setter and an ability to crawl.

And now we can construct the dragon but we're going to construct it differently and by the way let me get rid of all of this code.

So the construction of the dragon is going to be like this will have a dragon but this time round we are going to keep the bird and lizard inside the Dragon but we're not going to do the straightforward aggregation where you automatically get all the members instead we'll have just field so we'll have a field for the bird and a field for the lizard.

So now that we have this what we can do is we can redefine the behaviors such as flying for all and simply redirect them or proxy them to the appropriate fields but before we do that we have to have the getter in the center for the age and this is going to be the same as I've done it when I was talking about how to figure out the situation with aggregation so the idea is fairly similar.

So when you do the setter you set both the bird and lizard age but noticed that unlike in the previous example these are lower.

So the idea is you shouldn't be able to access the fields directly instead you only operate on behaviors and then of course you define those behaviors so you define the behaviors in the following fashion.

So when it comes to flying the dragon can fly into fly.
It just uses the bird path to fly.

And similarly the dragon can crawl and it uses the lizard part to perform the crawling.

Of course now what we have is we have a situation where the dragon has to be initialized by providing instances of bird and lizard and this implies that you have to have some sort of factory functions for doing that.

So here I have a function called new dragon which creates a dragon and it initialized as the bird and lizard parts and now we can put everything together and we can start using it.

So ask before I can make a dragon and this time round I can set the age to 10 for example and I can call the methods I can call the fly and I can call data crawl and if we run this we should hopefully get the right result so here we're flying because we're 10 years old and if I set this to fly for example then we're going to be crawling hopefully.

Yep that's exactly what we get.

So what is the situation here and how does it relate to the decorator.

Well in the Dragon class that you see here the dragon struct rather we have constructed a decorator.

We have constructed an object which extends the behaviors of the the types that we have right here.

But what it's doing really is it's providing better access to the underlying fields of both.

But then the lizard because you have to set the consistently and you also have to be able to get the H consistently and in addition it combines the behaviors of both the bird and the lizard by providing the interface members with the same names and simply promising over the coals to the underlying fields.

### Multiple Inheritance code: structural.composite.multipleinheritance.go

```go
package main

import "fmt"

/*
type Bird struct {
  Age int
}

func (b *Bird) Fly() {
  if b.Age >= 10 {
    fmt.Println("Flying!")
  }
}

type Lizard struct {
  Age int
}

func (l *Lizard) Crawl() {
  if l.Age < 10 {
    fmt.Println("Crawling!")
  }
}

type Dragon struct {
  Bird
  Lizard
}
*/

type Aged interface {
  Age() int
  SetAge(age int)
}

type Bird struct {
  age int
}

func (b *Bird) Age() int { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

func (b *Bird) Fly() {
  if b.age >= 10 {
    fmt.Println("Flying!")
  }
}

type Lizard struct {
  age int
}

func (l *Lizard) Age() int { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

func (l *Lizard) Crawl() {
  if l.age < 10 {
    fmt.Println("Crawling!")
  }
}

type Dragon struct {
  bird Bird
  lizard Lizard
}

func (d *Dragon) Age() int {
  return d.bird.age
}

func (d *Dragon) SetAge(age int) {
  d.bird.SetAge(age)
  d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
  d.bird.Fly()
}

func (d *Dragon) Crawl() {
  d.lizard.Crawl()
}

func NewDragon() *Dragon {
  return &Dragon{Bird{}, Lizard{}}
}

func main() {
  //d := Dragon{}
  //d.Bird.Age = 10
  //fmt.Println(d.Lizard.Age)
  //d.Fly()
  //d.Crawl()

  d := NewDragon()
  d.SetAge(10)
  d.Fly()
  d.Crawl()
}
```

## Decorator

To demonstrate the use of the decorator design pattern we're going to consider a very simple example an example where we have a bunch of geometric shapes that we have in the system and we want to extend the functionality of those geometric shapes by giving them additional properties.

And the question is how do you do this.
Well first of all let's set up the scenario.

So let's suppose that we have some sort of interface called Shape and this interface is going to allow a shape to render itself OK.

Now since we're doing a text based demo I'm going to render it to a string as opposed to a graphical object like you would do in the real world.

And now we can implement a bunch of different types of.

For example I can have a type circle which is going to be a struct with a single member called radius which is going to be a floating point number.

Let's implement the let's implement the shape interface on the circle.

So here let's just print out let's just print out the fact that we have a circle of a particular radius.

I'm going to say that we have a circle of radius center F..
And here I'll just say C dot radius.

So where we're turning a string that string is going to kind of emulate the fact that we're really rendering it somewhere and we can also have a member of a circle like for example here I can have a member called resize where there's some factor by which you resize the circle.

So here see that radius gets multiply assigned by that value factor like so.

Okay.

So now we can also have a square which is also going to be a shapes that once again type square struct.

Okay side float 32.

And then we'll have the implementation of the shape interface as before.

So here I'm going to sprint to have something different so F.A. return FMC the s print Taff I'm going to say that I have a square with side percent F where SDL side is being used and that's pretty much it.

Okay.

So imagine you have these circles and squares operating in your system and what you want to do is you want to color them for example.

So how do you make a colored circle now or indeed call it square.

Now one approach is that you jump back into this truck and you add additional members you go ahead and you add additional members here.

However the real problem with this is it breaks the open closed principle.

Remember we have this assumption that once you've written and tested your types you don't want to jump into them and modify them because you don't know how this is going to affect their behavior.

So for example maybe you're a C realizing this types to a file and you've changed this structure thereby you might break the serialization or it might be some other operation maybe somewhere in your code you're using size of to measure the size of the entire struct and you're going to be breaking that code.

So it's a lot better to just extend these types and the question is Well how do you extend them.

Because one possibility is that you can simply aggregate and thereby extending the type.

So for example if I went to college square I could say type colored square struct where I have a square and I also have the color as a string.

Now on the one hand this could work if you have a small number of struct that you want to extend this way.

There is no reason why it wouldn't work on the other hand if you have entire groups of struct.

So if you have lots of different drugs like circles and squares and triangles and rectangles and all sorts of other shapes then having a counterpart colored class for every single one of those types just doesn't make sense.

It's just too much effort and we don't really want to do it this way because unfortunately without generics we cannot specify that this is going to.

This is going to include this some some type provided as a type parameter.
There's no functionality like this and go.

So what we can do instead is we can use a shape in here and we can say that instead of having colored squares and colored circles we'll have a colored shape and that's going to be our first decorator.

So the color shape is going to refer to the shape it's decorating shape shape and it's going to specify the color.

Now color shape itself is also a shape so it has to implement the shape interface.

So as I implement this interface what I can do is I can use the underlying implementation of the interface remember shape itself also implements the shape interface and so it has its own render method so we can reuse that render method and we can also add the color.

So here when it comes to rendering the shape I can say I can return F.A. dot as print f I can use the underlying render percentiles and then I can say it also has the color another percent yes.

So the first one is going to be CDO shaped up render.

So we use the underlying shape to actually render what what.
Whatever it is that we have.
And then we specify the color see the color.
Okay.

So this is our first decorator and I want to show you how you can start using it.

So imagine you have a circle.
So we have a circle with a radius of 2.
And then I can let's just print line the circle.
That's print line the circle.

So I do circle that render remember render comes from the shape interface.

Now what we can do is we can make a red circle so a red circle is going to be a color shape where you specify appointed to the shape you're actually decorating and also the color like red for example.

And then I can once again print line I can do red circle dot render so let's render the red circle and as I run this you can see that in the first case I have a circle of radius 2 and in the second I have a circle of radius 2.

And then here is the decorated part.

So the decorator has added this additional path that the circle has and the color red.

Okay so the decorator might look might seem really wonderful but there are certain things that you lose with a decorator.

Now let's go up into our definition of circle.
So here is a circle.

Now you'll see that on the circle I've defined this method called the resize which allows you to resize the circle.

So if you're operating on the circle right here then certainly circle dot resize is not a problem you call a circle dot resize provide an argument and that's pretty much it.

It's going to I'm going to resize the circle and you'll get a circle of radius for not.

Here's the problem the problem is that once you've made a decorator once you've put colored shape over the ordinary shape what you cannot do is you cannot say a red circle dot resize because the resize method is no longer available unfortunately and there is no real solution to this because you are not aggregating anything you've lost that particular method.

The only way you can restore it so to speak is if you added again.
So if you added two color shape.
But here's the problem.

The problem is that how do you added without also adding it to the interface because remember it's only the circle type that has the resize method.

The Square type for example does not have a precise method so you cannot add this to the interface.

Unfortunately that's a real life limitation of the decorator approach.
OK so.

So that's one downside but one upside is the decorators can be composed which means you can apply decorators to decorators there's no problem doing this.

So let's make another decorator so let's have a transparent shape so a transparent shape is going to be very similar to the way that we made the colored shapes so you'll have a shape shape which is the underlying shape that you're decorating and you also have some transparency value which is going to be a flow 32 let's have a value from 0 to 1 that will normalize into sort of integer value.

Now we'll implement the render interface once again.
So this the shape interface rather with the render method.

So here what I'm going to do is I'm going to add 50 aspirins half so I'm going to say that a particular shape once again calling the underlying implementation of render has a certain percentage of transparency.

By the way the double percent is the way you get an ordinary percent as opposed to a formatting flag.

So that's how you do it.
OK.

So here we use the underlying shape so to shape that render and then let's take the transparency value and multiply it by 100.

So that transparency multiplied by hundred maybe even converted to an end.

OK so now that we have this what we can do is we can use a transparency decorator over the college shape decorator so we can apply a decorator over another decorator.

So here I can say that we have a red half transparent shape called Circle which is going to be a transparent shape.

Here I provide a pointed to Red Circle setting not point five.

Transparency and I can print line.
So let's print line a chess circle render.
Well let's take a look at what we get here.
OK so the conversion didn't work so well.

Well let's let's undo this for now and just just try it with a floating point values OK.

Here we go.

So so the last part is where we get the circle of these four and then the first decorator gets applied.

So we has the color red and then the transparency decorator gets applied and we have 50 percent transparency.

OK.
So as you can see decorators can be composed.
So you can apply decorators to other decorators again.

This does not do any kind of detection in terms of circular dependencies or in terms of a repetition.

So for example you can apply a color shape to a color shape that's not going to be a problem if you want to start detecting those things if you want to start detecting the repetition of decorator as then is going to be a lot more work and to be honest I'm not sure if it's really worth it in some situations in some very rare situations it might be worth detecting and duplicated application of a decorator but in most cases it's OK.

In those cases you just allow any layering of decorators one on top of another.

And that's absolutely no problem whatsoever.

### Decorator code: structural.composite.decorator.go

```go
package decorator

import "fmt"

type Shape interface {
  Render() string
}

type Circle struct {
  Radius float32
}

func (c *Circle) Render() string {
  return fmt.Sprintf("Circle of radius %f",
    c.Radius)
}

func (c *Circle) Resize(factor float32) {
  c.Radius *= factor
}

type Square struct {
  Side float32
}

func (s *Square) Render() string {
  return fmt.Sprintf("Square with side %f", s.Side)
}

// possible, but not generic enough
type ColoredSquare struct {
  Square
  Color string
}

type ColoredShape struct {
  Shape Shape
  Color string
}

func (c *ColoredShape) Render() string {
  return fmt.Sprintf("%s has the color %s",
    c.Shape.Render(), c.Color)
}

type TransparentShape struct {
  Shape Shape
  Transparency float32
}

func (t *TransparentShape) Render() string {
  return fmt.Sprintf("%s has %f%% transparency",
    t.Shape.Render(), t.Transparency * 100.0)
}

func main() {
  circle := Circle{2}
  fmt.Println(circle.Render())

  redCircle := ColoredShape{&circle, "Red"}
  fmt.Println(redCircle.Render())

  rhsCircle := TransparentShape{&redCircle, 0.5}
  fmt.Println(rhsCircle.Render())
}
```

## Summary

All right.

Let's try to summarize what we've learned about the decorator design pattern.

So a decorator embeds the decorated object and then it adds whatever utility fields or methods to augment the objects features.

And the decorator is actually often used to emulate something that we would call multiple inheritance in other languages in the sense that you basically acquire the behaviors and the fields of not just a single object but multiple objects.

That is what embedding does for us.

But there can be problems like for example if you embed two objects which have the same field for example a field with the same name then you have two you end up having to manage the consistency across those different objects because those two fields are going to be unique that will not be merged into a single field.

And so I've shown you the kind of problems that arise as well as how you can solve them.
