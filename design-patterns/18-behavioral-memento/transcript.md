# Memento -- Keep a memento of an object's state to return to that state

## Overview

Let's talk about the memento design pattern.

So what is this all about.

Well as you have objects in the system those objects go through changes.

For example if you have a bank account the bank account balance goes up and down depending on the number of deposits and withdrawals in the system.

Now there are different ways of actually navigating those changes and we already looked at some of them.

So one way is to record every change by using let's say the command design pattern and then you can also teach the command to undo itself.

And this is part of something called Command query responsibility segregation an approach which sometimes is also coupled with event sourcing and that is something that we would typically discuss as part of enterprise patterns as opposed to conventional design patterns.

But there is an alternative approach to using the command.

Another approach is to simply save snapshots of the system.

So whenever you need to you can save a snapshot of the system and you can return to that snapshot.

So the memento is some token which represents the state of the system and it allows us to roll back the entire system to the state when the token was actually generated.

Now tokens have different behaviors you may or may not directly expose the state information.

Typically you probably want to hide it if you can but it really depends on the implementation.

The key idea here is that you are given some token and you can use that token to return the system to that particular state when the token was produced.

## Memento

All right.

So what we're going to take a look at now is a simpler implementation of the memento design pattern.

So the idea of a memento is that basically whenever you have a change in the system you can take a snapshot of the system and you can return that snapshot to the user so they can subsequently return the system to that state.

So for example let's imagine you have a bank account.

So a bank account is just going to have a single member called Balance which is going to be an int and let's suppose that we want to allow the depositing of money onto the banking out.

So func bank account is going to be a method called deposit where you specify the amount number of dollars you want to deposit but instead of this returning nothing we're going to return a memento.

So what is this memento.

Well the memento is going to take a snapshot of our system and our system is rather simple we only have the balance of the system so we can have a type memento.

It's going to be a struct which is going to have a balance has an end.

And so when we return this balance this is going to be the balance initialized after the operation has been made.

So first of all we modify the account we say beat up balance plus equals the amount.

So we increase the balance by the amount and then we return a memento which preserves this new balance which sets the balance to beat off balance.

So this is how you perform some modification on the system.

And then of course you need a mechanism for actually restoring the system to a state represented by the momento.

So here we can have another function.

Another method on the bank account and this one is going to be called Restore restore where we take the memento and we restore from that memento.

So we said bank account balance to the balance specified by the memento.

OK.
So how does all of this work.
Well let's take a look.

So first of all I make a bank account so it's going to be a bank account with 100 dollars in it and then I will lose deposits a money a deposit.

Let's say 50 and but remember this is not.
This does have a return value.
So we're going to have a variable for it.

So this is going to be m 1 the first Memento and then we'll have M2 which is going to be well similar to this.

So we'll have M2 which deposits 25 for example so we're starting with 100 ballot dollars then we get to 150 here and then we get to 175 here so we can just print line the bank account.

It doesn't have a stringer interface is just going to be a very raw output.

And then we can try restoring the state of the bank accounts to the mementoes which are stored here and here.

So let's first of all restore the bank account to M1 so we can save the aid or restore to the M one and then we can copy this line and just print the output and then we can do this once again this time restoring to M2 and we can once again print out the state the bank account.

So let's actually run all of this and let's take a look at what we get here.

So the starting bank account has one hundred and seventy five dollars.

Then we restore it to the first Memento and the first memento remember that's where we deposited 50.

So we're back to 150 and then we restore back to the second memento.

So we are now back to 175.
So this is great for restoring a system to a previous day.
There's only one problem though.

That problem is that whenever we actually have this kind of setup we don't have a memento for the initial state of the bank account and if you want to have a mentor for the initial state of the bank account then you'd have to improvise somehow because remember if you make a typical method for a typical function for just initializing the bank account there's going to look something like this and the problem with this setup is that what you basically you return the state of you return the bank account that you've constructed but you're not returning the memento here.

And of course you could modify this.

You could for example say well let's also return the memento.

And then what would happen is you would actually make them the mental here and you would you would effectively return to things so you'd return the bank account but you would also return the actual memento where the bank account balance is also balance so something like that this you would be able to do this kind of thing and it's suddenly possible.

And let me just clean this up a little bit so that it's a bit more readable but this is how you would do it.

It's possible to do store both of these it's not particularly convenient though but it is possible to return the initial memento the memento representing the initial state and if you do that then we can do something like the following so we can have M0 here and this is going to be a new bank account with a balance of 100 and then eventually you can restore to m 0 as well so be it or restore to M0 and this should give us a if we just copy this over.

This should give us in the state of just 100 dollars hopefully so.

So let's just run this let's take a look and that's exactly what we're getting here towards the end.

So this is how you implement a very simple the mental design pattern.

### Memento code: behavioral.memento.memento.go

```go
package memento

import (
  "fmt"
)

type Memento struct {
  Balance int
}

type BankAccount struct {
  balance int
}

func (b *BankAccount) Deposit(amount int) *Memento {
  b.balance += amount
  return &Memento{b.balance}
}

func (b *BankAccount) Restore(m *Memento) {
  b.balance = m.Balance
}

func main() {
  ba := BankAccount{100}
  m1 := ba.Deposit(50)
  m2 := ba.Deposit(25)
  fmt.Println(ba)

  ba.Restore(m1)
  fmt.Println(ba) // 150

  ba.Restore(m2)
  fmt.Println(ba)
}
```

## Undo and Redo

So one natural consequence of returning a momento for every single change in a struct is that you effectively have the whole history of the system as it's modified.

And as a consequence of that what you can do is you can implement undo and redo functionality as part of the momentum design pattern meaning that as soon as you have this history of events you can implement undo and redo because you have every single state preserved in a memento.

Obviously this is not always practical sometimes your memento will simply store too much information.

But in our example with a bank account we can actually set this kind of thing up.

So let's actually try doing that.

So once again I'll have a bank account but this time things are going to be different.

That's why I'm typing things once again.
So we'll have a memento.

It's gonna be a struct which just preserves the balance and we'll have a bank account.

Now this time round the bank account is going to have three members so first of all we'll have the balance as before balance and we'll also have a storage basically an array of every single memento that's been generated by the bank account.

So I'm going to call this changes because that's effectively what they are and is going to be a slice of Memento pointers pointers are useful here we're going to make use of the fact that they are pointers in just a moment.

And then of course we'll have to have some sort of indicator saying which Memento is the current memento.

So remember as you are doing undo and redo you're basically moving a pointer inside this array of Memento pointers.

So we'll have current as an integer indicating our position within this array.

Okay.

So now that we've said to all of this off let me actually implement the string Stringer interface on the bank accounts so we can print a few things.

So let's just implement Stringer like so.

So what I'm going to do here is I'll just F.A. that s print so I'm going to say F.A. s print.

So I'm going to say that the balance is equal to veto balance.

And I'll also print out the current pointer that current variable that we have so current is equal to the current just so that we get to see where that thing is actually pointing to and then we'll have a constructor for bank account.

So.

So the constructor is going to initialize just the balance but also we'll do other things here because remember we need to add this initial change to the set of changes so we need to generate a memento.

So here we'll have the variable B.

That's our that's going to be our new bank account but also we'll take BDO changes and we'll append the first memento ever so we'll append a memento where we save that initial balance and then and only then did we return b.

So this is how we're going to set the whole thing up.
So now let's have a deposit operation.

I'm not going to have withdrawal operation but we'll have an operation for depositing money to the bank account and you'll see that it's rather more complicated than what we had before so we'll have a function on bank account.

So it's gonna be a method of bank account called deposit where deposit is certain amount int and we return a memento pointer.

There we go and lots of things are going to happen here.

So first of all we increment the amount be the balance so we increment the balance by the amount then we manufacturer the memento to preserve this change.

So we say m is equal to memento with BDO balance.

There we go we now have in the mantle we need to add this momentum to the set of changes on the account so be that changes.

We just append to it basically append this particular Memento and then we also need to move the pointer forward remember this pointer the current integer we need to we've just added an additional step.

So we're moving to the right.
So we need to increment this.
So we say veto current plus plus.

So as you can see there's lots of stuff this four different things that have happened here and let's also just print out some diagnostic information so we deposited a certain amount.

And so the balance is now.
And then just print the balance and that will go.

And then finally we return the memento of the memento pointers to be specific.

OK.

So now that we have the memento we can obviously restore to that memento.

So once again let's define this function.
But now it's going to be slightly different.
So we'll have a method on bank account called Restore.

So you specify a memento pointer and you restore to that memento.

However there's gonna be a critical difference here and the difference is that we're going to check that the memento is not equal to nil.

So we're going to be using nil values in situations where you have a an invalid memento that's a memento that doesn't point anyway.

So so and new memento basically means let's just ignore this whole thing.

You'll see why it's needed in just a moment if the momento is not know then we take the balance and we decrease it by MDL balance.

So we kind of restore the system to a previous state.

Actually no we didn't decrease it.
We set it to the to the momento.

And then what we do is we need to also record this change.

So this is a change of balance so we record this change so beat up balance append BDO changes rather by the changes so we append this change to the set of changes because every single application of the momentum to restore is itself a change which in and of itself might generate a memento.

But but we just keep it here internally.
And then finally we set the pointer.

So the current pointer is now pointing to the lost element of that whole thing so beat changes length.

So we take the length of the changes and we subtract one and that gets us a pointed to the lost element.

So this is how you restore a memento and now we can implement those fancy undo and redo operations and then I'm going to be particularly simple.

I'm afraid so.

First of all let's implement undo so func B bank account undo returns a a memento pointer.

Now here's the thing.
If we don't have anything to undo that means the.

For example let's suppose that BDO currencies equal to zero so we cannot undo anything so we're going to return nil.

And remember that nil is going to be a momento pointer and that's why we have a Neil check right here.

Because sometimes undo and redo will generate no values and if you feed those Neil values to restore you still won the whole thing to operate correctly.

So assuming that b the current the current pointer is greater than zero that means there is something for us to undo so we say B the current minus minus that we take the memento from BDO changes changes at BDO current so we take the previous memento so to speak.

We set the balance to the mementoes balance and then we return that memento.

And if the whole thing is nil right here then we return nil.

So if the argument is actually nil then we simply.

But so I know we don't return any mementos for this operation because there is no operation for us to do and the same thing goes for the reader.

So let me just type this rather quickly so readers is also method of bank account redo also returns a memento pointer and here if so we need the current pointer plus one to be less than the length of

changes in order for us to basically move forward one step so if ft or current plus 1 is less than the length of BDO changes then we can do something we can move current forward one place so we can get the memento Memento is going to be done.

Changes add veto current and then we can say beat up balance people's end of balance and return em like so otherwise once again where it's a new meaning that there is nothing for us to read.

So if you're at the end of the stack of changes or a list of changes now.

Case then there is nothing for you to redo there's nothing there so you just return nil there.

So let me show you how we can put all of this together.

So bank account is going to be a new bank account with a balance of 100.

I'm going to deposit 50 and I'm going to deposit 25 so let's print let's print line.

The State of the bank account and now we're going to undo a couple of times.

So be a or undo undo like so and let's actually print that after undo 1. We have the following state of the bank account let's duplicate this and let's do this again.

So after undo 2 you should be able to go even further in terms of history and then let's redo let's say be a redo and let's once again print line the state bank account after the redo.

Okay so let's see if everything that we've written actually works.

So as you can see first of all we have the actual deposit operation so we deposit it 50 and the balances 150 with deposit 25 the balance is 175 so we print out the bank account the balance is 175.

The current pointer it points to elements 2 because we effectively have 3 changes so the last change has position to the first changes when we deposit the initial 100.

The second change is right here the third changes right here so we have three changes so the pointer has the value of two.

And now what we do is we perform the undo.

So we jumped from 175 back to 150 and notice that the current value changes to 1 and then we undo the whole thing once again and we jump to a hundred and the current value is now 0 because the only change that's been made is the change for setting the initial balance.

That's that was change no zero.

And then we redo the operation of depositing 150 so the current now changes to one.

So in this demo I've illustrated how you can leverage the momentum design pattern in order to implement undo and redo functionality.

### Undo and Redo code: behavioral.memento.undoredo.go

```go
package memento

import "fmt"

type Memento struct {
  Balance int
}

type BankAccount struct {
  balance int
  changes []*Memento
  current int
}

func (b *BankAccount) String() string {
  return fmt.Sprint("Balance = $", b.balance,
    ", current = ", b.current)
}

func NewBankAccount(balance int) *BankAccount {
  b := &BankAccount{balance: balance}
  b.changes = append(b.changes, &Memento{balance})
  return b
}

func (b *BankAccount) Deposit(amount int) *Memento {
  b.balance += amount
  m := Memento{b.balance}
  b.changes = append(b.changes, &m)
  b.current++
  fmt.Println("Deposited", amount,
    ", balance is now", b.balance)
  return &m
}

func (b *BankAccount) Restore(m *Memento) {
  if m != nil {
    b.balance -= m.Balance
    b.changes = append(b.changes, m)
    b.current = len(b.changes) - 1
  }
}

func (b *BankAccount) Undo() *Memento {
  if b.current > 0 {
    b.current--
    m := b.changes[b.current]
    b.balance = m.Balance
    return m
  }
  return nil // nothing to undo
}

func (b *BankAccount) Redo() *Memento {
  if b.current + 1 < len(b.changes) {
    b.current++
    m := b.changes[b.current]
    b.balance = m.Balance
    return m
  }
  return nil
}

func main() {
  ba := NewBankAccount(100)
  ba.Deposit(50)
  ba.Deposit(25)
  fmt.Println(ba)

  ba.Undo()
  fmt.Println("Undo 1:", ba)
  ba.Undo()
  fmt.Println("Undo 2:", ba)
  ba.Redo()
  fmt.Println("Redo:", ba)
}
```

## Memento vs Flyweight

So one thing I wanted to mention is the differences between the momentum pattern and the fly away pattern because they're kind of similar both patterns provide some sort of token that the client can hold on to but the momento is used only to be fed back into the system so it doesn't have any public mutable state it doesn't have any methods.

It's just a token it's just this tiny little piece of data that you can submit so that the system would return or roll back to a previous state whereas in the case of a fly away it is kind of similar but it's a reference to an object.

So in actual fact it can mutate state.

There is no problem in the fly with mutating state and it can also provide additional functionality so the fly wait can have its own fields or methods which actually do something.

## Summary

So let's try to summarize what we've learned about the memento design pattern.

So mementos I used to roll back states arbitrarily.

And a memento is quite simply some sort of token or handle.

And it typically has no methods of its own.

It typically doesn't have any behavior.

And if it does have fields then we might want to try and hide those as much as we can because while changing a memento is an idiomatic typically a memento is the kind of redone the object.

So a memento is not required to expose the state to which era of the system and in actual fact it can be problematic if it does because if it exposes some state then somebody could modify it and then return the system to the state in which it never was.

So that's a particular problem of the memento.

Now a memento can also be used to implement under and redo although this implementation is somewhat clunky because remember in order for this to work work you basically have to save every single memento.

So you have to save the state of the system at every point in time which is not particularly practical at least with the command pattern you just save the changes put the memento you save the entire state.

So if your system is really simple then you can record ever single State of the system it's no problem but if your system is complicated you will be taking huge snapshots with lots and lots of data being replicated over and over again for each Memento and there's just too much memory traffic and too much computation for us too to make this approach realistic.

So overall I would say that undo and redo is better handled by the command design pattern and event sourcing and all that sort of thing rather than by mementos.

But if your system is simple then yeah you can use mementos instead.
