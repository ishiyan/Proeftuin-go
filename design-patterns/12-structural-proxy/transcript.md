# Proxy -- An interface for accessing a particular resource

## Overview

Let's talk about the proxy design pattern.
So what are proxies for.
Well let's suppose you're doing something simple.

You have a struct in a variable called food and you're calling the method bar.

So you're saying that by now there are some assumptions behind the scenes here.

So this invocation assumes that food is for example in the same process as bar but does it really have to be the case.

And what if later on as you kind of you improve your system you make it smarter.

You want to put all the food related operations into a separate process.

So instead of calling it in the process that you're currently in.

You wanted the invocation of food up bar to happen let's say in a different process or maybe on a different machine somewhere on the network or maybe halfway around the world in some cloud or something.

The question is can you avoid changing a code if that is the case.

And that is where the proxy pattern can come to the rescue because the idea of the proxy is that you provide the same interface but the behavior is entirely different so you would still have a variable called Fool and it would still have a method called bar but this method would not call something in process it would call something somewhere else for example.

So what I'm describing right now is called the communication proxy but it's really only one of a number of examples of proxies that you encounter are there so you can have logging proxies and virtual proxies.

We'll take a look at some of those proxies in this section of the course but generally the proxy is a type that functions as an interface to a particular resource.

So it's kind of like a decorator in a way in that it wraps things.

And the idea is that the resource itself can be remote.

It can be expensive to construct it may require logging or something some other related functionality.

And so you provide something which gives you the same interface as the original but also does additional things kind of behind the scenes.

## Protection Proxy

A protection proxy is the kind of proxy which performs access control It basically tries to check whether or not the object it's trying to proxy is actually allowed to be accessed.

So let me show you a very simple example let's suppose that whereas simulating the process of cars and other vehicles being driven.

So we have some sort of driven interface and this interface just has a single method called Drive.

Now what you can do is you can construct a car just as an ordinary struct and not even going to add anything in here and then you can implement this interface inside an ordinary car.

So in the case of just just an ordinary car without any checks you can print line the fact that the car is being driven.

OK.

So this is your starting point now imagine that you want the car to only be driven if you have a driver and if that driver is actually old enough.

So what you do then is you build a protection proxy on top of the car.

So once again you would be reusing the car somehow but you would also be specifying the driver.

So here you would have first of all you would have a driver as a separate struct driver might have an age for example so we can check whether the driver is old enough to drive the car.

And then we can build a car proxy.

Maybe there should be a better name for this like a safe car or verified car or something in the car proxy can for example store the car.

It's really up to you whether it's a story by value as for a pointer it doesn't really matter.

You can also have the driver Driver driver.
Let's have the driver is a pointer here.

So the idea is that whenever somebody wants to making new car what they in fact get is they get this car proxy so we have a function called New Car proxy where you have to specify the driver and it returns a car proxy object and we return a car proxy Car and Driver like so.

OK.
So this car proxy just like the ordinary car it has to implement that driven interface.

So let's go ahead and quickly implement this interface here.
So what can we do.

Well in this case we need to verify that people are actually allowed to use this car so we can say that if the car driver age is greater than or equal to 16 for example then yeah you can drive the car.

So we do see the car drive. Otherwise we may be outputs some sort of error message.

So print line for example a driver too young.
Now we go.

So this is our protection proxy and we can try working with it so we can make a car which is a new car proxy.

Well you have to specify the driver so you can say the driver is for example twelve years old and then you can try driving the car this way.

And if we now invoke it we should get an error message because well the driver is obviously too young and if I go back here in a modify the age to 22 for example and I run this once again then you can see that now the car is being driven.

So this example shows a very common kind of pattern where first of all you're starting out with an object or some sort of structure which can be used as is without any verification but also subsequently you want to have additional verification additional checks being made whenever somebody actually uses this drug.

So this is how you do it.

You you can have it as a new factory function and you can also introduce dependencies.

So remember the original card doesn't have any parameters that you have to feed it.

It doesn't have anything at all whereas here you have to explicitly specify the driver without this driver well you wouldn't be able to even construct this structure in the first place.

So this is how you implement a protection proxy.

### Protection Proxy code: structural.proxy.protectionproxy.go

```go
package proxy

import "fmt"

type Driven interface {
  Drive()
}

type Car struct {}

func (c *Car) Drive() {
  fmt.Println("Car being driven")
}

type Driver struct {
  Age int
}

type CarProxy struct {
  car Car
  driver *Driver
}

func (c *CarProxy) Drive() {
  if c.driver.Age >= 16 {
    c.car.Drive()
  } else {
    fmt.Println("Driver too young")
  }
}

func NewCarProxy(driver *Driver) *CarProxy {
  return &CarProxy{Car{}, driver}
}

func main() {
  car := NewCarProxy(&Driver{12})
  car.Drive()
}
```

## Virtual Proxy

A virtual proxies the kind of proxy that pretends it's really there when it's not necessarily.

Let me show you how it can work.

Let's imagine that we have an interface called image so an image is something that you can draw.

And in our case we're just going to emulate the process of drawing an image.

So imagine you want to have a bitmap.

So you load the bit up from a file and then you draw that bitmap so we can have a type called that map.

And we can give the bitmap a file name.
Let's do it lowercase file name string.
OK.
So you can construct a bitmap and then you can draw the bitmap.
So func B that map draw.

So when you draw the bitmap let's just emulate the whole thing.

So just print line the fact that we are drawing an image and then here is the file name.

OK so let's have a constructor for the bitmap which simply initialize as it with a file name.

But in addition what we're going to do is we're going to specify here that this is the point where we actually load the image from some file so we load the image we loading the image from and then specifying the founding.

We obviously need the image to actually construct the bitmap.
OK.

Now let's imagine that somewhere down below we have a function for actually drawing some image.

So you pass in the image interface and you get to draw that image.

So we'll have a function called draw image where you specify the image and then here's what we're going to do we're going to output a bunch of diagnostic calls.

So we're going to say that we're about to draw the image.

So we're about to draw the image then we're going to perform the actual drawing and then we're going to say we'd done drawing the image.

And in between we take the image and we call draw on it.

OK.
So far so good.

And what we can do is we can make a bit map and we can feed that bitmap into this draw image function.

So I can say BNP is a new bitmap with demo dot PMG and then we can do BNP a draw.

OK.

So let's actually run this let's take a look at what it is that we're getting here so you can see where loading the image and we are drawing the image actually let's feed it into this draw image function because that's better because we get to see some additional output so let's try this again and this time round we'll have more information.

So as you can see the first thing that happens is we're loading the image and then we're using the image.

So we're about to draw the image we draw the image and we are seeing that were done during the image.

Now what is the problem with this scenario.

The problem is what happens if you never draw the image in the first place if you never actually invoke this.

So imagine you have the invocation of new bitmap just going to put it into an underscore and that's it.

You never actually need the image.

Now if you run this you'll see that there is a fairly obvious problem that we are still loading the image even though we never draw it.

So one attempt to fix this might be to introduce some kind of lazy bitmap the kind of bitmap where the image doesn't get loaded until you actually need to render it.

So how would you implement this.
Well you can implement this using a proxy.

So a lazy bitmap is something that is going to wrap an ordinary bitmap.

And it's also going to implement the image interface and provide the draw method.

But it's going to do it differently.
So let me show you how this can work.

So we make a lazy bit tab and the lazy bitmap is going to store the file name but it's also going to reuse the underlying bitmap functionality because remember we don't want to re implement the process of drawing some bitmap.

So we're going to have a bitmap pointer here and we're going to use this in just a moment.

So the idea is that when you make a constructor let's have a constructor which initialize as the file name.

So when you make this whole thing you don't specify the bitmap yet because this bitmap right here is going to be lazily constructed.

What this means is that it's only going to be constructed whenever somebody needs it.

So now we can implement the the image interface on our new Struct and here what we want to do is we obviously want to use the underlying bitmap this bit app but we need to make sure that it's constructed because at the moment the pointer has a value of nil.

So if the pointer El bitmap has a value of nil then we need to construct the thing.

So we say held up that map equals new bitmap and we provide LDA file name so we specify the file name.

And now that the thing has been constructed we can do l dot file name dot Jewel there by calling Well it's El adult bitmap the draw there by calling the underlying implementation of the bitmap drawing algorithm.

OK.

So with all of this what we can do now is we can have a different so here I can say BNP is a new lazy bitmap with the file name demo that PMG has before and then we can draw image BNP.

And this time round let's take a look at what's actually going on because the order of the output is going to be different.

So you can see we don't load the image prematurely.
We are about to draw the image.
So we're inside this joy image function of ours.

Then when it's time for us to draw the image we load the image from the file then we draw the image and then we are done drawing the image.

Now of course I can go back until the source code then I can replicate this line.

I can do draw image twice and run this once again just to show you that the loathing is in fact lazy so the loathing only happens at this particular line.

And subsequently when you're about to draw the image the second time there is no loathing happening because we've already allowed that the image and all we have to do is just droid.

So the demonstration here shows how you can build something typically called a virtual proxy.

The reason why it's virtual is because when you create a lazy bitmap using the new lazy bitmap function it hasn't been materialized yet.

Meaning that the underlying implementation of the bitmap hasn't even been constructed and it's only constructed whenever somebody explicitly asks for it.

In this case asks to actually draw the bitmap.

That's when the whole thing gets constructed and subsequently used behind the scenes.

### Virtual Proxy code: structural.proxy.virtualproxy.go

```go
package main

import "fmt"

type Image interface {
  Draw()
}

type Bitmap struct {
  filename string
}

func (b *Bitmap) Draw() {
  fmt.Println("Drawing image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
  fmt.Println("Loading image from", filename)
  return &Bitmap{filename: filename}
}

func DrawImage(image Image) {
  fmt.Println("About to draw the image")
  image.Draw()
  fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
  filename string
  bitmap *Bitmap
}

func (l *LazyBitmap) Draw() {
  if l.bitmap == nil {
    l.bitmap = NewBitmap(l.filename)
  }
  l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
  return &LazyBitmap{filename: filename}
}

func main() {
  //bmp := NewBitmap("demo.png")
  bmp := NewLazyBitmap("demo.png")
  DrawImage(bmp)
}
```

## Proxy vs Decorator

So one thing worth mentioning is the difference between the proxy and the decorator because I mentioned that they are kind of similar but they really do serve different purposes.

So the proxy design pattern tries to provide an identical interface.

Now it's not always possible sometimes you will mess up the interface somewhat because for example you would have a factory function with a different name but generally the proxy in terms of the set of methods tries to provide an identical interface to whatever it is actually controlling to whatever resource it's actually using behind the scenes whereas in the case of a decorator you don't want to stick to an identical interface.

You want to enhance that interface.

You want to have more fields and more methods and should just provide additional functionality.

So the decorator also typically aggregates or has a pointer to whatever it is decorating now the proxy doesn't really have to the proxy doesn't have to use the underlying object at all.

It could be doing something completely different.

For example the proxy might not even be working with a materialized object.

So the object that the proxy sort of the object for which the proxy works might not even exist it might not even be constructed but the proxy could still be usable you could still call the proxies methods is just that the behavior would be completely different.

## Summary

So let's try to summarize what we've learned about the proxy design pattern.

So a proxy has the same interface as the underlying object and to make a proxy you simply replicate the existing interface of an object.

That's all that you have to do and then you add the relevant functionality to the redefined methods for example.

So you add whatever it is that you wanted to add there and you have different types of proxy.

So a proxy is not just a single pattern with a single purpose you have different kinds of proxies and they all have completely different behavior.

So for example we looked at the virtual proxy which lazily creates an object.

So that is an example of specific behavior of a proxy and you could have proxies other than the proxies I'm listing here.

There are lots of possibilities but hopefully you get the general idea of what the proxy pattern is actually for.
