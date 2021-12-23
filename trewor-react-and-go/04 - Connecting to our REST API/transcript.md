# 04 - Connecting to our REST API

## 48 - Setting up CORS middleware

So things are starting to take shape here, and there's probably at least one person who's decided to

go and try to hook up the react front end to our rest API back end.

And you probably wrote all of the code and JavaScript absolutely flawlessly and still couldn't get it

to work and kept seeing an error, something about c o r s.

And that's because our react front end is actually running on an application on PT..

Three thousand and our react back end is running on Port four thousand.

And it's doing that because you can't share ports with two different application servers.

So when that happens, the built in security feature for most modern browsers says, well, unless the

backend you're connecting to explicitly says allow connections from wherever you currently are, then

deny it.

So we need to fix that and we can fix it really easily in our go back end by adding a little bit of

middleware.

So I'm on my go application right now and in my command API folder, I'm going to create a new file,

which I'll call the middleware go

and I'll give it its package declaration.

And it's in the main package and I'll have one function here.

Did simple middleware and it has the application receiver and I'm going to call this middleware enable

Congress cross origin requests and like all middleware and go, we have to give it a parameter, one

parameter which I'll call next, and that's a type HTTP dog handler and it has to return and the handler.

So to do that, we're going to have a dead simple return statement return and it has to be an active

beta handler func, which takes an inline function as its argument with the parameters of W response

writer

and R, which is a pointer to a request.

And inside of that will simply said the header header dot set and we're going to give it the key of

access with a capital a hyphen control with a capital C hyphen, allow with the capital a hyphen origin.

And the argument or the value for that key is simply to allow all requests.

And then we just as you have to add a handler or in a middleware next dot serve http and give it the

W and the R response operator and a request.

So that is our middleware and it's dead simple.

Now there's one more change we have to make, actually two changes in the routes file and they're pretty

simple.

First of all, we don't want to return this type from the HTP router package instead will return HDB

to a handler, which the one from the HTP rotor package satisfies anyway.

So this actually will work even with no other changes.

But we want to enable the middleware.

So we're going to do that by returning not just Rotar.

Delete that.

We're going to return app, enable course and pass it the router and that adds the middleware to every

single request that comes to our application.

And that's the only change we need to make.

Now, allowing cross origin requests from everywhere may be exactly what you want.

Maybe you have a public API or perhaps you want to narrow it down a bit and only allow it from specific

locations.

And we'll talk more about how to tighten up the constraints on that middleware a bit later in the course.

But this will get us going in any case.

So let's move on and try connecting and react frontin to our go back end.

## 49 - Getting the list of movies

So this time around, we want to connect our react front end to our go back end, and I'm going to do

just one of the rest API reports, the one that lists all movies.

In order to do this, I have to have my REACT application open in one window in visual studio code and

in the other one, I need to have my back end.

Let me find that this one right here, I have to have my back end actually running.

So I have my back in the go application open in one window and I'm going to run it, as you can see

here.

Go run dot slash cmd slash API.

That will start the application listening on four thousand.

And I'm going to hide that.

I just want it to run in the background and back here in my react application, I'll open a terminal

window and run npm start.

And that should fire up a browser window with our application running and there it is.

So let's go back to our coat and start making some changes.

So hi, my terminal window now, right now I'm looking at the component movies, and this, of course,

is the one that lists all of the movies.

And I need to make some changes here, because right now in the component dismount, I'm just hard coating

this value.

So I'm going to delete that entirely.

I don't need that anymore.

And I'm also going to make one change to the state.

I'm going to add a member to that or out of value to that.

OK, I'll show you why in a second.

So movies will start as an empty row.

And I'm also going to add this because I'm a paranoid sort of person belt and suspenders when it comes

to any kind of network programming, I'm going to add a value here that says is loaded and it's just

a boolean and it will default to false.

OK, so that's available to us in our state and inside my component did mount.

I'm going to do this.

I'm going to call fetch, just like we did in the first react application we built.

And I'm going to fetch our route htp colon slash slash localhost colon four thousand.

And I want version one slash movies, plural, and that's the list of all movies.

And then of course I have to call them and say, when you get the response, turn that into response,

Jason.

And after that's done then at this point I'm going to say Jason is fed to this code.

And what I'm going to do with it is pretty straightforward, this dot set state.

And I'm going to set the state for two things.

The only two things in my state.

The first one is movies.

And that's not equal to Jason.

Don't forget, we wrap that.

And so what's inside of Jason Dot movies, right?

Our rapper is called Movies in the Jason File, and now all set is loaded to true.

And here's why I'm doing that.

OK, so this component did Mount Vires.

It goes to the API, it gets the necessary information.

It populates two values in our state and down here in the return or in the render function, I'm going

to read a couple of values.

So I'll declare constant and I'm going to take from state these two values movies and is loaded.

Now, I could just go state dot movies and so forth, but this just makes more concise code.

So those are equal to this state.

So now I have two variables I can use, and I'm going to execute a simple if statement if and this is

why I'm having these loaded variable, if not is loaded, then I'm going to return.

I'll just return a paragraph of text loading like that, OK else.

And I'll return this.

And of course, we have to make a change down here as well.

So that goes here to close this.

Let's format it all and make sure I have my braces in the right spot.

Looks good.

So I'm not going to go over this stuff straight up movie star map.

I can just say I can get rid of this dark state and say movies, dark map and everything else should

stay pretty much the same.

So let's save this and switch to our Web browser and see how it's working.

So everything should have reloaded.

I'm going to open the JavaScript console just in case there's an error and I'll say go to movies and

there are all of my movies.

And you can see in the console here, we got our request using X HRR, so it fires off his Ajax and

we get a response and everything looks exactly as it should.

There's all of the information that I need.

Well, perfect.

Well, that wasn't difficult at all.

Now there's some other things we need to do here.

As I said, it's good to be paranoid when you're doing any network programming.

So we probably should add some error checking into this and we'll take care of that in the next lecture.

## 50 - Checking for errors

So displaying the list of movies seems to work really well, and we can verify that right now by going

to our Web browser and just looking at the list of movies and there they are.

So I'm going to go back to the homepage, then I'm going to go back to my code and I'm going to simulate

a network error.

And I'll do that just by changing the URL, by adding a couple of X's at the end of it.

So that's actually going to generate a 404 page, not found error from our back end.

And if I go back to my Web browser now and I have the JavaScript console opens, you'll see there will

be an error.

But look at what the end user sees.

It just goes to loading and it stays there forever.

Now, the actual error is Jason Persse, unexpected non whitespace character and so on and so forth.

That's because it's not getting Jason back.

It's actually getting a page not found 404 error.

And we have no way of letting the user know that something went wrong.

So we probably should address that.

Someone go back to my code and the very first thing I'm going to do is I'll add a new item to my state

and I'll call it error.

And its initial value will be no like that.

And I'm going to fix this URL right now and we'll change it.

Back in a moment.

The next thing I'll do is down here in my render function.

I want to have access to that error.

Variable soldiers say they're like that and that will grab it from state and populate a local variable

named error.

Now, in my component did malfunction.

I'm going to comment this first line and we'll get rid of it in a minute.

And I'm going to replace it with something else then.

And I'm going to take a response and use an arrow function.

And inside of that, first thing I'll do is say console dot log status code is and then response status.

And that will be set to two hundred the string two hundred if everything worked as expected.

But if we get a 404 page not found then it will be the string four or four.

If there's an error of five hundred we'll get the string five hundred.

The point is this is something I can actually test against.

It's already in its statement here.

If response status

is not exactly equal to two hundred, then something went wrong.

So what I'll do is generate a new error, an actual JavaScript error.

So I'll create a variable which I'll call for, and that's equal to error.

And that gives me a new empty JavaScript error and I'll populate its message error message.

Is equal to invalid response code, and then I'll just append response status.

And then I'll set the state, this state set state and I'm setting the state for error will be equal

to her like the.

And then after the if statement, I'll just return response to Jason now, there's one more change that

we have to make here.

OK, I'll say a couple of more changes.

But in this last then statement, then Jason Carroll function after its closing curly bracket and before

the closing parentheses.

We need to add this.

Error and an error function.

This starts set St..

And we'll set a couple of things here we can say is loaded is true, and just because I'm like this,

I'll set error as well.

OK, and we need a semicolon right here.

Now, the last thing I'll change is down here.

We need to use this error function.

So this becomes an elusive and the line just before we can say if error and I'll move this curly bracket

down here so it's a little more attractive.

And here I'll return a div error and then my error message.

Error message.

Now, if we did everything right, right now I have a valid URL, so I should be able to go back and

actually look at the list of movies and get it.

And there should be no error in the JavaScript console.

So let's try it.

Movies we've got status code is two hundred and it displays the list of movies.

Let's go back to the homepage and let's put an invalid URL in there and now we'll come back and say

movies and it says Error Invalid Response Code 404.

Now that is a much more useful way of giving meaningful feedback to the user.

And this is just a simple example of how you can trap for errors.

There are many ways of doing it, but this actually works pretty well.

And one other nice feature about this is the fact that we're using our error checking or doing our error,

checking this way on line twenty eight through thirty three instead of using a catch block that's available

to us in JavaScript.

What would happen if we used a catch block is that we might have actual bugs in our component that we

never see and we think it's a network error instead.

And this makes sure that if there's a bug in the component, I'll actually get meaningful feedback from

react.

And if there's a problem with our network transport, then we'll get an error message and it keeps our

error handling much, much simpler.

All right, let's move on.

## 51 - Displaying one movie

So we have a list of movies displaying properly, and then we want to use the same logic and apply that

to displaying an individual movie with some details.

So I'm looking at movie star Jess right now.

And just to save some time, I'm going to copy this entire fetch statement from here and go over to

the component rendering one movie, which is called One Movie Dojos, and in its component did mount

or delete all of this and paste in what I just copied and then fix it.

So the euro is going to change and it's not going to be the one movies.

It'll be one v one movie slash.

And then I have to append this props, just like we did before match forams.

Thought I'd.

OK, so that's the correct euro.

Now, what else do we have to change here?

We definitely want to add a couple of things to our state.

So we'll have is loaded and that will default defaults and error.

Which will default.

No.

OK, so this stays the same.

This stays the same here.

It's not movies, it's movie singular.

And Jason Bourne movie singular because that's our wrapper in the Jason we're feeding back and that

looks good.

OK, now down in our ORENDER, let's extract the variables we want.

So we want CONSED and we want movie.

We'll extract that to a local variable is loaded, an error, and they all come from this state is equal

to this state there.

So now we have three local variables, which means we can do things like this.

We can, first of all, do a little bit of logic.

So back in our movies, not just in our render function, we had this kind of stuff.

So let me copy that, because we can use that same logic here and go back to one movie, Jackass, and

have it like this.

And then, of course, we have to have a closing parentheses right down here, and that should fix the

errors.

Let's format it and let's go back up and fix our if statement.

So that will be fine.

If there's an error display that if it's not loaded display loading, that's fine.

Now, let's clean this up.

We don't need to go this state dot movie because back on line three, we extracted a local variable

movie from the state and we don't want to display the ID here, but what we can do instead is maybe

put the year after the title of the movie.

So I'll put an opening parentheses and then movie dot year and a closing parentheses and that looks

good.

And underneath that, I want to display the empty AAA rating and also the genres, if there are any.

OK, now there's a problem with that.

Right now.

I want to use the map function to go over the genres.

And if you recall, I actually simplified the way that our JSON was being displayed so it wasn't returning

an array.

And if it wasn't a I could have just used a map function here.

But I like to do things in a way that makes sense to me.

And if you don't, you can go back and modify your rest API to return the genres as an array.

But I like to do it this way.

So I'm just going to, first of all, make sure that the movie has some genres, because right now only

one movie in our database has genres.

In my case at least, that's a Shawshank Redemption.

So I'll just say if movie Genaro's, if that is set, then I'm going to do something, otherwise I'm

going to default the value of MoVida Jarrah's.

Try that again.

Movie genres will just be an empty array, otherwise I'm going to convert it to an array and I can do

that using the standard JavaScript syntax, which is as simple as movie genres, is equal to object

values.

And we convert our Java JavaScript object, MoVida Genaro's into an array and now it's an array.

So that means I can do something with it.

So down here after the title, I want to display, first of all, on the left hand side of that portion

of the screen, the NPA rating.

So I'll use bootstraps div and its class and of course I have to use class name because I'm in X is

equal to float start and inside of that I'll just display.

And don't do it in small text movie DOT, MPLX, AAA rating, and before I'll just say rating and then

over on the right hand side.

So I'll use div class name equals float and I'll display the genres if any.

And that means they have to exist.

So I do that using the same math logic we did when we were listing all movies.

And now that I have genres in an array, I can do that.

So movies, MoVida Genaro's Dot Map and I'm going to map, I want to pass it to things.

M will be the current iteration of the array and I'm also going to pass it the index and I'll show you

why in a moment.

OK.

And that is equal to the arrow function and open parentheses.

And inside of that I'm going to put a span in.

I'll make it as a nice little bootstrap badge span class name equals badge, big dash secondary and

I'm going to give it a little bit of margin on the right hand side.

So m e for end, I'll just use one space that'll be big enough and now I'll have the name of the current

genre.

And I passed the index because if you recall, any time we're doing this sort of mapping or listing

through things in but in react, it expects to have the key attribute.

And I actually don't have a key here, so I'll just pass index and it doesn't need quotes indexed like

that now because I floated, start and floated and I have to clear fix so things don't get all messed

up on the screen.

Give class name equals clear fix and then just to make things a little more attractive, put an error

in there which I can leave that way, or I can do it this way, which is a little less text.

So I'll do it this way.

Now let's clean some things up.

We don't have to use this dot state because we have a local variable movie.

We don't have to use this state here either.

And maybe between these two, I'll put in another row with a TD and I'll put it strong like the rest

of them description.

And then I'll write the description here,

which is just MoVida description.

OK, so this should work unless I have a typo somewhere.

Everything should display as expected.

So let's go back to our Web browser and let's list our movies.

And first of all, I'll go to American Psycho, which I know doesn't have any genres in the database.

And that looks right.

There is the title, There's the year it was published and this remains empty over here.

But if I go back to movies and choose one that has genres, it has to I believe they show up as drama

and crime here.

And that's pretty straightforward.

Now, obviously, if you wanted to put more information on this table, you're more than welcome to

do so because we have more information about the given movie in the database.

But this should make it clear to you how you can do that sort of thing.

All right, let's move on and we'll take care of categories next to display a list of categories from

the back end and then filter.

Ultimately, when you click on a category, name will show only the movies that match that particular

genre, which means I'm going to have to go at genres to every movie or at least to a few more.

But we'll get started on that in the next lecture.

## 52 - Getting started with Movies by Genre

So we want to start filtering by genre, and before we go too much further, you might have noticed

that at some points I'm calling the various types of movies mystery comedy, so on and so forth categories.

And in other cases I'm calling them genres.

And that's going to lead to confusion down the road.

And now is the time to pick a name.

So I'm going to call it genre, which means I need to do some refactoring.

So in my apt jazz, for example, I'm importing right here components, categories.

And what am I doing with that?

Let's find out right here.

I'm using it here and that's matching us.

So what I'm going to do is just start refactoring things to start with.

I'm actually not going to use this category jazz at any point, so I'm going to delete that right now

because we're just going to replace it with something better.

So let's delete that.

And that would cause us to have some JavaScript errors and some compilers.

But we'll take care of those in a few minutes.

So back in Apgar's right away, let's find the things that we're calling.

For example, we can get rid of this.

We're not using it anymore.

But let's find the things where we're calling things category.

I'm going to call this genres

and down a little bit further.

Do we have any other calls?

Yes, we have calls to categories.

So I'm are going to delete that entirely because we're not going to use that in here by category becomes

genres.

And that will link us to something called genres, so let's create a new component.

A new file, which I will call genres jazz, and I'll do I am Ursy

and I'll define a class export default class genres extends component and I'll just give it a render

function right now just so it'll compile and we'll say return each to genres like that.

OK, so that gives me something to import that I can use back in my apgar's.

So let's go up here and import genres from DOTZLER components, genres and then down here where I have

this problem right here instead of a category page.

We'll go to Jonas, OK, and I'm going to get rid of this entirely as well, because that's causing

me some grief.

And then down here, we don't need this function category anymore because that's actually what we're

going to do this time around, is build our category page, which will now be called our genre page.

All right.

So this should be compiling without errors at this point.

And it is.

So that's good.

So let's keep going.

Now, what do we want to do this time around?

What we want to have a page.

If we go back to our Web browser, this should be compiled.

Let's relabel that to John knew is something

genre's.

All right, so we want to be able to click on this and list our various genres here and then click on

an entry from the list it appears here and show all of the movies that are of that particular genre.

That's our goal.

So there's some things that have to happen before we can do that, but we can get started.

So let's go back to this new page for this new component called genres.

And it's obviously going to have a a state and it's also going to have a component did mount.

So let's get started with that.

Let's give it a state to start with.

And it'll be very similar to what we had before.

So I'll declare or define the variable state to be equal to it'll have genres which will initially be

an empty array.

We'll also have an is loaded, which will initially be false, and we'll also have an error which will

initially be no, OK, then we're going to need a component did mount and I can just grab that from

movies to start with.

So I'll copy this and change it as necessary.

And again, we're going to be defining a route that we go to and our back end that doesn't exist yet.

Who will fix up before too long.

So let's just paste in our component demand and we'll change this movie's two genres.

Aretha doesn't exist in our back end, but we'll get to that probably in the next lecture.

OK, so we're going to grab some information from this.

Obviously, we're probably going to have a rapper called Genres so we can change that.

This should stay the same.

This should stay the same.

Perfect.

So let's go down to a render function and see what we're going to have here.

Obviously, we're going to grab a few things from from the state and we'll extract them into local variables

consed and we'll say genres and is loaded and error all come from state.

This state and again will have well we can just return something here for right now and fix it up later

on so we could say return and I'll wrap it in a fragment

and above this will put the title

each to genres.

And below that will have a unordered list and then we'll just go over our numbers, genres of math and

we'll pass it on inside of that, we'll have an ally with a key equal to him ID, which will be grabbing

from the Jason and I will have a link like two equals and we're going to build our You URL and it will

be a you earlier that we're not handling right now, but we'll add that later on genre and then dollar

sign Emdur ID.

And then we close the opening link tag and put an end genre, name and genre underscore name, I think

it is Goodis and our link is closed.

Now we need to import link so we can actually use that import link from react Rotar.

I should fix the error down here and it does.

OK, so we're going to have to put in our logic to check for loading and check for error and all of

that in the render function.

But this is enough to get us started.

So when I load this, we should have no errors.

And when I go back to my Web browser, I should be able to click on that.

Once I wrote to it, did I wrote to it in The Abduh?

Yes, I think I did.

Yeah.

I'm linking to genres here and here.

I'm going to genres.

So everything compiled.

So I should be able to go back to my Web browser and not see anything in the genres page, but it shouldn't

give me an error and that's good.

All right.

So now next, we need to go modify our back end to actually give us the genres as we want them.

And we'll take care of that in the next lecture.

## 53 - Getting Genres from back end

So now that we have genres set up more or less or at least a good start on it in our front end, we

need to go to our back end and actually extract the appropriate genres from the database and send it

back to our front as Jason.

So we'll start in our movies, Dash Dbag Go, which is inside our models folder.

And at the very bottom, I'll just add a new function right here and I will call it funking.

That has the receiver of type TV model.

So we have access to the database and I'll call it genres and it's important that you export it by starting

with a capital letter as I did here.

And this will take no parameters and it will return a slice of pointers to genre and potentially an

error.

And again, we need to have our context.

So I'm just going to copy and paste the context from up here to save some typing

and come back down to my new function and paste that in and write my query.

And my query will be assigned the value of this, a very simple query, select ID genre name created

and updated from genres order by genre name and we'll read into the variable rows and potentially an

error from emerg query context and the context intended or query.

Check for their.

There is not equal to nil return nil, and the error and defer are rows closed to avoid resource leaks

to Feroze door close.

This should be an s close and will define a variable to hold our entire slice of genres will be a slice

of pointer to genre.

And now we go through a rose for Rose Dot next.

And to find a variable for each iteration, I'll call it G of type genre and error is a sign the value

of Rose starts again and we just scan into our variable ampersand ID that's a lowercase G ID and duplicate

that a few times.

The second one is genre name.

The third one is created

in.

The last one is updated.

And if there's an error, if error is not equal to nil, return nil and the error otherwise we simply

say genres equals append genres.

A reference to G and down here return genres.

And so we have this function now.

Now we need to go create a handler.

So back in our seemed fine movie handlers and down here, I'll put it right here, funk app of type

application has the receiver of type applicants a pointer to an application and I'll call this get all

genres and it's a handler.

So it takes w http response writer.

And a pointer to a request, oops.

OK, so this is dead simple, we simply say in the variable genres and potentially an error.

Go to the database.

So its app DOT models, DB dot get all genres.

I call it no, I didn't call it that, I called it genres, I'll try that again.

And this is a lower case, M..

And we check for an error, so I'll just copy this error check and paste it in here and then we send

the JSON back if there's no error or equals update.

Right, Jason?

And it expects to have the response or a status, the value we're converting to Jason and a wrapper,

which I'm going to call genres.

And again, we check for an error.

And the last thing we should have to do is to go to our roots file and to find a route for this.

So I'll put it right here, Rueter dot handler and it's htp dot method get.

And we're going to go to V1, slash one slash genre and we'll call the handler.

We just created a dot, get all genres.

Now I should be able to stop my application.

It's not running so I'll start it.

Go run dot cmd.

Slash API.

And it's starting, so I should be able to go to my Web browser, slash VH1 slash genres and see my

list of genres and there they are and everything is arranged exactly as it should be.

They're all in the correct alphabetical order.

Perfect.

So in the next lecture, we'll go back and modify our front end to actually consume this JSON and display

the list of genres.

## 54 - Displaying the list of Genres

So we're left off right about here last time, getting the Jason from our rest API that lists all of

the genres we have in our database, and you probably noticed something that I should have but did not,

which was rather sloppy on my part.

But here in this Jason feed and I'm looking at the the representation that Firefox gives me, but I

can look at the raw data.

So let's go to raw data and pretty print.

I do have a wrapper of genres which is good, and I have one entry for each of the genres, but I only

have the genre name and I need the ID here.

So let's go back and fix that.

And that's a really simple fix.

So back in my rest API, in my back in code written and go, I'm looking at the models file and here

for the type genre I'm telling go not to include ID in any JSON that it produces and I'll fix that just

by putting it there.

Instead, stop this and start this and go back to my web browser.

And when I reload this I should have ID and I do exactly as it should.

So that was an oversight on my part.

Now we want to go back to our front end.

So right now we're we do have a genres page, but it doesn't display anything at all.

We need to do now is to go and hook up the necessary calls to our back end to actually produce this

information.

So most of it is there already.

So let's go look at it and I'll go to my source code right here.

OK, so I'm looking at genres, jazz, and there's one other thing that I overlooked, which again,

was sloppy on my part.

And it's this on the fetch statement on the first, then the only got the second then sorry, where

I say set state, I'm actually telling it to take Jayson's genres, field the rapper and put it in movies.

And I'm not actually using that property anywhere.

There is no property movies.

That's supposed to be genres like that.

OK, now when I come down here I'm actually using genres but I'm not using is loaded or er and the easiest

way to do that is to copy code.

That's exactly the same.

So in movies just let's get these lines in the render function, copy them, go back to genres jazz,

paste them in here.

Right here.

Paste and give this a closing curly bracket and format it, and now when I go back to my Web browser,

there are other genres, so home genres and that works exactly as I hoped it would.

Now we can make this a little bit prettier and we'll do that later on.

But right now I'm only worried about functionality.

Adding some styling is it's really not that difficult and we'll do a bit later on.

But the next step, of course, is to have when I click on one of these, to go and display a page that

lists all of the movies for a given genre.

And right now we only have this one drama and this one crime with any information whatsoever.

But we need to look for the root slash genre, slash some ID and then display a component that renders

the movies that fall within a given genre.

And we'll get started on that in the next lecture.

## 55 - Getting movies by Genre

So now we need to get all of the movies for a given genre from the database and feed that to Jason and

we already have a function in movies to go and I'm looking at to go back and we already have a function

called All.

And if we're going to get all of the movies from the database for a given genre, it would seem to me

that that's almost identical to this function.

Now, I could, if I wanted to add a purist, might do this, create an entirely new function called

all by genre, or I can take advantage of this function and just modify a little bit.

And I do it really easily.

If I make the all function, take one parameter genre and that parameter is actually very atic, dot,

dot, dot.

So you can put zero or in essence as arguments in this.

So that means that all doesn't actually require that parameter, but it will take it if you supply it.

And then down here after the defer cancel, I just declare a new variable called where that's an empty

string.

And then I say if the Lenn of genre, which is the parameter of this function, takes a bit greater

than zero, that I have a parameter that's been supplied and I'll just redefine where to be equal to

format as print F and I should actually let Visual Studio Code import that for me there.

That's easier.

And give it the string where ID is in where id in open parentheses, select movie underscore ID from

movie movies, underscore genres where genre ID is equal two percent the and then I just substitute

genre zero and that will be whatever ID was supplied when we call this function.

So that gives me a where clause and then I just modify my main query to say format s per def.

And just here after from movies, I put a string in their string placeholder and then I'll move this

up to the previous line so it's not quite so ugly and put comma where and if no parameter is supplied,

it will be a substituting an empty string, four percent s.

But if an int is given to this all function when we call it, then it will put the word clause in there.

And that's really all I have to do at that point.

There's no other change required so we can now go and create a handler for this.

So there are movie handlers.

Don't go.

Let's write a new handler that says get all movies by genre so I can just copy this one as a starting

point

and paste it in here and call this.

Get all movies by genre, and it's a very simple handbook.

Oh, we're going to do is get our programs and we'll use the HTP.

Rutter's built in function to do that.

Grammes is a sign.

The value of it should be rever not.

Programs from context, and we handed our context or context.

Then we convert our genre ID, which we're going to get from a route that we'll make in a minute and

check for an error and that of a too high, and we want to get from Pyramus by name.

We'll look for genre ID and that's what we'll put in our route and we'll check for an error and we'll

just copy and paste an error from here and error check from here.

Then we get our movies and potentially an error from the database by calling app models DBI to all,

and this time we handed an entire genre.

Check for an error and then, right, our Jason, there is a sign the value of our arrows equal to Jason

and we handed W status

our variable movies and wrap it in movies and check for an error.

So that's it for the handler.

Now, let's go to the roots.

And right here, let's just duplicate this line and change it from movies to movies slash and we use

the same key we specified in our handler genre ID and we call get all movies by genre.

So if I start my application and start my application and go back to my Web browser and open a new tab

and say localhost four thousand V1 movies and give it some genre, I'd say give it to I don't even know

what genre that is.

And there is a syntax error at end of input.

Well, let's go find what we did wrong in our database function.

So it has to be this one

might want to close that parenthesis.

There we go and start this up again and go back and reload this.

We get The Shawshank Redemption.

So that seems to work the way that it should.

So our next step is to go in to write the necessary components and logic in our front end to display

the list of movies by genre.

And we'll get started on that in the next lecture.

## 56 - Displaying movies by Genre

So now that we're getting our movies by genre from our rest API, our back end, we should be able to

modify our front end to actually display those movies.

And we're going to do that by starting in Apte James and looking at our roots.

So down here in the switch section, we have our path for movies, ideas and for movies.

And we have one for genres, but we don't have one for genres ID.

So what I'm going to do is just copy this line and paste it right here and have that go to a nonexistent

component, which I'll call one genre.

And instead of movies, it will be genre ID.

OK, so now we need to create this component one genre, and I'll do that by going to my components

folder and creating a new file called One Genre Address, and I'll use my standard.

I am Arcy to get my first line in there.

And I want to use a fragment too.

So I'm able to put that in right now fragment.

And I'm going to also use a link so I have to import link from react router dom right there and I will

create my class.

I like to end these with semicolons.

I know you don't have to, but I just like to export the default class and I'll call it one genre and

it extends component.

And in there we're going to have to have state and state will be equal to.

And I'll have just like I did my movie file one called movies, which starts as an empty item, then

I'll have is loaded, which is set to fault's initially and I'll have er which is set to know initially

and I'm also going to have a component did mount.

Now I'm going to save some time, I'm going to go to movies Jairus and find that component did mount

and copy this entire thing and go back to one genre.

Actually there should be an array not an object and paste that in here.

There it is, and we're not going to go to VH1 movies, we're going to go to VH1 movies plus a parameter,

plus this props that match Perama ID and that will give us our movies once I put the slash in here.

There.

So that's correct.

And nothing else here needs to change at this point.

So can I go back to movies, jazz and just get the render part, copy this whole render function and

we'll just change it as required because they're very, very similar.

So one genre and after component did mount.

Let's paste in our render function and see what we have to change here, if anything.

Well, first of all, we can't assume that when we search for movies by genre that we're going to get

anything back.

And absolutely, that's true right now because I only have one movie with any genres at all.

So what I'm going to do here after my constant and where I populate these three variables from the state,

we'll just check to see if we have anything in movies, if not movies.

And I'll just set movies equal to an empty array, OK?

And of course, if I'm going to overwrite that value, I have to use a let here.

Not a constant.

That's another change.

Now, if there's an error display the error message, if it's not loaded yet, displayed the loading

message.

Otherwise it's pretty much the same.

The only difference is, rather than choose a movie.

Let's put genre and then I want to display the genre name.

And how am I going to do that?

Well, I'm not going to worry about that right now.

I'm going to just leave that alone and we'll take care of getting the genre name in the next lecture.

But for right now, is there anything else I want to change?

Well, let's just make this a little more attractive.

And I said I wouldn't do this until later, but let's get started right now.

I'll make this a div instead of a ULLE and I'll give it a class name using one of the bootstrap classes

list dash group.

And then I'll get rid of this div and change this to a closing div tag.

And inside here, instead of having an ally, I'll just get rid of that and get rid of that.

Instead I modify this link, I will modify it by giving it some bootstrap classes just to see how it

looks.

OK, so format everything and I'm going to add the class name and I'll do it right here.

Class name equals list dash group dash item list.

Dash group dash item dash action.

I shouldn't have to do anything else.

So now I have this.

Let's go back to our apgar's.

We're going to have to import one genre of course.

So here I'll just duplicate this and import one genre from one genre.

OK, no errors there, no errors anywhere.

Let's see how this works, see if we missed anything.

So my application is running.

So let's switch back to our Web browser and go to go watch movies and go to genres.

OK, and now let's look at one that I know.

There's something for drama, The Shawshank Redemption.

And there that's very nice.

I roll over it and it highlights and that link is going to an existing route for our movies.

So this should work, too.

And it does.

Now, if I go back to genres and choose comedy, I know I don't have anything in comedy.

It just displays an empty page, which is perfect.

So the only thing that's missing at this point for basic functionality is displaying the genre name

right here.

And we'll take care of that in the next lecture.

## 57 - Showing Genre name - an alternative to lifting state

So this time around, we want to be able to display the genre in the title when we're looking at the

one genre component right here on line fifty one of my code.

And you're probably thought to yourself when I said, how are we going to do that?

You probably thought he's going to go through lifting state once again.

And I could.

And that would be a perfectly valid way of solving this problem.

But I'm going to do it a little bit differently this time.

I only go to genres, jazz, and I'm going to pass some parameters in this link.

So rather than having the link as it is right now, I'm actually going to wrap everything in the two

section here in curly braces so it'll start and end with two curly braces.

And I'm going to name this entry in that half name, which is a required.

If you're going to do this, you have to use the word PADF name here and I'll put this on its own line

so it's a little easier to see.

Put a comma here and give myself two more lines.

OK, then I'm going to give it a second entry.

I'm going to pass an additional parameter so I can pass it like this.

Just pick a parameter name and we're looking for genre name.

So I'll call this genre names and its value will be a genre name because at this point when I'm rendering

the list of genres, I know the name, so I've added that value and I'll just format everything so it's

much easier to see.

And then back in one genre, I'll make some changes.

Here at the very top of this file, I'll add something to my state called genre names like that and

I'll make it empty by default.

And then down here in my component did malfunction.

I'll simply populate that value.

So we'll say this set state we're setting movies we're setting is loaded and now we're going to set

genre names.

And that will be equal to this props location, dot genre, the name that's available to us because

we're using the react router.

So I have a genre name populated in my state wants the component did mount and that means all I have

to do is come down here and put genre name and extract that up here in this genre and that should take

care of it.

So make sure your application is running.

Mine is and everything appears to have compiled.

Let's switch to our Web browser and let's click on genres and let's choose as genre.

I'll choose drama.

And there it is.

Drama shows up.

Let's try a different genre action, which has nothing in it, but should still display the genre,

any genre action.

And that is perfect.

OK, so now we are able to look at the home page, which has no exciting content but does render we're

able to look at the list of movies entirely.

And again, we're going to want to padget this at some point.

We'll put some more entries in our database so we can actually show five or ten movies per page or whatever

we want.

Genres are listing appropriately and we can look at entries within individual genres and then go directly

to that movie.

So this seems to be working really well.

The next step, and this will take a little while, is to allow us to manage our catalog to edit, add

or delete movies directly from this application.

And of course, that means at some point we're going to have to implement some kind of security protocol

that will require people to have the necessary rights to make changes to the catalog, because we don't

want random people to be making changes to our movie catalog.

And it'll take us a while to get there.

But we're now at the point where we can move on.

## 58 - Code clean up

So before we move on to the next major section there, just a little bit of code cleanup I want to do,

and it's just housekeeping and nothing here is terribly exciting, but it's all something that I would

typically do at this point in a project.

So if you look at our code right now or look at our product right now on the Web page, if I look at

all movies, they're listed just using a standard bulleted list.

But if I got a Shandra's and look at, say, crime is this nice functionality, I want to clean it up

so it looks a little bit better.

So let's go back to our code and we'll look at two files.

First one is one genre.

And here's where we have the div class Nomikos list groups.

So I'm going to copy that and I'll go back to movies and change this.

You will to a div with that class name and this will be a div as well, which means I don't need the

elai at all.

So let's get rid of that and instead we'll go link and we have to give it a key attribute.

So key is equal to my ID and then we just add class names.

So I'll go back to one genre and copy those class names like this back here back to movies and put the

class names right in there and I'll just format everything and that should fix it for that page.

So let's go back and look and look at movies that looks a little bit better, OK?

And genres do the same thing with genres.

So let's go back to genres and change that.

We'll just add the class name here like this.

And put our key in there as well

and ID and get rid of these allies because we don't need them anymore and change the top part to a div.

Copy this.

And go back to genres and replace that you will with this div tag and this becomes a closing tag

and I should do it, so it's formatted OK and go back and look at genres and they look much better this

way.

All right.

So that's just a little bit of cleanup now.

There's lots of other things we could do to make this more attractive.

And I'll leave that to a little bit later in the course.

But right now, it's time to move on to manage catalog.

And managing a catalog is actually going to require us to learn a bunch of new things.

First of all, how do you do forms in react?

How do you submit information and grab them, grab it back?

How can you send a post request from react to say, change the details of a movie or add a new movie

and how do we delete a movie.

And again, that's going to take a little while.

None of it is terribly difficult and everything you've learned so far will be helpful as we go through

the next section.

So let's move on.
