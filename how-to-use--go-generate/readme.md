# How to Use `//go:generate`

The go generate command was added in Go 1.4, "to automate the running of tools to generate source code before compilation."

Before explaining different ways to use it, let's look at a simple alternative to using the `go:generate` directive that is also built into the Go tool.

## Using `-ldflags`

Suppose you wrote a program and you want it to print out its version number.
How could you bake the version number into the executable?
One way might be to just remember that the version is stored on line X of the source code and manually update that line periodically, but that's not very automated.

Fortunately, the Go toolchain comes with a way of overriding individual string values.
The link tool has the following command line flag: `-X importpath.name=value` which will "Set the value of the string variable in importpath named name to value."
Of course, most of the time you program in Go, you're not using the link tool directly.
Instead, you're either running `go run`, `go build`, or `go install`.
Fortunately, those commands can pass a flag to their internal use of the Go link tool with the flag `-ldflags`. (Yes, you pass a flag to pass a flag.)

Here's a quick example. Given this content in `main.go`:

```go
package main

import (
    "fmt"
)

var VersionString = "unset"

func main() {
    fmt.Println("Version:", VersionString)
}
```

We can use `ldflags` to override `VersionString`:

```sh
$ go run main.go
Version: unset

$ go run -ldflags '-X main.VersionString=1.0' main.go
Version: 1.0
```

Even more useful, we can use Git’s commit hash as an automatic version number:

```sh
$ git rev-parse HEAD
db5c7db9fe3b632407e9f0da5f7982180c438929

$ go run -ldflags "-X main.VersionString=`git rev-parse HEAD`" main.go
Version: db5c7db9fe3b632407e9f0da5f7982180c438929
```

Similarly, you can use the standard Unix `date` command to pass the build time into a file.
One confusing point is that, since we’re passing a flag to set a flag, if you want to pass an argument with a space in it, you have to use quotes inside of quotes:

```sh
$ go run -ldflags "-X 'main.VersionString=1.0 (beta)'" main.go
Version: 1.0 (beta)

$ go run -ldflags "-X 'main.VersionString=`date`'" main.go
Version: Sun Nov 27 16:42:10 EST 2016
```

The more linker flags we want to pass, the more complicated our Go build commands will be, so they will need to be saved somewhere to standardize the build process, perhaps in a simple Bash script file or a Makefile.

## `go:generate` basics

Makefiles and Bash scripts are time-tested means of code development, but if you’re a Go programmer, you are an expert in Go and not necessarily an expert in Makefiles and Bash scripts.
Wouldn’t it be nice if we could automate our build process using the Go tool alone?
This was the problem that `go:generate` was built to solve.

What follows is a brief description of `go:generate`.
Please [consult the docs](https://golang.org/cmd/go/#hdr-Generate_Go_files_by_processing_source) for a full specification of the flags and options it accepts.
When you run the command `go generate`, the Go tool scans the files relevant to the current package for lines with a “magic comment” of the form :

```go
//go:generate command arguments
```

This command does not have to do anything related to Go or code generation.

```go
package project

//go:generate echo Hello, Go Generate!

func Add(x, y int) int {
    return x + y
}
```

```sh
$ go generate
Hello, Go Generate!
```

Note that while `!` is a special character in Bash, Go ignores it and just processes the command normally.

There are already a number of tools available that are designed to be run by the go:generate directive, such as [stringer](https://godoc.org/golang.org/x/tools/cmd/stringer), [jsonenums](https://github.com/campoy/jsonenums), and [schematyper](https://github.com/idubinskiy/schematyper).

## Using `go:generate` to run a Go program

As mentioned, the command run by `go:generate` does not have to meet any particular requirements.
This may lead to a situation where building some Go code requires running go generate first, but running go generate requires some third-party dependency that may not be installed.
We can probably rely on other computers having Git and date installed, but they may not have other tools, particularly if the tool in question is one we wrote ourselves.
To work around this, the Go team recommends that you distribute your generated source files along with your handwritten source files, but this can be inconvenient or impractical in some cases.

One solution to this dependency problem is to have go generate use go run to run a Go program that we create.
Because we know that any computer building Go must have Go installed, we don’t have to worry about pre-build dependencies.

Here is a concrete example.
Supposed I want my binary to output a hard coded list of contributors to the Go project with “Carl” in their name.
To automate the building of that binary, I can have go generate download the list of contributors from Github and output that to a Go source file in the project. Here is what that could look like.

```txt
    `-- cmd
    |   `-- carls
    |       `-- main.go
    |-- contributors.go
    |-- gen.go
    `-- project.go
```

`cmd/carls/main.go` is separate just to make the example more clear.
All of the code it relies on is imported:

```go
package main

import (
    "project"
)

func main() {
    project.PrintContributors()
}
```

`project.go` is where the `//go:generate` directive is, as well as the active code for this project:

```go
package project

import "fmt"

//go:generate go run gen.go

func PrintContributors() {
    for _, contributor := range Contributors {
        fmt.Println(contributor)
    }
}
```

When you’re using code generation, you want to be sure to keep the code that needs to be generated separate from the code that doesn’t need to be generated, so that the revision history of the project will be clear for future maintainers.

`gen.go` is run by go generate and does the downloading of contributors and outputting their names into a Go file.
Because all Go files in a given directory need to belong to the same package, we will have to add a build constraint magic comment to the file we run to prevent it from being considered a part of the package.

```go
// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates contributors.go. It can be invoked by running
// go generate
package main

import (
    "bufio"
    "log"
    "net/http"
    "os"
    "strings"
    "text/template"
    "time"
)

func main() {
    const url = "https://github.com/golang/go/raw/master/CONTRIBUTORS"

    rsp, err := http.Get(url)
    die(err)
    defer rsp.Body.Close()

    sc := bufio.NewScanner(rsp.Body)
    carls := []string{}

    for sc.Scan() {
        if strings.Contains(sc.Text(), "Carl") {
            carls = append(carls, sc.Text())
        }
    }

    die(sc.Err())

    f, err := os.Create("contributors.go")
    die(err)
    defer f.Close()

    packageTemplate.Execute(f, struct {
        Timestamp time.Time
        URL       string
        Carls     []string
    }{
        Timestamp: time.Now(),
        URL:       url,
        Carls:     carls,
    })
}

func die(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
// using data from
// {{ .URL }}
package project

var Contributors = []string{
{{- range .Carls }}
    {{ printf "%q" . }},
{{- end }}
}
`))
```

Another feature of `gen.go` is that it uses a timestamp to make it clear when the file was machine generated.
Using a text template provides for easy extension of the file if the need should arise in the future.

Finally, `contributors.go` is the machine generated file:

```go
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2021-04-20 15:27:35.532276 +0200 CEST m=+0.333478401
// using data from
// https://github.com/golang/go/raw/master/CONTRIBUTORS
package project

var Contributors = []string{
    "Carl Chatfield <carlchatfield@gmail.com>",
    "Carl Henrik Lunde <chlunde@ifi.uio.no>",
    "Carl Jackson <carl@stripe.com>",
    "Carl Johnson <me@carlmjohnson.net>",
    "Carl Mastrangelo <notcarl@google.com>",
    "Carl Shapiro <cshapiro@google.com> <cshapiro@golang.org>",
    "Carlisia Campos <carlisia@grokkingtech.io>",
    "Carlo Alberto Ferraris <cafxx@strayorange.com>",
    "Carlos Alexandro Becker <caarlos0@gmail.com>",
    "Carlos Amedee <carlos@golang.org>",
    "Carlos Castillo <cookieo9@gmail.com>",
    "Carlos Cirello <uldericofilho@gmail.com>",
    "Carlos Eduardo <me@carlosedp.com>",
    "Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>",
    "Carlos Iriarte <ciriarte@gmail.com>",
    "Carlos Souza <carloshrsouza@gmail.com>",
    "David Carlier <devnexen@gmail.com>",
    "Dustin Carlino <dcarlino@google.com>",
    "Juan Carlos <juanjcsr@gmail.com>",
}
```

As of Go 1.9, there is a standard format for comments to designate a file as machine generated.
A generated file should have a comment matching this regex: `^// Code generated .* DO NOT EDIT.$`.

To keep your generating code from being mistaken for generated code, make sure the magic comment isn’t at the start of a line.

Now we can put it all together and run it:

```sh
$ go generate

$ go run cmd/carls/main.go
Carl Chatfield <carlchatfield@gmail.com>
Carl Henrik Lunde <chlunde@ifi.uio.no>
Carl Jackson <carl@stripe.com>
Carl Johnson <me@carlmjohnson.net>
Carl Mastrangelo <notcarl@google.com>
Carl Shapiro <cshapiro@google.com> <cshapiro@golang.org>
Carlisia Campos <carlisia@grokkingtech.io>
Carlo Alberto Ferraris <cafxx@strayorange.com>
Carlos Alexandro Becker <caarlos0@gmail.com>
Carlos Amedee <carlos@golang.org>
Carlos Castillo <cookieo9@gmail.com>
Carlos Cirello <uldericofilho@gmail.com>
Carlos Eduardo <me@carlosedp.com>
Carlos Eduardo Seo <cseo@linux.vnet.ibm.com>
Carlos Iriarte <ciriarte@gmail.com>
Carlos Souza <carloshrsouza@gmail.com>
David Carlier <devnexen@gmail.com>
Dustin Carlino <dcarlino@google.com>
Juan Carlos <juanjcsr@gmail.com>
```
