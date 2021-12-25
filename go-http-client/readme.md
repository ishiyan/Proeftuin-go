# How to develop a productive HTTP client in Go

## Working on the core

### Introduction to Go modules

- Go < 1.13 : No modules available. DEP was the dependency management system.
- G0 = 1.13 : Modules as BETA.
- Go > 1.14 : Modules are the standard for managing dependencies.

- Go < 1.8 : GOPATH must be defined.
- Go = 1.8 : Default GOPATH ~/go

```text
GOROOT
    Go installation folder

GOPATH
    Go workspaces
    src
        github.com
            username
                repo
    pkg
    bin
```

If you use `GoLand`, add the following to your `.gitignore` file.

```.gitignore
IDE folders
.idea
```

```bash
cd go-httpclient
go mod init github.com/federicoleon/go-httpclient
# now add dependencies to go.mod
go tidy
```

### Go basics: Structs, functions, interfaces and methods

### Adding basic behavior

### Defining custom & common headers

### Dealing with the request body

### Testing, testing and testing

### Be careful with code coverage

### Dealing with timeouts

### Allow timeout customization

### Allow timeout disabling

### Builder pattern applied

### Refactoring our builder implementation

### Making the client concurrent-safe

### Using our custom response implementation
