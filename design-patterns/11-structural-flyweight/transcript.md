# Flyweight -- Space optimization

## Overview

When I'm going to talk about the Flyaway design pattern so flyweight is essentially a space optimization technique.

The idea is to avoid any kind of redundancy when storing data.
So here's an example.

If you think about any kind of massively multiplayer online game not necessarily an RPG any kind of multiplayer online game you're going to have lots of uses.

And if you ask the users to enter their first and last names you're gonna to have lots of users with identical first and last names.

You're gonna have lots of people called John.
Lots of people with the last name Smith.

That's just that's just a statistical expectation of the set of people's names.

There is really no point in storing the same textual data like the same first and last names over and over and over again that is wasteful and we know that we can avoid this.

Now what you can do instead is you can store just a list of names externally and then for every single player they can refer to those names like using indices or pointers or whatever mechanism you prefer.

Another example is like if you have a text editor for example you are formatting text and you want to make some chunks of text make some parts of text like bold or italic or something.

There are different ways of going about this because what you can do is you can take every single character and also give it formatting information like split the text into individual letters and each letter will have a boolean flag indicating if for both the boolean flag for italic and so on.

But this is once again wasteful so what you really want to be doing is to be operating on ranges so you would specify the starting and ending positions of the piece of text and then you can customize that piece by saying this is bold this is italic and so on and so forth.

So the fly would design pattern is a space optimization technique that allows us to use less memory by storing the data associated with similar objects externally.

And then you can customize the data to actually well for example to provide text formatting as we'll see in the demos.

## Text Formatting

In order to demonstrate the flyway design pattern we're going to take a look at a very simple example where let's imagine that we have some sort of text editing application.

So you're working with plain text but you want to add a bit of formatting maybe you want to make text bold or italic or something like that.

And the question is how do you actually implement this in code.

So what we're going to take a look at is we're going to take a look at a very inefficient version of doing that and then we'll take a look at a more efficient version which uses the flyway design pattern.

So let's imagine that we have some time called formatted text.

Now at the moment the only kind of formatting I'm going to have is capitalization I'm basically going to specify that certain regions of text have to be made using all capital letters.

So we'll have the text itself plain text as a string and in addition we'll have capitalization.

So the question is how do you indicate that a particular chunk of this plain text has to be made capitalized.

So one approach one very naive approach is to make a boolean array.

You basically make a boolean slice gold capitalized.

So this boolean array the underlying boolean array is going to be of the same size as the number of letters in plain text.

So we're going to basically allocate a boolean value for every single letter indicating whether or not this letter has to be capitalized.

And surprise surprise this is actually going to work it.

Let me let me first of all make a constructor so here.

When it comes to making a new formatted text what I'm going to do is I'll specify the second parameter.

So the first parameter is obviously just the plain text.

But the second one is going to be a call to make where it will make a boolean array and the size of that the rate is going to be the length of plaintext.

So that is pretty much all that we all do we need.
That's that's all there is to it.
Let's get rid of plain text here and we are done.

We are done with initializing full method text but of course the real issue is how do you render it.

So you've got this plain text and you want to render it to the screen including all the capitalization rules being applied.

So let's implement the stringer interface for formatted text so it's going to look like this.

So what I'm going to do is I'll make a string builder and then I'll go through every single letter I'll check whether or not this letter has to be capitalized.

I will capitalize it if necessary and I will added to the string builder and return the the final result.

So here is the string builder so the string style builder that will go.

So I'm going through every letter I starting from zero.

I is less than FDR plain text length I plus plus like so.

So I'm gonna get the letter a letter C character C so it's F plaintext add position I.

And then I have to find out whether or not this letter has to be capitalized.

So if I do want to capitalize it there will be a boolean value inside my array so if that's the case then what I'm going to do is I'll write a rune where I'll simply do Unicode to opera.

So I'll take it to uppercase and then I'll just cast of the cast that point to a room so Room C like so otherwise I'll just write it as is.

So I'll say as be the right rune Room C like this and then I'll return the entire thing so I'll return as B string thereby putting all the parts together and returning them as a single string and this is going to work.

This is something that is going to work just fine.

So let's have a bunch of text here so the text is going to read.

This is a brave new world and let's suppose we want to capitalize the word brave.

So here I'll make formatted text so let's use that construct and you formatted text specifying the text I say FTD capitalize.

And by the way we need a utility method for capitalizing things we need the utility method that would go into this array and actually set those boolean flags so let's add it here so func f formatted text.

Let's go to capitalize so we'll define the start and end position.
The positions that you want to make capital.

So here we'll just go in a loop for I from start I less than or equal to N so it's going to be an inclusive range.

And here I'll say after capitalize at position i equals true.

So this is the utility method that we can use to capitalize individual letters.

So here I can say F T's dot capitalized with the capital C of course it's a method not the not the internal array that we are storing and I can say starting from the tenth letter to the 15th letter let's just capitalize that and then let's print line the whole thing.

So FTD dot string and let's run this and let's take a look that's what we get here.

So as you can see we do have the right operation.
This is a brave new world with the word brave being capitalized.

Unfortunately this approach is extremely inefficient it's inefficient because we are specifying a huge boolean array one element for every single character inside plaintext.

The problem with this is that if imagine you're reading an ed text like a war in peace and inside war and peace you only want to capitalize a single word outside of thousands upon thousands of words you're going to be allocating lots and lots of values lots of billion values that you don't even need.

So there must be a better way of capitalizing or indeed applying any kind of formatting in a text editor scenario.

And in fact there is what we can do is we can introduce an idea of a text range which is simply the starting and ending positions of a range inside this rather large set of letters that we have.

So let's take a look at how this will be implemented.

So you would have some sort of type called text text range text range struct.

You'll have the start and end positions which are integers and you'll also have any kind of formatting information that you need so for example you can have information about capitalizing letters you can have bold italic any kind of formatting you want.

And these are going to be boolean flags.

So essentially if you set it to true you want this range of letters to be bold or italic or whatever.

Now we can also have a utility method for figuring out whether it's X range covers a particular point.

So here func t text range covers some position which is an integer and it returns a boolean so we're basically checking that the position is greater than the start of the range or and less than the end of the range including the end points.

So a return position greater than our equals Tito start and start with a capital S and position less than or equal to C dot.

And there we go.

So this way we're checking that the range is actually covering a particular position and we can use this as subsequently when we make a different implementation of our formatted text struct.

So we'll have a better formatted text struct.

So here as before we'll have plaintext which is a string but instead of having this huge boolean array and imagine if you had if you had to capitalize and make both and make italic you'd have several of these huge boolean arrays but we're not going to have that.

Instead we'll have formatting which is going to be an array of text range pointers.

Now the reason why it's pointers is because you also want to share these text ranges you want to be able to return them to the use it to operate upon.

So what you can do is you can return the client a pointer to a text range and they can manipulate that text range which is really wonderful we're going to do this in just a moment but first of all let's let's do a few utility things so let's have a constructor as before so we'll have a constructor which initialize this plaintext and here that's really all that you need to do in the sense that you don't need to perform any additional operations apart from apart from doing this actually the specify here should remain because I'm not I'm not providing the formatting information right here.

So we said the plaintext that's pretty much it.

Now what we want to be able to do is we want to be able to construct and return a range inside this text.

So we'll have a function so function on a better formatted text called range.

So you're going to be able to construct and return a range object given the starting and ending positions like so so we're going to define a range as a text range with start and end all of the other flags are basically going to be false to begin with so false false false.

You don't have to specify them right here.

So we do two things we add this range to the set of formatting on the text and we also return this range as well so here I say BDO formatting and I just append this range to the formatting and I also return the range pointer so that the user can operate upon it and this allows the user to take the range and customize it to their heart's content.

They can even store it somewhere and customize it later on.

OK so would this set up the last thing we need to do is to implement the stringer interface on better formatted text so that we actually have a string method to work on.

OK.

So what's happening here while this situation is going to be somewhat similar to the situation in the in the previous example except that we're going to be searching through this set of formatting flags that we have.

So here's how it goes.
First of all make the string builder as before.
Strings builder.

There we go and then we'll go through every single character inside plaintext.

So once again for I starting in 0 5 less than the length of the plaintext I plus plus.

OK.
So we get the character.
So C is beta plaintext at position i.

Now what I'm going to do here is I'm going to take lots of liberties with conversions from Unicode to ordinary bytes for example in the real world.

This would be a slightly more complicated.
But I just want to skip over all the complicated parts.

So essentially now that we have this character we need to go through every single formatting specifying every single element inside this array and we need to find out whether or not that particular text range actually covers this character.

So we go through every single one.
So for underscore a comma are in the range of B dot formatting in this range.

So we say if r covers the position I and R specifies that somebody wants to capitalize a letter then we capitalized that letter we say C equals and then there's gonna be lots of manipulations here because

I'm going to call you in 8 on Unicode not too upper on rune on C.
So lots of character related stuff that's not critical to our discussion right now.
But basically I'm making this character uppercase.
That's all I'm doing.

And then once we're done with this inner loop here I can as be that right rune once again convert C to a rune which is going to be fine in our case because we're using mainly just ASCII characters and we're going to return as B adult string which is going to return the entire thing.

So what have we done here with this event essentially with constructed a flyweight so the text range that we've defined here is a flyweight.

It's essentially an object that allows manipulation and a certain scale but it's very compact it tries to save memory.

So instead of making huge boolean arrays indicating the formatting of every single character inside a piece of text we have this very compact representation we have just the starting and ending points and then some customization like making a text capitalized for example.

So let's take a look at how we can use all of this.

So here I'll make better formatted text DFT knew better formatted text once again specifying that string of this is a brave new world what I can do now is I can say FTD Dot and I can grab a range of characters.

Now let's suppose we want to capitalize the word new so the range will go from the 16th to the 19th character and I can say capitalized equals true.

OK.
So so this probably requires an explanation like what's going on here.

So when you call range do things happen at the same time so arrange is constructed it gets added to the set of ranges inside the formatting field of the better formatted text but it also gets returned to the client.

So we when we call range we actually get a pointer to the text range object.

So this thing gives us a pointer and on that pointer we can performing amputations like for example we can capitalize the whole thing.

So now if I do a print line and I do BFD to string or just string like so we're going to get hopefully the right result let's actually run this end as you can see that's exactly what we're getting so we've capitalized the word new.

Now the critical distinction between this approach and this approach is of course the savings in memory and this is what the flyweight pattern is actually for is essentially a trick to try to avoid using too much memory and to instead maybe introduce additional sort of temporary objects like we're doing with text ranges but these temporary objects allow us to save a lot of memory and that's always a good thing.

### Text Formatting code: structural.flyweight.textformatting.go

```go
package main

import (
  "fmt"
  "strings"
  "unicode"
)

type FormattedText struct {
  plainText  string
  capitalize []bool
}

func (f *FormattedText) String() string {
  sb := strings.Builder{}
  for i := 0; i < len(f.plainText); i++ {
    c := f.plainText[i]
    if f.capitalize[i] {
      sb.WriteRune(unicode.ToUpper(rune(c)))
    } else {
      sb.WriteRune(rune(c))
    }
  }
  return sb.String()
}

func NewFormattedText(plainText string) *FormattedText {
  return &FormattedText{plainText,
    make([]bool, len(plainText))}
}

func (f *FormattedText) Capitalize(start, end int) {
  for i := start; i <= end; i++ {
    f.capitalize[i] = true
  }
}

type TextRange struct {
  Start, End int
  Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
  return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
  plainText string
  formatting []*TextRange
}

func (b *BetterFormattedText) String() string {
  sb := strings.Builder{}

  for i := 0; i < len(b.plainText); i++ {
    c := b.plainText[i]
    for _, r := range b.formatting {
      if r.Covers(i) && r.Capitalize {
        c = uint8(unicode.ToUpper(rune(c)))
      }
    }
    sb.WriteRune(rune(c))
  }

  return sb.String()
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
  return &BetterFormattedText{plainText: plainText}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
  r := &TextRange{start, end, false, false, false}
  b.formatting = append(b.formatting, r)
  return r
}



func main() {
  text := "This is a brave new world"

  ft := NewFormattedText(text)
  ft.Capitalize(10, 15) // brave
  fmt.Println(ft.String())

  bft := NewBetterFormattedText(text)
  bft.Range(16, 19).Capitalize = true // new
  fmt.Println(bft.String())
}
```

## User Names

We're now going to take a look at what is probably the most classic implementation of the flyway design pattern the storage of people's names.

So why is it such a big deal.
Well let me show you.

Let's suppose we have some massively multiplayer online game which has a bunch of users in it.

So the user gets to type in their name and well that's their name.

Maybe they're required to type their proper full name like first name and last name for example.

So we'll have full name as a string here.
So what can we do.

We can have a constructor which actually sort of factory function which initialize as the.

Use it to their full name and then we can start making these uses down here somewhere.

So for example I can have John John a new user John Doe.

I can also have Jane have Jane.
So there's gonna be Jane Doe and.

Well let's have another Jane that's have also Jane whose name will be Jane Smith for example.

Okay.
So what is the issue with all of this.
Well there is no problem if you have just three users.
If you have three users there is really no issue.

But in the in the real world out there if you take let's say 100000 thousand different players in an online game or any online system for that matter you're going to have people with similar names you're gonna have lots of people called John and you're going to have lots of people with the last name Smith.

And every time that you store one of these you're effectively duplicating memory.

So you're effectively duplicating things because every time that you have like say Simon Smith or something that's a new string that's a new string that needs a new chunk of memory and in a very large user setting where you have lots and lots of people it is simply impractical to store all of these objects considering that they are very very similar.

Why do we want to waste so many bytes on the storage of all this data why can't we somehow compact ify this data like store it inside a single container like store all the first names and last names inside arrays and then just have indices into those arrays.

So that is the another approach taken by the flyway pattern that can save you lots of memory.

So let's take a look at how you would implement this.

So as you can probably guess the implementation is going to be more difficult.

You don't just add a string as a member and say oh we're done here.

It's gonna be a bit more difficult.

So what I would do here is I'll make a variable called all names which is going to be a String array no strings Liszt and then I'll have a new type let's call it user tune and use it to is going to be a bit more frugal when it comes to the use of memory because instead of storing the actual strings the names and notice I'm saying names here we have just a single full name here we have names as in the several different parts of the name.

The names are going to be stored as simply you end eight values.

So there's gonna be a bunch of these values like John and Jane all of these will become integers.

So what's going to happen is we'll take the full name of the user.
We'll split it up using space as the separator.
So we'll get separate strings like John and doe.

We'll put all of those strings inside this global array and then we'll have indices into this array being done from here.

So that means that the constructor for User 2 is going to be rather complicated.

OK.
So let's do that.
So.

So we'll have func a new user to where you specify the full name and you get a use it to as the end result.

Okay.
So so this is where things will get rather tricky.

So what we want to be able to do first of all is we want to have some sort of function which either gets an element from all names or if that element doesn't exist.

It adds that element and returns you the index.
So we're using U.N. aid here just to save some memory.
So we'll have a function called Get or add.

So this is going to be a function which takes a string s and it says a you end eight for that element.

So what do we need to do.

Well first of all since this is just a slice it's not like a dictionary or anything.

What we're going to do is we're going to go through every single element inside that slice and if we find a match for this particular string we'll just return the index of the position where that match was found.

So I'll say for i in the range of all names if all names at position high is equal to S then we return the index I cast to a U.N. date once again here I'm using U.N. dates I'm making the assumption that there's only gonna be 256 unique names at most and that is done to additionally safe memory because well if you make it an end then all of a sudden every single one of these names elements because becomes eight bytes 64 bits.

So I'm doing you an eight here.
So the idea is if we found it then we just returned the index.
Everything is okay.
But if we're here we did not find it.

So all we do here is we append so we append the string s and then we return the length of the whole thing minus one because we want an index of the last element that would be the length of the whole array minus one.

So we return you indicate of all names grab their length minus one.

So that is an inner function called Get or add what we now need to do is we need to make a use it to we need to split the full name that was provided as the argument.

Right here we need to split that full name into the separate parts so the parts are going to be strings that split.

So we take full name and we split it on a space separator and then for every single one of these parts we call get or add on that constituent parts and we put the element into names.

So we put the index into names so let's do that.

So for this call my p in a range of parts result up names equals.

Well it's just to like this append get or add P.

So in this case P is every single one of the parts so for example somebody specified John Doe so 1 on the first call P would be John on the second call P would be DOE and every single one of these would be fed into this.

Get R add function.

This function will return the U.N. date that will subsequently append the set of names.

And then finally we'll return the overall result.
So this is how you construct a new user 2.

Now of course now we have a problem here if you wanted to get the full name of user you just acquired full name.

That's pretty much it.

But in our case there is no full name as a single element instead of we have a bunch of integers inside a member called names.

So whenever somebody needs a full name we have to basically provide a function which reconstitute those.

So let's have a method on user 2.
So methadone use it to called full name.
It's going to return a String.

So here we'll have an array of parts then we'll go through every single element of the names and try to acquire that part.

So for on this score comma Ivy in range of you names.

What we do is we save parts.
We append the part from all names though.
So we take the index.

The idea would feed it to all names and we get the actual string and then we append that string to the string.

Supplies that we have here and then we use string that join on the sly so we return.

Strings don't join.

We provide the path slice and we join them using a space as the separator.

Okay so let's take a look at how fast all these work and how the other elements work.

So just to kind of illustrate a print line.
So let me print line for example.
So.

So here in the case of John print line John Doe full name and you get John's full name.

There's really no problem of course I'm using variables which haven't been which are not being used actually.

So here is John Doe the full name.

Getting the full name is easy but when it comes to using User 2 then all of a sudden it becomes a bit more complicated.

So John 2 is going to be a new user to call John Doe John Doe.

And then of course if you want to print line his name you have to say John to full name with round brackets because now it's a method it's not just a field.

So I run this and I get pretty much the same output now as you can imagine a bit more work has been carried out than before.

But on the other hand there are certain memory savings so how much memory are we actually saving.

Well let me actually drop in a kind of initialized demo here and we'll talk about what's going on here

I'll get rid of some of the extra print lines right here.
Okay.

So in the first example I'm making three uses I have John Doe Jane Doe and Jane Smith so you can see some of their names overlap like Jane here and Jane here and Doe here and Doe here.

So we're calculating the memory taking by these users as basically taking their full names turning them into binaries and looking at the length of those binaries.

So that is going to give us the number of bytes taken up by John and Jane and also Jane.

Okay.

In the second case we're doing exactly the same thing but we're using the user 2 as opposed to ordinary users.

So we're making John Doe Jane Doe and Jane Smith.

But here are the memory calculation is different because first of all we have to take all the names inside the system and we have to calculate their total length.

So we take every single element we take the length and bytes we added to the total memory and then of course we also add the length of every single one of those bytes race.

Now if these were let's say integer arrays you would have to multiply this by some unpleasant value like four or eight for example if you defined it as an int but here we're using U.N. date.

So we're not we're multiplying them by 1 effectively.

So we're adding the length of all of these three elements.

Now we can run this and we can see that comparatively how much memory stick.

So you can see that in the second case we have a saving of 4 bytes.

Now as you can imagine if you had a really large scenario really thousands upon thousands of users you would save huge amounts of memory.

It wouldn't just be a couple of bytes it would be a couple of thousand maybe a couple of million bytes because essentially you are storing bytes byte arrays of just just typically two bytes per user.

So if we if we kind of ignore all the old names here you're storing 2 bytes per user for everyone else and then the the combined all names array would be a lot smaller because of all these repetitions between the first names and last names and so on and then of course some users have a unique first names and some unique last names and that's not a problem either because the system handles it just fine.

But this is yet another illustration of how the flyway design pattern allows you to save memory and to just sort of to it to optimize the system not one question you might have is where exactly is the flyweight.

So here are the flyweight is inside names.

So names just like in the previous example where we had text formatting and went had that text range construct.

These are also kind of like ranges they're kind of like pointers into all names here except they they are indices into all names.

So instead of operating on strings directly you operate on integers which are representations or are pointers into this overall array.

### User Names code: structural.flyweight.usernames.go

```go
package main

import (
  "fmt"
  "strings"
)

type User struct {
  FullName string
}

func NewUser(fullName string) *User {
  return &User{FullName: fullName}
}

var allNames []string
type User2 struct {
  names []uint8
}

func NewUser2(fullName string) *User2 {
  getOrAdd := func(s string) uint8 {
    for i := range allNames {
      if allNames[i] == s {
        return uint8(i)
      }
    }
    allNames = append(allNames, s)
    return uint8(len(allNames) - 1)
  }

  result := User2{}
  parts := strings.Split(fullName, " ")
  for _, p := range parts {
    result.names = append(result.names, getOrAdd(p))
  }
  return &result
}

func (u *User2) FullName() string {
  var parts []string
  for _, id := range u.names {
    parts = append(parts, allNames[id])
  }
  return strings.Join(parts, " ")
}

func main() {
  john := NewUser("John Doe")
  jane := NewUser("Jane Doe")
  alsoJane := NewUser("Jane Smith")
  fmt.Println(john.FullName)
  fmt.Println(jane.FullName)
  fmt.Println("Memory taken by users:",
    len([]byte(john.FullName)) +
      len([]byte(alsoJane.FullName)) +
      len([]byte(jane.FullName)))

  john2 := NewUser2("John Doe")
  jane2 := NewUser2("Jane Doe")
  alsoJane2 := NewUser2("Jane Smith")
  fmt.Println(john2.FullName())
  fmt.Println(jane2.FullName())
  totalMem := 0
  for _, a := range allNames {
    totalMem += len([]byte(a))
  }
  totalMem += len(john2.names)
  totalMem += len(jane2.names)
  totalMem += len(alsoJane2.names)
  fmt.Println("Memory taken by users2:", totalMem)
}
```

## Summary

All right.

Let's summarize what we've learned about the flyway design pattern.

So the idea is very simple.

You take all the data that can repeat all the sort of common data and your story externally so you don't store it at the position of this truck to that needs this data you store it somewhere else and then you basically specify some sort of index or pointer into the.

So no data store so that objects can actually share this data and can access this data where necessary.

You can also in the in the context self text editing you can define this idea of a range.

So you can define a range on any homogeneous collection and just store the data related to the range rather than storing data related to any individual elements.

This is another use of the flyway design pattern.
