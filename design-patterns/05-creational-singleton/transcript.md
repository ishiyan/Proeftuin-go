# Singleton -- a design pattern everyone loves to hate... but is it really that bad

## Overview

All right.

We are now going to talk about one of the most hated design patterns out there -- the Singleton.

So about 10 or 15 years after the Gang of Four book published the authors of the Gang of Four.

The gang of four authors they met together to discuss which patterns are still relevant.

And here is a quote from Eric Gamma.

So when discussing which patterns to drop we found that we still love all of them.

Not really I'm in favor of dropping the Singleton its use is almost always a design smell.

So is the single thing really that bad.
Well we're about to find out.
But first of all we have to discuss what it actually is.

So for some components in the system it only makes sense to have one instance of such a component in the system at any given time.

So a simple example is a database repository.

Let's suppose that at the start of the program you're reading all of the database into memory.

Now the question is do you want to do this more than once.
Well not really.

You don't really need a second copy of the database and it would be wasteful both in terms of computational resources and also in terms of memory because you have to allocate twice as much memory.

Another example is something as simple as a factory.
So you have a factory which knows how to create objects.

But if the factory is itself stateless meaning it doesn't have any fields or anything then there is no point in having more than one instance of this factory.

It doesn't benefit anyone and there's no reason to allow people to create instances of the factory.

Why not just reuse the instance that you already have.

So why would you encounter such a situation well for example the construction of the object might be expensive like with the database so when you call the mechanism to load in the database that might be expensive in terms of both the processing time as well as memory and you only want to do it once you want to somehow whenever somebody actually addresses the repository for the first time you only want to initialize it once and then you want to work with the data that's already there.

So the idea is very simple we somehow want to give everyone the same instance of the object rather than creating a new instance every time somebody calls let's say a factory function.

So with this setup what we want to be able to do also is to prevent anyone from adding or creating additional copies because obviously anytime you create an original copy that causes you to do extra work.

And while maybe it's just a limitation of your system that no more than one object of a type is created at any point in time.

So we need to take care of many things one of those things is lazy instantiation you don't really want

to make this expensive constructor call until somebody actually needs it.

You don't want to sort of preemptively load a database if nobody's going to use that database.

So that's something that we're going to take care of as well.

So the singleton is quite simply a component which is instantiate and only once.

## Singleton

All right.

So before we start looking at various problems that people encounter with a single term we need to talk about how to actually construct a singleton and what is the basic motivation for doing so.

So here I have a text file with a bunch of cities as well as their respective populations.

So imagine this as a kind of database that you want to load into memory.

And obviously if you're loading it into memory you want to want to do it once you only want to have one instance of whatever struct is actually keeping this memory.

So how can we ensure that this is the case.
Well let's take a look.
So first of all I'm going to make a struct called Singleton database nodes.
No notice that I'm using the lowercase has here.

So that is a hint that people shouldn't be creating this directly they should be using some sort of factory function for initializing this so this is gonna be a struct which is just going to keep a single map of the various capitals although I think I saw in New York City.

That's not a capital.
But anyways we'll call it capitals.

So it's going to be a map from string to ends and it's going to basically store the name of a city and a respective population in that city.

So now you want some sort of utility method for actually getting to this data.

So you might have some sort of let's say Singleton database get population method where you specify the name of the city and it gives you just an integer without any kind of validation Eric checking you just return deadbeat capitals at a name.

We'll return to this method later on because it's actually kind of important to define it like this.

But now the problem is that we want people to only ever access one instance of this struct.

So how can we do this.

Well we can do this using the sink once a structure as well as just keeping a pointed to the one and only instance so we'll have a variable called once which is going to be of type sink once that's this construct that ensures that something gets called only once and then it's done.

And we'll also have an instance just as a singleton data is pointer.

There we go.

So now we'll have a function for actually getting or creating that that database of ours.

I'll call it get Singleton get Singleton database like so and it is going to return that Singleton database pointer.

So this is where we get to use this a once function.
Now there are actually different options here.

So in terms of making the whole thing thread safe because threat safety is important you don't want to threads to start initializing this object at the same time you want to control for it.

So in this case there are two options.

So one of them is using sync once and the other option is using just the package level in IT function.

But the other feature that we want so.
So this is a threat safety.

Another feature we want is laziness and laziness basically means that you only construct the database you only read it from a disk to memory.

Whenever somebody asks for it.

So laziness is not going to be guaranteed in the init function unfortunately but it can be guaranteed using single once inside our own function.

So that's exactly why we are taking this approach right here.

So here is how it goes.
You use once dot do.

And this is where you specify the function that you want to be called exactly once and this is where you get to initialize the database.

So this is quite a process.

Let me actually cut and paste that chunk of code which reads the data from a file so I'll paste it right here.

So this is a rather large chunk of code which reads data from a file given a particular path and just loads up a map from string to it and I'm not going to go into the details there's absolutely nothing sophisticated here.

So we're going to use this function here to actually construct the one and only instance of the database.

So here I'll say DV is a singleton database.

And here I have to provide the Capitals and you get those by obviously reading the data.

So here I'll say read data I will specify the path to capitals start to theme.

So let's let's make variable I'll call this caps like so.

And then just provide the caps as an argument to the Singleton database creation.

So now I have this database I can ensure that there are no errors.
So.

So if there is no era then I can say the DOT capitals equals caps like so and then finally I can return.

Well I can set the instance so I can set instance equals pointed to the B and after we're done with this one time initialization we can just return the instance pointer and we keep returning the instance pointer if there are several callers who are attempting to get data from this database.

So this is it.
We're on now have a fully functional Singleton we can try using it.

So here I'll make a database or get the database or I'll call get Singleton database and then this tried to get the population of Seoul so population is equal to DB they'll get a population of Seoul like cell and let's just print line let's print line.

The population of Seoul is equal to and then population.

So let's run this let's see if it actually works all right so we're getting seventeen point five million here or so everything seems to be functioning correctly.

So this is the implementation of the Singletons which is both lazy and thread safe.

So you get all of the benefits in just one single demo and in the subsequent demos we're going to discuss the problems with the Singleton design pattern as well as how those problems can be overcome.

### Singeton code: singleton.go and capitals.txt

```go
package main

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
  "strconv"
  "sync"
)

// think of a module as a singleton
type Database interface {
  GetPopulation(name string) int
}

type singletonDatabase struct {
  capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(
  name string) int {
  return db.capitals[name]
}

// both init and sync.Once are thread-safe
// but only sync.Once is lazy

var once sync.Once
//var instance *singletonDatabase
//
//func GetSingletonDatabase() *singletonDatabase {
//  once.Do(func() {
//    if instance == nil {
//      instance = &singletonDatabase{}
//    }
//  })
//  return instance
//}

var instance Database

// init() — we could, but it's not lazy

func readData(path string) (map[string]int, error) {
  ex, err := os.Executable()
  if err != nil {
    panic(err)
  }
  exPath := filepath.Dir(ex)

  file, err := os.Open(exPath + path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  result := map[string]int{}
  
  for scanner.Scan() {
    k := scanner.Text()
    scanner.Scan()
    v, _ := strconv.Atoi(scanner.Text())
    result[k] = v
  }

  return result, nil
}

func GetSingletonDatabase() Database {
 once.Do(func() {
   db := singletonDatabase{}
   caps, err := readData(".\\capitals.txt")
   if err == nil {
     db.capitals = caps
   }
   instance = &db
 })
 return instance
}

func GetTotalPopulation(cities []string) int {
  result := 0
  for _, city := range cities {
    result += GetSingletonDatabase().GetPopulation(city)
  }
  return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
  result := 0
  for _, city := range cities {
    result += db.GetPopulation(city)
  }
  return result
}

type DummyDatabase struct {
  dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
  if len(d.dummyData) == 0 {
    d.dummyData = map[string]int{
      "alpha" : 1,
      "beta" : 2,
      "gamma" : 3 }
  }
  return d.dummyData[name]
}

func main() {
  db := GetSingletonDatabase()
  pop := db.GetPopulation("Seoul")
  fmt.Println("Pop of Seoul = ", pop)

  cities := []string{"Seoul", "Mexico City"}
  //tp := GetTotalPopulation(cities)
  tp := GetTotalPopulationEx(GetSingletonDatabase(), cities)

  ok := tp == (17500000 + 17400000) // testing on live data
  fmt.Println(ok)

  names := []string{"alpha", "gamma"} // expect 4
  tp = GetTotalPopulationEx(&DummyDatabase{}, names)
  ok = tp == 4
  fmt.Println(ok)
}
```

```txt
Tokyo
33200000
New York
17800000
Sao Paulo
17700000
Seoul
17500000
Mexico City
17400000
Osaka
16425000
Manila
14750000
Mumbai
14350000
Delhi
14300000
Jakarta
14250000
```

## Problems with Singleton

Okay so in the previous demo we constructed a single tunnel which is lazy and fat safe and everything is perfect.

Now I want to show you the problems with a single then so why is it that a singleton is not always the best idea.

And the problem is that a singleton quite often breaks the dependency inversion principle which we talked about at the beginning of the course.

But let me show you how this might affect you.

So let's imagine that you want to perform some sort of research you want to get the total population of several cities so you go ahead and you make some sort of function which given a bunch of cities as just a string slice for example you get the total population of those cities so you have a function called get total population where you specify a bunch of cities and you return an integer.

So this the implementation is rather easy.

So you say result equals zero and then four underscore comma city in the range of cities what you do is you feed it to the Singleton So you say result plus equals and then you get the Singleton database.

That's the function that we created in the previous demo and then you get the population for a particular city so it might seem okay.

It might seem like everything is fine here and we can we can try testing this so we can write something of a unit test to make sure that it does in fact work so I can say cities is equal to and then let's have Seoul and Mexico City.

So I'll get that total population.

So T.P. is equal to get total population of cities and then I have some expectations.

So I want to make sure that the result is actually okay and the result is okay.

If the total population is equal to and then I have to look up those values.

So I go into capitals DST and I see that.
Let's see Seoul is what seventeen point five million plus.
Let's find Mexico City somewhere seventeen point four million.

So those are the total population so what I would expect to be the result and I can just F.A. print line whether or not we did in fact get the right results so I can run this and you can see that the result is true so the test obviously passes but there is a huge problem with the tests that I'm writing right now.

The huge problem is that the test is dependent upon data from a real life database and in real life software engineering you almost never test against a live database because well these values they are essentially magic values the database can change at any time so somebody can go into our capitals file and they can change something they can modify this data according to the latest census data from Mexico City and then all of a sudden our test would break because the test is dependent upon real life data and also there is a performance consideration because even though you just want to test this you want a unit test against this function here.

That's all you want you want a unit test this particular function but instead your test is turning from a unit test into an integration test because you are actually going into the database so you're testing not just that the sum up of the cities and their populations is correct but you're also testing that the database reads correctly and stuff which is totally not what you want you want to be able to supply some sort of fake database with predictable data which doesn't imply actually reading from disk.

So that is the kind of issue that we have right here in that with a singleton this is difficult.

And the reason why is difficult is because of this because essentially if you look here you are depending You're depending directly upon the Singleton database so here you are depending upon a concrete instance of the singleton database.

And here I want to remind you about the dependency inversion principle.

So one of the ideas of dependency inversion principle is that instead of depending on concrete implementations you want to depend upon abstractions which typically implies depending upon an interface and certainly if we decide to depend upon an interface.

This problem of reading a life database for the purposes of unit tests he kind of goes away and this will be going to take a look at in the next lesson.

## Singleton and Dependency Inversion

We really want our tests to depends on some sort of abstraction so that instead of depending on the concrete Singleton database we can substitute this database with something else and just provide a different implementation of CAP population.

So how can we make a dummy database.

Well first of all we have to introduce some sort of abstraction which has something in common between the real database and the dummy database.

So here I'll introduce an interface called database and this interface is going to have just a single method called Get population which is going to take the name as a string and return an integer.

Now interestingly enough we already have Singleton database implement this interface right here you can see get population.

So what we can now do is we can slightly change the way that we access the the the way that we calculate the total population so that we now introduce a dependency.

So here the total population is unfortunately hard coded to use the single then and that means there is no way for us to substitute something in here.

So what we can do is we can make a better function for calculating the total population.

I'll just put it down here.

So this is going to be getting total population X and is going to take a dependency upon the database like so.

So you provide the database as opposed to just expecting that it will find the database for you.

So here instead of get Singleton database you just use DB instead.

But apart from that nothing really changes the end operation is the same.

But now you have additional flexibility so if you do really want to test against the real life database you would call get told population X and you would provide the database right here yourself.

So you would say get Singleton database comma cities.
That is how you would operate.

But really what we want to be able to do is to get away from using a real life database and use some sort of dummy database.

So let's take a look at how we can do this using this new abstraction approach that we have created because what we want is we want a database with predictable values like instead of having real life cities we just want a map of let's say three values let's say alpha beta and gamma.

So what I can do is I can make a type called dummy database now a dummy database is also going to have data but this will be dummy data dummy data and that's going to be a map once again from string to end and we'll initialize it using once again the get population method.

So here I'm just going to go into my cogeneration tools and I'm going to implement the methods for for the database on a dummy database.

There we go.

So here is a dummy database I really do prefer to use pointers here as opposed to just value.

So this is how I would define this.
And so what do we want to do here.

Well you want to fill in dummy data and we want to do it lazily so I'm going to say if these are dummy data if the length of dummy data is equal to zero then we initialize it to the sum some values which are predictable so I can say that dummy data and then I can just make a map from string to end where we'll have three different values we'll have our phone with a value of one we'll have data with a value of two and we'll have gamma with a value of three and then we'll return that data at a particular index so calibrate should be here.

And then we return the dummy data at Nate.
There we go.

So now we're operating on dummy data and as a result we can write a proper unit test not some flaky integration test which is going to go into a real database will define a bunch of names.

But here I can have predictable names like alpha and gamma.

These aren't going to change and we don't have a brittle test because they are never going to change.

Then I can calculate total population.
And this time round I will say T.P. is equal to get total population X..

So this is the modification where you have first of all have to provide the database so here I can just make a dummy database like so I don't even need to for this database to be a single another I could make it the single done as well.

And then I provide the names of the cities and then we expect the result to be equal to for so FMC the print line we expect the total population to be equal to 4.

Let's why is it four by the way.

Well if you look at the data Alpha is 1 and gamma is 3 so one plus three equals four and now find gamma are in fact the values we are feeding into the test so let's actually run this.

Let's take a look at what we get.

And as you can see the output here is true so everything is working correctly.

So what does that take away from this example.
Well the takeaway is that the Singleton isn't really that scary.

The scary part is depending directly on the Singleton as opposed to depending on some interface that the Singleton implements because if you dependent the interface you can substitute that interface but of course in order to be able to substitute something you have to provide an API aware that something can be plugged in.

And that's exactly what we did here by making our modification of the ghettos population function which also takes the database up on which it operates.

And this is the thing that has allowed us to write proper unit tests.

## Summaty

All right so let's talk about what we saw about the scary Singleton design pattern.

So lazy one time instantiation of the singleton is possible using sink once and that solves all our problems relating to lazy instantiation as well as thread safety by the way so thread safety is handled by sync once as well.

But we do want to adhere to the dependency inversion principle so that our code remains testable and whatnot.

And the idea here is that instead of depending directly on the singleton and using that Singleton directly in your code what you want to be able to do is you want to have your Singleton implement some interface and then you depend on the interface.

And the reason why it's important is because then you can substitute the implementer of this interface.

You can replace the Singleton with let's say some sort of test dummy that you can use in your tests without let's say if you're a real Singleton goes into a life database.

You obviously don't want that in your test.

You want to provide alternatives and with this approach with the approach of depending on abstractions you get exactly that.

So the takeaway I think from this session this section of the course is that the singleton is not a scary pattern but you do have to be careful.

You do have to be careful because if you use it directly if you use the singleton and you depend directly on the Singleton then it's too strong a dependency is the dependency that might come to hurt you in the end so the dependency inversion principle is particularly relevant when talking about the Singleton.
