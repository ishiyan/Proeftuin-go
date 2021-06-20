# How to Use //go:embed

[source](https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/)

## Version information

`//go:embed` gives us an easy way to include version information from a version.txt file.

```go
package main

import (
    _ "embed"
    "fmt"
    "strings"
)

var (
    Version string = strings.TrimSpace(version)
    //go:embed version.txt
    version string
)

func main() {
    fmt.Printf("Version %q\n", Version)
}
```

For a more complicated example, we can even include version information conditionally based on whether a certain build tag is passed to the go tools.

```go
// version_dev.go
// +build !prod

package main

var version string = "dev"
```

```go
// version_prod.go
// +build prod

package main

import (
    _ "embed"
)

//go:embed version.txt
var version string
```

```sh
$ go run .
Version "dev"

$ go run -tags prod .
Version "0.0.1"
```

## Quine

A quine is a program that prints out its own source code.
Let�s look at how easily we can write one with `//go:embed`:

```go
package main

import (
    _ "embed"
    "fmt"
)

//go:embed quine.go
var src string

func main() {
    fmt.Print(src)
}
```

When we run this program, it prints out itself.

## Embedding complex structs (GOB)

If we have some complex information we want to precompute, we can save that information in our Go project by writing it out into a .go file with `//go:generate` and templates, or we can save it to a serialization format understood by Go and load up the serialized data on start up:

```go
package main

import (
    "bytes"
    _ "embed"
    "encoding/gob"
    "fmt"
)

var (
    // File value.gob contains some complicated data
    // which we have precomputed and saved.
    //go:embed value.gob
    b []byte
    s = func() (s struct {
        Number   float64
        Weather  string
        Alphabet []string
    }) {
        dec := gob.NewDecoder(bytes.NewReader(b))
        if err := dec.Decode(&s); err != nil {
            panic(err)
        }
        return
    }()
)

func main() {
    fmt.Printf("s: %#v\n", s)
}
```

```sh
$ go run create.go

# observe file values.gob being created

$ go run decode.go
```

## Website files

This is probably going to be one of the biggest application for `//go:embed`.
We can now include all the static files or templates needed for our website in a single executable.
We can even toggle between reading files on disk and reading embedded files on the fly based on command line arguments:

```go
package main

import (
    "embed"
    "io/fs"
    "log"
    "net/http"
    "os"
)

func main() {
    useOS := len(os.Args) > 1 && os.Args[1] == "live"
    http.Handle("/", http.FileServer(getFileSystem(useOS)))
    http.ListenAndServe(":8888", nil)
}

//go:embed static
var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
    if useOS {
        log.Print("using live mode")
        return http.FS(os.DirFS("static"))
    }

    log.Print("using embed mode")
    fsys, err := fs.Sub(embededFiles, "static")
    if err != nil {
        panic(err)
    }

    return http.FS(fsys)
}
```

Note that we need to strip off the directory prefix from the `embed.FS` with `fs.Sub` so that it matches what is produced by the `os.DirFS`.

Here is another example that shows an embedded template:

```go
package main

import (
    "embed"
    "os"
    "text/template"
)

//go:embed *.tmpl
var tpls embed.FS

func main() {
    name := "en.tmpl"
    if len(os.Args) > 1 {
        name = os.Args[1] + ".tmpl"
    }
    arg := "World"
    if len(os.Args) > 2 {
        arg = os.Args[2]
    }

    t, err := template.ParseFS(tpls, "*")
    if err != nil {
        panic(err)
    }
    if err = t.ExecuteTemplate(os.Stdout, name, arg); err != nil {
        panic(err)
    }
}
```

With `en.tmpl` having the contents Hello `{{ . }}, how are you today?` and `jp.tmpl` having the contents `こんにちは{{ . }}。お元気ですか。`, we get this output:

```sh
$ go run ./main.go
Hello World, how are you today?

$ go run ./main.go jp
こんにちはWorld。お元気ですか。
```

## Gotchas

There are some gotchas with embedding to be aware of.
First of all, you must import the embed package in any file that uses an embed directive.
So a file like this won’t work:

```go
package main

import (
    "fmt"
)

//go:embed file.txt
var s string

func main() {
    fmt.Print(s)
}
```

```sh
$ go run missing-embed.go
# command-line-arguments
./missing-embed.go:8:3: //go:embed only allowed in Go files that import "embed"
```

On the other hand, the usual Go rules forbidding unused imports apply.
If you need to import embed but not refer to any exported identifiers in it, you should use `import _ "embed"` to tell Go to import embed even though it doesn’t look like it’s being used.

Second, you can only use `//go:embed` for variables at the package level, not within functions or methods, so a program like this won’t compile:

```go
package main

import (
    _ "embed"
    "fmt"
)

func main() {
    //go:embed file.txt
    var s string
    fmt.Print(s)
}
```

```sh
$ go run bad-level.go
# command-line-arguments
./bad-level.go:9:4: go:embed cannot apply to var inside func
```

Third, when you include a directory, it won’t include files that start with `.` or `_`, but if you use a wildcard, like `dir/*`, it will include all files that match, even if they start with `.` or `_`.
Bear in mind that accidentally including Mac OS’s `.DS_Store` files may be a security problem in circumstances where you want to embed files in a webserver but not allow users to see a list of all the files.
For security reasons, Go also won’t follow symbolic links or go up a directory when embedding.

There are no limits to the possible applications of embedding.
For example, what if you read the documentation and licensing info for your command line application out of the repo’s README file?
How about storing database queries for your application as embedded .sql files?
Or you could write an overlay FS to combine a built-in `embed.FS` with user-supplied override files…
These are just a few ideas scratching the surface.
I’m sure we’ll see a lot of clever and unexpected uses for it as time goes on.
