# Builder -- when construction gets a little bit too complicated

## Overview

All right so we are finally discussing design patterns.

I thought it would never get to this point but we're done with the SOLID principles so let's talk about the first of the design patterns which is called builder.

So what is the builder for.
Well some objects that you create.

So some strokes are rather simple and you can just create them using either you know a constructor call a factory function call if you have one or if you don't have one you can just create them by initializing their fields and you don't have to initialize all the fields you can partially initialize some of the fields.

So in some situations this works just fine.

But in other situations objects actually require a lot of ceremony to create.

So for example if you have some sort of factory function with 10 arguments in it this isn't really productive you're forcing the user to make lots of decisions within a single expression within a single statement and that's never a good thing.

So we want to somehow work around this and make the construction process a kind of a multi-stage process.

So we construct an object piece wise as opposed to trying to do everything in a single factory call.

So instead what you can do is you can all four piece wise construction that's when you take an object and you kind of build it piece by piece and the builder pattern is basically all about providing some sort of an API for constructing an object step by step as opposed to trying to construct it all at once.

So the build a design pattern is what you do when you want piece wise object construction when it's complicated.

Then you make a builder and you provide some sort of API for ensuring that the construction is understandable and it happens succinctly.

## Builder

Let's take a look at how you can implement the builder design pattern.
And we're going to begin by using a builder that's actually already built into go.
And there is the string builder.

So let's imagine a situation where let's say you're writing a web service so web server is supposed to serve HTML out.
It also serves other things like javascript but let's just stick to a for the time being.

So the idea is that you need to build up strings of HTML from ordinary text elements so for example you have a piece of text and you want to turn that text into a paragraph.

So the simple way to do it would be something as follows.
Let's suppose you have a piece of text like let's say you have just a string Hello for example.

So if you want to build up a paragraph and they stem him a paragraph from this string what you would probably use is something like strings builder.

So as B is going to be a String start builder so a strings builder is a built in component it's a component which is part of the go SDK which actually helps you concatenate strings together so it allows you to write several strings one after another into a buffer and then get the final concatenated result.

So here what we can do is we can use as b the right string to actually add a couple of strings to the buffer so we can add the opening paragraph tag and then we can add the string itself in our case the string is called Hello.

And we can have the closing paragraph tag and here I can just do FMC don't print line and we can print line the contents so I can take as bead on string just called the string method on it and see what we get and hopefully we get our paragraph.

And that's exactly what we're getting here so we're getting the paragraph with the word hello.

So this is a very simple example you just have a single word you put that word into a paragraph what if you had a lot longer text longer than one word that would work as well.

Let's take a look at something more complicated.
Let's suppose you have a list of words and you want to put them into a list.
So I have a bunch of words.
Let's just make a couple of words.
Let's have.
Hello.
And the world like so you can think of you know bigger lists with more words.
So essentially we're going to reuse the strings builder that we have up above.

So I'll say SB the reset so we'll reset the contents and then what we want to build is we want to build an unordered list.

That's going to be basically the U L tag followed by a bunch of list item tags of where you provide the text.

You close the tag and then you repeat it for however many words you actually have.
And then at the end of it all you put in the closing unordered list tag.
So how do you do this using the strings builder.

Well it's also rather easy we continue using that whole right string business.

So here I can put the opening tag and somewhere towards the end I'll put the closing tag and then we can iterate the words.

So for underscore comma V in range of words what we do is we once again write string.

So we write the opening list hide and tag we write the closing tag here and in the middle we just put the word itself.

So that's also fairly simple stuff.
And then once again I can definitely go to print line to show you what we got here.
So as B that string will give us R on ordered list.
All right.

As you can see we got pretty much what we wanted but the whole process of building up a city email by using these tiny little pieces of text and using the string builder is a bit too complicated is a bit too much too much work basically because you know here that you want an opening tag with Ally and then you want a closing tag with Slash ally.

So why not take everything that we have here and just put it into structures put it into corresponding structures with methods which are more more flexible to use and just just easier to use basically.

So this is the reason why you get builders in the first place.
So the idea is that you have a some sort of object that you want to build up in steps.

So here is a step and here is a step in here is another step and there's lots of steps to building up an object and you want to make it convenient.

You want to make it easy for the Cline to for example build H2 AML elements.

And without them having to tie pull the strings with the opening and closing symbols and all the rest of it.

So you can represent the entire HDMI construct as basically a tree and then just learn how to print that tree and give give the user a nice builder component where they can add elements to this tree without necessarily being aware that the tree is actually there.

So that's exactly what we're going to implement and we'll start by defining the idea of an H timeout element.

So typeface to AML elements struct.
So that's gonna be a structure way each element has a name.
So a name would be something like P if it's a paragraph for example but HDMI elements also have text in Texas.

What goes in between the tags.
So here is an opening p tag you as a closing P Dag.
So this part would be the text.

So both of these are strings but also in addition to these two fields you have to realize that a simple elements can also contain other elements.

So let's include this as well.
So elements each demo element like so.

So we're going to have those and now what we need to do is we need to be able to what we need to do a couple of things first of all we need these H2 AML elements to be principal.

And here unlike in our console we're going to have all the nice trappings of HDMI like indentation for example.

So I'm going to define an in then size here.

Let's just define indentation as two spaces and then I will define a string method for the stringer interface so that we can print each HDMI out element.

This is actually reasonably convoluted steps I'll just cut and paste something I've written before so essentially the string the string your implementation with a capital S is going to call the method with a small s and that's going to do all of the work because this is a recursive method which needs to take the indentation level to print the correct indent all the correct size.

So that's what's happening in here on the going to go through this code.

Just trust me that it prints the HDMI elements correctly but we're here to talk about the builder design pattern much more interesting.

So here I'll make a tie called a team builder which is going to be a stroke.
Okay so what do we need inside the HDD email builder.

Well the HMO builder basically just cares about the root element so long as you have a root element which is an HMO element so long as you have one of those you can call root adult string and you can get the representation.

So in addition to that we'll keep another thing we'll keep the name will cache the root name separately because sometimes you want to reset the builder you want to take the builder and you want to wipeout everything in it.

However in each demo you always need a root element and that root element has to have a name so we'll keep root name here as a string.

Okay so I'll just make a utility function for making a new HDMI l builder.
So let's have that new age demo builder.

We just specify the roof name as a string and will return pointed to an HMO builder like so so here
I'll just make a I'll just basically return an HMO builder.

So we provide the root name as the first argument and then we have to create the actual root so a root is an HMO element where the name is Ruth name.

We're not going to specify any text and the final argument.
The set of elements is just going to be default implementations though.
So a default slice of each demo elements elements like so so.

Well that's that's pretty much how you would implement a kind of new method or new function for creating an HMO builder initialized with a particular root name.

But now of course what we want to be able to do is we want to have utility methods for actually populating HTML out elements like for example here when we create the list items you want to be able to take the route you well element and you want to be able to add a child.

So let's add this and also we need to have a string representation for the HMO Boulder itself which is just the representation of the root element.

So let's do that quickly so B is in case the builder.
So we're gonna have the string method which returns a string.
And here we return BDO roots dot string.

So we simply return the string implementation call for the roots and then let's have a utility method for adding a child.

So func B H2 builder we're going to have a method for adding a child so a child is going to have a child name.

That's the name of the tag and also the child text and both of these are going to be strings.
So here what we'll do is we'll create the H email element let's call it E.

So it's going to be an HTML element where you provide the name the child name rather child name child.

Text and then the lost argument is once again an empty set of elements HMO elements like so.
So that's that's pretty much all that you have to do.

And then once you've created this element you simply say well builder dot root not elements equals append beat up root elements comma E that's all that you have to do.

So armed with all of this what we can do is we can start using this builder to make this entire complicated process simpler because all you have to do now if you want to just just have a couple of elements is you make a builder so you say B is going to be a new HDMI builder with the root of unordered list and then for every single child you just say beat ought to have child and you specify the type of the the name of the tag as well as what the text contains.

So like hello for example or world in this case all we could go through the loop one by one because where they have it in a in a slice somewhere and then I can just f empty the print line the whole thing.

So once again Vito string should hopefully give us the kind of text that we're looking for.

And here it is it's actually rather pretty as you can see we have everything nicely formatted with the indentation and everything and the the end user the consumer of the builder they just need to care about the utility calls they don't care about anything else.

So one thing I want to show you which shows up quite a bit inside the builder pattern is the use of fluent interfaces so a fluent interface basically is an interface that allows you to chain calls together now changing calls and go isn't really that convenient because you kind of you leave the DOT hanging at the end instead of at the beginning.

But but we can't do this so let's just copy our child and we'll make a new method called add child fluent.

Now the only difference between this method and the previous method is that this method is going to return the argument the return the actual receiver so the receiver here is the HDMI builder.

So we'll just return an HCM builder.
And here we simply do return b.
Now why did we do this.
What is the point of doing this.
Well the point is that you can change these calls together.

So now if we come back down here you can see that instead of calling our child and then doing Beatport and child what I can do is I can put a dot here I can remove all of this completely and if I put fluent and obviously fluent here then everything continues to work as before and this allows you to kind of have a single statement effectively where you simply do the call but then you do the next call immediately after and then you can continue this chain of calls infinitely so fluent interfaces basically returning the receiver or appointed to the receiver at the end of the method is something that you're going to see in many design patterns you'll see it in many locations because it's just a convenient way of helping the user along.

So once they call something they know that they can reuse the result of that call to continue calling on that thing again and in the case of a builder that means you can continue calling the different build methods there by sort of building up your object until it is finally complete and you are ready to use it.

## Builder code: creational.builder.builder.go

```go
package main

import (
  "fmt"
  "strings"
)

const (
  indentSize = 2
)

type HtmlElement struct {
  name, text string
  elements []HtmlElement
}

func (e *HtmlElement) String() string {
  return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
  sb := strings.Builder{}
  i := strings.Repeat(" ", indentSize * indent)
  sb.WriteString(fmt.Sprintf("%s<%s>\n",
    i, e.name))
  if len(e.text) > 0 {
    sb.WriteString(strings.Repeat(" ",
      indentSize * (indent + 1)))
    sb.WriteString(e.text)
    sb.WriteString("\n")
  }

  for _, el := range e.elements {
    sb.WriteString(el.string(indent+1))
  }
  sb.WriteString(fmt.Sprintf("%s</%s>\n",
    i, e.name))
  return sb.String()
}

type HtmlBuilder struct {
  rootName string
  root HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
  b := HtmlBuilder{rootName,
    HtmlElement{rootName, "", []HtmlElement{}}}
  return &b
}

func (b *HtmlBuilder) String() string {
  return b.root.String()
}

func (b *HtmlBuilder) AddChild(
  childName, childText string) {
  e := HtmlElement{childName, childText, []HtmlElement{}}
  b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(
  childName, childText string) *HtmlBuilder {
  e := HtmlElement{childName, childText, []HtmlElement{}}
  b.root.elements = append(b.root.elements, e)
  return b
}

func main() {
  hello := "hello"
  sb := strings.Builder{}
  sb.WriteString("<p>")
  sb.WriteString(hello)
  sb.WriteString("</p>")
  fmt.Printf("%s\n", sb.String())

  words := []string{"hello", "world"}
  sb.Reset()
  // <ul><li>...</li><li>...</li><li>...</li></ul>'
  sb.WriteString("<ul>")
  for _, v := range words {
    sb.WriteString("<li>")
    sb.WriteString(v)
    sb.WriteString("</li>")
  }
  sb.WriteString("</ul>")
  fmt.Println(sb.String())

  b := NewHtmlBuilder("ul")
  b.AddChildFluent("li", "hello").
    AddChildFluent("li", "world")
  fmt.Println(b.String())
}
```

## Builder Facets

In most situations that you encounter in daily programming a single builder is sufficient for building up a particular objects.
But there are situations where you need more than one builder way.
You need to somehow separate the process of building up the different aspects of a particular type.

So let me show you how this can work.
Let's suppose that we have a time person which is a struct.
Now let's suppose that a person has two particularly particular types of information that you want to build up about them.

So one type of information is related to their address so you might have a street addresspostcode city and so on.
And those of these string.
So this is related to where a person lives.
And another would be related to where they work.
So their job information so you might have things like company name or position or you might have annual income for example that would be an end.

So imagine you want to have a separate builders for building up the address information and for building up the job information.
The question is how do you do it.
Well it's actually not that complicated.

Let's imagine that we have some sort of type called Person builder which serves as a starting point for building up a person.
So we're going to have a type called Person builder which is going to be a struct.
And here what we're going to do is we're going to have just a person point.
So this is going to be appointed to a person that is being built up and obviously we have to initialize it somewhere because it's nil by default.
So typically you would have a function somewhere maybe have a function right here for making a new person builder.

So when you make a new person builder which returns a person builder point of what you do is you make a person builder where you also initialize the person.
So you have to initialize the person to something just just a default default set of values for that person.

Now what we can do is we can have additional builders for the address information for the job information so instead of putting anything inside the person builder what we do is we make additional builders which aggregate the person builder.

Now as soon as you aggregate the person builder you automatically get a pointer to person.
So you aggregate person builder and then you have this field available to you.

So let's do that.

So we have two different aspects.
We have a job builder for the job information and we have this address builder.
So we'll have type person address builder which aggregates person builder.
And similarly we have a type well that's actually just duplicate this whole thing.
So we have an additional type called Person job builder which also aggregates person builder.
So what do we want to be able to do.

Well our starting point when building up a person is the person builder but from person builder we want to be able to provide interfaces which are provided by the address builder and the job builder respectively.
Aldi here and same goes for this location.
OK.

So how do we do this.
Well in actual fact the first thing that we can do is we can provide utility methods on person though there which give us a person address builder and person job builder respectively so the utility method for providing a address builder is going to be lives.

So we say func be person builder lives.
And this returns a person address builder.
There we go.

So we simply construct a person address builder and notice that because person address builder aggregates person builder we have to initialize this pointer.

So we initialize this pointer by simply saying you know let's put Star been here so we make a copy of the builder we put it into person address builder and that copy is the pointer that is pointing to the object we are currently building up.

So this thing is called lives and we're going to have a similar thing duplicated a similar thing called works for the person job builder.

So here we'll put person job builder.
And of course we copy and put it here.

So now we have ways of transitioning from a person builder to either a person and trust builder and a person job builder but you have to realize that effectively in effect person Joe builder and person address builder are both person builders and as a result when you have a person and dress builder you can quickly use the book's method to jump to a person job build and vice versa.

You can jump back to the person in dress builder using the lives method.
So that's very convenient and not what we can do is we can actually populate the methods of person job builder in person and dress builder.

So let me show you how this can look so for example for the person in dress builder you would have a builder methods similar to the following.

So this should come as no surprise so person a dress builder has a method called hat so you can say you live at a particular street address you can live in a particular city and you can have a particular postcode.

And in all three of these cases all we do is we assign the appropriate value.

So we say it that person the postcode and I should rename that it to something like B for builder.

So let's rename B like so.
Actually I should rename it everywhere.
Just just for consistency sake.
So let's have be like so.
OK.

So having set this up we can do the same thing for the job build.

So once again I'm not going to type out all these building methods I'll just paste them down below and you can take a look.

So effectively for the person job builder we have methods such as ADD so you work at a particular company has a developer manager whatever earning a particular amount.

So we've set up effectively a kind of tiny DSL a tiny domain specific language for building up a person's information.

And now of course all we have to do is we have to somehow provide a build method where we actually yield the whole thing.

So once the object is built up we return it so let's add it well it's handed down at the very bottom so here we'll have func be person builder build person.

We simply return b return it a person.
That's pretty much all that we have to do.

And now we can use all of this to effectively construct a person but construct it no using one builder but effectively using two builders.

So APB is going to be a new person builder like so and then we can say APB dot and here I'm going to use indentation on the next level so you can see how everything is grouped together.

So first of all we want to provide information about where the person lives.
So we say lives.
And once again I'm going to put a dot here and indent things once again so a person lives at 123 London Road.
Let's put a dot and they live in London once again for the dot here with postcode as W 12 B.C. for example.
Not sure if it's a real postcode or not.
OK.

So they live at this particular location and then we can provide the works work information.
So works.
So person works at Fabry cam and they work has a programmer network as a programmer and they.
Let's say one twenty three thousand.
OK.

So now that we have this we can use the though there to build a person so I can say person is equal to PVR out build like so and then I can just print line the person just get the default output for it so let's run this let's see first of all if it works or not.

And as you can see we're getting pretty much what we expected.
So we fully initialize the person with all the relevant data.

So once again let's recap of what's happening here because it's not immediately obvious what's going on.

So instead of using one builder we're effectively what effectively we have three builders to be honest because person builder is also a builder.

It doesn't do anything apart from giving us the lives and works methods but it's still there.

And what happens now let's just go back to the full initialization and I'll walk you through it I'll walk you through what happens.

So when you call lives effectively you're calling lives on person builder but lives gives you a person address builder.

And once you complete this dot here you're working with a personal address builder so you can provide address information like where they live in what city with what postcode.

But remember that a person address builder happens to be a person builder and a person builder also has a works method.

So at any point in time like here for example or here or at any point in time you can call the works method and we're doing it right here to switch from one builder to a completely different builder and now we're working with the person job builder.

So here we specify the company name and the position of the employee as well as how much money they earn.

So this demonstration shows that it's possible to set up a situation where you have several builders who are working cooperatively because they share their they both aggregate a type which actually has the pointer of the object being built up this is possible in order to be able to build up an object by using several builders as opposed to just a single builder.

## Builder Facets code: creational.builder.builderfacets.go

```go
package main

import "fmt"

type Person struct {
  StreetAddress, Postcode, City string
  CompanyName, Position string
  AnnualIncome int
}

type PersonBuilder struct {
  person *Person // needs to be inited
}

func NewPersonBuilder() *PersonBuilder {
  return &PersonBuilder{&Person{}}
}

func (it *PersonBuilder) Build() *Person {
  return it.person
}

func (it *PersonBuilder) Works() *PersonJobBuilder {
  return &PersonJobBuilder{*it}
}

func (it *PersonBuilder) Lives() *PersonAddressBuilder {
  return &PersonAddressBuilder{*it}
}

type PersonJobBuilder struct {
  PersonBuilder
}

func (pjb *PersonJobBuilder) At(
  companyName string) *PersonJobBuilder {
  pjb.person.CompanyName = companyName
  return pjb
}

func (pjb *PersonJobBuilder) AsA(
  position string) *PersonJobBuilder {
  pjb.person.Position = position
  return pjb
}

func (pjb *PersonJobBuilder) Earning(
  annualIncome int) *PersonJobBuilder {
  pjb.person.AnnualIncome = annualIncome
  return pjb
}

type PersonAddressBuilder struct {
  PersonBuilder
}

func (it *PersonAddressBuilder) At(
  streetAddress string) *PersonAddressBuilder {
  it.person.StreetAddress = streetAddress
  return it
}

func (it *PersonAddressBuilder) In(
  city string) *PersonAddressBuilder {
  it.person.City = city
  return it
}

func (it *PersonAddressBuilder) WithPostcode(
  postcode string) *PersonAddressBuilder {
  it.person.Postcode = postcode
  return it
}

func main() {
  pb := NewPersonBuilder()
  pb.
    Lives().
      At("123 London Road").
      In("London").
      WithPostcode("SW12BC").
    Works().
      At("Fabrikam").
      AsA("Programmer").
      Earning(123000)
  person := pb.Build()
  fmt.Println(*person)
}
```

## Builder Parameter

So one question you might be asking is how do I get the uses of my API to actually use my builders as opposed to stop messing with the objects directly.

And one approach to this is you simply hide the objects that you want your users not to touch.

So for example let me show you a very simple example let's suppose that you have a an API of some kind for sending emails.

So you have a function called sent email.

And certainly if you actually have a structure for representing an email you could put the structure here.

So here you would say something like email e-mail and then you would actually have a structure somewhere like type e-mail struct where you would have some fields like from two subject body and so on and you would send that unfortunately the the problem or at least one of the problems is that you want your emails to be fully specified.

Then you can suddenly write a validator so you can write a component which validates each email or what you can do instead is you can create a builder so that the user can invoke methods on that builder in order to build out that email so we can keep the email with the lowercase letter here so we don't let it kind of bleed out from outside the package.

But what we can do is we can create a more available type called email builder so an email builder is something which is going to build up an email but it's not going to it's not going to expose the different parts of the email directly.

So here we'll have just email e-mail.

So we'll effectively aggregate all the information about the e-mail and then we can have utility methods for actually building up that e-mail.

So for example you can have a function function which specifies the from field so you specify who the e-mail is actually from we'll make a fluent interface as the force out with the e-mail builder here and here you will typically see like beat out e-mail down from equals from and you can return b just to provide that fluent interface.

But in addition this is another location where you can perform some sort of validation.

So a simple validation on an e-mail would be something like Well if for example the the string from doesn't contain the character then we can panic we can say not panic email e-mail should contain an ATH character.

So this would be an example of something that you you'd want to check against.

Let's just put the curly braces in here.

And similarly you would sort of populate the other builder methods for a given e-mail so you would have the e-mail builder and have different things like who to send the e-mail to and what the subject is and what the body is and so on and so forth.

So now we come to the important part.
How do you actually get people to send your email.

Because obviously somewhere behind the scenes somewhere you want to have a function called send mail impulse for example which actually takes the e-mail.

And it does whatever you need to do with it but you don't want your clients to actually work with the e-mail object.

You only want to work with a builder.
So how do you do this.

Well it you can do this by using a builder parameter.

And that's basically going to be a function which sort of applies to the builder.

So you have to provide a function which takes a builder and then does something with it typically sort of calls something on the builder.

So we define a type.
Let's just call it build.

It's going to be a function which takes an e-mail builder pointer and we're going to use this function in our publicly exposed function called send email.

So send e-mail is the function that people are meant to be using.

And here the argument the only argument is called action and it's a type build.

So what happens is whenever somebody calls send e-mail they have to provide the body of the function which takes an e-mail builder has the first and only parameter and doesn't return any values.

So in here what we would do is we would initialize the builder.

So it's a builder is an e-mail builder like so then we would perform the action on the builder or on the builder point rather so pointed to the builder.

And then we would do the internal kind of send mail impulse stuff just providing the e-mail part.

So builder that e-mail.

OK so how would all of this work what would this actually look like.

Well here's the idea from the client's perspective.

They have to call send e-mail.

But they see that there is a function that needs to be provided here so their I.D. if they use the co-generation of their idea they get something like the following.

So they will get the following generated code where you have to and I want that as a b where you have to provide a function which takes an email builder and this is precisely the location where you don't have access to the email object itself.

You only have access to the builder and you can use that builder to build up information about the email.

So here I can say beta from foo at bar dot com.

I can also say that the email is to bar at bars dot com.

Once again the dot here I can say that the subject is meeting and I can say for example that the body of the email is

Hello.
Do you want to meet.

So that's how I would define my email.

So what really happens is that when we go here when we call this whole thing we create the builder which is an email builder.

And then what we do is we apply the action which is the entire body of this function.

So we've defined a function we pass the function in to send email.

So we invoke this entire function on the build a pointer.

So basically the Build a pointer.

This part gets passed in here and gets used in here to actually initialize the whole thing.

And then by the time we do the input of the builder has been initialize our builder.

That email has been initialized with the right stuff.

Hello we call send mail impulse taking the email object that object that the clients are not supposed to be seeing.

We have everything initialized correctly with all the validations and all the rest of it.

So this is yet another approach to how you can use the builder how in fact you can force declined to use the builder as opposed to providing some sort of incomplete object for initialization for example.

So another viable approach and it's also used in in many places in real life code.

## Builder Parameter code: creational.builder.builderparameter.go

```go
package main

import "strings"

type email struct {
    from, to, subject, body string
}

type EmailBuilder struct {
    email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
    if !strings.Contains(from, "@") {
        panic("email should contain @")
    }
    b.email.from = from
    return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
    b.email.to = to
    return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
    b.email.subject = subject
    return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
    b.email.body = body
    return b
}

func sendMailImpl(email *email) {
    // actually ends the email
}

type build func(*EmailBuilder)
func SendEmail(action build) {
    builder := EmailBuilder{}
    action(&builder)
    sendMailImpl(&builder.email)
}

func main() {
    SendEmail(func(b *EmailBuilder) {
        b.From("foo@bar.com").
            To("bar@baz.com").
            Subject("Meeting").
            Body("Hello, do you want to meet?")
    })
}
```

## Functional Builder

So one way of extending a boulder that you already have is by using a functional programming approach.

So let me show you how this works.

Let's suppose that we have a simple type called the person and you want to have a builder for that person or at least an aspect of that person and then you want to extend the whole thing.

So let's suppose a person has a name and a position in a company.

So let's first of all make a person Boulder which initialize is just the person's name and not their position.

So you can advertise called Person builder which is just going to be skyrocket.

And what I'm going to do here is I'm going to have a list of actions a list of modifications that are going to be applied to this person.

So essentially a person modification is just going to be a type let's call it plus an MoD which is going to be a function which takes a person pointer and does something to it.

So change is something on that person and then inside the person the will have a list of actions.

So we'll have a list of actions which is just a slice of person mod like so.

OK.

So what we want to be able to do in the builder is we want to be able to specify that a person is called something so func the person builder so a person has called some name and we return a person builder just so that we have a fluent interface.

Now we go.
OK.

So here instead of just performing the modifications straight away what we do instead is we add this modification to the list of actions.

So we say we don't actions equals append vetoed actions and then we create a function which takes a person and performs the modification on that person.

So we say pedo name equals name.
So we create this function.

There's going to be called in the future when I'm calling it right now and then we're just returning the builder in a fluent way.

So let's also have a method for actually building up the person once all those actions have been compiled so func B.

Person builder is called build and it returns a person.

Now we go so we create a new person like so and then for every single action in the range of beat on actions what we do is we invoke the action on a pointed to that person and then we return a point to that person.

That's how you go through the actions and apply every single one of them.

And we can the look at the way it works.

So here I can make a builder see you as a person builder and then I can say that a person is going to be b dog called Dimitri dot works well we haven't got works yet but just build and that gives us a person that we can print line subsequently.

So let's just print line star P.
Let's take a look at what we get here OK.
So we have James.

And there is a gap here because of the position that this argument is currently an empty string.

So this a setup.

The benefit of this setup is that it's very easy to extend the builder with additional build actions without messing about with making you builders which aggregate the current builder and so on and so forth.

What we can do is we if we want for example to extend person builder to specify the place of work is we just make a new function.

So we make a function called Make function called works as a.

And here you specify the position as a string and once again we return a person value pointer.

There we go.
So here be the actions.

That's actually correct the indentation so beat up actions equals append beat don't actions and then once again we have a function which takes a person.

And here we say speed up position equals position and then we return b.

And what this allows us to do is it allows us to effectively change the course here so after saying

B dot called name is James.

I can see that works as a developer for example dot build like so.

So this lets actually run this let's take a look at what we get here.

And as you can see we're getting my name and then we're getting the job title as well.

So this setup all it does is it illustrates that effectively what you can do is you can have a kind of delayed application of all of those modifications so your builder instead of just doing the modifications in place it can keep a list of actions a list of changes to perform upon the object that's being constructed.

And then when you call build what you do is you create just a default implementation of the object and then you go through every single action and you apply that action to that object that you are returning and then subsequently just return that object.

## Functional Builder code: creational.builder.functionalbuilder.go

```go
package main

import "fmt"

type Person struct {
  name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
  actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
  b.actions = append(b.actions, func(p *Person) {
    p.name = name
  })
  return b
}

func (b *PersonBuilder) Build() *Person {
  p := Person{}
  for _, a := range b.actions {
    a(&p)
  }
  return &p
}

// extend PersonBuilder
func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
  b.actions = append(b.actions, func(p *Person) {
    p.position = position
  })
  return b
}

func main() {
  b := PersonBuilder{}
  p := b.Called("Dmitri").WorksAsA("dev").Build()
  fmt.Println(*p)
}
```

## Summary

All right let's try to summarize what we've learned about the build a design pattern.

So a builder is a separate component that's used for building an object and to make a builder fluent which is one of the things that is sometimes useful.

You can return the receiver so you have a method and that method can actually return the receiver so return the pointer which allows chaining so you can have several builder method calls one after another without any kind of ceremony.

Of course goes for matter kind of breaks things a little bit because you have to be careful with the line breaks you want to do line breaks but you have to be careful with them.

So different facets of an object can also be constructed with several different builders.

So if you have a really really complicated object and if you have different aspects of that object being constructed as separate concerns of separate aspects that you might want to sort of separate then you can have several builders which can who can all work in tandem via some sort of common common struct that they're aggregating and they can work in tandem to build up the different aspects of this object.
