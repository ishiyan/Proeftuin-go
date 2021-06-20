# Bridge -- connecting components together through abstractions

## Overview

We are now going to talk about the bridge design pattern.
So what is this all about.

Well a bridge design pattern solves a particular problem and that problem is what's typically called a Cartesian product complexity explosion.

So where can this come from.

Well let's take a look at an example let's suppose that you want to have a thread scheduler but in actual fact what you need is you need a whole family of these threat schedules because threat schedules can be preemptive or call primitive and also through our schedules can run on either Windows or Linux.

So if you just tackle this problem straight on you'll end up with a two by two scenario.

So you'll have four different structures you'll have a Windows preemptive schedule a Unix preemptive scheduler a Windows cooperative and a Unix call price schedule.

And this is just a two by two if you had let's say three operating systems and three types of schedulers you would end up with nine different structures which is completely unmanageable and something that you really want to avoid in code if at all possible.

And this is what the British design pattern is actually for.

It tries to avoid this complexity explosion.

So if we look at the scenario initially you would represented something like this.

So you would have the four different chandeliers at the bottom of this diagram.

And what you would do instead is you would try to make sure that you have some you have some top level construct like a thread scheduler which in turn can be aggregated inside a pre-emptive scheduler or a corporate scheduler.

But on the other hand you also have a reference or a pointer or some sort of inclusion of a platform scheduler.

And this would in turn be implemented in the schedules for the different platforms.

So essentially the idea isn't particularly complicated there's just this idea that if you are facing this complexity scenario where you have an eight times B times C whatever then you try to manage this using just separation of hierarchies if you will.

So that instead of having one huge tree you have a couple of flat trees or lists a hierarchy like lists.

So the bridge is nothing more than a mechanism that decoupling and interface or an interface hierarchy from an implementation or implementation hierarchy.

## Bridge

So imagine that you're working on some sort of graphical application then this graphical application should be capable of printing different objects like circles and rectangles and squares and that sort of thing.

However it needs to be able to render them in different ways.
You want one to render those shapes in vector form or in roster form.

So if you were to sort of let this situation out of control you would have lots of different cross products of the different shapes so if you have shapes like circle and square and if for example you have let's say a roster renderer and a vector renderer then if you plan your application badly you'd end up with lots of different types you'd have things like circle like a roster circle let's say across the circle a vector circle across the square and vector square and so on and so forth thing you can imagine that if the number of types in any of these groups is larger than 2 then you would just have an explosion of all the different types.

So the question is how can we simplify this how can we actually reduce the number of times that we need to introduce and the answer is that you would typically do this using the bridge design pattern.

So let me show you how this works.

So instead of having shapes specialized for different render as we would take those renderings and we would maybe introduce an interface like type renderer interface and inside this renderer we would have a definition of methods for rendering different kinds of shapes.

For example if you have circles in your application you would have a render circle method which would take some radius for example.

So and similarly you would have like round the square and ran the triangle and that sort of thing and now what you can do is you can define types which take care of the rendering in a different format.

So for example for the vector rendering you would have a type vector renderer which would be a struct I'm not gonna put anything in here you can imagine that this would have any kind of utility information related to how you want the vectors to be constructed and then you would have a function for actually rendering a circle so we would implement the the renderer interface inside a vector renderer.

So we would do it like this.

And here what we would do is while in our case since I'm doing in demo I'm just going to make a bit of output so I'll just F.A. print find that we are drawing a circle of radius and I'll specify the radius.

That's all that we're going to do here.
So this is a vector renderer.

But similarly we can have a roster renderer which also knows how to render a circle but it knows how to render it in a different way so we'll have type roster renderer which is going to be a struct and once again you can have all sorts of different options here like for example if you're rendering a shape in terms of pixels you need to provide the API for that shape.

So that would be maybe one of the fields that you would have in here.

And once again you would implement this render interface so here we would implement the the renderer interface on this.

So we'll just do something very similar.

So I will just go F.A. print line and here I'm going to say that we're drawing pixels for a circle of radius and then specify the radius here that that's all that we're going to do for the demo.

Okay.
So now we can define the circle.

So you'll notice that we've just defined an ordinary circle will not define like a vector circle in the Rasta circle and all the rest of it will just have a circle which subsequently refers to or it has a bridge to the renderer.

So we'll have type circle struct.

And then if we want to have rendering handled for us all I'm going to do here as I'll just have a reference to the renderer.

And in addition I'll have the radius as a float 32 here.
Okay.
So now we have a circle and we need some sort of functionality on this circle.
First of all let's have some sort of factory function.
So maybe I can just use the idea to generate it.
So let's generate generated like this.
Use circle where we'll have the render and the radius.
Just just the default implementation of the whole thing.

And then we'll have some functions for let's say drawing the circle and you'll notice that this is precisely where we would if you if you have the separation into like across the section and a vector circle this is where you would do these sort of specific kind of rendering but here all we have to do it because we've defined a rendering member here we can use this render remember to actually render the thing that we want.

So here I would say see Dot render Dot and I would say render circle and we would render it with a particular radius.

So speed up radius like this and also you can have other like for example for resizing the circle just to show you how that would look.

So if you want resize by some factor which is a flow 30 to see that radius multiply equals factor. We're going to use this in just a moment.

So let me show you how the whole thing works.

So essentially what you need to do to get everything to operate is you need to make a renderer which is a separate component and then you provide that render into the circle.

So you sort of introduce it as a dependency.

So here we can have a roster renderer like so or we could have a vector render so it's really up to us which one we want to use and then when we create the circle we create a circle providing.

So we call a new circle and we provide that render in there.
So for example the roster render and we provide a radius of five.

Now this has to be a pointer so let's put an ampersand in here and then I can do circle draw and then I can for example also resize resize let's say doubling the size of the circle.

So resize by a factor of two and then circle draw again.

So now we're using the roster and the rows all comments on the vector render for a moment and we can see how this works.

So as I run this off we get the output so drawing pixels full circle of these five and then we're drawing pixels of a circle of radius 10.

If I now do it like this.
So if I now use the Vector Ender and now we'll have a different output.
So here we have just drawing a circle of these five drawing a circle various 10.

So as you can see what's happening is we've avoided this complexity explosion we've avoided this too by to set up of classes because obviously if you have let's say if you have two different renderings in two different shapes you have four different potential types that you want to you might want to make although not really.

And we've avoided this whole thing by essentially introducing a sort of dependency here on the render and then just reusing that renderer to actually render something of course what this implies and this is not the best of things.

Is that when you introducing you shape like a triangle for example you would by necessity introduce a new rendering method similar to the one that we have here so the renderer interface would have render square and rendered triangle and all the rest of it.

And then of course because you want this interface to be implemented by both vector render and roster render it it would result in a cascading set of functions or methods rather being created on those renderer.

So if I introduce a render square here for example I would have to add render square to a vector render unto a roster and or add any kind of renderer that I have in my system.

But this is the price to pay for the additional flexibility and hopefully you can see that this is better than allowing a complexity explosion of an infinite number of types.

### Bridge code: structural.bridge.bridge.go

```go
package structural_bridge

import "fmt"

type Renderer interface {
  RenderCircle(radius float32)
}

type VectorRenderer struct {

}

func (v *VectorRenderer) RenderCircle(radius float32) {
  fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
  Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
  fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
  renderer Renderer
  radius float32
}

func (c *Circle) Draw() {
  c.renderer.RenderCircle(c.radius)
}

func NewCircle(renderer Renderer, radius float32) *Circle {
  return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Resize(factor float32) {
  c.radius *= factor
}

func main() {
  raster := RasterRenderer{}
  vector := VectorRenderer{}
  circle := NewCircle(&vector, 5)
  circle.Draw()
}
```

## Summary

All right let's try to summarize what we've learned about the bridge design pattern.

So the idea here is that for the bridge we decouple the abstraction from the implementation.

I know it sounds a bit scientific like like what are we doing here.

But that's really what's happening behind the scenes so we decouple some of the abstraction from the implementation and we create a parallel hierarchy for that and thus we prevent the the whole complexity explosion and both the abstraction as well as the implementation they can exist as separate hierarchies meaning you have one hierarchy which also makes use of another as opposed to this vector product kind of approach you can think of the bridge pattern really as a stronger form of encapsulation.
