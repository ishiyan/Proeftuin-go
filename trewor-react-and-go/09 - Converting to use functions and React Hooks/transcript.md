# 09 - Converting to use functions and React Hooks

## 107 - About this section

In this section of the course, we're going to start working with functions instead of classes, pretty

much everything we've built in our projects so far uses react components, which are classes.

So as you can see here in my editor on the right hand pane, I have the edit movie component open and

it exports a default class, which extends component, and that's a class.

So that means we have access to a constructor, to the state, to all those things that you don't appear

to have access to in regular JavaScript functions.

But with the use of react hooks, we have access to all of those things.

So the two most important hooks we're going to be using are use effect and use state.

I'll be going over those shortly.

But as you can see on the left hand side of my editor, I have another function called edit movie func,

which duplicates exactly the functionality that we have over here in the Riak component.

It duplicates it using a standard JavaScript function.

So as you can see, we still have our import statements and we're we're importing almost the same things,

but we're not importing component like we are over here.

Instead, we're importing something called use effect and something called use state.

And one of the great advantages of working with react using functions instead of components is usually

wind up with less code.

For example, over here in my component version of added movie I have, this component did mount function,

which starts on line forty three and ends all the way down on line eighty nine.

And the same functionality is achieved using the use effect which starts on line twenty four in my function

based version and ends on line forty nine.

And that's a lot less code, which means a lot less opportunity for mistakes and a lot less code to

maintain.

Now it's important to understand how to use react components as classes for the simple fact that there

are millions of lines of code out there in production right now that use them.

And there are also many, many organizations that still do most of the reactor work using classes because

their developers are used to it and it works really well.

That means that there are many employers looking for people with that skill set.

And having said that, it's a useful exercise to convert our existing project over to functions and

hooks.

And that's what we're going to do in this section of the course.

Now, I have an application started already, and it's just to show you a couple of things.

So let me show you my Web browser.

What we have here.

We have a simple react app and I just created using NPCs Create React app and gave it some name and

it just displays one title hook test.

So let's go back and look at the code for that right here.

So this is the entire application.

It is one statement of one function called app, and all it does is return some X.

Now, if I wanted to use state in this, how would I do that?

Well, it's pretty straightforward.

You would come up here inside your function and it has to be inside your function because state is still

limited to the current function and I'll declare it constant and I'm going to call my state variable

favorite color.

But the syntax to create ID is to use square brackets and give it your name.

So favorite color, then a comma, then the name of the function used to set the value of that state

variable.

And the convention is to start it with the words set and then use your variable name, favorite color

and outside of the square brackets, an equal sign.

And then you use the react hook called use state.

And that takes an optional parameter, which is the default value.

So I'll set it to now.

As you can see here, I have an error because I have an import of that.

So I'll come up here and say I'm here to import react and you don't actually need to import react.

But it's good practice to do so.

At least I think so.

And then I'll import the hook use state.

Now, this is available to me so I can come down here after my title in my jail sex and maybe put an

H.R. in nature, and then I'll say in a paragraph, your favorite color is and then I just use the local

variable, which is called favorite color.

Which is pretty straightforward.

Now, if I go back and look at my Web browser, it says my favorite color is red.

So how do I change the value of that state variable?

Well, that's where this function set favorite color comes in.

So what I'll do now is put another paragraph tag in and I'll add a button button and I'll give it a

class name of, say, button button primary from bootstrap and I'll give it an unclick handler.

And that's going to be equal to and I have to use this syntax or it won't work and I'll just call set

favorite color and this time I'll say it's going to be green.

And close up and say change color.

Now, if I save this and go back to my Web browser and it now says your favorite color is red, and

if I click on change, color is change it the green.

So that's how I manage state.

And it's very straightforward.

Now the syntax is a little bit different and it takes some getting used to.

And to be honest with you, it feels a bit like cheating because this seems to be so much simpler than

using a class based approach.

And it is.

But the fact that you've gone through and already figure out how to work with classes should make this

ever so much more meaningful to you because you know exactly what's going on behind the scenes.

So that's the first one we need to work with now.

There's actually a second hook we'll be dealing with quite regularly as we go through the remainder

of this course.

And it's called use effect.

So I'm going to it use effect and use effect, despite its name, actually has a lot in common with

functions in classes like Component did mount, for example.

In other words, it will run the very first time and every time this component loads and is rendered

on the screen and we use it like this so we can say use effect and it's just a function.

But it uses the Arrow syntax, which you're familiar with by now.

And what I'll do here is duplicate this constant and change this one to favorite number and I'll give

its function to set its value set.

Favorite number.

That's what I'll call it.

And I'll give it a default value of, say, 10 like that, OK?

Now, in use effect, even though I've given this a default value of ten, I can actually overwrite

it every time the component looks.

So I can say something like set favorite number two eleven.

Now, that sort of works or looks like it should work, but you're not.

So when I did this, I got a little warning overuse effect.

And that says the warning is react hook use effect contains a call to set favorite number without a

list of dependencies.

This can lead to an infinite chain of updates.

And that's something you really don't want to have happen.

And we get around that by simply passing favorite number to use effect, because favorite number is

associated with the function set favorite number and use effect actually has an optional parameter,

optional second parameter, which comes in the form of an array.

And all I have to do in this case is to pass your favorite number to it.

OK, so now I've got my state variable favorite number, which defaults to a value of ten, and I've

passed this favorite number to use effect and told it to whenever this component renders set favorite

number two eleven.

So I can come down here now in my ex and say comma and your favorite number is favorite number.

So remember that I'm defaulting the value of favorite number to 10, and then I'm using the hook use

effect to set it to 11.

So let's see what we get.

We'll go back here and it says your favorite number is 11, so you can see the use effect hook ramp.

Now, again, these are really simplistic examples, but over the course of the remainder of this section,

we'll be using both the U.S. state and the use effect to a great deal to duplicate the functionality

that we have in our classes.

So let's get started.

## 108 - Converting the Movies.js component to a function with hooks

So in this section of the course, we're going to look at moving away from using react components as

classes into using functions and hooks, which apparently is going to be the way that react moves forward.

Up until very recently, the documentation stressed knowing how to use components in classes and so

forth, and said use hooks only in certain situations where it's appropriate.

But the documentation is actually changed and it's become clear that the future for react is going to

be using functional components and hooks.

So what we're going to do in this section of the course is convert our project to use function and hooks.

And we're going to start with this one with movies, Dot Jass.

So this is what our file looks like right now.

It is a standard Riak component and we use the component did malfunction, so forth.

And what I'm going to do is create a new file in my components folder, which I'll call movies and I'll

just call it functions.

So it's a movie function.

And just to make things a little simpler, what I'm going to do is split my editor.

So I'll view the editor layout to be two columns that I'll move movies over this column and then apparently

close that there.

So now we have an empty movie function file and I'm going to rewrite my react component movies just

to be a function instead.

So we definitely need to import react.

Let's just go if I am or in part react.

And I know I'm also going to be using a few things from react.

I'm going to be using use effect, which is a hook used state, which is the other half we're going

to use.

And those are far and away the most common hooks you will ever use.

And I'm also going to use fragment.

And the other thing I know I'm going to use is I'm going to be import importing link from react rubber

door right there.

OK, now to make this function, we just declare a standard JavaScript function and I'll call it the

movies func and it will take the argument of props because we might be getting properties from whatever

calls this.

And inside of this, we're going to return something.

So for the sake of argument right now, I'll just say return.

And in parentheses, I'll put H to movies or choose a movie.

That's what I call it.

The other one, choose a movie.

OK, and now of course, I need to export default and I'm going to export movies for now.

Let's go back to abduct Jass.

Which is right here, an import that so import the one we just created, Louise Funk, and down here

where we link to movies instead of going to movies, will go to movies.

OK, now if I have my application running, apparently I do.

This should now work so I can go to movies and it gives us our empty or very simple functional component.

That's what's been returned.

So let's go back to our code and now let's start modifying our code in movies funk to actually do the

things we needed to do.

Now I'm going to be using use effect and what use effect does for the most part, at least for our purposes

right now, is to take the place of component mount.

So what I'm going to do is declare two constants const and the first one is going to be R movies.

So what we're going to store the list of movies we get from the back end server M and that actually

has to be constructed in an unusual fashion, the variables called movies.

And to change the value of this, I'm going to use something called set movies and this can be whatever

I want it to be.

This can be called whatever I want it to be.

But set movies seems logical and that's going to be using the state hook.

And by default, I want movies to have no value at all.

So I'll just put an empty array and I'm also going to declare a second constant and this one will be

for errors in case there's an error or we call the back end.

And again, how do I change the value of that with a function which I'll call set error, which seems

to be illogical.

And I want that to have a default value.

You state of let's you signal OK, or we can make it an empty string string there.

So now I have declared those two constants.

Now I'm going to call the back end.

And over here in movie Dargis, we use the fetch statement and component did mount.

But here in the functional version of this, I'm going to use the the hook use effect and we use the

Arrow syntax to define what happens when this is called.

And what I'm going to do is just copy this entire fetch statement and modify it as necessary.

So I copy this out of movies, dojos and paste it right in here.

Now, this Fitch statement will be exactly the same then response.

Well, I'm actually not going to be able to use this statement, so let's comment this part out.

And instead, if we have an invalid response, if our response is not exactly equal to and in this case,

it's going to be the number two hundred probably should have been in the other one as well.

We'll call Sarah and we'll set her to invalid response code and then its response status or status,

if you prefer.

OK, so clearly that's less code.

Three lines here.

We're down to one line.

Now, if there is no error, I'm going to say else I'm going to call Sarah and make it to.

And the reason I'm doing that is when I make this call, I might get an error.

Something system is busy, resources, too many resources being used.

I'm getting something other than two hundred.

But then the next time it calls it, I need to set error back to normal.

So I do that and then we have return response.

Jason, same as before.

Now in this case, when we get a valid response, we obviously can't use this dot set state because

it's not available to us in a function.

So instead, all I have to do is say set movies and we're going to set up to chase and dot movies again,

a lot less code.

So that's a good thing.

Now, there's one other thing we need to do here.

There's an optional callback unus effect.

And what that optional callback says is what's the default value?

And I want that to be an empty array.

And I'll show you why we have to set that before too long.

So I'm just going to put a comment here.

Must set default value now down here in the return, we can actually use the same logic we did before.

You can see that we have this kind of logic.

And what I'm going to do is just put in a statement here.

I'll say F error is not exactly equal to no.

Then do something else, do something else, and I'll delete this and put this down here and then I'll

copy this entire fragment out of Movistar Jazz and replace this with that.

OK.

And I'll just have this all in so it looks a little bit better now if there should be if error, because

that's what I call the variable.

And if there's an error, possibly return, give error and then I'll put in error message.

OK, so let's see how this works.

Let's go back to our Web browser.

And there we have some movies that worked very well.

So if I go to home and movies, no problem at all.

Now let me open the console, OK?

And let me come back here and delete this default value like that.

Now, let's go back and look at the console and you see how it's gone into an endless loop, getting

that information from the back end.

There's no way we don't see any error here.

Everything seems to work.

But if you don't set that default value in your use effectively, you're going to get this endless loop.

So let's go back here and put that back in.

There's an empty array and let's delete all of our commented code just to see how it looks.

One thing you'll notice right away is that at least in this case, and it's certainly true in many cases,

when you write your components as functional components and use hooks, you actually wind up with less

code.

And despite my I mean, I actually do prefer using components.

And I think that's become abundantly clear to you as we go through this course.

I will confess that you're going to wind up in a lot of cases with less code when you use functional

components and hooks instead of standard react class components.

So I suspect that as time goes on, I'll become more accustomed to using these instead of classes.

But right now I still kind of like classes, but that's a matter of personal preference.

So this is clearly not that difficult.

And as I said before, these are the two hooks you're going to use far and away the most off use effect

and use stick.

OK, so there's our first one converted.

Now we can go back to Apgar's and we can delete this component luvvies, because we're not using it

anymore.

All right.

In the next lecture will convert another component to a function.

## 109 - Coverting the Genres.js component to a function with hooks

So we've managed to convert our movie's component to a functional component that uses hooks, and I

want to move on to the genres, which is pretty much the same logic as we saw in the last lecture.

As you might have noticed, we don't need to worry about the form components like input just because

those are already simple functions.

And we also don't need to worry.

I don't think about the alert component because again, it's a simple function.

So I'll close those two and get started on genres.

So we're going to react.

And just like last time, we're going to use the use effect and use state and fragment.

And we also need to import like,

OK, that again.

OK, so let's make our function, which I'll call just like I did last time, the name of the old function

followed by func.

And it takes the argument props in case we're passing properties to it and I'll declare two constants,

same as last time.

So a concept and this one will be called genres and it's callback function to set.

It will be set genres again, that uses state and we'll give it a default value of an empty right now.

And again we'll have er const error and set error so we can set it for the state and that will be equal

to use state and we'll give it a default value of nil.

Last time we did an empty string which worked fine and it's not func.

That's my go background.

Coming into play its function, that's better.

And again we're going to use effect and as you can see I have the existing genres class open over here.

Last time we used components did mouth.

This time we'll use use effect with an arrow function and we're going to copy this entire fetch statement.

From genres, jazz and pasted in here, and I'll format this a little better and move this over to their

OK and we'll get rid of all of this.

And change it to if response status is not exactly equal to two hundred, and we'll do one thing, otherwise

we'll do something else.

And the ELT's is simply set error.

No.

And if the response is not 200, then we want to say set error, invalid response.

And it's just response that says, again, less code down here.

All of this goes away and it gets replaced with a very simple set genres.

So we set the state for this function to Jason genres.

OK, so there's our use effect that now we want our render.

So we'll say just like we did last time, if error is not equal to null, then we'll return div there

and just error message.

Otherwise we have something to display and all I can do here is come down and copy this entire return

statement

and paste it in there and now export the default function export default

genre genres func and go back to abduct gess, import it import genres func and we'll get rid of this

genre because we're not using it anymore.

And then scroll down to where the error now exists and change that to genres.

And if everything compiles and it did, I should be able to go back to my Web browser, reload this

and click on genres.

And if I go to say mistery, I should see one entry.

Maybe not, mister.

Maybe it was crime.

Crime.

There it is, The Shawshank Redemption.

And everything works again.

Very simple to make this kind of conversion.

Some of them will be a little bit harder, but I thought I'd start with a couple of really easy ones.

So let's move on to the next part and convert another component to a function that uses Hux.

## 110 - Converting the OneMovie.js component to a function

So we have displaying the list of movies working as a function and we have the displaying the list of

Gene was working as a function.

Now let's display one movie as a function.

So the component that displays one movie is called One Movie.

Dargis and I have that open in the right pane of my editor and I have an empty file called one movie

Funk Jazz Open in the left pane of my editor.

And these are both in the components folder, in the source folder.

So let's get started.

And this is pretty straightforward.

We're going to import react and we're also going to want, as we did for the other ones, use affect

you state and fragment.

And we have no links in this, so we don't need to import link.

So let's create a function function, one movie function and it takes props and we're going to have

two constants, same as before.

But the first one is for a single movie and we'll again use set movie as our function to set its value.

And that's equal to use state and it has an empty object as its default value.

And we'll have another constant same as before error set error, and that is use state with the default

value of null.

Now we're going to use effect same as before and we'll use the arrow function syntax.

And I'll copy the statement from one movie dogs and paste it in here and we'll get rid of everything

in here.

In the if response status is not equal to two hundred, give it an ELT's clause and change 200 to a

number

and if we have an error, we'll just set error, set error, invalid response

with the code response to status.

Otherwise we'll set the error to no.

And down here, we'll get rid of the entire this statement and have a one liner set movie to Jason movie.

Now for the callback on this one, we actually need to pass a specific value.

If you recall, we're getting the idea of the movie from, as you can see here, this prop, Scott,

matched up premise that but we don't have it.

This is what we'll do, is get rid of this entirely and grab that from props.

But we need to pass that to this function to use effect.

And we do that by passing an array consisting of prop match, dot param.

OK, so that's our use effect statement.

And if you come down here, we do a couple of other things in the render function in one movie.

Yes, we have this, for example, if movie Genaro's, if that set, then we convert movie genres to

an object of values, which is the format we need it to be in.

Otherwise we pass at an empty array so we can just copy all of that and put it after this.

So I'll put it right here.

OK, and that will make sure that movie genres is set to a value.

We can actually use the now we do our render the equivalent of a render function in a react component.

And that's just if error is not exactly equal to null, then display the error, otherwise display the

information so we can simply copy this return of error bit and paste it right here and we can copy this

return fragment bit all the way down to here.

And paste it in here.

Let's make sure we have everything here, MoVida Jean-Christophe Map looks good, movie title, all

of that should work.

So if we go back to abduct jazz and import that import one movie, oh, we've got to do our default

export, of course.

Let's go back to one movie funk.

And export default, one movie from there.

Now I can come back here and import that one movie and get rid of one movie

and scroll down to where the error is right here and make that one movie function OK.

Now, if this compiled looks like it did, this should work.

Let's find out.

Back to our browser, go to the list of movies and display American Psycho.

And there it is.

No problem at all.

So a little bit different for that one.

We had to pass the necessary information to allow it to get the ID from the URL, but pretty much the

same logic as we've done in the previous show.

All right, let's move on.

## 111 - Converting the OneGenre.js component to a function

OK, so we seem to have the listing of movies, the listing of genres and the displaying of a single

movie working properly, using react functions.

However, I did make a slight mistake.

If you look at the genres, funk, jazz file, I left out the second optional parameter and empty array

in this case.

And the problem, of course, is if I don't have that in there, we get into that endless loop where

it continually pulls information from the back end.

And that's something I tend to do now and again.

So it's back in there now and life should be good.

So I have, as I did the last time around, open the one genre component in the right hand painting.

My editor and I have an empty file called one genre funk jazz in the left hand ed page.

So let's get started.

We're going to import react same as before and we're also going to use use affect you state and fragment.

OK, and we're also going to in this case, have links.

So let's import the link there and we'll give our functioning function one genre funk and the argument

as props.

And I'll do the export default before I forget export default.

One genre funk there and inside of here.

I'm going to have this time three constants or three variables, one constant and two, the change.

So the first one will be our list of movies.

And I can't use a constant here because the list of movie will change depending on which genre is clicked

on social issues left let movies and my function to set it set.

The state will be set movies and that's equal to you state and we'll give it a default value of an empty

array.

Then we'll have a constant and this will be an error and set error and it's equal to you state with

the default value of nil.

And finally, we're going to have genre name now for genre name.

If you look on one movie or one genre jazz, we're getting that from the props from pop star location,

genre name.

So what we can do here is select genre names because this will change.

Therefore we can't use a constant and I've never been happy with the way the JavaScript works with constants.

But there you go, genre name and we'll give it a function set genre name and that will be equal to

you state and I'll give it a default value of an empty string.

OK, now we'll have our use effect, same as last time and we'll use our arrow syntax.

And in here I want to use the same fetch statement modified as I did in component did mount in one genre

just so it's copy and paste it in here and then fix it.

So I'll just have all this stuff over there.

So if the response status is not exactly equal to the number two hundred, then I'll set an error and

I'll say set error, invalid response

and that'll just be response status.

Otherwise we'll set our original set error null.

OK, and down here, we can't use this set state.

Let's get rid of that.

And we need to do two things here this time.

The first one, let's fix this to response there.

The first thing we're going to do is we're obviously going to set moving so we can do that, set movies

to Jason movies.

But we also need to get the genre name because we want that to display above the list of movies so people

know what genre they're looking at.

So all we're going to do here is set it in state set, genre name.

And the value we're setting it to is going to be props dot location.

Doug, Jean-Rene, OK, and up here, will we make the call to fetch obviously we can't use this.

We have to get rid of that, but that should be OK.

And down here, we need to set a couple of things for our actual information.

We're passing to our use of fact function.

And the first one is we're grabbing the ID from the U.S. or else.

So we need to pass props, match DOT Hiram's ID and we're also using props, location,

dot genre name.

OK, and if we don't do that, nothing will work.

So we have our use effect function that seems to be reasonably complete here.

We'll find out shortly, I guess.

And after that, we need to do some other things as well.

For example, here in one genre, jazz, we're actually checking to make sure that movies is actually

a value, that it has something in it.

And if it's not, we set it to an empty array.

And of course, that's a situation where we're looking at a genre that doesn't have any movies in it.

And that's certainly true for my database.

It may not be for yours, but it is for mine.

So I'll just copy and paste the Tabit back to make it look a little better.

OK, and now we can do our equivalent of the render function.

So if error is not exactly equal to know, then we're going to display an error.

Otherwise we're going to display our information.

So I'll copy the error message from here.

And paste it in here and then I'll copy the entire return fragment from one genre right there, copy

it and paste that in here and see if we get everything right.

OK, so we'll go back to after Jess and we'll import our function.

We just created import one genre funk.

We'll get rid of one genre up here.

Scroll down to the air, which is right here, where I'm rooting to a specific genre and change that

to one genre punk and let's open a terminal, everything appears to have compiled.

Let's go back to our Web browser, reload this and go to home and then go to genres and click on crime.

And there we are.

We have our genre, we have our crime, and we have The Shawshank Redemption.

And this should link right through to one movie.

And it does.

OK, that's it for this time.

So we'll pick it up in the next lecture and continue our conversion from Riak classes to react functions.

## 112 - Converting the EditMovie.js component to a function

So we're coming along, we have movies working as a function instead of a class, we have the display,

a single movie, the genres and display individual genres saw is working.

And now we want to work on this part where we actually manage the catalog.

And we'll start with a catalog item which has a form.

So let's convert that one to a function.

So go back to my code.

And as you can see, I have edited movie, which is the component in question.

I have that open over here and we're going to recreate it as a function over here.

So what I'll do, first of all, just to save some typing, is copy all of this and change it to start

import statements.

So we're not going to import component.

We're going to import use effect.

And use state, OK, everything else we're absolutely going to need, so let's create our function,

function at a movie funk, I'll use the same logic I did before, the same syntax, and it takes props.

Then let's get our expert here, because I always forget to do that.

OK, inside of here, we're going to set up a few constants, so concert movie for the individual movie

and the function to set it in state will be set moving and that will be equal to you state.

And the default value will be an empty JavaScript object.

And we'll create another one for error set error and it will be equal to you state with no the default.

And we'll do another one for errors because we're doing form validation here and we wanted an array

of errors.

If there are any const errors, set errors, plural will be equal to use state and an empty array is

the default.

And finally we'll have const alert because we're going to be displaying alerts when things go wrong.

Set alert and I'll be equal to use state.

And its default will be a hidden alert, so as a JavaScript object type, the class to hide things is

done none in bootstrap and the message will just be an empty string.

And finally, we're going to need another one for MPLX options over here in the component.

We're putting that as part of the state.

So what we'll do here is create a another variable concept and it will be called MPLX options.

And that will be equal to an array of I'll just copy all of this information.

And then we could just call it as a local variable.

So there's our constants.

Now we'll set up our use effect.

Use effect.

Using the arrows, syntax and what do we have here in our component did mount, but we have this first

check to make sure that someone is logged in.

So let's copied and pasted and modify it as necessary.

Now, obviously, we don't have this, but we do have props so we can just get rid of this.

And that should work just fine for the first part, so let's just format this properly.

And that will verify that people actually have a valid JWT before going any further, if they don't,

we just knock them out of there, then we have a constant over here and that it movida us the ID, which

I can copy and modify.

And again, it's pretty simple to modify that.

We just get rid of this because we do have props available to us, then we have this long if statement,

so let's copy that whole statement and modify it as necessary.

And I'll just copy this and leave the ELT's out just to make things simple.

OK.

So what needs to change in here?

Well, obviously, we're going to convert this to a number again.

We're going to get rid of all of this and just do the standard set error.

So something went wrong.

We'll display the error message and it's just going to be as simple as invalid response.

And then it is response status, same as before.

Otherwise, we do need an elsea will say set error.

No, there is no error.

So we'll reset error to no.

OK, so that's our first part.

The second part is also going to be a lot simpler.

We'll get rid of this set state because we don't have that available to us.

And we'll have a one liner here.

Actually, it'll be more than one line.

So we do have a release date const release that we need to get this date in the correct format.

So we have JSON available to us.

And then we really need to do here is to say, Jason, movie release, underscore date is equal to and

then I'll just copy this over here, get it in the right format, copy paste and get it in the format

we need it to be.

And then we just set movie to Jason movie.

So all we've done here on this line line forty five is overwrite the default format we're getting from

the database to the format we want it to be to display it on the form properly.

So that should take care of that.

Now are we finished with this use this use effect?

I don't think so, because we're actually using a number of things in here that we need to tell the

use effect function of.

OK, so we're using props JWT, for example, and we're using props, DOT matched up parameters ID and

we're also using props to our history up here.

So we just need to tell you the effect of both that.

And again, that goes in the form of an array and we put in props, dot history props, dot JWT and

props dot match dot programs.

And now it has the information it needs in order to move forward.

So that's our use effect.

And as you can see, use effect is actually now a lot shorter than component did mount.

And the more I'm working with hooks and functions and more, I'm liking it because you wind up with

less code, which is pretty good.

So the next thing we have to worry about is our let's see here.

We have components amount.

So that's done handle submit.

We'll do that next.

Let's take care of handle change right now.

Now, over here, we just use the arrow function because we're in a class.

We can do it this way.

But since we're in a function over here, let's go to this line after our use effect hook and declare,

which we call handle change.

And now we can use the arrow function.

If I could type it, there we go now and handle change, what is it going to handle?

Well, it's going to have.

There's something missing here.

Oh, yes.

I need to give it the event that's better over here and handle change, we have all of this logic.

Let's copy that, paste it in here and modify it as necessary.

Now, name and value, let let value equal event target value and let name equal event or target name.

That's fine.

But this set statement obviously has to change because we can't do it that way.

We don't have event.

We don't have access to this state because we don't have access to this.

So instead we use the set movie function.

And what we did in the other one is to populate it with what exists from the state, which we call previous

state or prev state dot movie.

But here we don't have that.

So instead we can just do it this way.

We can say grab all the values from the local variable movie and then take the name and override it

with the value which are value in name from lines fifty two and fifty three.

So exactly the same logic.

But instead of using set state we're just using set movie and that should work fine.

So there's our handle change.

That should be finished now let's work on Handal submit, let's find that where's handle submit?

Right here, so line ninety one in that movie in my code is the handle submit function and we want to

rewrite that to work with functions.

So we're going to use the same logic.

It'll be cost, handle, submit.

Is equal to and we're going to have the event and then our area function and what I'll do is copy everything

out of here and modify it as necessary.

Copy that and paste it in here.

Now let's go fix it.

First of all, let's tab things over a little bit so we can see it.

So first thing we're doing is preventing default, and that's fine.

That can stay the same.

And then we have our errors.

Let errors equal a..

That's fine.

And then we have if this state movie title again, we just want to use movie because we're not getting

it from state.

It has a local variable and we'll say errors push error for title.

And here where we're saying this dot set state errors, errors, we can't do that.

But we can say set errors, plural.

Make sure you get the plural one and just hand it errors.

And then we check to see if the length of errors is greater than zero, if it is, there's an error,

so we return false.

That's fine.

And now we actually need to worry about submitting this.

Now, what do we have to change here?

So we passed.

So post info.

The data is fine, the payload is fine, the headers is fine.

And then we append our content type application, JSON good and we append authorization.

It's not this stop stuff data, but it's prop stuff.

JWT, so that's fine.

Now, if there's an error after we submit to this YORO, then we can't set state if we say if data error,

let's delete that and it in and change that to set alert.

And we'll simply pass it the kinds of things we want, alert column type, alert danger and the message

is dated error message.

And put a comma after that just in case, because sometimes JavaScript doesn't like it when there's

no comma there.

OK, so there's that.

Otherwise, we redirect them to the admin URL and we just get rid of this before props.

And type things in so they're a little more attractive.

OK, so that works or should work anywhere.

All right, so that should take care of submitting it and data validation.

And as you recall, we're only validating on title because it's simple enough for you to add your own

validation for whatever field you want to validate.

So let's go a little bit further.

Next, we need to worry about confirming a delete.

We want to delete this.

Let me find that.

So confirm.

Delete.

Where is that?

Here it is.

So confirm, delete.

Let's create another function here.

And again, it'll be a constant and I'll call it confirm delete.

And that will be equal to passing an E, which is the one I used before.

And it's sort of that we want to confirm alerts.

So let's copy that, confirm or confirm alert function and modify it as necessary.

Copy paste.

And what do we need to change here?

So, again, we're not going to have this prop stock, Judy, but we'll just have prop stock JWT.

We're not going to go to this state movie.

It is just a movie movie, Dot, Heidi.

That takes care of that one.

And then again, we're not going to set the alert this way.

We'll delete that and change that to.

There is an error, so alert.

Type

alert danger and the message is just data error message.

We'll put a comma there, not a comma, no, we don't need a comma there when you semicolon there.

Then I'll copy this and change it for the success message.

So this goes away and gets replaced with alert success and the message is just movie deleted.

And the redirect, we just get rid of this in front of props, and that should take care of that.

OK, so that one works.

All right.

And we also need a has error function, which we're using up here.

So let's copy that.

And we can put it anywhere we want, so has error now, we can't do it this way, obviously there has

to be a function and we don't want to return this stuff.

State the errors.

DUDNIK So we can just get rid of this DOT state because errors is a local variable.

OK, so let's format this a little better.

So there's has error and at this point, we're ready to do the equivalent of a render function in a

standard react component.

So we'll say if an error does not exactly equal to null, we're going to display their faults.

We'll display the form which should be populated with information if we're editing a movie and should

have no information in the form if we're adding a movie.

So let's find the render function right here.

Will return this by copying that and return the error message if there is an error.

Otherwise we're going to return everything here.

So let me copy this entire return statement and modify it as necessary.

Copy.

And paste, OK, now there are a few things that have to change here.

For example, we're not going to have this state dot alert.

It will just be alert.

So we'll get rid of that.

And we're not going to have this dot stay here.

So we'll get rid of that.

And it's not this dot handle.

Submit, when we submit is just handle submit there.

And we also need to change this on change.

It's not this dot handle change, its handle change and we pass it title.

And down here, it's not this stuff handle change, it's handled change.

And we have it.

This one's actually moved.

It'll never change, but I like to be consistent.

And down here, this is tough to handle change title in the air or do becomes this, this goes away,

it's just call that function directly and that's fine.

And up here, we get rid of this.

That should take care of it for title and for real estate.

We'll just take care of this by putting handle change and release date.

That's fine, and this becomes handled change.

Runtime

and MPLX rating, we're not going to hand something from the state, we're going to have that constantly

declared right at the start of this lecture, MPLX options that stays the same.

Then this becomes and will change MPLX, a underscore rating.

That's fine, and they'll change here, becomes and will change and we pass it rating.

And finally, description and change

and description.

OK, so that's the same this story confirmed the leak becomes confirmed delete.

I think that should take care of it.

So let's go back to our abduh as an that function so important.

And it's called Edita movie Folke.

And we'll get rid of this one at a movie.

And that should show us an error way down here, which becomes a hit movie.

OK, now let's open our terminal, make sure everything compiled it did, and let's have a look.

So back to our Web browser and I'll go to the home screen.

We're logged in.

Good.

And we'll go to manage catalog and I'll look at American Psycho.

And there it is all populated with information.

And if I delete the title and try to submit it, my form validation should stop me from doing so.

And it does.

And now if I call it American Psycho and put the number two after it, which will fix in a minute and

save it, it should save it and update the title.

And it did.

Perfect.

And now if we come in here and delete that and save it again, it should go back to the correct information.

All right.

Now, if we go back and look at our code and hide the terminal, my existing moviegoers react component

has two hundred and ninety five lines of code and my shorter function, one only has one hundred and

forty three.

And I'm beginning to see the appeal of using functions and hooks because it's less code and less code

is always less code to maintain and less opportunity for errors.

So the more I work with this, the more I'm actually starting to like it.

All right.

We still have some work to do, so let's move on.

## 113 - Challenge: convert Admin.js to a function

So this time around, I have a bit of a challenge for you, and it's a pretty straightforward one,

and you should have no difficulty with it because it's covering things we've done repeatedly in this

section of the course, and that is to convert the admin jazz component from a component to a function.

And one other thing I need to mention, last time I said I like to be consistent when working on the

edit movie function.

And here in the return with this fragment, you probably noticed this.

I actually put movie ID for the on change handler and that should actually be handled change.

And then it's just it and that never gets cold, but I like to be consistent, and that was just an

oversight on my part.

But back to admin dojos.

So in this challenge, I want you to actually take the necessary steps to convert this into a reactor

function just to a plain function that uses hooks.

So the steps necessary to do that are as follows.

What we're going to do is create an empty admin function file inside the components folder.

Then you'll do the necessary import.

So you'll import, react and use effect and use state and whatever else you need.

Then you'll create the necessary function, which is just going to be named admin funk probably.

And it will take one argument props.

Then you'll declare your necessary constants to take care of the use state hook functionality and then

you'll convert.

The component did mount function from the react component into a use effect hook and make the changes

as necessary and finally recreate the render function from the component as plain old JavaScript.

So give that a whirl and I'll show you how I did it in the next lecture.

## 114 - Solution to challenge

So how did you make out with the challenge, I suspect you didn't find it terribly challenging.

Well, here's how I completed it.

So first of all, I created an empty function or empty file in the components folder, inside the source

folder called admin function, just using the same naming convention I've been using all along in this

section that I did my imports and I'm using react, use, state use effect and fragment.

And I'm also using link from the react router that I created my function and I created two constants,

both using the use state hook.

One was for the movies, the list of movies we're getting from the back end, and the other was for

any error we might encounter along the way.

Then I created the recreated the component did mount functionality using the use effect.

OK, first of all, I checked for log in to make sure that people have the necessary durability in order

to access this and I remove the this from the props.

So where it used to say this prop JWT on line nine, now it just says prop star JWT and the same for

prop star history on the next line that I modified my statement and it became much simpler.

First of all, we check for an error.

So if the status from the response is not exactly equal to two hundred, we just set the error.

Otherwise we set the error to null using the use use state hook.

And then finally I set movies to JSON movies and gave the necessary arguments, secondary argument in

the form of an array because we're using props, JWT and props for history, we need to pass that to

the use effect so it knows what to do.

And then finally, I just recreated the render function, which is exactly the same logic and almost

exactly the same code as the render function on the component itself.

And finally, hopefully you didn't forget to do this, because I often do, as I'm sure you've noticed,

I exported default admin --.

So that was all the change necessary in the admin function file.

And back in Apte Jazz, I just modified the import to use import the new function we just created.

I got rid of the old admin and then I changed the link down here, wherever admin is to use admin func

instead of admin and that's all I had to do.

All right, let's move on.

## 115 - Convert Login.js to a function

So this time around, we want to take care of this, where we're trying to log into the system and this

one's a little bit different than the ones we've done so far, but it's not that difficult.

So let's get started.

I'll go over to my ID and I have on the right hand side of my window the existing logging component.

And over here on the left hand side, I have an empty logging function file, which exists in the components

folder, inside the source folder.

So let's do the imports IMER and we're going to import react and we're going to want you state.

But this one actually doesn't have a component, did malfunction in the component.

So we don't need use effect, but we do want fragment and we also want these two, which I'll just copy

from logging, just the import of input and alert and then we'll give our function a name function,

log in func and it takes props as an argument.

And I'm going to do my export now.

So I don't forget export default logging.

OK, now we want some constants and we're going to have for this time around and I'm going to have one

constant for email, the input email on the form and one for password.

So let's do those const email and the function to set it is set email and that's equal to use state

with a default of an empty string and I'll duplicate that.

Make this one password and this becomes set password and it has a default as well as an empty string.

And also because this is a form, I'll have some validation.

So I'll have errors with set errors and that will be an array.

You state default to an empty array.

And finally, we're going to have alerts so const alert and set alert just to show the invalid login

message and that will be use state with the default value.

That's a JavaScript object type is deduction on bootstraps, hidden class and an empty string for the

message.

So those are Constance.

Next, we're going to have a handle submit now over here we have this for our handle segment, but of

course, we need to use the syntax for our function.

So we'll just declare a constant called handle submit and it will be equal to and will handle the event

and user error syntax.

So the first thing we'll do is an event to prevent default.

So I don't think we did the last time around and it seemed to work, but we really shouldn't have that

there.

Now we'll check to see if we have both email and password.

So if email is exactly equal to an empty string, then we want to add an error errors, push for the

field email and I'll duplicate this.

And change it for password.

Those are two required fields because you can't log in without both of those.

And we'll set our senators in the state set errors, which is just errors, and we'll check to see if

there are errors.

If errors length is greater than zero, then just return false.

Don't go any further because they don't have a username and password typed in.

Next, we actually do our authentication check.

So let's copy all of this from login doorjambs and paste it in here and modify as necessary.

So let's tab this stuff over so it's a little more readable.

They're the first two are fine, request options is fine or call here is fine and this is what we want

to change.

So I'm just going to get rid of this part if data error.

Well.

We can't set state this way, so let's delete that and instead do alert, just like we did before,

and it takes a JavaScript object type will be alert.

Gosh, danger, which is boot straps, alert class with a nice red background.

And I'll specify a message and I'll just say invalid logic like that.

That's all we have to do there, I believe.

Now, if this is correct, I'm not going to console all of the data.

If we have the valid login, then we can't call this DOT handle, GW Change, JWT change, and we don't

even have a function named JWT change.

But we'll do that in a minute.

Everything else here should be OK.

So that's our login.

Now the actual handle JWT change over here in login.

J.S. is this.

Let's copy that and come over here after a statement.

Actually, after our handle's segment, which is right here, type the keyword function and paste and

get rid of the this on props.

And since we're getting that function, handle JWT change as a prop., that's all we need to do.

So that should take care of that for us.

We also need Hazra, as you can see over here, we have Hazra.

So let's copy that.

And paste it after the keyword functional and get rid of this state on errors, and that should be fine

for that.

Now we need to create some handlers for both of our email and password functions for the inputs and

a function, and I'll call it handle email, and it will take an event as a parameter.

And all we need to do is to say set, you know, to target value.

So whatever they type in, the input will automatically update everything as appropriate.

And then I can copy this and paste it and change it to handle password

and change this.

To set password.

And that should take care of that.

So the last thing we need to do is our actual return and we're going to return our favor, which I will

copy from over here.

This entire fragment and then make the changes as necessary.

OK, so let's get started.

The title is Fine, the alert can't be this stop state alert, it's just alert type because we're not

in a class and alert is now a local variable and its state is managed by the U.S. state hook.

We can't do handle, submit or this to handle submit.

We have to do submit, just handle submit by itself.

Handle change becomes.

Handal email, and we get to this on Hiser, and this has air and down here, this becomes Handal password

and we get rid of that.

This is.

And I think that should be it.

So if we go back to Abduh chairs and we import that function, we've just created import, log-in,

funk and get rid of this log-in.

And then go down to where the air is marked.

Right here and changes to logging folk, open our terminal, make sure everything compiled, it looks

like it did.

Let's give it a try.

So back to our Web browser.

Refresh this and go to the home screen and click on login and let's give it invalid data.

First you at somewhere dot com with a password of test and that should give you my alert at the top

of the form.

And it does.

And now if we go here dot com with the password password, we should log in.

And I have this is undefined.

So there's an error somewhere.

This dot props dot history to push.

We can fix that easily enough.

Let's go back to our login funk and just search for the keyword.

This.

There it is.

And search again and it's nowhere else.

So we should be able to log in now.

So let's go back to the home screen, log out, log in

here dot com with the password.

Password.

And we're fine, perfect, and we should be able to see a movie and cancel and everything works the

way that it should.

So we have pretty much everything converted to a function.

Now, the one thing that's not a function is right over here.

I thought, yes, that is still a component and we can actually convert this to a function.

Honestly, I would probably stop at this point and say I'm happy, things are good, but in order to

make this a complete course, using functions and hooks in the next lecture or two will convert this

to a function as well.

## 116 - Convert App.js to a function

So the final step in our conversion of our project from using classes to using functions is to convert

the abduh component, the top level component for our project.

There is still the graphical stuff, but I'm not going to worry about that because that's a simple process.

And we have many, many examples of how to do that in this section of the course.

So I have open on the right hand side of my editor at my Abduh file and I have a blank app, dot funk

jazz file open in the left hand side of my editor.

And this is inside the source folder right beside Abdah Jazz.

OK, so let's start will import import react and we want react and we also want you state because we're

going to use state use effect because we have a component dismount and fragment, because we have a

fragment in our return.

And we also want all of this other stuff.

And I can just copy and paste it from jazz to funk ducktails because this is exactly the same import

that we need and now we'll have a function.

So we'll export a default function, which I'll call funk, and it takes the argument problems.

And instead of that, we're going to have first of all, we don't have a constructor because this is

not a class, but we do have one constant we need to keep track of in state.

And that's our JWT.

So we'll call it JWT and we'll call the function to set its value set JWT all uppercase like that.

And that's equal to use state or use state hook and it will default to an empty string.

OK, and as you can see over here, we actually have a component did malfunction.

So the functional equivalent is use effect and we'll use our arrow syntax and we'll do exactly the same

sort of thing we did here.

So let me copy this and paste it in.

And of course, this determines whether or not someone is logged in at any given point in time.

So we'll just format things to make it a little bit better.

And there we go, and we can't say this set state, so we'll get rid of that and change that to such

a beauty.

And the argument will be Jason Persse TE.

And of course, since we're using this JWT and here we actually need to pass it as the second parameter,

which needs to be an array.

So we'll put JWT in there and that takes care of our use effect.

So the next thing we need to worry about is this handled JWT change, so let me copy that and we'll

modify it.

First of all, we'll start with the key word function because we're using functions and paste and we'll

just change this to an ordinary function, which will work just fine

and change this to set JWT.

And the argument is just JWT and that's all we have to do for the handle.

JWT change.

Now, what else do we have?

Let's fix our formatting here.

Next, we have our Loga function well, again, keyword function and let's copy this and paste it in

here and change it into an ordinary function.

So we're not going to set state this way, we're going to set state or set JWT state to attempt to strengthen

and remove that from local storage if it exists.

Perfect.

And we'll format this again.

There we go.

OK, and we don't need a semicolon there.

So now, if you look at our render function, we're actually constructing using an if statement, our

login link, which is log when you're not logged in and log out when you are logged in.

So we'll do that outside of the return statement and we'll say if JWT is equal to an empty string so

we don't check the state because it's a local variable.

Now, that's fine.

Let's just fix the formatting a little bit.

And fix this formatting a little bit there, and it's not this stop log logout is just.

So that was pretty simple.

So we have logging link and logging link.

Those are fine.

Now we want to return RSX, so we'll just put some brackets in there and I have a place to paste it

and I'll copy from our log in just the entire Rohter section there, there and paste it in.

And now it's just a matter of searching for the keyword this.

So I'll search for this and we change that to plain old JWT and cloverleaf G to find the next occurrence.

It just becomes handle due, but change and this is what will lift our state for us.

So we're passing this handle JWT change attribute or property.

We're passing a function.

So we read that in the log in function.

And whenever there's a change, we actually change the state at the top level.

So that's how we're lifting state.

So let's keep looking for this.

There there's another one and that just became becomes JWT and this becomes JWT.

Are there any more?

This is there's one right there.

So this becomes JWT.

And let's look for this again.

And there are no more occurrences.

OK, so that should be all I have to do for this function.

Now we go to index jazz and we import.

Up func.

There it is.

I don't know why we have a twice, but there you go and we get rid of this one and we change this to

APOC and if we did everything right, it should compile.

Looks like it did.

So let's go back to our Web browser.

So far, so good.

So let's try movies there.

Let's try logging out a change to log in.

Let's try logging in me here dot com with the password.

Password.

Perfect.

So now we have managed to convert our entire project from classes to nothing more than functions and

hooks.

And I think it was a useful exercise.

But as I said a while ago, it's still very important to know how to use classes because there are millions

of lines of code out there.

React to applications built using classes, and they're still being built using classes by many people.

So it's important to know both approaches.

And certainly you wind up with less code using functions and hooks.

But there are lots of employers looking for people with the skills to work with reactor components as

well.

All right.

There we go.
