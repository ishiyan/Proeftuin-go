# Adapter -- getting the interface you want from the interface you have

## Overview

We're now going to talk about the adapter design pattern.

So an adapter is something that you should find familiar because you find adapters on sale in flight stores in duty free in all of the airports everywhere.

So the idea is very simple like in the real world.

So you have electrical devices which have different power requirements.

So for example you might have different voltage requirements or different socket types but we have European sockets and UK sockets and USA sockets and also other like Australian sockets and so on and so forth.

Now the problem is that if I have a laptop with a plug I cannot modify my laptop plug to suddenly support every single possible interface.

I cannot change it so that it works in the in the UK for example.

So European plug wouldn't go into a UK socket so it is unrealistic to modify home gadgets to support this interface.

So what we do instead.

Well we do get some support like you can support different voltages but in general the interface itself is unsupported.

So what we do instead is we use some special device which is called an adapter which is designed to give us the interface that we require.

So if you're in the UK it gives you an interface to plug into the UK socket from the interface that we actually have.

So if I have a European plug I would put the European plug on one end and I would stick the adapter into a UK socket then everything will work well with software.

It's a similar thing.

So in software engineering we have this idea of an adapter as a construct which adapts some existing interface X to conform to the required interface.

Why so suppose you use some external library which requires you to give it an interface why and you don't have it because it's not your interface.

So you take your existing interface X and you try to put some in-between layer.

So one becomes compatible with the other.

## Adapter

Many design patterns books and videos show a really abysmal implementation of the adapter design pattern that doesn't really show you the complexity of the kind of stuff that you face when you have to build an adapter for something.

So what I'm going to do in my lectures is slightly different.

I'm going to build up a scenario that's sufficiently complicated I would say it's not going to be the simplest scenario of all.

And then I'll show you how to actually build an adapter and why you'd want to build an adapter in the first place.

But we're going to start by setting up a scenario where an adapter is actually required.

Let's imagine that you're working with some sort of API for rendering of graphical objects and let's suppose that that API is completely vector based.

What I mean by that is that all of the images are defined as basically a bunch of lines and everything is built up out of line.

So we can define a struct for a line so we can define a struct and for a line you have the starting point so X1 y 1 and also the ending point so x to y 2.

So these are going to be integers.

Now let's imagine that a vector image is quite simply composed out of several of these lines so we can have type vector image struct and that vector image is going to be composed out of a bunch of lines.

So just like that.
Okay.

Now let's suppose that we're consuming this API so we're being given this API by some external developer external system and we have a function for making new graphical objects so we have for example a function for making a new Rectangle.

So you specify a rectangle in terms of width and height and you return you get a vector image a pointer has the output.

So here the implementation would be something like the following.

Now the reason why there is a minus one here for the width in the height is that typically things are zero based.

When you start them out.

So if you say an image you want an image that's with a 5 it has to go from Position ZERO to position for I'm just making this simplification right here but all we're really doing is we're making four lines because a rectangle is composed of four lines for the top part.

The leftmost part the right most part and the bottom part.
So we make these four lines.

We put them into a vector image and that's what we return.
So let's imagine that this is the interface that you're given.

Let me actually be explicit about this here.
So this is the interface you are given.

Now let's suppose that you cannot work with this interface.

And the reason why you cannot work with this interface is for example that you don't have a way of putting things graphically.

So take me for example if I want to show you a rectangle I can show you a rectangle using characters

I can really draw you a proper rectangle.

So let's suppose that the interface we have is somewhat different.

Once again gonna be explicit here so this is going to be the interface we have.

So the interface we have is going to deal strictly in terms of points not in terms of lines like line going from one point to another but we're going to do in terms of points so we're going to have a type called Point which is going to be composed of just X and Y being integers and we'll define this idea of a raster image so a roster image is an image that's defined by the points the pixels on the screen somewhere.

So we'll have roster image and we'll have this as an interface so we can have different implementations.

And the only thing you need from a roster image is a set of points so you can actually draw the image.

So we'll have a method called get points which is going to return a slice of points.

So that's how we're going to define it.

And that of course we'll have some sort of utility function for actually drawing the points because well that's how our system operates.

So we'll have a function for drawing the points so define it as draw points where you provide the roster image let's call it owner.

And we return the string like so.

So we're essentially drawing a bunch of points but we're not drawing them in a graphical way.

Instead we're drawing them using just characters on the screen sort of simplification.

Now the implementation of this code is going to be rather tedious.

So once again I'm not going to bore you with the gory details of how this is done.

Let me just paste in this huge chunk of code.

So basically what we're doing is we're allocating a two dimensional array of runes effectively just characters that you can fill in and then we go through every single point and we make sure to put a star where a point is where it actually is added and then we just use the strings builder and we put everything and just just get one big string with line breaks in it.

So that's how we can take a roster image and we can convert it into effectively a string represent patient which I can actually print out so I can print this to the console and have you look at it.

Okay so this introduces an obvious problem in the entire system because we are given the interface right here so the only way to create a rectangle is by making a vector image but unfortunately the only way to print something is by providing a roster image.

So what do you need in this case.

And the answer is well obviously you need an adapter you need something that takes a vector image and somehow adapts it into something which has a bunch of points in it so that those points can be fed into the roster image and subsequently fed into the draw points function.

OK so let's let me first of all write down what what we actually want to achieve.

So we want to make a new Rectangle.
So I make a new Rectangle new Rectangle let's say six by four.
Keep it simple.

And then what I want to do is I want to do something like F.A. dot print.

And here I want to call draw points because draw points is precisely the function that actually turns an image representation into a bunch of points.

But unfortunately I cannot feed in a rectangle.

What I have to do is I have to convert a rectangle into something which is actually a roster image.

So somewhere in between I need an adapter.
So I need an adapter.
Let's call it vector to raster.

So I'm going to feed in my rectangle into this adapter and then I take this adapter and I stick it into draw points.

So that's how everything is going to work.

And now we need to define this construct this vector to raster adapter this precisely the adapter pattern.

Basically something that's going to be able to take all the lines because remember here we are operating on lines.

Take every line turn it into a set of points and then return the sum total of those points.

Sounds easy right.
OK.

So let's try to solve this problem so let's put a solution.

This is where we're going to try and solve it So essentially we'll have an adapter.

But remember the adapter that's being passed here basically has to be just any any raster image.

So we want an adapter to implement the roster image interface in the adapter itself doesn't have to be public so I can have a type called vector to raster adapter and struct like so it's going to be lowercase is going to be private.

So I'm not going to expose this and this adapter is going to have just a bunch of points.

So there's gonna be a bunch of points in it that will be generated from the lines that we consume inside this adapter.

OK.
So what do we need.
We need some sort of way of.
First of all getting the point.

So let's generate an implementation of the interface which in our case is roster image.

So here we go.

This is our vector to raster adapter and we need to provide points somehow that's the use of pointing and that's rather easy because this is simply veto points like.

So remember we are.

We have these points calculated somewhere and we simply return them.

But the real issue here is how this thing gets constructed because well we're keeping it private so we need some sort of factory function to actually construct a new vector to raster adapter.

So let's have a function called vector to raster where you provide the vector image you want to convert and you get the roster image as the output like so.

So this is where we create the adapter
like so at some point begin to return this adapter.
Obviously we're returning the adapter as a roster image.

So we're returning the interface but in between what's going to happen is we are going to take every single line from the vector image and use the adapter to generate a suitable set of points.

So I'm going to say for underscore a line in the range of the idle lines remember if the eye is our vector image and we can say lines here actually that let me just check that for a second did I misspoke.

Yeah that should be lines multiple.
Okay.

So in range of the idle lines what we do is we take the adapter and we say well let's take the adapter let's add a line.

We don't have this function yet but we're going to add it in just a moment.

So for every single line the adapter is going to add a line.

So it's going to take that line and turn it into a bunch of points and then we return the overall thing we return the overall adapter.

OK.

So another thing that I have to fix a little bit is that to here and get points I used a pointer and if we look at the definition of the interface Well I think we're gonna be fine.

I think we're going to be fine for the moment.
So let's.
Let's get rid of this.

I'm going to get rid of his point for now let's just have get points like this.

OK.
So what's going on with our setup.
We need to implement add line.
So we need to implement some way of converting from a line.

As you know the starting and ending coordinates to every single point in this line or at least as many points as we can reasonably create.

So let's add a function here so we can have a method defined on vector to raster adapter is gonna be called add line.

So you pass a line line line like so.
And what's happening here.

So we need to take the line we need to decompose it and set up a bunch of points which once again I'm not going to go into the actual implementation of the algorithm.

It's not really important.
But let me just drop it down here.

I'm also using a function called min max that i need to add and this is a particularly sad function because the reason why this function exists is just finding the minimum and the maximum of two values and returning those is the reason why this function exists is because of the lack of ternary operator in the Go programming language.

But anyways coming back here what this function what this method rather does is given a picket line it basically decomposes the line so it looks at the left right top and bottom parts of this line it finds the edges so to speak and then it fields fills in the points.

So we have this collection of points and we just do an append.

So we basically generate a set of points for every single line and at the end we made some points and we can even sort of printout that we've generated a bunch of points and we can we can say how many exactly we generated.

So this is this is how you build an adapter basically so we've just built something which takes one API.

So that's the API related to vector images and we've adapted it to a completely different API that only wants to deal with points only wants to have some sort of get points thing that that works.

So let's actually try running it just just to show you that it does in fact work.

So first of all we wanted a six by four rectangle as you can see here we got exactly that.

And we also got some output here.
So as we went through every single one of these walls.

So we went through this wall and this wall and this wall and this wall we generated six points then it went up to 10 points 14 points and finally up to 20 points.

I means we'll rename this to something more sensible like we have because it's a running total.

OK.

So this is how you implement an adapter basically it's it's it's a reasonably simple idea.

So you're given one API in this case where given this API we have this API to work with but we have a completely different API.

So in our case we have draw points.

It takes a roster image and we're given a vector image problem we need to make a bridge somehow make a connection and that connection is an adapter.

So you basically build a new type or in this case with built a type and we also kind of exposing it using an interface.

You build this type have some sort of factory function for actually making the thing and then you use that instead.

So this extra line performs the adaptation from something which is a vector object into something which is just a bunch of points so where a roster object that's really all there is to building an adapter.

## Adapter Caching

OK so there's one thing that you need to be aware of when you're building these kinds of adapters like the adapter that we've built just now.

And that problem is the creation of too many temporary objects.

Now remember in order to actually draw our lines as pixels we turned every single line into essentially a bunch of points and these points were added here and here.

Now it's not really a big deal it's not a big issue.
We just created a bunch of points.

We store them somewhere and then we supply those points to whoever was interested.

It becomes a bit of a problem if you try to do it more than once.

Like for example let me let me just you know if I replicate this line right here and I say well let's make another adapter.

So this is a second adapter and let's actually run this let's take a look at what we get right now.

OK.

So if I if I kind of scroll up here you can see that we have six points 10 14 20 and then here we go again 6 10 14 and 20 points.

Now my question to you is this Are these operations necessary.

So so the first set of operations on a six by four rectangle is it makes sense.

We need those points but then we went ahead and we regenerate those points.

Once again that was unnecessary we didn't have to do this again because we already have the points up here.

Why did we regenerate those points.

Well it's obvious we made yet another adapter so you might say well it makes sense.

You know you might modify one adapter and the other adapter stays the same.

So why do we keep extra data.

But you can avoid making this extra data if you don't need it.

For example if you assume that your adapter is immutable meaning nobody gets to edit the internal state of the adapter then it makes perfect sense to implement some sort of caching so that this kind of thing doesn't happen so that we don't get this ridiculous duplication.

So how can you implement this kind of caching what is the what is the algorithm for implementing this kind of thing.

Well the simplest thing you can do is you can just build a very simple cache so I can make a point cash.

So the point cache is basically going to be a map is going to be a map from a special calculated value the hash of a particular line mapping to a set of points that this the points that represent this line.

So it's going to be a map from a sixteen byte array.

Now the reason I have a 16 by the way is because I'm using MDG 5 to actually perform the hashing and it's going to be a map on to a slice of points like so.

So this is our cash and now we need to change our ad line function so that it doesn't add those points if they've already been generated.

So let me just go ahead and I'm going to copy this entire function and I'm going to duplicate this.

So we keep the original but we're going to have yet another implementation let's call it ad line cached.

So Adeline cash is going to be slightly different because before adding a line before generating the points for a line and we're going to calculate the lines hash.

So here I'm going to say a hash is equal to and let's just let's just define it in a function.

So we're going to have just a function which takes anything so interface and returns is 16 by the way because like I said that's what MDG 5 gives us.

So a 16 by the way it will be returned now the first thing we can do is we can take the object that is provided into this function.

So the object right here and we can just use Jace on two to write it to a string basically or write it to a set of right.

So I can say bytes comma underscore is equal to json.Marshall.

And here I can just provide this object and I'll get a string representation and then what I can do is I can do MDA five dot sum and I can have that string representation being turned into in calculating this empty five hash.

So here I provide the bytes and I get the 16 by the right.

So this is my hash function and now I need to use it obviously.

So how am I going to use it.
I'm going to take this line.

And before I allow this line to be processed I'm going to calculate the hash value for this line.

So I'm going to calculate the hash value here and then I'm going to say well if this hash value is already in our map remember we made that cache map.

If it's already in the map then let's just use the points that have already been calculated.

Let's not recalculate them again.

So here I'm going to say if PTSD comma okay is equal to point cash at position H.

So if we have the entry with the key Hage if everything is OK then we're going to go ahead and use those values so for underscore karma points in range of points.

Let's actually just take eight points and let's append to them.

So let's append this particular point.

So instead of calculating the points again we have these points already pre calculated.

So we simply add them to the adapter and then we can we can return right here.

So if the Hash was found that we returned from this position and none of the other code none of the code down here actually gets you execute because it's unnecessary.

OK.
So we execute all of this.
We execute all of this.

Only if we don't have an entry in the cache then of course we need to add it to the cache.

So we say point cash at H is equal to eight points.

So if we haven't added those points to the cache yet we need to do it now and then we can print out that we have so many points.

Okay.

So with this setup all we need to do is we need to make sure that we are using the new function so add line cached here and let's run this once again with two adapters being created and let's see if we get the same output as before hooray.

As you can see we have six points 10 points 14 points 20 points and then nothing.

Nothing else happens because even though we try to do the whole thing twice for the second adapter right here it does matter because on the second run what happens is we already have those lines cached so we already have those points generated.

We simply copy over the points and we are done and you can you can improve the situation even further by not storing those extra points by using points pointers instead or something like that you can really improve the situation in terms of just data storage the data being generated but at least in our implementation we used caching to make sure we don't do any extra work.

We don't do any extra calculations and so you can save on processing time you can save on memory and this is just something to watch out for because essentially not every single adapter is going to generate temporary objects temporary data but if you do generate temporary data like we do here we generate temporary points then it might make sense for you to investigate how you can make sure that this data isn't getting generated redundantly like for example when somebody makes the same adapter twice on the same object that you avoid generating the same data avoid doing the additional storage and additional calculations.

### Adapter Caching code: structural.adapter.adapter.go

```go
package main // structural.adapter.adapter
import (
  "crypto/md5"
  "encoding/json"
  "fmt"
  "strings"
)

func minmax (a, b int) (int, int) {
  if a < b {
    return a, b
  } else {
    return b, a
  }
}



// ↑↑↑ utility functions

type Line struct {
  X1, Y1, X2, Y2 int
}

type VectorImage struct {
  Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
  width -= 1
  height -= 1
  return &VectorImage{[]Line {
    Line{0, 0, width, 0},
    Line{0, 0, 0, height},
    Line{width, 0, width, height},
    Line{0, height, width, height}}}
}

// ↑↑↑ the interface you're given

// ↓↓↓ the interface you have

type Point struct {
  X, Y int
}

type RasterImage interface {
  GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
  maxX, maxY := 0, 0
  points := owner.GetPoints()
  for _, pixel := range points {
    if pixel.X > maxX { maxX = pixel.X }
    if pixel.Y > maxY { maxY = pixel.Y }
  }
  maxX += 1
  maxY += 1

  // preallocate

  data := make([][]rune, maxY)
  for i := 0; i < maxY; i++ {
    data[i] = make([]rune, maxX)
    for j := range data[i] { data[i][j] = ' ' }
  }

  for _, point := range points {
    data[point.Y][point.X] = '*'
  }

  b := strings.Builder{}
  for _, line := range data {
    b.WriteString(string(line))
    b.WriteRune('\n')
  }

  return b.String()
}

// problem: I want to print a RasterImage
//          but I can only make a VectorImage

type vectorToRasterAdapter struct {
  points []Point
}

var pointCache = map[[16]byte] []Point{}

func (a *vectorToRasterAdapter) addLine(line Line) {
  left, right := minmax(line.X1, line.X2)
  top, bottom := minmax(line.Y1, line.Y2)
  dx := right - left
  dy := line.Y2 - line.Y1

  if dx == 0 {
    for y := top; y <= bottom; y++ {
      a.points = append(a.points, Point{left, y})
    }
  } else if dy == 0 {
    for x := left; x <= right; x++ {
      a.points = append(a.points, Point{x, top})
    }
  }

  fmt.Println("generated", len(a.points), "points")
}
func (a *vectorToRasterAdapter) addLineCached(line Line) {
  hash := func (obj interface{}) [16]byte {
    bytes, _ := json.Marshal(obj)
    return md5.Sum(bytes)
  }
  h := hash(line)
  if pts, ok := pointCache[h]; ok {
    for _, pt := range pts {
      a.points = append(a.points, pt)
    }
    return
  }

  left, right := minmax(line.X1, line.X2)
  top, bottom := minmax(line.Y1, line.Y2)
  dx := right - left
  dy := line.Y2 - line.Y1

  if dx == 0 {
    for y := top; y <= bottom; y++ {
      a.points = append(a.points, Point{left, y})
    }
  } else if dy == 0 {
    for x := left; x <= right; x++ {
      a.points = append(a.points, Point{x, top})
    }
  }

  // be sure to add these to the cache
  pointCache[h] = a.points
  fmt.Println("generated", len(a.points), "points")
}

func (a vectorToRasterAdapter) GetPoints() []Point {
  return a.points
}

func VectorToRaster(vi *VectorImage) RasterImage {
  adapter := vectorToRasterAdapter{}
  for _, line := range vi.Lines {
    adapter.addLineCached(line)
  }

  return adapter // as RasterImage
}

func main() {
  rc := NewRectangle(6, 4)
  a := VectorToRaster(rc) // adapter!
  _ = VectorToRaster(rc)  // adapter!
  fmt.Print(DrawPoints(a))
}
```

## Summary

OK so let's try to summarize what we've learned about the adapt her design pattern so implementing an adapter is rather easy.

Basically you determine the API that you have and you determine the API that you need for the whole thing to actually work and then you create some sort of a component which aggregates or has a pointer to the ADD T.

And then you have the the whole thing provide the appropriate data and sometimes yes sometimes you'll see a situation where you have intermediate representations and they can pile up and in this case you have to use caching and other optimizations to make sure that the amount of temporary data that you're generating as you're providing the adapter is manageable and it doesn't go out of bounds.
