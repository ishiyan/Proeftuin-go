package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type htmlElement struct {
	name, text string
	elements   []htmlElement
}

func (e *htmlElement) String() string {
	return e.string(0)
}

func (e *htmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type htmlBuilder struct {
	rootName string
	root     htmlElement
}

func newHTMLBuilder(rootName string) *htmlBuilder {
	b := htmlBuilder{rootName, htmlElement{rootName, "", []htmlElement{}}}
	return &b
}

func (b *htmlBuilder) String() string {
	return b.root.String()
}

func (b *htmlBuilder) AddChild(childName, childText string) {
	e := htmlElement{childName, childText, []htmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

func (b *htmlBuilder) AddChildFluent(childName, childText string) *htmlBuilder {
	e := htmlElement{childName, childText, []htmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func main() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Printf("%s\n", sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	// <ul><li>...</li><li>...</li><li>...</li></ul>'
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	b := newHTMLBuilder("ul")
	b.AddChildFluent("li", "hello").AddChildFluent("li", "world")
	fmt.Println(b.String())
}
