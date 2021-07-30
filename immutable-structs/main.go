package main

import "fmt"

type Person interface {
	WithName(name string) Person
	Name() string
	WithFavoriteColors(favoriteColors []string) Person
	NumFavoriteColors() int
	FavoriteColorAt(i int) string
	WithFavoriteColorAt(i int, favoriteColor string) Person
	AtVersion(version int) Person
}
type person struct {
	name           string
	favoriteColors []string
	history        []person
}

func (p person) WithName(name string) Person {
	p.name = name
	return p.nextState()
}

func (p person) Name() string {
	return p.name
}

func (p person) WithFavoriteColors(favoriteColors []string) Person {
	defer func() {
		p.nextState()
	}()
	p.favoriteColors = favoriteColors
	return p
}

func (p person) FavoriteColors() []string {
	return p.favoriteColors
}

func (p person) NumFavoriteColors() int {
	return len(p.favoriteColors)
}

func (p person) FavoriteColorAt(i int) string {
	return p.favoriteColors[i]
}

func (p person) WithFavoriteColorAt(i int, favoriteColor string) Person {
	p.favoriteColors = append(p.favoriteColors[:i],
		append([]string{favoriteColor}, p.favoriteColors[i+1:]...)...)
	return p
}

func updateFavoriteColors(p Person) Person {
	// return p.WithFavoriteColorAt(0, "red")
	return p.WithFavoriteColors(append([]string{"red"}, p.(person).FavoriteColors()[1:]...))
}

func (p *person) nextState() Person {
	p.history = append(p.history, *p)
	return *p
}

func (p person) AtVersion(version int) Person {
	return p.history[version]
}

func NewPerson() Person {
	return person{}.
		WithName("No Name")
}

/* type personMock struct {
	Person
	receivedNewColor string
}

func (m personMock) WithFavoriteColorAt(i int, favoriteColor string) Person {
	m.receivedNewColor = favoriteColor
	return m
} */

func main() {
	me := NewPerson().
		WithName("Elliot").
		WithFavoriteColors([]string{"black", "blue"})

	me2 := updateFavoriteColors(me)

	fmt.Printf("%+#v\n", me)
	fmt.Printf("%+#v\n", me2)

	// mock := personMock{}
	// result := updateFavoriteColors(mock)
	// fmt.Println(result.(personMock).receivedNewColor)

	// We discard the result, but it will be put into the history.
	updateFavoriteColors(me)
	fmt.Printf("%s\n", me.AtVersion(0).Name())
	fmt.Printf("%s\n", me.AtVersion(1).Name())
}
