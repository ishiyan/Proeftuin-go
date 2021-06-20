# Strategy -- System behavior partially specified at runtime

## Overview

Now we're going to talk about these strategy design pattern.

So what is the motivation.

Well many algorithms that are out there can actually be decomposed into the higher level and lower level parts.

And what we mean by that is we can define an algorithm in terms of the general strategy as well as the implementation detail.

So for example take making a cup of tea for example.

You can decompose this into the process of making a hot beverage where you just boil the water and you pour it into a cup and then there are the specifics like for example if you're making a cup of tea you have to put the tea bag into the water if you're making coffee you have to grind the coffee and put it into water and maybe add cream and whatever so you can see there is this separation between the high level stuff that happens in all cases and the low level implementation details that you need to provide.

So the high level algorithm can be reused so you can make a high level algorithm and then you can reuse it for.

In our case making tea or coffee or hot chocolate or something else.

So you basically have different beverage specific strategies so you have the high level algorithm of making a hot drink and then you have like the specific strategies like a T strategy coffee strategy and so on and so forth.

So the strategy pattern is a representation of this idea.

Basically the idea that you separate an algorithm into the kind of skeleton the kind of high level abstract definition as well as the concrete implementation steps which by the way you can vary at runtime so at runtime you can actually switch from one strategy to another.

## Strategy

All right.
Let's take a look at how we can implement the strategy design pattern.

So I'm going to do a demo where we're going to print a list of textural items and we're going to do this printing using different formats.

So we're gonna have a type called output formats and we're going to have some sort of text process which takes a list and it can print it using either markdown or HDMI.

So we'll have some constants for both of those.

So we'll have markdown and we'll have HDMI as well.

So the idea is that you have some strategy for how to print a list.

We can define a time for this.

We can have a type called list strategy which is going to be an interface.

Now the idea is that whenever you printed list there is typically three things that you need to know about how to print a list.

There is sometimes a start of a list.

So in the case of the HDMI for example you're going to have that you al element.

Then you have every single list item.

And this once again depends on what what kind of format you're using.

And then you have the closing of the list.

So there are three elements which means that we have to have three parts of our strategy we'll have a method called start.

Now all of these methods are going to take a strings builder so that we get to accumulate the text as it's being output.

So here I have builder as strings builder like so I'll say I'll have the start method.

I'll also have the end the method.

And in addition I will also have the add list item.

So add list item is where you add the individual elements and here in addition to the builder itself you also need the item.

So we're going to have the item as a string.
OK.

So now that we have this we can build different strategies for constructing lists using the markdown format and using the 80 mile format then they're going to be different they're going to be significantly different.

So first of all let's implement the markdown strategy.

So markdown list strategy.
So this is gonna be an empty struct.
There's not going to be anything here.

Typically you would have some service information here or some formatting setting specifically for markdown maybe.

But we're going to keep this empty.

All I'm going to do here is implement the methods of the list strategy interface like so and then we get to look at what's actually going to happen in here.

Now the first thing to note is that when you're making lists of markdown elements they're typically done like this.

So you have one and two and so on and so forth.
So there is no preamble.

There is no start of the list and end of the list like we have an HDMI.

So both the start and add methods are basically going to be empty which means we simply get rid of any code that's in here and we just let them be everything that has to happen happens inside the add list item.

So this is where we get to we get to do the following.

So we say builder dot right string and here we write the string so we put that asterisk before every single item.

Then we add the item and then we add the line break.

So that's all there is to it and now we can implement also the HDMI list strategy so we can have a type called a CML list strategy.

It's going to be a struct and then once again all I'm going to do here is I'll implement the men methods of the strategy interface or yes the interface of course.

And then for the start in the end we are actually going to have concrete code because we have to have that opening and closing list item HDMI.

So here I'll say builder dot right string and I'll specify unordered list you l followed by a line break and I'll do the same thing here except here is gonna be the closing tag as opposed to the opening tag and then let's not add those.

And then finally when it comes to adding the list item we're going to just add the item with a bit of indentation.

So here I'm going to say builder dots right string.

So some indentation then list item.

Then I'll have the item itself then I'll have the closing tag for the list item as well as a line break.

There we go.

So this is how you can write both HDMI Mel as well as markdown.

And now let's imagine that we have some sort of text process.

So the text processor is this configurable component which you can feed this component to a list and specify the strategy that you want to take.

So we'll have a text processor as a struct now we're going to have a builder.

This is gonna be the strings builder that we're gonna be using and passing it to all of those methods.

And in addition we'll have the list strategy.

So this strategy is going to be of type list strategy.

This is the strategy that we're going to take when we actually go ahead and write the list.

So now let me make a constructor and now in the constructor we only pacify the list strategy we don't specify the the builder so in terms of the builder here you would simply making you strings builder and then specify the list strategy.

So this is how you would initialize the thing and then of course in addition to being able to set the list strategy at the beginning meaning you set it right here in the factory function well you can also do as you can switch from one output format to another.

So once again it's really up to you how to do this here but you'll notice that when we got started I defined a bunch of constants here and we can use these constants to force a switch from one strategy to another so that means that somewhere inside the tax processing you would have a method and this method would set the output formats depending on the output format that you specify.

So here we just switch on the format and if it's marked down and then we simply replace the list strategy member so here we say t that list strategy equals markdown list strategy and like so and if it's HDMI L then t that list strategy is equal to HCM l This strategy like so.

So that's all there is to it.
And of course now the critical part.

So you want to be able to tell the text processor that you want to append a list of different elements.

So when you append the list we're using that list strategy.

That's where all the fun happens.

So let's have a method on text processor called append list where you take a bunch of items just as strings and you append them using the strategy that you are using.

So here I can say that S is going to be to dot list strategy so that's our strategy.

And then I can say Ask Don't start and I provide obviously teed up builder like so.

Actually let me let me just think maybe we can maybe we should like this but then of course yeah.

That that kind of complicates things let's not do that so so we start the process and then we go through every single item four on this core common item in range of items.

So for every single item we take the strategy and we add the list item using the builder that we have specifying the item.

There we go.

And then towards the end once again we take this strategy and we call and on it we specify to the builder and that's all there is to it.

So this is how you can append an entire list.

Let's also add a reset method because we want to be able to reset that internal strings builder so func t text process or reset.

That's just going to do teed up builder.

So we just reset the builder like so and that's also have a string representation where once again I'm just going to implement the stringer interface for the text processor.

So the idea here is that we return we take the builder and we get the builder to return that string that's been accumulating there.

Okay.

So now let's take a look at how all of this works so I can make a tax process or new tax processor and here I can specify the strategy.

So in this case for example I can specify the markdown list strategy and then I can append a list T.P. dot append lists so I can just make a bunch of items foo bar bars like so.

And then we can print the whole thing so we can just take the tax process or I don't even need to do and you can care because it implements the string your interface and we can run the whole thing and just take a look at what the output is well as soon as it's done compiling that as any as you can see we're getting the right kind of stuff.

So we're getting a markdown list now what you can do is you can obviously change this from a markdown this strategy to an HMO strategy but you can also reuse add text processors so you can take the text processing you can reset it.

This is why I created the reset method.

You can call said output format thereby switching the format using a constant.

So in this case I'll use the constant HDMI email that we defined at the beginning of the file and then

I can append that list once again so we can take this list and we can append it once again.

And then once again we can print out the contents of the text processor so let's run this and let's take a look at what we get here.

And as you can see what we're getting here is we're getting the expected results so we have an unordered list and then we have the list items for each of the H2 AML list elements.

So this is an illustration of how strategy works.

So essentially what you do is you have you have a member which can be defined to different kinds of strikes.

So in this case we have a list strategy which is just an interface.

And this can be implemented by either the markdown list strategy the 80 mls strategy so you can switch one from another.

You can have one defined at initialization like we do here.

So here this strategy is defined when you initialize the text process but you can also have methods for switching dynamically switching from one strategy to another.

### Strategy code: behavioral.strategy.strategy.go

```go
package strategy

import (
  "fmt"
  "strings"
)

type OutputFormat int

const (
  Markdown OutputFormat = iota
  Html
)

type ListStrategy interface {
  Start(builder *strings.Builder)
  End(builder *strings.Builder)
  AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct {}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {

}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {

}

func (m *MarkdownListStrategy) AddListItem(
  builder *strings.Builder, item string) {
  builder.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct {}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
  builder.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
  builder.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
  builder.WriteString("  <li>" + item + "</li>\n")
}

type TextProcessor struct {
  builder strings.Builder
  listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
  return &TextProcessor{strings.Builder{}, listStrategy}
}

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
  switch fmt {
  case Markdown:
    t.listStrategy = &MarkdownListStrategy{}
  case Html:
    t.listStrategy = &HtmlListStrategy{}
  }
}

func (t *TextProcessor) AppendList(items []string) {
  t.listStrategy.Start(&t.builder)
  for _, item := range items {
    t.listStrategy.AddListItem(&t.builder, item)
  }
  t.listStrategy.End(&t.builder)
}

func (t *TextProcessor) Reset() {
  t.builder.Reset()
}

func (t *TextProcessor) String() string {
  return t.builder.String()
}

func main() {
  tp := NewTextProcessor(&MarkdownListStrategy{})
  tp.AppendList([]string{ "foo", "bar", "baz" })
  fmt.Println(tp)

  tp.Reset()
  tp.SetOutputFormat(Html)
  tp.AppendList([]string{ "foo", "bar", "baz" })
  fmt.Println(tp)
}
```

## Summary

All right so let's try to summarize what we've learned about the strategy design pattern.

So the idea is very simple actually you define an algorithm at a high level and then what you can do or one of the approaches is that you can define an interface that you expect each of these strategies to follow and then you somehow support the injection of the strategy into the high level algorithm so that could be just a parameter being passed or it could be a field that gets assigned at the initialization of the algorithm.

And subsequently you use that field's methods to actually do something.
