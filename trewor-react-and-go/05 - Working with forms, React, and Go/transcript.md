# 05 - Working with forms, React, and Go

## 59 - Building a form in React

So this time around, we want to create a form that allows us to add or edit a movie, and I'm looking

at my code right now.

And if we look at our application, here's what I want to do.

I'm going to add a new link here between genres and manage a catalog that allows us to add a new movie.

So all that has to do is to display a form.

And we need to, first of all, create a component and a route to it.

So let's get started.

So back in my idea, I will go to Apter Jazz and I'll create a new link right here.

So I'll just copy this one and modify it.

I'll call this ad movie

and it will link to slash admin at and down here.

I'll simply create a new path right here or a new route right here.

And the path will be equal to slash admin slash ad and the component will be one that doesn't exist

yet.

And I'll call it edit movie.

So let's go create that path or that component sumai components.

All right, click and choose new file and I'll call that edit movie Jass and I'll use my shortcut.

I am Arcy and I'll export a default class edit movie which extends component.

OK, so give us some state and in there I'll put a movie which will eventually be empty and I'll put

new loaded which would be false and error, which should be no

and then we'll have a render function.

So the first thing we'll do is extract our movie variable, even though initially when we're adding

a movie the form will be empty.

We could reuse this component to edit a movie later on.

So let's extract movie.

And we'll extract that from the state and we need to return something, so let's put it in a fragment

or wrap everything in a fragment and let's give it a title and we'll just call this and edit the movie

and maybe put an HRR under it.

And I will create our form and we'll give it a method, even though we're not using this form.

We're not going to submit it the normal way.

We'll submit it using JavaScript and just put something in here.

So we'll put in a div class name equals maybe three.

Just give us margin and a label for equals title

and we'll give it.

The class name is equal to form a label, which is bootstraps class and the title will simply be titled.

Then we'll have an input.

Type equals text, class, name equals form, dash control.

So it uses bootstrap styling.

And we'll give it an I.D. title, and the name title will also give it an initial value, which will

be empty when we add a movie, but it won't be later on.

OK, so now if we go back to our abduh, just that error should go away once we import it.

And there the error is gone now.

So let's go have a look at our code and see if it works.

So at a movie and it gives us our form control.

Perfect.

So let's go to the rest of our form

back and edit movie.

We can just copy this section.

And pasted in there, and the second thing will be the release date.

We'll call this release day

and copied and pasted here and pasted here and paste it here, and that should give us another entry

on our form.

Let's go make sure that it is not really I want this to be bolded.

I don't like the non bold text there.

So let's just quickly come in here to components and create a new file called Edit MoVida Success and

in that will have label which are going to spell right.

And we'll give that a font weight of old and go back to edit movie and import that success.

That should fix that, and it does that looks much better.

OK, so we have really state next comes the runtime, I believe will copy this

and paste it.

And this will be for run time.

Copy paste, paste, paste.

And name this runtime.

OK, and next, we'll have the empty AAA rating, so this will be another day of class name will be

equal to maybe three segments before and now will have a label four equals MPE A rating and the class

name will be equal to form the label.

And we'll call this simply a rating and we'll make this one to select because we know there's only a

predetermined set of values you can use.

Select class time equals form Dasch select and we'll give it a value as well, which will allow us to

have this set to the correct value when we're editing a movie.

MoVida MP a rating.

And in here, we'll make the first one.

Option.

Class name equals form dush select, and we'll just say choose and we'll duplicate that and give this

one a value.

Value equals G and there should be G for general.

Now, we're also going to have PG, PG 13, R and NC 17, so this one will be NC 17

and its value is NC 17.

Then R and R, PG 13, PG 13 and PG.

So now we should have a select let's go make sure it looks all right.

It does.

That's what we want now, our star rating, and that can just be a text input.

So I'll copy this, paste it here called this rating.

And again, this will be like a one to five stars or something.

And this is a no, this is rating.

And copy paste, paste, paste, and then finally the description

div class name equals M.B three.

And we'll have a label, I'll just copy this one.

And this just for description.

And now we want to text area.

Text area, class name equals form, dash control and the ID will be description and the name will be

description and let's give it three rows.

Rows equals three and we'll give it a value of MoVida description.

And then underneath that will have an HRR format, the whole thing.

Make sure it looks good.

It does.

And let's put a in here so.

But class name equals DTN, VCM primary, and we'll give it a value of, say, save.

OK, let's go see how this looks.

That all looks pretty good now we can also.

Give it a component, did mount, functional component did mount, and here we can just say, just to

make sure this is actually going to work, this state equals and we'll put it in a movie and that will

have a title of The Godfather.

And just to make sure the select works as well, we'll say ampe, a rating is set to our.

And I should say that she

did not have an equal sign.

It looks better.

So now what?

I love the form it should have The Godfather with an MP, a rating of Ah.

Let's see if it does and it does.

OK, so that form actually works.

And of course, because this form is a component, we can use it to add a new movie and have the form

blank.

And of course we'll just get rid of this state in the component did malfunction.

And we can also use it to edit an existing movie to pull that information from the database, to extract

it from our back as Jason and use that to populate the form so we can modify it to our heart's content.

So we'll get started on that in the next lecture or two.

## 60 - Making our form a controlled component, and binding it to state

So we have our form displaying properly, but there's still a fair bit of work to do because the way

that react handles forms is a little bit different than the way you might be used to working with straight

HTML and JavaScript.

So this is our form right now.

And let's go back and look at our code.

Let's make a couple of changes.

So the very first thing I'm going to do is go to the very bottom of this and in between the closing

form tag and the fragment, I'm going to put a little bit of HTML in here.

And what I'm going to put in is a div with a class name equal to empty dash three margin top three from

bootstrap and I don't put a price tag in and then I'll put Jassam String Afie in curly brackets.

And I'll pass it three arguments, this state No.

And three

there.

Now, if we go back to the top of this form, there's another oversight.

And this is something I do all the time when I'm working with ASX, where we have label for equal title

for is actually a reserved word.

And we have to use the substitution that we act expects to find, which is HTML four.

So let me copy that and put it on all of the places where we have a label with a four attribute.

There's one.

There's one.

There's one.

Class name there and description should be the last one.

OK, now let's go back and look at our code and see what we have here.

And you can see at the very bottom of the page we actually have a representation of the current state

and we actually need to make some changes to this.

If you recall, this comes from the component did mount.

So let's get rid of that and do something a little more intelligent.

So the top of this form, I'm just going to delete this entirely.

So we're not going to use that right now.

And I'm going to add instead a constructor and I'm going to do that because we need to actually allow

react to be the single source of truth for all information that's displayed on that form.

Now, there's another way to do this, using refs, but it's not recommended.

And I'm not actually going to cover that because you should never do that.

So let me add a constructor

and it has an argument.

Props.

And of course, we have to call super props.

Any time you have a constructor in a react component, you have to call super props.

And I will set the state.

This state will make that equal to and we'll have our movie.

And I'm going to put in all of the attributes that are associated with a given movie and I'll set them

to default values.

I will be zero title will be an empty string and duplicate that a few times.

Next comes release date.

After that comes runtime, after that comes empty, AAA rating, duplicate that, then we have rating

and finally we have description and we also want is loaded and we'll set that default initially and

error will be set to know.

OK, now I'm going to put two more things in our constructor.

I'm going to come back and add more a little bit later on.

But for now, I'm going to do pretty much the same thing or something very similar to what we did when

we lifted state.

I'm going to have to handlers one for any time, a form element or input changes and one for when you

submit the form.

So this DOT and I'll call it handle change, which seems to make sense or will bind that this dot handle

change, dot bind and hand it this duplicate that and change handle change to handle submit.

And this becomes Handal submit now, those are functions that don't exist yet, so we have to go create

those, of course.

And again, it's very, very similar to what we did when we lifted state so as to handle, submit first

handle, submit

and I'll pass it and event

and using the Arrow syntax, we'll say for now when a form is submitted, I'm just going to say console

dialog form was submitted and then I'll prevent default or prevent default.

So that's our handle submit.

And again, when we actually submit the form to our back end, we're going to have to make some changes

there.

But now let's do our handle change.

And this is very, very similar to the lifting state.

So that's going to be equal to or ahead of the event and use the Arrow syntax.

What I'm going to do is grab the name and the value from the given form input that is handed to this

function.

So let value equal the target value and let name not left that name, equal the target name and now

in a bind this to the state.

So whenever a change takes place, we'll say this starts set state and I'm going to hand it an argument

which I'll call previous state by state and using the Arrow syntax,

what we're going to do.

Is simply say movie, because that's what we're dealing with, is the movies the only thing in the form

that we care about and that will have the value of the previous state probe, state dot movie.

And then in square brackets name, and that will be substituted with the name of the current input that

we're dealing with and that will be set to value.

OK, so now we have this handled change event and of course, we need to do something with it effectively.

All we have to do is go behind this event to each of our form inputs.

So let's go up here.

And what the first thing we're going to have to do, well, we'll just get rid of this method equals

post makes no sense.

And while I'm here, I may as well just do this bit on submit will be equal to.

This not handle Sidmouth, so that will handle our form submit and then since I'm here anyway, we have

another attribute we need to worry about and that will be the ID of the given movie.

So we'll add a hidden input.

Input type equals hidden name equals ID.

ID equals ID value will be equal to movie ID.

And we have to bind that and that would be unchanged is equal to this door handle change.

And then we close our tech now I'll copy this line and put that on every other form input.

So right here for title and down here for release date and down here for runtime.

And up here on the select, we can put it right there.

And for rating.

And we're going to have to also change this text area, because this needs to be a controlled component.

So what I'm going to do is delete this, make this is self-closing tag and I'll paste in that on change

event and I'll set the value equal to movie value.

OK, now that should be all that we need in that.

OK, I don't see any errors.

Looks good, let's go check our Web browser.

So far, so good.

Now, if I click and title and start typing now this is my title, you see that the title actually changed

down here where it's displayed and the same thing for the release date.

So if I put this to 10, it changes to 10.

Not what we want.

We want that to actually be a date component.

But we'll get to that shortly.

And runtime that changes MPLX rating changed to PG and it changed it to that didn't quite work PG.

It added one.

So let's go back here and see what we missed on that.

I suspect that the Select doesn't actually have a name associated with it.

Yeah.

So the select name should be MPLX, a underscore rating.

OK, now let's go back, see if our page reloaded.

It did.

Let's try MPLX rating PG 13 14.

I got a typo there somewhere.

Value should be PG 13.

How I did that.

OK, so now if I choose PG 13, PG 13 rating, make it five and description.

This is good.

Perfect.

So now we've got our form elements bound to the state.

So every time something changes in the form element, our state changes.

And that's exactly what we want.

And let's try to submit.

I'll open the console and I'll try saving this form.

And all I should get is a message in the console form was submitted.

Perfect.

OK, this is getting much closer now.

There are still things we can do to improve this form.

For example, let's go back and look at this.

We can actually make some changes so that we can use this same form to edit information from an existing

movie in our database.

And that, of course, would be the component did mount function where we make a call to the appropriate

route in our back end and populate the form when it's displayed.

We can also take advantage of the fact that REACT is component based.

You'll notice that a lot of these inputs, this one, for example, title, there's a lot of HTML in

here that's repeated over and over the same thing here, the same thing in the runtime.

We can actually create components and use those as reusable dumb components in our form.

So we'll get started in those things in the next couple of lectures.

## 61 - Making form inputs reusable components and a Challenge

So this time around, I want to take our various form inputs and turn them into some reusable components.

But before I do that, I had a slight typo.

And I'm sure you noticed that last time around here in edit MoVida J.S. on the input for the text area

description, the value shouldn't be moved up.

Value should be movida description.

OK, now let's change this form so that, for example, this text input becomes a dumb reusable component.

And I'm going to do that by going to my components folder and I'll create a new folder in that.

Which I will call form dash components.

And inside of that, I will create a new file, which I'll just call input thought Jass, and this is

not going to be a react component.

This will be a very stupid, dumb, reusable component called const input.

And that's going to be equal to and it will accept props.

And I'll use the Arrow syntax and I'll just return some jassa, and inside of that, I'm going to return

effectively exactly what a text input looks like so I can just go over to edit movie and choose a text

input like this one and copy it and go back to my input ojars and paste it in there and just change

it as necessary.

So a format of the HMO for will not be title.

It'll be equal to a property that I saw.

Pass it.

Props the name.

And the class name will leave alone the title, I'll just pass it, props that title and the type.

I'll make that.

Crops don't type, so I could, for example, make this a text input or a date input.

I'll leave the class name alone.

The idea, I'll make that exactly the same as the name.

Crops don't name, and the name will also be propped up name.

Props tightening the value, make it props.

The value

and the on change are not always going to be using the same names or actually it props that handle change.

And then if you want to, for example, in some cases, you might want a placeholder, so put a placeholder

in it, which I won't use this time around, but it's something I might use later on, perhaps the placeholder.

OK, and then, of course, I have to export the default, so I will export default and I'll call it

input, of course, because this is called input.

That's what I'm exporting by default.

So I now have input.

So if I go back to one movie to edit a movie, that is what I can do, is just comment this one out

and at the top I'll import.

I'll call it input

from form Dasch components, input.

Now, I can use that like this.

So right here.

I will simply say input

and has its various arguments, its properties, so I'm going to give it a title which will be equal

to in this case title, because it's the title of the movie.

I'll give it the title.

I want this to be an input type equals text, so I have to do it like this, actually.

And title should also have curly brackets around it.

And the name will be equal to the title all lowercase, and the value will be equal to movie titles

and handle change will be equal to the start and will change.

And if I got everything right, if I imported this properly, when I go back to my form and reload this,

it would still work.

I should be able to click in here and type the new movie and it works beautifully and that's much cleaner.

And it also means I can copy this and replace all of the other basic ones.

So this time I'll post that in comment this one out and change this title to release date.

And the type that in this case will be the date and the name will be released and the value will be

a release date

and handle change stays the same.

OK, now I should be able to come down here for runtime and do the same thing.

So I'll paste this in, copy this out and just say, you know, we're going to have to make a change

to the way that the data is displayed here, because JavaScript dates and JSON dates don't play nicely

together.

But it's an easy change and we'll take care of that in a moment.

So this one will be runtime type is text, name is runtime value is movida runtime and handle changes

the same.

Now we'll take care of this one later because it's it's a selection.

We have to create a new component for that, but we can definitely do rating.

So we'll come down here, paste this in, comment this one out and change this to rating.

And the type is text and the name is rating and the value is rating.

And the last one, of course, is another one we're going to have to build called text input or text

area.

So let's you see if the ones that we've changed right now actually work.

Let me go back here and reload this and try the title ABC.

I worked.

Let's try the runtime.

One hundred that worked.

Let's try the rating for that worked.

Now MPLX rating and description will have to be modified to be their own DOM components.

And here's what I'd like you to do.

I'd like you to do.

Don't worry about the select.

We'll do that one together.

But you should be able to take the description, the text area and create a reusable dume component

and substitute that dumb component for this code right here.

This should be replaced with a another dumb component, perhaps called text area, just like the input

one.

And you have enough information at your disposal to do that right now.

So give that a whirl and I'll show you how I did it in the next lecture and we'll get started doing

the one for the text area as well.

## 62 - Solution to Challenge

So how did you make out creating the reusable text area component?

Hopefully you didn't find it terribly difficult.

So what I did was create a new file in my form components folder called Texta at our Jazz.

And it is, as you can see here, virtually identical to the one for input Duchesse.

We just change the input to a text area and add the appropriate properties.

Right here are the attributes right here.

So very straightforward, one constant, which is an arrow function exported as the default function.

And then on edit MoVida J.S., I import it right here on line four and down at the very bottom of the

form.

I left the old one here, commented out and I can delete it now because I don't need it anymore.

And I put it in text area.

And as you can see, these dumb components make things very short and sweet.

So the only one we need to worry about right now, apart from handling the format of the date, which

we'll get to shortly, is this one and that's the select.

And we'll go through creating a reusable dome select component in the next lecture.

## 63 - Creating a reusable select component

So this time around, we want to replace this entire section of sex with a reusable select component.

And not surprisingly, the first thing I'll do is go to form components and create a new file and I'll

call it select dojos, and I'll use the same syntax and logic that I did for the other components.

I'll create a constant which I'll call select, and it will be equal to and will take as a receiver

or an item and make it a standard arole form and will return some X.

So return and I'll just close this so I don't forget to.

And here we'll have a div with a class name equal to that again class name equal to maybe three

and we'll give it a label with a team of four equal to prop up name and we'll give it a class name of

form label

and then I'll have an empty space here in case there is no title.

Otherwise I'll put in Prop Start title.

And again, if there is no title, give nothing.

OK, so there's our label now let's create our select element

select, and its class name will be equal to form select because I'm using bootstrap, so I don't mind

her, including the class name.

But you can make that a prop. if you want to.

The name will be equal to proper name

and the value will be equal to prop value.

And the change will be equal to crops dot and will change.

And we'll close that.

This dude doesn't belong there where it belongs down here.

OK, so let's format that just to get a little cleaner and inside of our select, of course we want

to have our options.

So first of all, put a default option in there.

Option value equals nothing and we'll put it in a proper placeholder.

And then I'm going to assume we're going to be passing a something we can use the map function on,

so props dot options, that's what I'll call it, and then we'll call map and we'll call each individual

entry option and use it arrow function, just like we did in the other many other cases.

And we'll return an option that has a class name equal to form select from bootstrap the key.

I'll make that option.

The ID and the value will be equal to option ID and the label will be equal to optional value.

And we'll close that option and here we'll just put option value down here, live export and export

default select, OK?

I think that all looks right.

We'll find out shortly.

So let's go back to edit movie digest and go to the top and we'll import that import, select from thought,

slash from components, slash, select and go back down and find where the HTML actually is.

So here it is.

Was comment this out.

And put in our solar component and then, of course, we have to populate the options, but we'll get

to that in a moment.

So you select and we'll close it.

And in the middle, we'll put title equals MPLX a rating.

And name is equal to MPLX, a rating also on options, and this doesn't exist yet, but will create

a 10 minute sequel to this state MP, a options.

And value is equal to movida empty AAA rating.

And handle change will be equal to this handle change.

And finally, we'll give it a placeholder placeholder.

I'll just make that equal to

choose.

OK, so there's our select component.

Now let's go back up and actually populate in the state default values for MPLX options.

And I'll just put them right between these so we'll have an entry called MPLX options, and that will

be an array so we can iterate over it using map.

But each entry will be an object ID G and value G.

Now you may have noticed I could have just made this a straight array of strings, but there are many

situations where you're not going to want to do that, where you're going to want to have, say, a

numeric value for the ID and the value and the label will be something else.

So there we go.

That's our first one.

And I'll just duplicate this a few times.

One, two, three, four.

So the next one is PGT and the value will be PJI and then PG 13, which will be PG 13, then ah, which

is R and then NC 17, which is NC 17.

Add a comma here and in theory this should all work.

So let's switch our web browser and find out.

So everything looks good down here, we have those necessary values for MPLX options, and I should

be able to choose a PG and see the MPLX rating changed to PG and it does.

OK, so the next step is to modify the component.

Did mount function in this class to actually fetch data from a remote API and pre populate this form

with an existing movie and we'll take care of that in the next lecture.

## 64 - Prepopulating the form with an existing movie

So this time around, we want to modify our code so that we can pre populate that form with an existing

movie, getting the data from a rest back end.

So the very first thing I'm going to do is in my edit movie dogs, we don't need these lines anymore.

They've been bothering me for a while and I keep forgetting to delete them.

So now they are gone because we're actually setting our state in the constructor.

Now, if we go back to our Web browser right now, we're getting to our movie to ad by saying slash

it means ad.

And I want to make that simpler.

And I'm going to do that by going back to my code and going to abcess and looking at the roots.

And right now I have down here on line sixty one in my code, this is the the path that we're matching.

So we're going to slash admin's ad and I'm going to change that to slash admin slash movie slash and

then I'll give it an ID that can be different.

It can be any number.

So when I'm adding a movie I'll make that zero.

And when I'm editing an existing movie, I'll make that the ID of the movie in the database.

So we change the path.

And if we come back up here, we need to change this to admin movie ad.

And that's the actual link.

Sorry, not at zero.

I'm getting rid of the ad.

That's the whole point of this exercise.

So that's our new path.

And that means back in in edit moviegoers I have access in my component did mount.

I actually can check to see what the ID is, so I'll declare a constant here consed and I'll call it

ID and I will be equal to this start.

Props match programs ID and that's what I called it was ID in the route.

Now I can just check if ID idea is greater than zero then I'll go get that movie, otherwise I'll do

something else.

OK, so how are we going to get that movie.

Well if you remember where we have a means of getting a single movie using Fetch Fetch and we're going

to go to colon slash slash localhost colon four thousand slash V one slash movie slash and then I just

append it and then in my then cause my first one I'll say then my response

and first of all check to see if there's an error, if response star status is not exactly equal to

two hundred, something went wrong.

So I'll do something.

I'll say let error or equal error or create an empty error and then say the message is equal to invalid

response code.

And put the response code, their response status.

With a lower case s.

And set the air, this starts set state.

Error is equal to her.

And of course, I have to return response to Jason, but if we get past that, then and we'll take Jason

Carroll function and here is where I actually need to do something with the date.

And we haven't looked at it yet.

But if you actually played with your form and tried to set the date, you notice it doesn't set quite

right in the form and we can fix that this way.

Const release date is equal to new date.

And here's where we take the Jason movie release date, convert it into a JavaScript date and now we'll

just set the status and I'll do it the long way just so we can see exactly what's going on this start

set state.

And we'll set, first of all, our movie and that will be equal to ID, which is the ID that we extracted

from the URL.

The title will be Jason, the movie title.

The release date, here's what I'm going to format, the date, the way that I needed to be, it will

be release date.

To ISO string, which is the format I wanted in, and I will split that on the tea that is found in

the string that we get and take the very first part of the resulting array, which will give me the

yea yea yea yea dash, dash month, month, dash day, day.

And that's exactly the format that I want.

And then we want the runtime, which is Jason movie runtime, and now we want the AAA rating and that's

equal to Jason movie, the empty AAA rating and the rating, which is Jason thought movie rating and

the description which will be Jason MoVida description.

And that's it for the movie now will set is loaded to true

and otherwise if there's an error.

Then we'll say this starts at St..

Is loaded, is true, but show us the error.

Now, in the Elz clause, the other thing we want to do is if we're not getting the data from the back

end, we still have to set is loaded to true.

Otherwise, we'll just see the loaded loading message.

So we'll say this state is loaded.

True.

And that we'll make sure that we see the empty form with nothing in it.

So now let's go down to a render function and extract all the things we need from the state.

And that would be the movie is loaded and error and we'll do a simple if statement, same as we did

before, if error.

Then we'll say a return divx error and error message.

Or else, if not, is loaded.

Then we'll return

loading whatever else, and we just need to move this curly bracket all the way down to the bottom.

Right here and format this and get rid of the commented out select statement, because we're not using

that anymore, so that's gone.

And unless we missed something, I should be able to go back to my Web browser and go to home and click

on ADD movie and get a blank form.

Perfect.

And now if I change this to a one, I should get the movie of ID one from the database and we do.

That looks good.

So the last thing we have to do eventually is come out of this section.

But I'm going to leave that there for right now just so I can look at the data that we're posting,

because the next step is to actually post a new movie to the database and have that information saved.

So we'll get started on that in the next lecture.

## 65 - Sending data to the REST back end

So things are progressing nicely with our form, and this time around, what I want to get started on

is posting information to the rest API to the back end.

So let's go back to our code and let's get started.

So I'm looking at edit moviegoers and we're down here somewhere.

We have handle submit things online.

Thirty five in my code.

And this is the code that will actually execute when the user fills out the form and clicks the save

button.

So what I'm going to do is write some code that will grab the information from the form, convert it

to Jason, and then we'll write the fetch code to call an API or a root our backend that doesn't yet

exist.

But we'll create that in a little bit.

So the easiest way to do this, first of all, what I'm going to do is move this event, prevent default

to the very top of this function.

So it gets called before anything else.

And I'll get rid this console log was submitted.

And what I want to do is grab the information from the form.

So the easiest way to do this is to create a constant and I'll call it data and that will be equal to

a new form data and we'll get the content for that or the argument for that from our event event target.

So now we have some data and now I'm going to convert that to a payload.

So const payload this is what we'll post is equal to object, the built in JavaScript object class and

we'll call the from entries method on that and we'll pass it data entries to get all of our formed here.

So that's now in a variable called Paillot.

And let's just say console dot log payload.

We won't go any further than that right now.

So let's switch back to our Web browser and I'll reload this to make sure it's actually current.

And I'll open the JavaScript console and I'll feel some information.

So title I'll just call it.

This is my title and I'll pick a date.

So release date will be whatever runtime is.

One hundred minutes.

The MPLX rating is PG 13.

We'll give it a four star rating and some description.

So everything seems to be working well down here.

Our movie State Variable is populated the way it's supposed to be.

Now let's save it and look at the console.

And what we get back is this.

We get an object with all of the necessary entries in that.

So now we have something we can actually post, which is good.

Now, there's no form validation yet and we'll get the form validation in a bit.

I'm not going to validate the entire form, but I'll show you how to add validation to one or two of

the input components and that should be sufficient to get you started.

So let's go back to our code and let's actually write the the request to post this.

So what I'm going to do is create another constant which I'll call request options and that will be

equal to method.

We're going to use post for a method, so we'll have to create the appropriate route on our back end

and we'll put body and we'll just called JSON String of five and pass that payload.

And now we'll just fetch and again, we're going to be fetching to a non-existent route's, we have

to go create that in a moment.

So we'll call fetch and we'll pass it.

HTP Coltons localhost four thousand.

We one and I'll call it admin and just edit movie.

And then we'll get a response which will convert to Jason

lower case, and then we'll just cancel the response.

OK, so this is now set up, except that I need to pass the request options here, which is why I created

those there.

So this is now set up to actually make the call to our back in.

But of course, we actually have to change our back in code to handle this particular route.

So let's go back to our back and code.

My application is running right now, so I'll stop it and hide the terminal and go to my roots file

and let's create a route.

I'll just put it right down here.

So Rohter handle func handler func and it will be htp dot method post

and you are all we want is one slash admin slash edit movie and we'll call a handler that doesn't exist

yet after a movie.

OK, so now let's go to our handlers, our movie handlers and go down here and I'll just take one of

the existing ones that was a kind of a stub code and I'll just call this what we're going to use.

Insert and update will be the same thing.

So I'll delete, insert and change this to edit movie.

And right now I'm just going to have it hand back some JSON.

So just to make this simple or create a type specific to this function and we'll get rid of this eventually.

Jason, response.

And that is a struct and it will have one member just OK, it'll be a boolean.

And in Jason we'll call this OK then I'll create a variable called OC of type Jason response and OK

will be true.

And now I'll just write by Jason, there is a sign the value of right, Jason scroll is up a bit.

We'll give you a status, OK?

And we'll hand it OK and we'll wrap the whole thing.

I'll just call a response.

This is just a temporary placeholder.

And again, if there's an error, I'll just copy one from up here, save some typing, paste this in

here.

And now we have a path we can call.

So let's start our application, go round CMD slash API that's now running so I can switch back to my

Web browser.

I'll clear the console and I'll reload this to make sure it's current.

And I'll just put some data in here.

ABC or AC, I guess, pick a date, pick a runtime NPA rating, put a number in there and some text.

And then when I post this, I should get a response back and I do I get a response right here and I

can look at it in detail and we'll see that OC is set to true.

So we've now successfully connected to our back end and we've actually posted some data there.

And what I want to do in the next lesson is just to add some basic client side validation.

We should have server side validation as well.

But right now I'm going to worry about client side validation.

And what I'll do is just maybe do validation to make the title a required field.

So you can't actually submit the form until you have something in the title field.

And if I do it for one input, you should be able to extrapolate that and apply it to all the inputs.

But we'll get started on that in the next lesson.

## 66 - Client side form validation

So this time around, I want to add some client side validation, and as I said last time, I'll do

one input and that should give you enough information to add validation to your heart's content.

And since we're using bootstrap, I'm going to take advantage of some of bootstrapped validation, logic

or success at least.

So I'm on the bootstrap site right now and I'll go to documentation and I'll go to forms and I'll go

to validation.

And I'm not going to go through this.

But I'll post a link to this page on the course resources for this lecture.

We're going to take advantage of bootstraps success.

So let's get started.

I want to go back to my idea and I'm looking at edit movie dogs and I'll go to the very top.

And the first thing I'm going to do is in my state, I'm going to put somewhere to hold errors and it's

just going to be an array.

So initially it'll be empty errors, plural, and I'll make it an empty array.

And I'm going to scroll down a bit.

And right here on the handle, submit wherever that is.

Here it is.

We're going to do our validation right here.

So I'll just put a comment in client side validation, OK?

And what I'll do is declare a local variable called errors, and it will be equal to an empiric.

And then I'll say, if this state thought movie title, which is the one field that I'm going to make,

requires, so do some validation on that, if that's equal to absolutely nothing.

It's exactly equal to nothing, then I'm going to say errors push title.

So I'll push the name of that field to my right.

OK, and then underneath that I will say this starts at state and I'm going to set errors

equal to my local variable errors, which will be empty if there are no errors but will have one value

in its title if there's nothing in the title field.

So I push that and below that to a simple if statement.

If errors don't length is greater than zero, then return false.

I won't go any further.

Otherwise we keep going, so we'll just do our post as normal.

So this is where the actual validation takes place.

Now clearly there's more information than that required.

So down here, after my handle change and before my component did mount, I'm just picking that spot

at random.

I'm going to write another function called Has Error.

And if there is an error and we're going to pass it, a variable or an argument key, what we're going

to do is return this state dot errors or index of key, not exactly equal to minus one.

So it's true or false.

If we find an error for that particular key, it returns the index of it in the array that's in our

state, the errors array.

And at this point, all we have to do is make some changes to our input for title right here.

And what I'm going to do is simply pass some more properties.

So we'll have to make changes to the input element or the code for that input in a moment.

But I will just push the information right now.

So I'm going to add class name as a property and that's going to be equal to this has error and we're

going to call it using title, the name of this field, and we'll simply do this.

So what I'm saying, if there is an error, then the class name, property two is invalid.

Otherwise it's nothing and is invalid is the class name we're going to use to change the appearance

of that particular field on the form with bootstrapped classes.

So typos stay the same name will stay the same.

We're going to go down here now and just add a little bit more information.

I'm going to add one called arrogant and that will be equal to similar to the first one.

This start has error and again will pass the title

and this time will add text Dange.

Otherwise, we'll add dash none.

It used to be hidden in bootstrap for, but now it's Denon, do not display this element.

And we're also going to add one final thing and that will be under error message or call it error MSG,

and that will be equal to whatever I want the error message to be.

Please enter a title.

OK, so that's it for the input.

And that's all the changes I need at this level to perform client side validation.

But I still need to make some changes to input charts.

So here I've opened up input just and what I'm going to do is below the overall input.

But before the closing div, I'll add another Dave and its class name will be equal to from my properties.

The error Dave.

And inside of that divide, output from my properties, the error message.

OK, so this is where the error message will appear.

And initially, because there's no error, it will have the class of Dynon, which will make it hidden.

So now back on the input, we need to add a few more things at the top here just before ID put it right

at the very top, just before type, I'll put one called classman.

Actually, it's right here I have class.

I just need to modify this one that's going to this will go away and we'll change it to in curly braces

and then in back to it, because I'm going to use JavaScript templates, form Dasch control that will

always be there.

And then in a dollar sign with curly brackets popped dot klutznick.

And that should be it for that one.

OK, so that actually should work.

So let's go back and look at our application right here.

Reload this just to make sure it's current.

And I'll try submitting this.

And as you can see, it says, please enter a title and gives the nice bootstrap formatting on this

required field.

But if I put something in there and this is my only required field at this point, it should colourist

back in and it does.

OK, so that is an example of client side validation.

And as I said, it's pretty simple to implement this for all of the various things that are on this

particular form or for any form for that matter.

But I'll leave that as an exercise for you because this is enough information to get you started performing

client side validation.

So let's move on.

## 67 - Receiving data on the REST back end

So reform seems to be working pretty well so far and we have validation for only one field and again,

you can add all the validation you want.

It should be pretty straightforward at this point.

So it would seem that we could go back to our backend code.

And I'm looking at movie handlers, don't go and go to the edit movie function.

And it would seem that all we'd have to do at this point is capture the Jason, read it into the movie

struct and then write the necessary code to put it in the database.

And that's not quite true.

But let's give it a try and see what happens here.

OK, so I'll put the code right here, right at the beginning of the function.

So the first thing I'll do is declare a new variable over and I'll call it movie and type models dot

movie right there and then just reading my JSON from my request.

So there is a sign the value of JSON New Decoder, and we want to decode the body of our request and

decode it into our movie variable and check for there.

And I'll just copy this error code right here and paste it up here and change this assigned to it equals

so we get rid of that error.

And if that work properly, I should just be able to go log print line and read movie, say total and

print that to the console.

So let's see what happens.

I'll stop my application and start it up and go back to my Web browser and open my JavaScript console

and I'll clear this out and go to home and then back to the ad movie page and put in my I'll just put

in dummy data test title.

The release date will be any validate.

The runtime will be one hundred minutes.

I'll choose PG for an NPR rating, give it a four four stars and some text for the description and I'll

save this and you'll see we get an error here.

Ad request.

Now why is this a bad request?

Well, it doesn't seem to make much sense.

Let's go back and add one more line to our code here and see what's going on.

So we'll do a log print line and we'll just print out the error.

So I'll stop my application, start my application, and this time I'll put some blank lines in here

so I can see the error message and I'll just save this and get the same bad request error.

And there it is, bad request, but it is posting the necessary JSON.

But here we have the error, parsing time, and that's the date that we're getting out of the JSON and

we can't pass it.

So there's something wrong here and actually will be even if we fix that error, if we back here and

look at the actual posted JSON, if you look closely, you'll see that ID is being passed as a string.

In fact, everything is being passed as a string.

Now, how are we going to fix that when we have a few options?

We could go back to our front end code and instead of just grabbing all of the form fields, as in one

line of code, we could manually construct the JSON in the correct format.

But that would be a pain.

What if we have to have this same form in a different handler somewhere else?

Well, that's a bit of a problem because that's an awful lot of code duplication.

And it seems to me that that's a lot of code to write.

Or we can go back to our backend code and look at our models and open the the movies.

Dasch DB.

No, the movies, Dasch go file and we could take this movie struct and write a custom JSON decoder

and encoder, in other words, a custom mercial and a custom unmerciful function and extend the actual

decoding or encoding of JSON.

But that's an awful lot of work and we'll run into a lot of headaches trying to do that.

So it seems to me the simplest thing to do here is to go back to our handlers and close this terminal

window and let's define a new type.

So I'll put the typewrite outside of this function in case I need to use it somewhere else and I'll

call this type.

Movie payload, and it will just be strapped.

And I'll give it the same fields as our movie type, so I'd but this time I'll make it a string because

that's what I'm getting from my JSON payload in Jason.

That's called ID.

And I'll duplicate that a few times.

And the second entry would be the title.

And in JSON that's called title all lowercase and to is description.

And that's called description all lowercase, and then we have the year, which is year.

You're sure it's not coming from Jason?

I'm going to put it in there anyway, then we have the release date.

Which is called release date, and I'll duplicate that again, then we have runtime called runtime.

And rating,

which is called rating and finally AAA rating,

which I'll call which is called AAA rating like that, and a format this now a rename this movie Variable

to Payload, just so I know what I'm dealing with.

And it's not of type models movie.

It's going to be of type movie payload.

And down here, I'll just rename this variable to correspond to the actual one we're decoding it into.

And now down here, all loved print line payload title and I'll save this to clean up the imports and

I'll stop the application and start it.

And give myself a few blank lines, and now when I go back to my Web browser and save it this time,

clear the console, we should not get an error and we don't.

And if we look at our terminal in our IDY, we actually have the title.

So now I've successfully read my payload, my Jason payload into my temporary variable called payload.

So at this point, all I really have to do is now declare a movie variable for a movie of type models

movie

and I'll just assign all of the content from my payload struct into my movie struck doing the type conversions

as necessary.

And I'm going to do this without checking for errors just to save some time, because I'm sure you know

how to check for errors and go.

So the ID will be used sterkel in package to and we're going to get payload ID and then we get the title

and that's just a string.

So MoVida String or movie.

MoVida title equals payload title

and the description

matches description and then the year movie year I shall do movie, don't release date

movie, don't release date and I'll ignore the error is equal to time to pass.

And what I'm getting is the the ISO standard format.

So I'll use goes rules for matching a date pattern and that comes from payload.

The release date now I'll do the movie, the year is equal to movie release date.

Yea, and I'll get the runtime MoVida runtime, and that needs to be an ant, so I'll use destruction's

package again, a two hour payload dart runtime and after that I get the rating slashed.

Just the rating.

And that has to be an informant.

And the last thing I need to worry about right now, reading from the Jason is MoVida MPLX rating,

and that's equal to Haloed dot MPLX rating and then movie created.

It will be equal to time dot now

and movie dot updated up is equal to time dot now.

And at this point we're almost ready to write to the database, but let's make sure it worked.

Logged on print line will say movie dot year because that's when we're converting from one format to

another.

We're actually passing the release date from a string into a time dot time and the year we're getting

from our release date in the movie variable.

So this should work.

So let's stop.

Our application started again.

And go back to our terminal or to our Web browser and just click save.

So no errors.

And now in the console, I should see the year twenty twenty one.

Perfect.

OK, so the next step, of course, is to write this to the database.

So we'll close our terminal window and we'll open up our movies.

Dush db dot go file where we have all of our database routines.

So we have a get, we have an awful lot to keep going and find out where we want to put this or put

it right at the very end.

We just couldn't function funk.

And it has the receiver of a pointer to DV model and we'll call it insert movie and it will take one

parameter movie of type, the movie and it returns potentially an error.

Now I'll just go up and copy the contact stuff to save some typing and come down here and paste it in

and let's write our statement.

Statement is assigned the value up and it's very straightforward.

Insert into movies and we want to insert title description here, release date, runtime rating, MPLX

rating and we'll go to the next line, created us and update it up.

And we're going to get these values which are placeholders and we have it looks like nine dollars and

one dollar sign, two dollar sign, three, four, five, six, seven, eight, nine, and close our

parentheses.

So that looks right.

Now, let's ignore the first return item, but check for the error and we call em DBE, the exact context

and the context and add our statement.

And now we just put in all of our entries.

So MoVida total and duplicate that a few times.

The second one is the description than the year.

Then the release date

and runtime

rating that again rating duplicator empty AAA rating.

Created and updated as

we check for an error, if error is not equal to nil, return error or otherwise return nil.

So now we go back to our movie handlers and at this point right here

will check for an error is equal to app models, DV, dot, insert movie.

And we had a movie.

And again, we'll check for errors or copy this this check for the error pasted right below that.

OK, now the next thing I want to do is I have this Jason response type and it's it's specific to this

function and I'm actually going to change that.

So I'll cut this

and at the top of the file I'll define that type and paste it in here.

And I'm going to add one field to it and I'll call it a message of type string and in JSON or call that

message.

Let's go back down here to our edit movie function and just clean up some of this stuff.

So that should be gone.

We don't need that anymore and we don't need that anymore.

And that all looks right.

So now I should be able to start my application, start my application.

Go back to my Web browser.

Clear the console and let's put an actual movie in here, so I'll put in, say, The Princess Bride,

which was released in nineteen eighty seven so far in nineteen eighty seven.

A long time ago in October.

And it was on the 9th and no, I don't have that information at my disposal.

I did, in fact, look it up and it's runtime is 98 minutes and it's PG and it's rating.

It's a five star movie.

And I copied and pasted the description.

So let me paste it in there and let's save this and see what happens.

And we have an error message.

So let's go back and check our error messages.

And it probably is in movies.

Dush db certain the movies are missing a comma.

That's what it is.

I think that all the time.

So let's try that again.

Stop it started.

And go back to my Web browser and clear the console and try it one more time, save and we still have

an error.

So let's try this long, not print one.

I think it's probably here,

undoubtedly a typographical error somewhere.

One more time save.

We should still get the error.

Operator does not exist.

Integer.

Look at that.

Tell me how that happened.

I guess I'm getting sloppy in my old age, third time is the charm.

Let's try this again.

Save that looks better.

So now if I go to movies, I should see The Princess Bride and be able to click on it.

Perfect.

OK, a little bit of sloppy typing there, but we finally got the job done.

So the next step is to actually allow us to edit an existing movie and we can use this same function

in our handlers.

We'll just have to have a different database function instead of calling in.

Certainly we'll call update movie and we can do the check based upon the ID that we get from the request.

So if our ID is zero, we're adding a movie.

If our idea is not zero, then we're updating an existing movie and we'll get started on that in the

next election.

## 68 - Providing feedback with a reusable alert

So we have our adding a movie working correctly, it saves the information to the database, it doesn't

provide any feedback once I click save unless I'm looking at the JavaScript console.

I don't know that anything happened.

So we want to address that.

And the other thing we want to work on is right now I'm looking at it means movie zero.

But if I'm looking at one, for example, it does in fact populate this form with the correct information.

But I want to click save.

It's actually going to save that as a new movie in the database.

And we need to modify our handler so that it updates the movie in the database.

So that's a couple of things we want to take care of.

Let's get started on those.

So right now, I'm going to switch back to my colleague and I will find my back end code, which is

right there, OK?

And my application is running.

So I'm going to stop that and I'll hide this.

And right here in the handler edit movie, I'm just going to make a little change right now just to

give us something to work with, OK, without changing our database a lot.

So I'm going to add an if statement here.

If MoVida ID is equal to zero then I'll insert the movie otherwise for right now I'm not going to do

anything.

OK, so let me format this.

That looks right.

So that's the one change I'm going to make at the moment.

So let me start the application again.

So if I hide this and go back to my Web browser and click save right now, it shouldn't save a second

copy of The Shawshank Redemption.

OK, so finished.

Let's go to movies and we only have one Shawshank Redemption, which is great.

So let's go back and look at that one again.

One.

So that's showing us The Shawshank Redemption.

So what I want to do right now before we update things in the database, let's figure out how we can

provide some feedback to the user and we'll get rid of this because we don't need this here anymore.

So let's go back to our idea and find our front end code, which is right here, OK?

So I'm going to hide my terminal window and I'm going to go to app or to edit movie dojos with the very

bottom.

I'll find that part where it's writing me this stuff.

I'm going to get rid of this and not show the session or the state anymore because we don't need that.

So let's think about how we can modify what's being returned right here.

How about this?

How about we put a message that displays when the appropriate entry in the state changes between the

title and the.

So right here.

So we're going to put something right there.

What are we going to put there?

I mean, I could if I wanted to just write it div and away we go, but instead let's make it a reusable

component.

So I'm gonna go back to my components folder and.

Oh, right, click on this and choose new folder.

I'll create something called UI components for user interface.

That's a folder.

And I'm going to create an alert and I'm going to take advantage.

Let me go back to my Web browser of these from Bootstrap, these nice alerts.

If something succeeded, if we have a success message, we'll use this one.

And if something failed, we use this one.

And it's really quite straightforward.

So I've got a copy, one of these examples, and go back to my form components or my UI components and

I'll create a new file in there and I'll call it Alert Jass.

And this will be very similar to our input components.

I'll just get a constant I'll make it dumb and pass properties to it.

So constant alert and it receives prompts and it's an error function and I'm just going to return.

And inside of that I'll paste HTML.

Just copy and I'll tab it over so it looks better and I'll make this part a property.

So instead of that I'll say props dot alert type and that will be the class name and it will be alert

dash success for the green one and it'll be alert class errors for danger.

This isn't quite right.

Let's fix this again.

This needs to be like this here and that's right.

And then this and then close it with that and get rid of these quotes and that should fix it out.

OK, and this rather than being a simple alert, check it out.

We'll just change that to prompt that alert message.

OK, so now we have an alert that we can use once we export export default alert.

OK, so now we have this component that we can use.

So let's go back to our edit movie Jazz and let's put a few things in here.

So right at the bottom of state, let's add an entry to hold our alert information.

We'll call it alert and we'll make it an object and it will have.

Two entries tight, and I don't want it to show up at all to start with, so I'll add the class dema

as the default means the alert will be invisible and message will be an empty string.

OK, now where do we want to set this information?

Well, obviously we want to set it after the form and submit it.

So let's go down to the part where we submit the form, which is right here.

Right now.

We're just doing a console log data and we need to make that a little more intelligent.

OK, so we know if everything succeeded, we're not getting back an error.

So I can just check to see if the JSON has a key of error.

Just by saying if DataDot error, then we have an error message display or else we don't.

So if data error, what I'm going to do is actually set my state variable for alert, I'll say this

dot set state.

And we'll make alert equal to type alert danger, which gives it the correct class name and the message

will just be whatever we got from the error message.

So data error message.

Otherwise, let's just copy this and pasted in here and make that alert success.

And I'll just say my message is going to be changes said.

Now, I want to use this state information for alert with the component we just built in alert.

So let me go import that so we'll say import alert from that.

You are components.

Now it's available.

Now we'll go down in my returned HTML after the title and before the hour right here.

And we'll pass it.

At the same level alert and I'll put it on the next line so it's a little more readable, we'll say

alert type equal

this dark state dot alert dot type.

And alert message equals this state alert dog message and then close that tag and format, everything

there.

OK, so that's saved.

Now if we go back to our Web browser and go back to watch movies and make sure we're looking at one,

we are looking at a movie slash one slash admin slash movie, slash one.

I'll open the JavaScript console in case there's any errors.

And I have it in Valladolid property class.

I made a mistake.

So let's go back here and look at alert Jass and change class to class name.

That's the danger of copying and pasting.

Go back, clear the console, reload to make sure everything is valid.

It is now let's try saving this change saved appeared up here exactly as it should.

OK, so that's some feedback and that's really good.

And I'll leave it as an exercise to you to introduce an error in your back end code so it can send an

error back, but it will absolutely show it in the correct format.

OK, so that's a good first step.

Now in the next section and the next lecture, what we'll do is write the necessary code in here.

Back in our back end window here.

We'll put an else statement here that allows us to actually do an update of a movie in the database.

And of course, that means we'll have to go write the appropriate function in the database routines

or movie stars.

Dargo but that should be really straightforward.

So we'll take care of that in the next election.

## 69 - Editing an existing movie

So this time around, we want to update existing movies and the logic for that's pretty straightforward.

So I'm looking at movies, Debe go in our back and project and I'm just going to copy this entire function,

insert movie and pasted and rename it to update movie and fix this query.

So we're going to not we're not going to insert into movies, we're going to update movies and we want

to set title equal to dollars on one description, equal to dollar sign two year equal to dollar sign,

three release date equal to dollar sign four.

And I want to make sure that comma is there this time.

And I don't have any hashmark supposed to have dollar signs.

So far, so good runtime equals to dollar sign five rating equal to dollar signs six MPLX rating equal

to dollar signs seven.

We don't need created out so I can delete that because that's not going to change on update and updated

at all.

Equal to dollar sign eight where ID equals dollar sign nine and that should be our query so we can delete

all this.

Now let's make sure these match we don't need created out.

So that goes away.

And the last parameter, the last substitution is movida by deed.

So title description, your title description year release date, runtime rating, release date, runtime

rating and a updated an ID and that's it.

OK, so that function now exists and that's pretty straightforward.

So let's go to our handlers and we're looking at the rear in the file movie dush handler Stocco when

we're looking at the function edit movie and we do need to make some changes here and they're not that

difficult at all.

Basically here I have this.

If movie ID equals zero, that's fine.

But what I really want to do is to make sure that if I'm updating a movie, I extract the information

that's already in the database and only update the things that have actually changed.

And I want to make sure that's as simple as possible.

And I make sure that this code is easy to update in the future.

Now, what I'm going to do is right here, after I declare that movie variable, I'll put an if statement

and I'll check to see if haloed ID is not equal to.

And remember, we get that as a string, so we check for the string value of zero that I actually want

to get the ID.

So I'll get the ideas and antibusing ID and I'll ignore the error is equal to stir confort a two and

you shouldn't ignore the error in production code.

But of course I'm trying to teach a concept here and we already know how to check for errors.

So we're looking for payload.

ID like that.

That gives me the ID as an integer.

Now I'll create a new local variable called M and again I'll ignore the error and that's equal to app

models.

I'll just get the current movie ID get ID.

So now we have that now get returns movie.

As you can see here in the preview or the helper text that my ID gives me, it returns movie as a pointer.

A pointer to models.

Movie and movie is not a pointer.

That's why I put it in this local variable m because all I have to do now is say movie equals star M.

OK, now I have the movie variable which is declared on line 119.

It's now populated with the existing values from the database.

The only thing I want to change here is MoVida updated and I'll set that equal to time.

Do dot now because that is the current time.

We're actually changing it at this point.

And down here, the first part will stay exactly the same.

We'll just add an else clot's else.

If we're not adding a movie, we must be updating a movie.

So error is equal to after models.

The DB update movie and had it movie and we'll check for an error copy of this code.

Paste it here.

And at this point we are actually done.

So I should be able to stop my backend application, start my back end application, everything compile.

That's a good sign.

Switch to my web browser.

I'll reload to make sure this is actually the current page and it is.

And I opened my JavaScript console to look for errors and I'll go to add movie and I'll bring up, say,

movie ID one, which should be The Shawshank Redemption.

And it is.

OK, the only thing I'm going to change right now, just to make sure this works is I put a period at

the end of this sentence which it should have just to be grammatically correct, and let's save it and

see if we get everything right.

Syntax error at or near dollar sign one also have a mistake, some we see that our error actually works

properly.

So let's go fix that back to our back end.

And I have a typo here somewhere.

Missing an equal sun.

And I was so careful what I typed.

I really was because I made mistakes last time around.

All right.

Let me copy or stop.

The application started up again and go back.

And I shouldn't have to reload the page.

I should be able just to save it.

And this should change if I get everything right.

Changes saved.

Perfect.

Now, let's go look at movies.

Now let's bring up The Shawshank Redemption.

And there is the period right there.

So updating actually works the way it's supposed to.

OK, so the next step is obviously we can add a movie, we can update a movie, we can look at a movie.

We need to be able to delete a movie.

So we'll take care of that in the next lecture.

## 70 - Deleting a movie

So right now, we can display a movie, we can add a movie and we can update a movie, and the next

step is to implement the functionality to allow us to delete a movie.

So I've got my application running both the front end and the back end.

And I'm looking at the ad movie screen at a movie screen and I'm going to change the URL to display

The Shawshank Redemption I'd want.

And what we'll do is we'll add a couple of buttons here, first of all, and we'll do this right away.

We'll add a cancel button here, which the users expect to see.

So we'll click on Cancel and we click when you click on it.

It'll take us to the Manege catalog screen, which right now just displays a title and nothing else.

But eventually it will display the list of movies and you can click on them and go right to the edit

screen.

And then if we're not adding a movie, but only if we're not adding a movie, we'll add another button

here that says delete and that will, of course, delete the movie.

So let's get started.

So I'll go to my code and I'm going to be looking at the front end right now.

We'll take care of the front end first and then we'll do the back end later.

And I'll scroll right down to the bottom of the screen just beside this save button.

And I'll add a link here.

And remember, we're using the reactor, so I'm not going to go a F equals.

Instead, I'll use the link functionality and we'll use a link.

And I let Visual Studio Code Auto import it for me.

And now we're going to go where we're going to link to and the world we want to go to on our front end

is slash and admin and now we'll add some classes to this class name equals and we'll make it a button

and we'll make it a button morning and I'll give it a little bit of padding on the left emergen M.

S Dash one and the text will just be canceled.

OK, so we'll save that.

And the auto import, of course, at the very top here, added import link from reactor at or dome.

OK, so now we have a cancel button and if I click on it, it takes me to this screen and we'll be working

on this one a bit later.

But it gets us where I want to be.

So let's go back to add movie.

And this time let's add once again one.

Now I want to display a delete button here, but only if the movie ID is greater than zero.

In other words, if my new URL is slash, it means movie zero.

We're adding something.

So I don't want to see the delete button here, but if it's not zero, we do in fact want to add that

button.

So let's go back to our code and right at the bottom here next to our cancel button, we'll just put

a little bit of conditional rendering logic here.

And it's a little bit odd the way the conditional rendering works, but let's have a look at it.

MoVida ID.

So in curly brackets, movie ID is greater than zero and then you have to add double ampersand, which

makes sense.

But I'll explain why it makes sense right away.

It is not, shall we say, intuitive.

Then in parentheses I'll put a F equals and I'll have it linked to nothing.

So I'll do that by going hash bank and then I'll add an unclick function unclick with a capital C equals

and in curly brackets we're going to use the S6 syntax this dot and I'll refer to a function that I

haven't written yet.

Confirm, delete and close my curly brackets and then we'll give it a classmate class name equals and

again we'll make it a button, but because of the delete button we'll make it button dange, which will

make it a nice read and again give a little bit of margin on the left hand side and the text is just

delete.

So I need to go create this can confirm delete function.

So I'll just do it up here at the top before my render function or add a function.

Confirmed delete is equal to and opposite the event and using the arrow syntax.

I'll just log to the console console dialog would delete movie ID and this state MoVida ID.

OK, so it's not doing anything yet but it will at least write to the console.

So if we get everything right we should be able to go back to our web browser.

And now I have a delete button and let's open the console and make sure it works.

So there's the console and we'll clear all those old errors from the last time I was playing with this

and I'll click the delete button and it says would delete movie ID one.

Now let's change that to zero.

And the delete button does not show up perfect.

OK, so in the next lecture, we'll start implementing the actual logic to delete a movie.

## 71 - Adding a confirmation step when deleting movies

So this time around, we want to continue to work on our delete functionality, but before we do that,

last time around I suggested that the inline conditional rendering to decide whether or not we display

the delete button, that it was less than intuitive.

But it makes sense.

And I forgot to actually tell you why it makes sense.

So here in my code on line two twenty four, I have the expression movie ID is greater than zero and

then an ampersand followed by something.

And in JavaScript.

True.

And some expression always evaluates to expression and false and expression always evaluates to false.

So when you think about it, this will always return.

True, if movie ID is greater than zero and it will always return false if covid is less than zero or

equal to zero.

So that's that.

So let's move on.

So this time around, what I want to do is when someone clicks on this button.

So let me go back to my Web browser.

When someone clicks on delete, I don't want it just to automatically delete it.

I want some kind of confirmation and there's lots of ways we could do that.

We could use the simple JavaScript confirmed, but that's pretty ugly.

So I'm not going to use that.

Or we could use the bootstrap model and put a yes and no button in that.

But that's a lot of work.

And honestly, there's a really simple package and I'm going to show it to you.

It's called React, Confirm, Alert.

And I'll post a link to this and of course, resources for this page.

This is a really simple one.

So what I'm going to do is just grab the install code, go back to my terminal, stop my application

and paste that in there.

And this will install the necessary package.

Now, I'm a bit of a purist and I install a few packages as I possibly can.

But you're probably aware of the fact that there are many, many packages available for react, including

some very good ones.

In fact, there's one for React Bootstrap, which makes working with bootstrap in react a little less

cumbersome.

But in this course I'm going to try to stay as close as I can to pure react.

Nevertheless, react, confirm.

Alert is really easy to use.

It's relatively small and it's extremely popular.

So let's give it a whirl.

All right, so you can see that I should be running NPM audit fix to fix some vulnerabilities, but

I'll do that later on.

Right now we have it installed.

So how do we use it?

Let's go back and look at the examples.

So lots of options we can play with and I'll let you read those at your leisure.

Let's go down here.

So here we have something that a class app extends Riak component.

That's exactly what I want.

So I'll copy this code.

I don't need the submit function, but I need the stuff in the middle and I'll copy that and go back

to my code and find my confirmed delete function wherever that is.

There it is right there.

So I'll leave the console log in there for now and then I'll just paste in the code.

I just copy and of course I need to import that.

So let's go back and see how it's important right here.

So I'll copy these two and go to the top of my code where the imports are and paste those in and get

rid of the comments.

And go back down to my confirmed delete function, which is just before my render function.

Right here.

All right, so I'll change the title from confirmed to submit to delete movie Questionmark and change

this to.

Are you sure?

Yes.

And then right now, it's just going to if I click yes, it will actually just alert click.

Yes.

And if I click no, it will alert.

Click No.

Well, I don't want to do anything on this right now, so I'll simply do this.

Just give it an empty function.

And on this one, I'll leave that right now and make sure this actually works.

So let's start our application, will clear the console and start.

He will fire up a new browser window and there it is.

So let's go to Ed movie, change it to movie ID one so we can see the delete button.

And there it is.

And this should give me some kind of confirmation alert.

And there it is.

OK, so you can style that to your heart's content.

The documentation is right in the link that's on the course resources for this lecture.

But as you can see, if I click yes, it should say click.

Yes.

And if I say no, it should just make the dialogue go away.

And it does.

OK, so there's a little confirmation step for us.

Now, one of the thing I want to draw your attention.

I'm going to open the console, OK?

And I'll clear it and I'll go back to my code and hide the terminal window.

And down here where we actually put this button right now, we have unclick equals.

And I use this syntax.

Now, watch what happens if I get rid of this part.

Cut that out there, OK?

Notice how, first of all, it through an error here.

And secondly, this just fired as soon as the page loaded.

Well, that's why we're using this syntax.

And when you're putting an unclick handler like this, make sure you use this syntax.

Otherwise, this is just going to fire no matter what happens.

So we'll go back here and everything should be good once I say no and reload this.

And that's exactly what we want.

So there's no new errors being generated over here.

So that's just something to bear in mind.

And it's an easy mistake to make.

I make it all the time and have to go fix it after the fact.

And almost certainly you will as well.

## 72 - Implementing delete on the back end

So let's write the code in our back end, necessary to delete a movie.

So I'm looking at the movies, dash DB go file.

And at this point in the course, this should be really straightforward for you or create a new function

and this will be our database function.

So it takes the appropriate receiver and I'll call it delete the movie and it will take one parameter

ID type it and it'll return potentially an error and I'll get my contact code just by copying and pasting

this because it's ever so much faster and pasted in here.

And I'll write my query.

So my statement will be, and it's very straightforward, delete from movies where ID equals someone.

And let's just scroll up a bit and now we'll do the actual call to our database driver and the exact

context.

And we're going to have the context and our statement.

And it takes one parameter ID and we'll check for an error.

If error is not equal to nil, just return the error, otherwise return nil.

Now let's go to our handlers and I believe we already have a stub.

Yes, we have a stub delete movie.

So here we want to get the ID from the URL so we'll do the same thing we did before.

Pyramus is a variable that comes from HTP Rohter and we're going to call Purim's from context and we'll

have the context context.

Then we get our ID and check for an error by calling Starcom A2A and we want to get Purim's by name

and we're looking for ID and we'll have to get that in the early when we set the route up.

In a minute, we'll check for an error and I'll just copy and paste this code.

And if there is no error, then we check for error again and errors equal this time to abduct models.

DB delete movie right there and we handed ID and we paste in our error check.

Otherwise we can just define a simple OK is assign the value of JSON response and give it an OK set

to true.

And we just right at our Jason file.

Error equals app dot, right, Jason?

And we handed W and http dot status.

OK,

what we want to encode is Jason and wrap the whole thing in a response and a page to my error check.

So there's our delete movie handler.

Now let's go to our roots file and set up the necessary root.

So I'll duplicate this one because that's a lot faster.

This will be method get.

There's no need to post this because we're getting everything we need from the URL and we're calling

delete movie

and we'll hand off the ID and we pass it to the right handler to the movie.

All right.

So that should compile.

Let's make sure.

Perfect, so our back end is now set up at least enough to get things working, so in the next lesson,

we'll go and hook our front end up to this point in our back end.

## 73 - Connecting our delete button to the REST back end

So our final step is to actually connect our delete button on our front end to our delete functionality

in the backend, and that's really simple.

So before we do that, I don't want to delete one of these movies we have right now because I'll be

using those.

So instead I'll add a test movie to delete.

So I add a movie test movie and the release date will be any validate because I'm going to tweet this

and the runtime will be 101 minutes and we'll make this one, say, PG 13.

And the rating is not a very good movie.

So we'll give it two stars and the description is some text.

So I'll save this.

OK, so it's now saved.

And what I'm going to do next is actually look at my list of movies and find out what it's ideas.

So test movie in my case is already 10, so I'll go back to add movie and change this to a 10.

OK, so it's been entered into our database and now I want this delete button to actually do something.

So let's go back to our code, our front end code and we'll look at this file edit movie digest and

we'll find the part where we actually delete it.

And if I recall correctly, it's just before the render function.

And there it is confirmed.

So let's get rid of this this council logs, we don't need that anymore.

OK, and and here this unclick function.

OK, we're not going to say click.

Yes, we'll get rid of that and we'll put some curly braces because we're going to do a number of things.

So what we'll do, first of all, is use fetch, fetch, and we're going to fetch to our new route in

our back end HDB colon slash slash localhost cohort four thousand slash.

That means slash, delete, movie slash.

And then we just need to append the idea of the movie, which is this state.

Movie ID, and just to be clear, I'll add the second parameter, just to make sure that everyone knows

when they're looking at this code that this is a get request.

Then we'll take our response and convert it to Jason,

and then it'll be pretty much the same thing we did before.

But I'll type it out here.

So we're getting data and we're going to hand that to all of this if there is an error.

So if DataDot error, then we'll do one thing else, we'll do something else.

And what we're going to do here is the simple alert functionality.

So we say this set state and we want to set alert to be of type

alert danger to give it the red background.

And the message is just whatever is handed to us in the Jason file dated error message.

And then we'll copy this and pasted here and modify that to alert success, to make it green and change

the message to movie deleted.

Now at this point, if we stop there, we'd get a nice, alert movie deleted appearing at the top of

the page and we'd stay there with the movie still displaying in the form.

And that seems kind of pointless.

So maybe we shouldn't do this alert at all.

Let's get rid of that instead.

Let's take them somewhere else.

And we're going to do that, of course, because this is a single page app.

We actually have to use the history.

So this props dot history, dot push and I'm going to push it to pathing and it has to be like that

all lowercase letter and I'll take them to the admin screen.

OK, so let's try this out.

Let's go back to our Web browser.

We'll look at the list of movies.

And as you can see, test movie is still there and it has an idea of 10.

So let's go to add movie and let's change this to 10 to bring that movie up.

Am I delete button should appear when I do this and it does.

And now when I delete, we we should be taken to the the manage catalog screen.

Yes, so there we are.

And now if we look at our list of movies, it's still there.

Why?

Well, we made a mistake somewhere.

Let's go back and look at and see what we called for a euro, because that's probably it.

So we left TV one out of our URL.

So there we are, the one.

Let's go back here and go to home and go to a movie and change that to attend

and delete it.

And now it should delete it.

So now if we go back to our list of movies, it's gone.

All right.

One typographical error isn't too bad.

So now we have to create we have show, we have update and delete or from the CRUD acronym, create,

read, update and delete.

All of those are now functional for our application.

Of course, we're not finished yet.

I mean, one thing we're obviously going to have to do is we don't want this ad, movie or managed catalog

to be available to just anyone.

We want only the people with the appropriate credentials to have access to that.

Also, we probably should do a little bit more validation on the Jason we're sending to our back end.

So on the back end, maybe we should implement a little bit of validation there just to be absolutely

safe.

So we still have some work to do, but it's coming along nicely.

## 74 - Challenge: displaying list of movies to edit

So this time around, I want to do a little tiny bit of cleanup and then I want to give you a challenge

and we'll do the cleanup first, you may have noticed and you probably fix this yourself, but just

in case, let's clean things up.

Whenever I start my front end application, I get these warnings that these are defined but never used.

And this is an abduh.

So let's get rid of the offending things, which would be these two things and would also be these two

things.

And that should take care of it.

So now let's start our application of the console and start the application and those warnings should

go away and then I'll show you what I want you to do.

And I already have it done.

So here's our application and everything works the way that it should.

We see the list of movies.

So there's a couple of things you've probably noticed by this point.

But let's clean them up now just to be safe.

So, for example, under managed catalogue, I want the list of movies to show up here and that's something

you're going to have to do.

And when I click on a particular movie, it should take me to the edit screen, and that's straightforward.

But the other thing you may have noticed is whenever you save a movie and I've already fixed this,

so I'm not going to click save if you if you're displaying the movie for the first time, it in fact

makes the changes.

And really it shouldn't leave you on this form.

We're putting a little success alert up here and that works fine.

But when I click on Save, I want you to take the user back to the Manege catalog screen.

OK, now none of that is terribly difficult.

And the places you're going to have to make these changes, of course, are let me go back to my editor.

The places you're going to have to make these changes would be in this file right here, admen, which

is and this is a rather heavy handed hint, going to be remarkably like this file movies.

So those would be the two places you're going to have to make the changes for or the one place you're

going to have to make the change for the admin screen.

The other place you're going to have to make the change, of course, is in edit movie after the user

clicks the save button, rather than simply displaying that little alert at the top of the screen,

take them to a different page.

All right.

So give that a whirl.

This shouldn't be terribly difficult for you, but it will get us much closer to where we need to be.

And in the next lecture, I'll show you how I did it.

## 75 - Solution to challenge

So how did you make out with the challenge, hopefully you didn't find it terribly difficult and hopefully

you actually tried it.

So how did I accomplish this?

Well, let's have a look, first of all, at mindspace.

And literally what I did for this particular file was open movies, not just copy the entire content

pasted over top of what was in and mean, not just rename the class to admin, then I just scroll down

here.

The render function, change the title to manage catalogue and updated the link to go to slash admin

slash movie, slash the idea of the movie and saved it.

And that's all that I did that was very straightforward.

Then in edit movie dojos under the function named Handle's Submit in the last then clause of the fetch,

I just changed the ELT's to be this start.

Props that history push and take it to path name admin instead of displaying the alert.

And that's all that I did.

So that was a pretty straightforward challenge and hopefully it didn't give you too much difficulty.

So it's time to move on.
