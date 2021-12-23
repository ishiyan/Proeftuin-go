# 01 - Getting started with React

## 08 - How React Works

So before we actually start writing react applications, it might be useful just to have a brief overview of how React works and how it's different from a standard website.

So this is a version of what will be building in the coming lectures in this course.

And it looks just like a regular website.

It's not particularly pretty, but I'm not putting a lot of effort into making it attractive.

I'm more interested in functionality right now.

So we have a menu on the left in each of these appears to link to a different page, and that's exactly what the end user sees.

We have a login screen.

We have all the things you would expect to find here.

The difference is that if I view source on this and I'm going to use the view source functionality in my browser, you would expect to see go watch a movie, find a movie to watch tonight, an image and so forth.

But look at what we actually see.

This is it.

We have the body which begins here on line 18 and ends here.

And in between we have an empty div and a little no script warning for people to try to hit the site without JavaScript and a link to a bundle J.S. And then a couple of other Jass files as well.

And that's it.

So there appears to be actually no content on this, but if instead we inspect the page, so I'll inspect it right now.

You can see that we actually have a whole lot of HTML.

Now, where does this HTML come from?

Well, it's actually constructed by the JavaScript.

So in other words, if you try to hit this site without JavaScript, you're not going to see anything.

Now, fortunately, almost nobody does anything with JavaScript disabled.

So that's really not much of an issue for us.

But this is literally one webpage and that's why this is called a single page app.

So when you look at this, we actually have one component that is the entirety of the Web page, everything that you see right now.

And inside of that, we have a component for the menu on the left, a component for the content in the center, and it actually gets more complex than that.

So let me log in.

This form, for example, consists of an overall component and then it has the login component.

And each of these is actually another component.

So let me log in.

And let's manage our catalog and click on a movie now, this form is itself a component within a component, but this title input and this runtime input and this NPR rating, these are all individual reusable components.

And that's part of what makes Riak such a powerful and useful means of building a website.

You can create reusable components and use them as many times as you need to and actually adjust them or modify them for a specific use case.

But they're all literally the same component.

So all of the text inputs on this form are the same component with a few different properties pass to them.

The dropdown here for a rating, that's another component.

The text area is another component.

And all of these can be used wherever I want in my Web application.

So we need to change the way that we think about building Web pages, plural, because we're really building one Web page singular that consists of a variety of components that are customized and deployed to meet our needs of the moment.

So let's get started.

## 09 - How to use the downloadable code

So from this point on in the course, virtually every lecture will have downloadable source code and the source code will be the code as it exists at the end of that lecture.

Now, if the lecture has to do with react, then just download the zip file, unzip it somewhere, open it up in Visual Studio Code.

And I've done exactly that here for the Hello World lecture, which you'll see in a little bit in order to get it run.

You can actually look at the read me and it tells you this is hard to read this kind of code, but to run it, all you do is open your terminal window and type NPM install and wait for it to install all of the things that are referred to in that package.

JSON file and I will take a little while, but once that's installed, you can run that just the same way.

You'll you see me run the code in the lectures that are up and coming.

So that's how you do it.

If it's a JavaScript application or react application and if it's go code, you just run it the way you would normally go code.

But you have to run this one step NPM install for react applications first.

## 10 - Our first React app

So let's create our very first react application, and I'm on my desktop right now, and I'm going to create a new folder and you can put it wherever you want, but I'm going to stick one on my desktop and I'll call it react apps.

OK, now I'm going to open a terminal and you can use a command prompt on windows if you're on windows.

But I'm on a Mac, so I will use my favorite terminal program and I'm going to change to my desktop folder.

And of course, in Windows you can right click on the folder you just created and say open a command prompt here, but somehow get to your desktop or wherever you put that folder and do an Alice or different Windows.

And I'm going to go to my REACT apps and I'll clear the screen.

And now I want to create my very first REACT application.

And fortunately, it's pretty easy because we installed no G.S. and all of those things.

We can just type in X, and that is not a typographical error.

It is an X and we're going to call Create React app and then we'll give our app a name, which I will call test app.

And this will take a little while, particularly the very first time you install a react application and has to download a few things and create it and get all the necessary components.

So we'll just wait it out.

OK, so that took a little while, but now it is installed our first Riak application and you can see it gives us some instructions here.

We can run NPM start.

First of all, you have to go inside the directory we just created and we'll do that in a moment and you can run and start to start a development server.

And that's exactly what I want to do.

We can also do npm run, build or NPM test and all those things.

But right now I want to do an LS and you'll see we have a new directory here.

And again it's Dyax on Windows and I'll go into my test directory and I'll run npm start.

And what this will do is fire up the default application that was just automatically installed for us and open up in a web browser on PT..

Three thousand.

So let's do that.

And there it is on part three thousand, we have a sample react application running now, it's a cute little animation and everything, but it's nowhere near what we're going to want for our final project.

But it is a starting point.

All right.

In the next lecture, we'll open up this directory, the one we just created in Visual Studio Code, and we'll start playing around with it.

## 11 - The obligatory Hello world app

So we left our application like this last time, and if you actually quit out of it, don't worry,

I'll quit out of it now and you can quit out of it.

Make sure you're in the correct directory where your test app or whatever you name it lives and just run npm start.

And it will start up a Web browser with the actual application running once again.

So there it is.

So I have this running and right now I want to make some changes to it.

So what I'm going to do is switch to Visual Studio code, which I have running right here, and start a visual studio code.

And if you have a welcome screen, you can just close it.

And the first thing I'm going to do before I do anything else is go to the code menu and find preferences

in a similar place on Windows.

But you can find it in what you're looking for is extensions.

And I'm going to install an extension.

And then what I'm going to look for is called Riak.

Just type in, react.

You want this one S7 React redux, so on and so forth, install that.

It will just make things a little bit easier for us when we're done.

So once it's installed, this will change to uninstall.

So you can close this and click on this icon on the upper left hand corner and we're going to open the folder we just created.

And I put mine in my desktop, so I'll go to my desktop, which is here somewhere.

There it is, and I called it react apps, and I'm going to open test up and open this folder.

You'll see right away there's quite a bit of stuff in here, we have a source folder, we have a public folder.

Let's look inside the public folder to start with.

Inside of that, I have a lot of files, a robot text, which is kind of standard when you're deploying something to the Internet to tell search engines how to pay attention to your content, the fav icon is just an icon.

We can ignore that.

This is the one I'm interested in indexed HTML.

So I'm going to open that file and you'll see there's a whole bunch of comments here that tell you exactly what is where.

And some of the things might look a little different to you if you're coming from straight HTML, for example, here on line five percent public underscore Eurail percent.

Well, that, of course, is substituted for the actual path to your public.

You URL at some point.

So you can ignore that for right now.

But if you actually look at this file, if I was to remove the comments, so I'll remove these comments and these ones.

And these ones.

You'll see there's almost nothing in there you have in between the body tags, you have no script tag that tells people that don't have JavaScript enabled that they need to enable JavaScript to see anything.

But the only other thing you have is div with an ID of root and nothing inside of it.

And yet when I look at my web browser, I have all this content.

So where is this coming from?

Well, it's actually coming from and this is the other important folder, the source folder inside of here.

I have a whole lot of JavaScript files and CSV files.

I have an SVG file, which is an image file.

And of course that's the spitting image we see on our Web page.

But there's a lot of things that we can change here.

Now, what I'm going to do and this will seem strange to you, but I'm going to delete a bunch of stuff.

I'm going to delete set up testifiers, I'll delete that, delete the logo, I'll delete the app, test Jass or delete the app J.S. and I'll leave everything else.

I'll leave, delete the access to delete all of that.

OK, so I'm going to right.

Click on it and choose delete it will move those to the trash and they're gone.

And now if I switch back to my web browser you'll see I have an error.

Failed to compile and it failed to compile for a really simple reason.

So let's go back to our idea.

Let's open up this indexed jass file.

And the reason it failed is here at the very top.

This is a JavaScript file.

And if you're familiar with JavaScript, there'll be nothing new here to you because you'll understand how imports work and so forth.

But if you don't, don't worry.

All of this is causing us some grief because it's referring to things that no longer exist.

So what I'm going to do is just delete everything in it, OK?

And I'll put instead this JavaScript directive console log.

I am running just so I know that something's running.

Now, this, as you know, doesn't actually write anything to the Web browser.

So if we switch back to our Web browser right now, you'll see that it's a blank screen and it's a blank screen because we're not actually writing any HTML.

We have an empty body.

We have a body tag with an empty div in there.

But if I open the developer tools, so I'll open those and look at the console, you'll see I am running is here.

So now it's actually doing something of value.

Now clearly if you're taking a course and using react and go and this is your end product, you're a little disappointed.

So we probably want to do a little bit better than this.

Let's go back to our idea and let's put our very first react component in this file.

So I'll select out and delete and I'm going to type some stuff that might not make sense to you.

But don't be concerned.

We'll be going through this in detail as time goes on.

But right now, I want to write a really simple react component.

And in react, you'll be using components all the time.

So first of all, I'm going to import something into my JavaScript file.

I'll import React, which was installed when we created our application.

So I'll import react comma.

And then in curly braces I will put component because we're going to be using a react component.

So we have to have that and we're importing that from RACT.

OK, so there's some logic which is now available to our application.

And down here I will create a class, a JavaScript class and I'll call it app.

You can call it whatever you want, but call it app if you want to follow along.

And it extends the one we just imported called component.

OK, and then in curly braces inside of that, every react component must have one function called render.

It has to have that or else it's not going to work.

So I'm going to create a function called render and render, just return some content.

So what I'm going to return is and this is going to look odd if you've written JavaScript before, but I'm going to return it looks like HTML and it's almost HTML, but it's something else.

But let's just type right now and see how it works.

I'll type div and then I'll return to get to the next line and then I'll put an H1 tag and we'll do the obligatory hello world logic that you expect to have in any time you're learning a new programming

language or a new development environment.

And then I'll end it with the semicolon.

Technically you don't have to, but I do it all the time and I may as well confess right now that even though it's not necessary, as this course goes on, I will almost certainly always put the trailing semicolon and it hurts absolutely nothing.

So I can save this.

And that's great.

It seems like it's right, but I have this render function, this this this app component with a render function that returns something that looks suspiciously like HTML.

But at no point is it actually writing this to the browser window, even though it's rendering it.

We've got some content we can put to a browser.

We need to explicitly tell react, put this in the browser window in order to do that.

I have to import another thing.

I have to report react Dom.

Now, Dom stands for a document object model.

And if you've done any work in HTML and JavaScript, you're familiar with the document object model.

But React actually has its own document object model and will be going over that in detail as time goes on.

But right now I want to import react dumb from react dom, not surprisingly.

OK, and down here below my class, I'm going to call a method on that react dom react dom dot and it's render and I have to tell it two things.

First of all, what component do you want to render and where in the hell do you want to render it now?

My component is called app, but the syntax for calling that component is actually like this app and then a closing point in brackets.

So an opening point bracket app and then a closing point bracket.

That's my first argument.

And I want to put this.

If you look at our indexed jass or indexed on HTML, I want to put it right here in the div with the ID of route.

So back in my index digest, my second argument to react Dom Render is give me a reference to the ID route in my document and we use the standard JavaScript document, get element by ID in case matters in JavaScript.

So pay attention to how you type this and then in parentheses root and closed my parentheses and type my trailing semicolon and save this.

And when I go back to my web browser now it puts Hello World in there.

Now that is our first react component and it's not terribly exciting, but we'll be going through this syntax in a lot of detail in the coming lectures.

And you're going to become intimately familiar not just with components, but also with something known as the react lifecycle, which is pretty important to know.

One final thing that I'm going to draw your attention, even though this file indexed Jass has the extension JS technically speaking, from a react point of view, this is actually called a J. S X file and JS X File supports this syntax, which looks suspiciously like HTML in a lot of cases in this case, certainly is simple HTML, but there's a slightly different syntax required for certain things.

For example, attributes in an HTML tag.

We can't use class because class is actually a reserved word in JavaScript.

So instead we use class name with a lower case C and uppercase M, but we'll be going through that as time goes on.

Right now, we have successfully installed our first reactant application.

We've deleted a lot of things that were installed by default and we have created and deployed our very first react component.

All right.

We'll be building on this project as time goes on and let's get started in the next lecture.

## 12 - Working with components

So we've successfully rendered one component to the browser window, but of course, a react application can have multiple components.

So let's make another one.

And what I'm going to do is create a new component that just writes a footer to the page.

OK, and it's simple exercise.

We'll be modifying this as time goes on, but we need to get used to working with components.

So right in our source folder.

I'm going to create another file in here, so I'll create another file, new file and I will call it

App Footer App Footer Dojos.

And inside of that I'll import, react, react and I'll also import component.

From React.

And I'll create a class, but this time I'm going to use the key word export and I'll use the key word default class.

So I'm creating a class that automatically exports itself when it's called and I'll call it App Footer.

And as was the case with our other component, this one extends component.

Now, if you'll recall, every component in REACT must have a method or a function called rendered.

So let's create the render and will return some HTML or actually return some XML, some JSM.

And this time I'll surround everything with parentheses.

It just makes a little more attractive.

So what do I want to return here?

Well, I want to return a horizontal rule.

And you would think since this is HTML, I could just do this, but you'll notice how my ID automatically added a closing tag.

Well, if you're familiar with HTML, HRR actually doesn't require a closing tag.

It's one of the very few HTML directives that doesn't have a closing tag.

But we're not returning strict HTML.

We're returning XML and XML requires that all things have closing tags so I can fix it like that.

So now we're turning a horizontal rule and now return a paragraph that just says something like copyright and I'll put the copyright symbol and then the year and then some company, Acme Ltd..

OK, now my ID is helpfully telling me this is an error and this is again, it looks like valid HTML.

You're allowed to have the slash before the closing pointy bracket on our tag.

But why is this throwing an error.

This seems like valid html.

Well that is valid html but it's not valid XML and I need to surround this entire selection of HTML with some kind of tag and of course I could put a div around it, but that might actually break the presentation.

So I'm going to go back up to line one and I'm going to import something else from react and we're importing something called fragment.

And down here, inside my return, I'm going to have the opening tag fragment and then I'll move the closing tag down here to the end.

And this is a special thing available to us in REACT that actually makes us have valid XML.

It now has an opening and closing tag.

So it's valid XML, but the fragment itself will never be generated as HTML.

So it's just something created by react that allows us to write our return statement properly.

OK, so here we have a new class called App Footer that actually returns some HTML that we can include in our main page.

So let's go back to our indexed Jass file and let's import this file and we'll import it by saying import at footer from and this is not built into react.

So I need to specify the path name to it and I just use DOT for the current directory slash footer and give it a trailing semicolon and now I can put that anywhere I want inside of my render statement.

So again, I'm going to make this a little more attractive by putting parentheses here.

And a closing parentheses right here, and that looks a little bit better to have this over with.

I want to put this well, I can put it say right here.

Let's see if this works.

I'm going to import it by saying happe footer like that, and again, I have my error and my error is once again, I need to have one root level tag that surrounds everything I'm returning.

So in this case, I can just say give and take this closing div and move it down here and hit return and I'll just have everything over.

So it's a little more attractive and that should work fine.

Now, I don't believe I have my application running, so let me run it.

NPM start.

And I'll switch to my Web browser, and there it is.

So now I have both hello world and a horizontal rule and my copyright statement.

OK, one other thing we can change here.

Let me just hide this.

And go back to my app footbridges, I don't want to go back and modify my code every time the year changes, so let's instead make that year a variable.

So inside the render statement, let's just create a constant, constant current year and I'll make that equal to the new date.

Don't get full year.

So I have this variable.

How do I use that down here?

Well, the syntax for using a variable in the fragment or the HTML you're returning is very straightforward in curly braces, put in the name of the variable you want.

Now, if I save that, get rid of this, save that and switch back to my Web browser.

There it is.

Once again, it looks exactly the same.

But now that it's dynamically generated and will update every time the updates.

So we'll go back here.

And just to show you that it actually is changing, I'll put, say, 20, 20 dash and save and switch back.

And there it is modified.

All right.

So now we have on our index J.S. we actually have our hello world being rendered as part of their main return statement.

We're including a second component in this component.

And presumably you can see how we're going to start building things and react.

We'll have a library of components, reusable components that we can incorporate into any page we want to create in our application.

So this is a good start.

In the next lecture, we'll start styling things and make things just a little bit more attract.

## 13 - Styling components

So let's have some styling to our application, just some basic styling so we can see how success works in react.

So right now I'm looking at my index file and we have our import statements at the top and we have our class app, which extends component.

And that just returns this fragment of X amount, which actually is translated into HTML on the Web browser.

So what I'm going to do is add some style to this surrounding div and what I want to do is centered the content, make it a bit narrower and maybe change the font and things like that.

So what I'll do is refer to a CSA style that doesn't exist yet and you might think you do it like this class equals and maybe call it app and you would think that would work.

However, it probably won't.

And it's certainly not good practice and it's not good practice because remember, even though this looks like HTML, it's really XML.

And that word class is actually a reserved word in JavaScript, its class that we use right up here on line five.

So instead in REACT, we use class name like this and there are quite a few directives that actually have to be modified a little bit, particularly for attributes on HTML elements.

So I have this class name app, but that doesn't exist anywhere.

There's no access that defines it.

So you might think that you could go to index or HTML and just put a style in here at inline style or load an external stuff.

And that might work, but it's not the way it's typically done in in react.

So back here instead, you'll notice that we have a file called Index Jass and we have a file right next to it called Index Dot Com.

Let's open that up.

This is where we conventionally put the excess associated with a given Jass or jass file in react.

So I'm going to delete everything in here because none of this is being used.

And instead, the first thing I'll do is I'll import a Google font.

And I happen to have copied and pasted the font I want right from Google fonts, the import statement

I saw pasted in, and now I'll define my app class.

What I'm going to put in here are some really simple things just so we can see if it works or not.

So I'll say max dash width and I'll make it 80 percent of the viewport and I'll put in a margin of zero auto to center everything.

And finally I'll change the font font family and I'll make it Raboteau, the one I just imported a moment ago, and I'll save this.

So will that make any changes to our application?

Let's go to our Web browser and find out.

So I switch over to my Web browser and reload this and everything looks exactly the same.

And there's a simple reason for that.

Let me go back to my code.

Although I've defined my course as rules, I've not actually told my JavaScript file to use them.

And we do that with a simple import statement import and I'm going to import just dot slash.

It's at the current directory level and the name of the file is index counts.

And I'll say that now when I go back to my Web browser and reloaded, look how it's all changed.

So now I have some success rules that I can apply to the entire application page in this case, because

I'm applying it to the root level of my application to index, not jazz, but I can define success rules for pretty much anything.

So, for example, let's go back to our editor here and create a new file in the source directory called

App Futer Dot Com.

And this is the convention.

You have an app Footer Jass, and if you want to have CSA styling for it, you call the file that's associated with it the same thing, but ended with success.

So let's put some rules in here.

I'll create one for dot footer.

I'll just put some basic rules in here, text a line, maybe make that center and let's put in color, change it to dark gray.

And of course, you can use the standard success coloring syntax here.

And I'll make the font size a bit smaller, maybe zero point eight M's just to make it eighty percent of the standard size.

And let's go to App Footer Jazz and here on line eleven for this paragraph, I'll give that the class name.

Remember, it's not class, it's class name and I'll make it footer.

And again, we're not finished.

We still have to import that import dot, slash that footer, don't success.

And if I save that and switch back to my web browser now I've applied some styling to that footer.

OK, so this is a good start and it shows us how we can style various elements of our application.

But what if you want to use a framework like Bootstrap or foundation or one of the many popular success frameworks?

Well, in the next lesson, we'll see how we can include, say, bootstrap intraplate.

## 14 - Using a CSS Framework

So this time around, we want to use a success framework and I'm going to use bootstrap just because I'm really familiar with it and I like it a lot.

But before we do that, let's give ourselves some place to play with some content.

So I'm looking at my index, Jaspal, right now.

And what I'm going to do is create another component and this one is going to be called app content.

OK, so what I'll do, first of all, is in my source folder or create a new file, which I'll call app content Gorgias.

So let's import react and we'll import a component from react.

And give it a semicolon and now we'll export default class and we'll call it app content.

And it extends component and as always, we have to have a render function and it will just return and I'll use parentheses and close this off so I don't forget and we'll just return some HTML.

This is the content.

OK, so I've created that now.

Important in my index jass.

Import content from slash AB content.

And down here, between the or after the header, I'll simply put in my contact.

OK, so if we save this and go back to our Web browser and reload this, we have the content showing up there.

And this gives us a place to play with some bootstrap.

Now, there's a couple of ways of installing bootstrap.

I could if I wanted to go to my indexed HTML and just add the link to bootstrap right here in the head section.

So I would go back to my Web browser, open up a new tab and say go to get bootstrap dot com, click on getting started and copy the system right here and paste it into the head of my document and that would work.

But the way it's conventionally done in most applications that depend on NPM is simply to use NPM, npm, install, bootstrap and it will go and get it and install it for us and then we have it available to us.

OK, so now it's available to us, but we need to import it, so I'm going to go back to my index dojos and at the very top, just like I'm importing index stocks, I can also import and I'll put it right here, import and then the path name to the access for bootstrap.

And I'm importing it from bootstrap, slash, dist slash, CSFs, slash, bootstrap and success.

And if I save that and hide this and go back to SERAP content, let's put something in here.

OK, so let's add a button and I'll add it right here.

I'll put it in and remember, I have to close it like that because every tag has to have a close of some sort.

Since this is Ximo I'll just put a button it button class equals button button primary and I'll give it an empty aircraft just so we can make sure the bootstrap is actually working.

And some text my but.

And if I save that and switch back to my Web browser and get to the correct tab and reload this, there's my button.

So it doesn't do anything right now.

But I've actually added some styling using bootstrap, and that is the way that I always include bootstrap in my react applications.

Now, if you're more comfortable just putting it into the head of your HTML file.

Feel free to do so.

But this is the way that it's conventionally done.

So now we have several components and we can style them either using our own styles or we can style them using bootstrap or any access framework that you choose to include.

All right, let's move on to doing a few more interesting things with react.

## 15 - More about the CSS Framework

So just a little bit more about bootstrap and how to include it in your application right now, I've included just the success and you might have been asking yourself here, unindexed jass online sex, how did he know to put bootstraps this success and so on and so forth?

Well, when we ran NPM install, bootstrap, it actually installed bootstrap in this folder node modules.

And if I open that and scroll all the way down to the BS, you'll find a folder called Bootstrap right here.

And inside of that is a folder called Dist.

And inside of that is a folder called CSFs.

And that's all the various CSFs files used by Bootstrap.

And the one that I wanted was the minified version of the entire bootstrap framework, which is why I chose Bootstrap Dortmund's success.

You might also want to include the JavaScript, and the one I would include is this one bootstrap bundle mindgames.

And we can include that JavaScript exactly the same way.

Import from bootstrap dist slash J.S. slash bootstrap bundled up mean just like that.

And if I save that for some reason, my ID gives me a warning, which I'm going to ignore because I know it's not wrong.

If I save that and go over, say back to my web browser and open up bootstraps website again, get bootstrap and find in the documentation something that I think depends upon JavaScript.

So I'll go to components and maybe accordion and I'll find an example.

So I'll copy this example and go back to my REACT app and then go back to my code and my code is over here and let's open up our app content.

And inside of this I'll get rid of the button we created and I'll just delete that and paste in the code.

I just copy.

There's a lot of it and I'll get rid of it shortly.

I just want to make sure this works.

So I'll tab it over and save this and go back to my ID or go back to my Web browser and reload this page.

There's an accordion and there should not collapse.

And there it goes.

And if I view source on this page, actually, if I view source, you're not going to see very much.

We'll just see our ID route and this static bundle that if I inspect this page.

So inspect the page.

Down to the bottom, I have my static just bundle my vendor's main chunk and my main chunk, jazz, this is all automatically compiled from the various imports that we have in our REACT application.

So that's where the JavaScript goes.

Now, again, if you don't want to import it that way, lots of people and there are many schools of thought on this, but lots of people, I'm going to delete this code because we're not actually going to use it anymore.

Delete that if you choose to import things using a CDN, there's lots of positive arguments or lots of valid arguments on why that's a good way to do it.

I just tend to import it using NPM.

That's my preference.

You can do whatever you wish.

Either way is absolutely fine.

All right.

That's it for bootstrap and installing it.

Let's move on with our exploration of the things we can accomplish in RACT.

## 16 - Components and props

So this time around, I want to talk about components and props or properties, but the first thing I want to do is a little bit of cleanup.

So I'm looking at my index, Jaspal, right now.

And you can see we have an extraneous div.

I want to get rid of that because I don't need it anymore.

And I also want to create another component we have right now, our main app component, and that is the parent to an app component right here on line 17 and another app component called App Footer Online 19.

And what I'm going to do is create another one called Petyr.

So I'll go to my source folder and create a new file called App Header Dot Jazz.

And instead of that, I'll do what I've done before.

So we'll go to Apter and we'll create let's just copy and paste save a little bit of time.

So I'll copy all of this and go back to app header pasted in fix the name to app header.

We don't have access for it, so we'll get rid of that and I'll make this return.

Nothing more than an H1.

It's like this

H1 and what do we call it in the next hour.

Jess hello world.

So we'll use that for right now.

Hello world like that.

And I actually don't need the fragment tags because I'm not returning a fragment, which means I don't need to import it, which means that's pretty much ready to go.

So let's go back to indexed Jass and import Apter.

So here import app header from Donziger header right there.

And we use it right here like this.

Delete that app header and we'll format things a little bit better like this.

So that should all work.

Now if I go back to make sure that this is actually rendering properly, I'll switch to my browser and that looks great.

OK, so we have everything working the way that we want.

Now I want to talk about properties.

I said a while ago that a lot of what you do in react is to build reusable components.

And the problem is that my app header is actually got a title hardcoded in it and I want to fix that and I'm going to fix it by passing a property to the component.

And it's as simple as this.

What do I want to call the property equals?

And we'll call this my app, a different title just so we know that it's going to work.

So that is a property that is now available to header.

So let's go to app header.

How do I get access to that property?

First of all, I don't need this, so let's delete it.

How do I get access to that?

Well, we can actually do it like this.

We can just say get rid of Hello World and replace that with a variable inside of curly brackets, the same way we use with variables and we call this props dot, whatever I called it, and I called it title.

So if I save that, I've actually have access to this title variable as part of a props objects on the current item called this.

If I go back to my Web browser now and reload it, I have my app which was passed as a property to that component app header.

So if I go back to my code and go back to indexed jazz and call this instead cool app with an exclamation and save it and go back to my Web browser and reload this page, it automatically updates it.

So that's how we pass properties and we can pass as many as we want.

But it gets kind of awkward to pass things like title and then subject equals something and then whatever it may be, if I have twenty five things I want to pass there.

Of course this is a really inefficient way of passing lots of information.

So in the next lecture we'll look at how we can pass content in a slightly more intelligent fashion.

## 17 - More about props

OK, let's talk a little bit more about properties, and as I said last time, passing a whole bunch of attributes this way or props like this is really inefficient.

So instead, let's create a variable called cost and I'll just call it my promise.

Doesn't matter what we call it right now, because we're just learning how they work.

And this will be a standard JavaScript object.

And inside of this, I'll have keys and values.

So I want to pass title, so I'll give it the key of title and then a colon and then what the title is going to be.

I'll call it my cool app this time, then a comma and I can pass a second value or as many as I want.

So I'll just pass to four right now.

Subject my subject until this past three and we'll call it favorite color red, which is not true,

but it'll suffice for our purposes.

So I've created this object, my props and now I can pass that right in here like this and I'll take advantage of the spread operator that's built into later versions of JavaScript, dot, dot, dot.

And my name is my props like that.

So that passes it to header.

So if I go back to my app Heterogenous, I have this dot prop title.

Let's see if that works.

So my application is not running right now because I closed it.

So let me run it and start.

And I'll switch to my Web browser, and there it is, my cool app showing up exactly as it should.

And of course, I'm only taking advantage of one of the properties that I've passed.

But if I go back and look here, I actually have access to everything I've posted.

So indexed jazz passes my props with all of these three objects and I can actually go back and change the app header just instead of using that.

I'll use my favorite color and save that.

I think that's what I called it.

Make sure your favorite color.

Yes, I spelled one the British way and one the American way just to be egalitarian.

So let's go back and look at our Web browser.

And it changed it to Red, which was my favorite color.

So that's a slightly more efficient way of passing attributes from our index, not just specify all of your properties inside of a custom, constant or variable of some sort and pass it using the spread operator.

All right.

Let's move on in the next lecture.

## 18 - React Events

So this time around, we want to explore, react to events, and I'm looking at my application right now and as you can see, I've changed the application total back to my cool app and we'll make it something more useful before too long.

We're just exploring core reactor functionality right now.

And this time around, what I want to do is add a button right up here and have that button do something when I've clicked on it.

And just to get things wired up properly, what I'm going to do is add a button here.

And when I click on it initially, all it's going to do is write something to the console over here.

So let's go back to our code.

And I'm going to look right now at this this file app content.

Just so first of all, let's put a button in here and I'll just put a button in a standard bootstrap button.

Button class name equals button button primary.

And we'll put some text in here.

We'll call it fetch data because that's what it's going to do before we're done here.

And I need to actually have this button wired up to some logic in my react file.

So what I'll do is create another function here in my code.

So I'll call it fetch list, even though that's not what it's going to do right now.

That's what it will do before too long.

And I'll use the seven syntax to create this, which you'll see all the time.

So that's a function and all it's going to do is console the log.

I was clicked, OK.

So now how do we connect this function to this button?

Well, there's a number of ways of doing it.

And the simplest way you'll see this a lot is to use the unclick handler handler with an uppercase C that's going to be equal to this dot fetch list.

And I close it like that.

So now if I save that switch back to my browser, you'll see that it's reloaded.

So I click the console and there's our button.

And when I click on it, I get I was clicked over here and that's all there is to it.

So it's very straightforward.

Of course, I actually want to do more than merely write to the console.

So let's make this a little more complex.

I'm going to put a horizontal rule in here just to get some space.

And remember, you have to close the tag like that and I'll define below that and unordered list, which I'll call you Will and I'll give it an idea of post list.

Doesn't matter what I call it, as long as that has an ID.

So that's an empty URL.

And what I want to do now is instead of actually just writing, I was clicked on the console.

I'll delete that.

Let's have this function go and grab some data from a remote server in JSON format and then write that data to the unordered list, one for each entry.

OK, so there's actually a site where I can grab some default data and it's called JSON placeholder dot type code dot com.

So what I'll do is use the JavaScript fetch function to go to that URL, which is https Colon's Jason Placeholder dot type the code dot com posts and you can visit that site and it will explain how it works.

But all I'm doing is calling a public API that gives me some dummy data.

So we call then in our fetch function we get our response, which I'm going to convert to response start JSON and then Jason is past two and I'll open and close my quotes inside of that right now.

Let's just say console dot log Jason, just to see what it looks like.

OK, so I'll save this and I'll switch back to my Web browser, clear the console and I'll click the button.

There it is, so I got some data back and you can examine that data and it's just dummy data, as you can see.

All right, perfect.

So we've got that data.

We know that our fetch is working now instead of just logging it to the console.

Let's actually get a reference to our section down here, which I called post dash list.

So I'll say let posts equal document dot, get element by ID and it's called post dash list.

And now we'll just for each our Jason Jason Dot for each function.

And I'll just call the argument for object and I'll say elai equal document dot create element.

And since this is an unordered list, I want to create an ally element to try that again, Ally and

Ally will append to the child.

Document, not create text A..

And we want to put up the title and poster and child, the elai we just created.

So if we go back to our Web browser.

And we fetch the data again, there's all of the titles from our various data that we've pulled from a remote system, and of course, this is the sort of thing we'll be doing later on in the course when

we create the go back in to send data back to our REACT application in the form of Jason.

So our response we're getting here.

Has an object with a user ID and ID and a title, that's what we're paying attention to is the title.

So we're grabbing this from each row.

If you think of it that way, in our JSON response, paying attention only to the title, creating elai elements for each entry and writing it to the screen here.

OK, so now we have a way of giving some kind of interactivity to elements on our screen.

So in our REACT application we have a button.

That button has a function associated with it and we called it using the unclick attribute right here and just passed it.

This fetch list.

This is the current class fetch list is a function in the current class.

All right.

So that's a good start.

Let's move on in the next lecture.

## 19 - More events

So in the last lecture, you might have noticed that I use the Arrow Syntax online, five of app content digest when I define the function, and you might even asking yourself, why didn't you just do something

like this?

So I'm going to copy this so I can restore it.

Fetch list and that will work.

We can run the program and it will actually do exactly what it did before.

But the problem is, let's say we had another function in this, so I'll just call this another function and I'll define it the standard old style JavaScript way inside of this function fetch list.

I have no access to that function.

Another function I can't go, for example, just say this dot another function.

And up here I write console dialog, another function.

So if I do this and save my code and switch back to my Web browser and clear the screen and try fetch data.

Look at the this is undefined and it tells me on app content, dangerous line 14, if I go back here and look at app content, just line 14 right there, I have no access to this function.

But if I restore this to the way that it was before and save it and go back and clear the screen and now I click on it, I actually do call another function.

So the problem with defining your functions this way, like another function this way, is that you have no access to this keyword for the class, and that's going to be a bit of a problem.

So when you're writing functions in your react application, this is the correct syntax to use.

All right.

So we've used one event so far and that was the on click event.

And we use that right down here in line thirty two.

But there are many other functions we can use as well.

So let's put some other things in here just so that we can see how this works.

So I will before my HRR right up here, I'll just put another in seem to be really fond of the errors these days, but there you go.

I don't want to put another element in here and I'll just put a paragraph, so I'll put it in a P.

This is some text, OK, and on this, I'm going to put another handler and I'll put on this time on most enter and that will be equal to this.

Another function.

OK, that's all I'm going to put there.

So I'll come back up here to another function and I'll define this the correct way just because I can and I should get in the habit of doing this there.

So I now have this function called another function, which we're going to get rid of before too long.

This is just to demonstrate how event handlers work in react.

So I'll go back to my application and we have a new paragraph called This is some text.

So I'll clear the JavaScript console and I'll roll over that.

And when I roll over it, see how it called another function.

I can go back here and I'll create a third function called left paragraph.

And that's equal to using the pointer syntax console log, left the paragraph and save this and go back down here and say on mouse leave, on mouse leave equals this dot left paragraph.

And if I save that and go back to my Web browser and clear the JavaScript console and roll into it, if you watch the console, it says another function.

And when I leave it left a paragraph.

So there are many, many handlers you can use or event handlers you can use and react.

And I'll provide a complete list in the course resources for this lecture.

But this gives you some indication as to how you can add additional functionality to your REACT application.

It's relatively straightforward and there are different syntaxes for binding things and we'll be looking at those before too long.

But that's enough for this time around.

## 20 - Refs

So last time we managed to call a remote API, get some data and display it on the screen, and that's great.

And you would have noticed here an app content just right here on line 19.

I'm using something that is remarkably familiar to anyone who's worked in JavaScript for a while, and that is the document get element by ID and I'm getting the element post list by ID.

And while that works for this case and it's fine for the examples that I'm doing right now, typically in a react application, rarely will you see anyone using document element by ID or referring to an element in the DOM, by ID or by name or by class.

And there's a reason for that.

And that's because REACT has its own lifecycle and elements appear and disappear in the dom without you really knowing when that's happening because it all happens asynchronously.

So in this example where I have app content and it's my main component being displayed and it never disappears, that works just fine.

But you should be aware of something known as refs.

And this is Rick's version of Get Element ID, for example.

So let's just change this and not have an ID on post list here and I'll do it just for this lecture.

Now let's use a ref instead.

So refs work in component classes and the way that they work is like this right at the top.

I'm going to use the constructor method.

Constructor in react.

Constructor always takes the argument of props.

You can call it whatever you want, but typically it's called props.

And the very first thing you must do every time you call a constructor component is this super props.

Otherwise, nothing is going to work.

And the next thing I'll do is create a reference.

And what I'm going to create a reference to is my you help and you do it like this, this dot and I'll call mine ref because I'm referencing a list and that's going to be equal to react dot create ref.

So now that that's available to me in my code and of course I need to put that somewhere so way down here where I have ID equals post list, I'll leave that there even though I'm not going to use it because I will say ref equals and then in curly brackets this dahlstrom now I have a means of getting this dom element, this ul in react with relying upon ID and I'll talk about why using ID is a bit of a problem a bit later on.

So back up here.

The only change I need to make at this point is online.

Twenty four of my code posts is no longer a document get element by ID.

Instead I go like this const posts equals this Lystra that is not just this dot lystra.

I have to add this keyword effort dot current.

Nothing else changes.

So when I go back now to my web browser and click this button, it still works exactly as we would expect.

Now the reason you don't use ID in reactant is because the basic principle behind REACT and one of its great strengths is you have reusable components.

Now, right now, this component app content, Gorgias, only exist once.

But what if this component was instead, say, a text input for a form?

And I have a form I'm building and I use a text input component five times in that form.

If I had something like this ID on my text input and I use it five times, suddenly I have five elements on my form that I'll have the same ID and that's really not a good thing.

Of course, ID must be unique for the current Web page being displayed and if I have five elements with the same ID, well, bad things are going to happen and RAF's gives us a way of getting around this.

Now, the fact that you have access to the virtual dom to reacts virtual dom using refs and that they behave very much like IDs do in standard HTML and JavaScript, it might give you some encouragement to use them all over the place.

And in fact, if you look at the Riak documentation for refs, they say this.

Your first inclination may be to use refs to make things happen in your application.

If this is the case, take a moment and think more critically about where state should be owned in the component hierarchy that we haven't talked about state yet, but we're going to momentarily.

The point is, use these sparingly.

There's almost always a better way to manage the things that you want to take place in your application rather than using refs, but instead of using ID whenever possible, and it's almost always possible, use a reference instead.

So this is just a quick introduction to them.

I will tell you right now, I am not going to be using refs very much in this course at all because I don't want to encourage you to use a shortcut.

To try to make things happen because you overuse of refs will almost certainly result in problems in your application as it grows in complexity, and the easiest way to avoid problems is to not use refs unless it's absolutely critical that you do so.

All right.

So let's move on.

## 21 - Simplifying things with state

So we successfully got some data from a remote end point, and this is my application right now, and I'll just reload the page to make sure it's the latest version that I see.

I have an error.

I have an eight year where it doesn't belong.

We'll fix that in a moment.

But in any case, I can now retrieve some data and get a list.

So what I'd like to be able to do next is to have these as links so that I can click on them and do something, whatever it may be.

So how are we going to do that?

And first, let me see what this error is.

H.R. cannot appear as a dependent.

OK, that's easy enough.

Let's go fix that.

So back to our code and find this will just change this P to a do and that should make that go away because you can't have an R as a child of a P, so let's just save that, make sure it's working and

I'll clear the screen.

Reload this.

Perfect.

OK, so what I want to do now is to somehow have this list populated with all of the entries we're getting from our JSON and be a link.

And you can start to see right away that because we're not writing simple JavaScript, we can't just build a string with the appropriate unclick handler's the standard JavaScript unclick handlers, not the unclick with a capital C that we have and react because we're using Jass.

So it becomes a little complex and it's not immediately apparent how you might do this and fortunately react has something called state.

And when we use state, this becomes remarkably easy.

So to start with, let's come up here and just use just declare this or write this line of state equals and I'm going to in curly brackets say posts.

That's my key.

And that will be equal to that is just empty.

Now that is a word that you need to become intimately familiar with state and you've got to make sure you never declare your own variable called state.

You just use the one that manages state in your react application.

The fact that I've given this class that state and given it one key posts means that I can greatly simplify my code.

For example, when we click that button and this fetch list function is called this last part of the chain, the last, then that becomes a lot simpler.

So what I can do is get rid of these two things because I'm not actually using them and I no longer need this this request to post or this this reference to posts, post lists.

I can get rid of that.

And I no longer need to go through this entire foreach in JSON.

All I really need to put in here is this.

This state, which is a function available to us from react, and I'm going to set the key of posts to be set to chase them, that's all I need to do.

And that's a lot less code.

I should put a semicolon there.

Now, I can come down to my post part here.

This doesn't even need the ID anymore so I can get rid of it.

It's just cluttering up my code.

And inside of that I'm going to write a simple Riak directive.

And it goes like this in curly brackets, this dot state hosts.

Which I spell right posts and I'm going to call map, and of course, Map is familiar to you if you've worked with JavaScript at all, and I'm going to put a variable C and point that two in parentheses and down on the next line here, I'll just write some just below and I'm going to put this attribute on it and I'll explain why in a moment.

Key equals and I'll just make that equal to see ID, which is straight out of our adjacent file, and then I'll put an atrip equals hash bang, which means it doesn't go anywhere.

And I'll put an unclick handler here on the click is equal to in parentheses using our arrow syntax this start and I'll just say clicked item and handed an argument to it id ok and we'll make sure that function is done correctly at the moment.

Then I'll close my opening a tag and here I'll just write the total title.

OK, let's make sure that our clicked item function actually takes an argument so I don't have a click function.

Let's create one so clicked function or clicked item.

I guess it was clicked item.

Is equal to and this will take one argument or arroz Syntex console log clicked and I'll pass the X as a second parameter to that right into the console.

OK, so let's go through this.

There's a couple of things that are new here to start with.

We have our state and all I've done is to find one thing in my state object and I can have as many keys as I want in there.

Then I created a new function which takes one argument ex and his rights to the console clicked, plus whatever Xs and down here in the US, which a little while ago I had an idea of post list or something like that.

I'm ranging over or iterating over posts and we're using the dot map which is built into JavaScript.

Each iteration of the map will create or populate the variable C with the current value of that iteration.

Then I have some standard X, which we're already familiar with my unclick handlers.

A little bit different, though, because I'm actually using an arrow function and empty parentheses arrow and then this dot clicked item because I want to pass a parameter and I have to do that.

I don't have access to this unless I use the Arrow syntax and I just write the title in between the atrip.

So let's try that and see what happens.

So back to my Web browser and I will clear the console to make sure there's no errors in there.

Reload this.

So far, so good.

Let's click fetch data and see what happens.

And there is my list.

Now, I should be able to click on one of these and get something written to the console from my new clicked item function I just defined.

And there it is, clicked one.

And this should be a different one.

Click three.

Now one other thing to note.

Back in our code in this little Jass X, where I'm going through the posts using my map function, I have key equals and then a unique ID.

And this is something that react really expects you to have there.

For example, if I take this out and save this and go back to my Web browser and clear the console and then fetch data, we'll get a warning.

And the warning is each child in a list should have a unique key property so that it gives you a link to it.

OK, it's still worked, but you're going to have problems unless you leave this right in there.

So the line, each entry in that range through the posts needs to have an attribute key somewhere.

And that key needs to be unique for each entry.

And of course, I have an ID coming in my JSON file and I know that's unique, so I can use that.

All right.

So that is how we manage this.

Now, we'll just do one more thing here just to show you that you can use state anywhere before my ule.

I'm going to put a paragraph together and I'm just going to say in drawing from my state, this DOT state that posts the length which again is available to us just because of JavaScript and I'll say items long.

And I'll just put the word in here, so I'll say posts is there and I'll save that.

And when I go back, should say posts is zero items long.

But when I fetch data because I'm using state, it won't just update the URL with a bunch of Ellyse.

It'll also update this line of text.

And now it says Posts is one hundred items along and it took care of this.

And that is what makes state so powerful and so convenient.

Enwrapped.

All right, let's move on and play with some more interesting things.

## 22 - More about state - lifting state to share data between components

So things are looking pretty good so far.

We have an application with three components, one for the header, one for the content and one for the actual footer down here.

And we implemented state last time around and we did it in our code under app content.

So we declared a variable here called State and we gave it one member posts and we access it by going this state posts.

And if we want to change the state, we can never just do something like this.

Where do we do this?

So we can never say, say here, this dot state dot posts equals JSON.

That won't work.

If I comment this line out and come back to my web browser.

Looks good so far, but if I choose fetch data, I'll actually get a problem.

And the problem is it's not going to work.

It's not going to work just because you can't set a state variable that way.

You must use that method as we had before.

So restore this like it was using state and now if we go back, it will work as we expect.

There it is.

OK, so we have that state, but that state exists only within the component where it's declared.

So right here we have it in our app content component.

So if I go over to app header and say, well, let's put another line in here, so I'll wrap this in a fragment because it has to have a root level element being XML and move this down here and then give it a paragraph underneath the heading that says P length of posts is this dot state dot posts, dot length.

I actually have to get an error and you'll see it pretty quickly.

Let's go back and look and we get that lengthy error and I'll just reloaded so we can see the error a little more clearly right up here.

Length of post is this, but the error, of course, is that this dot state is no, that's because each component manages its own state.

So let's go back and get rid of the offending codes.

So at least run get rid of this fragment and this line and this fragment and we no longer need to import fragment.

OK, so now it runs again.

So how am I going to share state between components?

Well, that's accomplished, in fact a number of different ways.

But one of the most common ways is something called lifting the state of a component.

So let's do that.

Let's go back to our code and let's go to Index DOT just because what we're going to do is lift the state to the nearest common ancestor of the components that I want to share it with.

And of course, we only have three components and their common ancestor is index.

So let's see how we can lift the state.

So the first thing we'll do in our index JS file is right after the class app extends component line.

I'm going to put a new function in here and it's just a constructor function.

So it takes one argument props and it does very simple things.

First of all, we'll call super to get our super constructor from the default component.

Then we're going to say on this for this we're going to call a function called Handal host change.

Because all we want to change right now here right now is the postdated.

And that will be equal to this dot handle, post change, dot bind this OK, which looks a little odd, but if you worked at all with object oriented programming, you're familiar with this.

We're binding one function to another and finally will say this dot state equals and we'll just give posts an empty value.

Now, of course, that Handal post change function I'm referring to on line 15 has to exist.

So let's create it handled post change and it will take one argument posts.

And we're just going to say this dot set state and we'll give it posts is equal to the argument we've passed to this function posts.

OK, so that's all that's necessary right now for the constructor bid.

Let's go down to the part where we're actually rendering our components right here.

So I better I want to share posts with app header, so I need to pass it.

Two more attributes.

Posts equals to this DOT state that posts we're just passing it property and handle post change, which is the function we just.

To find this equal to this dot handle post change, so that's to share it with the app header and of course, we also want to share it with the content.

But all we have to do here is call Pendle post change is equal to this dot handled post change.

So we're binding those two functions together.

So that's all the changes necessary in index dot.

Let's do the header next.

So our header here, we're again going to give it a constructor constructor and it will take one argument which I'll call prompts and again, super props and this dot handle post change, even though we're never changing the function in here, it's good practice and it's actually necessary to have this equals this dot handle, post change, dot bind and this.

And of course, we have to have the function handle host change.

And it takes an argument which I'll call posts and we're going to say this dot props, dot handle post change props.

Try that again, posts.

OK, so there's the function defined that is necessary because of this this call to it on line seven.

Now, down here in the return, let's once again wrap this in a fragment fragment and I'll take the closing tag and move it where it belongs.

And let's put our paragraph back it up.

There are this props, the post length.

Entries in posts.

OK, so the next place we want to change is, of course, is an app contest and there's a few changes required here as well.

So let's go back to app content.

And first of all, let's clean it up a little bit.

We don't actually need this paragraph four on most enter and so forth.

So let's get rid of that, which means we don't actually need the left paragraph or another function so we can delete those just to clean things up a little bit.

OK, now, again, we're going to have to have a constructor constructor and I'll give it the argument of props.

And again, super props to call the parent constructor.

This dot handle post change will be equal to just as was the case in the last one.

This dot handled post changed of mind

and this and it's all we need in the constructor.

So once again, let's create that function, handle post change.

Now here we are actually changing this.

So posts and we'll say this props the handle post change and hand it the argument posts.

So what other changes do we need to make here?

Well, we actually have to do one down here after we fetch the list, when we get our new value from

our call to Jason Placeholder, that type of code dot com slash posts, the value we're getting back is Jason.

And we're setting the state locally here.

But we also now need to call that function, this dot handle post change, and we're handing it the

value that we got from our call to the remote API.

So once we've done that, is there anything else we need to do?

Well, first of all, let's go find out and see if this worked.

So back to our Web browser and we'll reload the scissors in here.

But that might go away.

It did.

So we have there are zero entries and posts up here that's in our header.

And we have post to zero items along here in the app content.

And now when we fetch data, this change to one hundred and that changed to one hundred.

So that's how you lift state.

It's actually not that difficult.

Once you've done it a couple of times, it'll begin to make sense to you effectively, your binding one function to another and lifting the state to the nearest common ancestor.

And we only have three components, all with one shared ancestor index dojos.

So we lift it up to the app level or the top level of our application and make sure that everything stays in sync.

## 23 - Functional components

So far, everything that we have in our project, you might have noticed, is actually a component.

And as I've said before, one of its greatest strengths is that it allows us to create reusable components.

So at the very top level in Index Digest's, our class app extends component, which is, of course, from the REACT package.

In the same way, the header is a component, the footer is a component and our content is a component.

But it's also possible to create functional components and react.

And those are plain JavaScript functions.

So let's do that.

What I'm going to do to start with is go back to Index Dot J.S. and I'm going to comment about the footer here, which means, of course, that down here I have to comment better because it doesn't exist and that works fine.

So if we go back and look at our Web browser, it should be the same, except that we have no footer and we don't.

There it is.

So what we're going to do now is create a functional component that takes the place of app footer.

So let's do that.

So in my source folder or create a new file.

And just to make it clear what it is, I'll call it App Footer, Functional Component Doorjambs.

And inside of that, I'm going to, first of all, create a function and I'm going to give a few blank lines because we have to do a few more things before this will work.

I'll say export default function and I'll call it app footer functional component and I'm going to take the argument props even though we won't use them.

But I'll talk about that in a moment.

And inside of that, I'll do exactly what I did in my app Footer.

So let's go to app for the and find the return state and copy that.

Actually get the constant to and copy that and go back to my new app for the functional component and just paste that right in there and notice right away.

We have a few things missing to start with.

I'm returning JSM, which means I have to if I want this to do anything, I have to import, react from, react like that.

And because I'm using fragment, I also want to import fragment.

So let's put that in there.

Fragment now.

I have no errors.

OK, so I've created this functional component.

How do I use it?

Well, one thing to note right away, because this is a functional component without the use of something like react hooks, which will look at in the last section of this course, we have no access to state.

We have no access to the this keyword.

It's really not nearly as intelligent as a true react component, but there are lots of situations where you just want a basic dume component and this is exactly how you do it.

You declare a function, you import react and whatever else you need from the react package and return what you need to return.

So let's format this and let's go back to our indexed OJARS right here and scroll to the top and let's import the new functional component.

We just create import and I'll call it type footer functional component from dot slash output or functional component right there.

And now that it's important, I can come down here in my return statement where I'm rendering Jass X and just say footer functional component, give it its self-closing tag and this will be almost the same as what we had before.

So let's go back to a Web browser and reload this and there it is.

Now, what's missing, of course, is the success and that seems a little unusual, but it's not because we need to go back to our functional component right here and just important success.

So we'll import and use the same CSS file agap footer.

Success, even though it has a different name, it will still work.

So let's go back here and relive that and there it is.

And it is styled exactly as it was before.

But now, of course, this is a dumb component, which for a footer is fine.

It doesn't need to have stage in our case.

It doesn't need to have any of the things that you would get from a standard react component.

So there's no reason not to use a functional component in that case.

Now, I mentioned that I'm passing props here.

And of course, this is something that you will probably use.

You want to be able to pass properties to your functional components in a lot of cases.

So if I go back to my index, such as add a property here, my property equals Hello world.

Then I have access to this in my functional component so I can come here and I can't say this stuff props because the reason of this keyword.

So I just say props dot and paste in my property.

And now when I say this, because I called it props, here is the argument to my functional component and I referred to it as props here.

And this matches the property that was passed to this unindexed.

J.S. When I go back to my Web browser, reload this.

There it is.

Hello World.

And that's all there is.

So functional components.

Use them all the time.

When we get to the react router part of this course, we'll be using functional components as Stobbs to start with.

So the key things to remember, if you're going to be using JSW and you almost certainly are, you must import react if you're going to use fragments or anything else.

That's part of the react package.

You need to import that as well and always pass this props argument, particularly if you're going to be using properties.

Otherwise you'll have an error.

I could take out this property and take that argument to the function declaration and it would work fine.

But you should get in the habit of passing props even if you don't use them, because you might at some point in the future, I'm going to revert the source code to the state that it was at the start of this lecture.

But this version with the functional components is available as part of the course resources for this lesson.

## 24 - Cleaning things up

So I just want to do a bit of application cleanup right now, and this is nothing major, but I like to keep things more or less organized so I know where to look for things.

And we lift our application like this.

And I'm looking at app content.

And here I don't think we actually need to declare this state variable on line 14 at all.

As a matter of fact, I want the state to be part of my app component in it next J.S. And that just makes sense because that's the central location.

It's the parent component that's managing, sharing state and information between the various components.

So I'm going to delete this line entirely and I'm going to assume that I'm going to get my necessary information as a prop., which means I don't even need to set state at all.

I can delete this line in my list entirely because that's all going to be handled at the app component level, because on line 19, every time that Jason changes and every time I call the fetch list function, I'm actually binding the handle post change function in this class with the one in the app class.

So everything is being managed automatically so I can delete that entirely.

The only other changes in this file will be down here anywhere.

I'm calling this state doorposts.

It should be this prop dot dot posts like that.

So here and here on line forty, this dot crops dot posts, which means back in index dot j.s..

The only real change I need to make, I don't need to declare the state variable because that's handled by my constructor on line 16.

I just need to do the same thing for app content as I do for app header and that is to pass this prop.

So I'll copy and paste.

And if I see this now and go back to my web browser and reload the page, everything should still work.

And it does.

All right.

That's just a little bit of cleanup just to keep things as organized as possible.

Let's move on.
