# 02 - Building the Front End

## 25 - What are we going to create?

So for the remainder of this course, we're going to focus upon building this application and it's an

application that gives a little database of movies and it's very straightforward.

The front end is built in react and the back end is built in go.

And here's how it works.

Obviously, not a lot of effort is going to be put into making it look pretty.

Instead, we're focusing upon functionality.

So we have a home page.

We can display a list of movies.

We can click on an individual movie, say The Shawshank Redemption and get some information about it.

And we can also browse movies by genre and all of these the movies and the genres page.

And when you look at an individual movie, say crime, it shows you The Shawshank Redemption.

All of that information is coming from a standard rest API where we send a request off in the form of

Jason and we get the response back in the form of Jason and we consume that on our front end.

This part graphical actually does pretty much the same thing, but it accesses the back end using graphical

and it makes a graphical query and sends that to the server and receives a graphical response.

So, for example, one of the things we can do is look at American Psycho and we spend a little bit

more time on this.

So, for example, I can look at this movie and now we're actually fetching a poster for this movie

from a remote API and displaying that when we display the information and back on graphics, well,

we can also sort things off.

For example, type the it it filters the list to only show the ones that have the letters in the title.

So we'll we're doing that.

We're also going to have a log in functionality with one user, nadhir dot com and a password.

And when I log in and watch what happens over here and watch what happens up here, those change.

I now have a managed catalog link and I have an add Movielink and the login change to log out and that

authentication is done using JSON Web tokens.

So we'll spend some time on that as well.

And once I'm logged in, I can actually modify the catalogs so I could bring up the Princess Bride and

change its release date or its title, its runtime, the description, all of these things.

Or I can delete the movie if I choose to do so.

So that's what we'll be spending the next few sections of this course working on.

## 26 - A note about React Router 6

See note in the top readme file.

## 27 - Creating our front end application and introducting the React Router

So let's get started by creating a new react application.

So I have my terminal open and I'm in a visual studio project, which is where I'm going to store this.

So I'll create our application using NPCs, create react app and I'll call it go movies.

So that will create our application, which will take a little while.

OK, so our project is installed now I'll start up a visual studio code and I'll close this window and

browse to the folder I just created, which is in visual studio projects.

And I called it Go movies there.

So I'll open that.

Now, there's a bunch of files in our source folder that I don't need, so I don't need a report where

vitals or set up tests or logo and app test and I want to get rid of access as well.

I don't need those, so I'll simply delete those moved to trash and let me go to our indexed Jaspal

and look, get rid of this line because I don't need it and I'll get rid of this because I don't need

it.

We're going to start with a really simple application.

So that looks about right now.

Let's go to ABC's I'm going to get rid of this import to a logo because I don't need that.

And I'll delete all of this and replace it with what I want in a moment.

Now let's put some HTML in here.

So this is going to be a really simple application and I'm going to use bootstrap once again.

But this time I'll just imported into my indexed HTML file right here.

Let's get rid of all these comments because we don't need this.

And I'll call this go watch movies just to change the title, OK, and I'll get rid of these.

Leave the ID root there because I need that and get rid of this comment.

OK, now I want to add bootstrap to this, so I'll go to my web browser and let's go get the and link

for bootstrap, get bootstrap dotcom, get started.

I'm not going to party using NPRM this time.

I'll just import using success because I don't do that very often and it's nice for a change.

So I copy the link and I go back to Visual Studio Code and I'll put that in right here.

All right.

Now I have access to bootstrap, so let's save this and close that and let's get started on our application.

So what I'm going to do is create a very simple layout and I'm not going to put much effort into making

it pretty right now.

I just want to get it functional.

So let's create a div class name equals container, which is the bootstrap container.

And inside of that I'll put a ro div class name equals ro with a lowercase R and instead of that I'll

just put a title H one and I'll give it a little margin.

On the top class name equals margin top say three and then I'll put in a title, go watch a movie which

isn't a very good title, but it's what I have to work with right now.

And I'll put a horizontal rule.

And again, I'll give that a little bit of space.

Class name equals say M B dash three margin of three on the bottom.

OK, so there's our first row.

That's just the title row.

Now we'll put our content in area and give class name equals row

and I'll put it in a div with a class name.

Equals, say, call me to just the narrow column on the left hand side, and this is going to be where

I'll store our navigation.

OK, and then I'll just put another empty column in div class name equals call and 10 the other 10 of

our 12 available columns.

And in that will be where we actually put the content that we want to display when we click on links

in the left hand column.

OK, so that's enough to get started with that.

So that's all this is going to do.

Now here's the thing.

I could, if I wanted to in my navigation area, put something like this in, I could put you well or

put an nav nav tag in first, which you're supposed to.

And I put you will class name equals list dash group, which is bootstraps nicely formatted list, and

then put a couple of Ellyse in their selye class name equals list group item.

And in that I could put a link, say a graph equals let's say we're going to link to the home like that

and then I could just put home fine.

Now a copy this and paste it twice just to give two more links and modify them after the fact.

So this one would be say movies, a list of movies, and I'll put here the title movies and down here

I'll have instead of movies, I will have admin and I'll call this Manege catalogue like that.

Fine.

Now, if I was to actually try to run this right now and I can do that, let's see if I have everything

right npm start.

OK, I do have one mistake that's access shouldn't be here, so I'll just get rid of that, OK, and

try this again and go back and reload this.

There it is.

There.

That's our lab.

So this will be our navigation.

This is where we're going to display the content related to a given navigation item.

But how do I make these links work?

I mean, slash movies, slash admin, how do I route that to whatever I want to display?

Well, that's a bit of a problem because so far we know how to display components on a page, but we

don't know how to link to a particular component or have a link tied to a particular component.

And that's where this package comes in, the react router.

This is a very, very popular and widely used add-on for react that allows us to actually link to different

things, to do exactly what we want to do.

So this is basically a collection of navigation components that allow us to very quickly and very easily

create links in our reactor application and route to those.

So we're going to use that package.

So I'll switch back to my ID and I'll open a new terminal window and let's install that component.

And you do that very quickly.

NPM install, react router dom.

So that will install it for us.

So it's installed.

And now how do I use this?

Well, it's actually pretty straightforward.

First of all, let's import some things that we know we're going to need in our apgar's.

So back on line one, I'm going to import react, which I'm definitely going to be using, and I'm also

going to import fragment, which I'll probably be using.

And if not, I'll remove it after the fact from react.

OK, now I want to use the router.

Now I'm going to import and there are two potential routers that we can use from react router dom.

There's the browser router and there's the hash router.

So let's use browser router as router, just an alias.

And I'm also going to import a few other things, switch route and link.

And those are all going to come from the package we just installed from react router dom like that.

OK, now, inside of our app function here or have function app, I'm going to actually change this

to just to make it a little simpler.

Export default function app, OK, which means I don't need this line down here just a little clearer

when we're going to use this router.

The very first thing you have to bear in mind is that everything you want to be wrote, it must be wrapped

in this tag reader with a capital R and I'll take the closing tag and I'll move that way down to the

very bottom of this particular X.

OK, so we have that.

The second thing to bear in mind is that you don't write links like a you don't write them like this.

Instead you do something like this.

So I'll leave that there for right now and we'll get rid of it in a moment.

Use this syntax link, which is converted to the appropriate link after the fact.

And I want to link this to home, which is slash.

OK, and then I put home.

That's how you create links.

So let's copy this and just paste it and fix the other two.

So I'll paste it here and I'll paste it here.

And I want this to link to movies, which will be our list of available movies, which we're going to

get from our go back in before we're done, but not in this lecture.

And this one will go to Edman, not delete that and call this Manege catalog.

Like that.

Now, where do I want these links to be or fix that HDMI out there?

Down here in the column where I'm going to have the main content for my site, I'm going to use another

react router, DirecTV called Switch, and it works just like a switch statement.

So it has an opening and closing tag.

And I'm going to say route path equals movies like that.

And that's going to call movies, which doesn't exist yet.

But we'll get to that in a minute.

Now, I'll copy this for other two paths and paste them in paste and paste.

And the second one isn't going to go to movies.

It's going to go to admin and it will call and render the component and admin.

And the last one will be our Hummer, and that will go to home.

OK, so now I have to have something that will actually take the place of these little tags so I can

do that pretty easily for right now just to get it up and running.

Let's just add some functions.

So I'll add a function called and these are just standard functions and they're just placeholders for

right now.

I'll call one home and it will just return to home and then I'll copy this and create two more, one

for movies and one for ADMET

movies.

And luvvies, and I'll call this one admin, and the title will be Manege Catalog.

OK, so let's open our terminal and see if everything is compiling the way it's supposed to.

So I'll switch to the one that says node.

OK, so fragment is undefined, but never use.

That's fine.

We'll be using it before too long.

So it compiled.

That's good.

Let's close this and switch to our browser right here.

So now when I click on movies, this says movies.

And now when I say Manege catalog, this manages the catalog.

That is an excellent start.

So if you go back and look at this, let's just go through how this is working again.

So Index Dot just does nothing more than render the app component abcess right here.

And this is a simple default function that exports some HTML, some joysticks, which is then changed

into HTML.

But this package is actually using browser Rohter from the REACT router Dom.

And inside of that, of course, we have to wrap everything.

We want to be rooted in our Retter tag and then we just define our links using the syntax link to whatever

path name you want to link to.

Down here where we want the content to show up, we have our switch statement and we list all of our

possible routes and there are only three.

Now, this is using browser routing.

So if we go back to our Web browser and look at this and look at the URL, if I go right to the home,

it's just the very you are l if I click movies, it changes to movies.

But let me go back to home and come back here and change this too harsh rhetoric and save the changes

and go back to a web browser.

OK, it's reloaded, but look at the URL.

It's changed to localhost three thousand slash hash slash and when I click on movies it goes to slash

hash slash movies in the same way you can probably guess what this one is going to be.

It ends with admin.

Now, this is using hash routing.

Now this is far and away the easiest way to build and deploy a react app because it requires no configuration

of the Web server.

If I go back and use this instead, if I say change this back to browser router like that and save it

and go back to my Web browser now when I go to home and actually would require that I make some changes

to whatever Web server I'm deploying this on and it's really straightforward to do it.

You can configure Apache or Engine X or Katti or your favorite Web server.

There are lots of examples for how to configure a given Web server for this kind of routing.

But right now I'm going to leave it back as Hasharon or just because I like that right now and it doesn't

matter which we're making, you can see how easy it is to change from one to the other.

OK, so we have an application that's running right now and it does have three Stubbe functions and

all of these will go away and be changed to something more appropriate as we go on.

But at the moment, we have an application, we have a default web page and we have three active links.

So in the next lecture, we'll see how to configure those links to render an actual react component

instead of just a simple return and then a little tiny bit of X.

## 28 - Routing to a component

So this time around, we want to figure out how to relate to a component, so, for example, if I click

on a particular movie from a list, I want that to go to a particular component.

And the same way, if I click on Manage Catalogue, I don't want to go to this Stubbe function.

It mean I want to have a new component that renders all of that for me.

So the first thing I'll do is go to my source folder and create a new folder, which I'll call components.

And I'm going to store my components in here, so let's create one, a new file, which I'll call movies

jazz, and that will be our movies component.

OK, so I create that.

And inside of that I'll import react and I'll also import component from react

and then I will the same as we did before export default class.

And this is called movies and it extends component and we have to have a render function and that's

all we'll have right now and we'll just return some RSX.

And we'll just put it to and I'll give it a different title so we know it's actually doing the correct

thing.

Choose a movie.

And back in our outguess, before we can use that new component we just created, we have to import

movies from and this time it's components, movies.

OK.

And down here, we can't have a function name movies if we're importing an exported function, name

movies, so we can just delete that.

And if everything worked as expected, once I save this and switch back to my Web browser when I go

to movies, it should give me the new title.

And there it is.

We're rendering a component using our REACT router, which is pretty straightforward.

Now, obviously, we also want you to display a list of movies right here, and then we want to have

an individual component to display the details for a given film.

And it'll be a little while before we get there.

But that is the direction we are going.

## 29 - Challenge: Route to components

So from this point on in the course, I'm going to try to give you more regular challenges, things

that I'll ask you to do and give you a chance to do them, and then I'll go over my solution in the

following lecture.

And this is our first challenge.

And I want to start with a really easy one.

So in the last lecture, we got rid of that little stub that showed the movies, page that stub function

and replaced it with a component.

And I want you to do the same thing with these two Stubbe functions, the home function and the admin

function.

Those should actually render true react components.

So just a couple of hints.

No.

One, you'll notice that I stored the component I created last time in the folder components and I would

put them in there if I were you.

But it's entirely up to you.

So get rid of these two functions and create two new reactor components and have those roaded.

Two, when people click on the appropriate link, so I click on slash admin, I should go to the component

that you create.

And the same is the case for the home page.

So give that a whirl.

It shouldn't be too difficult for you.

And I promise I'll give you more challenging challenges later on in the course, but I thought we'd

ease right into it.

So give it a go and I'll go through my solution in the next lecture.

## 30 - Solution to Challenge

So how did you make out with the challenge?

I'm sure you didn't find it terribly difficult.

I have it half done and I'm going to go through the second half in this lecture because I want to show

you a couple of shortcuts, things that you may have stumbled across already, but things that I normally

don't introduce to people until they've been writing the reactor components and so forth by hand for

a little while.

So the first thing I did in my components folder, I created an admin dark's, and it's pretty straightforward.

It's virtually identical to movies.

Torture's just the name of the class and the content of the two tag changed.

Then back in my apgar's, I imported that component and then I simply just deleted the function named

admin down here.

And I want to do the same thing with the home function.

We'll get rid of that in a moment.

So I've already created an empty home dojos in my components folder and it has nothing in it.

And here's the shortcut I was talking about a little while ago.

Do you recall that extension we installed for Visual Studio Code lo these many lectures ago?

Well, one of the reasons I installed it is because I can do things like this.

I am.

And then notice you have some autocomplete here and the very first one is I am Arcy import react component.

I can just type, I am Arcy and return and it types that line for me, which is pretty straightforward.

I'm going to put a semicolon after it.

So let's export our default class and we're going to call this one home and it extends component

and we have to have a render function render and it just returns some ESX.

So I'll give it an H2 with the title and I'll just call it.

This is the home page.

I will replace this with more meaningful content a bit later on.

So I say this.

Go back to us, import home, so import home from that components home and give it its trailing semicolon.

And then of course I have to delete this function.

And if everything worked according to plan, once I save this and go back to my web browser when I click

on already did it for me.

This is the home page.

So Hollingsworth's moves these works and manage catalog works great.

So that's pretty straightforward and nothing too difficult in that at all.

What we'll do next is when I click on the movies link, I want to display a list of movies here because

eventually we're going to be getting a list of movies from our go back end as adjacent file processing

that and displaying the list of movies in our database on this screen.

So we'll get started on that in the next lecture.

## 31 - More about routing (and a bit about the React lifecycle)

So this time around, we want to talk a little bit about nested rooting, and it's going to take us

a while to get there, probably a couple of lectures, but let's get started.

So this component movies shows up here when I click on movies in my application.

And what I want to do is display a list of movies.

And as I said last time, ultimately that list of movies will come from our go back end as adjacent

file.

So I need to simulate getting adjacent file or getting something that can be converted into something

I can use in this lecture.

So we're going to do that by going back to our code.

And the first thing I'm going to do is import fragment, because I know I'm going to need that in a

minute.

So fragment.

And the next thing I'm going to do is set up some state.

And for this state, I only care that it's available in this component because this is the only place

I'm going to use it right now.

So I will say and we might change that later, but right now I'm just going to put it here.

I'll say state is equal to and I'll have a key called movies and its value will be empty to start with.

OK, now down here in this render function, I'm going to return this and I'm going to make this a little

bit easier by wrapping this in parentheses, like the just to give it a little more cosmetic appeal.

And I'm going to wrap everything in a fragment.

So fragment and my closing fragment tag gets moved down to the bottom like that.

OK, let me format this there.

Now, inside of this, I'm actually going to like we did in the last section, put a U.

L in here and opening you will tag in a closing you will tag and I'll just use the map function to go

over the movie's key in my state.

So this dot state dot movies, dot map and then I will use the variable M to keep track of the current

iteration and use my arrow functionality and open a parentheses.

And inside of this I'll just have elai key equals and open closing curly braces and I'm going to call

my ID which will generate in a moment.

So don't worry about that.

Then we'll put the name of the movie which will just be my title.

OK, now I need to populate my state, my my movies key in the state variables somewhere and I can do

that really easily.

I mean I have an empty one up here, so this will just generate generate an empty you will tag an opening

and closing.

You will type with no Ellyse in between.

But I actually want to put a few movies in there.

So what I'm going to do is after I declare this, I'm going to call a new function something that's

available to us and something that's part of the react life cycle.

If you recall in the last application we built, you had to click on a button to fetch some content

to display it on the screen.

And of course, I don't want users to have to.

And let me go back to my Web browser.

I don't want them to have to click on movies to say choose a movie, then click on a button to get the

list of movies, then click on the movie they want.

That's an extra step.

When I click movies, I want the full list displayed here.

So how am I going to do that?

Well, we're going to do that by using a function that's related to react lifecycle, which I'll talk

about in a moment.

But let's right at first the function is called component did mount like that and it takes no arguments

and this will get executed after the component is rendered to the screen.

And that's part of the react life cycle that we need to be aware of.

And I'll show you a helpful diagram and give you a link to it in a little bit.

First of all, let's just right here is where I would say the component did mount.

Now go to a remote server, get the list of movies and populated in the state.

We're going to fake that by just saying this set state

and we'll just write some movies in here.

The key is movies that, again, movies and it's equal to.

And I'll just put three movies in here, so I'll put in one with an idea of one.

The title will be, say, some movie, The Shaw Shank Redemption.

Actually, I have a list of them next to me and it's runtime is one hundred and forty two minutes.

OK, so there's my first one and I'll just duplicate this twice and change everything.

So make this to make this ID three and instead of The Shawshank Redemption will make it The Godfather,

which has a runtime of it's a long movie.

One hundred and seventy five minutes.

And finally, the last one will be The Dark Knight, just enough to give us some data.

The Dark Knight, which has a runtime of one hundred and fifty three minutes.

OK, so I've now set this up.

So really, when you think about what's going to happen here, we're going to click on from our.

Where is it here from our abduh, as the users will click on this link right here to go to movies,

which gets translated to either the movies, if we're using the browser router or the one with the hash

tag in it, if we're using, as we are, hash router and that takes them, that actually renders this

movie's component.

When this component renders after it's mounted, then we set this variable in the state called the key

in the state variable called movies and populated with this information.

And then it just writes this information to the screen.

So if I go back to my Web browser right now, you can see it's already working.

I've got a home and I go to movies and there it is.

It lists all three movies.

So what we'll do in the next lecture and I'll talk a little bit more about the life cycle before I end

this, but we'll do the next lecture is actually make these active links, and that's where nested routing

comes into play.

But right now and I'll talk more about this as we go on through the course, this is a really helpful

chart that shows you the life cycle of a react application.

So we have the mounting, we have the updating and we have the unmounted.

And really, when you look at what we're doing here, we actually call the constructor first, then

we render it, and then we set the state and we actually update the domes and reps at this point.

So this this function right here component did mount is called at this step in the mounting phase.

When we update things, there's the order that things take place from top to bottom.

And when we unmount them, we only have to worry about this.

Now, if you work and react a lot in the future, you probably should check this little box that says

show less common life cycles.

That just gives you a bit more detail on things that you normally don't run into when working with a

react application.

And we won't in this course.

But if you're doing a deep dive or working in a really sophisticated react application, at some point

this part of the chart will be of value to you.

All right.

So for right now, the only thing I'm interested in is when does the component did mount fire?

And it fires after the constructor, after everything is rendered, after react updates the dome and

its references to the dome, then we know the component is available.

So at that point, it is safe to write our our information that we just hardcoded those three movies.

It's safe to write that to the browser window.

OK, all right.

In the next lecture, we'll start actually getting some nested rooting working by making these active

links.

## 32 - More about routing Part II

So things are coming along and we want to talk a little bit more about routing and nesting.

So right now we have a home page that works.

We have a movies page that lists movies.

And I want these to link to a page about each of the individual movies.

And we also have this Manege catalog.

So let's go back to our code.

And the first thing I'm going to do is make a change to the way that we're routing.

So in abcess, just because it's clear, it takes the hashes out of it.

I'm going to change this back to browser rubber and save that go back to my code or my browser and I'll

just click on the homepage.

So now we don't have that hash messing things up and this will just make what I'm going to try to do

a little bit clearer.

So let's modify the bit of code that produces this list to make this a link.

And we'll just take it to a particular page that just shows the idea of whatever link we've clicked

on.

So let's go back to our code.

And the first thing I'm going to do is in our switch statement, I'm actually going to add before slash

movies.

I'm going to add another route and this one will match path equals slash movies.

And I'm going to put a colon idea in here.

And that's a kind of placeholder that's actually going to be substituted by whatever idea it happens

to be.

We use numbers just to make things clear.

So back on this page, for example, we'll have this big one.

The Godfather is already two and The Dark Knight is already three.

OK, so I'm going to inside of that, I'm going to link to something that doesn't exist yet, which

is called movie, and I'll just make that a simple function.

So down at the bottom of this file, which is Apgar's or create a new function just called movie.

And what I need to do is to get the I.D. That's part of the URL.

So slash movies, slash one, for example, that I can do that by saying let and then in curly brackets

I'll just put it, which is what I'm looking for, and that matches the colon ID in the path that we

just typed in.

And I'll make that equal to the function that we get from the reactor recruiter use programs.

OK, and now I'll just return just to make sure this works H2 and I'll say movie ID and then just ID

what that.

OK, so when I typed use programs there, hopefully it imported for me.

It did it right up like that.

So let's go back to our code and reload this just to make sure it works.

No errors.

That's good.

Now in our movies page where we list those movies, we need to make some changes to the code there to

actually build the URL for us.

So let's go to movies just right here.

And this is a react component and we're using that component did mount function to simulate a call to

a remote API.

So down here, we need to link this title.

We need to build a link there.

And of course, since we're using the router, we're actually going to use the link function link and

we'll say two equals.

And then you'd think you could type it like this movie, slash whatever it is the ID, but you can't

that actually won't work.

Instead, you have to use this syntax in curly brackets with a tactic, slash movies, slash and then

dollar sign open curly bracket and ID, which is our movie ID, and then we close our back text.

It's just a JavaScript template and we close our curly quote and we close our tag and I'll move this

link tag to the end.

No link is showing up as an error because it's not defined.

So back up at the top, let's just import it.

All we need is link.

And we're importing that from react router dom right there, OK?

So if I save this, there's no errors.

Let's go back and look at our browser and see how it works.

Now we have links.

And if I roll over the first one, it should say slash movie, slash one in the lower left hand corner.

And it does.

Then that's movies, too.

And that's movies three.

Perfect.

And if I click on one of these, let's see if our switch statement actually works.

Movie ID three.

Let's go back to movies.

The Shaw Shank Redemption movie ID one.

Perfect.

Now, this is actually an example of nested routing because we don't have one.

We have movies, slash one slash movie, slash three, so on and so forth.

But just to make it clear, in the next lecture we'll add another menu item here by category that allows

us to look at movies simply by category.

So comedy, drama, whatever it may be, and we'll get started on that in the next lecture.

## 33 - More about routing Part III

So this time around, we want to add a menu item on the left here, maybe just below movies, and it

will allow us to browse our movies by categories.

So drama, comedy or whatever it may be.

So this will give us a couple of things we can look at.

The first is more experience with nested retting, and the second is since once we click on categories,

we'll have the categories listed here.

We should be able to switch between categories.

So it will be a list of all the available categories and we'll just put a couple in for right now.

So we're going to have to have that page and we're also going to have to have the page that lists the

movies that match a certain category.

So this will give us a couple of things we can learn about.

And one is more about nested writing.

And the second one is how can I pass information to a component that's a true react component using

the react router.

So let's get started.

So go back to my code and I will look at approaches and I'm going to put my link right here right after

the link to movies.

So I'll copy that and paste it and just modify it.

And I'll make the URL by category and I'll call the label movies, which is by category.

Maybe just categories be as concise as I can.

All right.

So I have a link that I have nothing in my switch statement that matches that.

So let's put one in here and I'll also introduce another word to you, another key word we can use here.

This will be root.

And I'm going to add the keyword exact and that will say only match this exact route.

Now, you may have noticed if you were playing with the react router that the order actually does matter.

And that's because when it's matching things, it doesn't actually look at the entire you URL.

It just looks at the first part of the URL.

And that can be a little confusing.

And the word exact, which is useful if you have a short menu like we do, actually forces the router

to match exactly what's in the path statement.

So I'll save match exact path, slash Vikash category.

And inside of that, I'll route to a non-existent function, which I'll call category.

Now, that doesn't exist.

So let's go down to the bottom and we'll just create a function here.

So we'll call the function category page.

And of course, it has to have parentheses and we'll simply say we're going to have a return statement.

We'll get to that in a minute.

But here's where I can use some of the information that's available to me because I'm using the react

router.

So I'm going to create two variables here.

Let and I'll just have them automatically filled from what the function I call returns Tath and you.

OK, and those will be equal to whatever this builtin function use root match returns.

And hopefully it imported that for me.

Let's find out.

Yes.

Use root match code imported up here.

So this function actually returns some information that's available to us because we're using the react

router.

So we have Path and YORO.

Now, path lets us build route paths that are relative to the parade route.

And you, Earl, lets us build relative links.

So we'll explore that right here.

So what are we going to return from this?

Well, all I want to return here is a list of categories I will just put a couple in just to test things

out.

So I'll return.

I'll wrap everything in it.

Do I'll return, first of all, the title H2 categories.

And then underneath that, I'll just put an unordered list and we'll worry about making it pretty later

on with a couple of Ellyse in there.

So the first one will be a link.

So I have to use the link keyword that's available to us from the react router it's going to link to

and what I'm going to link to I'll just give myself some space here is has to be in curly brackets because

I'm building this string, OK, and I'm using a JavaScript template.

So back tactics and I'm going to link to dollar sign path.

The variable I got right up there on line seventy two slash and we'll start with drama.

OK, and then I'll close this and say drama and I'll just duplicate this line and change this one to

comedy and I'll call this one comedy.

OK, and I'm not using you are all right now I'm using path so let's go see what that does.

Let's go back and reload this page and click on categories.

And now I have a link to comedy and drama.

I want to inspect this just so I can see what was generated.

So the link here is by dash category comedy and the next one is by dash category slash drama.

And that's because if you look back in our code I used path and path lets us build route paths that

are relative to their parent route and the U.

Or L instead lets us build relative links.

Now, sadly, because we only have one level deep here, if I change this one to you earlier, shouldn't

be any difference.

We'll go back here and reload this page and inspect it again, so I'll inspect this one.

And that is by category drama and the one above it is by category comedy.

And that's exactly what I expected to see simply because my application is only one level deep at this

point.

So let's make it a little bit deeper.

I'm going to go back here and up in my switch statement after by category.

I have written a couple of times just to give us some space so we can see what we're doing.

Well, create another rigt and this one is again going to be exact route and I'll put it on the next

lines, was readable, exact and on the next line path will be equal to, in this case by dash category

slash drama all hardcoded so you can see exactly how it's going to work.

And this time I'm not going to just close my route tag and put the link to an appropriate function or

whatever in pointy brackets.

Instead, I'm going to use this, this approach render inline, I'm going to render and I'm going to

make that equal to and then opening curly braces.

I'm going to pass some properties along here.

OK, here's how I can pass properties from the react router to a react component, which doesn't exist

yet, but we'll build it in a minute.

So I use my pointy syntax and now I'm going to say root this to opening pointy bracket categories,

which doesn't exist, but it will shortly.

And I'm going to use the spread operator to populate the variable props and the property I'm going to

give it is title.

And in this case, title will be equal to in curly brackets.

And then I'm just going to put it in context, drama.

So it's just a string.

OK, now I closed my categories tag and now I close my overall point in brackets.

OK, so there is that one.

And the last thing of course I have to do is on the next line.

I close my route because it's one entry kind of mistake there.

Oh, there it is.

An extra type of there.

Get rid of that.

That looks better.

OK, so catagories doesn't exist yet.

And categories we need to create a react component named categories.

So I'll put it in my components folder.

I'll create a new file called Categories Dojos and inside of that I am Ursy import reaction a component

and I'll export a default class named categories which extends component.

And it has to have a render function of course.

But what I'm going to render this time will actually include let me just type the render function.

I'll return and I'll just say to category colon.

And now because I passed that property, I can go this dot props that title.

So this should work once I go back to abcess and import import categories from that, again from dot

slash components categories.

OK, so that should work.

So let's switch to our web browser and see if this all works as expected.

So I go back to my homepage, then I choose categories and now when I click on Drama, I should see

this content area replaced with a component we just wrote taking the word drama from our props and it

works.

And if we go back to categories and choose comedy, this is the home page because we didn't do that,

Rukiya.

So let's go to a root for comedy.

Back to our Apgar's.

Find our switch statement.

I'll simply copy.

And paste, I'll move this over where it belongs, and this one will be comedy.

Comedy with a capital C. And now when I go back and choose categories and click on comedy, it should

work and categories drama and it should work perfect.

All right.

So the next thing I want to do and then we'll have enough of our front end built that we can start working

on our back end.

The next thing I want to do is to have these do more than just, say, movie I'd want.

So again, in our live application, when I click on this link, what will actually happen is our component

will use the function that allows us to get data after the the component is rendered.

We'll call our remote API and say, give me all of the information you have on movie ID one as JSON.

Then we'll pass that Jason and populate a nicely formatted page that gives us all the information about

that movie.

And of course, we're going to simulate that in our next step.

But it will give us something to work with and allow us to move on and start building our go back end.

And that's the part of this course that I actually enjoy.

I enjoy go far more that I enjoy JavaScript or react.

## 34 - Displaying one movie

So what we want to do this time is to allow when I click on an individual movie from this list, to

do more than just display the ID here and the way that we're doing it right now, let's go back to our

code and look at Abduh as when we're going to a movie.

We're saying a root path is a movie, some placeholder ID, and then we're calling this function movie,

which exists right down here.

So I'm going to get rid of that entirely because I'm not going to use it.

And instead I'm going to change this route.

I'll delete this entirely and changes to something else.

And I'm going to take advantage of rooting to a component with the react reader.

So I'll say route path equals the same as it was before movies and then the placeholder ID.

But this time I'm going to use this keyword component and it's going to go to in curly brackets and

I'll call it one movie component that doesn't exist yet.

OK, and then I'll close to tech.

Now let's go create this one movie component.

So in components I'll create a new file which I'll call one movie dojos.

And inside of that I'll go by.

I am P for I am RC just to get the correct import put in and we'll go export default.

Class one movie extends component as usual.

And inside of this I'm going to have my state.

So I'll set up state just for this component.

I don't care about sharing it with other components at the moment, so I'll declare a variable state

is equal to and we'll give it one key movie and we'll make that an empty JavaScript object.

OK, and now I'll have a render function, which I have to have and that will return.

I'll wrap everything in a fragment which should auto import, and it did, and in here I'll put, first

of all a title, so I'll put H two, which I've been using right along for titles.

So I may as well be consistent and I'll just put some placeholder text to one movie, actually will

just say movie and then ACOA.

And inside of this I want to put the name of the movie.

OK, now how am I going to do that?

Well, as I said, when we go live, when this is hooked up to an actual back and rest, API will be

grabbing that from our remote server.

So we'll do it the same way we did with the other list of movies we used to component did mount

and we'll just fake calling.

An API will say this set state and the key is movie and we'll put some information in there.

So in curly brackets we'll give it an ID and will the ID that we want actually is passed to us because

we routed to a component and it's part of the properties.

And if you look at the documentation for the react router, we get it this way.

This dot drops.

So it's a prop and we're going to match from the programs in the early.

We're looking for ID and of course, ID is what we have in just right here, Colen ID.

So back to one movie that gives me the ID, so I'm just faking it right now.

The second thing, I'll give it as a title and I'll call it some movie.

So no matter what movie we click on, we're always going to get this information.

But this is just faking things and we'll give it a runtime of one hundred and fifty.

OK, so now we have some state variables we can use.

So here I could do something like this state dot movie title and then I'll put a table underneath it,

say table class name and I'll use bootstrap styling for tables, table table, compact tables straight

and I'll give it an empty head because I don't have any headings.

And in the Tebaldi I'll put in a row and a cell and I'll put in a strong title.

And on the nexted, right beside it, the total TRD, and it is again this dot state movie title,

and then I'll copy this entire row and modify it for the runtime.

Runtime will make it runtime

and put the word in its after that and back it up to us.

All I have to do now is import this component so the top

import and I want to import one movie

from Dotzler components one movie.

And if I save this and go back to my web browser and reload this.

So there is one movie.

So if I go to home it gives me the home page.

If I go to movies, it gives me The Godfather.

Now, no matter which one of these three I click on, I'll get the same information, but I should get

some information from all of them.

So there is some movie now we can actually verify that we're getting a little bit different information

because the IP or the ID is set using our props.

Let's go back to our code and go back to one movie and we'll just put here after the title just to demonstrate

that it works, this dot state dot movie ID.

And if I reload that in the Web browser, I have to for this one.

This should be ID one and it is and there should be ID three and it is OK.

So now we're at a point where we can start to work on our go back end and start generating some JSON

and sending it back to our REACT application and we'll get started on that in the next section.
