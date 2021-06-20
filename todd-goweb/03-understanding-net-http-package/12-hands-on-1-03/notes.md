# Hands-on

Take the previous program and change it so that func main uses http.Handle instead of http.HandleFunc

Contstraint: Do not change anything outside of func main

Hints:

[http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)

```go
type HandlerFunc func(ResponseWriter, *Request)
```

[http.HandleFunc](https://godoc.org/net/http#HandleFunc)

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

[source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)

```go
  func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
      mux.Handle(pattern, HandlerFunc(handler))
  }
```
