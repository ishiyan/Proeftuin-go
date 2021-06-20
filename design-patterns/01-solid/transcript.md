# SOLID Design Principles

## Introduction

All right.
Welcome to the course.

Now before we begin our discussions of the different patterns that no other related topics I wanted to give you some general information about what to expect from this course because obviously we're talking about the Go language.

So there are a couple of limitations that will encounter because well first of all design pattern are typically a object oriented discipline.

Design Patterns are particularly relevant in object oriented languages and of course the problem with go is that it's not strictly object oriented in the sense that you don't have things like inheritance for example the idea of encapsulation and go in is rather weak although it exists.

So you're going to see me kind of trying to adapt go into an object oriented paradigm whereas it's not really entirely object oriented.

Then sometimes this will blur the lines between different patterns and that's just that's just an expected kind of side effect of the fact that we're working with go as opposed to something which allows let's say multiple inheritance like C++.

So one thing you'll see you'll zeal seem a rather permissive approach from my side with regards to visibility and naming.

So I will try to make things consistent and name some things with a lowercase letter and start some things with the uppercase letter but you'll see me sometimes ignore these rules just for the sake of being able to demonstrate a particular pattern without worrying too much about the naming.

I'm also going to be using some terminology which comes from Object Oriented Programming Languages and which might not directly be applicable to go.

So for example I use the term hierarchy and when I'm gonna be talking about hierarchies that go.

That would mean simply a set of related types.

There is no inheritance and so we don't talk about inheritance hierarchies but in go you can have types which share an interface and you can also have embedding.

So when I will talk about hierarchy I will talk about hierarchy in this context.

Also I will talk about properties once again properties are not an idiomatic go construct but property is quite simply a combination of the getter and setter methods and in most cases we'll get away with not using properties.

So there's nothing to worry about but I think there's one or two lectures when we do actually use properties because their use does in fact make sense.

OK.
So a couple of presentation notes about the way that this course is actually presented.

So first of all lectures are going to be presented using go land and Connecticut so Golden is the idea that I will be using but you never actually see the idea itself.

You are only going to see the code because on top of Go land I'm going to be using the kinetic rendering engine which makes sure that the code is rendered nicely and the letters are large and you can read this and follow the courses whether you are using a computer with a large screen or whether you're using a mobile device.

So each of the lectures has a corresponding go file it is attached to that lecture you can downloaded.

You'll have to rename it because it ends with Delta 60.

That's just a limitation of the Udemy platform but you can download each of the go files and you can run it and see what's going on.

Now there are unfortunately no coding exercises and that is once again not my fault.

Typically I add coding exercises for all of my design pattern courses but unfortunately the DME platform does not support go.

So that is sad.

So as soon as it does and you're welcome to write to support it you damage dot com and ask them to add support for go as soon as it does.

I will add the coding exercises but at the moment unfortunately I cannot because there is no support for them.

So all the recordings that I've posted are 10 ADP.

So if you go to full screen you should be able to see them in 10 ADP resolution.

And in addition I want to just to give you a suggestion to consider watching the videos at a higher speed like one point two five or one point five.

You'll find those settings in the player and they will help you digest information a bit faster because I try to talk quickly but you know in some cases I might not.

And so you have the speed control and I suggest you use it if it makes sense for you because it just makes everything it makes your experience a lot better.

So with all that said let's actually jump into the lectures and let's begin by talking about the solid principles.

## Overwiew

So we're going to begin our discussion of design patterns by not talking about design pattern.

So before we get to design partners I want to talk about the solid design principles because they are kind of relevant to the discussions of many of the design patterns that we're going to see in this course.

So what are the solid design principles.

Well the solid design principles are just some of a number of principles introduced by Robert S. Martin also known as Uncle Bob and they were introduced in a set of books that he published on a well.

They use different programming languages so they don't use go.

They use languages like C shop and Java.

But those principles are kind of useful and they are kind of universal.

In fact they're so relevant that they are frequently referred to in design patent literature.

So we're going to be saying things like oh this particular pattern breaks this particular principle and the solid is just five out of a number of principles so here they are.

We have the Single Responsibility Principle the open closed principle the list of substitution principle the interface segregation principle and the dependency inversion principle every single one of these principles also has a three letter acronym that I'll be using so that I don't have to type the whole thing again and again.

## SRP -- Single Responcibility Principle

The Single Responsibility Principle states that a type should have one primary responsibility and as a result it should have one reason to change.

That reason being somehow related to its primary responsibility so let's take a look at how you can both adhere to the Single Responsibility Principle as well as how you can break it.

Let's imagine that you're writing a very simple application you're making a journal so you're making a journal where you're going to record your most intimate thoughts.

This journal is composed out of entries which I'm just going to keep has a string slice so I'll also count the number of entries so we'll have entry count equals zero.

Here we can modify this variable later on and then let's imagine that we want to have some ways of actually adding the entries maybe removing entries stuff like that.

So let's have a function for adding an entry to the journal called add entry.

So we're gonna be providing a string that we're going to add to the Journal and we're going to return just just the position of the entry we're going to assume that this is a.

In reality could be something like a key in a hash map or something.

But let's just let's just stick with the current model so here I would increment the entry count and then we can actually construct the entries so here I can print out some stuff so I can for example as print f.

So that's as print half the index of the entry followed by the entry itself.

So we'll have entry counts followed by text and then we can add it to the entries so I can say eight out entries equals append Jada add entries called that entry and then return the entry count.

So this is the kind of thing that you would add to a journal because it is the primary responsibility of the journal to actually keep the entries and similarly you could have a function for example for removing an entry from the Journal maybe something like remove entry given a particular index way you would also perform some core functionality of the journal.

And this is OK everything is fine so far with simply working with a journal we're sort of adding and removing entries and we are adhering to the Single Responsibility Principle because the Journal has a responsibility of keeping entries and managing those entries so everything is OK but we can easily break this Single Responsibility Principle by adding functions which for example deal with another concern.

There is another term in addition to the Single Responsibility Principle and that is a separation of concerns.

We unfortunately we cannot use the SLC abbreviation because that stands for system on a chip.

But separation of concerns basically means that different concerns or different problems that the system solves have to reside in different constructs of whether attached to different structures or residing in different packages.

It's really up to you but they have to be split up so you cannot just take everything and put it into a single package for example because that is an antsy pattern and this and the pattern is called God object.

So a god object is basically when you take off everything in the kitchen sink that you are doing and put it into a single package for example that would be a terrible idea.

So we want to do separation of concerns.

But let me show you how that can break down.

So let's suppose that you are working with a journal and you decide that you also want to implement persistence you also want to save the journal entries to a file and maybe load the journal entries to a file and ensure you can do this I mean there is it's not exactly difficult to do you make you make some sort of method on the journal called Save where you specify for example the file name.

And here you just go ahead and you use Io you till dot right file for example to save the file name you just take all the entries so you sort of turn all the entries into a bytes and by the way if you want to have the journal turn into a string that is another concern that you can sort of keep in the journal meaning that it's fine to have the string your interface implemented in genuine in actual fact let's implement this.

So I'm just going to quickly have the string or input interface implemented.

So here I would return something like strings don't join.

Just joining every single entry.

So taking all the entries and joining them using some separated like backslash and for example so we would use it here.

We would specify the permissions for the writing to a file.

We could similarly implement methods for loading let's say from a file so you would have something like this Load File Name String.

I'm not going to implement this and you could simply go ahead and expand this idea.

So you might say well let's not only load from a file let's also load from the Web so you could have once again just add a method on the journal called load from web where you would specify the U.R.L.

like so.

And once again you would provide some implementation here.

So what we're doing here is we're breaking the single responsibility principle because the responsibility of the journal is to do with the management of the entries the responsibility of the journal is not to deal with persistence persistence can be handled by a separate component whether it's a separate package or whether for example you want to have an actual struct that has some methods related to persistence.

But the point is you don't necessarily want to have persistence as part of the journo as methods on the journal and you might be wondering well why not.

It seems so convenient to do this but in actual fact imagine that you also have other types in your system and those types also need to be written to a file or loaded from a file and there are some common settings to both the way you load journos and save journals and also the way that you load and save other widgets and types and you want to keep those somewhere.

So that is one of the reasons why you might want to take everything that we're looking at here.

Take all the persistence information and just put it into a separate component whether it's a separate package or indeed a separate separate type.

So let me show you some of the ways that this can work.

So if you were to choose the package approach you would simply have a freestanding function so you would have a function called save to file a way you would take the Journal and the file name and you would perform pretty much the same operation as you would here so we can just copy this over but let's imagine that persistence across your different objects has a couple of settings like for example here when we persist stuff we might want to specify I didn't know a different line separator for the way that things are persisted.

So what you would do is you would specify this as a variable maybe sort of calling it the line separator with some default value like backslash and for example and then instead of using J the string you would do string stuff join and you would actually take the entries so you would say J that entries and then you would use that line separator to persist the entries according to the.

The operating system that you're working with so on windows you would use backslash are backslash and instead so you can see here what's going on so we're separating the concern of persistence into something else into some other component.

So this is no longer a method on journal.

This is just a thing that exists by itself it could exist in a different module called persistence for example and you could have different settings like for example settings for the line separator that you would subsequently use not just to save journals to a file but you to also save other types of objects to a file as well.

So in addition to this what you can do is you can of course turn persistence into a struct so you can say type persistence struct where you have a line separator field line separator string like this and then you would once again implement the safe to file method.

But this time around that will be a method rather than a function.

So you could add a receiver here soapie star persistence and that's that's pretty much all that you would need to do and then you would take the line separator obviously from the persistence object rather than from the global global variable.

So this is another approach to how you can you can deal with all of this.

So just to recap just to see that it works you have your journal.

So the journal is constructed like this you can add a bunch of entries the way that we've set up the API so you can have an entry I cried today you could have another entry where it says I ate a bug and then well we can print these entries so F.A. the print line we can just see that this part works.

So I can I can take for example the journo and they use that Stringer interface sage dot string.

If I run this now you'll see that we are getting pretty much the same output as we would expect.

And so the idea is that instead of doing J don't save when you want to actually persist the Journal to a file for example you either use a separate function so you call save to file where you specify a pointed to the Journal as well as the file name or alternatively what you can do is you can actually have a separate struct that you create.

So p becomes a persistent struct where you specify the line separator as maybe backslash ha backslash and and then you say Pete it's safe to file and you specify pointed to the journal once again and the file name after that.

So just to recap the whole premise of the single responsibility principle is that your class or your type your whatever constructor you're using your package has a single primary responsibility.

So in this case the Journal has a primary responsibility of storing entries and allowing some manipulation of entries like adding or removing and maybe representing those entries as a string and that's it.

And when it comes to other concerns like persistence we adopt the idea of separation of concerns.

So we take those concerns and we put them somewhere else because we realize that those concerns can be cross cutting concerns they can influence not only the way journals are saved but also the way I do know books are saved or manuscripts are saved or some other types of structures are saved so that way you take out the concern and you keep it in one place in that way.

You have common settings like for example line separator settings that you can control not in a million places in your code but in just one place you just control this in a single place in a single setting and you can change the setting and it affects everything that your application does with regards to persistence.

So that's the basic premise of the Single Responsibility Principle.

### solid.srp.go

```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/url"
    "strings"
)

var entryCount = 0
type Journal struct {
    entries []string
}

func (j *Journal) String() string {
    return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
    entryCount++
    entry := fmt.Sprintf("%d: %s",
        entryCount,
        text)
    j.entries = append(j.entries, entry)
    return entryCount
}

func (j *Journal) RemoveEntry(index int) {
    // ...
}

// breaks srp

func (j *Journal) Save(filename string) {
    _ = ioutil.WriteFile(filename,
        []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

var lineSeparator = "\n"
func SaveToFile(j *Journal, filename string) {
    _ = ioutil.WriteFile(filename,
        []byte(strings.Join(j.entries, lineSeparator)), 0644)
}

type Persistence struct {
    lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
    _ = ioutil.WriteFile(filename,
        []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}


func main_() {
    j := Journal{}
    j.AddEntry("I cried today.")
    j.AddEntry("I ate a bug")
    fmt.Println(strings.Join(j.entries, "\n"))

    // separate function
    SaveToFile(&j, "journal.txt")

    //
    p := Persistence{"\n"}
    p.saveToFile(&j, "journal.txt")
}
```

## OCP -- Open-Closed Principle

We are now going to talk about the open closed principle OCP so the open closed principle basically states that times should be open for extension but closed for modification.

What does this all mean.

Well we're going to take a look in in second and in this lesson you're sort of getting two for the price of one because in addition to talking about the open closed principle we're also going to talk about an enterprise pattern called specification.

The reason why we're also going to cover this enterprise pattern is because the discussion of open closed principle is really well illustrated using the specification pattern.

So let's imagine the following scenario.

Let's imagine that you're operating some sort of online store you're selling widgets of some kind just real physical objects.

And let's suppose that you want the end user of your Web site to be able to filter those items like filter by price or filter by size.

And that sort of thing.

So you have a specification you have a sort of product description that states well our Web site has to allow the user to be able to filter by certain criteria.

So let's build a scenario let's suppose that we have a product which is let's say it's composed out of the name.

So the product has a name but it also has a color.

And it also has a size so we can implement these as just ordinary integers.

So I'll have type color ends and then let's have a bunch of constants.

So we'll have red color and green and blue.

And similarly we can have size and we can have once again a bunch of constants for those so we can have a small size and medium and large.

That sort of thing.

So the product is composed out of name and color and size and let's suppose that our first requirement is to create some sort of filter type which has an ability to filter by color.

So we have a type called filter and let's suppose that somewhere in here you would have various filtering settings like for example settings on where you want to filter on the CB on where you want to filter on the GP you or something to that effect doesn't really matter.

But the point is filter has a method called filter by color.

So you have filter filter by color and you let's suppose you just get a slice of the product.

So you have a product as a product slice and you have color and let's suppose that the return type is a just just a slice of product a race product.

Point is product not product or race product pointers.

So here all you need to do is you need to grab all the products as they are provided here find all the products that match a particular color and simply return those as effectively an array of pointers.

So we'll have results being make let's make this array and then for I call my V in our range products.

All we need to do is we say if video color is equal to color and I think there's a mistake here somewhere let me just take a look what's what's going on here.

Oh yeah.
The ones that curly here.

So if you know color equals color then let me just realign everything so everything is nice and neat here.

So video color is equal to color then we add it to the results or result is equal to append result and we get the pointed to the right product products at index i.

And then we return the result.
Okay.

So you've got this implementation of filtering by color and this is something that you can already use straight away.

So let's set up a very simple scenario for just just a bunch of objects which all have different size and color.

So we'll have an apple which is going to be a product called Apple and it's going to be green and small.

We'll also have a tree to be a product called tree which is going to be green and large.

And finally we'll have a house.
So let's just magic.
You can buy entire houses online.
So you have a product called a house which is blue and large.

Now it goes so I'll make a slice out of those so I'll say products equals not a slice of products.

So Apple tree how what we can do now is we can somehow try filtering this whole thing.

So we have our type of cold filter and we have a method called filter by color which gives us all the right products.

So let me just F.A. you ought to print Taff here.

So I'm going to say that we're going to get all the green products and that's going to use the old method because subsequently we'll change everything and we'll do things completely differently.

So here I can say for on this score column A V in a range of.
So we we need to create the actual filter somewhere.
So here I'll say F is a filter.
Okay.

So here you would say effed up filter by color you specify the products and you specify the color green.

And that's how you're going to get your result.

So here we can say well let's use F.A. FMC the print f once again.

So here I'm just going to put a dash and then the name of the product is green and then a line break and that's put Vito name here.

So that's not say it really much.

We can we can actually run this we can just get to see what the output is.

So as you can see the system is finding two products is finding and an apple and the tree and both of those are green.

Surprise surprise.
OK.
So it might look like everything is OK.

However imagine that you implement this filtering by color and it goes into production and then your manager comes back and they say something like Well you know we also need filtering by size.

So what does this mean.

What does this mean that you also need filtering by size it means that you have to jump into a type.

So we have a product type and we have some methods on it you have to go back to that type.

You have to basically change the type adding another method to it.

So we have yet another method here filter by size for example you specify the size size here and then here once again you say v dot size equals size.

So you go ahead you you do this thing and yeah it does modify the actual type the filter type because it now has a new method.

But apart from that everything is okay or at least it looks okay.

So sometimes it passes and then your manager comes back and they say something like well we've got filtering by color and we've got filtering by size.

Can you please implement filtering by size and color.

And so you say well okay if that's what you want.

You know I can just cut and paste once again so we can grab this method we can paste it here we can say filter by size and color.

And now of course we also have to add color color here and we have to add a check here so v dot color is equal to color.

But apart from that yeah everything continues to work.

So what I'm showing you right now is a violation of the open closed principle because what we're doing as we're going back into the product type and we're modifying we're adding additional methods on the product is we are sort of interfering with something that's already been written and already being tested and the open closed principle is all about being open for extension.

So you want to be able to extend a scenario but by adding additional maybe additional types maybe additional just freestanding functions but without modifying something that you've already written and you've already tested.

So basically what I'm saying is you really want to leave the filter type alone you want to leave the filter type alone you don't want to come back to within keep adding more and more methods to it.

And and all that sort of thing you want to basically have some sort of extendable setup and that's exactly what we can do.

What we can get if we use the specification pattern so the specification pattern is somewhat different because it has a bunch of interfaces it has a bunch of additional elements done for flexibility.

So let's let's start implementing specification the first thing you do is you implement the specification interface so the specification interface is just going to have one method called is satisfied and you provide P which is a product pointer and it returns a boolean.

So the idea behind the specification interface is you are testing whether or not a product specified here by our pointer satisfies some criteria.

So for example if you want to check for color you would make a color specification make a color specification way you'd specify the color you want to filter on for example and then you would have a function defined on the color specification so that you conform to the specification interface so it would be called is satisfied.

Let me actually just just generate this quickly.
So on the color specification we want the specific.
That's not the right pattern.
Let me try this again on college specification we want the specification pattern.
There we go.
Okay.

So here I've generated the sort of placeholder for checking whether or not the product matches a college specification.

Here I just check that pedo color is equal to us.
See the color.
That's pretty much it.
That's all that you would need to do.
And you have a color specification.

So in a similar fashion what we can do here we can sort of cut and paste here and we can make a size specification size specifications so here you would say size size and obviously here you would have a size specification like so and you would say Pitot size equals size this way or s that size rather.

So this is how you said the whole thing up.

Now we're going to build a different filter so remember we have that filter type.

Now we'll have a time called better filter so better filter is a different time because it's the type that you're unlikely to ever modify.

In actual fact if you don't have anything any settings for the better filter you don't even need a type you can just have a freestanding function.

But I'm just doing it as a type just just to illustrate things.

So basically the idea is that given a better filter you also specify a method for filtering on products.

But this time around you just have to do it once so you have a function which takes a better filter.

So the receiver's better filter is called filter.
And here you take two things.
So you take the products.

So just a slice of products and you also take the specification so you take the specification and what you return is once again just just an array of product pointers.

There we go.

So the implementation of this is very similar to the implementation of the way that we've done this before.

The only difference is that now we have this specification object that we're working with.

So let's have a result make.
So then let's iterate the products.

And for each of the products we basically take the specification and we use that is satisfied member in order to check whether or not we are satisfying a particular specification.

So here we say you've spec his satisfied value then result equals resolved on a Pens result.

Well result equals append result comma products at position i.

And then we return the result.
That's pretty much all that you would do.
So this is the the setup that we have.

And let's actually take a look at how to use all of this because now things do become a bit more involved because you have to make that specification but they become a lot more flexible.

So let's find all the all the green products in this new way.
So green products in the new implementation.
OK.
So here.
Before doing anything you have to make that specification.

You have to say green back is a color specification with a color green.

But apart from that the actual implementation here is very similar.

So instead of filter away using better filter and now what what we can do is we can.

Well actually that that that's called B.F..

So here instead of filtered by color we call obviously just filter.

And that's that's pretty much it.

Now instead of passing green we pass the green specification like so.

OK.
So let's run this let's just see that the new approach does in fact work.
OK.
So as you can see we're getting identical results on both of these.
So everything is fine.

But the second approach the approach with the specification pattern gives you more flexibility because if you want to filter by a particular new type all you have to do is you have to make a new specification.

So for example here we have a color specification but you decide that you want to filter by size so all you have to do is make a size specification and make sure that it conforms to the specification interface.

That's pretty much all that would have to do.

And that follows the open closed principle so the types in this case that the interface type is open for extension meaning you can implement this interface but it's close to a modification which means that you are unlikely to ever modify the specification interface and in a similar fashion you are unlikely to ever modify better filter because there is no reason for us to do so it's very flexible.

It takes a bunch of products and a specification then that's pretty much all there is to it.

Now of course a different question is what if you want to filter something other than products.

Well yeah.

Then you can either extend better filter or you can make a new type but it at no point in time are you modifying something some structure that you've already created before.

So one thing you might be wondering is well hold on up here.

We have this method for filtering by size and color.

How would you implement this using the specification pattern.

Well it's actually not that difficult because all you have to do is you basically have to make a composite specification and that is and that is actually an illustration of the composite design pattern that will also talk about as part of this course.

But let me show you how you can actually work.

So basically a composite specification is just a combinator it just combines two different specifications so you'd make a type called an specification.

And similarly you could have or specification for example so you'd have this shrugged.

And here you would specify the first and second specifications life's like so and then you would once again go ahead and implement the specification interface on an specification.

So let me just kind of generate this behind the scenes.

So here in order to check that and specification isn't fact satisfied you have to check first and second.

So here we return a dot.

First is satisfied with the product and a second is satisfied with the product as well that that's all that you really have to do.

So let me show you how this can work.

Let's suppose that we want to find all the large blue items so large green items.

So we are they have a green specification we can also have a large specification.

So that's going to be a size specification of large.

And then we can make a large green specification large.

Let's just go with our G spec large green specification.

So there's gonna be a and specification where you specify Greens back and large spec.

So you combine the two specifications together and once again we can do this whole thing.

We can do the printout obviously.

So here we can find all the large green products and the way to do this is a you would once again write a for loop.

So here we would say for on this Gore column A V in range better filter dot filter.

And here you would provide the products and you would provide that large green specification and then you would go through all the products and you would once again print them out as we did before our copy this over and put it here so we can say this is large and green.

All right.

So let's actually run this let's see this final piece of code in action and you can see that we're getting large green products well there's really only one there's the tree so we're getting exactly that.

So what is the takeaway from this entire lesson.

The open closed principle basically states that types are open for extension so you can for example grab an interface and implement this somewhere in your code.

But they are closed for modification which means that once you've designed and tested the API of a particular type you shouldn't really keep jumping into it and modifying it and modifying it again because then you end up.

Well first of all you end up putting too much into one time.

Possibly but also it's just very inconvenient because well you've already got that type working you've either got clients relying on that you don't want to surprise them with additional methods for example.

So it's better to just use this this idea of fudge just implementing interfaces and making new types rather than just working with a single type and just extending it over and over and over again.

So that's pretty much the gist of the open closed principle.

### solid.ocp.go

```go
package main

import "fmt"

// combination of OCP and Repository demo

type Color int

const (
    red Color = iota
    green
    blue
)

type Size int

const (
    small Size = iota
    medium
    large
)

type Product struct {
    name string
    color Color
    size Size
}

type Filter struct {
}

func (f *Filter) filterByColor(
    products []Product, color Color)[]*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if v.color == color {
            result = append(result, &products[i])
        }
    }

    return result
}

func (f *Filter) filterBySize(
    products []Product, size Size) []*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if v.size == size {
            result = append(result, &products[i])
        }
    }

    return result
}

func (f *Filter) filterBySizeAndColor(
    products []Product, size Size,
    color Color)[]*Product {
    result := make([]*Product, 0)

    for i, v := range products {
        if v.size == size && v.color == color {
            result = append(result, &products[i])
        }
    }

    return result
}

// filterBySize, filterBySizeAndColor

type Specification interface {
    IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
    color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
    return p.color == spec.color
}

type SizeSpecification struct {
    size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
    return p.size == spec.size
}

type AndSpecification struct {
    first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
    return spec.first.IsSatisfied(p) &&
        spec.second.IsSatisfied(p)
}

type BetterFilter struct {}

func (f *BetterFilter) Filter(
    products []Product, spec Specification) []*Product {
    result := make([]*Product, 0)
    for i, v := range products {
        if spec.IsSatisfied(&v) {
            result = append(result, &products[i])
        }
    }
    return result
}

func main_() {
    apple := Product{"Apple", green, small}
    tree := Product{"Tree", green, large}
    house := Product{ "House", blue, large}

    products := []Product{apple, tree, house}

    fmt.Print("Green products (old):\n")
    f := Filter{}
    for _, v := range f.filterByColor(products, green) {
        fmt.Printf(" - %s is green\n", v.name)
    }
    // ^^^ BEFORE

    // vvv AFTER
    fmt.Print("Green products (new):\n")
    greenSpec := ColorSpecification{green}
    bf := BetterFilter{}
    for _, v := range bf.Filter(products, greenSpec) {
        fmt.Printf(" - %s is green\n", v.name)
    }

    largeSpec := SizeSpecification{large}

    largeGreenSpec := AndSpecification{largeSpec, greenSpec}
    fmt.Print("Large blue items:\n")
    for _, v := range bf.Filter(products, largeGreenSpec) {
        fmt.Printf(" - %s is large and green\n", v.name)
    }
}
```

## LSP -- Liskov Substitution Principle

The Liskov substitution principle which is named after Barbara Liskov isn't really that applicable to go because it primarily deals with inheritance.

So what the list of substitution principle basically states that if you have some API that takes a base class and works correctly with that base class it should also work correctly with the derived class.

But unfortunately in go we don't have base classes in derived classes.
This concept simply doesn't exist.

But I am going to try and show you a variation of the list the substitution principle that does in fact apply to go.

So let me show you an example let's suppose that you're trying to deal with geometric shapes of a rectangular nature and you decide to have an interface called sized that allows you to specify certain operations on those types of constructs like for example you might have the getters and setters for the width and the height.

So you would have an interface where you would have get with you would have a set with you would have get height and maybe set height.

So now let's suppose you have a type called rectangle so a rectangle obviously has a width and height.

We can actually store them in just just ordinary fields.

I use lowercase here and then what you can do is you can go ahead and implement the sized interface on a rectangle like so.

Except what I would typically go for is having the pointer types here is gonna help us out later on when you work with the interface.

Obviously when you work with interfaces it's a lot more convenient to have pointers here.

So I hear you would return a width here you would say are width equals width here you would return height.

And here you would say our height equals height like so and so far so good.

I mean this is something that you can work with and we can actually start experimenting with it.

So for example I can write a function called use it which uses some sized object sized sized.

So let's take a look at how this can work.

Let's suppose that I first get the width of this object so width is equal to size they'll get WIDTH.

OKAY SO THAT'S GOOD.
AND THEN LET'S SUPPOSE I set the height to 10.
So sized dots at height.
Let's add it to the value 10.

Now if you were to calculate the area of this sized object you would expect the area to be 10 multiplied by the width.

Because I've just said hi to 10 and we got the width right here so the expected value of the expected area would be 10 multiplied by the width.

But we can also calculate the actual area.

So the actual area would be size they'll get width multiplied by size and get height like so.

So we have an expected area and we have an actual area and we hope that they are the same value because if they are not that would be a big problem.

So here I would just do F.A. that.

So we expected an area of the expected area but instead we got actual area like so.

OK.

So we can already try out this entire thing just just create a rectangle and try plugging it into this.

This use it a function that we created.
So let's make a rectangle.
Let's put an ampersand here rectangle two by three.
And then we can call use it with this rectangle that are OK.

So running this.
Let's take a look at what we get.
So we expected an area of 20 and we go to an area of 20.
Now it looks like everything is correct.

Everything is working fine and until you try to break the list of substitution principle everything is OK but let's imagine that you decide to be smart and you decide to make a type called Square now a square operates just like a rectangle because well it has the same members.

So you can say something like type square struct which simply aggregates a rectangle because Square also has the width and the height.

However let's imagine that you decide in your infinite wisdom that a square is going to enforce this idea of with being equal to height always.

So you're always going to enforce this.

So you will have some sort of constructor so a new square which takes a size like so so here you would say Eskew is a square and Eskew width equals size as gear height equals size return.

Ask you something like that but in addition what you'll do and this is really the part that violates the Liskov Substitution Principle is you'll have methods for set width and set height that set both the width and height in both of those situations so you'll have a function which takes a square which has set width and here is the insidious part.

So not only do you set the width but in order to keep this a square.
You also set the height and then you decide to do the same thing for the height setter.
So you decide that you're going to have set height where are you now.
But the height here where you set the width and height to the height value.
OK so what is the problem with this.

Why is this breaking the Liskov Substitution Principle.

So the risk of substitution principle basically states the following if you were expecting some sort of something up the hierarchy so to speak in.

In this argument then it should continue to work even if you proceed to extend objects which are already size.

So we took a rectangle and we decided to sort of extend the rectangle and make a square.
It should also continue to work.
Everything here should continue to work but unfortunately it doesn't.

And if we try to plug in the square you'll see how it badly goes wrong because if I make a square over the side of five for example and I call use it we expect this whole thing to actually work.

We expect all of this to operate correctly but if I now run this you'll see that the results are a bit disappointing.

We expected an area of 50 and we got an area of one hundred.
So what happened here.

Well obviously what happened is this call to set height actually set not just the height.
It also set the width.

So the internal width of the square became inconsistent with the value of this variable right here and as a result we're getting different values for expected area and actual area.

So our telescopes that petition principal basically states that if you continue to use generalizations like interfaces for example then you should not have inheritance or you should not have implementers of those generalizations break some of the assumptions which are set up at the higher level.

So at the higher level we kind of assumed that if you have a sized object and you said its height you are just setting its height not both the height and the width and here what happened is we broke this assumption by setting both the width and the height which is a noble goal.

I mean you can see how somebody would try to enforce the square invariant by setting both within the height.

It's a noble goal but it doesn't work.
And it actually violates the Liskov Substitution Principle.

So the Liskov Substitution Principle is actually one of those situations where there is no right answer there is no right solution to this problem.

I mean you can take different approaches to how you would take this for example we can say that squares don't exist that since every square is a rectangle we don't work with squares at all.

Or for example you could do is you could explicitly make make illegal states unrepresented all so to speak.

So basically you can say that a square doesn't really have width and height square has a size and that's pretty much it.

So you could have some time called Square 2 which would have a size which is an int which would double as both a width and height.

So it would represent both of these states so instead of aggregating a rectangle you have its own member.

And then if you want to represent the square as a rectangle well you can have your cake and eat it too.

You can have a method which just is called a rectangle for example which returns a rectangle.

And here you would just construct a new rectangle with as that size comma as that size.

That's pretty much all that you would have to do so this is one approach to the problem.

But just just to recap basically the idea of the risk of substitution principle is that the behaviour of implementers of a particular type like in this case the sized interface should not break the core fundamental behaviours that you rely on.

So you should be able to continue taking sized objects instead of somehow figuring out in here for example by doing tie checks whether you have a rectangle or a square it should is still work in the generalized case.

And so a prime example of the violation of this problem is what we've done here.

So we've broken certain assumptions about the type and as a result we got incorrect behaviour.

So that's what I wanted to show you about the list of substitution principal.

### solid.lsp.go

```go
package main

import "fmt"

type Sized interface {
    GetWidth() int
    SetWidth(width int)
    GetHeight() int
    SetHeight(height int)
}

type Rectangle struct {
    width, height int
}

//     vvv !! POINTER
func (r *Rectangle) GetWidth() int {
    return r.width
}

func (r *Rectangle) SetWidth(width int) {
    r.width = width
}

func (r *Rectangle) GetHeight() int {
    return r.height
}

func (r *Rectangle) SetHeight(height int) {
    r.height = height
}

// modified LSP
// If a function takes an interface and
// works with a type T that implements this
// interface, any structure that aggregates T
// should also be usable in that function.
type Square struct {
    Rectangle
}

func NewSquare(size int) *Square {
    sq := Square{}
    sq.width = size
    sq.height = size
    return &sq
}

func (s *Square) SetWidth(width int) {
    s.width = width
    s.height = width
}

func (s *Square) SetHeight(height int) {
    s.width = height
    s.height = height
}

type Square2 struct {
    size int
}

func (s *Square2) Rectangle() Rectangle {
    return Rectangle{s.size, s.size}
}

func UseIt(sized Sized) {
    width := sized.GetWidth()
    sized.SetHeight(10)
    expectedArea := 10 * width
    actualArea := sized.GetWidth() * sized.GetHeight()
    fmt.Print("Expected an area of ", expectedArea,
        ", but got ", actualArea, "\n")
}

func main() {
    rc := &Rectangle{2, 3}
    UseIt(rc)

    sq := NewSquare(5)
    UseIt(sq)
}
```

## ISP -- Interface Segregation Principle

The interface segregation principle is a really simple principle is rarely the simplest principle out of the solid design principles.

Basically what it states is that you shouldn't put too much into an interface.

You shouldn't try to throw everything and the kitchen sink into just one single interface and then sometimes it makes sense to break up the interface into several smaller interfaces.

So let me show you a very simple somewhat contrived example let's suppose that you have some sort of document type so just some information about the documents and you want to make an interface that allows people to build the different machines different constructs for operating on the documents so doing things like printing the document or scanning the document or sending the document as a fax.

That sort of thing.

So one approach you might take is just make just a single interface type machine interface.

And in this interface you would have methods for printing the documents and also maybe faxing the document and scanning the document.

So this is generally okay it's an okay interface.

If what you are looking for is a kind of multifunction printer so if you have a multi function printer which can both scan and print and fax documents then everything is OK.

You simply implement the struct and then you go ahead and you actually implements all the all the interface members.

So you just grab the machine interface and you generate basically all of this stuff for the printing and for the faxing and for the scanning as well.

So there's absolutely no problem here.

However imagine a different situation imagine a situation where somebody is working with an old fashioned printer an old fashioned printer doesn't really have any scanning or faxing capabilities but because you want to implement this interface because maybe some other API is rely on the machine interface you have to implement this anyway.

You are being forced into implementing it so you go ahead through all the similar motions to implement the machine interface and you end up with the same stuff as you would for a multifunction device.

Except there is a bit of a problem.

So certainly when you're working with an old fashioned printer implementing the printing capability makes sense because a printer can print that.

That's what it does but it doesn't.
You don't really know what to do.
In the case of the faxing and in the case of this scanning as well.

So one thing you can do is you can certainly leave the panic messages in here except that you would probably say something more meaningful like operation not supported because that's really what's happening here it's not the case of implements me it's just that we don't support scanning from an old fashioned printer.

You can also as an additional measure you can add a comment which begins with the word deprecated.

So once again these methods they're not really deprecated where lying to the user a little bit.

But the consequence of having deprecated here would be that if you're using it let's say you made this old fashioned printer you made this old fashioned printer like so and then you try to do.

Oh p dot dot scan.

Some I.D. will actually cross out the scan option and they will tell you that this method is deprecated.

You shouldn't be calling it.
So that's certainly what happens in my idea.

So that's one way that you could deal with this situation but really we've created the problem by putting too much into an interface.

So we put both print and facts and scan into a single interface and then we expect everyone to implement this even if people don't actually have this functionality as part of the classes.

So they want the support of this interface because perhaps the interface is used in some sort of APIs but they really don't have anything to put into some of the implementations.

So how can we deal with this well we adhere to the interface segregation principle.

So the interface segregation principle basically states that tried to break up an interface into separate parts that people will definitely need.

So there is no guarantee that if somebody needs printing they also need faxing.

So in this particular example it might make more sense to split up the printing and scanning into separate interfaces so you would have maybe the printer interface where you would have the printing method.

You would also have the scanner interface where you would have the scan method and so on and so forth.

And this way this allows you to compose different types out of interfaces that you actually need.

So for example if you just want something which is only a printer.

So this is only a printer and nothing else it doesn't scan and doesn't do anything in this case what you do is you simply implement the printer interface.

So you take that printer interface and you implement it like so everything is okay.

You don't have to implement scan you don't have to implement fax.
Everything is fine.

And then let's suppose that for example you have a photocopier.
So a photocopier can both print as well as the scan.

So all you do here is you go ahead and you you implement the printer interface like so and you also end the same time implement the scanner interface.

And that way you you basically get both of the functionality and you you now have the photocopier which is both a printer as well as a scanner.

So let's not forget the fact that you can actually combine interfaces so you can compose an interface out of other interfaces so if you want an interface that represents a multifunction device then you can have your cake and eat it too because you can make a type called Multifunction Device which is an interface and into this interface you can put all the stuff from a printer and all the stuff from a scanner.

And you know if you had like a fax for example you would add that interface here as well so combining interfaces is fairly easy and not really complicated.

And of course what you end up with if you adopt this approach is if you want to build a multifunction machine and you've already got let's say a printer in a scanner implemented as separate components what you can do is you can use the decorator design pattern and we'll talk about when we get to the decorator design pattern.

But let me show you how this would work.

So if you want to build some sort of multi function machine which is both a printer and a scanner what you can do is you can simply have the print apart and you can have the scanner apart like so and then you can implement the necessary interfaces so you would have a function which takes a multi function machine for printing for example where you would simply reuse the functionality of the printer that you already have so you would say amd up printer not print the document and you would do the same thing for for the scanner.

So you would have a method called scan which would just say am dot scanner dot scan and pass that in so you can see that with the interface segregation approach what you can do is first of all you have very granular kind of definitions.

So you just grab the interfaces that you need and you don't have any extra members in those interfaces.

So if you're just building an ordinary printer you just get the print method and that's pretty much it.

And of course you or you always have an ability of combining the interfaces.

So here you can combine the printer in the scanner and have a kind of interface aggregate and you can subsequently have APIs which actually use this interface aggregate in your code.

So that's all there is to be said about the interface segregation principle.

### solid.isp.go

```go
package main

type Document struct {
}

type Machine interface {
    Print(d Document)
    Fax(d Document)
    Scan(d Document)
}

// ok if you need a multifunction device
type MultiFunctionPrinter struct {
    // ...
}

func (m MultiFunctionPrinter) Print(d Document) {
}

func (m MultiFunctionPrinter) Fax(d Document) {
}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
    // ...
}

func (o OldFashionedPrinter) Print(d Document) {
    // ok
}

func (o OldFashionedPrinter) Fax(d Document) {
    panic("operation not supported")
}

// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
    panic("operation not supported")
}

// better approach: split into several interfaces
type Printer interface {
    Print(d Document)
}

type Scanner interface {
    Scan(d Document)
}

// printer only
type MyPrinter struct {
    // ...
}

func (m MyPrinter) Print(d Document) {
    // ...
}

// combine interfaces
type Photocopier struct {}

func (p Photocopier) Scan(d Document) {
    //
}

func (p Photocopier) Print(d Document) {
    //
}

type MultiFunctionDevice interface {
    Printer
    Scanner
}

// interface combination + decorator
type MultiFunctionMachine struct {
    printer Printer
    scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
    m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
    m.scanner.Scan(d)
}

func main() {
}
```

## DIP -- Dependency Inversion Principle

We're now going to talk about the dependency inversion principle and the first thing I wanted to say is the dependency inversion principle doesn't have anything directly to do with dependency injection.

Those things need to be kept separate.

They do have a relationship but we're talking about the principle now and not the not the actual mechanism.

So what does the dependency inversion principle actually state.

Well it states two things it states that high level modules should not depend on low level modules and then both of them should depend on abstractions.

So what does this mean.

It sounds really cryptic and I'll try to show you what this means so let's imagine that you're doing some sort of genealogy research and you want to model relationships between different people.

So we'll have different types of relationships so I'll have a relationship just being in and we'll have a bunch of constants here.

So one person can be a parent of another person.
So that's one possibility.

Or you could be a child or you could be something else like a sibling and so on and so forth.

And then we can model different people so here a person would be a struct you would have a person's name and you would also have other attributes like date of birth and other types of information that we're not going to care about in this demo.

So how do we model relationships between different people.

If you have people who are parents to one another one person is a child of another.

Well you could have some sort of info struct away you have a relationship so you model of relationships from one person you specify the type of the relationship relationship like so and then to another person so you can for example say that John is a parent of Jill that's other thing.

Now let's take a look at how you would actually go about storing this information.

So you want to have all the data about the relationships between different people stored in some sort of type.

So you have a type called relationships which would just have relations as an info slice.

OK so what do we want to get from this time.

Well maybe we want to for example find all the children of a particular person just as a kind of as a kind of research.

But the first thing we need to do is we need to be able to add those relationships so we can have a function relationships called add parent and child.

So you specify the parent and the child and they are both person pointers.

And let's let's actually go ahead and implement this.

So here I can say are dot relations equals append either relations and then create this new info where you have parent being the parent of a child.

You can also do it the other way round by the way.

So we can you can sort of duplicate this and I can I can say for example that child so child here is a child of parent that sort of thing doesn't really matter it's not so important for our purposes.

But so what we're doing is we have this kind of storage element for storing the relationships between different people at the moment it's a slice.

But it could be something else it could be some specialized collection for example.

And then what we want to be able to do well we want to be able to perform some sort of research on this data.

So we want to be able to research perform research on this date and the question is well how can we actually do this.

So first of all research in our model is what you would call a high level module and relationships is a low level module.

The reason why it's low level is because it kind of storage.
Now I mean this could be in a database or something.
It could be on the web or somewhere.

So it's basically the storage mechanism and this is the high level module decide designed to operate on data and perform some sort of research.

So here what we can do is we can break the dependency inversion principle.

Remember the dependency inversion Principle states at high level modules should not depend on low level modules.

So we're going to break this principle by depending on a low level module.

So here I'll say relationship relationships relationships.
And so we have a high level module depending on a low level module.
And this does work.

This does allow us to for example perform research so we can have some sort of function called investigate where we say relations.

So in order to be able to actually perform any research we have to go into relationships and look at the relations.

So we would say relations and that will be defined as our adult relationships not relations.

So you can see us getting getting rather convoluted and then we can go through them all.

For on this call come around in arrange relations we can for example find all the children of John.

So if a real dog from the name is equal to John and Ralph that relationship dot is equal to parent.

Then we have John's child so we can print line.
For example John has a child called
Real dot to that name.
So this actually works.

I mean what I've written right now is going to work just fine.
So if you if you want to see it in action so to speak.
Let's let's do this.

So you would basically create this research module so you would say aha.

Is equal to research and then you would say ah dot investigate pretty much it.

Let's actually run this.
Let's see what we get here OK.
So.

So we're not getting anything because we haven't actually set up the the people the people that are participating in this research so let's do that quickly.

So I'll say parent is a person with the name John and then we'll have child one being a person within name Chris.

And that's half child 2 which is going to be a Matt for example.

And then we can add those people so we can create the low level path relationships equals relationships like so and we can call relation ships dot at Parent and Child with parent and child one we can duplicate this and our child too as well.

And now we can perform the research and now research is going to rely on that low level module relationships and we can actually do the investigation.

So this this time round hopefully everything works correctly.

Let me just verify that the the research mojo is specify correctly because it seems that it seems that it might not be so like this.

Okay so let's let's just run this.

Let's see what we get here and we get the right result.

So John has a child called Chris and John has a child called Matt so everything is OK everything is fine so far but there is a major problem with this entire scenario because at the moment what's happening as you can see is the research module here is actually using the internals of the relationships module so relationships is a high is a low level module and it's using literally its slice to get data from that slice.

Now imagine if relationships the low level module decides to change the storage mechanic from a slice to let's say a database.

So what happens then.

And the answer is that the code which depends on the low level module actually breaks.

So all of this is going to break because for example you can no longer use a for loop here.

So that's obviously something we want to avoid.

And that's what the dependency inversion principle is trying to protect us from these situations where everything breaks down.

And in actual fact it could be argued that the finding of a child of a particular person is something that needs to be handled not in a high level module but in a low level module because essentially if you know the storage mechanic you can do an optimized search like for example if this was a database you would just select all the people where parent is John and so on and so forth and you could optimize those queries.

So what we're going to do is we're going to rewrite this code to adhere to the dependency inversion principle.

How do we do this.

Well first of all we know that the dependency doesn't have to be on a low level module directly.

Remember the second part states that both should depend on abstractions and that's exactly what we're going to be doing here at least at the high level because we're going to depend on some sort of an abstraction.

And I'm going to call this abstraction a relationship browser.

So we'll have a browser which is a relationship browser.

So we're going to avoid the situation where we're exposing the internals of relationships and instead will define this new interface called relationship browser and will implement this interface on relationships.

So let's put it here so we'll have a time called relationship browser and it's just going to have a single method for now called Find All children of some person whose name we have.

So it's just going to give us a bunch of pointers to the children of whoever we specify right here.

So now what we can do is we can take relationships and we can actually implement this interface on it.

So we take relationships and we implement the relationship browser interface for some reason it got constructed.

Here let me just just move it as always to a better position somewhere right here for example.

OK.

So we want to find all the children of relationships I'll use I'll actually use a pointer type here.

It's much better for future use.

So we want to find all the children of a particular person now here because we are currently in the low level module we can actually depend upon the internal mechanics of how you would go about accessing that low level storage.

So here I can say a result is what actually we can sort of try Can we try copying some of this.

Well actually no.

Let's do this from the ground up so a result is going to be an empty slice of person pointers.

And then for I call V and arrange the relations.
Actually that should be relationships with an S. ah dot relations.
Here we go.

So we say if v dot relationship his parents and V.

From name is equal to name then we return the point it to the child.

So resolve equals circle that should be an ampersand here.

So result equals append the result.
And then just get the relationships and take the relation.
So the low level parts and get index i dot to.

So that way we get we get an actual pointer to the thing and then we return the result.

So now all the finding of the children is actually put into the low level module and then we can rewrite the high level mojo so the high level module now depends on an abstract just like we wanted.

We depend on the relationship browser and then the investigation becomes different because the search for the children is actually done in the low level path and all we have to do is we have to handle it somehow so we have to perform the actual investigation so here what I can say is I can say for on the score column A P in range in a range R dot browser dot find out children office on No this way using the browser now we're using that interface member.

So we're looking for all the children of John and once we find them we can actually print line something.

So F.A. that print line John has a child called and then just pedo name.

So everything became much simpler.

And also you'll notice that the low level implementation details are not exposed to the high level research module we're just using the browser abstraction to find all the children with a particular name.

And then we can jump down here and we can change all of this and now research depends on relationships as an interface.

So let's put the ampersand in here.

But apart from that everything stays pretty much the same and we can run this and hopefully get the same result and as you can see we are getting the same result.

Okay.
So that's the dependency inversion principle.

Basically it states that high level modules should not depend on low level modules typically by low level we mean kind of closer to the hardware sort of data storage and you know communication sort of system level stuff and high level would be the business logic stuff.

And both of them should depend on abstractions by abstractions we typically mean interfaces at least in go in other languages you would talk about abstract classes and base classes in go that would typically imply interfaces so we would depend on interfaces rather than concrete classes and this does require certain small modifications like that ampersand that we put right here because when I'm using an interface type.

But apart from that it's a really simple transition from one to another and you are protecting yourself against changes because now for example if we decide to change the storage mechanic of relations from a slice to something more sophisticated then you would only be modifying the methods of relationships you would not be modifying the methods of for example research because it doesn't depend on the low level details so that's it for the dependency inversion principle.

### solid.dip.go

```go
package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions

type Relationship int

const (
    Parent Relationship = iota
    Child
    Sibling
)

type Person struct {
    name string
    // other useful stuff here
}

type Info struct {
    from *Person
    relationship Relationship
    to *Person
}

type RelationshipBrowser interface {
    FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
    relations []Info
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
    result := make([]*Person, 0)

    for i, v := range rs.relations {
        if v.relationship == Parent &&
            v.from.name == name {
            result = append(result, rs.relations[i].to)
        }
    }

    return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
    rs.relations = append(rs.relations,
        Info{parent, Parent, child})
    rs.relations = append(rs.relations,
        Info{child, Child, parent})
}

type Research struct {
    // relationships Relationships
    browser RelationshipBrowser // low-level
}


func (r *Research) Investigate() {
    //relations := r.relationships.relations
    //for _, rel := range relations {
    //    if rel.from.name == "John" &&
    //        rel.relationship == Parent {
    //        fmt.Println("John has a child called", rel.to.name)
    //    }
    //}

    for _, p := range r.browser.FindAllChildrenOf("John") {
        fmt.Println("John has a child called", p.name)
    }
}

func main() {
    parent := Person{"John" }
    child1 := Person{ "Chris" }
    child2 := Person{ "Matt" }

    // low-level module
    relationships := Relationships{}
    relationships.AddParentAndChild(&parent, &child1)
    relationships.AddParentAndChild(&parent, &child2)

    research := Research{&relationships}
    research.Investigate()
}
```

## Summary

So let's summarize what we've learned about these solid design principle so we talked about the Single Responsibility Principle basically this idea that a type should have just one primary responsibility and therefore should have just one reason to change.

We also talked about the idea of separation of concerns.

The idea that if you have different concerns different areas of responsibility in the system then you should probably be putting those into different types or different packages rather than trying to stick everything into a single type.

Then we talked about the open closed principle the idea that types should generally be open for extension and by extension we mean things like those types being aggregated being used inside other types but the types should be closed from modification.

Basically this idea that if you've written and tested your code you shouldn't jump back into the code and modify it if it's possible to extend it instead.

Then we talked about the list of substitution principle.
This one is a bit tough.

It's a bit difficult to describe this in go as opposed to other languages but the idea is that if you have some type which kind of aggregates another type and thereby acquires all of its methods and so on then any API which uses the type you've aggregated should also be able to take the the elements from your current type and the explanation of this is actually a bit difficult.

So I suggest if you don't get this look at the code again because it's much better explained in code than explained in text then we talked about the interface segregation principle.

This is an easy to understand principle for a change.

Basically the idea is that you shouldn't be putting too much into an interface and you should probably split split your API into interfaces if for example your interface is getting too big for some reason.

There's also another acronym shall we say called Jani stands for you in going to need it and it's very relevant and when discussing the interface segregation principle because basically if you put too much in an interface you'll force people to implement things that they are not going to need and which is always a problem.

And finally we talked about the dependency inversion principle the idea that high level modules should not depend upon low level modules that both should depend on abstractions and typically by abstractions we mean interfaces because there isn't much choice in terms of the language features inside the go language.

So these are the five solid design principles and I will be referring to them quite a lot as we discussed the different design patterns.
