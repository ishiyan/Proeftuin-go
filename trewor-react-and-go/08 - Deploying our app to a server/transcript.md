# 08 - Deploying our app to a server

## 099 - Getting the React application ready for deployment

So in this section of the course, I want to talk about deploying our product to a production server,

and right away you realize that there are two things we need to deploy.

We need to deploy our react front end, which will take care of in this lecture in the next couple.

And we also need to deploy our backend, which is written and go.

So let's get started with the front end.

And the very first thing you will encounter when you try to put this to a production server is the kind

of thing you see here on movies online, 12 in the component did malfunction and that is this you URL

because this euro will work in development.

But clearly it's not going to work in production and we need to somehow modify this.

So it'll work on our remote server.

Now we could go in and manually update the link to what it's going to be in production for every request

that we have in our app.

But then, of course, we'd have to maintain two versions of our code, one for development and one

for production.

So the easiest way to do this is to use environment variables.

And I'm going to take you to a Web page right here.

This is uncreate dash react dash app Dev, which is a great site.

And under the environment variable section, it tells you how this works right here.

In order to have something from the environment available to our application, we create an environment

variable that begins with react, underscore Apte, underscore all uppercase letters.

So let's do that.

So I'm going to go to my terminal and I'll simply on a Mac or Linux type export, react, underscore,

app, underscore and then whatever I want to call it, I'll just call my own API underscore Yoro and

I'll make that equal to and for right now because I'm in my development environment, I'll just make

it equal to HDP Colon's localhost colon.

Four thousand like that.

Now I can use that environment variable right in all of my requests in my code.

So here, for example, what I'll do is replace this.

With and you'll see this right on that page I just was showing you, I'll delete that and I'll put a

dollar sign process dot on dot and then the name that I just chose, which was react, underscore and

underscore API, underscore your URL, and then to make this actually do the appropriate substitution,

I'll change these double quotes to tactics.

And now it's a JavaScript template.

So if my application is running and it's not so let's start at.

And there it is, this link should still show me the list of movies, and it does so that seems to work

really well.

Now, obviously, you're probably going to want to put that environment variable in the appropriate

script in your startup folder or in the slash profile file.

So you don't have to reinitialize it every time you start development.

But that's pretty straightforward and there are lots of resources online that show you how to do that.

And chances are, if you're taking this course, you probably already know how to do that.

Also, if you're on a Windows machine, you have to do it slightly differently.

And this link right here, let me show it to you.

And this is only course resources for this lecture.

This shows you how to set up environment variables on windows, and it's pretty straightforward.

OK, so the first thing you have to do, of course, is to make this kind of substitution for every

request to the URL that you have in your application.

And there's no point in you watching me do that.

But if you download the source for this lecture, you will find all of the updated requests in that

code.

So let's move on to the next step.

## 100 - Building the production ready React application

So I've made all the necessary changes in my source code to use process, sort of don't react to underscore

app and so on and so forth.

So those have all been done and now I'm ready to build my application for production.

Now, one thing to bear in mind, even if on my remote server I set up an environment variable named

React, underscore, app, underscore API, underscore URL and have it point to what my production API

is going to be if I just build my application now and deploy it.

That's not going to work.

These environment variables are actually substituted at build time.

So if I build it right now and copy the source up to my remote server and try to run it, it's still

going to think that it's set to whatever my current environment variable is set to.

So instead, before I do the step where I actually build this, what I'm going to do is clear the screen

here, my terminal.

And I'm going to export, react, underscore up, underscore, API, underscore you are all to be equal

to in single quotes and I'm going to put my production, you are always going to be HDB Colon's, and

in my case I'm using api dot lern dash code dossie.

So I'm going to export that variable now and now from my root level of my source front end.

Now I can actually build this npm run build.

And we'll take a little while, but what that does is create a new folder at the root level of my application

called Build, and inside of that will be production ready code.

So it's optimized for production and it's done.

So now you can see over here in my explorer that I have a new folder called Build and there are all

of the files and these are much smaller and optimized for speed and production use.

So what you need to do next is to compress this build directory, tear it up or zip it up or zip it

up and remember where you put it, because we're going to be copying that up to the production server

before too long.

And of course, now I can re-export export, react actually, and probably just use it up here.

OK, there is change that back to localhost.

Four thousand just so I can continue to work on this if I need to make changes and is to restore things

as they were prior to the most recent change to environment variables.

All right, let's move on.

## 101 - Getting the Go project ready for deployment

So just like we had to make some changes to our react front end before deploying to production, we

also probably should make a change to our back end.

And I'm looking at the back end source code.

I'm looking at main go.

And if you recall, we actually encoded right into our source, our JWT secret, and that's not a great

idea.

You really don't want to put secrets like this into your source code, particularly if you're pushing

it to a remote repository.

Now, if you're if you've already done so, all you have to do is generate a new secret and you'll be

fine.

But as I said when we first did this, this is probably not the best approach.

So I'm just going to comment that line out for right now.

And instead of reading that from a flag, a command line flag, what I'm going to do instead is right

here.

I will read JWT secret from the environment.

And it's really simple to do that and go.

I'm just going to say cfg dot jwt dot secret is equal to and I'll use goes built in ability to read

things from the environment.

OS dot get on and I'll let this right there.

I'll let my ID do the import for me and I'm going to look for an environment variable which I'll call

go underscore movies, underscore JWT like that, and that doesn't exist yet.

So I'll open my terminal window and scroll up so I can see that.

Copy this and I'll simply say export go underscore and whether I call it movies underscore JWT equals

and then a single quote pasted in there, close a single quote and have return again to make my life

easier as a developer.

I'll add that to the appropriate startup script in my home directory, and that's simple enough to do.

But in any case, that's there now.

So now I can delete this comment that a line and run this just to make sure it works though.

Runs cmd or dot slash cmd slash API, get my back end up and running.

Perfect.

Now I'll switch to my front end code which I have right here and it's not running there.

So let me clear the screen and run npm start and away from my browser to fire up.

And now I should be able to log in and if I can't and I did something wrong, so me out here, dot com

password.

And it worked fine.

Now let's go to manage, catalog and make a change to The Godfather just to make sure I'll just add

a few dots after this.

And if I can save these and go look at it and see those changes, then everything works as expected.

Perfect.

So the next step is to compile a version of our go back end to a single binary for the appropriate operating

system, and we'll do that in the next election.

## 102 - Building the Go back end for our remote server

So now that we've made the necessary changes to the source for our back end, let's get the application

ready for distribution to a leading server.

So I'm in my back in source code folder and you can see that I have the CMD directory, the go on file,

so on and so forth right there.

And I want to compile this for distribution to a Linux server and I want to assume it's Linux.

If you're going for a different platform, it's relatively straightforward and the process is pretty

much the same.

So let's compile this, not for the Mac that I'm on right now, but for Linux and Engo.

That's really straightforward.

Make sure you're at the root level of your application source code.

And I am right here, so I'm in the correct directory and type on and then go os x equals Linux, go

arch equals AMD sixty four and that's suitable for distribution on Ubuntu for example, on Digital Ocean

or Lynard or any of the popular services.

Then we want to go build Duchow and we're going to build out to a coal mine go movies like that.

That's the name of the binary.

And what do we want to compile ACMD API hit return and it will compile for that platform.

And if I do, unless there is the file right there go movies, a single binary.

And that's one of the great things about Go is everything compiles to a single binary.

All right, so that's ready.

Let's move on.

## 103 - Copying files to the server

So let's copy the necessary files to a remote server and get started making our application live, so

I'm looking at the source code for my front end.

And what I've done already is I've taken this folder build and I've zipped it.

So zip that up and put it in a directory and copy the binary we made in the last lecture, which I call

good movies, copy those into the same folder.

And when you do that, open up a terminal and go to the folder where those files exist.

So I'm in one call to transfer and you can see I have two files there, Bill Dot Zip, which is my react

front end all zipped up from the build folder and go movies, which is the binary we made in the last

election.

Now I want to copy those to a remote server and there's lots of ways of doing this.

You could build up a workflow that automatically updates and get repository and does that build on your

server.

But I'm going to keep things really simple right now.

So I'm going to copy these two files to my remote server.

And mine is called it lern dash code gutsier.

So I'll type Sepi for secure copy and want to copy, build the zip to TC's at the name of my server,

which happens to be learned dash code here.

And of course you might be using an IP address.

You're certainly not copying it to learn Dasch code gutsier.

I'm going to put that in my home directory home TC's so that will copy that file up.

There it is and I will do the same thing but instead of build zip I'm copying up go movies.

Go to movies and that will copy that up.

So now if I go to my remote server, SSA takes at learned code and clear the screen and do an LS.

There are the two files.

So I have my bill zip right here and I have good movies right there and I can actually try to start

go movies.

It won't run because it doesn't have a database set up, at least try to run and there it goes.

So it did try to run.

So the application is now functional.

So the next step is to copy up the database, put it into Postgres on the remote server, and I'm going

to assume you have that running and if not, you'll have to install it.

And that's pretty straightforward.

And we'll get this application running and we'll take care of that in the next lecture or two.

## 104 - Setting up the production database

So before we can make our application live, we need to set up the database on the remote server, and

I want to assume for the sake of simplicity, that your postgrads server is on the same machine as the

actual backend application.

So I'm in my two transfer folder right now, and I still have those two files that I just copied up

and now I want to dump the database.

So I'm going to type PJI dump to dump my existing database and I'm going to dump it with no owner like

that, no owner there.

And I want to dump go underscore movies, which is the name of my database on my local machine to GM

DOT sequel.

And if anyone else, there it is.

Now, let's copy that out the same way we did before GM got the sequel to your username at your host.

And I want to put it on my home directory.

So there it's all copied over, and if I switch to another tab where I'm already logged into that machine,

you'll see that I have my Gmod sequel right there.

OK, now I need to set up the database on the remote machine.

I don't want to use postcode to do that, but you can use whatever format you want, so I'll connect

to my remote server.

I'm using an S.H. connection to get to it and I'll go to the top level and create a new database, which

I'll call go underscore movies.

OK, so that's created.

Now let me go back to my terminal and I need to actually dump that database or take that database dump

and populate the new empty database I just created.

And that's as simple as pseudo Dirceu and my user name is Postgrads and I want to go Acuil dush d these

go movies for the database I just created and the file I want to dump in there or move into there is

GM dot sequel and type in my pseudo password and it's done.

So now if I go back to my post to go and look at Google movies, I should have some files in there and

I do and there's everything copied over.

OK, so now we can actually try and see.

Let me clear the screen.

So let's run go movies with the Dash H flag to see what our command line parameters are supposed to

be.

So DSN is the one we're interested in.

We want to try to start this and just make sure it can connect it to the database.

So let's try that dot slash go movies and I'm going to use the flag DSN and that's going to be equal

to and I'll put it in double quotes so my host will be equal to localhost.

My port will be equal to five, four, three, two.

My user will be user will be equal to Postgres.

My password is the singularly unsecure.

Very secret.

And the DB name will be equal to go under score movies and my SSL mode will be equal to the same.

Let's see if that works and there it is.

It's actually connected.

And you'll notice I used a different syntax for the connection string, and that's because I'm using

Lempicka and it accepts a couple of different formats for doing that.

But our server is now running and now we're ready to set it up using supervisor and then configure a

Web server for both our front end and our back end.

And we'll take care of that in the next couple of lectures.

## 105 - Setting up the web server

So now we have our files up on the server and it's time to set up our Web server and get everything

working in production.

So the Web server I'm going to be using for this lecture is Catti, and it's a very powerful Web server.

It's written and go.

It's extremely fast.

It's dead simple to configure and don't feel compelled to use.

Cadie, if you like Engine X or Apache or something else.

Feel free to use that.

It's very, very simple, similar to what I'm doing right now.

And I'll post sample configuration files for the most popular Web servers to this lecture.

But I'm going to be using Cadie because it works extremely well.

And one of the things I really like about it is when I set up a virtual host on Cati, it automatically

goes and gets let's encrypt security certificates for me and renews them as necessary.

And I don't have to do a darn thing and I really like that.

So this is the website for it.

There's lots of documentation, but let's get started setting up our application.

So in my terminal I'm logged on to my remote host.

And if I do, unless you see, I have my.

So I'm already logged on to my remote hosting my terminal, and if I do, unless there's my build zip

and there is my go movies file, and we've already, of course, imported our skill into our database.

So let's unzip, build, zip.

And it's as simple as unzip, build, zip.

Now, you might get an error if you're on Ubuntu, for example, saying you don't have one zip installed,

but the single command to install it is right there in the error message.

Just Apte get on zip.

So that creates a new folder in my home directory called Build.

Now I'm going to rename that, I'm going to move it to w w w I'll move build two w w w dot learn dash

code gutsier and that's the domain name I'll be using to serve my content up.

So that moves it.

Now I'm going to move this entire directory to the standard web server directory on a buncha pseudo

move.

We don't learn Dasch code dossie to via w w w hopefully that's in there now and there it is.

Good.

So that gets me my source code for the front end in the correct location.

Now let's go to our web and pseudo make idir api learn dash code dossie.

So I've made a directory for learn code for API.

I learned how code.

Okay, let's go into it.

Let's go into it again and now I'll copy from my home directory, my binary from my back and go.

Movies to hear, and I got to put a pseudo c.p.

Go movies to hear that's better.

OK, so if I go up one level, I should now have two directories.

I had some there before, but ignore those ones.

So I have one for my API API learned Ashkenazic and one for W-W.

We don't learn dash codes yet.

Now I need to tell Cadie how these files.

Now, we need you now I need to tell Katie.

About those two directories and those two virtual hosts, so let's go to where Katie is installed and

that's in etc., Katie, that's where its configuration files live.

Inside of that, you'll see that I have a couple of files there.

One is called Cadie File.

That's the live one that configures Cadie.

The other one is called Cadie File Dist.

And if you just installed Cadie, your Cadie file will look like this.

Try that again.

It looks like this, and I just moved the default caddy file to a back up called caddy file artist,

and what I've done is I've modified The Cafferty File to look like this, this one.

So it consists of four sections.

The very top is just an email.

And that's the email that's used to when you go to get an let's go and caddy goes and gets a let's encrypt

SSL certificate for you.

It just put your own email address in there.

Then I have to define one called static.

And that's just something I can use by referring to the key word static in my configuration files for

individual hosts.

So there's one for static files and there's one for basic security.

And this is standard stuff for anybody that's ever run a Web server before.

You don't want to reload JPEG or GIF files or CGS.

You want those to be cached for a long time.

And security just enables Høst to ensure that people never hit an unencrypted version of your site and

set some no sniff headers.

And finally, a referrer policy.

And then the last line import config D Start tells Cadie, go look inside the directory in the current

directory and look for any file that ends in dot com and a little dot as well.

And of course those are our virtual hosts.

So I already have a config directory.

You probably need to create it.

So pseudo McDeere config D and then go inside of it.

And what we're going to do is set up, first of all, one virtual host for our front end and a virtual

host for our back end.

So to do that pseudo Vij and I'll call mine w w w learn dash code for the front end.

And instead of that, the very first line, you tell Keddie what domain names or what IP addresses do

you want to look for, for this virtual host.

And I want w w w I think that's lern dash code Dossie and Espace learned dash code gutsier.

Both of those should resolve to whatever is inside these curly brackets.

Very simple to configure something in Cadie first line in Code Z, CTD and Jizan.

So those are two compression algorithms, one which is new and almost unsupported but will be eventually

that's set as TD or if you're American zc Z.

I can even say it Z.

S TD and G zip, which I better spell right, is the one that's used by pretty much every browser up

there.

So that's my first line.

Now I import static, which was defined in my Cadie file, and I import security.

And finally, they're not finally, what's the route where do the files for this live route star, which

says for everything Inver W-W dot ww w learn dash code gutsier and what do you want to do with those?

Act as a file server.

That one is actually done OK.

So now we'll do one from my API, even though it's not live yet, we can configure it, so we need to

pseudo Vij API lern dash code dot com and it's very similar now for this one.

I'm calling it API.

Learn, learn code orzio one name and it's directive's or even simpler I will encode.

Z, Steed and Joseph.

And I will import I don't need Stata because it doesn't have static files, but I'll import security

and I'll also say reverse proxy.

To localhost 4000 now, if you recall our back end application lessons on four thousand by default,

and the reason I'm reverse proxy from Cadie to that application is so I don't have to worry about setting

up SSL certificates.

Katti will take care of that for me.

So all this, this little directive says, is forever a request that comes into API.

Don't learn code dossie.

Forward it to the back end reverse proxy it to the localhost part four thousand.

And that connects our web server to our backend API application.

So I'll save and quit this.

And now, even though I won't be able to see a list of movies, I should be able to start cati.

Which I I'd stopped before this election, so pseudo service Cadie start.

And make sure there are no errors by going pseudo service katti status.

And it started so I should be able now to go back to my Web browser, open a new tab and say w w w don't

learn the code dossie.

And there it is now, it might take a second or two the first time you hit it for the SSL certificate

to be issued, so just give it a few seconds and reload the page and you should have it.

Now I can show things.

I could display the login screen, of course, but if I try to list the movies, it's just going to

say loading because it can't actually connect to the back end its connection to the web server.

But there's nothing running on the back end.

So we have a few more things to do.

So I was going to close this tab, go back to my terminal, and what I'm going to have to do next is

to start our back end application running on Port 4000 and make sure that it stays running and we'll

take care of that in the next election.

## 106 - Running the Go back end with supervisor

So the last step to get our application actually in production is to make sure that our backend server,

our API, is actually running all the time.

And there are lots of ways to do this.

But the way that I typically go with it is to use something called supervisor, which is available for

pretty much every distribution of Linux.

Now I happen to be an Ubuntu and if I wanted to install that, I'm logged into my server in my home

directory, I would just go and install supervisor and I will tell me I already have it installed,

but that's the command.

So it's already installed.

Now, once you have it installed, go to slash etsi slash supervisor and if you do an there, you'll

see that there is the configuration file which is fine by default.

And the Conforte.

Let's go in there.

And if I do now, as you see, I already have one thing running in supervisor, but what I would need

to do is to create another configuration file which will start our API running and keep it running.

And you do that by typing pseudo VI and give it some name.

I'll call my API and it has to end in dot com and if so, do that.

Now this is a really simple file.

It begins in Scherba Square brackets.

You have the word program colon and then what you want to call a combined API.

And if you have multiple things running under supervisor, of course they all have to have a unique

name.

So mine is API.

The next line I'm going to actually paste.

I've copied this already and I'm going to paste it because here's where I'm starting my application,

which if you recall existence w w w slash API and it's called go movies.

But I'm going to set the environment variable for our Jason Web token and it's a long one, so I didn't

want to type it by hand, so I'm just going to paste it and there it is.

So it begins with command equals.

And then I specify at runtime the time you start to specify an environment variable called Go Underscore

Movies, underscore JWT and then I have the token.

Then I specify the full path name to the binary, which is virus with API, single movies and then the

command flags.

So I have my DSN flag and my production flag and I'm not going to bother with port because I'm listening

on Port 4000 by default.

So the next line is directory.

And again, that's just our w w w slash api dot go dash, learn the dash code, learn dash code Nazia.

Which reminds me, I actually copied this and the directory is not slash API.

So let's go all the way over to this and fix that API dot learned code dossier, that's better.

So after the directory line we have two options ought to restart if it goes down.

Do you want to try to bring it back up?

Yes, I want this running all the time, so I'll make that true.

And do you want it to auto start?

Yes, I want that to start on application or server reboot.

And finally, let's create a log file so t underscore log file and I'll just put that in first.

WTW slash API learn dash code gutsier and I'll call it API dot log.

OK, so I've created that and let's see if I got it right in order to tell supervisor about this file

will run pseudo supervisor CTO and it's not there.

I see ERPs, the one that I had before I started the lecture.

So let's go reread it and update.

And status, and there it's up and running, so now I have my back and running, I have my front end

running.

Let's see if it all works.

Let's go to our Web browser and go to we don't learn learn the code dossier, which is what I called

mine.

OK, that's up and running.

Can I see the list of movies?

I can can I look at The Dark Knight?

I can.

Let's make sure we get the JWT correct log in the ad here.

Dossie and password.

And there it is, one password wants to save it for me, but it is working so I can manage the catalog,

I can look at a movie, I can go to craft.

You will.

I can look at The Dark Knight and there it is.

Everything is working exactly as it should.

So as I said, I have posted configuration files for other popular Web servers to the previous lecture,

but this should be enough to get you up and running.
