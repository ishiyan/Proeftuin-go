package main

import "fmt"

// combination of OCP and Repository demo

type color int

const (
	red color = iota
	green
	blue
)

type size int

const (
	small size = iota
	medium
	large
)

type product struct {
	name  string
	color color
	size  size
}

type filter struct {
}

func (f *filter) filterByColor(
	products []product, color color) []*product {
	result := make([]*product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *filter) filterBySize(
	products []product, size size) []*product {
	result := make([]*product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *filter) filterBySizeAndColor(
	products []product, size size,
	color color) []*product {
	result := make([]*product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// filterBySize, filterBySizeAndColor

type specification interface {
	IsSatisfied(p *product) bool
}

type colorSpecification struct {
	color color
}

func (spec colorSpecification) IsSatisfied(p *product) bool {
	return p.color == spec.color
}

type sizeSpecification struct {
	size size
}

func (spec sizeSpecification) IsSatisfied(p *product) bool {
	return p.size == spec.size
}

type andSpecification struct {
	first, second specification
}

func (spec andSpecification) IsSatisfied(p *product) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

type betterFilter struct{}

func (f *betterFilter) filter(
	products []product, spec specification) []*product {
	result := make([]*product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := product{"Apple", green, small}
	tree := product{"Tree", green, large}
	house := product{"House", blue, large}

	products := []product{apple, tree, house}

	fmt.Print("Green products (old):\n")
	f := filter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	// ^^^ BEFORE

	// vvv AFTER
	fmt.Print("Green products (new):\n")
	greenSpec := colorSpecification{green}
	bf := betterFilter{}
	for _, v := range bf.filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := sizeSpecification{large}
	largeGreenSpec := andSpecification{largeSpec, greenSpec}

	fmt.Print("Large blue items:\n")
	for _, v := range bf.filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
