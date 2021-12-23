# 03 - Setting up Go Backend as a REST API

## 35 - Installing the necessary software

So it's time to start working on our go back end, and before we can do that, there's a little bit

of software we need to install and you may already have it installed and that's fine, but I just want

to be sure.

So obviously, the first thing we need is go itself.

So go to this URL, go langue and click on the big blue button here that says download go and you're

going to download the version for your operating system.

So Windows or Mac or Linux and the sources is available.

I don't encourage you to work with the source.

Download one of these versions.

Instead, it will make your life much simpler.

So I'm going to Mac, so I would click on Apple Mac OS and that would actually start the download,

which I'll cancel because I already have go installed.

But download the installer and run it.

And down on the page here a little bit, you can see specific instructions for each operating system.

So there's Linux, there's Mac and there's Windows.

So once that's installed and after it's installed, not before once it's done, start up your ID Visual

Studio code and you want to go to the code menu on a Mac and find preferences and look for extensions.

So the first thing we want is called go.

And as you can see, I have it here already.

So I would choose that and I would install it.

This would say install, so install that and once it's finished and not before it's finished, but once

that extension is installed, it command shift P on a Mac or control shift on windows to bring up this

pallet and look for that go Colen space install update tools and Shusett and then click the checkbox

at the top to make sure all of these are checked and click.

OK now I already have these installed so there's no need for me to install them again.

But wait for that to install.

It will take a little while, you'll have a little pane that opens up at the bottom of your window and

it'll tell you what it's done.

And once that's done, I usually restart my ID just to be sure that everything is working properly and

at that point you're ready to go.

So let's move on.

## 36 - Setting up the Go project

So now that we have go installed and visual studio code configured, let's start working on our backend

application and we'll start with a really simple version of it, which will improve as time goes on.

So I have Visual Studio code open and I have an empty window here with no folder open.

So let me close this and open a folder and I'll create a folder in visual studio code projects and I'll

just call it back and dash out and I'll open this folder now because we're starting this as a modern

go application.

We're going to use go modules.

So I open my terminal in visual studio code and I type go on in it and I can call it whatever I want,

but I'll just call it back end, which is sufficient for our purposes.

And that creates, as you can see over here, go mode file.

Now I'm going to create a folder to hold the main part of my application.

And the convention for a lot of people, myself included, is to create a folder at the root level of

your app called CMD, and then inside of that create a folder because we're writing a rest API.

I'm going to create an API folder.

If this was a standard Web application, I would probably call that folder web, but I'll call it apart.

And inside of the API folder, I'll create a mango file and there's a mango file.

So it requires a package declaration and I'll just call it main.

And the very first thing I'll do is define a constant.

And this is going to be a constant that just keeps track of the version of our application, which will

send back as part of our response in certain cases.

And that will be equal to one point zero point zero, just a string.

OK, the second thing I'll do is create a type and I'm going to call this config and it will hold the

application config just a struct.

It's a type struct and it has two members port which is in it.

What port will our application listen on and I'll call this one on and it will just be a string and

that's the application environment.

So it'll be something like development or production or staging, whatever it may be.

I'm going to use that so I can pass information around to various parts of my application, not create

my main function function, and we'll make a really simple one.

So right away I'll create a variable called CFG, which is of type config, which we defined just up

there on line five.

And now I'm going to assume that when I start this application, I'll be reading things like the port

and the environment from the command line as a flag part of the application.

So I'll use the built in flag package, which is part of go.

And I'll first of all, I'll read for an ENT, so I'll find the eyes right here.

I'll just type into.

Inver, that's the one I'm looking for, so I'm going to read from the command line some integer value

and I will store it in my config variable that I just created in the member port.

And it's going to be called port on the command line.

That's the flag will be passing.

I'll give it a default value of, say, 4000.

OK, that's the port we're going to be listening on and a description server port to listen on.

This way, I don't need to specify a command line flag of four thousand.

That will be the default and that's what we'll use for development.

Now, Fleg, let's get a string stringer and this will give me my environment, which I'll read into

CFG and on and the flag will be on and the default will be development and we'll give it a description

application environment and we'll say the default so that you should choose either development or production.

We'll just go with two for now, OK.

And now that I have those defined, of course I have to pass the flags flying past.

And if you've done any work with Go, you've almost certainly encountered that.

Now we'll just say format, print line just to make sure everything is working, running.

So I'll open my command line and now I want to execute this so I can just type it because I'm in the

root level of my application.

Go run cmd APIs mango and it should just print running to the window and exit and it does.

OK, so that works.

No errors so far.

Now we actually want to serve some information.

All I want to accomplish this time around is just to get a basic web server running and serve some kind

of JSON content to the browser.

So we'll make this as simple as we possibly can.

So we'll create an HTP dog handler, func a handler function.

And this is a function that will listen for specific URLs passed to our server.

So it requires a few arguments.

The first one it requires is A is the path it's going to listen to and I'll just make it status for

now.

And then it requires a handler func.

Well, I can inline that.

So we'll just say func and it requires the argument of w htp response writer, as all handlers do,

and there is a pointer to http request.

And inside of that, all I really want to do, apparently I'm missing something here.

I have an R here hdb handle function.

That's better.

OK, so what I'm going to do here now just to make sure this actually works, it's just say format dot

f print.

I'll just print right to my response writer w and I'll just give it status just to make sure this actually

works.

OK, and now below that let's start a web server.

So error is assigned the value of HDB, don't listen and serve and we're going to listen to the part

that was passed to us so I can use format dot print F because it expects a string and I'll give it the

colon.

It wants then an integer value in that.

I'll put cfg dot port and the second argument I'll just hand it nil.

So I'm not passing any data to it and check for the error.

If error is not equal to nil log print line error.

OK, so now I should be able to run this, open a web browser and go to localhost port for thousands

of status.

So let's try it.

Go run cmd api slash mean don't go.

OK, so it's running.

So let's go to a web browser and let's go to localhost four thousand slash status.

OK, so it's working now unfortunately that's just plain text.

And what I actually want to send is Jason.

So let's go back to our web router.

I'd stop our web server by hitting control, keep or close this and let's make this a little more useful.

And this is just a starting point.

We're going to be changing this.

First thing I'll do, create a type up here, because I want to actually send some formatted JSON about

the current status of the server to whoever requested it.

So create a type which I'll call app status.

And it's a struct that has a few members and one I'm going to put in there, first of all, is status

and that will be a string, OK, and then I'll put in environment and then will also be a string and

version which will also be a string.

And you will notice that I.

Exported all of these members.

They begin with a capital letter and you need to do that now for Mathes.

OK.

And down here in my Handal funk, I'll get rid of this line because I'm not going to use that anymore.

I want to return some Jason.

So what I'm going to do is create a variable called current status, and it is of type AB status.

And I'll populate its members as follows status.

I'll just make it available.

A string and environment.

I can read that right from my config variable cfg n.v. and finally version.

Well, that can come from the constant version that I defined way up on line 10.

OK, so I have this variable defined and it has those three members in it.

Now I want to convert that variable which is a struct into JSON and I'm going to do that by calling

Jason Marshall in debt.

OK, so the result will be stored in a variable called Joce.

I will check for an error and I'll call from the builtin package in the standard library and JSM dot

mercial indent.

Now, this is the one we'd be using in production.

This is the one that we're using right now because it's more readable.

So Marshall Indebt requires a few arguments.

First of all, what do you want to convert into JSON on its current status?

And do you want to have any kind of prefix?

An empty string is good.

And how much do you want to indent it?

I'll just indented one tab by using backslash T, then I'll check for an error.

If error is not equal to nil, all do right now is logged print line the error.

Otherwise I want to write that Jass right to the browser as Jason so I already have it in JSON format

is stored in the variable jass but I need to set the header header set and I want to set the content

type with a capital C going to capital T just like that to application.

Jason and I also want to send some kind of status and because this I'm assuming this is going to work,

I'll just write the header and what I'm going to write is http dot status, OK, which is two hundred

an integer and that's a built in constant in the HDB package.

And now I just write mine, just write to my response writer w write us.

OK, so let's try this again.

Let's run our application, go run cmd api Mingo and it's running.

Let's go back to our web browser and just reload this page.

And now I'm using Firefox so we get a nicely formatted representation of the content of that.

Jason but if I click on Raw Data you'll see it actually is a JSON file, but I don't really like the

status, environment and version to begin with capital letters, although I want those to be lowercase

letters.

How am I going to do that?

I'm going to go back to my JSON here and then switch back to my ID.

And what I'm going to do is stop this, OK?

And back up here where tried to find app status.

I'm going to give some extra information to each of these in tactics.

I'll say when you're rendering this as Jason JSM Colon, then in double quotes, call this status.

And for the next line, when you're rendering this as Jason in quotes, double quotes, I'll call it

environment with a lowercase E and close that double quote, which is something I always forget to do.

And finally, for the last one, when you're rendering this as Jason Kolu double quotes version, OK,

let's run our application again.

And go back to our Web browser and reload this page and now you'll see that everything looks much nicer

and that's why I control what's going to show up in this Jason feet.

Now, there's a couple of other things you can do when you're rendering Jason this way.

But when we go a little bit further, we'll cover those.

But for right now, we have successfully created one path, which is status that's handled by this handler,

Funk, and we've actually returned a response as valid, Jason.

So it's a good start.

Now, clearly, there's a lot more to do, but we'll get started on that in the next lecture.

## 37 - Installing a router and creating better handlers

So we have a very basic application up and running right now.

It listens on Port 4000 on the localhost and listens for one route, which you can see right here in

main go status.

And it returns to Jason that we're actually hard coding right now.

So it's not doing anything functional, but it is at least working and it's a good start.

But right now, we're using the rooting functionality that's built right into the standard library and

it works.

And right here on line three, we have a function within a function in our main func.

We have HDB handle func, which listens for status and returns the necessary JSON and that works.

And if we really wanted to, we could build our entire application this way.

But as you might expect, things will quickly get out of hand once we have more than a handful of routes,

legal status, whatever it may be, once we get just more than a few of those.

This application would become rather unwieldy and really hard to support and maintain.

So what we're going to do is we're going to install a router, a router, a third party package.

And I'm going to go to my Web browser right now.

And the one we're going to install is right here, Julians Schmitt's HTP router.

And it's a very good router, very fast.

And it's really well suited for backend APIs, rest APIs like we're building now.

The one thing to be aware of, if you scroll down the page here, this paragraph that says only explicit

matches, this router only listens for explicit you or else.

So you can't have roots that match more than one.

You are out now.

Other routers like the key router or the router, they actually will allow you to have pattern matching

rules in your routes file.

And Julian Schmitt's HDB router does not.

However, since we're not building a Web application that builds Web pages, we're building a rest API.

This is almost perfectly suited for our purposes.

So we're going to import this and I want to import it by copying everything in the euro, starting with

GitHub dot com right until the very end, not the HDTV bit.

So I copy that and I go back to my idea.

Let's install it.

So open terminal window and type, go get Dirceu and then paste in that you are out and I return and

that will install it and it modifies our go mod file and there it is.

So it's now available to us.

So we already use this.

We're going to make some changes to it the way that our application is organized.

And the first thing I'm going to do is create a new file.

Actually a couple of new files not in my command folder, but in my command API folder.

So.

All right, click and create a new file.

I'm going to call this one rootstock.

And just so it doesn't throw an error, I'll give it its package.

Declaration package.

Meaning no, create another one that is specific for our handler, our status handler.

I'll create a new file and I'll call this one status handler, Dongo.

And again, I'll give it just its package declaration for now.

Package me.

OK, back in our main file, let's make some changes to the way that we're actually organizing things.

So the first thing I want to do is I want to share certain kinds of information between my various handlers

in my in my application.

And I can do that with global variables.

But that's kind of awkward.

So let's take advantage of Go's receivers and let's create a new type, which I'll create right here

on line twenty.

After line twenty two, we'll create a new type called application and this will hold our application

configuration and it's just a struct.

And right now we're going to put a couple of things in there.

We'll put it in the config, which we already have.

It's the type config which we defined on line thirteen.

We'll also put centralized logging it.

So a logger is the name of the field and it's a pointer to the built in package log logger from the

standard library.

OK, so that's my type.

Now a little bit further down here in the main function after line thirty four in my code where I finish

passing my flags, let's create a longer logger is assigned the value of and longer.

It's just a variable name I'm choosing again from the built in long package will say logged on to and

we're going to write just for, for now.

We'll write to OS standard and of course you can write that to a file.

You can do whatever you want with it, but we're going to go with ostinato and I hate return so it should

do the import for me and now specify the log format.

We're going to have log eight right there and then a pipe and log time just so we can know the date

and time when we're writing to the log.

So that creates our logger file.

Now let's create a variable, which I'll just call out because it's our application config.

It's a reference to application, the type we just define that again application, the type we just

defined unless populates members.

So config is obviously the CFG variable, which we defined on line thirty one and we populated it's

two fields with the flags on lines thirty three and thirty four.

So that's config.

And our logger is just the leader we just created back on line thirty seven.

So now I have this application, this variable application and I'm going to use that as a receiver on

my various other parts of my application.

So let's go over to our roots file and create a new function in here.

And this is where we're going to take advantage of Julian Julian Schmitt's HDB router.

OK, so I'll create a function func and it will have the receiver of app a pointer to application.

I'm going to call this Roots and this is where we'll put our application roots and it returns one thing.

A pointer to Julian Schmidt, Tretter HDB router, the router, OK.

And of course this type satisfies the necessary interface to be a monks' or a server mux.

So we can use that with no changes to any other parts of our application.

So it will create a new router, which I'll call router.

It's a sign the value of a router new and that takes new arguments and we'll create nothing.

Right now.

We're going to add our roots in here in a moment, so we'll just return the router.

OK, so that is our first application with one import and that's all that we need at this point.

So let's go back to our main dock, go and look at how we're going to make this entire thing work.

So we have this handler funk right now.

And what I want to do is create in status handler.

I want to duplicate this functionality, but tie it to my application config by creating a handler with

the receiver.

And I'm going to do that in status handler Dutko.

So inside of this, let's create a function with the receiver of app, which is a type pointer to application,

and we'll call this status handler and it is a handler.

So it requires the same two arguments that every handler Engo does.

First of all, I'll need my variable W and that's an active response writer and it also requires a pointer

to an HDB request, which I'll call our pointer to request our Q there we are.

And inside of that I'm going to duplicate the same functionality that I have back in Maine.

Go so I can come in here and just copy all of this or copy it and go back to my status handler ago and

paste it in.

OK, now there's a couple of changes I have to make to start with.

I can't use CFG.

Instead I'll use my receiver, which is a simple enough app config on and that's a string.

So that's fine.

This is a constant at the package level, so I can use that and now I just need to import my JSON.

So I'll just cheat by saying Jason Mercial indent and return and let it do the imports for me.

And again, the last change for our handler is this log.

We don't use that one.

We use Apte logger, print one there.

So here we're using a receiver and here we're using a receiver and this receiver will become very,

very useful as we continue building our application.

It's just an easy way to share things among the various components of our application in a clean and

logical fashion.

So I have this handler.

Let me go back to my roots file now and let's actually where's my roots?

They're there.

Let's actually do something with this roots file rather than just returning an empty root.

So what we'll do is define our only application route that we have right now.

And we do it like this router, which is the variable we defined online.

So we give it a handle func.

Handler Funk, and we need to specify the method.

Well, this is going to be a get request, so we use HDB dot method, get right there and that's just

a constant defined in the standard library.

We're going to listen for status and we want to go to Abdul statis handler with no parentheses after

it, and that should give us our route.

So we're not calling this route's function yet.

So let's go back to our main dog, go and make the last change we need to make.

First of all, I can get rid of all of this because we're not using it anymore.

I'll change that.

What's on line forty four to give you a more meaningful method shortly.

And we don't need that in line to handle funk because we have a route's file.

We need to change the way that we're starting our Web server.

So I'll delete this line.

And what I'm going to do is to find a new variable right here, which I'll call serve, and that's assigned

the value of a reference to from the HTTP package server.

And inside of that, we need a few members.

The first one is what you want to listen to our address exactly the same as we did before.

We're going to listen to localhost for four thousand.

So format S, print F and my string is Colen percent D for an integer and replace that with the value

from config part.

The second thing is what handler do you want to use.

Here's what we call our roots roots.

The third thing is how long do you want to time out for an idle connection so we'll sit idle, time

out and we'll make that one minute time dot minute right from the time package in the standard library

and ah, read time out.

We should give it a sensible value there.

If you can't read things within ten seconds, something's gone wrong.

So 10 times timed second and the right time out.

I'll make that one a little bit longer.

Right time out of sensible value is 30 seconds, 30 times time second.

So we've created that variable.

And now down here, first of all, let's print something out to the log so we'll use our longer, longer

print line.

And I'll just print a message that says Starting server on port and I'll just put CFG port and then

we'll start the server.

Error is a sign the value of serve the variable we defined on line forty six and we take advantage of

one of its built in methods.

Listen answer.

And that should be enough to start this application.

So first of all, let's make sure there's no errors.

Yes.

I'm still importing encoding JSON or encoding JSON.

I don't need that, so I'll just save and that should clean up my imports.

And it did.

So no errors here, no errors in roots and no errors in status.

Handler So I should be able to start my application.

So because we have multiple go files in the command API directory, we can't just say down here in the

terminal go run cmd slash API may not go because we actually have more than one goffer.

It's on a Mac.

You do it like this.

Go run cmd slash api slash start uko and that will start up in the Mac.

So let's run it just to make sure it runs and it does.

So I'll quit out of it.

If you're on windows it's a little bit different.

You type go run dot for the current directory, slash sumed slash API slash dot and that's how you run

it on windows assuming you're using the built in standard windows terminal.

If you're using power shell you probably already know how to do this.

I don't because I don't use power shell very often and if you're using git bash you would run it exactly

the same as you do on a Mac.

But if you're just using the standard Windows command prompt, you would do it.

As I show you here, go run dot slash cmd slash apps slash done.

But of course I'm on a Mac.

So go back to this first.

Let's make sure it actually works.

So starting server in four thousand.

So far so good.

Let's go back to our web browser and let's go back to our standard localhost column.

Four thousand slash status and return.

And this should work and it does.

OK, so that works really well.

Now there's a lot of work to do yet, but what we've done so far in this lecture is we've simplified

things immensely.

So I'm going to hide this and just go through it one more time.

First of all, we created a new type application and this is where we'll put the information we need

to share with our handlers and other components of our application.

And that will become incredibly useful as this course goes on.

Then we actually down here separate our routes so we don't have to hard code them into our main function.

Instead, we're putting them all over.

We can put each handler in its own file or group the handlers into a single file when they're logically

related.

And we have a route's file that makes our routes ever so easy to manage.

All right.

That's enough for this time around.

Let's move on in the next lecture and start making this application a little more functional.

## 38 - Models

So what I want to do this time around is to create some types, to hold the kind of data we're going

to be serving.

And as you know, we're serving some movies, so let's create some types to hold a movie.

And then we can actually establish a single route to serve one movie as Jason and another route to serve

a collection of movies.

So we'll do that.

I'm going to do this in a new package.

I'm going to set my hopes up by creating a new folder here, which I'll call Mobbles, because these

will eventually become database models.

And I'll create a new file in the models folder called Models Don't Go Great, and I'll give it its

package decoration models and I'm going to set up three types.

So the first one will be a type movie and it's just a struct.

Then we'll set up another one called type genre, which is also a struct, and finally I'll set up one

for type.

I'll call it movie genre, which is also a struct.

And let's put some field in there so we might change these later on.

But right now let's just get something so we can serve some content.

And I'll call this ID and it will be an INT and I'll give it its Jason description of it and then I'll

create the title for the movie, which would be a string and it will be JSON title like that, and then

I'll give it its description, which will also be a string JSON description.

And then I'll give it the year it was released, which will be an event and in Jason, I'll call that

year.

Then the actual release date, just to be specific, release date, even though I can get the year from

the release date, I like to have things so I can search very quickly on one end column.

Release date will be timed on time and let the import happen time to time there.

But the important for me, Jason, release date.

And I'll also give it, say, the runtime.

Which will be an event and that will be Jason runtime

and then I'll give it a rating in case I want to use star rating at some point, which I can end and

Jason rating.

So that might be one star, five stars, whatever.

It's going to be the empty AAA rating, which is important for a lot of people, which will just be

a string in our case and I'll call it Jason moppy a rating, the Motion Picture Association of America,

I think it stands for and then I'll give it credit, which will be time of time.

Jason created and updated at.

OK, and finally, one movie can have multiple genres, so I'll put in the movie genre, which will

be a slice of movie genre.

And in Jason right now, I'm going to just say ignore that, because I don't want to deal with that

at the moment.

I want to get one thing working first.

And I have a typo here.

They're.

OK, so for genre, I'll just have a list of genres just to look up table indie and Jason.

It is definitely going to have a created an update out, so I'll copy those and paste that right here

and then I know I'm going to have the same columns in movie genre, so I'll copy that and pasted in

here to save some typing.

All right.

What else do I want in genre?

Just a genre name, which would be a string.

And in Jason it will be a genre.

And down here, that's all I need for genre.

I'll have the movie ID, which is an aunt and Jason movie ID and then I'll duplicate that and make that

genre ID and here I'll call it genre ID.

And then finally, just because I'm probably going to use actually put the genre in there as well or

not to type genre and in Jason I'll just call it genre like that.

OK, let's format this.

Don't see any errors.

That looks good.

So there is our models and I can import that wherever I need it and use it and I stick it in its own

package just because it's clean and I like to do it that way.

If you wanted to put these in the main package, you feel free to do so, but I'm going to leave them

in their own package.

So we have that defined in.

The next step will be to actually set up a route that allows us to serve one movie.

And we'll just hardcoded the data right now.

We'll be looking it up from the database before too long.

But I want to set up a new route that grabs one movie or serves one movie as JSON, and we'll get started

on that in the next lecture.

## 39 - Setting up a simple API route

So we have our model set up, and this time around, I'd like to set up a new application route that

will allow us to display a single movie.

And the first thing I'm going to do is I had a little typographical error here.

This is this should not be updated at with a capital A..

It should be updated at like that.

And of course, naturally, I copied and pasted it twice.

So I have to fix it in two more places.

All right.

Just cosmetic, but it makes me feel better.

So what I'm going to do is in my command API folder, I'll create a new file, which I'll call movie

dash handlers Dutko and I'll put all of my movie handlers in here and we'll give it its package declaration,

which I'll call Main, of course, and I'll define two functions.

The first one, we'll get one movie and the second one, which will be a stub for now, will get all

of our movies so func.

And it has the receiver of app, which is a pointer to application and I'll call this one get one movie.

And because it's the handler, it takes a response operator and a request as arguments.

So w the response writer.

And there will be a pointer to a request, and now I'll just copy this and give it a different name

and this will be get all movies.

OK, then let's go back to our Roots file and let's set up roots that go to those dub handlers.

So here I'll say Rueter Handler Funk and this will be a get request HGP method get and I'm going to

go to not movie and then the I.D. And this is how we we put a placeholder for a numeric or any kind

of ID that will be different based upon the route we're matching instead of on a prepend.

The whole thing with one, because this is a convention when you're building a rest API occasionally

or API has to change.

And if it introduces a breaking change, typically you would go from version one to version two.

So I'll use V one and we'll root this too, after I get one movie right there.

Now, I'll duplicate that route and change this to movies and read that to get all movies.

OK, so now we have two new routes, but of course the handlers aren't doing anything.

So let's go back to movie handlers and we're only going to take care of this one right now.

So what we're going to do, first of all, is you'll notice back in our routes we have Colen ID.

And of course, this is the syntax we're allowed to use because we're using Julian Schmitt's, HDTV

Rohter.

So let's go see how we're going to work with that.

Back in our movie handlers, the very first thing I'm going to do is to try to figure out what that

ideas.

And I'll do that by creating a variable called programs like that programs is assigned the value of.

And I'll call one of the functions that's available to us from HTTP Rotar.

And the method we're looking for is called programs from context right there.

And I chose it from the list, so it imported for me.

And that takes one argument, which is simply the context which we get from our context.

So that gives me a program's variable.

And what can I do with that?

Well, quite a bit, actually.

We can take that Perram variable and look for a specific one by its key.

So I'll say it and error are assigned the value of Sterkel because it's coming in as a string.

So I need to convert it to an end Sterkel with a tour and I will go for programs dot by name.

And the idea I'm looking for or the name I'm looking for is ID like that and I'll check for an error.

If error is not equal to nil, then for now I'll just say after lugger dot print line and I'll write

a new error.

Errors new and I'll just say invalid ID parameter because we only want numeric values to come in that

way.

Now in a moment we'll actually do something else here and stop execution, but for now we'll just leave

it like that.

OK, so just to make sure this is working, let's say app logger dot print line ID is and will print

out the ID.

So let's start our application.

If it's running, we'll stop it, see if it's running.

It is.

So I'll stop it all started.

See if it compiles and it does.

Let's go back to our web browser.

And status should still work.

So let's refresh this.

It does.

Now let's go v one movie and I'll give you the idea of 10 and I go look at my log file in my ID and

it says it is perfect.

We've got that part working on right now.

What do we actually want to do?

We want to do more than print out the I.D. So I'll stop this.

I just want to see if I can get one movie to show up as Jason.

So what I'll do is I'll create a movie from that model's movie type that we define a little while ago.

And I'll put some information in here.

So I'd and I'll call that ID the one we've got from our parameters.

We'll give it a title no matter what the ideas for now, it'll be given the same title and other information.

So I'll just put placeholder information in some movie is the title, description is some description

and we'll go with year say twenty, twenty one and release date will be a new date.

So time don't date and we'll give it a year.

Twenty, twenty one, first month, first day and first hour and then zero zero zero and the local time

zone that'll get us working.

OK, so there's a release date and we also want to run time and I think that's one hundred minutes should

do it.

I have an error here.

What is this.

Models or didn't important for us.

OK, let's try that again.

Models dot movie and down here we'll go with rating, say, five stars.

And the next line, ampe, AAA rating, we'll go with PG 13 and finally created a will be timed now

and updated will be timed out now and get rid of my extra comma.

OK, now we have a movie so we can, as we did before, convert that to Jason and send it back.

But this seems like it's the second time we're having to create Jason and send it to the browser by

hand.

And that seems really inefficient.

So let's create a function to do that for us.

What I'll do is in my command API folder, we'll create a new file.

I'll just call it utilities, Dongo.

And inside of that we put my package declaration.

What I want to do here is to write a simple function that will write Jason to the browser, so what

sort of function?

Func and we'll give it the receiver app, start on application.

And this will just save me from duplicating code and I'll just call it right, Jason.

Like that and I'll make a few arguments.

First of all, you need a response writer.

Obviously, because you have to have somewhere to write it, then you need a status code, which will

be an hand and you need the data to convert to Jason.

We don't know what that's going to be, so we'll just call it interface like that.

And finally, I'm going to give it something called wrap, which is a string, and you'll see why in

a moment.

So this wrap is going to be used to wrap our Jason with some kind of key, something that describes

the kind of content that's coming out of there.

And the easiest way to do that is simply to create a wrapper, which is a variable.

And it's going to be nothing more than a map with the key of string in any kind of content you want

interface.

So that's my variable now when I call this function, what I'm going to do is wrap the data that's passed

as the parameter in line five.

I'm going to wrap that with something.

So I'll say wrapper.

And the key will be wrap is equal to data.

And now I just do what I did before.

I'll say Jass an error or assign the value of Jason Marshall.

And this time I want Mercial indent and I'm going to Marshall Wrapper and I'll check for an error.

If error is not equal to nil, then I'll say return an error, because this function really should return

an error like that there.

Otherwise I'll just send it off as Jason.

So I'll say that header dot set and this is going to be a content type just like before application

JSON and then I'll write my status right header and that will be status which we receive as a parameter

and then we write drafts and if everything worked, will return nil.

So now I have a function that I can call to send Jason back to the end user.

So let's go back to our handlers and let's try using that function.

Really, all I'm going to have to do here is to say on this line.

Error equals arthritis and hand, it is required parameters, which would be a response rider.

This would be status, OK, HDB, not status, OK,

our data, which in this case is just moving.

And then what do we want to wrap it in a movie like that.

So let's just try this bit and see if it works.

So I run my application.

It started let's go back to our Web browser and try going to be one slash movie, slash 10 and there

is my content, absolutely perfect.

You notice that it's wrapped in the key we used as movies.

So if I look at the raw data, it comes as a single line because I use Jason Marshall, not Jason Marshall

and Dent, but because I'm in Firefox and changes to print and it looks a little better.

So everything is wrapped with this key movie and all that does is make the Jason more readable.

So that is one way to get information sent to our end user.

Of course, if I change this 10 to 20, I should get the same content.

But now the idea is 20.

So that seems to be working really well.

So the next thing I want to do is to do more in this part of the handler.

So if I go back and look at our code, if there's an error, for example, if I put a there instead

of a number, then it should stop.

It should generate a good, nicely formatted JSON error message and send it back to the user with the

key of error at the top level.

So we'll take care of that in the next lecture.

## 40 - Improved error handling

So as I suggested last time, there's still a bit of a problem with our handler right now, and that's

found right here on line 16 and following.

So I'm expecting my parameter to come in numeric formats.

So movie one, for example.

And if I use a movie, something that's not a number, well, things don't work the way you would expect.

So let's go to our Web browser and I'll change this from slash 20 to slash Alpha.

And you see it still throws back JSON with an idea of zero, which is not useful at all because when

we're actually querying the database, none of this information will be there.

Since we are not looking up things in the database or we won't be by an alpha numeric value.

We're looking up only by an integer, so we need to fix that and that's pretty easy to do.

So I'm going to go back to my code and I'm going to open my utility stock go file and I'll add a new

function right here and I'll call it func and it has the application receiver.

So our application and I'll call it error, Jason, and it will take a couple of parameters w which

will be a response rate.

So we have somewhere to write to.

And it will also take error, which is an error, and it doesn't return anything.

And all we're going to do here is to find a type specific to this function, which I'll call type Jason

error.

And it will be a struct and it will take one message, which is a string.

And in JSON, I'll call that not surprisingly, and then I'll create an error variable.

The error is assigned the value of adjacent error and its message will be the error that we received

and we'll call the error method on it, which gives us a string value.

And then I'll simply write this act right, Jason, and give it the W it's expecting and give it I'll

call it HBP status bed request and we might come back and change this later.

But this will work for now.

The error variable itself and my wrapper will be error.

And all I have to do next is back in my handler right after I write to the log, I'll call after errors

and handed my response writer and the error I just generated.

And of course, I don't want to go any further, so I'll return now.

If we stop our application, if it's running, let's see if there's one marked as go here there is.

So I'll stop that and clear the screen and run it again and go back to my web browser and go to localhost

four thousand slash v1 slash movie slash one.

This should give me a valid JSON and it does.

But if I change that to say made up instead of one, now I get a nicely formatted error message.

OK, parsing beta invalid syntax.

And if I look at the raw data and say pretty print, that's how it looks.

And that is a much more useful response.

Now, as I said, we might go back and change that handler to allow us to specify a specific status

response instead of just bad request.

But this is sufficient for our purposes right now.

So let's move on.

## 41 - Creating the database

So in this lesson, I'd like to create a database that we can connect our arrest back end to and on

the course resources for this lecture, you will find a file called Go Underscore Movies, Dot School,

and that's a database dump from a Postgres database that gives us a very, very simple database that

we can work with.

So download that and put it somewhere that you can find it and open a command prompt.

So I downloaded mine to a folder called Escudo, and inside of that I have the file go underscore movies,

DOT sequel, and I want to import that into a database.

So I first I need to start up my database client and I'm using postal code.

But you can use whatever client you want, Deaver or whatever postgrads client you're happy with or

familiar with.

Connect to your local host and create a database and I'll call mine, go underscore movies and that's

done.

So now I'll go back to my command prompt and type Pascual Dash D, the name of the database I just created,

which I called Go underscore movies and dash F for file and that would be go movies sequel, go underscore

movies, start school and that will create the database.

Now if I go back to my client and in my case again, I'm on post to go and connect to my local host,

I should be able to look at that database right here go movies.

And as you can see, it's super simple.

There are three tables, one for movies, one for genres, and one that links movies to one or more

genres.

Very straightforward.

So another that's created.

We can actually write some code to connect our arrest back, end to the database and we'll get started

on that in the next election.

## 42 - Creating our connection pool and connecting to the database

So now that we have an actual database, let's write the code necessary to connect our application to

that database and that's all will accomplish this time around, is just to connect and make sure that

our connection actually works.

And it's pretty straightforward to do that.

Now, the first thing we need to do, of course, is to install the necessary driver for PostgreSQL.

And there are a couple to choose from and there's a lot of contention about which one to use.

So to be honest, it wouldn't matter which one I picked.

I'd have some people saying you should have used the other one and some people saying you should use

the one that we're going to use.

So at least half of you will be happy.

So I'm going to use a really standard one that's been around for a long time.

So open your terminal window and we'll go get Deshu and we're going to get github.

Com slash lib slash p q and to specify the actual version, just to make sure that everyone is working

with the same version, will go with the one point ten point zero.

OK, so go get that and it we'll go get it and add it to our go mod file and install it and now it's

available to us.

So I'll close my terminal window and I'm looking at Mango and I'm going to add a few imports up here.

So after a right before Fleg, I'll put in context, which is built into the standard library and that's

something we'll be using momentarily.

I'm also going to install or import database slash sequel because we're connecting to a single database

at the very bottom.

I'm going to put a blank line here and then type in the blank identifier the underscore and then GitHub

dot com

p q OK, so it's showing some errors right now because we're not using context or database secobarbital,

but we'll be doing that shortly.

So let's scroll down a bit and let's go to our config variable or type where we're defining our config.

And what I'm going to put in here is one more thing we'll put in the DB, which is a struct, and we'll

be modifying this before too long.

And that only has one field right now or one member DSN string.

And that's our connection string.

That's where we're going to use to connect to our database.

So down a little bit further, we need to get that connection straight and yours will be different than

mine.

But it'll be really easy to follow what I'm doing now at the moment, and this is a short term solution.

I'm going to read the connection string from a command flag when we start up and I'm going to put a

default connection string that will allow me to run the application without specifying any command flags.

So we use a flag and we're going to get a string bar and we're going to call this or write it into CFG.

DB DSN, which we just added to the type, OK, and we give it it's it's flag name, which we'll just

call it DSF and then we'll give it a default value.

Now, this is where yours will be different than mine.

I'm going to take Postgres and so should you call and then whatever username you can you use to connect

to your database.

So if mine's my initials and if you have a password, you would go Colen, whatever your password is.

But mine is set up so that I don't actually don't have to use a password.

OK, then at and your host in our case it's localhost because our database is on the development machine

or whatever.

You named your database.

Mine's called Go Movies and then it questionmark just to be safe.

You might not need this, but it's not going to hurt it.

SSL mode equals disable just to tell the connector that we're not using SSL to connect to this database.

And the last thing is the postgrads connection string.

OK, so that will get the DSN into our configuration.

All right, let's go a little bit further to the bottom of the file.

We're going to add a new function at the bottom of the file.

And I'm going to call this function just open DB and it will take one hour, one parameter CFG and it's

a type config and it will return to possible things, a pointer to SQL DB, which will be a connection

pool.

All of our connections we're going to use to connect to the database and potentially an error.

So first thing we'll do is open a connection to the database and you do that as follows.

DB is one return value check for an error there.

Assign the value of sequel DOT open and specify what connector you're using versus postgrads.

And that tells the connector which database driver to use.

And then the DSN and our case.

That's just cfg dot dot dot dsn and check for an error.

If error is not equal to nil then we'll just say return nil.

You don't have any connections and our error.

Otherwise, let's go a little bit further.

Here's what we're going to use, the context will be the variable we're creating cancel is used to cancel

a context and we'll go to context, which is a built in package right from the standard library with

time out.

And we're getting that from the contracts that is always available, which is context on background.

And we give it a reasonable time out, say, five times time dot second like that.

And then we defer.

Cancel.

OK, now we try to ping our database so we can use the air variable, we just create a little while

ago, we don't have to reinitialize it and we'll say DB ping with context and handed the context we

just created.

We'll check for an error.

If error is not equal to nil, then return nil and error just as before.

Otherwise, return our connection pool and no error.

OK, so that's the function we're going to call.

Now let's go back up here in our main function just after the longer line.

And what we'll do in here is call that function.

We just created DB an error or assign the value of open DB and head to the config check for an error

if error is not equal to nil, actually want to die at this point because if I can't connect to the

database, nothing's going to work.

So let's say logger fatal and write the error, OK, and then as you always have to do, once you open

a connection pool, you need to defer closing it.

DB Duclos.

OK, so the last thing we'll have to fix is down here on line sixty eight.

My code where I call serve, listen and serve.

That can't be in an assignment operator because we've already declared it an error variable back on

line forty seven.

This just becomes equal.

Now if all went well I should be able to open my terminal and run this application.

And just one thing to note, Postgres must be running or we won't be able to connect.

So let's run, go run cmd slash api slash StarTalk, go on a Mac or Linux and of course it's go run

dot slash command slash API slash done on Windows.

So let's run it and see if we get everything right.

And I have connected perfect, so we have successfully connected to the database and now it's possible

to go on and to begin reading information from the tables in that database and serving actual content.

And we'll get started on that in the next lecture.

## 43 - Database functions and a challenge

So we've managed to connect to our database, now let's do something useful with the connection pool

that's available to us.

So let's go to our models directory and open models Dutko, and we'll start setting up the things necessary

to connect our handlers to our database.

So the first thing I'm going to do is create a new type called models, and that's just a struct.

And it's going to have one member, which I'll call DB for database.

It's going to hold something called DB model, OK, and that doesn't exist yet.

But don't worry about that.

We'll get to it in a minute.

And underneath that, I'm going to create a new function called new models.

And that is a function that takes one parameter, which I'll call DB and it's a type pointer to see

called DB and I'll choose it from the list here.

So does the auto import for me.

And there it is and it returns the models type we just created.

So in that it just has one return statement and it returns a models struct that holds as its member

DB DB model, which doesn't exist yet.

But we're getting to that in a moment and it has a member DB that's of type DB or has the DB that we've

passed to this as a parameter.

OK, now let's go create that DB model.

To do that, I'm going to create a new go file in my models directory and this is where I'll put my

database functions.

So that file I'll just call it movies.

Dash, dot, go.

Doesn't matter what I call it, as long as it ends and go.

So it has a package declaration and inside of that I've created type.

And this is where I put my DB model type and it's a struct and it has one member called DB.

That's a pointer to school DB and again I'll allow the import to be added automatically.

So in here we're going to have our database functions and I'll just create a couple right now and just

put Stobbs in.

So this will be a funk with a receiver of M that's my what I'm calling my receiver and it's a pointer

to DB model, the type we just created, and this will be called Get In.

It takes one parameter ID it and this will return a pointer to a movie model and potentially an error.

Right now I'll just say return nil nil.

So it'll compile then I'll copy this entire function and paste it here and call this one all and it

returns a slice of appointers to movies.

So the first function, of course, is how we're going to get one movie from the database.

And the second one is how we're going to get all movies from the database.

So now that we have these setup, if we look at our models, the error should go away and we probably

should.

And content comments in here just to be good programmers.

Models is the wrapper for database and new models

returns models with DB Cooper and we do these ones at the same time.

Anything that's exported is supposed to have a comment.

So we will do that is the type for movies for and genre

is the type for genre.

And finally movie genre

is the type for movie genre.

And back in our movies DB let's give our two Stobbs comments

returns one movie and error if any.

And all

returns, all movies and error, if any.

OK, so now back in our main function, we have to make some changes here.

We've created those those functions and so forth in our models package and now we need to use them.

So the first thing we'll do is in our application struct our type for application, we'll add a model's

field, remember models of type models, the models, and then down here where we populate our applications

struct.

Models will be assigned the value of it, and we'll just call that function.

New models and handed the DB variable that we created up here on line forty nine.

So let's try running this or just compiling it to make sure it all runs.

Go run cmd slash AP slash start out.

Go see if we get anything wrong.

Now, everything compiled, right, that's good, so let's go back and do something useful in our movies,

DB right here.

So right now let's just get this one function, get working.

Now, we're going to take advantage of the context, as you should when you're working with a database.

So this is a variable I'm creating and cancel.

And these are populated from the context package.

And we'll say with time out and I'll pick up from the list, so does the import.

And we're getting it from our default context in the background, context, dot background.

And we want a time out of, say, three seconds, three to three times time dot second.

OK, so that creates two variables and we'll defer or cancel.

So if things go wrong, it will just cancel the query.

Let's write a query now or query is pretty straightforward.

We're going to make a really easy one to start with.

Select ID and title and description and year and release date and runtime and MPLX rating.

And we'll also get created out and updated at from movies where ID equals dollar sign one and that's

our placeholder.

So that's our query.

Now we'll take advantage of that query to populate a row variable.

Lonhro is a sign the value of from our M variable from the receiver dot db.

And now we have query row context and that's the one we want.

We want to get one row if it exists from the database and time out if it takes more than three seconds.

So we have to hand that our context that we created back up on line 15 and our query.

And we want to substitute the ID we received as a parameter to the get function for the dollar sign

one that's in our query.

So we just put it and now we create a new variable ver and I'll call it movie of type movie and we'll

check for an error and try scanning from this row row scan and we'll scan right into ampersand movie.

Dot ID is the first row and I'll just duplicate this a few times to say some typing in the second value

is our title.

The third value is the description.

The fourth value was the year.

Then the release date.

Then the run time and a few more duplications, MPLX AAA rating

created at.

And updated, and I think I missed a row.

I think there's also rating, let's just see if there is ampersand movie rating.

Yes, there is.

So rating on that should come after release date.

So right here, we'll put in rating.

OK, if there's an error, if error is not equal to nil, return nil.

And our error, which is error.

Otherwise return the movie and nil.

Now we can't return the movie because we're returning up here a pointer to a movie.

So we'll make this a reference to the movie by putting an ampersand in front of it.

And that should work.

So now if we go back to our handlers, movie handlers and we look at this part, we're actually grabbing

one movie.

I can just comment this entire section out where I'm hard coating that movie and instead called the

database.

And that's as simple as populate the variable movie and potentially an error by calling app models,

which we added a little while ago.

We don't want the URL function.

We want the jet function and we're going to pass it.

The ID that we grabbed from the URL and nothing else should change.

So let's see if this works.

Let's open this up and run our application.

Go run systemd slash web.

No slash StarTalk, go.

And we have some problems here in movie handlers, of course, I need to save that to get rid of these

extra imports, so save it and run this again.

And it's listening.

So let's go to our Web browser and we'll go to localhost for thousands of VH1 slash movie, slash one

and see what we get.

And there is The Shawshank Redemption right from our database.

Let's change it to two and there is The Godfather.

So this works really well.

Now, there is one problem with this.

Of course, if we go back and look at our models, we actually have in our description for a movie,

we have movie genre.

And I commented that I didn't comment on that.

I told Jason to ignore it.

I don't want that to be ignored.

I actually want that to be populated with genres, if any, from a database.

Now, right now, if I look at my database, I have four movies in the database, but other movies,

genres, there's nothing there.

So here's a bit of a challenge for you.

What I'd like you to do is in your code, back in your movies go or deep movies dash deep, dark go.

I want you to modify this get function so that after you create and populate this movie variable and

before you return it, in other words, right here, get the genres.

Now, presumably you've worked with databases and go before or you wouldn't be taking this course because

as it says in the course prerequisites, some familiarity with go and with postgrads is required.

But I'll get you started anyway.

You're going to call you're going to populate a new query so you can reuse that query variable that'll

be equal to whatever you want.

Then you'll populate a variable maybe called rows and you'll check for an error if you want to.

But I'm going to just comment on it right now.

But to leave it out by saying ignore the errors, you will call em dot db dot query context and you'll

find that the context and your query and the ID for the movie that you want and this will be populated

with these genres, if any.

So all you really need to do is to write the query to create a variable that is a slice of movie genre

and then populate your movie variable with that value movie genre.

So give that a go and it might take you a little while, but it shouldn't be that hard.

The sequel is very, very simple.

So give that a whirl.

And in the next lecture, I'll show you how I solve this problem.

## 44 - Solution to challenge

So how did you make out with the challenge, hopefully you didn't find it terribly difficult, but we'll

go through my solution, which might be different than yours.

But if you got the end result you wanted, then that's fine.

So I'm in my movies, Dash Go file and I'm out of line 43 just before the the return statement in the

function get.

And I'm going to first of all get genres for a given movie if there are any.

So you create a variable rose and I'm going to ignore the error for now and I will call IMDB dot query

context right there and I'll it the context and I'll handle the query which I'll write in a minute and

I'm also going to hand at the ID.

So the query I want will be reused.

A variable I have above query is assigned the value of or is equal to and I'll call simply select and

I'm going to select.

I'll just get the and I'll use the alias mmHg for movie genre img id img movie underscore id, m.g.

dot genre ID and I'll get the genre name from the genre table.

And I'm going to select that from movie genre, movies, genres with the alias M.G. and I will left

join the genres table

with the alias of G on and I'll do the join on genre I.D. is equal to movie genres genre

and then my Where clause says just do it for one movie where mmHg movie idea is equal to dollars and

one which is my placeholder and that will be substituted with Edem.

Then I get my rose.

And after that, of course, I have to defer my rose close so we don't have a resource leak and I'll

create a variable which I'll call genres, and that will be a slice of movie genre.

Now, I'll loop through my rose if there are any entries in the rose for Rose Dot next, and I'll create

a one variable to be populated on each iteration through the Rose Bar and I'll just call it M.G. and

it's type movie genre

and I'll scan from my Rose Rose dot scan

and I'll scan into my variable M.G. ID and then mmHg Dot movie ID.

And scan doesn't require parentheses.

I don't know where that came from.

Try that again

like this.

That's better and duplicate that and make this one genre idee.

And duplicate that, and this time we're going to put it into the genre, a member of the movie genre,

and it's a genre name that I will check for an error if error is not equal to nil and I'll just return

nil and the error otherwise genres is equal to append to genres.

My current entry, M.G., and then just before the return statement, I'll take my movie variable and

populate its movie genre filled with genres that might be empty.

And certainly the first time I run this, that will be empty because there's nothing in the movies.

Underscore genres table.

But let's see if this works.

So we'll open our terminal and type go run CMD, slash API, slash Stargirl.

And actually I think you can do it this way on both Mac and Windows, which is something I learned yesterday.

Go run CMT dot slash CMT API just like that.

I did not know you could do that, but apparently the command is the same on both Windows and Mac,

I learned something yesterday.

All right, so this is running.

So let's go to our Web browser and make sure that it actually works.

And I'll just refresh this and it should give me the same information.

Perfect.

Now, of course, the genre's field is not showing up in this chasten feat.

So let's stop our application and go back to our models and find movie.

Movie genre in the movie struck.

And let's return genres, so this time it'll populate the Jason, even though there's nothing in there,

I should at least get something in my Jason.

So let's run it again.

And go back and reload this.

And there genres is no it has no values whatsoever, so let's go back to our database right here.

Open up post, connect to localhost and let's put some values in movie genres.

So I'm looking for good movies.

There it is.

And my first movie is The Shawshank Redemption.

And that's ID one and my genres right here.

That is both drama and crime.

So let's populate some rose in this field or in this table for movie ID one.

We're going to give a genre ID one and I'll put in some date.

Twenty twenty one zero five nineteen twenty twenty one zero five nineteen.

OK, so I'll save that and go back to my date of birth to my web browser and reload this and that.

Nulls should change.

And of course, that's already, too, so let's make this Eid one.

There so we have a genre's entry and it is an array adjacent array with the first entry being zero and

it has all that information.

Now, the only thing I really want in there is Jean-Rene name.

So can I actually make this simpler?

Sure I can.

Let's go back to our code and find movie genres.

So I'll just hide this and go to a movie genre and say I don't need to show the ID in the JSON.

That's not useful, nor is this, nor is this.

That's OK.

Nor is this, nor is this.

Now, if I save this and restart my application.

Stop this, run that.

We'll see what it looks like now, so we'll go back and refresh this.

OK, now let's just let's just get rid of the ID created an updated ad inside the genre model find genre.

I don't want to show the ID in the Jason.

I do want the name and I don't want this.

I don't want this.

As a matter of fact, back up in movies.

This is not useful information either.

So let's get rid of created an updated out from that and let's save this stop.

Our application started again and now we'll go back to our Web browser and reload this.

And that's much cleaner.

Actually, if you look at the raw data, you can see just how nice it looks or pretty printed and that

is much cleaner.

We're only given the information that's actually going to be useful to our front end, which doesn't

care about created or updated ad or the IDs, that sort of thing.

So this is a much cleaner way of serving that information.

Now, obviously, you can go and add genres to all of the movies in the database tables as appropriate

and play with this using different kinds.

But the next step we have to follow is to go back to our code and go back to our movies, Dash, DB,

Dongo.

And in here we actually now want to get this function working to return all movies and an error.

And of course, if you have a database with one hundred thousand movies in it, returning all the movies

is extremely inefficient.

But we'll take care of that a bit later on.

Right now, let's just get this working so we can have two functional things that we can try to connect

to on our react front end.

## 45 - An aside: cleaning up our JSON feed

So I had suggested in the previous lecture that we were going to do the get all function or the all

function for movies, but I think I want to improve or get a single movie first.

I'm looking at the Web page right now for The Shawshank Redemption Feed, where I get the Jason file

for a single movie.

And you can see that I've added a second genre there for crime.

And I want to simplify that just a little bit.

I don't want that to show up as an array.

I want to make that as clean as I possibly can.

I'm going to do that by going back to my code.

And finding the section in our models where we specify movie genre and I'm going to change that from

movie genre, a slice like this just to just to make it a map of string.

And you'll see why in a moment.

So let's format this and let's go back to our movies, TV, which now has an error.

And down here where I'm getting the movies, instead of having this version as a slice of movie genre,

let me instead change that to a genre is a sign the value of which means I get rid of her.

Make that a map.

Make a map of it string.

And down here, instead of appending to this slice or simply say genres and then M.G. ID, which gives

me my aunt is equal to M.G. dot genre, genre name like that.

And that should fix that problem now with no other changes if I start my application.

And started again.

And go back to my Web browser and reload this page, look how much cleaner that is.

We don't have to worry about traversing an array.

We simply have everything in a nicely formatted fashion.

And I like that much better.

OK, sorry about the interruption, but I just thought I'd clean that up a little bit before we go any

further.

It will make working with JavaScript on the front end a little bit easier.

So let's move on and fix that next database or complete that next database function.

## 46 - Getting all movies as JSON

So this time around, we want to complete the all function, and this is a function that will return

all movies currently in the database.

And as I said a while ago, if we had a really large database of movies, this would not be appropriate.

And we will change this later on.

But for right now, I just want to get a feed adjacent feed that shows me all of the movies that are

currently in our database.

So let's get started.

So the first thing I'll do is use the context, as you should whenever you're working with a database,

particularly with the Web application or an API, because the user might just lose connection partway

through the transaction, whatever the case may be.

But we'll give this a reasonable time out.

So context with time out and we'll get it from context background

and we'll give it a time out of three seconds

and will defer, cancel and let's write a query.

Now, the nice thing about this is the query we want is already pretty much written for us up here.

We just have to modify it.

So I'm going to copy this and go back down to the bottom.

And just replace this with that query and remove this where I.D. equals one and change that to order

by title, order our movies by title, and now we want to get our Rose.

So Rose and er assign the value of the query context right there and add the context and that our query

and we have no other parameters to parse, so that's enough for that.

We'll check for an error.

If error is not equal to nil, return nil and error and we'll defer rows close.

OK, so now we have our query so let's create a variable to store our movies in VER and returning a

slice of pointers to movies.

So let's create a variable called movies and make it a slice of pointers to movie.

And now we'll iterate through our rows, four rows dot next.

And the first thing we'll do is create a variable to store the current iteration, which I'll call the

movie.

And it's just a movie.

And I will say error is a sign the value of rows dot scan and we'll just scan it into our movie variable.

So ampersand movie the first variable as ID and I'll duplicate that a few times.

The second field is title.

The third field is description.

The fourth one is here, the fifth one is released.

The next one is rating.

After that comes run time,

after that comes MPLX rating

and then created and updated, which we don't need, but I like to be thorough, created and updated

at.

OK, so if there's an error, we'll return it.

If error is not equal to nil, return nil and error.

Now we want to get the genres, now we can just go up and copy the code we used before and modify it

as necessary, so we'll say get genres, if any, and we'll just copy all of this down to here and go

back down here.

And this is code duplication, and I'm aware of that.

And if you want to clean it up and break things out into functions, you feel free to do.

So I just want to get this running for right now.

So I'll paste this in and we have to make some changes, obviously.

So let me take this over and format the whole thing.

We can't call this query because we already have a query, so I'll call this genre query and that all

looks the same.

So this becomes an assignment operator and we copy this and we can't call this Rose because we already

have a gross variable.

So we'll call this genre rows and we don't want to defer the clothes.

You shouldn't defer a close inside of a loop, of course.

So this becomes movie ID for the current movie and this becomes Genaro's.

Oops.

Try that again.

Copy paste, paste.

And I think that's all pretty accurate there.

Down here, after we finish the for loop, we want to say genre roads close or manually close that because

that won't execute until this loop is finished.

That should be fine.

And we have movie genres, equal genres, and down here, instead of return nil nil, we actually want

to append.

Let's see what we call that variable movies.

OK, down here, we assign that and then just before we continue the loop, we say movies equals append

movies and a reference to movie because this is a slice of pointers and now we return movies.

All right.

So if I format that at all, looks good.

We have one error here, John.

That one has to be this query right here.

OK, so let's just go through this and see if we've got it right.

So we set up a context.

We have our query, which will get us all the movies ordered by title.

We store that query result in the rose variable.

We defer a rose close.

We define our movie's variable to hold the final result.

We iterate through the rows variable and scan the current row into our movie.

That's right.

Check for an error.

Then we create a new query for genres and that's going to use movie ID, which is set up on a line ninety

five.

So that's correct.

Then we create our genres map and then we loop through our genre rows, scanning the genre information

into a movie genre variable.

That's correct.

And then we set our map value and then we close our rows once we're done looping through the genres.

That's correct.

Then we assign genres to the movie genre field of the movie variable and append a reference to the movie

variable to our slice of movies and return it.

That looks good.

OK, so now we need to go to our handlers.

So let's find our handlers, movie handlers and we need to make a change here.

First of all, I can get rid of this comment and code because you don't need that anymore.

But we definitely should check for an error.

So we check for an error.

But we never did anything.

If there's an error, we say if error, we're just fixing up the get one movie.

If error is not equal to nil, then I want to right adjacent error.

So I say AFG error.

Jason and I handed the response writer and our error and I'll just return even though don't need to.

OK, so here this should be a really short and simple function.

We declare a variable movies check for an error and we call upon the models DB all like that.

So it looks like I have an error in my movies.

DB only go back there.

Am I calling for a parameter I don't need.

Yes, I get rid of that ID that shouldn't be there at all because we're getting all the movies.

So back to our handlers.

OK, so now we check for an error so I can just copy and paste this right here.

OK, so once we have those we just call error is is equal to iReport, right, Jason.

And we hand it our response writer, our status hdb dot status.

OK, and a wrapper which I'll call movies.

And again, we check for air, and I should be able to just paste that in, OK?

And of course, we need to give this the actual data.

So I think that should work.

So let's start our application.

Go run cmd API.

And let's go to our Web browser and see if we can call that route.

So right now we're doing version one, movie one.

Let's just make sure that still works.

It does.

Now, let's change that to movies.

So maybe one or two movies hit return and there it is.

Perfect.

So there's no genres for this one or this one or this one.

But there are the genres for The Shawshank Redemption and everything is formatted nicely.

So that was relatively straightforward.

Now we do have to come back and work on this a little bit later because we're going to want to check

the results, maybe 10 movies per page or something like that.

But for right now, we have enough to move along.

## 47 - Next Steps

So before we move on to the next section, there's just a little bit of cleanup I want to do and some

Stubbe functions I want to write and I'm going to start right here in the movie handlers don't go file.

And the first thing I'll do is get rid of this Apte longer print line.

It is because we don't need that.

That serves no useful function whatsoever.

I'm also going to create a few stub handlers, so I'll create one func app with the receiver of application

and it will be called Delete Movie and it will take a response writer and a pointer to a request

and that will be the handler.

We used to delete a movie and I'll just copy this and paste a couple of more.

So one will be insert movie will allow us to add a movie from our react front end and we'll say update

and a better spell update.

Right.

We're going to need that.

So basic basic crud methods which we'll get to eventually, but it'll take us a little while to get

there.

And we're finally going to want to have another one that says search movies.

And again, that will allow us to search movies and get some kind of response.

Now, these are just stubs that will remind me that I need to go create these methods.

And there may be more that we're going to add, but we're definitely going to need these in the same

way.

Back over in our movies, DB, at least HDB, Duco, there's a few more methods we're going to want

here as well.

But I won't bother creating those stubs right now.

You can probably guess what they are.

In any case, at this point, we have enough functionality for our back end and it's just a start and

we'll take care of that in a later section.

But we have enough right now that we can for the first time actually connect our front end written and

react with our backend written and go.

And we'll start that process in the next section.
