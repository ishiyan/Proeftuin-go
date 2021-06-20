# Visitor -- Allows adding extra behaviors to entire hierarchies of types

## Overview

With all that said we've reached the final pattern of this entire course and that is the visitor design pattern.

So what is the motivation for using the visitor.

Well one of the problems that we sometimes try to solve is we try to define a new operation not on this single type but on several types at the same time.

So for example if you have a document modeled so you have a document represented as lists and paragraphs and other things let's say we want to add printing functionality.

And the question is how do you do this efficiently because you don't want to keep modifying every type.

In this hierarchy you don't want to go into let's say 10 or 15 different strokes and give each one of these trucks a method called to print that is just too difficult.

And also it doesn't conform to the Single Responsibility Principle because if we think about printing as a separate concern then applying separation of concerns.

We want to have the new functionality separate.

We want some sort of printer which knows how to print a list and print a paragraph and so on.

And we don't want to spread this functionality across our hierarchy.

So this approach is quite often used for traversal like let's say you have a binary tree of elements and you want to traverse those elements in a particular way.

Well in this case your elements can actually help you.

And this this can be an alternative to the iterator because the hierarchy members like the tree elements for example they can help you traverse themselves.

So you want to take a hurricane and you want to take every single element and teach it how a visitor.

So how some components can actually visit every single one of these linked entities but you don't want to keep modifying those entities.

And this presents a particular problem that the visitor pattern is actually trying to solve.

So the visitor is a pattern where some component or visitor is allowed to visit or traverse the entire hierarchy of types and there is a clever trick to the way this is done and then it's implemented by changing.

So you do change the types but you propagate just a single method called except throughout the entire hierarchy and then you use that single method to add additional functionality whether it's functionality for traversal or just something else some sort of other processing of a set of related types.

## Intrusive Visitor

All right so we're going to take a look at different implementations of the visitor design pattern.

And the first implementation that we're going to take a look at is what's typically called an intrusive representation.

So let's take a look at how this can actually work.

So what I'm going to do is I'm going to set up a scenario.

Now in this scenario we're going to be taking a look at expressions such as one plus two plus three for example.

So we'll be having numeric expressions but they will be represented in terms of struct what I mean is that we'll have some sort of an expression interface and we can then think of expressions as being composed of other expressions so if you look at this expression right here what is it composed of.

Well there are only two constituent parts.

Either you have a number so you have a value of some kind basically a double precision expression or you have a binary expression which has a left side and a right side.

And we can represent this entire thing in terms of those two constructs.

So now that we have the expression interface we can have a type called Double expression.

I'm calling a double because we're going to be using double precision floating point numbers.

So the value here is going to be a float 64.

We are storing the value of this expression.

And in addition what we're going to have is in addition expression.

So type addition expression also a truck.

This is where you basically have pointed to the left side and the right side.

So for example you can think this whole thing as an addition expression with the left side itself being an addition expression with values 1 and 2.

And on the right hand side you'd have a double expression with a value of three so an interesting expression has a left side and the right side and they are both off type expression where of course expression is an interface that we've defined up here.

So now that we have this setup what we can do is we can construct such an expression.

So let me let me actually do an expression one plus and then two plus three so slightly different arrangement to what we have up here so this is going to be an expression composed of an addition expression where the left side is a double expression with a value of one and that has to be a pointer and we could as well do it do a pointer here as well.

So the left side is this and the right side is an addition expression which is made up of parts double expression with the value of two and a double expression with the value of three love expression with the value of three so hopefully this all makes sense because this is how we set up the whole thing and now that we have the left and right parts what we want to be able to do is we want to be able to print an expression for example and before we get onto the visitor design pattern I want you to pause and think as to how you would actually go ahead and prints this whole thing because remember this is effectively an abstract syntax tree and you have to traverse this tree in order to be able to do anything but where is the printing done exactly.

So the first implementation of visitor is an intrusive visitor.

What we mean by that is whenever you call something intrusive it basically means that it intrudes into the structure that you've already created and any intrusive approach is by definition a violation of the open closed principle because remember the open closed principle basically states that once we have defined a double expression and an addition expression we shouldn't be able to jump back into those trucks and give them methods for example.

So that is one of the things that we're trying to avoid but since we're doing an intrusive approach we're going to modify this entire hierarchy so we're going to modify the interface so that we tell the interface that it knows how to print itself and then we're going to have implementations of double expression and addition expression.

So we go into the expression interface once again violating the open closed principle and we add a method called print.

So for every given expression you should be able to print the expression string representation into a builder into a strings builder like so.

So then of course we need to implement this interface in both a double expression and addition expression.

So let's do that first of all for the double expression.

It's not particularly difficult because all we have to do is we have to take that value and we want to print it in a nice way.

What I'm going to do is I'll just take the builder and I'll do right string on it right string and in here I'll use s print half so as print f I will use the G formatting.

You can also use the F form I think but then you'll have lots of zeros after the decimal separator.

I want the G formatting so if I have full 5.0 it doesn't write 5.0 it just writes five.

So here I use data value.
So this is how you print just a single float 64.

So this is easy now let's implement this for the edition expression.

So once again I go to implement methods I implement the expression interface here and so I have to have this print method.

Now here we're going to surround this expression in round brackets so I'm going to say SB dot right rune and I will write the rune for the opening bracket and that's duplicate this couple of times so we'll have it for the closing bracket and we'll have it for the plus in the middle but here and here what we need to be able to do is to print the left and right sides.

And of course that's not really a problem because the left and right sides are of type expression and expression has a print method defined on it.

So we can use that print method.

So I can take the left side and I can say left dot print providing these strings builder right here and I can do the same here so I can take the right side and print using the strength builder.

OK.

So this is something that's actually going to work and we can take a look at how it works.

So here I can make a strings builder like so and then I can say I've got prints so I take the expression and then I print into the string builder provided here as a pointer and then we can take we can just take the string builder get its string member and print that as the end result.

So let's take a look at what we get.

You can see that here we get round brackets around the whole thing because remember every single edition has round brackets around it.

And then we have one plus round brackets two plus three which is exactly what we wanted in the first place.

This was the expression that I defined right here.

So on the one hand we managed to implement a kind of visitor.

Now one question you might have is who exactly which component exactly is the visitor.

Which component is actually visiting every single element of this abstract syntax tree and then allocates the visitor here is the string builder so the string builder is the one that gets passed into every single print method.

And so it gets to visit additions expressions as well as double expressions.

That's why it's called a visitor on the other hand.

I said that this visitor is intrusive so it's not the best visitor that you can build because implementing this visitor implies that you modify the behaviour of both any interfaces that you have as part of the element hierarchy as well as the elements themselves.

So every single element suddenly has to have this additional method.

And imagine for a second that you want to have another visitor that actually calculates the value.

So one plus two plus three is equal to six.
You want to calculate this value.
How would you do this.

Well unfortunately in this setup what you'd have to do is you'd have to go into the expression interface.

You would have to add another method.

Basically another method called evaluate maybe which returns float 64 and then you would have to implement this in both double expression and additional expression.

Now another very important concept that we need to cover here is the idea of separation of concerns and single responsibility.

You might say well it's kind of the responsibility of each expression to print itself but not necessarily.

It would make more sense if you had a separate component.

Let's say a separate struct which knew how to print double expressions Ed. expressions as well as any other kind of expression you would have as part of your program.

So this is what we're going to take a look at in the next lessons we'll try to take out the print methods from these expressions and just have them exist in a separate component.

### Intrusive Visitor code: behavioral.visitor.intrusive.go

```go
package visitor

import (
  "fmt"
  "strings"
)

type Expression interface {
  Print(sb *strings.Builder)
}

type DoubleExpression struct {
  value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
  sb.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
  left, right Expression
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
  sb.WriteString("(")
  a.left.Print(sb)
  sb.WriteString("+")
  a.right.Print(sb)
  sb.WriteString(")")
}

func main() {
  // 1+(2+3)
  e := AdditionExpression{
    &DoubleExpression{1},
    &AdditionExpression{
      left:  &DoubleExpression{2},
      right: &DoubleExpression{3},
    },
  }
  sb := strings.Builder{}
  e.Print(&sb)
  fmt.Println(sb.String())
}
```

## Reflective Visitor

In the previous lesson we looked at an intrusive visitor a visitor that causes you to modify existing structure.

So imagine if you wanted to do it differently.

Imagine if you wanted to concentrate all the print methods in a separate component in a separate tract or indeed just to a separate function imagine you didn't want to modify these types and give them additional methods can you do this well it turns out you can.

So let's actually do that.

So first of all I'm going to go into the expression interface Now delete the print method and I'll delete all the implementations of the print method from these struct because we're going to have it separate now for this particular demo.

I'm not going to have any separate struct for printing different types of expressions.

Instead we'll just have a function it doesn't really matter for this particular demo we'll just keep things simple.

So the idea is that you have a function called print so you take some sort of expression and you print this expression into a string builder strings dot builder.

OK.

So the kind of visitor that we're going to build here is going to be what's called a reflective visitor.

Why is it reflected well.

Because typically there is this construct called Reflection.

That's when you look into a type and you look at what the type actually is what kind of members it has and go has a certain amount of support for reflection now.

One of the trademarks of reflection is that you check the actual type.

So here we have an expression but we don't really know what kind of an expression it is.

Luckily for us there are casts or we can try to do a cast in order to figure out whether this expression is a double expression.

In addition expression or something else.

So let's actually go ahead and implement the print method.

So the idea here is that at this point we don't know if this expression is a double expression or in addition expression but what we can do is we can do a couple of these casts so here I can say if the e come up okay is equal to and then I take e I take the expression and I tried to cast it to a double expression pointer.

So there are two outcomes here.

Either I succeed in which case ok will be true or I fail in which case OK we'll be false.

So if I succeed.

So if I get an ok then we perform the operation assuming that the variable D E is a pointer to a double expression.

So I can now start using it.

So I can do the right string for the string builder.

And here I can s print f once again using % G and I can take the value.

So you'll notice that the handling of the printing of the different types is now handled in a separate function as opposed to inside each of these expressions right here.

Okay.

So the other alternative is that the expression is actually in addition expression.

So we do a.
Come on.
OK.

Equals to e dot and then we tried to cast it to an edition expression pointer and we check whether or not this operation is actually OK and if it's OK we do the same thing as before.

So remember we had this whole write run thing going on well we had this three times so the opening the closing and the plus and in-between what we do is we continue to recursively call the print function.

So we call the print function we pass in a dot left.

So we pass in the left side as well as the string builder.

And then of course let me copy this and paste this we pass in the right hand side.

OK so now we have a self-contained function specifically for printing of the expression given a particular string builder so we can start using that.

Let's actually do this.
Now this code is going to change slightly.

Well we'll still have the strings builder but then we just call the print function passing in the expression as well as appointed pointed to the string builder.

And then finally we can take string builder and write it.

So the console once again.

So first of all before we discuss this let's actually run this.

And as you can see we're getting the correct result as before.

OK.

So this approach is a somewhat better approach than the approach we've taken in the previous demos.

The reason why it's better is because we've taken out this particular concern the printing concern.

So we're following this idea of separation of concerns.

We have a concern called printing the idea of outputting some types to text and we've taking it out and put it into a separate function that we could have created a separate struct if you would prefer

to have some sort of printer struct for keeping this function.

But well I'm just keeping this function at the top level right here.

So this function is a basically an implementation of reflect a visitor because we are checking the type.

So we take the variable e here and we check its type.

We try to cast it to this pointer if it's okay.

We do one thing and otherwise we try to cast it to some other pointer and we try to do a different thing.

So this approach is much better but it's not without its downsize the most obvious downside is what will happen if there is a third type.

So we have double expression edition expression.

Let's suppose you also have subtraction expression.

Now the reason why it's problematic is because all of a sudden you have to add additional code here.

So you have to go into this function which has already been written.

It's already been tested.

People might already be using it and you would have to add another else if inside this function which once again unfortunately we are breaking the open closed principle we want to be able to extend things we don't want to modify existing code that's already been written and tested.

So another possible problem is what happens if you forget.

Because imagine if you forget to write that if else imagine if for example at this point a dot right is actually a subtraction expression.

So you're going to be calling print on a subtraction expression.

So the subtracting suppression gets passed in here.
It fails this check here.
OK.

Is false it fails this check and then absolutely nothing happens so it's possible to actually have a an expression which doesn't get process because you forgot to write the code the process is it.

So the classic implementation of the visor that we're going to take a look at next.

It makes it pretty much impossible it makes it impossible to set up a scenario where you suddenly forget to implement a particular particular support for a particular type of an expression.

### Reflective Visitor code: behavioral.visitor.reflective.go

```go
package visitor

import (
  "fmt"
  "strings"
)

type Expression interface {
  // nothing here!
}

type DoubleExpression struct {
  value float64
}

type AdditionExpression struct {
  left, right Expression
}

func Print(e Expression, sb *strings.Builder) {
  if de, ok := e.(*DoubleExpression); ok {
    sb.WriteString(fmt.Sprintf("%g", de.value))
  } else if ae, ok := e.(*AdditionExpression); ok {
    sb.WriteString("(")
    Print(ae.left, sb)
    sb.WriteString("+")
    Print(ae.right, sb)
    sb.WriteString(")")
  }

  // breaks OCP
  // will work incorrectly on missing case
}

func main() {
  // 1+(2+3)
  e := &AdditionExpression{
    &DoubleExpression{1},
    &AdditionExpression{
      left:  &DoubleExpression{2},
      right: &DoubleExpression{3},
    },
  }
  sb := strings.Builder{}
  Print(e, &sb)
  fmt.Println(sb.String())
}
```

## Dispatch

So one thing that I wanted to mention in the discussion of the visitor pattern is the idea of dispatch what is dispatch and why do we care.

Well dispatch answers the question of which function should we be calling at a particular point in time.

So it's a decision which is taken at compile time not what we typically work on in most programming languages and certainly in go is the idea of single dispatch.

So the function that you want to invoke the exact function that you need to call at a particular point in time depends on the name of the request and the type of the receiver.

So if the type of the receiver is foo and the name of the method I want to call is bar then what we're going to be calling is the method bar on a struct foo.

That makes sense but in some cases that's not enough in some cases we need something called double dispatch.

Now double dispatch is the situation where the selection of which method to call depends on the name of the request.

So you obviously need the name of the actual method but you also need the type of two receivers not one receiver.

So the first type of receiver is the element that you're calling the thing on but the second receiver is the type of the argument.

And remember God doesn't have method overloads.

So this adds additional problem and you're going to see how those problems are solved in the real world.

## Classic Visitor

All right.

So we looked at the reflective visitor the idea being that you take some general expression and then you perform a bunch of typecast and then an attempt to figure out what kind of expression you actually got.

So if we are going to implement the classic double dispatch implementation of the visitor we need to discuss what dispatches because this patch is a very strange word you might not have heard of it but basically the idea of this patch is this patch is all about figuring out which function or which method you're actually going to call.

And in some cases it's easy.
But in other cases it's pretty much impossible.
So let me show you a very simple example.

Let's suppose that instead of having this print function you decide to have a separate print function for each type of an expression.

So for example for the double expression you decide to leave this chunk of code right here and for the additional expression you decide to have a similar kind of signature.

So you just copy this over and you have let's see in addition expression here like so and you know you do you try to do something like this.

Now unfortunately this isn't going to work.

You probably know why I'm going to work whether actually multiple reasons why this wouldn't work.

So one of the reasons of course is that you cannot overload functions.

Here we have a function called print here we have a function called print so if I try to compile this you you'll see a fairly obvious message so print gets re declared.

So you're not allowed to do this.

But imagine for a second if you were allowed to do this if you were allowed to.

So so to speak.

Overload the print function and would the code work then and the answer is no the code would still not work.

The reason why it wouldn't work is because of these problems right here.

Now whenever the compiler encounters either left it knows that e the left is an expression it doesn't know that or left is an additional expression or a double expression and it cannot figure it out at compilation time.

And the compiler wants to know the static type of either left in order to be able to do something with it to call a particular function but because it doesn't and that's pretty much impossible.

So if you were to call some function which has which has a way of accepting a double expression this wouldn't work because it's not a double expression is just an ordinary expression.

So this is a limitation of this patch.

So this patch is all about choosing which function to call and at the moment because either left and either right are expressions we cannot make any choices.

We cannot say that we have to call this particular function or we cannot say that we call this particular function it's simply impossible.

And that's the reason why this idea of doubled this patch is you.

So the idea of double dispatch is that instead of calling something like this which is impossible you take a dot left and then you call something on e dot left you call some sort of an accept method.

Now inside this accept method you do know who the owner is because that's the receiver parameter you do know who the owner is.

And so you can jump from an accept method back into some sort of a visit method like what print would be a visit method for example right here.

So by performing this double jump you implement something called double dispatch.

So double dispatch is being able to choose the right method not just on the basis of some arguments but also on the basis of who the caller is.

So we're going to implement this whole thing and we'll try to do it slowly enough so you can understand what's going on.

Now the first thing I'm going to do is I'm going to modify the expression interface.

Now some of you might say well hold on.

You just told us that the open closed principle implies you don't do this.

Well the thing about the double dispatch classic visitor is that you can modify this interface but you only do it once.

You don't have to do it several times you simply do it once and then that operation is leveraged for the many different kinds of visitors.

So that operation is called accept.

So we're going to have a method called accept and it takes a visitor.

So in this case I'm going to call it an expression visitor.

So an expression visitor is nothing more than an interface.

So expression visitor is going to be an interface and this interface is going to have different methods for visiting particular types of expressions because we have different types of expression we have a double expression we have an addition expression.

So we're going to have a method for each will have visit double expression and we'll similarly have a visit addition expression where we obviously take it and addition expression

has the argument like so.
OK.

So now that we have this interface what we can do is if we want to make let's say a printer for example we need to implement this interface.

But before we do that we have to take the expression interface and implemented implements specifically the accept method in both double expression as well as addition expression.

So let's just do this quickly Sol implements the expression interface on the table expression.

So the idea here is simple that we perform.

So we've just jumped into an accept method on an expression.

So here the parameter is double expression.

So that means if you take an ordinary expression you can call except on it because except this part of this interface.

So we're here and we want to return the control back to the visitor back to the expression visitor.

So we simply say evey dot visit.
Double double expression the like so.

So we return the control back but we return the control specifying what kind of expression we actually have.

So we're invoking the right method and we do the same thing for the ed. expression.

So here I'll go to implement methods and I have an expression here like so and we just say evey dots visit Ed. expression a.

So this is how we perform the double jump but now of course what we need to do is we need to do a demo we need to have some sort of printer.

So once again I'll make an expression printer.

So an expression printer is going to be a struct where there is a strings builder inside that we can use.

So when it comes to visiting the different kinds of expressions what we can do is just fix a small error here what we can do is we can implement that visitor interface so we'll implement the expression visitor interface and then gets us the visit double expression and a visit Ed. expression.

So here I'll just say X while actually expression printer.

So E.P. is a better name and of course I'll take it as a pointer.

So here the idea is that the printer is visiting a double expression and gets a double expression and it gets to print it.

So here we can just take the string builder and we can.

Well it's easy isn't it.

We take the string builder and we call the right string right string.

And here we do s print f has before as print f where the format is % G.

And the value is the expression that value.

So here we just print the ordinary value in this particular case we do the same thing as before.

So we say EPD out SB dot right runes so we write the opening the closing and the plus but in between something interesting happens.

So in between we say that we take the expression we take the left hand side of that expression and we call accept on it.

So this is where the double dispatch magic happens.

Let's see if you can follow me as I explain what's going on.

So we call E dot left dot except the reason why we can call this regardless of what e dot leftism is because either left is an expression and an expression is an interface that defines a method called accept so we know that the method accept is there.

Now when we pass an accept method something we pass in E P so we pass the visitor into accept.

So we go into except for the left side.

So we either end up here for the addition expression all we end up here for the double expression.

So one of these two things and depending on that we either call the visit double expression on the visitor or we call the visitor edition expression on the visitor when you call one of these methods you end up back where you started you end up either here or here but you end up in the correct overload with the correct information about whether you have a double expression or indeed you have an edition expression.

So we do this for the left side and we also do this for the right side.

So this is the trick that allows us to basically by doing this double jump we have all the information about the colour and the.

So we have this double dispatch approach.
So now that we have this.

Let let me just make a constructor for the new expression printers so I'll do that.

So a constructor for the men I'm going to pass a string builder just make a new one like so case.

So strings dot builder there we go.

So this is a constructive for the expression printer.

Let's also implement the stringer interface on an expression printer so expression printers implement the stringer interface like so and here.

On going to do is take the string builder and return the contents of string builder like so.

So where they have something that's going to work and let's actually take a look at how this example will change on the basis of this double dispatch visitor.

So I'll make a new expression printer new expression printer like so and then I will say E.P. dot visit addition expression so this is one thing I can do I can see E.P. dot visit addition expression.

But imagine if I didn't know whether I had an edition expression or something else.

So imagine if I didn't know the type of this particular variable then what I could do is I could take this variable and I could call except on it passing in the expression printer.

That's that's all that I would have to do.

And then of course the expression printer would still have the right information so we could actually take a piece of string and we can print the contents.

So if I run this we get one plus two plus three so everything is okay everything is as you would expect it to be.

Now let's talk about the sensibility of this approach because we've been fighting for having the support of the the open closed principle so the idea is the open closed principle is things are open for extension both closed from modifications but still some modification might be required.

Like imagine if you also have a subtraction expression then unfortunately you would have to go into the expression visitor interface and you would have to implement a visit subtraction expression.

But as soon as you did this and the interesting thing would happen.

So all of the all of the visitors which you wrote would you would be mandatory for them to support subtraction expression.

And this is a drastic difference to the previous example where you could actually forget to handle a subtraction expression and everything would still work.

In this particular scenario you cannot forget to handle a particular type of an expression.

So that is the modification that you would have to make in both the interface as well as the implementers of this interface like for example the expression printer.

But on the other hand the situation when you need a new visitor is much better.

The implementation of having a new visitor is easier and it does follow the open closed principle.

So imagine if you want to not only print the expression but also give it a value evaluate its final value.

It's very easy for us to write an additional visitor let's call it an expression evaluator.

So we'll have type expression evaluator which would store some result.

Let's say a float 64.

So in this particular case what you would do is you would once again implement the expression visitor interface.

So here we go.
The only thing I would change is the naming here.

I want this to be called E expression evaluator like so.

OK.
So how would you do this.

Well in the case of visiting a double expression you would simply store the results so e dot the e the result is equal to the result all the value of the double expression and in the case of the calculation well there would be a bit more work here shall we say.

Because you basically need to store the result of visiting the left side but you also visit need to visit the right side and both of these actually modify the result variable.

So how do you do this.

Well first of all you check out the left side so you don't left don't accept.

So check out the left side passing in the expression evaluator and then you store the results somewhere because it's going to be overwritten on the next call.

So I say X is equal to e the result then we called the right side.

So we say e dot right Dot accept.

So we try it with on the right hand side.

We once again have a result and we add that result to x.

So we say x plus equals E the result and then X has the sum of the left side and the right side.

So we can say e the result results equals x.

So this is what we want.

And now now we have the completely evaluated expression so let's take a look at how this can be used.

So here instead of this printout I'll do something different so we'll make an expression evaluator.

I don't need a constructor here because there is no builder to initialize or anything so expression evaluator like so I take the expression to accept and the expression evaluator and then I can print both of these.

So here I can for example let's do a print f so I can say percent S equals percent G.

So percent S is going to be the expression printer and percent G is going to be the result of the expression evaluator like so.

So that's actually run this.
Let's take a look at what we get.
Okay.

So we are now getting not only the textual representation of our extra abstract syntax tree but was are getting the evaluated result as well.

So this approach the person that you've seen here is the classic double dispatch visitor it is the most common kind of visitor you're likely to see and hopefully you can appreciate the benefits it gives and the flexibility it provides as you create additional visitors for example.

### Classic Visitor code: behavioral.visitor.classic.go

```go
package visitor

import (
  "fmt"
  "strings"
)

type ExpressionVisitor interface {
  VisitDoubleExpression(de *DoubleExpression)
  VisitAdditionExpression(ae *AdditionExpression)
}

type Expression interface {
  Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
  value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
  ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
  left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
  ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
  sb strings.Builder
}

func (e *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpression) {
  e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

func (e *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpression) {
  e.sb.WriteString("(")
  ae.left.Accept(e)
  e.sb.WriteString("+")
  ae.right.Accept(e)
  e.sb.WriteString(")")
}

func NewExpressionPrinter() *ExpressionPrinter {
  return &ExpressionPrinter{strings.Builder{}}
}

func (e *ExpressionPrinter) String() string {
  return e.sb.String()
}

func main() {
  // 1+(2+3)
  e := &AdditionExpression{
    &DoubleExpression{1},
    &AdditionExpression{
      left:  &DoubleExpression{2},
      right: &DoubleExpression{3},
    },
  }
  ep := NewExpressionPrinter()
  ep.VisitAdditionExpression(e)
  fmt.Println(ep.String())
}
```

## Summary

Let us summarize what we've learned about the visitor design pattern.

So the idea is that you take every single element in your hierarchy and you propagate a method you propagate a method called accept which takes some sort of visitor and then you just stick it throughout every single member of the affected target.

And then you make a visitor with a bunch of visit methods so you would have a visit to visit bar and so on for if every single element in the.

And then you connect the two together so every single accept method actually calls visitor dot visit something.

And this allows you to basically make new visitors and have those visitors add functionality to every single element of the hierarchy.

And this is useful for both traversal as well as any kind of other concern where you need to go through a set of related elements and get some information about them.
