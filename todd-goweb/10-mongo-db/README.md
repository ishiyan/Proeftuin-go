# Mongo DB

1 Start your server

```bash
go run main.go
```

2 Enter this at the terminal

```bash
curl http://localhost:8080
```

This code forked from Steven White and his excellent articles:

- `https://stevenwhite.com/building-a-rest-service-with-golang-1/`
- `https://stevenwhite.com/building-a-rest-service-with-golang-2/`
- `https://stevenwhite.com/building-a-rest-service-with-golang-3/`

## Organizing code into packages

We are first going to lay groundwork for building an application which works with Mongodb. Our initial groundwork will be to setup a server using Julien Schmidt’s router.

We will then create functionality to get a user. We will also use packages to organize our code.

## Create user & delete user

Continuing to build our application, we will now stub out functionality for creating and deleting a user. We will test this functionality with curl

## MVC design pattern - model view controller

Model–view–controller (MVC) is a software design pattern for implementing user interfaces on computers. It divides a given software application into three interconnected parts, so as to separate internal representations of information from the ways that information is presented to or accepted from the user. Traditionally used for desktop graphical user interfaces (GUIs), this architecture has become popular for designing web applications and even mobile, desktop and other clients.One of the seminal insights in the early development of graphical user interfaces, MVC became one of the first approaches to describe and implement software constructs in terms of their responsibilities. Trygve Reenskaug introduced MVC into Smalltalk-76 while visiting the Xerox Palo Alto Research Center (PARC) in the 1970s.

### MODEL

The central component of MVC, the model, captures the behavior of the application in terms of its problem domain, independent of the user interface. The model directly manages the data, logic, and rules of the application. A model stores data that is retrieved according to commands from the controller and displayed in the view.

### VIEW

A view can be any output representation of information, such as a chart or a diagram. Multiple views of the same information are possible, such as a bar chart for management and a tabular view for accountants. A view generates new output to the user based on changes in the model.

### CONTROLLER

The third part, the controller, accepts input and converts it to commands for the model or view. A controller can send commands to the model to update the model's state (e.g., editing a document). It can also send commands to its associated view to change the view's presentation of the model (e.g., scrolling through a document).

Although originally developed for desktop computing, MVC has been widely adopted as an architecture for World Wide Web applications in major programming languages. Several commercial and noncommercial web frameworks have been created that enforce the pattern. These software frameworks vary in their interpretations, mainly in the way that the MVC responsibilities are divided between the client and server. Early web MVC frameworks took a thin client approach that placed almost the entire model, view and controller logic on the server. This is still reflected in popular frameworks such as Django, Rails and ASP.NET MVC. In this approach, the client sends either hyperlink requests or form input to the controller and then receives a complete and updated web page (or other document) from the view; the model exists entirely on the server. As client technologies have matured, frameworks such as AngularJS, EmberJS, JavaScriptMVC and Backbone have been created that allow the MVC components to execute partly on the client using Ajax.

The use of the MVC pattern in web applications exploded in popularity after the introduction of Apple's WebObjects in 1996, which was originally written in Objective-C (that borrowed heavily from Smalltalk) and helped enforce MVC principles. Later, the MVC pattern became popular with Java developers when WebObjects was ported to Java. Later frameworks for Java, such as Spring (released in October 2002), continued the strong bond between Java and MVC. The introduction of the frameworks and Django (July 2005, for Python) and Rails (December 2005, for Ruby), both of which had a strong emphasis on rapid deployment, increased MVC's popularity outside the traditional enterprise environment in which it has long been popular. MVC web frameworks now hold large market-shares relative to non-MVC web toolkits.

## Install mongodb

To install mongodb, go to `https://www.mongodb.com/` and find the download area. Select the community server. Choose the operating system of the computer upon which you will install mongodb. Look for a link that offers instructions. Read and follow those instructions. When you have mongodb running, it will listen for connections on port 27017.

```bash
docker pull mongo
docker images

# By default, MongoDB stores data in the /data/db directory within the Docker container.
# To remedy this, mount a directory from the underlying host system to the container running the MongoDB database.
# This way, data is stored on your host system and is not going to be erased if a container instance fails.
mkdir -p /mongodata
mkdir c:\mongodata

docker volume create mongodata
docker volume inspect mongodata
docker volume rm mongodata

# Start the Docker container with the run command using the mongo image.
# The /data/db directory in the container is mounted as /mongodata on the host.
# Additionally, this command changes the name of the container to mongodb:
# -it – Provides an interactive shell to the Docker container.
# -v – Use this option to attach the /mongodata host volume to the /data/db container volume.
# -d – Starts the container as a background process.
# --name – Name of the container.
docker run -it -v mongodata:/data/db -p 27017:27017 --name mongodb -d mongo

# Once the MongoDB server starts running in a container, check the status
docker ps

# Always check the Docker log to see the chain of events after making changes
docker logs mongodb

# The container is currently running in detached mode.
# Connect to the container using the interactive terminal instead
docker exec -it mongodb bash

# Start the MongoDB shell by typing mongo in the interactive terminal
mongo -host localhost -port 27017

# Type exit to leave the MongoDB shell and then exit once again to leave the Interactive shell
exit
exit

# Stop
docker stop mongodb

# Start again
docker start mongodb
```

### Install driver

You will also need to download and install a driver for mongodb using

```bash
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
```
