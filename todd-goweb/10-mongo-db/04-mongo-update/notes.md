# Install Mongo

## Go get driver for mongo

```bash
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson
```

## Step 1

Don't run this code

Just making updates - a several step process.

We will need a mongo session to use in the CRUD methods.

We need our UserController to have access to a mongo session.

Let's add this to controllers/user.go

```go
UserController struct {  
    session *mgo.Session
}
```

And now add this to controllers/user.go

```go
func NewUserController(s *mgo.Session) *UserController {  
    return &UserController{s}
}
```

And now add this to main.go

```go
func getSession() *mgo.Session {
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://localhost")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}
```

and this

```go
uc := controllers.NewUserController(getSession())  
```

## Step 2

Don't run this code

Just making updates - a several step process.

IMPORTANT:
Make sure you update your import statements to import packages from the correct location!

In this step:

MongoDB represents JSON documents in binary-encoded format called BSON behind the scenes. BSON extends the JSON model to provide additional data types and to be efficient for encoding and decoding within different languages.

We will update our user model to change the type of our `Id` field to be a `bson.ObjectId`

Add this to `models/user.go`

```go
type User struct {
    Name   string        `json:"name" bson:"name"`
    Gender string        `json:"gender" bson:"gender"`
    Age    int           `json:"age" bson:"age"`
    Id     bson.ObjectId `json:"id" bson:"_id"`
}
```

## Step 3

Don't run this code yet

We create an `ObjectId` using the `bson` package.

We do this in `controllers/user.go` in func `CreateUser`

```go
    // create bson ID
    u.Id = bson.NewObjectId()
```

Second, we store the user in mongodb.

We do this in `controllers/user.go` in func `CreateUser`

```go
uc.session.DB("go-web-dev-db").C("users").Insert(u)
```

## Step 4

In this step:

We will get a user from mongodb

First we will get the user id from the URL

```go
id := p.ByName("id")
```

Next we will Verify that the id is an `ObjectId`

```go
if !bson.IsObjectIdHex(id) {
    w.WriteHeader(http.StatusNotFound) // 404
    return
}
```

`ObjectIdHex` returns an `ObjectId` from the provided hex representation.

```go
    oid := bson.ObjectIdHex(id)
```

Now we will create an empty user to store the results in

```go
    u := models.User{}
```

And then we will get the user information

```go
if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
    w.WriteHeader(404)
    return
}
```

### Run this code

1. Start your server

#### POST a user to mongodb

Enter this at the terminal

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"James Bond","gender":"male","age":32}' http://localhost:8080/user
```

-X is short for --request
Specifies a custom request method to use when communicating with the HTTP server.

-H is short for --header

-d is short for --data

#### GET a user from mongodb

Enter this at the terminal

```bash
curl http://localhost:8080/user/<enter-user-id-here>
```

## Step 5

In this step:

We will delete a user from mongodb.

This is identical to what we did in the last step to GET a user.

First we will get the user id from the URL

```go
id := p.ByName("id")
```

Next we will Verify that the id is an `ObjectId`

```go
if !bson.IsObjectIdHex(id) {
    w.WriteHeader(http.StatusNotFound) // 404
    return
}
```

`ObjectIdHex` returns an `ObjectId` from the provided hex representation.

```go
    oid := bson.ObjectIdHex(id)
```

Next, add code to delete the user

```go
if err := uc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
    w.WriteHeader(404)
    return
}
```

### Run the code

1. Start your server

## DELETE a user from mongodb

Enter this at the terminal

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
curl -X POST -H "Content-Type: application/json" -d '{"name":"James Bond","gender":"male","age":32}' http://localhost:8080/user

# on windows curl:
curl -X POST -H "Content-Type: application/json" -d "{\"name\":\"Miss Moneypenny\",\"gender\":\"female\",\"age\":27}" http://localhost:8080/user
curl -X POST -H "Content-Type: application/json" -d "{\"name\":\"James Bond\",\"gender\":\"male\",\"age\":32}" http://localhost:8080/user
```

```bash
curl http://localhost:8080/user/<enter-user-id-here>
```

```bash
curl -X DELETE http://localhost:8080/user/<enter-user-id-here>
```
