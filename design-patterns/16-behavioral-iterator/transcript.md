# Iterator -- How traversal of data structures happens and who makes it happen

## Overview

The iterator design that is all about traversing data structures.

So what is the motivation for using the iterator well iteration or traversal is actually a core functionality of various data structures and and iterate Tor is a type.

So it can be a struct for example that facilitates the traversal.

So basically what the iterator does is it keeps a pointer to some particular element within a collection and then it has mechanisms of advancing so the iterator knows how to move from one element to the subsequent element and go allows very simple iterations using of the range keyword.

So you have seen those if you're using things like arrays or slices or any of that stuff you have built in support for iteration of those without any extra steps.

But sometimes we need to support that in our own structures as well.

And there are different ways of doing that.

So the iterator is quite simply an object that facilitates the traversal of some data structure.

## Iteration

Let's talk about the general idea of iteration of something so to iterate something means to go through every single element or maybe most elements or maybe a selection of elements.

So let me show you a very simple example let's suppose we have a type person and this person has a first name a middle name and a last name.

And they are all strings.

So the question is a what if somebody wants to iterate all the names.

What if somebody wants to go through every single name in this sequence.

How would you actually get this working.

So one of the approaches is that you can simply expose an array so you can take every single name you can put it inside an array and just return that array.

And because arrays are either a ball there is really no problem.

So we can have a function.
We can have a method.

In fact we can have a method called names which just returns three strings so we can return three strings three strings pedo first name pedo middle name and Peter last name.

So that's already something that we can work with in the sense that we can make a person we can give this person names and we can iterate those names.

So let's choose a famous American Inventor let's say Alexander Graham Bell.

There we go.
OK.

So what I can do is I can say for undiscovered common name in the range of PDA names and PDA names is an array so we can use the range keyword here like so.

And we can go through every single one and we can print it for example.

So I can print the name like so.

OK so let's run this let's take a look at what the output is.

As you can see the output here is correct.
So this is one approach.

Obviously this approach is not free because you're effectively copying over a string so it's not as efficient you might want to switch to pointers for example and there are also additional complications like let's suppose for example that if the middle name is empty you don't want to return the middle name as part of the names.

Let's suppose that it can be empty because some people don't have a middle name.

So how would you implement this.

Well in this particular paradigm it's difficult.

You would have to jump from a using an array to using a slice but it's still manageable it's still doable.

So this is one approach.
Now another approach is to use a generator.

So it's this whole business of using channels and go routines and all that stuff so we can have a function.

We can have a method once again.
So a methadone person called names.
Well it's code named generator.

So this is gonna be a generator which is going to give us a bunch of strings.

So here I'll make a channel.

So out is going to be a channel which yields strings.

And then let's make a go routine.

So go func like so and here I'll defer the CLO closing of all the output channel.

And here I will just return the first name.

And then what I can do if I don't want to return the middle name if it's empty I can check the length.

So I can say if the length of beat up middle name is greater than 0 then yeah we're going to return it.

So we say out paedo middle name and then out P..

Not last name.

So this is gonna be running in parallel and then we return the output channel like so.

So I can use this instead down here.

So what I can do is I can instead of saying range P names I can see a range P names generate.

Now we go and now of course I need to I need to make a small change because now it's going to be just going through the names.

There is no counter variable because it's not a it's not an array or a slice it's just a generator of strings so we consumed this generator we get every single name and it's actually run this once again.

So we get Alexander Graham Bell and if I decide to get rid of Graham here and run this once again and I just get Alexander Bell because we are ignoring that middle name that is while the check is being done right here.

So let me undo this for a moment so this is the second variety of iteration using a generator.

And then of course there is the third the most complicated variety of iteration.

That is when you use a separate struct.
Now this approach is a very un idiomatic.

It's it's the kind of approach they use in C++ but you can also use it in go if that's what you need.

So for example what I can do is I can make a new type called Person name iterator.

So this is going to be a struct that's going to have a pointed to the person that's being iterated plus in person and in addition it will have a value and integer value in creating the current element that we're supposed to be giving out as as you are iterating through this whole thing.

So here let me make a factory function which just initialize as the person and we'll also initialize the current value to minus one.

The idea being that as you start iterating you move the current value to zero so that would be first name then you move it to 1.

Then you move it to 2 and you cannot move it beyond 2.

So then we can have some sort of a function let's call it move next.

So this is gonna be a function on a person named iterator called move next.

Now this moves the iterator forward and it returns a boolean indicating whether there is actually something to consume because remember we don't have an infinite number of names we only have three names.

So here we say Peter occurrence plus plus.

So we move that pointer if you will to the element that we're supposed to return.

And then we say that it's only valid if this pointer is less than three because when we reach the value 3 there are no more names that we can give.

So once we move this whole thing we can then it retain its value so we can have another method.

We can have another method called value which returns a string.

And here because of the way things are stored we just look at the current pointer.

So remember this pointer which has an end.

I'm calling it a pointer because it's effectively a pointer it's an index into the set of names.

That person has.
So here we can check which one it is.

I'll switch Pete current and in a K zero we return p dot person dot first name.

And similarly let me just duplicate this in a case of one.

We return a middle name and in the case of two we return last name.

And of course we shouldn't have any other cases but if we do we're going to panic.

We should not be here because obviously we cannot iterate beyond the third element.

Okay.

So now that we have this set up what we can do is we can use this iterator instead.

So this would imply that you make get rid of this so you make an iterator so you say for it's being equal to a new person named iterator where you pass an appointed to the person and we continue iterating while it not move next so move next returns a boolean and we check whether or not this is valid.

So while while this is valid we can print line it's got value.

This gets us the value of the name.

So once again I can run this right here and I get the full name printed out.

So this is the third a variety.

The third approach to how you can implement duration and typically when we talk about the iterator design pattern we mainly talk about explicitly constructed iterator is like this one we talk about separate structures which are used to track the position of where we are in the object that's being iterated and obviously we have appointed to that object so that we can go into it and get some information that we actually need to.

So for a full implementation of an iterator we'll take a look at the next lesson.

That's what we're going to do tree traversal but that's it for now.

So there are three ways in which iteration is possible in go.

### Iteration code: behavioral.iterator.iteration.go

```go
package main

import "fmt"

type Person struct {
  FirstName, MiddleName, LastName string
}

func (p *Person) Names() []string {
  return []string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
  out := make(chan string)
  go func() {
    defer close(out)
    out <- p.FirstName
    if len(p.MiddleName) > 0 {
      out <- p.MiddleName
    }
    out <- p.LastName
  }()
  return out
}

type PersonNameIterator struct {
  person *Person
  current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
  return &PersonNameIterator{person, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
  p.current++
  return p.current < 3
}

func (p *PersonNameIterator) Value() string {
  switch p.current {
    case 0: return p.person.FirstName
    case 1: return p.person.MiddleName
    case 2: return p.person.LastName
  }
  panic("We should not be here!")
}

func main() {
  p := Person{"Alexander", "Graham", "Bell"}

  // range
  for _, name := range p.Names() {
    fmt.Println(name)
  }

  // generator
  for name := range p.NamesGenerator() {
    fmt.Println(name)
  }

  // iterator
  for it := NewPersonNameIterator(&p); it.MoveNext(); {
    fmt.Println(it.Value())
  }
}
```

## Tree Traversal

So like I said the concept of an iterator Ingo isn't as common and as it is in other programming languages.

But I want to show you a scenario where it is in fact relevant and where you cannot get away with without using an iterator and what we're going to take a look at is very simple binary trees so a binary tree is basically a data structure which has some sort of root like let's say we have the number one at the root and then it has two branches.

So every element can either have one two or zero branches are here the number one has branches and we have the nodes 2 and 3 on the end neither to No 3 have children of their own but they could they could have children either one child or both children.

So that is that and with this what we want to be able to do is to set up a data structure where we actually construct these nodes right here and specify the connections between the nodes.

So let's do that we'll have a type node of which is a struct will have the value of the node.

And now I'm going to be using integers for my particular case.

And in addition we'll have we'll have pointers to other nodes.

So this particular node also can have children and can has a can have a child on the left.

So for example here the node with the value one that has a child on the left with a value too and also a child on the right with the value three so left and right are going to be node pointers but we also sometimes want a pointer to the parent node so let's have the parent and all of these are going to be node pointers.

So we have each node has a value and it also has pointers to other elements of the tree.

The left and right children which by the way can be nil.

For example here the node with the value to doesn't have any children.

So those values will be nil and the parent value which most nodes will have a parent value except for the root note.

So the topmost node of a binary tree doesn't have a parent.

So here the parent value will be nil.

And now we need to have a bunch of factory functions for actually initializing all of this and they're not going to be that simple.

So first of all let's make a factory function where you just have the value.

You don't have any information about the parent the left the right side.

Just like this.

So this is a factory function that's usable for the terminal notes the nodes right here because they don't have any children.

And when I'm going to specify the parent right here because well it's it's difficult we don't know who the parent is.

So we'll call this new terminal node so a new terminal node is this path and then we'll generate a more fully initializing factory function where we initialize the value as well as the left and right sides.

But we need to do additional operations here.

So this is the know that we're going to return but before we return it what we want to be able to do is we want to take the left and right sides of the node and we want to set their parents to the current node so we say left not a parent equals and right adult parent equals and and only now can we return the actual node.

So this is that we have now two factory functions one for the terminal node and one for it.

No that isn't terminal.

So these are terminal nodes and this is a more fully specified node.

So we're going to construct those in just a moment actually let's set up the scenario that set up this scenario here.

So the root node is going to be a new node with a value of one and then it has two terminal nodes with a value of 2 and the value of 3.

Now we go.
So this is our route now.

What we want to be able to do is we want to be able to traverse this binary tree and there are different algorithms different ways of traversing this tree.

So the three most common ones are in order pre order and post order.

So in the in order traverse so you would start at the leftmost node and then you would follow from left to right following the nodes and maybe printing them out or maybe just collecting them in some sort of slice or something.

It's really up to you how you handle this.

We're not going to handle preorder and post order but they just want to illustrate that there are different approaches and different approaches will result in you having different iterator.

So what we're going to try and do is build an inward iterator so an in order iterator obviously has to have reference to the root node so that would be appointed to the root node and it also has to have other things besides so let's have a type called in order iterator.

So we're going to have a pointed to the current node.

So this is what you will typically interact with as the iterator is going through the elements.

We also want to have a pointer to the root of the entire tree.

That's going to come in useful later on and remember when we had that previous example we started with an index of minus 1 Well there is a way of avoiding that and starting with the proper index and that is to have just a boolean variable indicating whether or not we return the starting value.

So that's what I'm going to have here I'm going to have a variable called returned start which is going to be a boolean value.

So now what we can do is we can create a factory function which initialize this iterator given a particular route.

And in actual fact there's gonna be lots of stuff to do here because what we want to be able to do is first of all we'll make an in order iterator.

Obviously we have to make it because we have to return it.

So let's let's do that.

So here we specify the current both the current as well as the root elements as root because we're starting with The Root and root is always root.

So we need to keep it like so.

And the return state is return start is initially false so we can do it like this but we're not done yet.

You see in the other iteration implies that you start from the left most element and we are at the moment at the root.

So current is equal to root.

So what we want to be able to do is we want to traverse the entire tree until we find the left most note so we're going to do it right here.

So here I'm going to have a loop.

So while I do a car and left is not equal to nil I'm going to set the current to the left.

So Ida current equals Ida current not left like so.

So this means that when we're starting iteration we're starting with the leftmost element.

So the first element that will be returned is two which is what we want because we are expecting 2 1 3 to be returned.

OK.

So that is how you construct an in order iterator.

But now we're going to have a bunch of utility functions for just working with a situation like for example you might want to reset the iterator.

So here you would have a method called Reset where you would set the current to the root so you would set the current to the roots and you would also say that we didn't return the starting position.

So we'll set false here.
OK so this is it.

And now we're going to have another move next method like we did in the previous demo.

So we'll have a function in the iterator called MOVE NEXT returns a boolean.

So the idea here is that first of all if the current value is nil for some reason we're going to return false because there's nothing we can do so find out current is equal to nil then we return false.

But that's a corner case I don't think we're likely to ever encounter this case.

However if we haven't returned the starting value already.

Fi if not I return and start.

Then we're going to set idle returns start to true and we're going to return true.

Meaning that whoever is iterating this object is welcome to take the current value because the current value has the starting value we've ensured.

Where are the leftmost nodes in this particular loop right here.

So we know that everything is going to be OK now from here on out things get really complicated because we basically have to traverse the entire tree and remember it can be a tree where some of the nodes don't have the left and right.

Children may be some nodes have only the left some nodes only have the right.

So we have to kind of traverse it from left to right correctly and I'm not going to go into the actual details of how this is done instead I want to show you the code for getting it done as you can see it's rather complicated.

So here we perform lots and lots of checks and traversal of different sides and so on and so forth and you're welcome to study this code because it's attached to this lecture.

But what we're going to do is we're going to assume that this implementation of move next is fundamentally correct a traverse traverses the tree in the right fashion and so we're going to start using it we'll start using this iterator right here.

So I'll make an iterator so new in order either later and then I'll pass in the root obviously and then we can start using it so I can have a for loop where I just move the iterator to the next element and that's it.

And for every single such move I can for example print out something let's print out.

So we're going to be printing it not current value and the formatting flag is percent the comma.

So we'll print them comma separated.
And of course we need a line break at the end.
So here I will just print line and backslash B.
Now backslash be a racist the Lost characters.

So that extra comma that we generate at the end of the very last element we iterate will be erased so let's run all of this and let's actually take a look.

Well obviously we forgot something specifically we forgot a return somewhere we forgot to return for the constructed object here.

Let's see if there are any errors.
Hopefully not.

Hopefully everything is okay and lo and behold we have the right output to 1 3 is exactly what we expected.

Okay.

So this is how you use an iterator and just just feed the iterator a starting position if you want to make everything pretty.

Let's suppose that you want to have a really nicely packaged implementation of both in order to reverse it as well as other forms of traversal.

You can go ahead and have a struct called binary tree because remember we are talking about a binary tree but we don't have a strong called binary too.

We just made a bunch of nodes which are linked to one another so I can make a type called binary tree.

A binary tree obviously has to have a root node and then you want to initialize it something like this for example and then what you want to be able to do is you want to the binary tree to actually give you an iterator so you can have a method on the binary tree called let's say in order which returns and in order iterator.

So here we simply return the new in order either eight or passing the root from this binary tree.

That's all there is to it.

Now if you want to iterate this whole thing you would do it similarly to to the way it is done here.

So you would say 4.

So first of all we have to make the binary tree obviously so t is going to be a new binary tree passing in the root.

Once again we don't need this iterator anymore so we create a new binary tree and then we say for iterator initialized to see dots in order.

So this is the initialization step and then the kind of condition step is where we perform no move next.

There is no increment because move next also increments the current pointer.

So here once again we can print out that same thing as before with I dot current dot value so that would be % D comma and then once again we can put a backslash b here and now let's run this up.

This implementation is take a look at what we get here and as you can see we're getting exactly the same output as before.

So this has hopefully been a useful illustration for why you would want to construct different iterator objects.

So in our case we have a tree.

There are different ways of traversing the tree and so we can implement different data structures which know how to traverse them.

So an iterator in this particular case is nothing more than some structure which has obviously a pointer to the elements of whatever it is traversing and it also has in this case a current pointer which is what you would use to actually access the elements on which the iterator is currently stopped.

So that's how you implement the iterator design pattern.

### Tree Traversal code: behavioral.iterator.brokerchain.go

```go
package iterator

import "fmt"

type Node struct {
  Value int
  left, right, parent *Node
}

func NewNode(value int, left *Node, right *Node) *Node {
  n := &Node{Value: value, left: left, right: right}
  left.parent = n
  right.parent = n
  return n
}

func NewTerminalNode(value int) *Node {
  return &Node{Value:value}
}

type InOrderIterator struct {
  Current *Node
  root *Node
  returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
  i := &InOrderIterator{
    Current:       root,
    root:          root,
    returnedStart: false,
  }
  // move to the leftmost element
  for ;i.Current.left != nil; {
    i.Current = i.Current.left
  }
  return i
}

func (i *InOrderIterator) Reset() {
  i.Current = i.root
  i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
  if i.Current == nil { return false }
  if !i.returnedStart {
    i.returnedStart = true
    return true // can use first element
  }

  if i.Current.right != nil {
    i.Current = i.Current.right
    for ;i.Current.left != nil; {
      i.Current = i.Current.left
    }
    return true
  } else {
    p := i.Current.parent
    for ;p != nil && i.Current == p.right; {
      i.Current = p
      p = p.parent
    }
    i.Current = p
    return i.Current != nil
  }
}

type BinaryTree struct {
  root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
  return &BinaryTree{root: root}
}

func (b *BinaryTree) InOrder() *InOrderIterator {
  return NewInOrderIterator(b.root)
}

func main() {
  //   1
  //  / \
  // 2   3

  // in-order:  213
  // preorder:  123
  // postorder: 231

  root := NewNode(1,
    NewTerminalNode(2),
    NewTerminalNode(3))
  it := NewInOrderIterator(root)

  for ;it.MoveNext(); {
    fmt.Printf("%d,", it.Current.Value)
  }
  fmt.Println("\b")

  t := NewBinaryTree(root)
  for i := t.InOrder(); i.MoveNext(); {
    fmt.Printf("%d,", i.Current.Value)
  }
  fmt.Println("\b")
}
```

## Summary

Let's try to summarize what we've learned about the iterator design pattern.

So an iterator specifies how you can traverse an object and what the iterator does is it moves along the iterator collection.

So it moves along the elements of an object indicating also when the last element has been reached.

Because as we traverse the object at some point we have reached the end and the iterator can tell you what element you're at right now but it can also tell you whether or not you've actually reached the end of the whole thing.

Now iterator is not particularly idiomatic and go there is no standard ignorable interface or something like that but they are still easy to build and use provided that you do actually need them.
