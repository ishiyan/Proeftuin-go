# Go embed and Angular

A showcase `Go 1.16` new embed package.
[taken from](https://shibumi.dev/posts/go-embed-and-angular/)
[and updated from](https://stackoverflow.com/questions/66600155/is-it-possible-to-embed-angular-inside-golang-application)

The project consists of two directories.

- server: includes all Go code
- webapp: includes the Angular App

To build Angular app, do the following.

```bash
cd webapp
npm install
./node_modules/.bin/ng build --prod
```

I have changed the output path for the Angular generated assets in the `angular.json` file via setting

```json
 "outputPath": "../server/static"
```

With this change Angular will move the static assets to a new static directory inside of the server directory.
Why do we need this?
We need this, because the embed package does not support `../`, `./` or leading slashes,
hence we cannot import data from the webapp.
One possible solution is to place a Go file in the webapp directory, but I have not tried this.

The `app.go` file.

```go
// First, I introduced a new global variable
//go:embed static
var static embed.FS

// .....

// Then I modified the start() function accordingly:
func (a *App) start() {
  a.db.AutoMigrate(&student{})
  a.r.HandleFunc("/students", a.getAllStudents).Methods("GET")
  a.r.HandleFunc("/students", a.addStudent).Methods("POST")
  a.r.HandleFunc("/students/{id}", a.updateStudent).Methods("PUT")
  a.r.HandleFunc("/students/{id}", a.deleteStudent).Methods("DELETE")

  // We need to strip the static directory from our path
  // for serving files in the index folder via the http.Fileserver()
  webapp, err := fs.Sub(static, "static")
  if err != nil {
    fmt.Println(err)
  }

  // We need to use Gorilla Mux' PathPrefix function here, because the Pathprefix
  // adds a wildcard to the route eg: /*, otherwise we would only route to "/"
  // Hence the error with 404-returning JS files before got thrown, because
  // Gorilla Mux had no route to these JS files.
  a.r.PathPrefix("/").Handler(http.FileServer(http.FS(webapp)))
  log.Fatal(http.ListenAndServe(":8080", a.r))
}
```

What is happening here? First I introduced a new global variable called `static` with type `embed.FS`.
The important part about this change is the go preprocessor-like statement before the variable declaration.
With `//go:embed static` we explain the Go compiler to embed the static directory in the current directory via the embed package.
Note: the missing space between `//` and `go` is important here!
The next modification is the `start()` function.
We are now serving content from a directory, for example `static/index.html`,
thus we need to strip the `static` directory name from it.
This happens via the `fs.Sub` method. The last change is the use of `http.FS` instead of `http.Dir`.
We are dealing with a filesystem now, not a local directory anymore.
If we now compile the Angular app and compile our Go binary the Angular generated assets will get included into our Go binary and we have a single binary for deployment.

The new `Dockerfile` makes use of Google�s distroless docker image.
Distroless images are basically like docker scratch images, with the difference that they provide tzdata and ca-certificates and other data applications might need.
Everything else (libraries, shells, busybox utils, etc) is missing in these images.
The final Dockerfile looks like this.

```dockerfile
FROM node:12.11 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@8.3.12
COPY webapp /webapp
WORKDIR webapp
RUN npm install && ng build --prod

FROM golang:1.16 as GO_BUILD
WORKDIR /go/src/app
ADD server /go/src/app
COPY --from=ANGULAR_BUILD /server/static /go/src/app
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=GO_BUILD /go/bin/app /
CMD ["/app"]
```

## CGO

If you are on Windows, you have to install `gcc` otherwise `go-sqlite3` won't compile.

```bash
go get github.com/mattn/go-sqlite3
# github.com/mattn/go-sqlite3
exec: "gcc": executable file not found in %PATH%
```

Follow the steps mentioned below.

- [Download the GCC](http://tdm-gcc.tdragon.net/download)
- Install (check setting to add to the PATH)
- Restart the shell where you are building.

## Build

```bash
# build webapp first, the output will be in the static folder

# copy app.go.outside to app.go
go build -o ng_outside.exe .

# copy app.go.embedded to app.go
go build -o ng_embedded.exe .
```
