# Building Immutable Data Structures In Go

[article](https://levelup.gitconnected.com/building-immutable-data-structures-in-go-56a1068c76b2)

[Check out this great wiki for slice tricks](https://github.com/golang/go/wiki/SliceTricks)

Immutable just means that the original structure is not changed, instead a new copy of the structure is created with the new property value.
Let�s first look at a simple case:

```go
type Person struct {
    Name           string
    FavoriteColors []string
}
```

Obviously we can instantiate a `Person` and modify properties at will.
There is nothing wrong with this approach, per se.
However, when you get into more complicated nested structures that pass around references, slices and copy through channels, shared copies of the data can be modified in ways that create very subtle bugs.

There are other benefits to immutable data structures beyond avoiding bugs:

- Since the state never changed in place, it�s great for general debugging and recording each step of the transition to be carefully inspected later.
- Undo, or the ability to �go back in time� is not only possible but trivial as it�s done with an assignment.
- Shared state is widely considered to be a bad idea because to be implemented correctly and safety requires the performance hit and complexity of carefully placed/tested memory locking.

## Getters and Withers

`Getters` return data, `setters` mutate state, `withers` create a new state.

With `getters` and `withers` we can control exactly which properties are allowed to be changed.
It also gives us a great way to record transitions (later).
The new code looks like this:

```go
type Person struct {
    name           string
    favoriteColors []string
}
func (p Person) WithName(name string) Person {
    p.name = name
    return p
}
func (p Person) Name() string {
    return p.name
}
func (p Person) WithFavoriteColors(favoriteColors []string) Person {
    p.favoriteColors = favoriteColors
    return p
}
func (p Person) FavoriteColors() []string {
    return p.favoriteColors
}
```

The important things to note here are:

- The `Person` properties are private so that external packages cannot circumvent the methods.
- The functions for `Person` do not receive a `*Person`. This ensures that the structure is passed by value and returned by value.
- Notice that I use the word `With` rather than `Set` to make the distinction that it's the returned value that's important and the original object is not modified as a setter would imply.
- The properties are still accessible (and therefore mutable) by other code in this package. You should never interact with the property directly, always use the methods even in the same package.
- Each of the `withers` returns `Person` so they can be chained:

```go
me := Person{}.
    WithName("Elliot").
    WithFavoriteColors([]string{"black", "blue"})
        
fmt.Printf("%+#v\n", me)
// main.Person{name:"Elliot", favoriteColors:[]string{"black", "blue"}}
```

## Handling Slices

So far it's still not ideal because we are returning a slice for the favourite colors.
Since slices are passed by reference we can see an example of a bug that might otherwise go unnoticed:

```go
func updateFavoriteColors(p Person) Person {
    colors := p.FavoriteColors()
    colors[0] = "red"
    
    return p
}
func main() {
    me := Person{}.
        WithName("Elliot").
        WithFavoriteColors([]string{"black", "blue"})
        
    me2 := updateFavoriteColors(me)
        
    fmt.Printf("%+#v\n", me)
    fmt.Printf("%+#v\n", me2)
}
// main.Person{name:"Elliot", favoriteColors:[]string{"red", "blue"}}
// main.Person{name:"Elliot", favoriteColors:[]string{"red", "blue"}}
```

We intended to change the first color, but it has the side effect of mutating the `me` variable as well.
Since this isn't something that would prevent the code proceeding in a more complex application trying to hunt down a mutation like this can be really frustrating and time consuming.
One solution is to make sure we never assign by index and always assign a new slice:

```go
func updateFavoriteColors(p Person) Person {
    return p.
        WithFavoriteColors(append([]string{"red"}, p.FavoriteColors()[1:]...))
}
// main.Person{name:"Elliot", favoriteColors:[]string{"black", "blue"}}
// main.Person{name:"Elliot", favoriteColors:[]string{"red", "blue"}}
```

This is clunky and error-prone in my opinion.
A better way is to never return the slice in the first place.
Expand your getters and withers to operate only on elements (rather than the whole slice):

```go
func (p Person) NumFavoriteColors() int {
    return len(p.favoriteColors)
}
func (p Person) FavoriteColorAt(i int) string {
    return p.favoriteColors[i]
}
func (p Person) WithFavoriteColorAt(i int, favoriteColor string) Person {
    p.favoriteColors = append(p.favoriteColors[:i],
        append([]string{favoriteColor}, p.favoriteColors[i+1:]...)...)
    return p
}
```

Now we can safely use:

```go
func updateFavoriteColors(p Person) Person {
    return p.WithFavoriteColorAt(0, "red")
}
```

## Constructors

In some cases we can assume that the struct defaults are sensible.
However, it�s best to always create a constructor so that if we do need to change defaults in the future it exists in a single place:

```go
func NewPerson() Person {
    return Person{}
}
```

You can instantiate Person however you like, but I prefer to always do state transitions through the setters to keep it consistent:

```go
func NewPerson() Person {
    return Person{}.
        WithName("No Name")
}
```

## Interfaces

Up to this point we have still been using a public struct.
This can be painful for testing as we are at the mercy of these methods and creating mocks may have unwanted side effects.
We can create an interface of the same name and make the struct private by renaming it to `person`:

```go
type Person interface {
    WithName(name string) Person
    Name() string
    WithFavoriteColors(favoriteColors []string) Person
    NumFavoriteColors() int
    FavoriteColorAt(i int) string
    WithFavoriteColorAt(i int, favoriteColor string) Person
}
type person struct {
    name           string
    favoriteColors []string
}
```

Now we can create testing mocks by only overriding the logic we wish to stub:

```go
type personMock struct {
    Person
    receivedNewColor string
}
func (m personMock) WithFavoriteColorAt(i int, favoriteColor string) Person {
    m.receivedNewColor = favoriteColor
    return m
}
```

The test code may looking something like this:

```go
mock := personMock{}
result := updateFavoriteColors(mock)
    
result.(personMock).receivedNewColor // "red"
```

## Recording Changes

As I mentioned earlier full state transitions are great for debugging and we can catch all or some transitions by hooking into the withers:

```go
func (p person) nextState() Person {
    fmt.Printf("nextState: %#+v\n", p)
    return p
}
func (p person) WithName(name string) Person {
    p.name = name
    return p.nextState() // <- Use "nextState" whenever you return.
}
```

If you have more complex logic, or if you prefer, you can use a `defer` instead:

```go
func (p person) WithFavoriteColors(favoriteColors []string) Person {
    defer func() {
        p.nextState()
    }()
    p.favoriteColors = favoriteColors
    return p
}
```

Now we can see the changes:

```go
nextState: main.person{name:"No Name", favoriteColors:[]string(nil)}
nextState: main.person{name:"Elliot", favoriteColors:[]string(nil)}
nextState: main.person{name:"Elliot", favoriteColors:[]string{"black", "blue"}}
```

You may want to add a lot more information to this.
Such as timestamps, stack traces or other custom context to make debugging easer.

## History and Rollback

Instead of printing the changes we can collect the states as a history:

```go
type Person interface {
    // ...
    AtVersion(version int) Person
}
type person struct {
    // ...
    history        []person
}
func (p *person) nextState() Person {
    p.history = append(p.history, *p)
    return *p
}
func (p person) AtVersion(version int) Person {
    return p.history[version]
}
func main() {
    me := NewPerson().
        WithName("Elliot").
        WithFavoriteColors([]string{"black", "blue"})
    
    // We discard the result, but it will be put into the history.    
    updateFavoriteColors(me)
    fmt.Printf("%s\n", me.AtVersion(0).Name())
    fmt.Printf("%s\n", me.AtVersion(1).Name())
}
// No Name
// Elliot
```

This is great for when introspection needs to happen at the end.
It's also useful to record all the history to be logged if something goes wrong later, otherwise the history can be just discarded with the instance.
