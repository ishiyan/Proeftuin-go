# Self Contained Blog Server with Go 1.16

[code](https://github.com/bketelsen/bktw)

Create a web server distributed as a static binary with all resources embedded in it, including

- HTML Templates
- CSS Files
- Template images
- Javascript files

To challenge myself a bit, I wanted to use [Tailwind CSS](https://tailwindcss.com/), a CSS framework that breaks from the traditional framework model and provides "utility first" classes.
Since I didn't really know what any of that meant when I started this project I was excited to learn more about Tailwind to see if I could understand why it's gaining popularity so rapidly.

This isn't an article about HTML or CSS though, so we're not going to be hand-crafting artisanal web pages.
Instead I found an [HTML template](https://github.com/tailwindtoolbox/Ghostwind) that mimics the [Ghost](https://ghost.org/) blogging platform's Casper theme, but uses Tailwind to style it.

## Getting Started

Because Tailwind requires some compilation and processing, I decided to start with the CSS side of things first.
I followed the [Getting Started](https://tailwindcss.com/docs/installation) guide on the Tailwind site, which guided me through setting up Tailwind CSS compilation using autoprefixer and postcss.
Let's not pretend that I have a firm grasp on what those things do.

After setting up CSS compilation I had a few configuration files and a src/css/main.css file that imported the Tailwind base classes.

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

Because I'm old-school, and we're mixing Node.js and Go in this project, I created a Makefile to manage the steps and dependencies required to build the app.
The target to build the CSS is here:

```makefile
assets/css/main.css: src/css/main.css
    NODE_ENV=production npx tailwindcss-cli@latest build ./src/css/main.css -o ./assets/css/tailwind.css 
```

Translation: The output file "assets/css/main.css" depends on the input file "src/css/main.css" and can be built with the command "NODE_ENV=production npx tailwindcss-cli@latest build ./src/css/main.css -o ./assets/css/tailwind.css"

I took the build command straight from the Tailwind installation documentation.

## Go Serve Things

To build the web server I ran 'go mod init' in the same directory to create a Go project.
I decided to use [Gin](https://github.com/gin-gonic/gin) as the web server because I hadn't previously used it and the API is lightweight and pleasant to work with.

To initialize a Gin server, you first declare a router:

```go
router := gin.Default()
```

I read the Gin documentation and examples and concluded that I'd need to use Go's 'html/template' package if I was going to take advantage of an embedded filesystem for templates.
So I created a 'templates' folder and put my HTML templates in it.
In the Gin app, I created a global variable called 'f' that represents the data of the filesystem I intend to embed in the compiled binary:

```go
//go:embed assets/* templates/*
var f embed.FS
```

The variable declaration is preceded by a magic comment that instructs the Go compiler to embed the contents of the 'assets' and 'templates' directory as a Filesystem into the 'f' variable.

In 'main' I create a template set by reading the templates from the embedded filesystem:

```go
templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
router.SetHTMLTemplate(templ)
```

Then I set up the embedded 'assets' directory to be served under the '/public' path:

```go
// example: /public/assets/images/example.png
router.StaticFS("/public", http.FS(f)
```

The rest of the app is standard plumbing... fetching data from the Strapi content server and making it available to the templates.
I'll let you read the source code for the implementation of those parts, since they aren't novel to this server.
