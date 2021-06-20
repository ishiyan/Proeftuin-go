# Composite -- treating individual and aggregate objects uniformly

## Overview

All right.
Let's talk about the composite design pattern.
So what is this all about.

Well sometimes you have a situation where objects use other objects fields or methods through the idea of embedding and composition lets us make compound objects so you can have objects made up of objects made up of objects to infinity almost.

So for example if you think about the mathematical expression that the expression would be composed of smaller expressions and those smaller expressions it would be composed of very simple expressions.

Well for example if you have a drawing application some sort of vector a drawing application you could have a group of shapes that you can drag around and resize.

But that group of shapes is composed of several different shapes so the composite design pattern allows us to treat both single objects and composite objects uniformly and by uniformly would typically mean that they would have the same interface.

So for example if you have an object of type foo and you also have an array or a slice of foo objects you can get them to have common API is common as in identical APIs.

And that's the goal of the composite design pattern.

So the composite is just a mechanism for treating both individual objects or a scalar objects and compositions of objects in a uniform manner.

## Geometric Shapes

If you think about a typical drawing application where you can draw all sorts of different vector shapes you know that you can take a shape and you can drag it around the screen but you can also take several shapes and you can drag them together around the screen as a single group.

So let's take a look at how we can implement a scenario very similar to what I've just described.

So imagine you have some sort of graphic object so a graphic object can be one of two things either it's a simple object like a square or a rectangle or something in which case it has a name and it has some color.

Let's keep both the strings but in addition a graphic object can be a combination of different objects that you selected and you also want to be able to like drag it around that resize it even though it's a group of different objects.

So in addition a graphic object can have a bunch of children which are sub objects also of type graphic object by the way so you can either have just a singular object or you can have a collection of object and it doesn't really matter here which one it is we don't make two separate types for this we just stick to a single type.

So now that we have this kind of graphic object we might want to.

Well first of all print the whole thing in and obviously printing it to a console for example would be completely different depending on whether it's just a single object in which case you can print the name and the color and that's it or whether it's a graphic object container with a bunch of children containing other objects and remember this is a recursive relationship so you can have a drawing.

For example you can define the drawing to be a graphic object and the drawing can contain a bunch of groups and a bunch of groups can contain a bunch of children so you can have a graphic object which contains a bunch of children which in turn contains a bunch of children themselves.

So in order to print this out we'll define a string method on going to.
I'm not going to actually write this out.

I would just show it right here so string actually uses a utility a method called print which also tracks of the depth of our recursion because we're going to go into objects and objects and objects.

There's plenty of code here for actually outputting the the sort of depth of the object so to speak before you actually print its children.

The implementation here doesn't really matter because what we want to be able to show is how the whole thing is composed.

Let's imagine that we have what we obviously have graphic object that we can initialize straight away but we can also let's say we can make squares and circles.

So we'll have a function called New circle where you provide the color and it returns a graphic object pointer.

So here we return the graphic object where the name is circle.
The color is color and there are no children like so.
So we'll put a new in there and we'll also have a function called New Square.
Actually I'll just copy this so just duplicate this and let's have you square.

So here the name would be square but everything else would be pretty much the same so would this what we can do is we can set up a scenario where we have a drawing which is a graphic object and we also have these psychos and squares on the drawing so I can make a drawing I can say that drawing is a graphic object where the name is my drawing.

There is no color as such and there are no children yet although we're going to add them in just a moment.

Actually what we can do is we can we can add them right here but let's add them on the next line so here I'll say drawing drawing the children and I can just append a new child for example I can make a new circle with the color blue or let's say the color red for example and then I can make a new square with the color yellow like so.

But in addition to having a bunch of shapes in a drawing what I can do is I can also make another set of shapes but I can add them to a group so I could say hey let's have a group which is going to be a graphics object let's call it group 1.

So Group 1 is not going to have a color it's not going to have any children until I actually add them and then I'll just say group the children dots and then I'll append let's say a new circle with the color blue and also a new square with the color blue.

And then what I can do is I can take this group and I can also added to that drawing so I can say drawing dot children dot append here and I can say let's append the group so we spend the entire group and what we can do now using that string method that I implemented is we can just print line the whole thing.

So here I can say drawing dot string.
There we go.

Okay so let's run this because this will be great way of illustrating what's going on so here you can see that we're using asterisks to indicate the fact that we have some depth so at the top level we have my drawing that is the graphic object that contains everything in our model and inside the drawing we have a red circle and a yellow sack but we also have a group and the group in turn contains a blue circle and a blue square.

You can see the two asterisks here showing that we have additional depth here.
That's why we have this additional depth parameter.

So the takeaway from this entire example is you can have these objects like a graphic object of infinite depth because a graphic object can contain any number of graphic objects to infinity and so a graphic object you can treated as a single shape or as a scalar shape.

So for example when you have new Circle or new square you're effectively making a scalar you're making a singular object not a collection of objects.

However if you do decide to manipulate the children of an object then all of a sudden it becomes a collection of objects.

And that is the goal of the composite design pattern that it doesn't really matter that it doesn't really matter whether or not your object is a scalar object or a collective object object which is composite which is made up of more than one element because ultimately you can write algorithms like the print method that we have right here which do not really care because all they do is they they perform the appropriate check.

So it looks a g dot children and if G dot children is nil then that's okay.

We just don't iterate anything but if you don't children is not nil then we perform the recursive iteration and we go into the into the depths of the object as far as is necessary.

So that is the gist of the composite design pattern.

### Geometric Shapes code: structural.composite.geometry.go

```go
package composite

import (
  "fmt"
  "strings"
)

type GraphicObject struct {
  Name, Color string
  Children []GraphicObject
}

func (g *GraphicObject) String() string {
  sb := strings.Builder{}
  g.print(&sb, 0)
  return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
  sb.WriteString(strings.Repeat("*", depth))
  if len(g.Color) > 0 {
    sb.WriteString(g.Color)
    sb.WriteRune(' ')
  }
  sb.WriteString(g.Name)
  sb.WriteRune('\n')

  for _, child := range g.Children {
    child.print(sb, depth+1)
  }
}

func NewCircle(color string) *GraphicObject {
  return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
  return &GraphicObject{"Square", color, nil}
}

func main() {
  drawing := GraphicObject{"My Drawing", "", nil}
  drawing.Children = append(drawing.Children, *NewSquare("Red"))
  drawing.Children = append(drawing.Children, *NewCircle("Yellow"))

  group := GraphicObject{"Group 1", "", nil}
  group.Children = append(group.Children, *NewCircle("Blue"))
  group.Children = append(group.Children, *NewSquare("Blue"))
  drawing.Children = append(drawing.Children, group)

  fmt.Println(drawing.String())
}
```

## Neural Networks

Let's take a look at yet another implementation of the composite design pattern and I'm sure you know that nowadays machine learning is all the rage and part of machine learning is the use of neural networks.

So let's see if we can model very simple neurons using the go programming language.

So what I'm going to do is I'll define a type called neuron and a neuron typically has connections to other neurons.

So basically there is incoming connections and outgoing connections and we can model those as just slices of neuron pointers.

So neurons connect to other neurons.

Now what you want to be able to do is you want to be able to connect one neuron to another.

So if we were to write a function for sort of neuron method called connect to.

So in your own wants to connect to some other neuron This is rather simple because you take the outgoing neurons of the current object and you just append to them the other neuron and then you take the other neurons incoming connections and you append the current neuron.

So that's that's pretty much all there is to it.

Now imagine this situation becomes more complicated imagine that it's not convenient to work with individual neurons you want to work with whole neuron layer as a neuron layer is basically a collection of neurons all stored together somewhere so we can have a type called neuron layer and then you're on layer is just a bunch of neurons.

So we'll have more neurons has your own slice.
There we go.

So this is in Neuron layer we can have a factory function for making a neuron layer of a particular size.

So new neuron layer a way you specify the count so the number of neurons you actually want like so.

And then we return a neuron layer where we simply make the number of neurons that we need.

So so this is how you would define it.
Obviously with curly braces as opposed to round braces.
There we go.
And that's pretty much it.
But now we have a problem.
So we want to be able to connect like we have this connect method.

We want to connect not just individual neurons but we also want to connect neuron layers together and furthermore we want to be able to interconnect neurons.

And you're on layers.
So let me show you a simple scenario.

Let's suppose we have two neurons you're on one and you're on two and I would just define those as neurons like so.

So they're going to be pointers and then we'll have layer 1 and layer 2.

They're going to be new neuron layer with three neurons and new neuron layer with four neurons for example what we want to be able to do is we want to have just a single function or indeed just a single method.

If we could somehow have a method for connecting neurons to neurons neurons to layers layers to neurons and layers to layer.

So what I want to be able to do is to connect neuron want in your on to here on to like.

So I want to be able to connect neuron one to layer one.
I want to be able to connect let's say layer two to neuron 1.

And finally I want to be able to into interconnect layers obviously interconnecting the layer implies that you take every single neuron from the first layer and you connect connected to every single neuron of the second layer layer one layer or two.

That's what it would imply and I want to be able to have all of this without writing for different functions.

So I want just a single function called Connect.
Now how can we actually implement this.

How can we get this to work considering that we need to somehow be able to iterate every single neuron inside either a neuron and I know that sounds weird like iterate every single year on the inside a scalar object.

That sounds weird or inside a neuron layer indeed within neuron layers it's a bit easier.

But let's take a look at how we can implement that.

So imagine you have some sort of interface like a neuron interface now and you're on the interface is going to have just a single method which you can call ITER or you can call it something else like collect for example or I dunno whatever you feel like it.

And that's basically going to give us a slice of neuron pointer.

So it's going to give us every single neuron that's contained inside a particular type of object.

Basically so when it comes to grabbing the objects of a neuron layer that's not really a problem because we already have a slice of neurons so we just collect their pointers and that's it.

So here let's actually go ahead and implement the.
Then you're on interface on the neuron layer.

So here what you would do is you would make result basically make a slice of a neuron pointers initially with size zero and then for I in a range of vendor neurons what you would do is you would just grab a pointer to every single neuron so you would append the result.

You would take and Dot neurons at position i and appended to the results and then return the result.

So that's simple enough.

That is how you would collect appointed to every single neuron inside in your own layer.

But another problem is we need to have the same thing on a neuron.

So let's actually implement this on the neuron as well so it would look something like this.

And here we would have to collect a pointer to ourselves because remember when we have a method that method has a receiver which is a neuron pointer so we can simply return that neuron pointer as part of a slice.

So here what I can do is I can simply say that we're going to return a slice of neuron pointers where the only pointer is ourselves the receiver and that's pretty much it.

Now you can see that now that we've implemented this interface in both a scalar object which has a single neuron as well as a collection object which is in your own layer we can write a single unifying connect function which would connect the two together.

Let me show you how this would work.

So you'd have a function called Connect which takes left and right which are both of type neuron interface.

Notice we're taking that interface type and then we simply iterate both of the Left dot either and right on it or so we say for underscore comma left in range of left dot either for on this call comma right in range right Dot.

Either we say l dot connect to Ah remember this connect to method is actually a method that is defined on a neuron so that interconnects one neuron with another neuron so we reuse this here and we have a unifying connect function so how does this work.

Well it works because ITER is defined on any type which implements you're on the interface and of course both neurons and your layers implement this interface.

So when we go and we iterate a neuron layer we get every single point her we get appointed to every single element of that neuron layer.

But if we iterate a scalar object if we iterate a single neuron what happens is we just get appointed to ourselves.

So we take the receiver and we pass that receiver as a slice thereby returning the one and only point your so we get in the way we get a scalar object to masquerade as if it were a collection.

So here regardless of what left and right are we iterate both of them we get the pointers for the left and right sides and then we use those pointers to interconnect the left and right sides and this is why all of these four calls now become valid they become completely legal because it no longer matters what you're passing in so long as every single type that you are passing in here implements this neuron interface that we've defined so this is another illustration of the composite design pattern.

### Neural Networks code: structural.composite.neuralnetworks.go

```go
package composite

type NeuronInterface interface {
  Iter() []*Neuron
}

type Neuron struct {
  In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
  return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
  n.Out = append(n.Out, other)
  other.In = append(other.In, n)
}

type NeuronLayer struct {
  Neurons []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
  result := make([]*Neuron, 0)
  for i := range n.Neurons {
    result = append(result, &n.Neurons[i])
  }
  return result
}

func NewNeuronLayer(count int) *NeuronLayer {
  return &NeuronLayer{make([]Neuron, count)}
}

func Connect(left, right NeuronInterface) {
  for _, l := range left.Iter() {
    for _, r := range right.Iter() {
      l.ConnectTo(r)
    }
  }
}

func main() {
  neuron1, neuron2 := &Neuron{}, &Neuron{}
  layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

  //neuron1.ConnectTo(&neuron2)

  Connect(neuron1, neuron2)
  Connect(neuron1, layer1)
  Connect(layer2, neuron1)
  Connect(layer1, layer2)
}
```

## Summary

All right so let's try to summarize what we've learned about the composite design pattern so we know that objects can use either objects via composition and some composed and singular objects need either similar or in fact identical behaviors.

So the composite design pattern actually lets us treat both types of these objects uniformly and we can support things like iteration.

So you can iterate composite objects but you can also iterate just singular objects.

The object will basically return itself when it comes to iterating it.

And for more information about iteration take a look at the iterator design pattern which is also covered as part of this course.
