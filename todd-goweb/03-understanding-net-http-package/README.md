# Understanding net/http package

## Handler - this is one of the most important things to know

[http.Handler](https://godoc.org/net/http#Handler)

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

## Server - notice that both of the functions take a handler

[http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)

```go
func ListenAndServe(addr string, handler Handler) error
```

[http.ListenAndServeTLS](https://godoc.org/net/http#ListenAndServeTLS)

```go
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error
```

## Request

See [http.Request](https://godoc.org/net/http#Request) in the documentation.

Here it is with *most of the comments and some of the fields* stripped out:

```go
type Request struct {
    Method string
    URL *url.URL
    //  Header = map[string][]string{
    //      "Accept-Encoding": {"gzip, deflate"},
    //      "Accept-Language": {"en-us"},
    //      "Foo": {"Bar", "two"},
    //  }
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    // This field is only available after ParseForm is called.
    Form url.Values
    // This field is only available after ParseForm is called.
    PostForm url.Values
    MultipartForm *multipart.Form
    // RemoteAddr allows HTTP servers and other software to record
    // the network address that sent the request, usually for
    // logging.
    RemoteAddr string
}
```

Also see the index showing type `[Request]()` from the http package.

Some interesting things you can do with a request:

### Retrieve URL & Form data

`http.Request` is a struct. It has the fields `Form` & `PostForm`. If we read the documentation on these, we see:

```go
    // Form contains the parsed form data, including both the URL
    // field's query parameters and the POST or PUT form data.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores Form and uses Body instead.
    Form url.Values

    // PostForm contains the parsed form data from POST, PATCH,
    // or PUT body parameters.
    // This field is only available after **ParseForm** is called.
    // The HTTP client ignores PostForm and uses Body instead.
    PostForm url.Values
```

If we look at **ParseForm**,

```go
func (r *Request) ParseForm() error
```

we see that this is a method attached to a *http.Request.

If we look at **FormValue**,

```go
func (r *Request) FormValue(key string) string
```

we see that this is a method attached to a `*http.Request`. `FormValue` returns the first value for the named component of the query. POST and PUT body parameters take precedence over URL query string values. `FormValue` calls `ParseMultipartForm` and `ParseForm` if necessary and ignores any errors returned by these functions. If key is not present, `FormValue` returns the empty string. To access multiple values of the same key, call `ParseForm` and then inspect `Request.Form` directly.

### See the HTTP Method

The `http.Request` type is a struct which has a `Method` field.

### See URL values

The `http.Request` type is a struct which has a `URL` field. Notice that the type is a `*url.URL`

Take a look at type `url.URL`

```go
type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string
    RawPath    string // encoded path hint (Go 1.5 and later only; see EscapedPath method)
    ForceQuery bool   // append a query ('?') even if RawQuery is empty
    RawQuery   string // encoded query values, without '?'
    Fragment   string // fragment for references, without '#'
}
```

### Work with the HTTP header

The `http.Request` type is a struct which has a `Header` field.

From the documentation about the `http.Header` struct, we see that:

```go
type Header map[string][]string
```

## Response

[http.ResponseWriter](https://godoc.org/net/http#ResponseWriter)

```go
type ResponseWriter interface {
    // Header returns the header map that will be sent by
    // WriteHeader. Changing the header after a call to
    // WriteHeader (or Write) has no effect
    Header() Header

    // Write writes the data to the connection as part of an HTTP reply.
    //
    // If WriteHeader has not yet been called, Write calls
    // WriteHeader(http.StatusOK) before writing the data. If the Header
    // does not contain a Content-Type line, Write adds a Content-Type set
    // to the result of passing the initial 512 bytes of written data to
    // DetectContentType.
    Write([]byte) (int, error)

    // WriteHeader sends an HTTP response header with status code.
    // If WriteHeader is not called explicitly, the first call to Write
    // will trigger an implicit WriteHeader(http.StatusOK).
    // Thus explicit calls to WriteHeader are mainly used to
    // send error codes.
    WriteHeader(int)
}
```

### Setting a response header

An `http.ResponseWriter` has a method `Header()` which returns a `http.Header`.

Look at the documentation for `http.Header`

```go
type Header map[string][]string
```

Look at the methods which are attached to type `http.Header`

```go
type Header
func (h Header) Add(key, value string)
func (h Header) Del(key string)
func (h Header) Get(key string) string
func (h Header) Set(key, value string)
func (h Header) Write(w io.Writer) error
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```

We can set headers for a response like this:

```go
res.Header().Set("Content-Type", "text/html; charset=utf-8")
```

## Web Programming Synonymous Terms

- router
- request router
- multiplexer
- mux
- servemux
- server
- http router
- http request router
- http multiplexer
- http mux
- http servemux
- http server

In electronics, a [multiplexer (or mux)](https://en.wikipedia.org/wiki/Multiplexer) is a device that selects one of several input signals and forwards the selected input into a single line.

The term multiplexer has been adopted by web programming to refer to the process of routing requests.

A web server has requests coming in at different routers and via different HTTP methods. For instance, we might have these requests:

REQUEST #1

- Path: /cat
- Method: GET

REQUEST #2

- Path: /apply
- Method: Get

Request #3

- Path: /apply
- Method: Post

Based upon the requests coming in, the server needs to determine how to respond to that request - for each request that comes in, different code will be run.

I've been using the word "server" but I could have also been using the word "multiplexer" or "mux". The server, or multiplexer, or mux, determines what code needs to be run in response to each incoming request

***

ServeMux is an HTTP request multiplexer.

A ServeMux matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

Patterns name fixed, rooted paths, like "/favicon.ico", or rooted subtrees, like "/images/" (note the trailing slash).

Longer patterns take precedence over shorter ones, so that if there are handlers registered for both "/images/" and "/images/thumbnails/", the latter handler will be called for paths beginning "/images/thumbnails/" and the former will receive requests for any other paths in the "/images/" subtree. Note that since a pattern ending in a slash names a rooted subtree, the pattern "/" matches all paths not matched by other registered patterns, not just the URL with Path == "/".

If a subtree has been registered and a request is received naming the subtree root without its trailing slash, ServeMux redirects that request to the subtree root (adding the trailing slash). This behavior can be overridden with a separate registration for the path without the trailing slash. For example, registering "/images/" causes ServeMux to redirect a request for "/images" to "/images/", unless "/images" has been registered separately.

Patterns may optionally begin with a host name, restricting matches to URLs on that host only. Host-specific patterns take precedence over general patterns, so that a handler might register for the two patterns "/codesearch" and "codesearch.google.com/" without also taking over requests for "http://www.google.com/".

ServeMux also takes care of sanitizing the URL request path, redirecting any request containing . or .. elements or repeated slashes to an equivalent, cleaner URL.

***

## ServeMux

[http.ServeMux](https://godoc.org/net/http#ServeMux)

```go
type ServeMux
    func NewServeMux() *ServeMux
    func (mux *ServeMux) Handle(pattern string, handler Handler)
    func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
    func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

Any value of type `*http.ServeMux` implements the `http.Handler` interface.

Remember, the `http.Handler` interface requires that a type have the `ServeHTTP` method.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

What this tells us is that we can pass a value of type `*http.ServeMux` into `http.ListenAndServe`

```go
func ListenAndServe(addr string, handler Handler) error
```

You can also see from the documentation of `*http.ServeMux` ...

```go
type ServeMux
    func NewServeMux() *ServeMux
    func (mux *ServeMux) Handle(pattern string, handler Handler)
    func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
    func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

... that we have a method `Handle` attached the the value of type `*http.ServeMux`. You can see that `Handle` takes a pattern and a `http.Handler`.

We can use `Handle` like this:

```go
    mux := http.NewServeMux()
    mux.Handle("/", h)
    mux.Handle("/cat", c)
```

The overall game plan:

We will create a NewServeMux, then attach the method `Handle` to it to set routes, then pass our `*http.ServeMux` to `http.ListenAndServe`.

## DefaultServeMux

ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:

```go
http.ListenAndServe(":8080", nil)
```

## HandleFunc

[http.HandleFunc](https://godoc.org/net/http#HandleFunc)

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

## HandlerFunc

[http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)

```go
type HandlerFunc func(ResponseWriter, *Request)
```

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

**This is just a nice thing to know about. You wouldn't do this in production code probably.**

## Third-party ServeMux

You can search [godoc.org](https://godoc.org/) for third-party packages.

Here is [a good third-party ServeMux](https://godoc.org/github.com/julienschmidt/httprouter) that allows easy access to methods for routing & path parameters.

### [julienschmidt/httprouter](https://godoc.org/github.com/julienschmidt/httprouter)

#### Match method & path

The router matches incoming requests by the request method and the path.

```go
 func main() {
     router := httprouter.New()
     router.GET("/apply", apply)
     router.POST("/apply", applyProcess)
     http.ListenAndServe(":8080", router)
 }
 ```

#### Named path parameters

The registered path, against which the router matches incoming requests, can also contain parameters. Parameters are dynamic path segments. They match anything until the next '/' or the path end.

```go
func main() {
    router := httprouter.New()
    router.GET("/blog/:category/:article", blog)
    http.ListenAndServe(":8080", router)
}
```

```text
Requests:
 /blog/go/request-routers            match: category="go", article="request-routers"
 /blog/go/request-routers/           no match, but the router would redirect
 /blog/go/                           no match
 /blog/go/request-routers/comments   no match
```

#### Catch-all path parameters

Catch-all parameters match anything until the path end, including the directory index (the '/' before the catch-all). Since they match anything until the end, catch-all parameters must always be the final path element.

```go
func main() {
    router := httprouter.New()
    router.GET("/files/*filepath", loadFile)
    http.ListenAndServe(":8080", router)
}
```

```text
Requests:
 /files/                             match: filepath="/"
 /files/LICENSE                      match: filepath="/LICENSE"
 /files/templates/article.html       match: filepath="/templates/article.html"
 /files                              no match
```

#### Using path parameters

The value of parameters is saved as a []Param

```go
type Param struct {
    Key   string
    Value string
}
```

The slice is passed to the Handle func as a third parameter.

```go
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("user"))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:user", Hello)

    http.ListenAndServe(":8080", router)
}
```

Retrieve the value of a parameter:

```go
user := ps.ByName("user") // defined by :user or *user
```

#### Performance

Julien Schmidt's router is [nicely performant](https://github.com/julienschmidt/go-http-routing-benchmark#static-routes)
