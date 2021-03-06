package config

import (
	"fmt"

	// _ is needed to use the driver
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

// DB is a database
var DB *mgo.Database

// Books is a collection of books
var Books *mgo.Collection

func init() {
	// get a mongo sessions
	//s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")
	s, err := mgo.Dial("mongodb://localhost/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")

	fmt.Println("You connected to your mongo database.")
}
