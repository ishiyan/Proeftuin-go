# 07 - Adding GraphQL into the equation

## 88 - What is GraphQL?

So in this section of the course, we're going to have a brief look at Grauwe QoL and that is a different

way of requesting information from your back end.

Effectively, graphical has a bunch of advantages.

We can get only the exact data we need without ever over requesting or requesting, and we can also

get multiple resources in a single request.

And the nice thing about this is that front end developers absolutely love you will, because they don't

have to contact backend developers and say, can you modify this query or can you give me a new query?

Or I need to get the data in this format.

The front end developer has a lot of control over the kind of data that's sent back by the server.

Now, there are many ways of implementing graphical and you'll actually find add ons for the Postgres

database that actually implement a graphical server with very little effort.

But in this course, we're going to have a look at how you do it in go.

And what we'll do is just to a few queries will establish the necessary endpoints on our rest API so

that you can make a request for movies, for an individual movie, so on and so forth, and we'll show

how to consume that information.

Using plain JavaScript.

There are add ons or imports like the Apollo client that make working with RFQ a little simpler.

But I think it's important to understand how it works.

So we'll do it using just plain old JavaScript within react.

So let's get started and we'll start with the back end.

## 89 - Setting up a schema and REST endpoint for GraphQL

So let's get started setting up our backend to handle graphical requests, and I'm in my source code

for my rest API.

I'm looking at the rootstock go file and I'm going to set up a root here.

But first, let's get rid of these comments about roots because we don't need those anymore.

There and let's set up a route to a nonexistent handler that will take care of our fuel requests, so

router dot handler function and it will be HTP method post because on the front end will be posting

JSON to this route and will set up the path name to the route being V1 and I'll just call it graph and

we'll just do a really simple one to get started here and we'll just say, listen, this will list the

movies that are in our database and we'll call a handler that doesn't exist yet, but we'll create that

in the next step app DOT and I'll just call it movies graphical doesn't really matter what I call it,

but we'll go quickly because that's the way they name it.

OK, so there's our route.

Now, let's go create a handler for this.

So I'll create a file right in command API.

So in this folder, Command API, we'll create a new file and I'll just call it graphical dutko.

And it is of course in package main.

And what we're going to do first before we start writing the code, we actually need to implement or

import a third party package that will give us a much easier way of dealing with graphics will.

We could do it all by hand, but there's no point in reinventing the wheel.

And there are a number of packages to choose from.

And it really doesn't matter which one I choose because people who are familiar with graphics will some

will say you picked a great package, good choice, and others will say, why didn't you use this one

instead?

So there's no way I can be right and satisfy everyone.

So I'm just going to pick one and the one I'm going to pick.

I'll open a terminal window and clear the screen.

So let's just go get one.

Go get from GitHub dot com, siggraph cual Dasch, go siggraph Kuo and there it is.

So that's been important.

And we're going to use this one for this example.

So the first thing I'm going to do is I'm going to create a package level variable and I'm going to

call it movies because we've never used that anywhere else in this package and it's going to be a slice

of pointer to models movie.

OK, so that did the import for me, so that's good.

Now I'll create the handler itself and then we'll have to do a whole lot of work to describe the kinds

of data we're going to be handling or sending back from our code.

So let's create the handler first just to get rid of the error in our routes.

File func and it takes the receiver of a pointer to application.

And I'll call this movie's graphical, which is what I called it, my roots file.

And of course, because it's a handler, it takes a response writer, response writer and a pointer

to a request to request.

OK, so that gets rid of the error in our roots file.

So we know that the very first thing I want to get, we just want to get a very simple graphical request

and response working properly.

And as you might have guessed from the name or the root name and or reads file list, I just want to

return a list of movies.

OK, so we need to actually get that information from our database.

And unfortunately, that's pretty easy.

And just to say some time ago movies and I'll ignore the error, OK, and that's going to be equal to

because you already have this movie variable declared, but it'll be equal to the same as we did before

after models dot DB and that will populate the package level variable movies.

OK, then I know I'm going to be getting the request from my front end in the form of Jason.

So let's read that we can do that with or without doing anything else.

Q Which is just a variable I'm declaring, and again, to save time I'm going to ignore the error and

you shouldn't, you should do whatever error checking is appropriate for your application.

And I'll just call the from the standard library.

I don't read all and I'm going to read the request but I read all returns a slice of bytes and that's

not useful for me.

So I'm going to cast that to a string.

So we'll create a new variable called Query and that will just be webcast Q to a string.

So now we have that.

And when I actually make the request for my front end, which we can't.

Yeah.

Because we haven't written that code yet, just to make sure everything is working properly, I'll just

say log print line and print that.

OK, so now this would, this will at least compile.

So we have a route right here, slash V1 slash graphical slash list.

It's not a protected route.

We're making this one publicly available and we're calling the handler movies graphical.

And all we're doing so far is saying get all the movies.

So populate that variable.

So this will actually populate the movies variable.

And then we're reading the request body, getting our query, which will come in the form of JSON and

we need to do something with it.

And this is where we actually need to describe the kind of data we're going to be sending back.

So there's a few steps that we have to follow here.

The first thing we need to do, and I'll do it right up here after my declaration of the movie's variable

variable is I need to create a new variable, which I'll call Fields, because I'm going to describe

the fields that will be using engraft.

You will.

And here's where we start using that graph Keywell package we just imported.

So I'm going to say that's equal to the type graph dot fields.

She will do the import for me.

And of course, this in curly brackets has to have some descriptions, so the very first thing I'm going

to say is this is going to be a movie, OK?

And that is going to be read into graphic.

You will not field.

And it's a type a movie type, which we don't have yet, but don't worry about that, we'll get that

before too long.

OK, and the description is going to be get movie by idea will define a couple of ways of accessing

the data to get the entire list of movies and to get a single movie.

And we'll use both.

And that has some arguments.

And again, this is from the graphical type.

So ERGs is going to be graphic.

Well, again, dot field config argument and we're describing what this is and the field is ID and the

it has to be a reference to graphical argument config for Whewell argument config and the type is graphical

and OK, I don't worry too much about the syntax for this.

This will make sense once we actually start doing things with it.

But as you can see, what I've described there in the arguments is to say I'm going to have a movie

and it's going to have an I.D. and it's going to at this point, I've got to put a comma there or this

will compile an a comma there and or this will compile.

Then we have resolve.

What is this resolved to?

Well, it resolves to a function and it takes one parameter and that's a type graph.

Q Will resolve problems, which is a built in type from this package.

Resolve Hiram's

and it returns an interface and unknown type or any type, an error potentially.

And from here all we're going to say is ID OK, standard checking a map is Peter ERGs looking for the

key.

Idy casting that to an end and we check to see if it's OK, if OK.

And now we loop through movies to see if we can find one with the given ID, the idea that supplied

to us.

So for and I'll ignore the index movie is assign the value of range through movies and if movie ID is

equal to the supplied ID, the parameter we're looking for will return movie and no error otherwise

down here.

Will returns nil and nil, and we need a comma right here or this one comp. And we need a comma right

here are this one comp. So that's our first one.

And we don't have I got a typo or G or a p h q we don't have movie type defined yet, but don't worry

about that.

We're going to get to that next.

So that's for looking for a movie by ID.

Now at the same level as movie, we can now specify another one which I'll call list and that's the

one that our current route is going to target.

So again, ampersand graph cual field and the type in this case is going to be a graph cual new list

because we're not returning a single item, we're returning a list of items.

And again, this is of type movie type which we'll get to shortly and we'll give it a description which

is optional but good practice description.

Try that again description

and that will just be get all movies.

And after that we have a very simple resolver.

Resolve, resolve the request to get a function again with programs and again, their graphics will

resolve, resolve Hiram's and it will return just the same as before interface and potentially an error.

And here all we do is return movies with no error and put a comma there and put a comma there, OK?

And I misspelled resolve, perhaps all the parameters.

OK, so now we have this is what's called our schema.

So I'll just put a comment there.

So you know what it is.

Grauwe Cuil schema definition.

This defines the schema and that allows us to permit remote users to request whatever they want or any

kind of data that they want.

So when they ask for a movie, if they just want the movie ID and the movie name or just the movie name

and just the movie release date, that's all we're going to send back.

So it makes it a very efficient method of constructing queries.

So let's take care of that movie type ver movie type.

And this is equal to and this is what describes what's in our database to graft you up.

So this is a graph, cool new object.

And in here we need to describe its configuration.

And this is basically going through the fields we want to expose to graphic will from our database.

So first thing we type is graphical object conflict.

And this configures the object.

And the first first item in that is name.

And we'll just call this movie with a capital.

And its fields consist of graphical dot fields, plural.

And inside of this we have one entry for each of the fields we want to expose.

So we want to expose ID and again, ampersand graph Kuo dot field singular and we describe which type,

which is the graphical, not boolean it.

OK, how to save a little bit of time.

Let's go over necessary commas before we forget.

There it is.

OK, let's copy this I.D. and paste it and modify the next thing I want is titled Straight from the

Database, and the type is a string in this case.

And if you look at the definition of our table in the database, again, the next one is description,

which is a string.

And again, I'll paste this in.

And after a description comes here, which is in it, and after a year comes release date and this one

is a date time.

That's how time stamps are handled in grad school.

And after that comes runtime, which is an end, and after that comes rating, which is in it, and

after that comes MPLX rating.

Which is a string in this case, and then we have created an update, which is a date time

and updated so popular

and that's a date time.

So at this point, we have no errors in our code, and that's good.

So let's go through this again.

First thing we do in the package where the handler to take care of the graphical endpoint exists is

we declare a variable to hold the data.

We're going to be querying.

In our case, it's a slice of pointers to models movie, and that's all we care about at this point.

Then we define our schema.

And this tells graphical what kind of queries we're going to be getting.

And in our case, we only have two one for an individual movie online, 17, and we're going to get

that by ID.

And the second one right now is just a list that says, give me all the movies.

We're starting with really simple stuff just to make sure we can get it working and then we'll improve

upon it from there.

And then we have to describe what kind of information in our variable movies, which is a slice of pointers

to models movie.

What kind of information do we want to expose to the end users?

So, for example, if I didn't want to expose created it created an updated I just wouldn't include

it in this definition and it would be ignored by graphics so people could ask for the created out or

updated ad, but they wouldn't get it.

So that describes our data and that's what we're really concerned with in this point.

And the next thing we need to do is to actually implement the logic, to pass queries and populate the

appropriate results and send that back as JSON.

And that's going to be taken of here like eighty four movies.

Graph QoL, our handler, and we'll get started on that in the next lecture.

## 90 - Handling the GraphQL request

So let's keep going with our handler, and as you're going to see, it takes a fair bit of work to set

up, a graphic will back end and it's kind of a trade off.

If you want people on the front end to do less work, you're going to have to do more work on the back

end.

But it's not that difficult.

So so far, all we're doing right now in this handler is populating our movie's variable or package

level variable movies.

And that's the data we're going to query using graphics.

Well, so we just call our back, end our old method, get all the movies into a slice, and then we

read the request body, which comes in the form of a JSON file and we're just logging it and now we

need to do something with that query.

So let's get started.

First thing you have to do is specify a root query and I'm going to create a local variable called Root

Query, which is assigned the value of from the graphical package we're going to call object config

for configuring an object.

And that takes in curly brackets a name which I'm going to call root query like that.

And it takes fields and that's just our local variable fields.

So now we have a query.

Now we need to create a schema config, which I'll call schema config.

And again, that's a variable name of my choice.

And again, from the graphical package we're going to call the method schema config, and that takes

query, which in this case will be a graphical new object.

And I'm passing it the value query, the variable we created on the line just before this one.

So now we have our schema config and finally we actually get the schema we can work with, which I'll

call schema and we're checking for an error.

And that comes from graphical new schema.

And we hand it schema config and we check for an error if error is not equal to nil.

I'll just say Aptera Jason and I'll hand it the response writer and errors new and I'll just put in

failed to create schema and I love the logo print line error and I'll return and I guess I need to import

the errors package but I'll admit do that errors and has a capital there.

But so now we have our schema.

At this point we have to see what are the programs we're getting in our request.

And we do that again by creating a local variable parameters.

And that's a sign the value of graphical parameters.

And we can figure that by saying schema is schema and the request string, which is something it requires,

is our actual query.

What we got back up on line eighty nine.

So we're finally handling our graphical request and we get a response.

A response is assign the value of graphical dot do and we handed prints.

So now we have a response and we check for errors.

So the way you check for errors here is by saying if the length, not length, if the limb of errors

is greater than zero, then we have at least one error.

Lenegan again, I'll just call APTA Error JSON and I'll handle the response right here and I'll give

it a new error error.

Start new and I'll just say format s print f fail and then I'll just put in the errors percent plus

the and just print whatever you have there.

Respeto errors.

And I got too many parentheses here.

Sloppy typing today and import the format package.

OK, now, if we get past that, we're actually going to get a response we can send back in the form

of Jason, so I'll declare a new variable J and I'm going to ignore the error.

And we just say Jason Marshall and I'll say Marshall Indents so we can actually look at it in the log

file and it will be readable and we're going to Marshall Recip with no prefix and a few spaces and we'll

write that to the end user header set.

Content type, it's going to be in the form of application, JSON and the right header and will write

status, OK, http dot status, OK, and we'll write to the response rate dot right, Gerke.

OK, so if this compiles and hopefully it will go run runs, dot slash CMD, then our back end should

now be set up to handle one kind of request.

And again, the only kind we're handling right now is from our two available ones in the fields.

In our schema definition, we're going to listen for either a movie which returns one movie or list

which returns all of the movies.

OK, so the next step is to go to our front end and write the necessary JavaScript and Jass X and so

forth to make the request, and we'll take care of that in the next lecture.

## 91 - Implementing GraphQL requests for all movies

So we have the back end set up to make requests using graphics.

And now we need to do that implementation on our front end.

And what I'm going to do is just add a new menu item here called graphical, and we'll play with that

to see what kinds of things we can accomplish using graphics.

Well, so let's get started.

So we'll go back to my idea and I'm looking at the front end source code.

And what I'm going to do, first of all, is create a component just to hold our our code as we experiment

with graphics.

So I'll create a new component, a new file, which I'll call you up

and I'll do my import.

PMRC and I will export a default class called Hubo, which extends component.

And we have to have a constructor which takes props and we'll call Superprofits and I'll set up some

state.

This stuff state equals and we're going to need a variable to horror movies.

So I'll put one in here called movies and it will start as an empty array.

Also put it loaded in which will default defaults and error, which will default to null.

And I'm probably going to use in Alert's or particularly in as well type initially hidden.

So Dynan the bootstrap class and the message will be an empty string.

OK, so there's our constructor.

Now the very first thing I'm going to do here is create the component did Mount Method and we'll use

that to construct and send a graph to our request off to the back end.

So component did mount

and the first thing I'll do is construct my query and I'll call that payload because that's what we're

sending to the back end.

And I'm going to use the JavaScript template syntax.

So tactics.

So let's create our first graphical query and we're going to call is the command list because we want

to get all of the movies so we do it like that list and then in curly brackets and the whole thing is

wrapped in curly brackets and went to I want to get from the back end.

Well, I want the ID and there are no commas here and the title and say the runtime and say the year

and that's enough for now.

So that's where I'm constructing my request and this is what I'm going to send to the back end.

So let's create now our fetch request.

So first thing we'll do is create some headers.

Const my headers, just like we did before, is equal to a new headers

and will spend a header to it.

My headers append and we're just going to tell the back end what kind of content you're getting.

Content type and it is application JSON.

Now we'll create our request options, const request options and that's equal to an object and it has

the method.

And in our backend we're expecting a post request.

So we'll send a post.

The body will be our payload and the headers.

Will be equal to my headers.

Now, let's do the actual fetch request, so just like we did before fetch and we're going to fetch

from HTP Colon's localhost PT. four thousand one SIGGRAPH slash list and we pass along our request options.

And then we have our first, then close.

Then we're going to take the response whatever we get back, convert it to Jason.

Response Tajikstan then.

And here's where we actually take what we're getting and populate a variable.

I'm going to call the variable the list because it's a movie, a list of movies.

So I'm getting data back.

That's what I'm calling it.

And I'm going to use an arrow function and create a new variable.

Let the list equal and we're going to take what we get back and grab what we need from it and populate

the list.

And we do that the same way we did before.

Object got values and we're going to get data dot data and you'll see why when we get the response in

a moment.

Don't list.

And now we return that return the list, so it's available in our next dot, then, dot then and now

we're going to get the list and using arrow function.

And I'll do right now is just say console that.

Log the list.

OK, and let's go now create our render function, because this won't compile without one render and

we'll just return and for now we'll just return to graphics.

OK, so hopefully everything compiled, let's find out.

It did perfect.

So now we need to create our menu item and actually display this component.

So let's go back to Abcess and let's find where we're going to put this, like I'll put it right here

just before the closing up.

So what I'm going to put in here is Elai, and it has a class name which is equal to list grouped item.

And inside of that will route to something that are linked to something that doesn't exist yet.

We'll create that in a moment.

Link to equals and we'll just put graphical.

And then down here, say right here, doesn't matter where it goes, but I'll put it right here, we'll

say route.

Exact path equals SIGGRAPH cool, and it will link to our graph culo, which I have to import and hopefully

did the import for me.

Let's make sure there it is.

So now we should be able to go back to our Web browser.

Everything is working, it seems.

I'll open the JavaScript console

and I guess I left something out of there.

Oh yes.

The text to the actual like.

Cool.

That's better.

And there's my link, so let's watch the console when I click on this and there it is.

We got an array back and inside of that array.

We have all of our movies, so let's go display those.

So I'll go back to my ID and find graphical dojos and let's improve this render function.

So I'll wrap everything and returning in a fragment and that should do the import for me and drag this

down here.

And inside of here, I will put say an hour and then I'll just deal with what's in my state the same

way that I did before.

So before the return statement, let's say let movies equal to this state.

So I have a local variable called movies I can work with.

And after the year, I'm just going to say do class name equals list group like we did on the other

page, where a component where we're displaying movies.

And inside of that I'll say movies, dot map and I'm going to map over M that will be my iteration.

And inside of this I will put an Alyque A and I'll just say key equals

M ID.

So I have the key so we won't get any warnings and the class name is going to be equal to list group

dash item list dash group dash dash action.

And I'll give it an F. I'm not going to link to anything you can link it to whatever you want to, so

I'll just put hash back and then inside the attack, I will put the title so strong and then my title.

And that should be enough right now just to make sure it works.

So let's go back and have a look at our Web browser and go to grad school, reload the whole thing and

go to grad school and nothing is showing up.

So what did I get wrong?

Oh, yes, of course.

Up here we consolo of the list, but we actually need to populate the state variable.

So I'll just say this starts set state.

And I'm going to populate movies to be equal to the list.

OK, let's try that.

That looks better.

So now I have all of these items and that's exactly what I want.

Now, I can do more than that, actually.

You'll notice when we're getting back here is only the information that we requested.

So for the first index in the array, for example, I'm getting the ID, the runtime and the title.

But what if I want the description or I can come back here and modify my query and say, include the

description.

Now I want to come back here and look at this in my console.

I have the description as well.

So that means I can actually modify the kind of information I'm displaying.

So let's go back to our idea and scroll down to the purple rendering this and inside that a tag.

Let's put a V.R. after this.

So we go to the next line inside the tag and now we'll say small class name equals and I'll make it

not so dark, say text dash muted.

And inside of that, I'll put in parentheses the year the movie was released and year and then outside

of that, a dash and then the runtime runtime minutes.

And I'll just get rid of this closing parentheses and move it over to here.

And then after that, I'll put another line, say here and I'll put some of the description, I don't

want the whole description to be there because it might be long.

So instead, I'll go of the description that slice from position zero to a maximum of one hundred,

and I'll put it outside of the closing brackets.

Dot, dot, dot.

So let's go back and have a look at it now.

And there it is.

Now, that's not bad.

So there is a very simple way that we can use graphics to get to the entire list of movies and display

it.

But of course, we can do much more than that.

We can get an individual movie, which is straightforward, and we can also and this is what we'll do

in the next lecture, maybe at a search field at the top that allows us to filter the lists of movies

based upon the characters that appear in the title.

So we'll take care of that in the next lesson.

## 92 - Adding a search endpoint

So this time around, we want to set up search functionality and allow the end user to filter the list

of movie based upon the title, and we're going to do that, obviously, first of all, in our back

end.

And that means we're going to add in our fields, in our schema definition.

We already have a movie and we have list.

We're going to add another one for search.

And right away I start to think about the route that I'm using to access this functionality.

Let's look at our reads file right now.

We have here on line twenty five and my code, we're listening to V1, SIGGRAPH Quill's list.

I'm going to get rid of that list part and just change this to graft.

You will, because the list is actually not meaningful.

The directive as to what we want to do, what kind of content we want to get back is actually part of

our payload.

So I've just modified my route and I'm going to start my back in the application and clear the screen.

And I won't start yet because we have more work to do.

So we'll have to remember to change the front and change the way you are recalling.

And that's pretty straightforward.

So let's go back to Graphical Dutko and I'm going to add a new function here, a new action, which

I'll call search.

And just like the other two, it has to be ampersand graphical dash field.

And we describe what it's going to consist of.

And I'll put my closing comma there so I don't forget.

So first of all, Tike, just like the rest of them, is graphical.

And we're going to get not a boolean a new list because we're getting more than one potential movie

here.

And its argument is movie type, just like the one where we're getting the list of all movies and we

give it a description which I'll call search movies by title.

And here we have to have arguments because we are passing arguments for this particular one and the

arguments are a graphical dot field config argument, and again, that's just a struct and we have to

describe its contents.

So we're going to and this is what we're going to be passing as part of our payload title contains,

which makes sense because we're going to be searching by the characters in the title.

And its type is ampersand of graphical argument config.

And its type is graphical strength, researching strengths, researching a character or one or more

characters, and then we have its resolver.

So down here.

Right there, karma resolved.

What is this result?

It resolves to a function which takes Purim's resolving the parameters.

Gravity will resolve problems and it returns interface and potentially they're OK.

And we need a comma there.

It inside of this, let's create a variable to hold the results of our search via the list, and it's

a slice of pointer to models don't movie.

So this variable will be populated with the results of our search.

And now we're going to get two arguments here or two very search and OK.

And that will be assigned the value of programs, dot ergs and in square brackets.

We're looking for the key of title contains and we're casting that to a string.

So now we've got our parameter, what we're searching for, and that will be whatever the end user types

into the search box on the front end.

And then we see if OK, if OK, then we have something to search through.

So we'll just do a for loop for and we'll ignore the index and we'll call the current iteration current

movie.

You can call it whatever you want, but I'm going to call it current movie.

And that comes from ranging through movies, our list of all movies.

And we just have a simple if statement here if and don't use the strings package strings DOT contains.

And we're looking for current movie title

at our search term.

What we're looking for is stored in the variable search.

If the string, the current title of the current movie contains the characters were searching for,

then just add that to the list.

So I'll just put a log print on here so you can see how it works.

Logged off print, found one and then I'll just append to the list is equal to append the list and we're

appending current movie.

And finally outside of this if statement we just return the list and no.

And that gives us another action we can perform with our graphical query.

So if this compiles, let's make sure it does go run dot slash cmd, slash API.

Looks good, so the next step is to go to the front end and change what we have there, or at least

add to it to actually access this particular graphical query.

And we'll take care of that in the next lecture.

## 93 - Implementing GraphQL requests for search on front end

So now that we have the search functionality set up on our back end, let's implement it on the front

end and I'm looking at the graph.

You will not just file in my front in code.

And the very first thing I'm going to do is here in my component did mount function.

I'll modify this URL to be the one that exists in our new back in code.

OK, so that change is done.

Now, the next thing I want to do is to come down here to the render function and we'll put in a place

where people can actually perform a search, where they can type in a search.

So let's get an input in there.

We may as well use that nice input component we made a while ago.

So input right there.

And that should do the input for me and I'll give it some blank lines and the type here.

Well, the title let's do the title first is going to be equal to in curly brackets and in quotes search.

So that's my label and the type will just be a text

and the name I'll call it search

and the value will be equal to something.

We have to set up this dot status search term, which doesn't exist yet, but will shortly.

And we'll have to also put a handle change in their handle.

Change will be equal to a function that we have to write code and will change.

OK, so that's our input.

Now up at the very top in our constructor, let's put it in a search term in our state search term will

default to an empty string.

OK, and I want to bind in my constructor this dot handle change, just like we did before, equals

this dot handle change, dot bind and this.

OK, so that's the binding.

Now let's create our handle change function.

Kendel change.

And we'll use the Arrow syntax and of course, we're going to grab the value of the form field and there's

only one value equal AVP dot target value.

And set the stage, this set state and again, just to be safe, previous state and the era function,

and we're going to set search term to value it.

So there's our handle change set up right away.

Now, the next thing we want to do is we actually want when something changes in our search box, we

actually want to perform a search.

So what we're going to do is create another function which will call perform search.

And it's going to be remarkably similar to what we have here.

So let's just grab the constant payload.

Copy that and modify it.

So we're going to have a payload that we're posting to our back end.

But it's not going to be a list.

It's going to be search.

So let's change that to search.

And it's going to take an argument in our back end.

We call that argument title containers like that and its value will be equal to and because I'm using

a JavaScript template, I can do the double quotes and put a dollar sign and then a curly bracket,

this dot state circuit and close that.

OK, and what do we want to get back?

Well, I'm going to get the ID, the title, the runtime, the year, the description.

I can get other stuff as well, but that's sufficient for our purposes right now.

So that gives me my payload.

And now I want to do pretty much the same thing.

So I'm going to copy the headers, the fetch statement on a copy, all of this, then I'll modify it

as necessary just to say some typing.

So back up in my performance search function, return a couple of times and paste that in.

Now there is going to be one slight change here and that is the final then clause.

First of all, let's shift to a console along the list.

We have that and I will put an if statement in here.

So I'll cut this and I'll say if the list length

is greater than zero like that and get rid of this because we don't need that, that's better.

Then I'll paste that in there.

Set the state of movies to the list, which is great.

But if we have nothing, I'm actually going to say else and I'll paste that in there, but I'll make

the list equal to absolutely nothing and an empty array like that.

And that way, if we're searching for something that doesn't exist, it won't leave values on the screen.

OK, so that gives our affection.

Is there anything missing here?

You notice the most important thing from our perspective is the payload where in one component did mount,

which is right down here we have list as what we want to do.

But on this one or we're searching, we have search plus a parameter and the value of that parameter

is actually grabbed from an input box.

OK, so there's that.

So we have our perform search function now back up in our handle change.

We actually want to perform that search so we can just do it like this for right now.

This don't perform search.

OK, now let's make sure that everything compiled and it did and let's make sure our back is running

and I'm going to stop and start just to be sure.

So control, see and run.

It's running.

So let's go back to our Web browser and give this a try and see if we missed anything.

So I'm in here.

I will clear the console.

We have our search box, so that's good.

So let's look for the word, though.

And we have some can't convert undefined to objects.

So I must have missed something.

Let's go back and have a look and we'll have our back end and we'll go back to our front end and see

what we missed.

So it has to be in the perform search.

Oh, yes, data data list should be data, data search.

And we'll see why in a minute here.

OK, let's just say console.

Don't log Jason right here.

Data on Jason and then down.

And the component did mount will do exactly the same thing here.

Console log

data.

OK, so let's go have a look at this.

This now we'll go back, will clear the console and I'll reload this page.

And you see that we have this back here.

We have object data.

Now let's look at the network and I'll clear this and try this one more time.

Reload the screen.

So first of all, I have this, and it gave me a response to this one here and it gave me a response.

So my request was list everything, give me the ID, the title, the runtime, the year and the description.

Now, this was sent to the back end on component did mount and the response we got was actually like

this.

Let's go to RAW.

We got a wrapper are to curly brackets to wrap the whole thing, then data, then list.

OK, now if I go search for something, let's say a search for the now it's calling that using our perform

search function and it is curly brackets data search not list.

So back in our code

here and this function is component did mount.

I'm looking for data data list and on perform search I'm looking for data data search and that actually

worked.

So we go back and look at our Web browser again.

When I look for the I see Dark Knight, but if I get rid of this and put in, say, show, it gives

me The Shawshank Redemption and when I back up and type dark, it gives me The Dark Knight.

So it works and it works really, really well.

Now, obviously, that's not a full implementation and there are things that could be done to simplify

this.

But this should give you a clear indication as to how useful graphics can be.

Now, some people love it, some people don't like it.

It certainly requires more work to set up on the back end.

But it also gives people working on the front end far more freedom to construct the kinds of queries

that they want and get exactly the kind of response that they want with no extra data whatsoever.

So there you go.

Time to move on.

## 94 - Displaying one movie using GraphQL

So a graphical functionality is coming right along, and right now we can list the entire catalog of

movies, we can type in the search box and filter the list of movies by name, and we might as well

complete the functionality to click on something from this list.

And that's the graphical link, of course, and actually load an individual movie using graphics graphical.

And we might as well improve it a little bit while we're at it, because right now these links don't

work.

But if I go to the movies list up here, then click on The Godfather, for example, I get some textual

information and it's trivial for you to add the rest of the information if you so choose.

But I really like to add a poster, say an image up here, and I don't want to do it manually because

there are posters all over the place.

So what we'll do and we'll do it just for the graphical component.

And it's easy enough to do the same thing in the standard rest API component.

But what we'll do is we'll go to this place, the movie database, and they have a free API that allows

you to query information.

So what you'll have to do, of course, is create an account.

And once you've created it, you'll get a validation email, click on that, then log in and click on

your little avatar up here.

If you have one, it might be just an image of a user, but click on that and click on settings and

then click on API and create an API.

And as you can see, I have one here and I'll leave these ones active.

You can use my key if you want to until it gets invalidated for too many requests.

If lots and lots of people use it, I encourage you to use your own.

It's free and it takes very little effort to do it.

And once we've done that, we can actually query things like this.

So I've gone to this URL up here and don't worry about that right now.

I'll provide it to you when we get to that part of the the course and we can get the response about

it in an individual movie by name.

So I searched for The Shawshank Redemption and it gave me this.

And this is actually a path name to the poster for that particular movie.

Actually, it's not that wants this one down here.

That's some backdrop.

I have no idea what that is.

This is the one we want poster path so we can actually modify our code so that when someone ads or updates

a movie, we can check to see if the movie debug has a poster for it.

And if it does, we can use that and display it when someone looks at the individual movie.

So we might as well do that at the same time.

Now, this is going to take a few steps and one of the things we're going to do first is modify the

front end to display the current information that we have.

Then we'll modify the backend to allow for the addition of a poster and that'll require a modification

to our database and our backend code.

Then we'll come back to the front end and put everything together.

So let's get started and we'll get started on the front end in the next lecture.

## 95 - Updating the front end

So let's get started displaying one movie with a graphic you'll call in the front in source code, and

that's where I am right now.

And as you might expect, this file, one movie dojos, which is our component for displaying a movie,

it's virtually identical to what we're going to need when we make a graphic.

What will change, of course, is the component did malfunction, but this is a good starting point.

So I'm going to select this all this code, copy it, close this file and I'll create a new component.

So a new file and I'll call it one movie graphic, which is not terribly inventive, but will serve

our purpose.

So now paste all the code I just copied in there and go back up and change the class name to one movie

graphical.

And here in the component did mount function.

We need to make some changes.

So what I'm going to do is say some time is go to the graphical Dutchess and I'll copy its entire call

to the back end.

So everything here from line forty seven in my code down to here, which is the fetch statements, I'm

going to copy that and go back to my one movie graphical dojos and replace this entire fetch statement

with what I need or my starting point is anyway.

OK, so we definitely want headers so my headers equals new headers.

We definitely want the content type application JSON, even though this isn't actually accurate or valid

JSON that will set the content type appropriately for our back end.

And then we have a request options which will be post and payload and my headers.

Now we need that payload variable.

So let's go back to graphical dojos and just copy a payload as a starting point.

So this is our starting point for payload for copy this and go back to one movie graphical and paste

my payload in there like that.

OK, and now let's fix this.

Of course, we're not using the search directory from our back end.

We can go look at our back in code right here.

So the one we want and this is great QR logo in the back end code.

It's this one movie.

So that's our DirecTV or our action.

And it has one argument ID.

So let's go back to our front end code and change this from search to movie and not title contains.

We want ID and ID is actually an integer so we don't use double quotes because that will tell everyone

that it's a string and we don't want this state.

That search term.

If you recall, this actually comes from our props.

It's an argument pass to us by the react rotor.

So we want this props dart match parameters, biometric Purim's ID and we're going to get more than

ID total runtime your in description.

We may as well grab release date and rating and empty AAA rating and that's enough for now.

I think that gives us pretty much everything we need for our purposes at the moment.

Then down here, we're not going to get the list.

We actually don't need this at all.

You can delete this one.

We don't need that one at all.

We're getting a response, Jason, which we'll call data and we'll get rid of this entire console log.

And our movies variable will come from Jassam data.

But this time it's not list and it's not search.

It's the name of our action, which is movie, OK, and movies.

This should be movie singular.

And we can just get rid of this else altogether, just get rid of that and get rid of this and we're

calling the data.

So this should be updated.

Jason, data.

Data there.

All right.

And of course, we have to close our function.

Let's let's format this whole thing that looks right.

We have no error showing, OK?

And back up here, we have state movie is loaded faults and error.

And down here we're getting a post.

That's the correct YORO.

And I think that's everything we need at this point.

OK, so that should be unless we missed something all of all that we need right now in one movie, Gravity.

Well done.

So now, of course, we need to route to it in Aptos.

And what we'll do is come down here.

We're not going to have anything in our menu that appears on the left so this part can stay the same.

But we do want to read it to our component.

So I'm going to duplicate this line and I'll just change the URL to graphical movies, graphical movies,

graphical colon ID.

And the component will be one movie graphical and I'll let my ID import that for me.

Now we'll go back to one to our list of movies that we're getting and that's in graphic.

You will just find the part where we're actually showing the movies, listing them, which is right

down here.

And we just need to change this a to a link.

So I'll use the link and let my ID be important for me.

Hopefully there, give it a closing link instead of closing a and the atrip becomes too.

And then the actual Parthenon that we want to link to in curly brackets with tactics is movies and of

S graphical graphics will slash dollar sign and make it.

And if everything worked as expected, let's make sure it compiled good.

Let's go back to our Web browser and click on graphical and try clicking on, say, The Dark Knight.

And it gives me a loading because I didn't set my loading variable.

Let's go back to our code, bring up one movie and we're saying is so at this point, we also want to

say is loaded is true.

Let's save that.

Go back and try it again.

And there it is so I can go to graphical click on The Shawshank Redemption.

And there it is.

So there that's simple enough to get the information using graphics.

You will now in the next lecture, we'll make the changes to the database structure to allow for a movie

poster and we'll make changes to the back end code that will actually request that information and send

it to the end user when they try to display an individual movie.

And that might take a couple of lectures, but we'll get started in the next one.

## 96 - Modifying the back end to handle poster images

So we have the front end set up to display information about an individual movie making a graphic,

you will request the now it's time to make changes to the back end.

And the changes I want to make here are to allow for and movie poster to be added to our actual list

of movies.

So the first thing I'm going to do is go to my database and I'll open up post poster girl or whatever

that is.

It is.

And I'm looking at my Go movies database and I'll open the movies table and look at the structure and

I'll add a column and I'm just going to call it poster and it will be a character area.

And if you're using Beever, it might call it Verkerk, but we want to put text in there and I'm not

going to set a default.

OK, and I will save that.

So now I have that column and if I look at the content for the movie, I should have known for every

one of the movies that are there and I do.

So that's good.

So let's go to our back end code and open models up.

Go right there and find a movie and I'll just add a movie right here on that poster right here.

So poster and it will be a string.

And in JSON, I will call the poster format everything.

And I'm actually going to set create an updated app back to created that and updated.

The reason being, if I don't, these will never be exported from my graphical fields for my schema.

OK, so we've added that now because we have this new field available to us, we have to go to our database

right here and update our get request.

So I'll just add a poster, but I'm going to do it like this coalesce poster if it exists and it's not

null.

Otherwise, an empty string that'll save me the problem of having to deal with nulls from the database,

which is always a pain when you're dealing with with GIO.

That's one of the things I do not like about this language.

But anyway, I've got this request in my query, so added here and populate movie dot poster.

And we're also going to need to make changes to the insert and update.

So let's find where we're inserting a movie.

There's all and John Resul and insert movie there.

So I'll put a dollar sign ten here for my 10th placeholder and I'll insert movie poster and for updating

a movie right here, nine becomes a ten and I'll put a poster equal dollar sign nine and then tweak

these last two items.

Movie poster.

OK, so there's our queries now we need to go to grad school and telegraph cool about this particular

field.

So down in our description of the fields, I can just add it right here.

So I'm going to have one called poster, and it is a graphical field.

And its type is graphical strength.

And I need a comma there, so that's been the last thing we have to worry about here is to go to our

movie handlers and we need some means of getting information about the movie poster.

So I'll add a function right at the very bottom of this file.

And we're not going to use search movies anymore because we're doing that through graphical.

So I can delete that stub and I'll add a function and I'll call it get poster.

And it doesn't need a receiver because we're not going to call the database here.

It will take one argument movie of type models, the movie, and it will return models dot movie.

So basically we'll pass our movie variable to this function, try to get a poster.

If we get it, we'll add that to the variable and return it back.

Otherwise, we'll just return back to the variable unchanged.

So this is pretty straightforward.

So the first thing we're going to want to do here is to describe a new type and I'll call the type the

movie DB, because that's where we're getting the information and it's a struct.

So how am I going to do this?

And also right at the end, just so we can get rid of this error return movie, I need to describe the

JSON that I'm getting back from the API that I'll be calling momentarily.

So let me go back to my Web browser and I'll make a little shortcut here.

So back in my Web browser, if you recall, this is the UFO that I'm calling and I'll post a link to

this so you can actually get it from my the course resources for this lecture.

And this feeds back, Jason.

So I'm just going to call copy this, Jason, and then I'll go to this rather helpful resource that's

also available for free online.

And there'll be a link to it as well.

What it does is it takes JSON, which I'll paste in the left side of the screen and automatically gives

me the correct struct that I can just copy and go back to my code.

This is a really useful little thing.

Back to my code, which is right here.

And I'll just replace this type undefined here with the one I just pulled and I'll call it again, the

movie DB.

So that describes the JSON that I'll be getting back.

Now it's just a matter of making the call to the remote API, which takes a little bit of code but isn't

difficult.

So we'll create a variable called client and that will be assigned the type of a reference to a client.

Right there

and now I need to put in my key, my API key, so I'll make that a string and I'll go back to my Web

browser and find my my API settings.

And again, you get there by going to click up here once you're logged in choosing settings.

So I click on this choose settings.

That takes me to this screen.

I click on API and you want the API key version three, so I'll copy and paste it in here.

And I'm actually not going to put a copy of this on the course resources because I want you to want

to encourage you to go get your own API.

The reason being, if everyone uses mine, it might exceed the rate limiting that's built into the API

and it might not work for you.

So if you want to use this, just type it out, pause the video, type out the API.

But it would probably be faster, simpler and safer for you to go create your own account on that website.

So now I have a client.

I have a key.

Now we need to build the URL.

So I'll call it the euro and don't call it a URL because that's a built in package.

So I'm calling my value URL and that will be equal to that.

And I got this link right from the API documentation API dot the movie Debug three, because we're using

version three of the API search movie Questionmark API.

Underscore key equals just like that.

And then I'm just to make sure everything works properly.

I'm going to say logged on pretty long just to print this.

So I'll print it out like this though.

You url plus key plus ampersand query equals and then I need to encode the name of the movie, the title

that will be passed as part of the movie variable to this function.

I need to encode it for you or else which is really easy and go use the URL package and you're looking

for query escape.

I've got to put a plus in other, otherwise I'm not going to get my autocomplete query escape there

and inside of that it's MoVida title.

And the second argument is no, look, I'm just picking that up.

So there is no second argument.

Now, that is just to print that you are allowed to the console.

So if there's a problem, I can copy and paste it and see what the error is.

Now we make our request, request an error or assign the value of a new request and we're going to pass

it the type get because it's a get request for this API.

And I just put in the euro, the euro plus key plus in quotes, ampersand query equals plus euro dot

query escape and it is a movie title.

And now I pass the second argument of no.

So I have that.

Let's check for an error.

If error is not equal to nil, something went wrong.

So I'll just say logged uprate long and print the error and now I'll just return the movie.

So this won't fail.

It'll just return back to the movie with no poster.

Now we'll add a couple of headers rec dot header dot ad and we'll tell the remote API.

We're accepting applications, Jason, and we'll add another header content type

application JSON again, that's just good practice.

Then we'll check for a response, response and error or assign the value of client.

Don't do request.

And again, we check for an error.

So I'll just copy this and paste it here.

And if there's no error, you of course don't want to resource leak, so you defer resp body close and

then we get the bytes from the body body bytes I'll call it, and error or assign the value of IO.

Don't read all and I'm using the IO package.

Let me import it.

Important for me.

Don't read all and I'm reading the response body and again check for an error.

So paste what.

It's already on my clipboard.

Now we'll create a variable which I'll call var response object and it's a type the movie DB.

And unmerciful my Jason, Jason and Marshall, and we're Amish and martialing body bites into our response

object, and now we check to see is there actually a valid response in there?

Now, if you look back up at the type which was helpfully created for us by that website, you're actually

going to have an array called results.

And inside of that you will find poster path.

So all we have to do down here, once we have mercial, it is to say if the length of response object

results is greater than zero, then we have a movie.

We have a poster.

So we'll just say movie poster equals response object.

The results, the first index, we'll just take the first poster, we get a poster bath and now at the

very end we return movie so I can just delete that format, everything, and then all looks pretty good.

So this is pretty simple.

It's fun.

We can call from wherever we need to.

So now let's go back up and find our edit movie function, which is right here.

And we could just put a simple check in, if we already have a poster, we don't want to go get it again.

So we'll say if movie poster is equal to nothing, which is what we'll get from our query if we call

the database and it will be nothing by default, an empty string by default or for creating a new one.

Then we say movie equals movie equals get poster and had a movie and that should be all that we need

to do.

So let's try compiling this.

So I'll start my application.

Start my application.

Everything of piled the next step is to go and update our front end to request the poster using graphic

you and show it if it exists.

And that should be pretty simple and we'll take care of that in the next election.

## 97 - Updating the front end to display the poster image

So we have our back and set up properly, and now it's time to modify our front end to display movie

posters if they exist.

So I'm looking at the front end source code and I'm looking at one movie graphical.

And the first thing I'm going to do is add poster to my graphical query.

And that's as simple as typing poster out at the back end is set up properly.

And the next thing I have to do is to come down to the render part and actually display that poster,

but only if it exists.

I'll do it right after the title.

Don't do it down inside of this float part because you'll screw up your floats.

So I'll put it right here after the title.

And what I'm going to do is do a check.

Movie poster is not exactly equal to an empty string, then the double ampersand, then the parentheses.

And if those conditions are met, in other words, if movie poster has some value, then I'll display

the image and I'll do it in a day.

I'll just representative.

And the syntax for displaying the image is pretty straightforward.

Image source equals.

And it's going to of course be like this because I'm going to use a JavaScript template and you know,

I got straight from the documentation for the API, which is really easy to find, and it goes like

this htp colon and also HDB colon slash flesh image dot tm, DB Dorji and then a slash and then T and

then a slash and then P and then slash W how many pixels.

Why do you want this to be.

I'll put two hundred and when we get the value back from the remote API, it already has the slash that

would come after W two hundred.

So all I have to put is dollar sign then curly brackets movie poster and outside of that entire curly

bracket I'll give it an all tag, all equals poster.

And if you don't react will complain and say it's not correct, it'll give you a warning and then close

it.

Now you would think that this would be enough.

So let's make sure everything compiled and it did.

And let's go back to our Web browser.

Here we are and I'll click on Graphical and I'll click on, say, the Princess Bride.

Perfect.

No errors.

There's no image.

But I expected that because we need to actually make a change to this record so it will check to see

if the image exists.

And if it does, it should update our database.

So let's go look at it, manage catalog and click on The Princess Bride.

And all I should have to do at this point is click save and it will make the request to our back end

and attempt to find a poster for the movie brought and update the database.

If it does so, let's save great.

No errors.

Let's go look at our database and there appears to be a value in there.

OK, so really you would think at this point I could go to grad school and click on The Princess Bride

and see my movie poster and we're not going to and we'll see why in a moment.

No movie poster, so why not?

I mean, I've added the necessary field to my graph, to a query.

I'm returning it from my back end.

It's all set up properly.

Now, let me show you why this isn't working.

So I've got a visual studio code and it's not in the front end.

Let's go look at the back end.

And I'm looking at movies, the movies, dash dbag go the part where we're querying.

And if you'd notice here, we made a change to the insert, to the update and to the get methods, but

we never did it all.

And back on our graphics will don't go right here.

The function movies, graphics will it populates this variable movies by calling the old method.

So let's go back to the old method and just fix this.

It's really easy.

After updated I did my query.

I'll put coalesce and if you don't do this, you'll get no no errors from the database driver for go

and if it exists use poster otherwise using empty string.

So I added that to my query.

That is the last thing I'm adding there.

So let me come down here and put ampersand movie, dot poster and a comma.

I will open my terminal, stop my application, start my application and go back to my Web browser,

go back to graphics, you will, and click on The Princess Bride and there is the poster.

Perfect.

So literally all I have to do now is receive every entry in my database and it should if it can find

the poster and display it like this.

So that was pretty straightforward.

A lot of steps necessary to make this take place and a lot of error checking, but not that difficult.

So there you go.

## 98 - Cleaning things up

So before we finish off this section, there's just a little bit of cleanup I want to do, and I have

both the back end and the front end running.

And I'm looking at the graphical component, and I'd like you to pay attention to this, the console.

Let me clear that off as I type search things here.

So if I type, for example, t it makes a request H it makes another request.

Now that's OK.

But honestly, you probably want to limit your search to three characters or more or something like

that.

So let's take care of that right away.

So let's go back to our front end code and I'm looking at Graph Keywell and I'm looking at the handle

change function.

So right now, no matter what happens if we type a single character, we're always calling this DOT

perform search.

So let's make that a little bit better.

Let's say instead, if value, length and value is what we're getting from the form component or the

form entry on the page, if that's greater than to then perform the search.

So I'll do that.

And I'll say else, this dot set state movies and I'll just set that to an empty array.

OK, so let's format everything and see how that looks.

So I'll go back here and I'll clear the console again and we have the full list of movies and if I type

t h p then it gives me the list back.

And that may not be exactly what you want, but that gives you some indication as to how you can make

the search seem a little more realistic.

So if I backspace over everything, it clears it off and now if I type K and I, it doesn't return anything.

So another thing you might want to do is to because I have to use uppercase K and I g and there it shows

up.

You might want to modify the back end and the front end to ignore case.

And the easiest way to do that is to convert the case of what you're searching and the case of what

you're searching for, to be the same thing.

So convert both to lowercase or both to uppercase, and that's pretty straightforward.

And I'll leave that as an exercise for you then.

The other thing I want to take care of here is on the home page.

We don't need to have this admit one on the bottom.

So let's just go back to the code for that and find the home component home mortgages and get rid of

this part right here.

Div class name equals tickets and this HRR, and that's just a little bit of cleanup.

So there you go.

Now, we have a functional backend at this point.

We have two means of getting information from our rest API, one using standard Jason and one using

graph.

Q Well, and both have their their strengths and their weaknesses.

And which one you use is entirely up to you.

Or you can if you wish.

And I usually do implement both.

So let's move on.
