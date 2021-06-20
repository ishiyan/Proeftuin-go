# Template Method -- A high-level blueprint for an algorithm to be completed by inheritors

## Overview

Now we're going to take a look at the Template Method design powder which is very similar to the pattern that we've met before which is the strategy pattern.

So we already talked about this idea that algorithms can be decomposed into the common parts and the specifics.

So a Strategy Pattern typically does this through composition.

So you have a high level algorithm which uses some interface and then you have a concrete implementation which implements this interface and then you keep some sort of pointed to that implementation and you you have a pointer of the interface type you provide the concrete implementations and then you use the whole thing.

So the Template Method is actually very similar.

It's very similar because in go there is no inheritance if we had inheritance then those two would be completely different.

But we don't.

So the Template Method is well typically it's a function.

So it doesn't have to be a struct with a reference to the implementation it can be just just a freestanding function which takes that implementation as a parameter we still use interfaces just like in strategy simply because there is no other way really for us to generalize this idea of implementation details.

But there is an alternative.
So we're going to take a look at both of those.

So one of those is when you have a function which takes some interface which makes it very similar to strategy but there is another alternative when you have a function which takes a bunch of functions as parameters.

And in this case you can do without struct or interfaces you can just work with functions and nothing but function.

So this is the functional Template Method approach.

So a template method is once again a skeleton algorithm defined in a function.

And these functions can either use an interface which makes them kind of equivalent to a strategy or they can simply be functional so they can be a higher order function that takes several functions as arguments.

## Template Method

All right.

Let's take a look at how to implement template methods.

So imagine that you're simulating different kinds of games.

Now these games all have very similar structure.

So if you think about games like chess or checkers or a card game then they're all pretty much the same in the sense that you have a bunch of players and each of the players takes their turn.

So one player goes then another player goes and so on and so forth they're all of course variations.

But let's imagine a simple scenario like a game of chess or checkers for example you simply have two players or maybe more than two players.

And every single player takes a turn one after another.

So what we can do is we can formalize this process using a template method.

But in order to do this we also need some sort of interface for what parts of the game we're interested in.

So here I'm going to define game as an interface and this interface is going to have a bunch of methods so first of all we'll have a method called start.

Now this method is going to be invoked whenever the game actually starts.

Now in addition we'll have a method for taking a turn.

So that is what happens when one of the players takes a turn.

So for example in the game of chess they might move the piece.

Then we'll also have a method for figuring out whether or not the game is over whether or not we do in fact have a winning player.

So we'll have a method called have winner which is going to return a billion value.

And in addition we need to determine the number of the winner like which of the players actually won the game.

So we'll have a method called winning player which is going to return in the end.

So we have this interface what we can do is we can write a template method which is simply a skeleton algorithm which makes use of this game interface.

So for playing a game we're going to have a function called Play game which takes G which is a game.

So the idea here is that we simply use the interface and invoke the interface members in the exact order in which we want to define the algorithm.

So to play a game first of all you have to start the game.

So we called you the start and then for as long as we don't have a winner somebody has to take a turn.

So we say for not g have winner we take a turn.
So we say Gee I'll take turn.

So we take turns and once we have a winner we might want to printout the winner's name.

So here I would do a simple print half where we would say player percent the wins like so and here I would specify the winning player once again using that game interface.

Okay.

So now that we have this what we can do is we can make an actual game we can make a game of chess for example and of course here I'm not going to implement all of the chess rules and whatever I just want to show you a simulation of how the game can proceed under this template method paradigm.

So we're going to have a tie called chess.

Notice I'm making everything lowercase I'm just going to return the game interface when I make the factory function.

So we'll have a game of chess and we'll have a bunch of just inner variables that we want to know about like whose turn it is the maximum number of turns.

Maximum in the sense that we're going to simulate a game rather than really play this game.

We'll have an index of the current player and all of these are going to be integers and then what we can do is we can have a factory function.

So we'll have a function called new game of chess.
Notice it's not going to return.

Chess is going to return game that way you can sort of hide all the internals of the chess struct that we've created up above.

And here we just return this chess struct with a bunch of parameters.

So the game is going to start on turn 1 there's gonna be a total of 10 turns that we're going to simulate and the current player will have index zero.

So we'll have two players with indices 0 and 1.

So this is how you can set up the whole thing.

And now that we have this whole thing we can continue and actually do something about it.

All right.

So there's just one thing missing here in the chess struck that and that is the implementation of the entire game interface because remember chess is a game so we're going to go ahead and implement every single one of these methods and we're just going to simulate the whole process.

So for example when the game starts we can just say that's that we're starting a new game of chess.

Like so when it comes to taking a turn we can increment the turn.

So we can say See that plus plus.

And then we can for example output information telling us that a turn number so and so was taken by a particular player.

So here I can do a print half.
So here I would say 10 percent D taken by player.
Percent The.
And the certain number is see that turn.

And the player number is see the current player and then we can increment current player so see the current player equals and we don't increment that we cycle it.

So the players take turns.

So it's player 0 then player 1 then player 0 again and so on and so forth.

So we could just say something like see the current player equals one minus see the current player.

That way you alternate between zero and one forever.

Okay.

So finally let's implement have winner and winning players so we're going to say that the winning player is after we simulate the 10 or so steps the winning player is the current player let's just return see the current player and the real world of course you would examine the actual conditions of the game and determine which player has won.

And finally when it comes to having a winner we're just going to say that we have a winner when we've gone through every single turns out here.

I'm going to return.
See that equals a seat.

Max Jones once again this is all a simulation.

This is not a real game of chess.
I'm just showing you how the template method can be used.
OK.

So now that we have all of the pieces what we can do is we can make a new game of chess.

So here is a new game of chess and we can actually play the game so we can invoke the template method passing in the chess object.

So let's actually run this let's see if it compiles and runs.

All right.

So as you can see you were calling the start method first of also starting a new game of chess and then we simulate the turns taken by the different players.

So it's play a zero play a one play zero player one and so on until the end and at the end we have player 1 play 0 taking a turn and then player 1 is supposed to take a turn and player 1 wins.

So that is how you set up the template method.

So essentially the template method is a skeleton algorithm so you can see that here we're using the abstract members in a way the members of an interface which you don't have a concrete implementation until you actually implement the game interface in some struct of yours and then you pass that struct into the play game function and you actually use it.

You actually use that high level algorithm with the you know definitions of the different members being defined in whoever actually implemented this interface.

### Template Method code: behavioral.templatemethod.templatemethod.go

```go
package templatemethod

import "fmt"

type Game interface {
  Start()
  HaveWinner() bool
  TakeTurn()
  WinningPlayer() int
}

func PlayGame(g Game) {
  g.Start()
  for ;!g.HaveWinner(); {
    g.TakeTurn()
  }
  fmt.Printf("Player %d wins.\n", g.WinningPlayer())
}

type chess struct {
  turn, maxTurns, currentPlayer int
}

func NewGameOfChess() Game {
  return &chess{ 1, 10, 0 }
}

func (c *chess) Start() {
  fmt.Println("Starting a game of chess.")
}

func (c *chess) HaveWinner() bool {
  return c.turn == c.maxTurns
}

func (c *chess) TakeTurn() {
  c.turn++
  fmt.Printf("Turn %d taken by player %d\n",
    c.turn, c.currentPlayer)
  c.currentPlayer = (c.currentPlayer + 1) % 2
}

func (c *chess) WinningPlayer() int {
  return c.currentPlayer
}

func main() {
  chess := NewGameOfChess()
  PlayGame(chess)
}
```

## Functional Template Method

So now that we looked at a structural implementation of the template method let's take a look at an alternative.

Let's take a look at the functional approach now in the functional approach we're they're going to have any interfaces or Amy's these trucks.

Instead what we're going to do is we're going to make sure that the template method operates simply on functions.

So here's what they can look like.

Let's have a function called Play game kind of similar to the play game function we had in the previous demo except this time round we're going to take a bunch of arguments so instead of taking an interface we're going to unwrap that interface and we're going to take a bunch of functions as parameters.

So we're going to have two functions start and take turn.

And these are going to be functions which take no arguments and don't return any values.

Then we're going to have a function called have winner that's going to be a function that returns a boolean value and then we'll have a function called winning player which returns an integer and then the actual implementation of the template method is going to be very similar to what we had in the previous example.

So first of what we call a start and then while we don't have a winner we take a turn.

So while we don't have a winner we take turns like so.

And then finally we can do the same output has before.

Let's just do a print f where we say player the winds and of course to get the actual player we use winning player.

So that's another function that we call.

So the idea here is that instead of taking an interface which has all of these methods what we do is we take these functions as arguments and as a result we can use this template method inside the main function of ours without really defining any interfaces or struts or anything like that.

So I can do everything right inside main So for example if I need a bunch of variables for storing the current turn the maximum number of turns the current player I can just have them as members right here so I can have them as variables so we'll have turn Max turns and current player is gonna be equal to 1 10 and 0 and then I can define the actual component function so for example for starting the game I can just define a function right here and this function is going to just say starting a game of chess.

And similarly we can define all those functions that we previously had inside a struct.

So let me just paste those in here so you can see that here I have a function for taking it turn and then once again this is just a function doesn't take any arguments and it incremented turn does pretty much the same thing as the Premier's demo did and the same goes for have winner and the winning player.

So there is no reference to any kind of struct here because these variables like turn and Max Jones for example they are defined up above they are defined right here so now that we've put everything together what we can do is we can use the template method so we can call play game and we can pass in those functions we can pass and start take turn have winner and winning player okay.

And if we run this well let's take a look at what we get as you can see you were getting pretty much the same output as before I was starting a game we're taking the turns and then we determine who is the winner.

So the takeaway from this demo is that you don't necessarily have to deal with interfaces and strikes if you don't want to instead what you can do is you can define a template method not as something which uses an interface full of functions but actually something which takes functions as arguments and then uses those functions to define the skeleton of some method and then of course once you've defined that skeleton what you do is you simply fill in the gaps so you create every single one of these functions right here and then you pass these functions into the template method and led the template method use them to actually implement the details of your algorithm.

### Functional Template Method code: behavioral.templatemethod.functionaltemplatemethod.go

```go
package templatemethod

import "fmt"

func PlayGame(start, takeTurn func(),
  haveWinner func()bool,
  winningPlayer func()int) {
  start()
  for ;!haveWinner(); {
    takeTurn()
  }
  fmt.Printf("Player %d wins.\n", winningPlayer())
}

func main() {
  turn, maxTurns, currentPlayer := 1, 10, 0

  start := func() {
    fmt.Println("Starting a game of chess.")
  }

  takeTurn := func() {
    turn++
    fmt.Printf("Turn %d taken by player %d\n",
      turn, currentPlayer)
    currentPlayer = (currentPlayer + 1) % 2
  }

  haveWinner := func()bool {
    return turn == maxTurns
  }

  winningPlayer := func()int {
    return currentPlayer
  }

  PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
```

## Summary

Let's summarize the things that we've learned about the Template Method design pattern.

So it's actually very similar to the strategy pattern and the typical implementation is that once again you define an interface with common operations and then you make use of those operations inside a function.

But there is an alternative functional approach that's where you make a function that takes several functions.

This is what we typically call a higher order function and then you can pass in functions that capture the local state.

So for example you can invoke the Template Method right from the main method for example because if you have any variables inside the main method they will just be captured by the functions you are passing in.

And what this means.

This functional approach is that there is no need for structure or interfaces you can just work with functions and you can get very similar functionality although maybe not as readable maybe not as easy to understand but it's still a viable approach that you're welcome to take.
