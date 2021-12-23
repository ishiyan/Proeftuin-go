# 06 - Securing routes in our REST API

## 76 - Generating JSON Web Tokens on the back end

So we're at the point now where we need to actually generate our Jason Webb token, and I should let

you know right now that I often call these JavaScript Web tokens, even though that's not the proper

name and that happens a lot in this part of the world.

For some reason.

I'm not sure why, but the proper name is Jason Web tokens.

And right now I'm looking at the Wikipedia page for Jason Web tokens.

And there's a link to this in the course resources for this lecture.

And as you can see, the token is composed of three parts, a header which says, what kind of algorithm

is used to generate the signature, the payload, which contains a set of claims and we'll be making

claims here in momentarily.

So don't worry about that.

And finally, we have this signature and that's what securely validates the token.

So in effect, you're getting one JSON file consisting of one entry that has these three parts.

So let's get started doing this.

So I'm going to switch over to my ID and the first thing I'm going to do is install a third party library

that actually gives us access to the necessary functionality to implement JWT.

So I'll open my terminal window and I will type go get and I'll just paste this in here.

Otherwise I'll make a mistake.

But it's GitHub dot p s c ldk élodie JWT.

So let's go get that.

And that will install the necessary library and added to our Gombauld file.

Now the next thing I'll do is set up a route, a route that points to a handler that doesn't exist yet.

But we'll add that in just a moment.

So right here, after my status, I'll add router handler func and the method will be http dot method

post and we'll go to slash V1 Slash Sinon, and that will be handled by a handler.

We have to create just a moment.

We'll call it sign in like that.

OK, now that doesn't exist yet.

So let's go over to our command API folder and I'll put this in its own file and I'll call the file

tokens.

Don't go.

OK, so let's give it a package declaration package main and let's start writing things.

Now, when you're actually doing this in production, you will almost certainly authenticate users against

a table in your database like a user's table, and that would have all kinds of information and ID their

first name, their last name, their email address, their password, maybe their physical address,

whatever it may be.

And just to keep things simple and straightforward, I'm not going to go through that entire creation

of the user's table.

Instead, we'll just make it and we'll do it like this.

I'll create a variable which I'll call valid user.

And this will stand for or take the place of retrieving a user from the database and it will be of type

models user.

Now, we don't have a user model yet.

So let me just close this and go over to my models and we'll just add one at the bottom.

So we'll call this type of user and it will be a struct and we'll just give this a few fields.

So ID, which will just be an email, which is what they'll use to log in and that will be a string

password.

And it's not their actual password.

It'll be a password hash.

You never store passwords in a database, but we'll get to that shortly.

OK, so let's format this and give it a comment.

User is the type for users that are better spelled out, right?

OK, let's go back to tokens and we have.

Let's just have this import that's better, OK, and we'll just have to use a dummy user, so I'll give

this person ID of, say, 10

and we'll give them an email of me at here dot com and then we need a password.

And the password can't be something like password.

It needs to be a hash.

So how are we going to do that?

Well, it's actually pretty straightforward.

What I'm going to do is go to my Web browser and open a terminal or new tab and I'll go to this address.

And this address exists on the course resources for this page, and it's played on Galangal.

And it's a very, very simple program that takes a password.

And as you can see here on line nine, the password I've chosen is the singularly insecure one of password,

and it generates a hash four.

So all we do is run this program and we copy the hash right here.

So I'll copy that and I'll go back to my ID and paste that hash right in there.

OK, so we format this.

So now I have a valid user.

I'm going to create another type here and this is what we're going to use to test whether or not someone's

allowed to log into the system and we'll call it credentials and it's struct and it just has two fields

username, which is a string.

And in JSON, I'm going to call this email because that's what we're receiving when people use our login

form and then we have a password, which is also a string.

And in JSON that will be called password.

OK, so we have a valid user and we have a type.

Now let's create our handler func app and it takes the receiver of APTA, which is a point of the application.

And we're calling this sign in, which is what we called it in the Roots, because it's a handler,

it takes a response writer

and it takes a pointer to a request.

OK, now we'll declare a variable var creds of type credentials, the type we just made, and we'll

get the JSON body that's been posted to us and decode that into credentials.

So we'll check for an inner ear is a sign the value of Jason from the standard library.

And we're looking for a new decoder and we're going to decode the request body and we're going to decode

it into Kritz and we check for an error.

If error is not equal to nil, then we'll just say apter error, Jason, and we'll handle the response

writer and create a new error.

Errors start new and I'll just put unauthorized.

And of course, we return because we don't want to go any further.

At that point, they didn't even give us a valid post to log in.

Now, next, this is the point where you would actually get the username and password, the email and

the password it in our case.

And you would query the database and check to see if there's a valid user that matches that username

and check the password, the password hash stored in the database against what was supplied.

So we're going to fake that by using our hashed will declare a variable called hashed password, and

we'll just get that from our valid user password.

OK, so that would be a database query in a live application.

Next, we want to see if the hash for the password they just gave us matches the hash in the database.

In other words, those two hashes could have come from the same password.

They may not be the same.

They might be entirely different strings.

They almost certainly are, but they both should be producible from the password that the user supplied.

So we do that by, again, checking for an error and calling decrypt, which is part of the standard

library.

And we're going to call compare hash and password.

And it takes two arguments, both of which are a slice of bytes.

The first one will be a slice of bytes for hashed password there.

And the second one is a slice of bytes from Krenz Password.

What did they supply?

And we check for an error.

And again, I'll just copy this and paste it here.

So if they've got past this point, we have a valid user and we have a valid password.

So now we can create our JWT and send it back.

So we're going to create a variable and now we're going to start using that package we just imported

and now we're going to start using that JWT package we just imported will declare a variable called

claims and we're going to make the claims original.

And it's a type.

JWT claims, and now we just start filling out the various claims we want to be included in our token,

so claims that subject is the first one and that's going to be equal to whatever we want.

We want to store in this.

And in this case, all I'm going to store is the user I.D. and it needs to go in the form of a string.

So I'll just say format as print and we want valid user ID and that will turn it into the correct format.

The second thing we want is when was this issued?

Claims are issued and that will be equal to now the current date and time, but it needs to be in numeric

format.

Unfortunately, this package, JWT gives us a new numeric time and we can just apply the argument of

time dot now.

And the next argument or the next to me in our claims is claims that not before this token is not valid

before.

Again, now, JWT new numeric time.

Time does it now.

And now we want when does it expire?

What?

We're going to make this expire in twenty four hours.

That should give us sufficient time to get things working.

JWT new numeric time and its time.

Now the ad will add twenty four hours so twenty four times time.

Hour.

And who issued this and who's allowed to look at it.

Claims the issuer.

And this is just the domain name and we'll use anything as long as the what the token we issue.

The audience includes the issuers domain, then we're fine.

So we'll just use my domain dot com and who can see this is claims dot audiences plural, because you

might be able to access this resta back end for multiple remote sites.

So it has to be a slice of strings and we'll just put one entry in there, my domain dot com.

So we've made our claims and now it's time to actually create the token.

So we're going to create a variable which are call JWT bytes because we're getting a slice of bytes.

We'll check for an error.

And that is a sign the value of claims dot h Mac sign.

And this actually creates and signs our token and we need to know what algorithm to use.

And we're going to use JWT dot h.

S two fifty six, which is just a constant in that package.

And we're going to pass it a slice of bytes.

And what we need to pass it is actually our secret and we don't have a secret yet, but we'll make one

in a moment.

But what I'll do is just finish writing this particular function call and then we'll go get our secret

and put it in our app config and we'll put it in app dot, config dot and we'll create a new field called

JWT and that will have a field called secret.

OK, so before we can go any further, we have to go make that secret.

So first of all, let's add the necessary fields to our config.

So back in main Dargo, we'll find our config and we'll add JWT, which will be a struct.

And I will have one field called secret, which is a straight OK.

So we've added that.

Now let's go create that secret.

And again, I have a link on the course resources page to this address.

Again, unplayed go langue.

And this is a really, really simple program.

And as you can see on line eleven, I have a singularly insecure secret.

It should be a much longer secret, but this is what's used to generate the secret.

We're going to put it in our program.

So just leave it like this and run it and copy this secret right here, the part after result and make

sure you don't get any leading or trailing spaces.

Copy that and go back to Main Go.

And what we want to do is assume for now that this is going to be read as a command line argument.

OK, so we'll just say flag, dot, string ver

and we're going to read it into our config variable JWT dot secret and it's going to be called JWT Dark

Secret on the command line and we'll give it a default value.

And this, of course, will not survive into the final version of our program.

But this is just for development purposes.

This is our default secret.

And then we'll have a little bit of descriptive text at the end and I'll just put secret there for now.

OK, so that gets us our secret.

And again, this is not how you do this in production and we'll certainly fix this before this course

is over.

Ideally, secrets like this should be stored in an environment variable.

So they're not visible when someone say at the command line types stash X or X, they can't see the

secret when they look at it this way.

So we'll start in an environment variable, but this will get us up and running for now.

So now we have to go back to our tokens.

Don't go.

And let's do something with that.

JWT bytes variable we created, first of all, will check for an error.

So let me copy and paste this.

And we're not unauthorized at this point.

We're actually going to put error signing and at some point, again, we're going to want to fix this

error, Jason.

So it's not always sending back a bad request header, but we'll do that later on.

It's a very simple fix.

And this this will get us up and running for now.

So the last thing to do is to call app, right, Jason, and write this JSON file to the user with the

token.

HDP status, OK, and we're going to pass it, JWT bites and we're going to wrap the whole thing in

response.

OK, so if this compiles, let's try running it.

Go run, dot, slash cmd, slash API that everything should be all right.

OK, and we have one thing missing here, so let's go get that.

Go get it.

And let's try running it again.

Go run, dot, slash command, slash API and our application is running.

So the next step is to go back and modify our front end code so that we can actually request and receive

that token.

And we'll get started on that in the next lecture.

## 77 - Changing App to a component, and setting up state

So now we have a means of authenticating users on the backend and we need to work on the front end to

actually create a login screen and then decide what we want to have protected and what we want to have

available to everyone.

So I have the application running right now and obviously we want this function and movie and this function

managed catalog and this delete button all to be protected.

And when I say the delete button, I mean, even though the button may not be visible, the route behind

it has to be protected as well.

So no one can call that route directly and delete movies without even seeing this screen.

So those are the things we want to be protected.

And in fact, I don't even want these two menu items to show up at all until someone has logged in.

So maybe what I'll do is I'll add a link up here that says log in and clicking on that will take you

to a login screen, which I'll write as a component.

And that login screen will allow the user to enter an email address and a valid password.

And if they successfully authenticate, they'll get their JWT handed back to them.

And at that point, these two menu items should show up and the login should change to log out.

So there's quite a few steps necessary to make that take place.

And one of the things we're going to have to think about, of course, is how to share that token.

Once we get it, we need to store it somewhere so we have access to it when we need to use it.

So obviously we're going to be lifting state as well.

So let's get started.

So go back to my idea and I'm looking at the front end code and I'm looking at app.

And of course, this is the parent component.

It's the one that loads other components.

And since we're going to be lifting state, perhaps we should convert this from a mere function into

an actual component.

Unfortunately, it's pretty easy.

So export a default class instead of a function.

And that means I don't need these parentheses and this has to extend component right there so that imports

it for us.

That's great.

And of course, this means we need to make a couple of other changes.

For example, we can't just return.

This has to be the render function.

So we'll change that into a function, get down to the very bottom here and add a closing parentheses

or curly bracket format.

This right.

So now it's a component that was pretty easy.

But of course, we want to do more than that.

First of all, let's give it a constructor,

which takes the argument props.

And then we have to call super props.

And I will set some state.

And the only thing I care about at this level is my JavaScript tokens.

So I'll just default it to this this start state equals and I'll call my entry JWT, which seems logical

and I'll make it an empty string by default.

So there's my constructor, but of course we're going to be lifting state.

So let's create a function here where we can actually lift the necessary things to our state.

And I'll call it Handal JWT change the same kind of logical naming convention I use last time and it

will take one argument the JavaScript or Jason Web token, and it just calls this dot set state.

And all we care about is JWT, which is set to JWT, OK, and now in our constructor will bind this

this dot handle to the beauty change takes the argument of this dot handle, JWT change, dot bind and

hand it this way.

OK, so we're going to need that for sure.

The other thing we're going to want is a logout function which will just create right now using the

arrows in text.

So we have access to this and I'm missing a pointy bracket there.

That's better this dot set state and all we have to do to log someone out is set JWT back to being an

empty string.

And that will be the the logic we use.

Now, we might change that later on, but I'm going to make the assumption that a JWT has a valid value.

Then the user is logged in if JWT is empty and the user is long up, which seems pretty straightforward.

So we'll go with that for now.

Now, the next thing I want to do is down here.

Let's change the top part so we can have that log in logout.

So that means the link is going to change depending upon the state.

If JWT has some value and it's a logout button, otherwise it's a log.

But so in my render function, I'll just create a variable, which I'll call let login link.

And that's just declaring the variable.

And then we'll just do a simple test set, let there if this starts straight, JWT is exactly equal

to an empty string, then set the one value else set another.

And what I'll do here is simply say Log-in link is equal to all US' link, which I'll allow my ID to

import for me if it's not there already.

And I think it is too.

And we'll we'll just use a non-existent room right now to log in and we'll start to log in.

Otherwise, we copy this and paste that link to and I'll call this one logout, which doesn't exist

and it doesn't matter because I'm actually going to call on Click

and that will be equal to this DOT logo and that will love user.

And of course, the label should be longer.

OK, so now I have this variable, I can use it.

So down here I'll simply take this H1 tag and put it in its own column.

Div class name equals call and I'll give it a little bit of margin on the top say empty dash three

and I'll get rid of this and move the closing div down where it belongs right here.

And then I'll add another call.

Div class name equals call again some merging of the top empty dash three and this one.

I'll also say put the text at the end to the right of other words

and in there I'll simply put log in like.

OK, so if I save that and go back and look at my screen now in my Web browser, I should have the login

screen and there it is, a login, but perfect.

Now, if I go back here now, if I put some values in the state, in my constructor, if I change that

to, say, some X's and go back and look at my Web browser, now, that should change to log out.

It does.

Perfect.

Now, let's also take care of these menu items.

So I'll go back down to where those menu items actually show up and I show up right here.

So I want these two allies not to be visible unless someone is logged in and we know how to do that.

That's pretty straightforward.

We can just put a condition in here.

So this DOT state JWT, if that's not exactly equal to empty and then use the double ampersands and

move this down to here and then wrap everything in a fragment.

Right here, fragment and then move the closing fragment down to here and hit return and format everything.

So right now in my state, I have it equal to Zakes.

So those should show up.

Let's go back and look up.

They're still there.

But if I remove these, I have the state default being empty and go back.

We have Log-in and those two main menu items have disappeared.

Perfect.

OK, so the next thing we want to do is to create a login component, something go back to my ID and

in components I want to create a new component.

So, all right, click on that and choose a new file and I'll just call it login dojos and I'll use

my shorthand to do the heavy lifting for me to start with.

And I'll say export default class login extends component and it's going to have a constructor which

will take props as an argument, which means I need to call super props or super props rather than OK.

And it also is going to have some state.

So I'll say this start state and what do I want to have in my state?

Well, I know I'm going to have to form values email which will default to empty and password which

will default to empty.

I'm also going to be checking for an error probably.

So I'll put that is null and I'm definitely going to have errors.

The array which will default to an empty array.

And I'm probably also going to use an alert that nice alert component we built a little while ago.

So if someone logs in with invalid information, I can give them some visual feedback and it will have

a default type of dash.

None.

So it's hidden to start with and it will also have a message which will be empty to start with, and

that should be enough to get a start.

OK, so that's our constructor.

We know we're going to have form validation.

So we're going to want to add probably the same kind of logic we did on our edit movie screen.

So I'll put in the handle change right now.

We're going to have that handle.

Change is a function and that will be equal to with the argument event, an arrow function, and just

like we did last time, will say let value equal this target value and let Nain equal this dot target

dot name.

So if we're coming from the email entry, for example, the name will be email and the value will be

whatever the current value of that form component is.

And all we're going to do is call this starts at state like we did last time, this dot set state,

and we're going to have previous state so we don't lose anything that's in there already and we'll use

an arrow function and it's a bracket like that, then a curly bracket and we'll put dot, dot, dot,

pre state and then whatever the name is followed by value and that should be it.

OK.

OK, so we have handle change.

That's all set up now back up in our constructor.

Let's do the necessary binding.

This dot handle change will be equal to this door handle change dot bind this.

And we're also going to have of course have a form submit.

So let's get that started right now.

Handle submit.

Instead of submitting the form, we're going to call this function and it will again take the EVP and

for right now we'll just say evt

prevent default and we'll take care of this in a subsequent lecture.

And again, we need to bind that this door handle, submit to this door handle, submit, mind this,

OK?

And because we have a form with potentially errors, we're also going to want to has error function,

which will take key as its argument and return this state to errors or index of whatever the key is,

is not equal to minus one, just like we did in our edit movie component.

And finally we have a render function.

OK, so what do we want to have in our render function?

Well, I know it's probably going to be in a fragment, so I'll wrap the whole thing in a fragment just

to be safe.

And if I don't need that, I'll get rid of it later, OK?

We're definitely going to want a title to log in seems logical, we're probably at this point also going

to want to have our alert component.

So let's let our IID auto import that for us.

Alert right there.

And I'll just close the tag down here and put in my two default values.

The properties alert type is equal to this state of alert type and alert message is equal to this state

of alert dot message.

OK, so those are our alert component.

Now we can start the form proper, so I'll put an H.R. maybe before the alert right up here just to

give it some nice spacing and so close that tag.

Now let's create our fourth form and I'll give it a class name, just a little bit of margin at the

top or padding at the top two or three.

That should be enough.

And we also want to have the on submit, which will be this dot handle segment right there.

So that will prevent our form from it'll give our form something to do, but it will not allow the form

to submit and do a server round trip.

Now, I only want to input in this form email and password, so it would seem logical to use our pre-built

input components.

So we'll do that and we'll let our ID import the necessary code for US input and I'll just give it a

blank line and close the tag and our first one will be the email.

So title equals email and then the type would be the HTML five email type and the name will be email

again.

And we'll also bind that to a handle change function.

Handle change will be equal to this door handle change

and we need to give it a class name and we'll use that same logic we did last time.

So equals this dot hetzer and we're going to pass at the name of the field email

and we'll do our ternary used invalid.

If there's an error, otherwise, nothing.

And our next one is the error do and that will be equal to this has error and again, handed email the

name of the field and we want it to be read.

So text dash danger if there's an error, otherwise hide, it does not send the message.

Error MSG is equal to please enter a valid email address.

OK, now a copy of this entire thing and just modify it for the password.

Paysite in their format, everything.

And let's change this to password and this two password lowercase and copy that and pasted here and

paste it here and here and change the message to please enter a password.

OK, now after those inputs we'll put an H.R..

And I will just put a button button class name equals button, say, button, dash primary.

And call it log in.

OK, so there is our form, so we've got a logging component, there's no air showing up at all.

So let's go back to abduct Jess.

What's important?

Import log in from.

Components plug.

Down here, we need to root to that, so.

I'll put it right here and we'll just use route like this one exact path equals log in and we want to

route this, but we want to route it with the props.

And we also want to bind it with the handle JWT change function so we can lift state.

So the syntax is a little bit different than, for example, one genre above, but not that different

component equals.

And we're going to hand it the props and we're going to use the logging component and we're going to

include all the props and we're going to use that when we, for example, redirect the user from the

login screen somewhere else after they successfully authenticated.

And then we need to bind, handle a JWT change equal to this handle JWT change.

And then we close our tag and we close a Retek.

OK, so that should be enough to get us started.

So let's go see if we can see our form log in and there is our form.

So another that we have a form.

The next step is to actually go back to our login component right here.

And actually do the necessary logic here to validate the data, provide feedback if there's a problem,

and then make a request to our back end and try to authenticate and we'll get started on that in the

next election.

## 78 - Getting the JSON Web Token from the back end

So we have our login form and now we need to work on this bit of code in logging J.S. right here, Handelsman.

So all we're doing right now is saying nothing, just prevent the default action document to form and

nothing happens.

And the first thing we want to do, of course, is some validation.

So we'll do that exactly as we did at our movie component.

So first of all, we declare variable errors, equal an empty array, and then we'll check for the existence

of our two required fields, the only two fields in the form.

If this dot state or email is exactly equal to nothing, then we'll push onto the errors, array errors,

push email and we'll copy this and do exactly the same thing for password.

And this is very simple validation.

You probably would want to add some validation for the email being in the correct format, but we're

not worried about that right now.

You can figure that out with two minutes on Google.

So password and password.

Now, if either of those fail than our array length will be greater than zero.

So if errors start length is greater than zero, then we'll say return false.

And of course, we also want to set the stage this set state errors, errors.

So now we have something that just to display.

So if the form doesn't submit properly, if they haven't filled it out properly, we should get some

visual feedback.

Let's see if that works.

So we'll go back to our Web browser and we'll try submitting an empty form.

And that's exactly what we want.

OK, so assuming we are past this point now, we actually connect to our respec end and try to get our

JWT by authenticated.

So we'll do exactly the same way as we did in our edit movie component would declare a constant called

data and that will be equal to a new form data.

And we'll just grab all the fields from the Ford Target

and our payload const payload is equal to the object from entries and we're going to pass it data entries

for all of the individual form elements and there should only be two now.

Never declare our request options.

So constant request options is equal to and it's an object.

And we're going to go with method of post because that's what we're listening for.

In our back end is a post request and the body will be chasing that string apart and its payload to

converted into Jason and I.

We do the actual request.

So we're going to fetch, as usual from HDB colon slash localhost PT. four thousand slash V one slash

sign in and we handed our request options.

Then we take our response and convert it into response to Jason.

Then we take our data and use an arrow function and we do our check, if data does error, something

went wrong or else things are good.

So if there's an error at this point, we want to set an alert.

So we'll say this starts at state and we want to set the alert to be of type alert the danger

with a message data, the error message.

Otherwise we successfully authenticated.

So at this point, all I'm going to do is just say console log data so we can see what we get back.

All right, let's try this.

Let's go back to our Web browser and let's open our JavaScript console to see if there's any errors

as we get any feedback.

And I'll just click on log in to make sure everything is current.

And it is.

So let's put the wrong information there at X.com.

And we have a problem right away.

This DOT target is undefined on handle change.

So log in just line twenty five.

Let's see where that problem is.

Log-in Geass, line twenty five, it should be evt target value that explains the OK to many.

This is in JavaScript, so let's go back and try it again.

We'll clear the console and we'll type X X.com with a password of X.

No errors this time.

That looks good and we have a bad request, which we're going to fix later on because this is technically

not a bad request.

But we did get an error from our post and we got our unauthorized, which is perfect.

Let's clear the console and let's try the correct information.

So me at here dot com is my valid user on my back end and the password is password log in and we got

our response and there is our JWT right there in the console.

Perfect.

So what do we want to do once someone is actually logged in?

Well, we'll take care of that in the next lecture.

But it would seem to me that the logical thing to happen here would be to actually set the state, of

course, and we'll take care of that and then to redirect the user to some other page.

So we'll get started on that in the next election.

## 79 - Handling a successful login

So we're looking for what seems to be working, we can authenticate, we can get Jason Webb token back

from the server and right now, of course, all we're doing is logging that to the console.

And that's not terribly helpful.

So what we want to do right away, of course, is to lift state, to take the response we're getting

from the server and put it back in the state on the parent component.

And to do that, of course, we need a function just like we did when we did our lifting state tutorial.

So we're going to add a handler here or a function which I'll call handle JWT, change the same syntax

we use last time around.

That takes one argument JWT and we're going to just call the handler that we received as a property

in the route to this particular class.

This props dart handle, JWT change right there and we'll handle the JWT.

And up here after we do our console log data, we'll call this DOT handle JWT change and we'll hand

it right now.

We'll just hand it data.

OK, and we can check to see if that works.

We can do that by going over to abcess and just giving a visual representation

of the state so we can go pre.

And then Jason Stringer for this state

No.

Three, just so we did a long time ago.

And that will give us a view of the state on our application.

So let's go have a look at it.

And you can see right now I have the login screen and there's my state, JWT is equal to nothing.

So let's log in here dot com password and see what happens.

So there is our duty.

But now that's not exactly in the format that I want it, but it's nice to see that it went from this

component up to the parent component.

And from there I can push it to any other component I want.

So let's just change this.

We don't want that response.

We just want JWT to be equal to a string.

So let's go back and make a slight change to our code.

So back in log in jazz, instead of saying this DOT handle, JWT change, we actually want to grab the

just the value of the actual token.

So there's lots of ways to handle this.

But the easiest way is just to call our standard object object values plural and hand it data and we

get that result back as an array.

So we want the first index of that array.

Let's try that.

Let's go back to work our Web browser and I'll just reload everything.

OK, so we have no value in JWT state.

So let's Lorean we at here dot com password and log in and that looks much better.

There is my token and now I can do anything I want with it.

So the last thing I'll do is redirect and we'll just push the user back to say the admin route.

This dot props, dot history, dot push.

And we want to go to path name for lowercase.

Let's try that and see if that works.

So reload.

Nothing in the JWT in the state.

So we adhere to our common password and we are redirected and we have our token in state.

OK, so the next step is to actually protect those roots.

Right now it appears to be protected.

But if you know the you are all you need to go to, you could in fact get to the screens that should

be protected.

And we need to protect those on the back end and we'll take care of that in the next election.

## 80 - Adding middleware to check for a valid token

So we've made some good progress, we can log in now and we're getting a JWT token, which isn't in

quite the correct format, but we'll get to that momentarily.

There was an oversight on my part and we can log out and everything seems to be working properly.

And this the managed catalogue page should be a protected group.

But really, if I log out and just put it in there, I still see it.

So we still have some work to do.

And part of what we have to do is to actually do something with that token.

So let's get started.

I'm going to go to my back end source code and the first thing I'm going to do is fix a mistake in the

tokens.

Don't go file.

And I'm just a little oversight on my part.

The very last line at the end of the function, I have approached Jason and I neglected to cast this,

which is a slice of bytes, JWT bytes.

It needs to be a string.

So I need to cast that to a string.

So I'm going to stop my application, start my application and go back to my Web browser and log in.

And it should still work.

But now I should have a valid token, something in the correct format.

So I go to the login screen and log.

It has me at here dot com with the password password and that looks better.

That's actually the correct format.

OK, so now we need to actually do something with a token.

So if you recall when we log in, we actually get this token and we lift state and put it at the top

level of our application, the app component, and there I can push it anywhere I need.

So what are we going to do with that?

Well, part of what we have to do, of course, is to add that token to our request that we send to

our protected roots on the back end and authenticate using that token.

So in order to make that happen, the authenticate part, we need to go back to our back in code and

open our middleware and we need to write some new middleware.

OK, so I'm going to call this middleware.

It'll be a function with the correct receiver and I'm just going to call it a check token.

And because it's a middleware, we have to have next SCDP dot handler

and it needs to return.

And a sheep dog handler.

And inside of that, just like we do in the one above, I'll just copy this entire section and modify

it to save some typing.

We have a return function and this part is going to go away.

And what we're going to do here is, first of all, add a header checking for authorization.

So W dot header ad and we're adding a header called very or a varying header and it's going to be authorization.

So the key is very and the value is authorization.

And then we'll grab the header that we're getting called authorization and we'll store that in a variable

called off header.

And that is our header doget.

And we're going to get authorization, a header that we're going to have to send from our front end.

And then we'll just do a check here if our header is equal exactly to nothing at this point, if you

wanted to, you could set an anonymous user.

So I'll just put could set an anonymous user and we're not going to do that.

But there may be situations where you want certain protected routes to be available to anonymous users.

And if so, this is where you would perform that logic.

Now we're going to take that off header variable and split it on spaces.

OK, so we'll declare a variable called header parts and that will be equal to from the strings package

we're going to call split and we're splitting off header

or splitting it on a space which will be in the the value of the header that we set in the front end.

And we'll get to that in a while.

And we're going to expect there is exactly two parts in that split in header parts.

So we'll say if Len header parts is not equal to two, then we'll just throw an error after error JSON

and we'll pass the response rate and a new error error or start new and we'll just say invalid auth

header and return.

So that's the first thing we're checking.

Next, we'll see if the first part of header part is not equal to what we expect it to be.

We'll throw another error.

So I'll say if header parts and the first index of the array from that split is not equal to error,

which is what we're going to set in that header in our front end, then we'll throw that again.

So I'll just copy this and change the error message.

To.

Unauthorized, no bearer, and you can make that whatever you want.

OK, next we'll look for the token itself, which was included in the header that we're going to receive.

So we'll declare a variable called token and that will be assigned the value of header parts.

And it's the second index in the array.

Now we begin checking that token and we have a lot of checks to make.

So, first of all, will declare your variable claims and potentially an error and those will be assigned

the value of from the JWT package.

We're going to call each Mac check and that expects a couple of parameters.

First one is a slice of bytes.

So we'll convert our token to a slice of bytes and bytes is not plural.

So let's get rid of Yasser.

And the second thing we want is our secret, and we get that again, we want to convert it to a slice

of bytes and it's going to come from app dot config, dot, JWT dot secret.

OK, so this is our Mac check.

And again, we're going to check for an error.

So we'll say if error is not equal to nil, then I want to return some errors here.

So I'm going to copy this and paste it in here and change that to unauthorized

failed.

And I'll just call it a Mac check.

You can make whatever error message you want there.

So that's the first thing we're going to check against the token.

The next thing is you want to make sure is the token still valid?

So we'll say if not, claims dot valid time dot now.

So is this token still valid at this moment in time?

And if it's not, then we want to return an error.

So in here again, I'll copy this error code and paste it in there and change the message to an expired.

And next, we want to make sure that whoever issued this is, in fact our application and we're not

getting a token from some other application, so we'll say, if not, claims DOT, except audience.

And the audience we're looking for in my case is my domain dot com.

And you might be using a different one, but just put whatever is appropriate for that.

And again, I'll paste in that area code and say invalid audience and we want to check the issuer.

If claims dot issuer is not equal to my domain dot com, then again, another error.

Invalid, sure.

Now, if we pass all of these checks and again, let's go through them again, so we're getting our

token and we're getting that from heter parts, which came from what was passed in the authorization

token.

And it's the second index in the string.

First of all, we make sure that the claim is valid by doing, in fact, check.

So if it's not passing that check and there's something wrong with this token and we're just going to

return an error, then we make sure that it's still valid that the token hasn't expired.

And we do that here from line forty nine to fifty two.

Next, we want to make sure that the audience for what we're going to send is accepted, that we're

actually allowed to are permitted to send this kind of information, whatever is behind this token,

whatever the token is protecting to the appropriate audience.

And we make sure that the token was in fact issued for our domain.

And if all of those things work, then we can actually do what is necessary.

So what I'm going to do is extract the user ID that was embedded into that token user ID and potentially

there are assigned the value of and we want to convert it from what it is into an end.

So Sterkel of that person isn't.

Claims subject, and we want it to be a 64 bit one,

so we have that and we check for an error if error is not equal to nil l paste in my area code and adjusted

and just say unauthorized.

I don't really know what to put here for an error message.

And this will suffice for the moment.

And just to make sure we have this, I'm going to put this to the screen, to the console logged print

line, valid user and then user, OK, otherwise I won't be able to compile this because I have to do

something with this code.

So that is our middleware.

And let's just run our application to make sure everything compiles properly.

And I have an error on middleware.

Don't go line twenty two right there.

We should have a comma or a closing parentheses like that and get rid of this one.

And get rid of the extra HDTV and I know where that came from.

That looks better.

Now we have a different air down here.

This should be a parentheses, not a square bracket.

Among my first coffee, please forgive me.

All right.

Now, we should be able to compile this and run it.

Good.

It's running.

OK, so we've got the middleware and now we actually need to use this middleware.

So we go back to our roots file.

This is where we use middleware.

Right now we have one piece of middleware enabled course and we're going to have to make a change to

that cause middleware in a moment so we could actually pass the necessary information.

But first of all, let's set this up so we can actually use more than one piece of middleware because

we want to use the course middleware on everything, but we want to use the check token middleware only

on some routes.

So how are we going to do that?

Well, there's a number of ways of doing it.

The simplest way far and away, in my opinion, is to install another package that allows us to change

middleware.

And it's a very popular package used in a lot of places.

So open your terminal window and let's go get it.

And we're going to type go get.

And the package we want is from GitHub dot com and it's from the user j u s t i and as Ellis, that's

what I want.

And that allows me to change the middleware.

And you can go to that you url GitHub dot com just in Alice and read all the documentation.

It's very short because it's a very simple package.

So that's installed, it's now available to us and I want to use it.

And in order to do that, in order to change things the way that we want, we actually have to write

another helper function here because we're using HTTP Rohter.

So it's a very simple thing.

I'm going to write a function.

It has the receiver app pointer to application and I'm going to call it wrap.

It allows us to wrap things, wrap our middleware and it's going to take one parameter which I'll call

the next door handler and it returns and htp router.

Dot handle.

And inside of that, we have a return, a function, we're going to return a function which takes the

response writer and it takes a pointer to a request or a cheap dot request and it takes place, which

those are the parameters, the parameters from HDB reader, this GDP Rohter programs.

Now, the first thing we're going to do is declare a variable called context.

We'll call it.

And it's a sign, the value of context with value.

And it takes as its arguments our context, the context from our request.

And we're looking for programs and we pass it.

Now we have the necessary X, we've added the things from our context into what we're returning.

Otherwise we won't be able to use any parameters in our HTP Roter package.

And then we just say next door to serve GDB handed our response writer and handed her with context and

give it the variable we just created.

So we're just adding what's necessary from the context back to our new htp rotor door handle.

OK, so that's our wrapper function and that means we now have a means of using more than one bit of

middleware in our roots.

So to do that, we could undo our roots function and we declare a new variable.

I shall do it after the HDB revenue we say secure.

That's what I'm going to call my chain is equal to the package.

We just installed new and we handed a check token.

Now I can put as many pieces of middleware into this variable secure because this is actually a chain

of middleware.

So to use that, I'm going to come down and I'm going to protect one route and the one route I'm going

to protect right now is this one.

And to do it, I could fight with this this syntax or I could use a different syntax that's available

to us in the HTTP Rutter package, which does exactly the same thing, but makes using this particular

secure chain much, much simpler.

So I'm going to comment that one out and I'll create a new route right here and it's going to be router

again.

But this time I'll call post like this the same functionality, just a different syntax.

And the route I want to protect that I want to match here is v1 v one slash admin edit moving.

What happens when someone fills out that form to add or modify a movie.

And here this is where I'm going to use the wrap function.

We just created Aiport Wrap right there again wrap.

And I want to wrap secure the middleware chain that I just declared above, then func, then I'm going

to call a function and the function I want to call is the same one I did before after edit movie.

OK, now with this change to our roots file, now I can add this little bit Abigroup secure to anything.

I want to protect the back end to require authorization in the form of a Jason Webb token.

So that part's done.

Now let's go back to our middleware.

We want to make one more change in our middleware at the very top here.

We need to modify our headers a little bit so we actually allow certain kinds of things to be passed.

And one of the things that we're definitely going to want is an authorization header.

And we may as well at the same time set our content type header to be something that's permitted.

So this is as simple as adding one line right here w the header set and I'm going to set the the key

to be access Dasch, control Dasch, allow Dasch headers exactly like that upper and lower case matching

exactly what I've just typed and what we're going to accept our content dash type than a comma with

no space and authorization.

I make sure you spell that right or it's not going to work for you and that's it.

So now when I start my application, it's already stopped.

Let's run it and make sure that everything compiles.

And it does so now we have middleware, we have a means of applying that middleware selectively to certain

routes in our application, and we're setting the proper headers that allow us to accept an authorization

header and a content type header, which we didn't have.

And everything still worked.

But it's proper practice and a good idea to actually accept and send the correct content type header.

So in the next lecture, we'll start working on the front end and actually see if this authorization

token we received does its job.

## 81 - Protecting the route on our front end

So this time around, we want to actually send the appropriate header from our front end to our back

end so that we can access the protected routes.

And if you recall, the one route we have protected is the one that happens when I submit this form.

So right now I'm logged in and it doesn't really matter.

I mean, if I try to submit this right now, open the console, that route is protected.

The one that happens when I try to save this.

So I'll put, say, an X after the text here and try to save this.

I actually get an error invalid author and that's sent back by our back end.

And then after I click, OK, we should get the same error up here, so we'll get rid of that alert.

We don't need that anymore.

But what we want is to send the appropriate web token, which we can see right here.

We have it at this level of our application.

And of course, this is at that level, but we don't have the token at this component.

So we're going to have to, first of all, push the information containing that token to the add edit

movie component and then add the necessary header.

And fortunately, that's pretty straightforward.

So let's go to our code.

And the first thing we'll have to do is push the JWT to the edit movie component, and that takes place

right here on line ninety eight.

So we have to modify this.

So instead of going component equals and just edit movie, we need to pass the props because we can

pass JWT as a prop..

So we'll add props here.

And then push it using the Arrow function and open up parentheses, and we'll make that movie a component

this way and include all of the properties that exist like this, dot props and then add another property.

JWT equals this state DOT JWT, which is where that's currently store and close that tag.

So we don't need this anymore.

And finally, close this and close the overall tax.

OK, so that will actually push it to the movie component.

We now have access to JWT as a property in our edit movie component.

So let's go to our edit movie component and down here, let's just make sure we actually got that.

So in my component did mount well, actually, just add this console, the log JWT.

In a movie component did mount

and I'll put this stuff props WWT just to make sure we actually have it.

OK, so let's go back and clear our console and we'll go to the homepage and we'll go back to our will

log in so we can see the list of movies up here, dot com password and log in and we'll go to The Princess

Bride and watch our console.

And there it is.

We have our token in our component did Mount Perfect.

So let's go back to our code.

OK, now we need to add the necessary header on the part where we submit the form right here, handle

submit.

So what I'll do is after I declare my data and payload constants, I'll create another one called const

my headers and call it whatever I want, but my headers works and that will be equal to new headers

straight from the standard JavaScript.

And I will append to headers to this my headers append and the first one will be our content type,

which we now allow in our cores middleware.

So the key is content dash type with a capital C and a capital T and the value is application JSON how

it works without that.

But this is good practice.

So I'm going to add that now we'll duplicate that line and add the second one as being authorization.

And this is the one that will hold our token authorization and make sure you spell it right or this

isn't going to work.

And this doesn't consist of just one thing.

Remember, we're splitting the content of this on a space.

So I'm going to add, first of all, what we expect to see there, beer and then a space, and then

I will attend to that.

Plus, this props JWT.

OK, so now I have my headers variable, which is called my headers.

And down here in request options, we're going to get rid of this comment, that line, which I don't

know if you actually saw me add, but it's gone now and we'll append our headers by saying headers,

my headers.

OK, so that actually includes the necessary information from our Jason Web token in the post to the

edit movie route.

OK, so we should be able now to actually edit this movie.

So let's go back and I'll reload everything to make sure it's current, OK, and then I'll have to log

in.

So I'll go back to my login screen and I will log in as a valid user.

Me out here.

Dot com password.

Password.

OK, I his type that password.

I'll clear the console log in and I will go to The Shawshank Redemption and I will add one word after

this one word.

There is actually two words but you get the idea.

So this should work if we've done everything correctly and we have an error.

Authorisation is spelled wrong after I told you to spell it right, I went and spelled it wrong myself.

So let's go fix that

authorization now.

It's better.

OK, so it's reloaded it.

I'm still logged in.

I'll clear the console and I should be able to add one word again, one word safe.

And it appears to have worked.

So if we go back and look at The Shawshank Redemption, there it is.

And I should be able to remove this and save it.

And it's gone.

Perfect.

So now we have a means of actually pushing our JWT to the components that need it as a prop..

And we have a means of actually successfully authenticating for all of the protected routes on our back

end.

And we're not done yet.

So you can see that if I, for example, look at The Shawshank Redemption, I'm logged in right now.

But if I copy this, you URL and then log out and then paste it back in.

I can still see that screen and I shouldn't ever see this form if I'm not logged in so clearly, we

need to make some kind of redirect on our front end to check to see if the variable or the value JWT

exists in the state.

And if it doesn't and we can do this all right.

In the app component, then redirect them somewhere else, perhaps to the login screen.

That would seem appropriate and we'll take care of that in the next election.

## 82 - Adding redirects for protected components

So this time around, we're trying to solve a very simple problem, and it goes like this, if I log

into the site and I have both my back end and front end running and so should you.

So log in.

Is Miette here dot com with the password.

Password.

And a look at American Psycho and then I'll copy the hero, OK, and then I'll go and then I'll paste

that you URL back in my location, burned my Web browser and return and I can see that screen and I

should never be able to see this form if I'm not logged in.

And that's a really, really simple fix.

So let's go back to our code and we'll take care of that for this page.

So we're looking at end of movie.

Now, if you search the Internet, go to Google and look for ways to do something before the page loads,

you'll see references to a function available to you and react native called component will mount.

And actually that's now deprecated and it's considered unsafe.

And that's not what we're going to do.

You're never going to use a deprecated function because it's just a bad idea.

So let's get rid of this console log instead.

We'll put the logic right in component dismount and we'll do it before we do anything else.

So here, right after the opening declaration of this function, I'll just check to see if JWT, the

JWT has a value.

If props start this DOT prop, JWT is exactly equal to an empty string.

And I just do a redirect that's as simple as that.

And the redirect is exactly like we do elsewhere in our code.

This props dot history, which I have to spell right the push, and I want to push to pathing all lowercase

and I'm going to take them to the login screen and it's as simple as that.

And then just because I'm paranoid, I'll return and it should work without that.

But the return will ensure that nothing else in this function happens.

So let's go back to our code and I'll go to the home screen and then I'll paste that you are old and

again and I'm not logged in, so it should redirect me.

It does.

It takes me to the login screen and that's really, really straightforward.

So obviously we're going to put this kind of code in every page that we want to be protected in our

application.

And there's some other things we have to take care of.

Let me go back here and get rid of that alert.

For example, on the on submit, if there's an error, we have this alert data error message.

I want to get rid of that because that's just extra code that serves no function because we're already

displaying a custom alert at the top of the page.

But right now, what we're going to do in the next lesson is I'll give you a challenge where you can

do the sorts of thing we have done for the edit movie page to the other parts of our application that

should be protected.

So that's enough for this time around and we'll take care of going over the challenge in the next election.

## 83 - Challenge

So this time around, I'm going to give you a bit of a challenge.

Nothing terribly difficult, but it will take you a bit of effort and take you a bit of time.

So I'm logged into our application right now and I'm looking at the add edit movie screen.

And as you know, this form, the front end is protected with a redirect.

So if this value JWT in the state is set to an empty string, you can't even see this form.

You get redirected to the login screen in the same way at the back end that actually does the update

or ad of the movie in the database.

That call to arrest API is protected with middleware on the back end and the one that populates this

form is not protected.

And it doesn't have to be because we're just getting the information as adjacent file.

And in fact, that call is exactly the same one that we see when we go to the public movies page and

click on show me this movie.

The same call to the rest back in is made on both of those screens.

But this screen managed catalog, this needs to be protected on the front end.

Nobody should see this list of films unless they're logged in.

So that's something that will take you 15 or 20 seconds.

And that one's really straightforward in the same way, although this screen is protected and the save

button, the call to the back end is also protected.

This one is not the delete button.

So you're going to have to protect that in some way, shape or form.

Now, I'll leave that as an exercise for you.

Try to figure this out.

Don't just skip ahead and watch the solution in the next lecture.

Do it yourself, because that's how you're going to learn to do this sort of thing.

But in the next lecture, argue over how I accomplish this.

Good luck.

## 84 - Solution to Challenge

So how did you make it with the challenge?

It wasn't terribly difficult, I hope, and if you did have a problem, don't worry, I'm going to go

through my solution right now.

So the first thing I'm going to do or the first thing I'm going to talk about is protecting that list

of movies that you click on to edit, in other words, the route admin.

So in my act, Jess, the first thing I had to do was ensure that I could push the JWT to that component

as a prop..

So the old route is right here, commented out.

And the new route is just a variation of what we did for this one up here.

I simply added the existing properties and then manually specified this state, JWT, to be equal to

the property JWT, which is then available to us in the admin component.

So that pushes the JWT to that component and then in admin mindgames on component did mount right here.

I put exactly the same logic I did in the other component.

So if we're not logged in, then the property JWT would be equal to an empty string.

So I just redirect them to the login screen.

So that one was really straightforward.

Now the next one was at a movie and we had to make changes both here on the front end and also here

on the back end.

So on the back end, all I did was change the route and the old one is here, commented out on line

39.

And the new one is a variation of the same thing I did for the edit movie route online 35.

But in this case, the method is get I'm going to a different path through to matching and I'm calling

the Dileep movie Handlock.

So on the front end back here, I had to add the necessary headers.

So in my confirmed delete function in the confirm alert, I specify my two headers, so I create a constant

my headers of type.

New headers gave it the content type of application.

Jason and then I just put the authorization header exactly the same as we did the last time around.

And then I specified here where I put method get I put headers, my headers and that's all I had to

do.

So the source for this is available for download on this lecture.

If you have any problems, both the front end and the back end, and it's time to move on.

## 85 - Saving our token when the user leaves the site

So things are starting to look pretty good, we can log in, we can log out, we can protect routes,

we can update and add and delete movies.

And that's all very good.

But there is one problem, something you've probably noticed, and let's have a look at it.

So I'm on my application.

Let me log in so I'll log in as we at here dot com with the password password.

And that's fine.

I'm logged in.

I can see my log out.

I've got the protected menu items and there's my state with JWT and I can browse around and do things

and that's great.

But let's go back to the homepage now.

Copy this URL.

OK, so I copied it.

Now I'm going to go somewhere else.

So I go to w w w dot go langue.

OK, so now I've gone to a different site.

Now if I actually try to go back to where I was a minute ago by pasting that you URL in, I'm not logged

in anymore and that's kind of irritating.

And these errors over here come from my ad blocker so we can just ignore those.

There's nothing to do with our application.

So what we want to do is to be able to save the fact that we're logged in.

And you might do that by allowing the user to check a box on the login screen and saying, remember

me or something like that.

But we're just going to assume that users want to stay logged in right now.

How are we going to do that?

Well, let's go back to our code.

And I'm looking right now at the login file, OK, the logging file.

And I'm looking at the part where we actually perform the log in the handle submit function.

What I'm going to do is at this point, after my fetch in the final then clause in the ELT's where we're

allowed to log in right here.

OK, so we're we're redirecting the user to admin.

Just before we do that, I'm going to do something else I want I'm going to do is take that token that

we've received back from the rest API and I'm going to put it into local storage and going to put it

into a cookie if I wanted to.

But I'm going to put it into local storage just because it's really convenient and every modern browser

out there supports it.

So after this DOT handle, JWT change and before the redirect, what I'll do is just call window local

storage dot set item and I'll call it JWT.

That's the key.

Will get used to get it back.

And I need to of course, string of JSON.

I need to call Jason Stringer on this item before I can.

Stored in local storage.

So Jason, strength five and I'm going to string a five object dot values data, the first index of

the array.

At that point I've put it into local storage, so let's see if that works.

Just to go back to our Web browser and I'll click on Log in and I'll clear the console and enter me

at here.

Dot com password.

And then over here in my JavaScript developer tools, I'll choose storage and I'll look at local storage

and click on Local Who's three thousand and there it is.

I've saved that item in my local storage.

Well, that's great.

I've got it in there.

But of course, I need to get it back out in order for this to be useful at all.

So let's go back to our code and this time we'll go to Abcess.

And there's a couple of things I'm going to do here right at the top to start with on logout rather

than just saying this.

Dot sets are this dot set state and setting JWT to nothing.

We also want to clear local storage at this point.

If they log out, if they're making the conscious decision to log out, we want to get rid of that item.

So we'll just call window local storage, remove item and remove it by key.

And I called it JWT, so that should get rid of the item once the user chooses to log out.

No.

More importantly, I'm going to add another function here and I'll add a component did mount function.

And at this point when we're loading this component, I'm going to check to see if there's an entry

JWT in local storage.

So I'll say let T for token equal window local storage douget item JWT.

And if that doesn't exist, T will have the value of no.

So I can just check if t so if it's not know it exists then I'll check to see if the user is already

logged in.

If this dot state dot JWT is exactly equal to an empty string then I'll try to log the user in again.

OK, so that would be as simple as this dot set state and I want to set JWT and since I've just pulled

this item out of local storage, I have to person Jason pass t.

OK, now that's very simple.

Let's see if it works, so I'll go back to my Web browser and I'll clear this item out of local storage.

So delete all repurchases, reload this, and go to the console and see if we got it right.

So let me clear this whole thing and let's look at me at here.

Dot com password.

OK, so we're logged in.

That's great.

Now, let me copy this URL and go somewhere else.

W w w dot say reactivates dog.

I think that's a UFO.

Yes, it is.

And now I'll paste my address back into the location bar and we're logged in and let's make sure everything

still works.

And it does.

Now let's log out and let's check local storage.

And there's nothing in there so perfect, so now we have a means of persisting our log.

Now, obviously this will only last for the lifetime of the token, which in our case is twenty four

hours.

And there are lots of ways to handle the situation where you try to use a token and redirect the user

to the login screen.

But none of that is terribly difficult, and I think you should be able to figure that out on your own.

All right.

That's enough for this time around.

Let's move on.

## 86 - Making better error responses from our back end

So before we move on to the next topic, there's just a little bit of cleanup I want to take care of,

and it has to do with the fact that every time we have an error on our back end, we're returning an

HDB status bad request, which is appropriate for some cases, but not for all of them.

And I just want to clean that up.

So I'm looking right now at the list of HTTP status codes on Wikipedia.

And as you can see, there's quite a few and most of the ones we're interested in are either in the

300 or the 400 range.

So we're sending right now bad request a lot.

And there are other ones we probably should be using instead.

So let's go back to our back end code and open up utilities.

Don't go and look at the error adjacent function right here on line twenty five in my code right now

that only accept two parameters.

So I could make this require three parameters, but that would mean I'd have to go and fix every single

call to error Jason in my code base and.

Well, that's not a terribly onerous task.

I think I'm going to take a simpler approach.

One I'm going to do is add a third parameter here called status, but I'm going to make that very ADIC

it's going to be an aunt one or more.

Try that again.

Zero or more ints, which means that this is actually not a required field.

And what I'll do inside the function body, the very first thing I'll do is declare a variable status

code, which is an end, and that's going to be assigned the value of HTP status.

That request.

And that's our default.

That's what we're sending right here on line thirty five.

But instead of sending that specific code, I'll send status code.

And then after I, after I sign that variable or create that variable status code, I'll just check

to see if the length of status, which is my very ADIC option up in the parameters for this function

if that's greater than zero and I have one.

So all I have to do at that point is say status code is equal to status index zero.

And now I have the option of passing a status code when I call this function.

So let's go over to middleware ago and look at where we're calling this and we're looking for every

time we call origination.

So right here in the first case, we're calling if the length of header parts is not equal to two,

then it has an invalid authorization header and that's sending nothing as a third parameter as nothing

is right now.

So it's going to call bad request.

So if I go back to Wikipedia and look at the status, that request, which is right here, four hundred

bed request, the server cannot or will not process the request due to an apparent client error.

For example, malformed request syntax size two large invalid requests, message framing or deceptive

request routing.

I think actually bad request is OK because we're not getting the appropriate content from our user.

So the next one, if the header parts first index is not equal to better, again, I think er four hundred

status bad request is fine for that one, but the next one where we have our claims and we're actually

trying to verify the signature of that claim here, we should be passing a different one and this one

should be htp dot status forbidden.

And if you go look at the Wikipedia page and there's a link to that in the course resources for this

lecture, you'll see that's the appropriate one for this particular case.

And I'm going to copy that because that is actually appropriate for this one where token is expired,

for this one where we've been invalid audience and for this one where we have an invalid issue, where

we're able to pass the content we're getting from our request, the header that's included in the request,

but it's not appropriate for our use case.

And the last place I'll put this is right here, status unauthorised again, it's status forbidden.

And that's simply because we can't get a valid user ID out of the request token and we should be able

to find one.

So there's an example of how we can modify our error JSON files that we're sending back to the user

and specify a different status when required or when appropriate.

And I'll leave the rest of them.

And there aren't very many as an exercise for you to determine what kind of response code you want to

send back when there's an error.

## 87 - Adding images

So one thing we've not done it all is add an image to our site and it's not as straightforward as you

might think it is, at least it's not the same as you would do in straight HTML, but it's really pretty

straightforward.

So what I'm going to do is just add an image here on the homepage and we'll do it two ways.

One, using the access files for this particular component and it doesn't exist yet.

So we'll create it.

And the other way is just to declare an image in the same way you would with HTML.

And the syntax is a little bit different, but it's very straightforward.

So let's go look at our code.

And I'm looking at the component for home, which is home dojos inside the components folder, inside

our source folder.

And you might also notice that I have a new folder at the same level as components, in other words,

right beside Apgar's, and it's called Images.

And inside of that, I have one graphic, a file called movie underscored tickets JPEG.

And that file is on the course resources for this lecture.

So go grab that file, create the folder inside your source folder called Images and put that image

inside of it.

And once you have that done, come on back here and here's how you include an image.

So what I'll do is wrap this entire return statement, the contents of that return statement in a div

and I'll give it a class name just to center everything equal to bootstraps text center.

OK, so I'm going to wrap this whole thing in a div just so I can have more than one thing in there.

And after the to the title, I'll put a horizontal rule and then I'll put my image and I'll write the

image tag first, then I'll show you how to set the source.

So image source equals and I'll leave that blank for right now.

And then the alt text old equals movie tickets or movie ticket like that.

And I'll make this a self-closing tag.

OK, now I need to put something in the sauce equals part, and I can't just put the direct path name

to that image.

Instead I import the image so I'll import and I'll call it ticket from.

And the location would be the current directory slash up one level because we're inside the components

folder.

Slash images, slash movie tickets, dot GPG.

OK.

And down here, I just use my standard syntax to use that import ticket.

So once that's done, if I go back to my Web browser, I should see those tickets and there they are.

OK, now that's one way of doing it.

Now, let's also create a file for our home page.

So into my components folder or create a new file and I would call it home dot success.

And then, of course, back in home mortgages.

I need to import that.

So I'll import.

And this time it's just dot slash home dot com.

Now let's go write some success in there.

And what I'll do is create a class called tickets and I'll give it a background image.

And this is the part we're interested in the URL for.

This is going to be dot slash up one level and then inside the images folder and then movie dash or

underscore tickets.

Dot jpeg.

And I'll give it a width of, say, two hundred pixels and a height of, say, one hundred and forty

pixels and margin left on make auto and margin right on my Court of Justice center it within the deth,

then back in Homburg.

Just I'll simply put a horizontal rule so I can see what's going on and then I'll have a div class name

equals.

Ticket.

And just leave it empty now, when I go back to my Web browser, she'd have to reload this, don't see.

So I missed something.

The s on tickets, I think that's what I called it.

Tickets, plural.

I will explain that.

Let's go back and look again.

And there it is.

The flight's not quite right.

And because I'm a perfectionist, I'll just add a couple of pixels to this one hundred and forty three

and see if we can get the whole image in there.

Not close enough for government work.

In any case, that's how you add an image to your various components in react.

Now, to be clear, this is only for images that exist within your REACT application.

If the image exists on some remote server, you can just include it the way that you normally would

with the source equal to the full you URL for the image.

But if it's going to be existing within the source for your reactor application, this is one way of

getting it in.

All right, that's it.

Let's move on.
