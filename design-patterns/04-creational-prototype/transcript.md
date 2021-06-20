# Prototype -- when it's easier to copy an existing object to fully initialize a new one

## Overview

All right.
Let's talk about the prototype design pattern.

So apart from being a really cool game now a bit outdated but prototype was a really good the game but prototype is also a design pattern and it's all about this idea of how complicated objects are constructed.

So if you think about objects in the real world like cars and iPhones and whatever they are never designed from scratch.

Instead what happens is you iterate existing design say you already got iPhone twelve and you're making iPhone 13 or something.

So you take an existing design and you try to improve some things may be changing slightly so everyone can see that you've got an updated model and then then you produce on the basis of that you don't design everything from scratch.

And so the idea here is that you take an existing design and it can be like a partially constructed older fully constructed design and that is your prototype.

That is what you want to make a copy of.
And then you make a copy you customize that copy and then you use that copy.

So this requires something called deep copy support the copy support is basically this question of if I copy a struct what happens to the pointers because then you'll have several pointers pointing to just one object.

So when copying pointers for example you have to be really careful about the way you copy them so you actually replicate the elements they point to and not just copying the pointers over because that's not productive.

So this process which is called the deep copying or clothing we sometimes can make it even more convenient so we can make a prototype factory a kind of factory which serves prototypes and maybe even lets you customize the prototype at the point of creation.

So the prototype design factory is this idea that you have a partially or fully initialized object that you copy or clone and then you make use of that copy.

## Deep copying

In order to be able to implement the prototype design factory you basically have to do one thing you have to be able to
perform something called a deep copy of an object.

So what is this all about and why do we care.
Well you're about to see.
So let's just build a very simple scenario.

Let's define a person struct and let's suppose that the person is a name and the person has an address.

And we're going to have an address pointer here like so.
And we're going to define the address struct up here.

So type address struct it's going to be composed of a street address and city and country.

And those are just going to be simple strings.

So with this scenario what you might want to do is you might want to make a bus and and then if another person let's say they live nearby or they live at the same address but they have a different name you might want to perform a copy but unfortunately this copying process isn't as obvious as one might thinkit is.

So for example let's define John.
So John is going to be a person.
So the name is John and then we'll define the address here.
So the address is going to be 123 London Road in London U.K..
Okay.

So John can obviously be copied and we can make a copy of John as follows Jane is equal to John.

So what do you expect happens here.
Well what really happens is the name gets copied.
So there is no problem with that.
We copy the name from Jane to John.
It's fine.

But also the address pointer gets copied and this is a problem because of what you might want to do is you might want to customize Jane because that's the whole point after all you want to give Jane a name.

So you say Jane that name equals Jane.
This is completely OK.
There's no problem.

You now have John with the name John and you have Jane with the name Jane but then you want to customize the address as well so you say something like Jane dot address dot street address equals three the one Baker Street.

So that might seem correct to you but let's actually print both of these.

So let's print John and let's print Jane like so and we also need I'm using f empty the print here.

We also probably want to print the address separately because otherwise we'll just get a point of value so I'll see John Doe address and Jane Doe the address here.

Just so you can see the actual contents of the address.
So let's actually run this let's take a look at what happens here.

OK.
So we have an obvious problem.
So John and Jane they have different names so that's okay.
But for some reason they have the same address.
They are both now living on three to one Baker Street.
That's not what we wanted.

That's not what we wanted when we customised Jane but unfortunately that's what you get and the reason why you get it is because address is of course a pointer and when you copied John to the variable Jane using the assignment you copied a pointer as opposed to made a new object a new address and copied the contents.

So that is the problem of the copying and of course you have to you have to handle this.

So the question is if we cannot do this how can we reliably copy John into a new variable Jane Okay so let's comment all of this out and let's see how we can do this.

Well we'll still want to have Jane equals John as the starting point.

So the starting point will still be this statement but obviously it is insufficient because the address gets copied as a pointer and we want to recreate the address as well.

So Jane Doe the address would have to be a newly constructed address where the contents of that address are taking from John.

So here you would say John dot address dot street address.
John dot address dot City and John Doe address dot country.
Now we go.
So this is how you would handle the situation.
And now if we do this once again.

So remember we do want to customize Jane actually we just don't want to do it here.

So if we customize Jane now and we run this whole scenario now you're going to see the correct output.

Now you'll see that John and Jane do in fact have different addresses because this object right here is completely different to the one that's used up here.

And it's not just a process of copying the pointer.
You can see their different values here as well.
So this is how you perform something called deep copying.

So what is deep copy deep copying is basically this idea that when you're copying the object you're copying you making copies of everything it refers to including all the pointers the slices all the rest of it.

Because obviously if you don't do this then any object which operates as if it were a pointer like a slice for example would be shared between the original object and the copy and so modifying either of those would affect the other one as well which is obviously not what you want.

However the problem with this approach is that it really doesn't scale if you have a person which is composed of address and address is itself composed of something else then you end up with this very complicated recursive structure and you end up having to write lots of code just to be able to copy objects.

So you probably want to organize this somehow and not do all of this work and that's what we're going to take a look at next.

### Deep Copying code: creational.prototype.deepcopy.go

```go
package prototype

import "fmt"

type Address struct {
  StreetAddress, City, Country string
}

type Person struct {
  Name string
  Address *Address
}

func main() {
  john := Person{"John",
    &Address{"123 London Rd", "London", "UK"}}

  //jane := john

  // shallow copy
  //jane.Name = "Jane" // ok

  //jane.Address.StreetAddress = "321 Baker St"

  //fmt.Println(john.Name, john.Address)
  //fmt.Println(jane.Name, jane. Address)

  // what you really want
  jane := john
  jane.Address = &Address{
    john.Address.StreetAddress,
    john.Address.City,
    john.Address.Country  }

  jane.Name = "Jane" // ok

  jane.Address.StreetAddress = "321 Baker St"

  fmt.Println(john.Name, john.Address)
  fmt.Println(jane.Name, jane. Address)
}
```

## Copy Method

So now that I've shown you the problem with DB copying you might want to somehow organize your code so the deep copying is easier.

Now how can you do this.

Well one very simple approach is that you take every single struct that you have in your model and you give this truck to a method called deep copy which explicitly performs a copy.

Here's what I mean like for example here we have an address so I can have a method which is on an address pointer called DB copy which returns an address pointer.

And this is precisely the location where you would basically make a new address like we're doing down here so you would make this address and you would return this address from the method obviously instead of John we'll just put the variable a which refers to the receiver and we don't need any of this stuff so let's just let's get rid of let's get rid of address in all of these cases.

So get rid of it here and get rid of the extra dots obviously.

There we go.
Okay.

So essentially we perform deep copy on the address and we also have to perform deep copy on person.

Let's actually add something else just to demonstrate that this is also necessary for other types.

Let's suppose that we have a list of friends.
So friends are going to be a list of strings.
Now it could be a list of people instead.
But I would just keep it simple.

Let's have it as a as a slice of as a slice of strings basically so now we have to perform deep copy on that as well so we would also define a deep copy method on person so func P plus and deep copy which returns person pointer.

So here you would perform the copy so you would say Q is equal to P so that defines a copy of everything that can be copied by value and then you have to copy the address.

So you say Q Dot address equals Pitot address dot deep copy so because it has the DB copy method we can use that and when it comes to the slice you just say copy to cued up friends from pedo friends so you copy of the list of friends and then you return and Q here.

So this is actually going to work at least in our scenario.

So let me get rid of some of this code and I'll show you how this works so first of all we added this idea of a person having a list of friends so.

So here I'll just have a slice with let's say Chris and Matt.

So they are John's friends and then we can take Jane and we can say Jane dot Well we say Jane is equal to John Doe leave copy.

So we perform the deep copy where every single collection of elements every single point here actually gets unwrapped and gets copied correctly.

So now we can customize Jane once again we can say Jane Doe named for example equals Jane we can customize Jane's street dress so we can say Jane Doe address dot street address equals three two or three to one Baker Street we can also customize Jane's list of friends we can say Jane Doe friends and we can say append chained up friends comma Angela for example and then we can output all of these so we can F.A. the print line both John and let's not forget John Doe address because otherwise you don't get the output and then we can do the same thing for Jane.

So let's have Jane here.
John Doe address and Jane Doe address that's actually run this.

And let's take a look at what's going on well obviously we have some error here.

So I'm trying to return AMP Q And it's saying that you cannot return Amcu as type star person in return argument it's slightly weird.

I wonder.

Oh of course yes there is should be a star here as we are making a copy and not of a pointer but of the actual thing.

So let's try this once again and this time around we are getting the output let me actually I'll just do a print line instead of print so you get to see it all in one line.

So here it is on two lines.

So we have John who is friends with Chris and Matt lives at 1 2 THREE London Road and we have Jane 0 is friends with Chris Matt and also Angela who is at 3 to 1 Baker Street.

OK.

So what is the takeaway from this example whether the takeaway is that you can organize your own objects to have some sort of deep copy method available on them and then this deep copy method can actually be invoked on any of your objects to perform the deep copying however it still leaves open the problem of what to do with times which you don't own like for example here you have a slice you cannot just go ahead and you know add additional behaviors to that slice.

So essentially you are stuck with having to call some sort of copy method or something to that effect or even if you could add the behaviors to the slice.

How would that change things.

It would still force you to basically double and triple check every single one of your struct and make sure that every single one of the members types has a deep copy method.

So it's a workable solution.

It works ok but it's not ideal because you have to still do a lot of work in order to get the copying done.

But you know this is one possible approach.

### Copy Method code: creational.prototype.copymethod.go

```go
package prototype

import "fmt"

type Address struct {
  StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
  return &Address{
    a.StreetAddress,
    a.City,
    a.Country }
}

type Person struct {
  Name string
  Address *Address
  Friends []string
}

func (p *Person) DeepCopy() *Person {
  q := *p // copies Name
  q.Address = p.Address.DeepCopy()
  copy(q.Friends, p.Friends)
  return &q
}

func main() {
  john := Person{"John",
    &Address{"123 London Rd", "London", "UK"},
    []string{"Chris", "Matt"}}

  jane := john.DeepCopy()
  jane.Name = "Jane"
  jane.Address.StreetAddress = "321 Baker St"
  jane.Friends = append(jane.Friends, "Angela")

  fmt.Println(john, john.Address)
  fmt.Println(jane, jane.Address)
}
```

## Copy Through Serialization

So in the previous lessons we looked at this idea of basically defining a deep copying method on every single struct that you have in your program or every single struct that you need to replicate.

And then invoking that in a kind of recursive fashion we also talked about the fact that it doesn't save you from handling all those structures which you do not have a deep copy method like slices for example.

And the fact that you would have to perform your own copying operations relevant to that particular data structure so you might have a question well is there some magic wand that we can just wave and have all of it go away and have a good replication of objects where you just say let's copy this struct and the struct is copied including all of its dependencies including the following of the pointers and all the rest of it.

And the answer is luckily yes it is possible to replicate a struct completely including all of these dependencies.
And this is done using the serialization.
So let's go into our imports here and we've going to add support for bytes and also support for encoding slash gob.
So that is binary serialization.

We're gonna be using binary serialization although of course you're welcome to serialize to something else like Jason for example.

That's up to you but I'll use binary serialization.
So what is the idea.
Well the idea is that serialization constructs are very smart.

If you give them something like in person the serialize there is going to figure out that you have a string which has to be safe there's a string obviously but that here that you have a pointer and you need to follow that pointed to take a look at the data that's actually in there and serialize that data and not just the point of value because you realizing a pointer value which is an address in memory is just meaningless it doesn't help anyone.

So this you realize they know how to unwrap a structure and serialize all of its members.

And so if you think about it if you serialize the person to let's say a file or just to memory you save all of it state including all the dependencies and when you do serialize it you construct a brand new object initialized with those same value.

So effectively you can have a deep copy performed automatically for you and so we can get rid of all of these methods because we no longer need them we'll have a different kind of deep copy method but we'll have a deep copying method which actually leverages this idea of using encoders and decoders to basically save all the data and then read from the data.

So let's have a utility now deep copy method on person.
So this is gonna be deep copy on person.
So how do we do this.

Well first of all make a buffer so B is going to be a bytes a buffer and then we make an encoder.

So E is going to be golf dot new and coder it's gonna be an encoder which takes a pointer to this buffer that we made and now what we can do is we can take E and we can encode whatever is actually whatever needs to be written to the buffer.

In this case it's the person.
So we take the bus and we ride the bus into the buffer.

Now if you're interested in what actually got put put into the buffer we can take the buffer wicking take the bytes turn them into a string and actually take a look.

So let's do that.

It's an interesting exercise so I will print line string of beat up bytes so that way you'll see some string representation that might be some buggy characters in there because not all of them are principal but you'll see something.

Okay.

And then what we do is once we've got our buffer and we've got the data in a buffer we use a decoder to read the data from that buffer into a new object.

But first of all let's make the decoder So Gob dot new decoder.

Once again let's have appointed to the buffer and then here's the result.

So the result is a fuss and object.

So I prepare the memory for the person object and then I use that memory to use the decoder and decode into that object.

So I give the API a pointer to the result thereby sort of initializing that memory with the data that's read from our buffer.

And then I can return that result like so.

So this is our deep copy implementation let's actually take a look at how this whole thing can work now.

So the way it works is exactly the same as before.

So we don't need to change anything here because notice I've also I called the method deep copy once again.

So you you call John the deep copy but now you don't have to have this recursive call so the address doesn't need a deep copy method unless you have an address that you need to save somewhere.

You you basically have just the same invocation customize the name the customize the address.

We change a list of friends and if I now run this well hopefully we see predictable results and that's exactly what we're getting here so you can see John and Jane and you know they have different friends list and they live a different dresses so it looks like we finally solve this tedious problem.

All for copying objects and then of course there is the small matter of discussing the prototype design pattern because the prototype design pattern is all about taking a pre configured object like John here making a copy like we do here in deep copy and then customizing it like we do here.

So that's all there is to the design pattern this approach works rather well and it keeps the number of code that you have to write down to a minimum.

### Copy Through Serialization code: creational.prototype.serialization.go

```go
package prototype

import (
  "bytes"
  "encoding/gob"
  "fmt"
)

type Address struct {
  StreetAddress, City, Country string
}

type Person struct {
  Name string
  Address *Address
  Friends []string
}

func (p *Person) DeepCopy() *Person {
  // note: no error handling below
  b := bytes.Buffer{}
  e := gob.NewEncoder(&b)
  _ = e.Encode(p)

  // peek into structure
  fmt.Println(string(b.Bytes()))

  d := gob.NewDecoder(&b)
  result := Person{}
  _ = d.Decode(&result)
  return &result
}

func main() {
  john := Person{"John",
    &Address{"123 London Rd", "London", "UK"},
    []string{"Chris", "Matt", "Sam"}}

  jane := john.DeepCopy()
  jane.Name = "Jane"
  jane.Address.StreetAddress = "321 Baker St"
  jane.Friends = append(jane.Friends, "Jill")

  fmt.Println(john, john.Address)
  fmt.Println(jane, jane.Address)
}
```

## Prototype Factory

All right so now that we've figured out how to get object copying to work the way that you expect it to work how about an example where we set up a system where it's easier to actually use these prototypes.

So we're going to build a scenario that's very simple to the person an address scenario.

This time around we'll have employee an address.

And the idea is that when you have employees they might work in different offices of one company.

So so they might share like a city for example or they might not.

And you want to quickly sort of customize these objects because essentially the problem is that you still have too much customization being done by hand and it would not be nice to take this customization and put this into some sort of a set of functions for example.

And this is what we typically call a prototype factory.

Now we already looked at prototype factories when we talked about factories in the previous section of the course.

But now we're going to take a look at a slightly different implementation and implementation where you.

### Prototype Factory code: creational.prototype.factory.go

```go
package main

import (
  "bytes"
  "encoding/gob"
  "fmt"
)

type Address struct {
  Suite int
  StreetAddress, City string
}

type Employee struct {
  Name string
  Office Address
}

func (p *Employee) DeepCopy() *Employee {
  // note: no error handling below
  b := bytes.Buffer{}
  e := gob.NewEncoder(&b)
  _ = e.Encode(p)

  // peek into structure
  //fmt.Println(string(b.Bytes()))

  d := gob.NewDecoder(&b)
  result := Employee{}
  _ = d.Decode(&result)
  return &result
}

// employee factory
// either a struct or some functions
var mainOffice = Employee {
  "", Address{0, "123 East Dr", "London"}}
var auxOffice = Employee {
  "", Address{0, "66 West Dr", "London"}}

// utility method for configuring emp
//   ↓ lowercase
func newEmployee(proto *Employee,
  name string, suite int) *Employee {
  result := proto.DeepCopy()
  result.Name = name
  result.Office.Suite = suite
  return result
}

func NewMainOfficeEmployee(
  name string, suite int) *Employee {
    return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(
  name string, suite int) *Employee {
    return newEmployee(&auxOffice, name, suite)
}

func main() {
  // most people work in one of two offices

  //john := Employee{"John",
  //  Address{100, "123 East Dr", "London"}}
  //
  //jane := john.DeepCopy()
  //jane.Name = "Jane"
  //jane.Office.Suite = 200
  //jane.Office.StreetAddress = "66 West Dr"

  john := NewMainOfficeEmployee("John", 100)
  jane := NewAuxOfficeEmployee("Jane", 200)

  fmt.Println(john)
  fmt.Println(jane)
}
```

## Summary

All right so let's try to summarize what we've learned about the prototype design pattern so to implement a prototype you partially construct an object and you store that object somewhere and then you implement deep copying however you do this you allow for deep copying of the prototype and then once you have a copy of the instance you customize the resulting instance and you can also make prototype factories which provide convenient API for actually using the prototypes.
