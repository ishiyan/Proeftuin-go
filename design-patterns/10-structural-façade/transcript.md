# Façade -- Exposing several components through a single interface

## Overview

Now let's talk about the façade design pattern.
So what is this all about.

Well you really want to balance complexity and also presentation or usability and this works not just in the software world but it also works in the real world.

So for example if you are buying a home now a typical house has is really complicated has lots of different subsystems.

There is electrical There's plumbing there's ventilation there's there's lots of stuff happening in the House but on the other hand this complex internal structure like all these multilayered floors and walls or whatever they the end user doesn't care so much about them.

So when you're buying a house you're not exposed to all the Internet was you just want to know that the House is generally in good shape but you don't care about the exact layering of the floors for example.

So this is the same with software.

Sometimes you have a complicated system or indeed a set of subsystems behind the scenes but you don't really care about anything you just want to have the flexibility you want to have the a nice API.

You just want some surface level API that will work for you.

We'll set up many of these systems behind the scenes and get them working together without you knowing about all the internal details.

So the façade design pattern is basically a way of providing a simple easy to use interface so a set of set of methods or functions over a large and sophisticated body of code.

So behind the scenes you could have lots and lots of struct and functions and whatnot but the façade can provide you a simple API where everything is more or less self descriptive and intuitive and you don't have to know about all the complicated internal details.

## Façade

The idea of a façade is that you hide a very complicated system behind a very simple interface.

Unfortunately we don't have time here to build a complicated system but I do want to show you a more real life example of something that I personally build and that is the idea of a multi buffering multi view board terminal.

So you know about consoles and terminals they typically present just a row of letters rows and columns of letters and black background white letter is nothing particularly fancy that is what we end up looking at.

Every time we do a print line like when when I do a print line and I said hello to the program essentially right here you're looking at a black background with letters but sometimes you want more than one buffer to store the text being output so imagine if I was running several programs at the same time you would want to take this into space and somehow split it into different parts.

So you have different viewpoints which in turn are attached to different buffers but you still want to work with a console and in as much as simplified manner as possible you want a simple API over something that's complicated.

So if I were to build such a thing I would start with making a buffer so I would have a buffer which would just store a bunch of letters.

So the buffer would have some width and height as integers and it would store essentially just on the array of runes.

So we can define it like this.

So this would just be an array then of course you could have a constructor which initialize is the width and height and you would also initialize the the size of the buffer depending on the width and height so obviously if you have the width and height defined what you can do is you can initialize the buffer just by calling make saying that you want to run array and the number of elements you want is width multiplied by height.

So this is how you would initialize the buffer and have a factory function which constructs the buffer and then you could have for example some utility method for getting a character at a particular position in the buffer.

So here if we take a buffer.

So this is going to be a method on the buffer called at where you specify some index as an integer and you return the ruin.

This is simply a look up into that slice of virus over it on the buffer at position index.

So this is one of the components of our rather complicated system and then of course we need to present this buffer on the screen and remember a buffer can be really large it can be hundreds of lines long if you want to preserve the history of the inputs to a particular terminal but you can only show a part of that buffer and for that we have a new construct called the viewport.

So a viewport basically shows you just a part of the buffer as a particular offset starting at a particular line.

So here you have a pointer to the buffer that the viewport is attached to and then you have an offset which is just some integer to some offset from the start of the buffer.

So once again what we can do is we can make a factory function which initialize as the buffer in the following fashion and then we can have another utility method for getting a character at a particular position incorporating the knowledge about this offset member of ours.

So you would have something like a function the viewport where you would call get character at a particular index and you would return a ruin.

So here you return.

You take the buffer first of all and then you take a character add a position defined by the offset plus the index that the person is interested in.

That way they get the character from the start of the visible area as opposed to the start of the entire buffer.

So as you can see we have a situation where we have buffer and viewport and you can imagine a console a multi buffer console being a kind of combination so you would have lots of you ports and you would

have lots of buffers but you also want a simple API for just creating a console which contains all of these view ports and buffers behind the scenes and that is where you would build a façade.

That is where you would build some type called console.

So a console is going to be a struct which is going to incorporate information about the different buffers.

So that's going to be just an array of pointers is going to have a bunch of viewed ports.

Also an array of pointers.

And once again you can have some offset somewhere for the current the current viewport or something to that effect.

Let's actually make a factory function for initializing this console.

We were not actually going to put any buffers in here.

What we're going to do is we're going to have an initialize there which creates a default scenario.

Now a default scenario is where a console has just a single buffer and a single viewport.

And if you look at terminals in Windows and Mac OS and Linux the default implementation of a console or a terminal is the one where you have just one buffer and one viewport not very exciting but this is the kind of simplified API that you would expect a façade to provide so here you would have a function called new console where you would return a console which is initialized with a single buffer and a single viewport so you make a buffer new buffer.

Let's say two hundred by one hundred and fifty for example.

You would make a new viewport attached to that buffer so you would call new viewport passing in the buffer and then you would return a console where you would initialize the buffer array and they viewport array.

So you would say let's make a new buffer array with the buffer you would make a new viewport array with the viewport and then you would specify the offset or any additional parameters that you might need.

And once again now that you have this console you can also have a high level get character at a kind of function for figuring out a character's position at a particular point in the console.

Now we know that to do this you have to grab a buffer and look into that buffer.

And for grabbing that buffer you might want the viewport because remember a viewport has this get character at a function so you would write something like the following you would add a method on a console called

Get character at add a particular index.

Once again this would return a ruin and here by default you would look into the first viewport or you might have a selected for the current view for the viewport where the user is currently focused and you would call get character at providing that index.

So the way you would use this is instead of working with low level constructs like buffers and consoles and buffers and view boards and whatever you would use a console so you would say see is a new console just like that.

And this creates lots of stuff behind the scenes it creates a default viewport attached to a default buffer and there's only one of either and then you can get a character add a particular offset which in turn does lots of things behind the scenes.

So the idea of a façade is basically providing a simple API over something that's complicated.

So here if you were to work with buffers and view ports manually would have to manage them yourself.

And it's really a lot of work and in most cases most people need a an implementation where there is just a single buffer and just a single viewport and that's exactly what we are providing as part of the new console factory function.

So that is where we provide a default implementation.

And of course just because you're making a façade doesn't mean that you have to obscure the inner detail.

So if somebody wanted to mess about with the buffers in the viewport they can do it through the console.

There is no problem in doing that or if they want to they can use the viewport type and the buffer type directly without even working with the console.

But of course the console is here if you need it.

It's a simpler way of working with this entire ecosystem and that's the goal of the façade.

Basically just to make things easier to use.

### Façade code: structural.composite.façade.go

```go
package facade

type Buffer struct {
  width, height int
  buffer []rune
}

func NewBuffer(width, height int) *Buffer {
  return &Buffer { width, height,
    make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
  return b.buffer[index]
}

type Viewport struct {
  buffer *Buffer
  offset int
}

func NewViewport(buffer *Buffer) *Viewport {
  return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
  return v.buffer.At(v.offset + index)
}

// a façade over buffers and viewports
type Console struct {
  buffers []*Buffer
  viewports []*Viewport
  offset int
}

func NewConsole() *Console {
  b := NewBuffer(10, 10)
  v := NewViewport(b)
  return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

func (c *Console) GetCharacterAt(index int) rune {
  return c.viewports[0].GetCharacterAt(index)
}

func main() {
  c := NewConsole()
  u := c.GetCharacterAt(1)
}
```

## Summary

Let's try to summarize what we've learnt about the façade design pattern.

So to build a façade.

Well the idea of building a façade is to provide some sort of simplified API over a set of components and those components can be quite complicated.

They can have lots of details and parts.
Basically just just a really complicated system.

And we may also wish to optionally expose those internals through the façade so the façade can try to hide them as much as it can.

But sometimes we might want to expose those internal details so that if we have a power user somebody who really wants to understand what's going on they can also manipulate those implementation details.

So sometimes you would have simple understandable APIs but you would also allow users to escalate the use of more complex API.

So for example you would have functions which take additional parameters where you can specify the advanced options shall we say and it's up to the client whether or not to use those or just just leave everything and the kind of understandable level or go in deep and customize this system to their heart's content.
