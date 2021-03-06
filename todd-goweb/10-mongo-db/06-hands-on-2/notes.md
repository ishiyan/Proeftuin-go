# Use a map

Use the code from the previous hands-on-1 folder.

There is a map that holds all of the user data.

Every time a user is created or deleted, write this map as JSON to a file.

Also, when your program starts, if there is a file with JSON data in it, load that data.

IMPORTANT:
Make sure you update your import statements to import packages from the correct location!

## Solution

1. Start your server
You must use "go build" as you need to build a binary that includes a dependency (models package).

```bash
go build -o cowgirl
./cowgirl
```

### POST users

Enter these commands at the terminal

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"James Bond","gender":"male","age":32}' http://localhost:8080/user
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
curl -X POST -H "Content-Type: application/json" -d '{"name":"Q","gender":"male","age":54}' http://localhost:8080/user

# on windows curl:
curl -X POST -H "Content-Type: application/json" -d "{\"name\":\"Miss Moneypenny\",\"gender\":\"female\",\"age\":27}" http://localhost:8080/user
curl -X POST -H "Content-Type: application/json" -d "{\"name\":\"James Bond\",\"gender\":\"male\",\"age\":32}" http://localhost:8080/user
curl -X POST -H "Content-Type: application/json" -d "{\"name\":\"Q\",\"gender\":\"male\",\"age\":54}" http://localhost:8080/user
```

-X is short for --request
Specifies a custom request method to use when communicating with the HTTP server.

-H is short for --header

-d is short for --data

### GET a user

Enter this at the terminal

```bash
curl http://localhost:8080/user/<enter-user-id-here>
```

### DELETE a user

Enter this at the terminal

```bash
curl -X DELETE http://localhost:8080/user/<enter-user-id-here>
```
