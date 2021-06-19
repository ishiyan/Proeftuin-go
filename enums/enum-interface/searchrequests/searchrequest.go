package searchrequests

import (
	"fmt"
)

type (
	SearchRequest interface {
		fmt.Stringer
		Int() int
		sr() SearchRequest // Merovious' trick from https://play.golang.org/p/0B-fmVhhZa
	}

	Universal struct{}
	Web       struct{}
	Images    struct{}
	Local     struct{}
	News      struct{}
	Products  struct{}
	Video     struct{}
)

var iter = [...]SearchRequest{Universal{}, Web{}, Images{}, Local{}, News{}, Products{}, Video{}}

func Iter() *[7]SearchRequest {
	return &iter
}

func FromInt(i int) (SearchRequest, error) {
	if i < 0 || i > 6 {
		return nil, fmt.Errorf("invalid search request int: %d", i)
	}
	return iter[i], nil
}

func FromString(s string) (e SearchRequest, err error) {
	switch s {
	case "UNVERSAL":
		e = Universal{}
	case "WEB":
		e = Web{}
	case "IMAGES":
		e = Images{}
	case "LOCAL":
		e = Local{}
	case "NEWS":
		e = News{}
	case "PRODUCTS":
		e = Products{}
	case "VIDEO":
		e = Video{}
	}
	if e == nil {
		return nil, fmt.Errorf("invalid search request string: %s", s)
	}
	return e, nil
}

func (Universal) String() string {
	return "UNIVERSAL"
}

func (Universal) Int() int {
	return 0
}

func (u Universal) sr() SearchRequest {
	return u
}

func (Web) String() string {
	return "WEB"
}

func (Web) Int() int {
	return 1
}

func (w Web) sr() SearchRequest {
	return w
}

func (Images) String() string {
	return "IMAGES"
}

func (Images) Int() int {
	return 2
}

func (i Images) sr() SearchRequest {
	return i
}

func (Local) String() string {
	return "LOCAL"
}

func (Local) Int() int {
	return 3
}

func (l Local) sr() SearchRequest {
	return l
}

func (News) String() string {
	return "NEWS"
}

func (News) Int() int {
	return 4
}

func (n News) sr() SearchRequest {
	return n
}

func (Products) String() string {
	return "PRODUCTS"
}

func (Products) Int() int {
	return 5
}

func (p Products) sr() SearchRequest {
	return p
}

func (Video) String() string {
	return "VIDEO"
}

func (Video) Int() int {
	return 6
}

func (v Video) sr() SearchRequest {
	return v
}
