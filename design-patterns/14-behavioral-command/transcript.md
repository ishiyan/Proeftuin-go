# Command -- You shall not pass

## Overview

Let's talk about the command design pattern.

So what is the command all about.

Well the motivation for using a command is as follows.

When you make or when you're either ordinary statements like X equals five those statements are perishable.

What do I mean by perishable.

I mean that when you do a field assignment for example you assign the field age the value of 20 you cannot by default undo that operation because there is no record of this operation.

And so you cannot for example serialize the sequence of actions or calls.

You cannot save the set of changes that were done to a particular type by default.

So what we want is we want objects which actually represent operations represent particular requests to some type or other to do something.

So for example you might have command which says that person to change its age to value 22.

And so you would define a structure where you would say who has to change.

In this case person what needs to be changed in this case the age and what the value should be in this case 22.

And similarly you could have a command which for example tells a car to explode.

So that would maybe specify which car and what the car should do.

So there are lots of uses of the command design pattern like if you're making any kind of user interfaces that is typically done using commands and also commands allow you to do multi-level undo and redo so you save every command that was sent to some system and then you can sort of roll them back and you can also group commands together and save a whole sequence of commands and these are typically called macros in goofy applications.

So these are some of the uses of the Command pattern so the Command pattern is quite simply some object which represents an instruction to perform particular action.

And hopefully this command contains all the information necessary for the action to be taken.

And then of course the command has to be processed.

So you also need to choose who gets to process the command.

There are different options here because the command can be processed by the object that the command operates on.

The command can be processed by itself.

It can basically have some sort of code method to apply the command or alternatively what you can do is you can have a special command process set where all the commands are sent.

## Command

In order for us to discuss the command design pattern we're going to set up a very simple scenario.

That scenario is going to be a bank account and a bank account is going to have withdrawal and deposit operations.

So we'll have a type called bank account.

It's just going to have a single member called balance.

I'll also set some overdraft limit let's say minus 500.

And that's just going to make things more interesting for us and then we'll have methods for depositing some money into the account and withdrawing some money from the account.

So we'll have a method for depositing a certain amount and this simply increases the balance so be it a balance plus equals the amount.

But we're also going to output some information just so that we get to see what's actually going on.

So here I'll put that line that we deposited a certain amount and now the balance is beat up balance.

There we go.
We'll have a similar method for withdrawing money.
And here we need to take into account of the overdraft limit.
There we go.

So this is where we get to check that if beat up balance it minus the amount is greater than or equal to the overdraft limit.

Then we can perform that withdrawal.

So here we can say beat up balance minus equals amount like so and once again we'll print line I'll actually just copy this over so we'll print line that we withdrew a certain amount and now the balance is beat up.

Balance.

So that is the implementation of the deposit and withdrawal methods.

And now what we can do is we can try to set up the command pattern.

So there are different ways of actually handling commands one way is where the bank account itself handles the command.

The other is when the command handles itself and we're going to consider this particular scenario we're going to consider a command which is just an interface which has a member called Call and when you call the command whatever modifications on the account that needs to happen do in fact happen.

Now in terms of the actions that we can take upon the account there's gonna be two of them.

So I'll have type action as an end and we'll have just a bunch of constants.

So we'll have a way of depositing money so deposit is one operation.

It's going to be a type of action and we'll also have a withdrawal operation withdrawal like so.

So now we can define a bank account command.
So we'll have type bank account command.
It's gonna be a strike.

Now in the bank account command will have the account which is going to be a pointer to bank account.

We'll also have the action that we want to take as well as the amount which relates to this action.

So if you want to deposit this is the amount you want to deposit.

And similarly for the withdrawing process and now what we want to be able to do is we want to be able to obviously call the action.

So we want to implement the interface.
So let's implement the action.

Well actually it's the command interface the command interface that we want to implement.

So here is the implementation of call here we simply look at the type of action that is required.

So we switch on beat out action and action like so and depending on the action we perform either the withdrawal or the deposit.

So here for example in the case of a deposit we say B dot accounts dot deposit and we deposit B on a mound.

Otherwise if it's withdrawal.

So case of withdrawal we say B dot account withdrawal and we specify the amount we want to withdraw.

So that's really all there is to actually setting up a bank account command.

We can also make a factory function for that just initializing all the different members like so.

And then we can start using this whole thing.

So here what I can do is I can make a bank account so I'm going to leave it at zero dollars to begin with.

And then we can make a new deposit command.
So let's make CMG as a new bank account command.

So here we specify the bank account to operate upon with specify that we want to deposit and we want to deposit a hundred for example and let me just replicate this once again.

So there's going to be CMG to where we try to withdraw 50 withdrawal 50 dollars.

Okay.

So here after I create the first command I can CMG you know call to actually invoke the command.

Then I can print line the actual bank account that's just going to output us a number.

Give us a number for the balance and then we can CMG to adult call just to see what is going on here and we can once again printout the account and and see how the whole thing goes.

So now let's try running this.

Hopefully it all compiles and gives us something as the output.

And here we go so we deposit a hundredth so the balance is now a hundred.

Then here is the printout of the state of the bank account and then we withdraw 50 so the balance is now 100 minus 50 so that is 50 and once again we get the correct output.

So this is the approach to implementing the Command pattern where the command set of calls itself.

So the command has some sort of call method and then you specify the actual commands which implement the command interface and then that's where you actually perform the call and then check what actually needs to be done then simply do that on the object which you are operating upon.

So that is a very simple implementation of the Command pattern.

### Command code: behavioral.command.command.go

```go
package command

import "fmt"

var overdraftLimit = -500
type BankAccount struct {
  balance int
}

func (b *BankAccount) Deposit(amount int) {
  b.balance += amount
  fmt.Println("Deposited", amount,
    "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
  if b.balance - amount >= overdraftLimit {
    b.balance -= amount
    fmt.Println("Withdrew", amount,
      "\b, balance is now", b.balance)
    return true
  }
  return false
}

type Command interface {
  Call()
  Undo()
}

type Action int
const (
  Deposit Action = iota
  Withdraw
)

type BankAccountCommand struct {
  account *BankAccount
  action Action
  amount int
  succeeded bool
}

func (b *BankAccountCommand) Call() {
  switch b.action {
  case Deposit:
    b.account.Deposit(b.amount)
    b.succeeded = true
  case Withdraw:
    b.succeeded = b.account.Withdraw(b.amount)
  }
}

func (b *BankAccountCommand) Undo() {
  if !b.succeeded { return }
  switch b.action {
  case Deposit:
    b.account.Withdraw(b.amount)
  case Withdraw:
    b.account.Deposit(b.amount)
  }
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
  return &BankAccountCommand{account: account, action: action, amount: amount}
}

func main() {
  ba := BankAccount{}
  cmd := NewBankAccountCommand(&ba, Deposit, 100)
  cmd.Call()
  cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
  cmd2.Call()
  fmt.Println(ba)
}
```

## Undo Operations

So one of the things we can implement under the current command paradigm is we can implement undo functionality so after you call the command you can sort of roll back the changes that the command introduced.

Now in order to do this we need to modify a couple of things. So first of all we'll add the undo method to the command interface.

But in addition one problem that we'll have is that some of the commands are actually going to fail like for example here when you withdraw a certain amount you want to withdraw it.

Only if the decreased amount is greater than or equal to the overdraft live and otherwise the command doesn't apply.

But if the command doesn't imply that means you shouldn't be able to undo it either because that will

leave the system in an unpredictable state.

So what we need to be able to do is we need to be able to have some sort of indicator whether the command succeeded or not.

So here I'm going to change withdraw adding a boolean return flag that will set to actually tell you whether the command succeeded.

So in the case when the there is enough money we return true.

But in the case when there isn't enough money we return false.

So another question is Well where exactly do we use this information.

And the answer is that we can use this whenever we actually perform the action.

So when we call the command we can also store some information about whether or not the command actually succeeded.

So let's go into command and we're going to add a new flag here succeeded so succeeded is going to be a boolean flag indicating whether or not the command actually went through fine or if it's false.

That means the command was never applied in the first place.

So when you come to undoing the command there is nothing for us to do.

So here in the case of deposit we're going to assume that a deposit always succeeds.

So we say beat out succeeded.
Equal is true.

But in the case of withdrawal we actually use the return result of the withdrawal method and store it in b succeeded.

So this is how we can set the whole thing up and now we can implement the undo method.

So now we can go ahead and once again I'll ask the system behind the scenes here to implement the the command interface for the bank account command.

So we get the undo method here and this is where we get to undo the command.

But you can only undo the command if it actually succeeded.

So if the command if the command did not succeed then you can simply return because there is nothing for us to do.

There is no operation that needs to be done here.

But if the command did in fact succeed then we can look once again at the action that was performed and we can try and doing the opposite of what was done.

So if somebody deposited a certain amount of money and then what we can do is we can say BDO accounts to withdraw and now I know that's not strictly correct.

You cannot consider these operations symmetrically.

But in our simple example that's going to work just fine.

So in the case of withdrawal for example we're going to just perform a deposit.

So we're going to deposit a certain amount.
Let's see them out as before.

So now that we have this kind of setup with the undo what we can do is we can start using it.

So after we call the command here and we can change this number to something like 25 so that there is seventy five dollars remaining what we can do is we can try undoing CMT too so we can try undoing CMT too.

And then we can print the bank account once again.

So that should hopefully get us to the initial state that we were in.

So we deposit a hundred.
The balance is hundred.
Then we withdraw twenty.

So the balance in 75 and then we perform a symmetrically opposite operation by depositing 25 set of undoing.

This particular command.

So the balance is now back to 100 is back to a state that it was in here.

And of course we could undo this command as well thereby getting a balance of zero which is the starting balance of the whole system.

So this is how you can implement undo operations under the command design pattern.

## Composite Command

Imagine that you want to transfer money from one bank account to another bank account.

You can consider this as a combination of two different commands and one of those commands would withdraw the money from the first account.

And the second command would deposit the money into the second account.

However there are a few limitations to this model.

For example they both should succeed or fail meaning that if you fail to withdraw money from account one then you shouldn't be able to deposit this money into account.

Two that simply doesn't make sense.

The whole transaction should fail a transaction consisting of these two different operations.

So what we're going to do is we're going to implement something called a composite come out now a composite command is a command which itself can consist of several other commands and it represents both the command design pattern as well as the composite design pattern.

But we need to once again make a few modifications to the command interface.

See now the problem is that every single command that we operate on has to be able to have this success flag it has to tell us whether it succeeded or not.

So the way to do this at the interface level is to add a method called succeeded which returns a boolean and also to add a setter so we add together and a setter like so so that whenever there is a concrete command that implements this interface it actually has to somehow operate upon this idea of success.

It has to have a well whether it has a flag or not doesn't matter but it has to have those members.

So coming down here to bank account command we now need to once again go ahead and implement the command interface because there is a bunch of stuff missing.

So we need to implement succeeded and said succeeded not because we have a field here called succeeded.

That's not going to be a problem.

We can just say BDO succeeded equals well it should be down here beat out succeeded.

Equals of value here.
And here we return b succeeded.

So there is really no problem in US satisfying this interface.

Now when it comes to performing the ordinary sort of deposit and withdraw operations there's really no change.

Their operations are pretty much the same so we can go ahead and we can build a composite bank accounts command.

So like I said a composite bank account command is going to be just a collection of commands so we can define it as follows We can have a type called composite bank accounts command.

And here you can specify the commands.

So that's going to be just a bunch of commands put together.

Okay.

So we now also need to implement the command interface on composite bank account command.

So let's do that.

So I'm going to go ahead and implement the command interface here.

Okay.

So we need to be able to call the composite command we need to be able to undo the composite command and we also need to set and get the sexy the flags.

So let's start with calling the command.

Now obviously in our case what's really happening is you're simply calling every single one of the commands.

So here we go through every single one.

So for underscore pharmacy M.D. in the range of seed commands and we simply do CMG in our call and that's all there is to it really there's nothing else for us to do.

Now we also need to be able to undo all the commands.

Now we need to go through them.

That's true but we need to go through them in reverse order starting from the last and ending at the first because the order in which they were applied is that is the reverse of the order in which they need to be undone.

So here I'm going to say for index in range of CDO commands and here what I'm going to do is I'm just going to get the commands from the end.

So I'm going to say CDO commands at position and then I'll take length of CDO commands minus index minus one and then I will on to it this way.

So we are going from the right most to the leftmost so to speak from the last to the first.

So this is how you undo a set of commands.

Now we also have this pesky annoying issue of getting and setting succeeded flags.

So how do we do this we need to set succeeded and get succeeded even though it doesn't really make much sense because the success flags are set by the invocation of the command.

But if you do want to implement this interface completely then you have to do it like this.

So in the Seder you just go through every single command and for every single command you call C.M. the set succeeded specifying the value that's specified right here and now when you want to check whether or not a composite command succeeded.

That implies that every single part has succeeded.

So you go through every single element and if that element has failed you return false.

Otherwise you return true.

So for underscore comma CMB in a range of CDO commands what we do is we say if not CMT dot succeeded.

So if the command didn't succeed then we return false.

And if we managed to go through this loop without returning false that means everything is fine and we simply return true.

So this is a correct way of implementing this succeeded together.

OK.

So now we have this composite command and we we can certainly start using it in the sense of just taking for example this deposit and this withdrawal and adding it to a set of commands just just making a composite out of it.

However that's that's not really exciting that's not why we did this in the first place.

We want to perform a transfer of money from one bank account to another bank account.

So that's going to be something that will aggregate the composite bank accounts command.

So so we're going to compose both have some kind of money transfer come out so money transfer command is going to be a struct where we'll have a composite bank account command inside it will also have specification of the bank accounts from which to take money and to which to deposit money.

So these are gonna be bank account pointers.

And of course we'll have the amount that needs to be actually transferred.

Now here is the interesting thing we need to make a factory function for this struct because it needs to be initialized correctly.

So if I just go ahead and I generate a factory function specifying all the different parts then it's going to look something like this.

So we have from Two and a mound but that's really not enough.

So the first thing we do is we make the money transfer command so C is going to be a money transfer command where we specify the from the to and the amount.

But we also need to initialize those commands because there's gonna be two two commands in this composite for withdrawing the money from the first account and depositing the money to the second account.

So I'm going to go see what commands and I will append to it.

So I'm going to append a new bank account command from a particular bank account.

We're going to withdraw this particular amount.

So that's command number one and then seed out commands the pen.

So I will append another command.

And this is going to be a new bank account command where we take it we take the money and we deposited to the second account.

So we do it like this and then and only then do we return the money transfer command that we've constructed.

Okay.

So this is our setup and we are not done yet.

Unfortunately there is one more thing that we need to do when it comes to transferring money.

You see if the first operation fails.

So if this command fails then it makes no sense for us to perform this command.

Makes absolutely no sense and it certainly makes no sense for us to undo this command later on because either both commands succeed or they both fail.

So we need to redefine the call method which is defined up here so if you if you look up here we have just a simple invocation of every single command.

But this doesn't work for us when we transfer money we need to perform the second call only the first call succeeded or not.

So so that is what we're going to implement so we'll have a function which is defined on on the money transfer command called call.

So this is going to be an alternative definition of the command invocation.

We'll have an OK flag here initially set to true and then we go through every single command.

So we go through every single command in the range of MDA commands and we check whether or not this flag is OK.

So this flag is going to kind of short circuit the whole invocation basically if we failed at some point then we're not going to continue.

But if everything is okay then we can take the current command we can call the command and then we said the Okay flag to whether or not the command actually succeeded.

However if we did not get an okay then that's that's pretty much all that needs to be done here so we would simply say that well everything failed.

So so we can see CMC dot set succeeded to false because everything has failed here unfortunately.

OK.

So with this setup what we can do is we can finally get a scenario where we're transferring commands from one bank account to another.

So let me wipe out all of this and we'll do this once again so we'll have an account called from.

It's going to be a do we have a.
No we only have it for commands.

So let's just make a bank account with a balance of 100.

We'll have another bank account called to where the balance of 0 here and then I'll make a money transfer command.

So money transfer command is going to be a new money transfer command where we transfer from the account to this account we try to transfer let's say a hundred dollars.

Actually let's make it something like 25.
So you can better see the end result.

So we take the money transfer command we call the command and then we can print line.

The State of the two accounts from.
And two.

OK so let's first of all run this and let's take a look at what happens here so we perform the transfer so we withdraw twenty five from the first account and we deposit twenty five onto the second account.

Remember it was empty initially and now it's at twenty five.

So the balance is seventy five and twenty five.

Now let's try the undo operation just to make sure that this also works.

So I'll do empty seat or undo like so.

And then we can print out the same info once again.

So let's go ahead and run this now so you can see that when we perform the undo operation we perform kind of symmetrical operations here and here.

And we're back to the starting balance so we're back to the balance of 100 in the first account and zero in the second account.

So there you go.

This is how you implement the composite design pattern so if you have other operations that you also need to handle as a single operation then you can simply aggregate the composite bank account command to get basically just the list of the list of commands as well as the default implementation of undo because undo is one operation which we didn't redefine and.

Well actually no that's that's not the right way of saying it.

We redefined call.

We left every other operation the same way we defined calls so that he would be internally consistent here in terms of just succeeding completely or failing completely.

But yeah you can you can extend you can use competent bank account command to make additional composite commands.

So if you have some complicated transfer that jumps through several banks for example you would define it as several commands.

You would simply stick them into the set of commands like we do here and then they would all either succeed or fail depending on what's actually going on.

So this is how you can marry the Command pattern and the composite pattern.

### Composite Command code: behavioral.command.compositecommand.go

```go
package main

import "fmt"

var overdraftLimit = -500
type BankAccount struct {
  balance int
}

func (b *BankAccount) Deposit(amount int) {
  b.balance += amount
  fmt.Println("Deposited", amount,
    "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
  if b.balance - amount >= overdraftLimit {
    b.balance -= amount
    fmt.Println("Withdrew", amount,
      "\b, balance is now", b.balance)
    return true
  }
  return false
}

type Command interface {
  Call()
  Undo()
  Succeeded() bool
  SetSucceeded(value bool)
}

type Action int
const (
  Deposit Action = iota
  Withdraw
)

type BankAccountCommand struct {
  account *BankAccount
  action Action
  amount int
  succeeded bool
}

func (b *BankAccountCommand) SetSucceeded(value bool) {
  b.succeeded = value
}

// additional member
func (b *BankAccountCommand) Succeeded() bool {
  return b.succeeded
}

func (b *BankAccountCommand) Call() {
  switch b.action {
  case Deposit:
    b.account.Deposit(b.amount)
    b.succeeded = true
  case Withdraw:
    b.succeeded = b.account.Withdraw(b.amount)
  }
}

func (b *BankAccountCommand) Undo() {
  if !b.succeeded { return }
  switch b.action {
  case Deposit:
    b.account.Withdraw(b.amount)
  case Withdraw:
    b.account.Deposit(b.amount)
  }
}

type CompositeBankAccountCommand struct {
  commands []Command
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
  for _, cmd := range c.commands {
    if !cmd.Succeeded() {
      return false
    }
  }
  return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(value bool) {
  for _, cmd := range c.commands {
    cmd.SetSucceeded(value)
  }
}

func (c *CompositeBankAccountCommand) Call() {
  for _, cmd := range c.commands {
    cmd.Call()
  }
}

func (c *CompositeBankAccountCommand) Undo() {
  // undo in reverse order
  for idx := range c.commands {
    c.commands[len(c.commands)-idx-1].Undo()
  }
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
  return &BankAccountCommand{account: account, action: action, amount: amount}
}

type MoneyTransferCommand struct {
  CompositeBankAccountCommand
  from, to *BankAccount
  amount int
}

func NewMoneyTransferCommand(from *BankAccount, to *BankAccount, amount int) *MoneyTransferCommand {
  c := &MoneyTransferCommand{from: from, to: to, amount: amount}
  c.commands = append(c.commands,
    NewBankAccountCommand(from, Withdraw, amount))
  c.commands = append(c.commands,
    NewBankAccountCommand(to, Deposit, amount))
  return c
}

func (m *MoneyTransferCommand) Call() {
  ok := true
  for _, cmd := range m.commands {
    if ok {
      cmd.Call()
      ok = cmd.Succeeded()
    } else {
      cmd.SetSucceeded(false)
    }
  }
}

func main() {
  ba := &BankAccount{}
  cmdDeposit := NewBankAccountCommand(ba, Deposit, 100)
  cmdWithdraw := NewBankAccountCommand(ba, Withdraw, 1000)
  cmdDeposit.Call()
  cmdWithdraw.Call()
  fmt.Println(ba)
  cmdWithdraw.Undo()
  cmdDeposit.Undo()
  fmt.Println(ba)

  from := BankAccount{100}
  to := BankAccount{0}
  mtc := NewMoneyTransferCommand(&from, &to, 100) // try 1000
  mtc.Call()

  fmt.Println("from=", from, "to=", to)

  fmt.Println("Undoing...")
  mtc.Undo()
  fmt.Println("from=", from, "to=", to)
}
```

## Functional Command

There is one thing I wanted to mention and that thing is that strictly speaking you can take a more functional approach to the way that you create commands as well as the way you create composite commands.

So here I have a bank account with a typical deposit and withdrawal methods.

There's somewhat more simple than what we had previously but it doesn't really matter.

So if you imagine a bank account as just a bank account with starting balance of zero what you can do is you can define the commands but instead of defining them as separate trucks which define the entire set of operations that need to be performed on a bank account what you can do is you can stick those operations inside simply a list of functions.

So what we can do is we can have var commands as just a bunch of functions and then we can append those functions.

So for example if I want to let's say withdraw so let's say I want to deposit 100 dollars and withdraw twenty five dollars the way I can do this is I can say commands and I can append to the commands and so here I can make a function.

And inside this function I can perform the deposit so I can take the bank account and I can deposit let's say 100.

And similarly I can do.

I can do a withdrawal here so far function and then withdraw.

So bank account withdrawal let's say 25 for example now that I have this list of commands what I can do is I can just go through this array.

So for underscore comma EMV in range of commands I can simply call CMG and this is going to perform the deposit and then perform the withdrawal let me actually print out the the end result cells print line let's just print line the bank account like so.

OK.

So let's just run this let's take a look at what we get.

So we're depositing a hundred then we're withdrawing twenty five and then the end result is seventy five.

So what does this give us.

Actually well on the one hand you lose all the information regarding what kind of operation you were actually doing so in the previous examples when we had the command you could save the command to a file or send it over the Y.

Here we are wrapping it into a function and so we're losing all the information about what's actually going on meaning we cannot just go and save this function somewhere or go into this function and look at what exactly is going on there because well remember.

Strictly speaking you can have more than one thing happening inside a single command.

You could have something like this happening inside a command.

On the other hand this is a more functional approach and it does have it does have its uses sometimes meaning that if you don't care about the structure and you just want to put put several invocations into a list for that list to be invoked or you know for example for the entire list to be undone at some point then a functional approach is equally viable

### Functional Command code: behavioral.command.functionalcommand.go

```go
package command

import "fmt"

type BankAccount struct {
  Balance int
}

func Deposit(ba *BankAccount, amount int) {
  fmt.Println("Depositing", amount)
  ba.Balance += amount
}

func Withdraw(ba *BankAccount, amount int) {
  if ba.Balance >= amount {
    fmt.Println("Withdrawing", amount)
    ba.Balance -= amount
  }
}

func main() {
  ba := &BankAccount{0}
  var commands []func()
  commands = append(commands, func() {
    Deposit(ba, 100)
  })
  commands = append(commands, func() {
    Withdraw(ba, 100)
  })

  for _, cmd := range commands {
    cmd()
  }
}
```

## Summary

All right let's try to summarize what we've learned about the command design pattern.

So the idea of a command is that you encapsulate all the details about an operation and you put them into a separate object the separate structure and then what you do is you define functions for applying the command.

And there are different ways of handling this you can put it into the command itself or you can put it elsewhere.

It's really up to you and optionally You can also define instructions for undoing the commands for returning a system to a state before the command was actually applied.

And you can group commands together you can create composite commands also known as macros and this is also a good illustration of the merger of the chain responsibility design pattern as well as the composite design pattern.
