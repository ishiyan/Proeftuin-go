# Learn how to build a single page application in React, with data supplied by a Go back end REST API
Trevor Sawler, Ph.D.

[github](https://github.com/tsawler?tab=repositories)
[web](https://www.gocode.ca/)
[twitter](https://twitter.com/tsawler)
[linked-in](https://linkedin.com/https://ca.linkedin.com/pub/trevor-sawler/8a/10b/381)

Useful resources.

- [The Documentation for React](https://reactjs.org/docs/getting-started.html)
- [React on W3Schools](https://www.w3schools.com/whatis/whatis_react.asp)
- [The React.js cheatsheet](https://devhints.io/react)
- [The React Lifecycle Methods Diagram](https://projects.wojtekmaj.pl/react-lifecycle-methods-diagram/)
- [React Router documentation](https://reactrouter.com/web/guides/quick-start)
- [GraphQL documentation](https://graphql.org/learn/)
- [Go By Example](https://gobyexample.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [DBeaver Community](https://dbeaver.io/)
- [Postico](https://eggerapps.at/postico/)

## 01 - Getting started with React

- 08 - How React Works
- 09 - How to use the downloadable code
- 10 - Our first React app
- 11 - The obligatory Hello world app
- 12 - Working with components
- 13 - Styling components
- 14 - Using a CSS Framework
- 15 - More about the CSS Framework
- 16 - Components and props
- 17 - More about props
- 18 - React Events
- 19 - More events
- 20 - Refs
- 21 - Simplifying things with state
- 22 - More about state - lifting state to share data between components
- 23 - Functional components
- 24 - Cleaning things up

- [SyntheticEvent](https://reactjs.org/docs/events.html)
- [Refs and the DOM](https://reactjs.org/docs/refs-and-the-dom.html)

## 02 - Building the Front End

- 25 - What are we going to create?
- 26 - A note about React Router 6
- 27 - Creating our front end application and introducting the React Router
- 28 - Routing to a component
- 29 - Challenge: Route to components
- 30 - Solution to Challenge
- 31 - More about routing (and a bit about the React lifecycle)
- 32 - More about routing Part II
- 33 - More about routing Part III
- 34 - Displaying one movie

```bash
npm install react-router-dom

# React router 6 recently came out, and once it's complete, with it's backwards compatibility logic,
# and some other bits that are not yet complete, I'll update the lectures for that version.
# If you use the first command, you'll get version 6, and that will cause you to have errors.
# I strongly encourage you to install version 5 of the React Router by using this command:
npm install react-router-dom@v5.3.0
```

[The React Lifecycle Methods Diagram](https://projects.wojtekmaj.pl/react-lifecycle-methods-diagram/)

## 03 - Setting up Go Backend as a REST API

- 35 - Installing the necessary software
- 36 - Setting up the Go project
- 37 - Installing a router and creating better handlers
- 38 - Models
- 39 - Setting up a simple API route
- 40 - Improved error handling
- 41 - Creating the database
- 42 - Creating our connection pool and connecting to the database
- 43 - Database functions and a challenge
- 44 - Solution to challenge
- 45 - An aside - cleaning up our JSON feed
- 46 - Getting all movies as JSON
- 47 - Next Steps

[httprouter](https://github.com/julienschmidt/httprouter)

### Go & postgres

```bash
docker pull postgres
docker images

mkdir -p /postgresdata
mkdir c:\postgresdata

docker volume create postgresdata
docker volume inspect postgresdata
docker volume rm postgresdata

# Start the Docker container with the run command using the postgres image.
# The /var/lib/postgresql/data directory in the container is mounted as /postgresdata on the host.
# Additionally, this command changes the name of the container to postgres:
# -it – Provides an interactive shell to the Docker container.
# -v – Use this option to attach the /postgresdata host volume to the /var/lib/postgresql/data container volume.
# -d – Starts the container as a background process.
# --name – Name of the container.
# default user: postgres, set POSTGRES_USER to override he username
docker run -it -v postgresdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=password -p 5432:5432 --name postgresdb -d postgres

# Once the postgresdb server starts running in a container, check the status
docker ps

# Always check the Docker log to see the chain of events after making changes
docker logs postgresdb

# The container is currently running in detached mode.
# Connect to the container using the interactive terminal instead
docker exec -it postgresdb bash

# Start the pql shell by typing in the interactive terminal
psql -U postgres
# or
docker exec -it postgresdb psql -U postgres

# Type \q to leave the psql shell and then exit once again to leave the Interactive shell
\q
exit

# Stop
docker stop postgresdb

# Start again
docker start postgresdb
```

#### Get driver

```bash
go get github.com/lib/pq
```

#### Create a db, create  a user, and grant privileges

```sql
CREATE DATABASE bookstore;
CREATE USER bond WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE bookstore to bond;
```

#### Switch to your bookstore database

You should already have a `bookstore` database:

list databases

```bash
\l
```

switch into that database

```bash
\c bookstore
```

directory of tables, if any

```bash
\d
```

#### Create table

```sql
CREATE TABLE books (
  isbn    char(14)     PRIMARY KEY NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);
```

directory of tables

```bash
\d
```

details of table `books`

```bash
\d books
```

#### Insert records

```sql
INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccol� Machiavelli', 6.99);
```

view records

```sql
SELECT * FROM books;
```

## 04 - Connecting to our REST API

- 48 - Setting up CORS middleware
- 49 - Getting the list of movies
- 50 - Checking for errors
- 51 - Displaying one movie
- 52 - Getting started with Movies by Genre
- 53 - Getting Genres from back end
- 54 - Displaying the list of Genres
- 55 - Getting movies by Genre
- 56 - Displaying movies by Genre
- 57 - Showing Genre name - an alternative to lifting state
- 58 - Code clean up

## 05 - Working with forms, React, and Go

- 59 - Building a form in React
- 60 - Making our form a controlled component, and binding it to state
- 61 - Making form inputs reusable components and a Challenge
- 62 - Solution to Challenge
- 63 - Creating a reusable select component
- 64 - Prepopulating the form with an existing movie
- 65 - Sending data to the REST back end
- 66 - Client side form validation
- 67 - Receiving data on the REST back end
- 68 - Providing feedback with a reusable alert
- 69 - Editing an existing movie
- 70 - Deleting a movie
- 71 - Adding a confirmation step when deleting movies
- 72 - Implementing delete on the back end
- 73 - Connecting our delete button to the REST back end
- 74 - Challenge: displaying list of movies to edit
- 75 - Solution to challenge

- [Bootstrap Validation](https://getbootstrap.com/docs/5.0/forms/validation/)
- [react-confirm-alert](https://www.npmjs.com/package/react-confirm-alert)

## 06 - Securing routes in our REST API

- 76 - Generating JSON Web Tokens on the back end
- 77 - Changing App to a component, and setting up state
- 78 - Getting the JSON Web Token from the back end
- 79 - Handling a successful login
- 80 - Adding middleware to check for a valid token
- 81 - Protecting the route on our front end
- 82 - Adding redirects for protected components
- 83 - Challenge
- 84 - Solution to Challenge
- 85 - Saving our token when the user leaves the site
- 86 - Making better error responses from our back end
- 87 - Adding images

- [JSON Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token)
- [Generate a password hash](https://go.dev/play/p/uKMMCzJWGsW)
- [Generate a JWT secret](https://go.dev/play/p/s8KlqJIOWej)
- [HTTP status codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)

```go
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Generate a password hash
func main() {
	password := "password"
	
	
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	
	fmt.Println(string(hashedPassword))
}
```

```go
package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

// Generate a JWT secret
func main() {

    secret := "mysecret"
    data := "data"
    fmt.Printf("Secret: %s Data: %s\n", secret, data)

    // Create a new HMAC by defining the hash type and the key (as byte array)
    h := hmac.New(sha256.New, []byte(secret))

    // Write Data to it
    h.Write([]byte(data))

    // Get result and encode as hexadecimal string
    sha := hex.EncodeToString(h.Sum(nil))

    fmt.Println("Result: " + sha)
}
```

## 07 - Adding GraphQL into the equation

- 88 - What is GraphQL?
- 89 - Setting up a schema and REST endpoint for GraphQL
- 90 - Handling the GraphQL request
- 91 - Implementing GraphQL requests for all movies
- 92 - Adding a search endpoint
- 93 - Implementing GraphQL requests for search on front end
- 94 - Displaying one movie using GraphQL
- 95 - Updating the front end
- 96 - Modifying the back end to handle poster images
- 97 - Updating the front end to display the poster image
- 98 - Cleaning things up

- [GraphQL](https://graphql.org/)
- [The Movie DB](https://www.themoviedb.org/)
- [JSON from movie DB](https://api.themoviedb.org/3/search/movie?api_key=b41447e6319d1cd467306735632ba733&query=The%20Shawshank%20Redemption)
- [JSON-to-Go converter](https://mholt.github.io/json-to-go/)

## 08 - Deploying our app to a server

- 099 - Getting the React application ready for deployment
- 100 - Building the production ready React application
- 101 - Getting the Go project ready for deployment
- 102 - Building the Go back end for our remote server
- 103 - Copying files to the server
- 104 - Setting up the production database
- 105 - Setting up the web server
- 106 - Running the Go back end with supervisor

- [Set environment variable in Windows](http://www.dowdandassociates.com/blog/content/howto-set-an-environment-variable-in-windows-command-line-and-registry/)
- [React - Adding Custom Environment Variables](https://create-react-app.dev/docs/adding-custom-environment-variables/)
- [Caddy](https://caddyserver.com/)
- [Configuring Nginx for React Router](https://www.barrydobson.com/post/react-router-nginx/)
- [How to fix BrowserRouter for React Apps on Apache](https://www.andreasreiterer.at/fix-browserrouter-on-apache/)

## 09 - Converting to use functions and React Hooks

- 107 - About this section
- 108 - Converting the Movies.js component to a function with hooks
- 109 - Coverting the Genres.js component to a function with hooks
- 110 - Converting the OneMovie.js component to a function
- 111 - Converting the OneGenre.js component to a function
- 112 - Converting the EditMovie.js component to a function
- 113 - Challenge: convert Admin.js to a function
- 114 - Solution to challenge
- 115 - Convert Login.js to a function
- 116 - Convert App.js to a function

## 10 - Where to go from here

- 117 - React and Redux

[React Redux](https://react-redux.js.org/)
