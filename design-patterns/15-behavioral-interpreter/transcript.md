# Interpreter -- Interpreters are all arount us. Even now, in this very room

## Overview

If you are a software engineer you're working with the interpret that design pattern every day.

Like all the time literally because it's used literally everywhere.

When it comes to the Structure and Interpretation of Computer Programs This is one of a number of recommended books if you're really interested in topic but essentially the idea is that the interpreter design pattern handles the situation where you need to process textual input and there are so many different situations where that is required like you want to get some text you want to turn it into some sort of data structure and typically what we call this data structure and especially if you are passing let's say programming languages you turn it into something called an abstract syntax tree and that tree can subsequently be traversed using for example the visitor pattern and then you can you can do things with it you can transform it you can compile it and so on and so forth.

So let me give you some examples of where textual input actually has to be processed somehow.

So the simple example is that the compilers and the interpreters of the languages that we use they are in the business of interpreting data.

Also the ideas as well.

So behind the scenes every single idea like go land for example would have a parser in there laxer for actually understanding what you type and then they perform analysis on top of that.

We have lots of specialized text formats like HDMI like somehow Jason and so on and those need to be interpreted somehow we have simple math like numeric expressions that you want to evaluate.

Remember you kind of just have the program before the string and understand that string that string has to be turned into an abstract syntax tree and then traverse accordingly.

A more sophisticated example is regular expressions that is essentially what we call a domain specific language so a language specific for solving one particular problem.

In our case the regular expressions solve a particular problem of whether or not a particular string matches a particular pattern.

So this is kind of like a language within a language and turning strings into these linked structures is a sufficiently complicated process.

In the case of programming languages it gets particularly difficult.

So when you're going to look at any super complicated examples about even the simple examples get somewhat non-trivial shall we say.

So the interpreter is some component that processes structure text data and it typically does it in two stages.

So first of all it takes the text and it turns it into a separate lexical tokens.

That's the process called lexicon and then the second part is a you take the sequence of tokens and you turn them into some sort of linked structure some sort of structure of an abstract syntax tree.

And this process is called passing.

## Lexing

All right.

So what we're going to do in our demo of the interpreter pattern is we're going to pass a simple numeric expression and then try to evaluate this.

And this is going to be split into two parts so the first part is the lexicon.

That's when you take the inputs and you split it up into separate tokens.

So let's define the input.

I'm just going to have a simple mathematical expressions let's say 13 plus four minus twelve plus one.

There we go.

And then what we want to be able to do is we want to be able to get a bunch of tokens.

So we want to get a bunch of tokens as a result of some lexical process so we pass the inputs into the Lexis process and we get our tokens.

Now the question while the first question is What are those tokens Exactly.

So where where do we get them.
Let's actually define a token type.

So it's going to be just an integer here and we'll have just a limited number of tokens so if we're looking at our expression here the kind of tokens we have are opening and closing parentheses we have pluses and minuses and we also have numbers.

So that's what we're going to stick to.
So here I'll just define a bunch of content.

So we'll have an integer and that's going to be token type.

We'll also have plus minus how left parentheses and right parenthesis.

So that's what we're going to have.
And then we can define the actual token.

So a token is going to have two things it's going to have the token type but in addition it will also have the actual text of that token because obviously when you pass a token such as the number 13 you don't want to lose information about the fact that it's 13 because if you were to just draw the end you would lose the information about what what what is the actual number so to speak.

So let me also define this Stringer interface on the token so the string your interface on the token.

I just want to be able to output tokens surrounded by backwards so you get to see what's what's actually going on there.

So if empty a dot aspirins have so percent yes.

And then here we we add the actual text of the token like so.

OK.

So now that we have this let's actually define the Lessing process.

So we're gonna have a function called Lexx which takes an input.

That's our string and it just returns a bunch of tokens.

Okay.

So let's make the result and then we're going to append to this result.

So I'm not going to be using a range construct here instead I'm going to be using an ordinary full loop because we're going to mess with the indices somewhat.

So we're going to go from zero to the set of tokens.

So so from zero to input length like so I plus plus they there go.

So we're going to take a look at each of the tokens in turn.

So I'm going to look at inputs at position i and depending on that input I'm going to perform certain actions.

So for example if I get a plus that means I have to append he plus token with the text being plus so I can see a result and I can append a token where the token type is plus and the text is also a plus.

Okay.

So now this is going to be quite a common theme here because we're going to do this four plus four minus

four opening and closing parentheses.

So I'm just going to do this quickly like so so we'll have minus here and minus here.

We'll have the left parenthesis.

So that's going to be health per n and left parentheses here and right parentheses here.

I'm got to be careful that the idea doesn't add additional closing parentheses or what not because it likes to do that.

So these are the four obvious cases the only case which is not obvious is the case of numbers.

Now why is it the case.

Well it's the case because a number can take up more than one character.

So 13 for example is composed of two digits which means we have to write special code in order to process this thing.

So here in the default case I'm going to assume that I got a number and I'm going to have to make a strings build their self strings dot builder.

Builder like so can we make a variable let's call it as beam.

So this string builder is going to be used to add all the digits together before we turn them into an actual number.

So I'll have a new counter J which is going to start at position i and j is going to go until the length of the input so until the very end.

J plus plus never.

So what we are looking at is we're looking at the digit at position J and we're saying what is this digit a is this in fact a digit.

If it's a digit we add it to the string builder and increment i.

If it's not then then we do something else.

So here I'm going to use if Unicode dot is digit so I'm going to convert inputs at position J to a rune.

And if it happens to be a digit then yeah we can write that rune right rune.

So once again here I'm doing some naughty things with simply converting the inputs to a room which is going to be just fine.

In our case so I'm going to write this into the string builder and I'm going to increment the index.

Now we'll we'll need to decrement this whenever we're done so in the LS case in the LS case we are now ready to actually set the result.

So we need to append to the result a new token but this token is going to be an integer.

So the type is end and the text of the token is whatever we have in the string builder.

So as beat out string.
So we add this token and then we decrement.

Because we incremented it too much in the process of going this.

And then we break b because we are effectively done with this whole thing.

So that is the legacy process and towards the end of the lesson process we return the result which is the array of tokens that we got.

And let's actually take a look at what we get as we kind of as we get the whole thing hopefully F.A. dots print line will actually take a slice and print it correctly so we'll print all the tokens right here.

Let's actually run this let's see if this expression gets interpreted correctly.

So for each of these U.S. pairs where the left side is the enum and the more rated thing so so ignore that.

So opening brackets then the number 13 plus for closing bracket minus opening bracket.

OK.
So everything works.

So the lexical process works the Lessing process and that is part one of our process of making an interpreter.

### Lexing code: behavioral.interpreter.handmade.go

```go

package main

import (
  "fmt"
  "strconv"
  "strings"
  "unicode"
)

type TokenType int

const (
  Int TokenType = iota
  Plus
  Minus
  Lparen
  Rparen
)

type Token struct {
  Type TokenType
  Text string
}

func (t *Token) String() string {
  return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
  var result []Token

  // not using range here
  for i := 0; i < len(input); i++ {
    switch input[i] {
    case '+':
      result = append(result, Token{Plus, "+"})
    case '-':
      result = append(result, Token{Minus, "-"})
    case '(':
      result = append(result, Token{Lparen, "("})
    case ')':
      result = append(result, Token{Rparen, ")"})
    default:
      sb := strings.Builder{}
      for j := i; j < len(input); j++ {
        if unicode.IsDigit(rune(input[j])) {
          sb.WriteRune(rune(input[j]))
          i++
        } else {
          result = append(result, Token{
            Int, sb.String() })
          i--
          break
        }
      }
    }
  }
  return result
}

func main() {
  input := "(13+4)-(12+1)"
  tokens := Lex(input)
  fmt.Println(tokens)
}
```

## Parsing

Part two of making an interpreter is basically taking these tokens and passing them turning them into more sophisticated structures because more sophisticated structures can be traversed in a recursive fashion and we can actually evaluate the numeric value of this expression.

So in order to do this we're going to introduce a bunch of trucks and interfaces so let's go up here and let's introduce those.

So we'll have a time code element which is going to be an interface and this interface is just going to have a single method for returning the value of something.

So it's going to return the value of either a simple construct like a number or a complicated construct like a binary expression.

Now why why binary.

Well because all that we have in our model is we have pluses and minuses and a plus and a minus both take to operate and so they are binary operations.

Okay but before we do that let's make a type called integer.

So this is going to be the primitive for every single integer token that we have.

So an integer is just going to have a value which is going to be an int.

Let's have a constructive for it just for the fun of it.

And what we need to be able to do is we need to be able to implement the element interface on integer so implementing the element interface here we simply return idle value.

And there is no magic here we simply return the value that is contained in the integer.

Now the interesting part is of course the binary operation.

Now there are two different operations.
There's addition and subtraction.
So we may as well define those as constants.

So I'll do it like I do it always type operation INT.

And then I'll have a constant block where we'll have the addition and subtraction.

Okay.

So now that we have this we can make a struct called binary operation binary operation.

Now a binary operation is of course operating on two different elements.

So we'll have the type of the operation and we'll have the left and right elements like so.

All right.

So now that we have this we want to be able to evaluate the value of the binary operation.

So we want to implement the element interface once again and get the value of the operation.

So this obviously depends on whether it's additional subtraction.

So we're going to switch on beetle type and take a look at the two cases.

So in the case where it's addition we can say that we return v the left dot value.

So we reuse that interface plus be the right Dot value.

Now we go and in the case of subtraction we can just copy and paste this.

So in the case of subtraction what we do is we put a minus here that's all there is to it then we can also have some default case just in case somebody entered something else.

We can panic here saying this is unsupported operation.

Now we go.
OK.

So now that we have all of this set up what we need to do is we need to write a new function called pass and this function is going to turn this set of tokens into a top level binary operation we're going to assume that every expression is a binary operation here it is binary because it's adding subtracting from this side and this is what's being subtracted.

So it's a subtraction operation at the top level.

So let's write a function code parse so parse takes a bunch of tokens and it returns an element.

So in our case is going to return a binary operation.

I'm just going to make this assumption for simple simplicity purposes.

So result is going to be a binary operation.

Now we go and when we come to fill in every single one of those operations that we pass we need to be able to know whether something we've passed has to go on the left hand side of the tree so to speak because a binary operation has a left hand side and right hand side.

So we're going to have a flag here called have l H.S. false.

So if we have the left hand side already it has to go to the right hand side.

Otherwise we just put it to the left hand side.

So once again we go through every single token and once again I'll just use use counter here so zero to less than length of tokens I plus plus.

Okay.

So I'm going to get the token at position ISO token is going to be.

Well let's get a pointer tokens at position i.
And then I'm going to take a look at what the token type is.
So switch token dot type logo.

So let's imagine that the token we got is an integer.

What we need to be able to do is we need to be able to convert that integer into a actual number because remember tokens just store Strings they still text but we need an actual number for calculation purposes.

So that number is going to be N and I'm going to assume that the operation will always succeed.

So I will not have an Hey flag here.

So we'll use SDR corn dots a two y to take token dot text and basically convert it to an integer and then we need to construct the integer object remember we made this whole big struct so integer as a structure is going to be integer initialized with value then.

And now we need to know where to put it.

So if we don't have the left hand side yet if we don't have the left hand side yet then result got left is this integer.

Notice it's a pointer here and we say have a play chess equals true otherwise result got right is appointed to that integer.

Okay so that's integer stuff done but that's not the only case that we have.

So so the different cases different types out there so you could have a plus for example if you encounter a plus that means the current operation is addition.

So we say result type equals Ed..
And similarly for the minus.

So if you get it minus then result type is is subtraction.
Like so.

And then of course we need to deal with the most complicated part of all the left and right parentheses.

Now what we're going to do is we're going to do the following when we encounter a left parentheses.

We're going to find the location where we encounter the right parentheses and then we're going to take everything in between the left and right parentheses and we're going to feed it recursively into the pass method thereby getting an element and that element is going to be stored.

So that's the idea.
Okay.

So when we encounter the left parentheses here's what we do.

So I'll make a new variable j which is going to be initialized with the value of AI and then we'll make a loop where a J goes until the length of tokens.

And so as soon as we encounter the right parentheses with we stop and that explains why the J variable is defined outside this loop because we want to use it later on so if tokens at position J has the type of right parentheses that means we're done and we can just break and then now J has the value of the location of the right parentheses.

So now that we have this what we can do is we can try grabbing the sub expression which is enclosed by parentheses because we have the opening and closing parentheses.

So let's make some expression like so token slice.
And then for.

Starting at position i plus one going all the way up until J.

But not including J obviously so K plus plus we append to the sub expressions so sub expression dot append tokens at position K another we've done this we have this sub expression we can pass it recursively so we can say element is equal to class sub expression now that we've passed sub expression we have the same thing as before with checking out whether or not it has to go in the left hand side though the right hand side so if we don't have the left hand side then resolved that left is equal to elements and have l H.S. plus true otherwise resolved all right equals true or result the right equals element rather.

Okay now towards the end of it all what we need to be able to do is we need to scrawl the value of VI to the point J because we've passed the whole subject impression so we say I equals J and the next element the next iteration of the loop will start at position j plus 1 which is exactly what we want in the first place.

So towards the end of it towards the end when we've done the whole thing we just return a pointer to the result because it's an interface type that we need to be returning and now we can use this parsing mechanism to actually pass the tokens and also evaluate them also calculate their value because remember we've already implemented the calculations.

Let me just jump up here.

So in the case when an integer is encountered you simply calculate its value.

We just store that internally in the case of a binary operation of the value calculation here performs the addition or subtraction depending on what was actually specified.

So we are they have the calculation mechanism is done which means we can put everything together so let's have past past being the invocation of pass on the tokens I seem to be missing a closing bracket here.

And once we've passed the whole thing we can we can print line both the input as well as the past value.

So here I can say percent has equals percent s so the first element is going to be the input that we try to evaluate and the second is going to be past dot value.

Now we go.
Get.
So let's see now actually these.

The second one is D because it's it's and no.

But but the first one is a string so everything should be everything should be okay here.

Actually that should be printed off.
Sorry about that.

That's some more correct way of doing things so let's let's run this let's see if we do in fact get what we want.

Okay.

So 13 plus four is seventeen and twelve plus one this 13 so 17 minus 13 is equal to four.

So we are getting there right now.
Both right here.
Okay.

So this has been a demonstration of how to actually implement in the interpreter pattern.

So the idea is that typically you split it into two parts.

There's the lexical part where you take your textual input and you set it up a you split it up into a bunch of tokens and then you take those tokens and you create more complicated tree like structures out of those tokens.

And those tree like structures can for example be traversed in all sorts of ways.

They can have their values evaluated.
They are easy to print.

So that is the end result that we want as a consequence of the interpretation process.

### Parsing code: behavioral.chainofresponsibility.brokerchain.go

```go
package main

import (
  "fmt"
  "strconv"
  "strings"
  "unicode"
)

type Element interface {
  Value() int
}

type Integer struct {
  value int
}

func NewInteger(value int) *Integer {
  return &Integer{value: value}
}

func (i *Integer) Value() int {
  return i.value
}

type Operation int

const (
  Addition Operation = iota
  Subtraction
)

type BinaryOperation struct {
  Type Operation
  Left, Right Element
}

func (b *BinaryOperation) Value() int {
  switch b.Type {
  case Addition:
    return b.Left.Value() + b.Right.Value()
  case Subtraction:
    return b.Left.Value() + b.Right.Value()
  default:
    panic("Unsupported operation")
  }
}

type TokenType int

const (
  Int TokenType = iota
  Plus
  Minus
  Lparen
  Rparen
)

type Token struct {
  Type TokenType
  Text string
}

func (t *Token) String() string {
  return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
  var result []Token

  // not using range here
  for i := 0; i < len(input); i++ {
    switch input[i] {
    case '+':
      result = append(result, Token{Plus, "+"})
    case '-':
      result = append(result, Token{Minus, "-"})
    case '(':
      result = append(result, Token{Lparen, "("})
    case ')':
      result = append(result, Token{Rparen, ")"})
    default:
      sb := strings.Builder{}
      for j := i; j < len(input); j++ {
        if unicode.IsDigit(rune(input[j])) {
          sb.WriteRune(rune(input[j]))
          i++
        } else {
          result = append(result, Token{
            Int, sb.String() })
          i--
          break
        }
      }
    }
  }
  return result
}

func Parse(tokens []Token) Element {
  result := BinaryOperation{}
  haveLhs := false
  for i := 0; i < len(tokens); i++ {
    token := &tokens[i]
    switch token.Type {
    case Int:
      n, _ := strconv.Atoi(token.Text)
      integer := Integer{n}
      if !haveLhs {
        result.Left = &integer
        haveLhs = true
      } else {
        result.Right = &integer
      }
    case Plus:
      result.Type = Addition
    case Minus:
      result.Type = Subtraction
    case Lparen:
      j := i
      for ; j < len(tokens); j++ {
        if tokens[j].Type == Rparen {
          break
        }
      }
      // now j points to closing bracket, so
      // process subexpression without opening
      var subexp []Token
      for k := i+1; k < j; k++ {
        subexp = append(subexp, tokens[k])
      }
      element := Parse(subexp)
      if !haveLhs {
        result.Left = element
        haveLhs = true
      } else {
        result.Right = element
      }
      i = j
    }
  }
  return &result
}

func main() {
  input := "(13+4)-(12+1)"
  tokens := Lex(input)
  fmt.Println(tokens)

  parsed := Parse(tokens)
  fmt.Printf("%s = %d\n",
    input, parsed.Value())
}
```

## Summary

All right.

Let's summarize what we've learned about the interpreter design pattern.

So apart from the really simple cases an interpreter typically acts in two stages.

So the first stage is when you take the input and you turn the text into a set of tokens so if you have something like three times four plus five you would take each constituent self-contained parts so you will take the three and that will become a literal he would take the star you would take the opening bracket and so on and so forth so you would split it into a sequence of tokens and then the second stage which is the passing is when you turn these tokens into NASD or you turned it into you turn it into some sort of aggregated structure or a structure with links to other structures so that the whole thing can subsequently be traversed because now that you have the pass data you can apply the visitor pattern and you can traverse this and for example evaluate the numeric expression or print the expression to the screen.

Stuff like that.

So this is the gist of the interpreter design pattern.
