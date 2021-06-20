# Techniques we will learn

- concatenate
- CLI pipeline - output to a file with >

## Code we will use from the standard library

### [os.Create](https://godoc.org/os#Create)

This allows us to create a file.

```go
func Create(name string) (*File, error)
```

***

### [defer](https://golang.org/ref/spec#Defer_statements)

The defer keyword allows us to defer the execution of a statement until the function in which we have placed the defer statement returns.

***

### [io.Copy](https://godoc.org/io#Copy)

This allows us to copy from from a source to a destination.

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

### [strings.NewReader](https://godoc.org/strings#NewReader)

NewReader returns a new Reader reading from s.

```go
func NewReader(s string) *Reader
```

### [os.Args](https://godoc.org/os#pkg-variables)

Args is a variable from package os. Args hold the command-line arguments, starting with the program name.

```go
var Args []string
```

## Type Template

### [template.Template](https://godoc.org/text/template#Template)

```go
template.Template
```

***

## Parsing templates

### [template.ParseFiles](https://godoc.org/text/template#ParseFiles)

```go
func ParseFiles(filenames ...string) (*Template, error)
```

### [template.ParseGlob](https://godoc.org/text/template#ParseGlob)

```go
func ParseGlob(pattern string) (*Template, error)
```

***

### [template.Parse](https://godoc.org/text/template#Template.Parse)

```go
func (t *Template) Parse(text string) (*Template, error)
```

### [template.ParseFiles](https://godoc.org/text/template#Template.ParseFiles)

```go
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```

### [template.ParseGlob](https://godoc.org/text/template#Template.ParseGlob)

```go
func (t *Template) ParseGlob(pattern string) (*Template, error)
```

***

## Executing templates

### [template.Execute](https://godoc.org/text/template#Template.Execute)

```go
func (t *Template) Execute(wr io.Writer, data interface{}) error
```

### [template.ExecuteTemplate](https://godoc.org/text/template#Template.ExecuteTemplate)

```go
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```

***

## Helpful template functions

### [template.Must](https://godoc.org/text/template#Must)

```go
func Must(t *Template, err error) *Template
```

### [template.New](https://godoc.org/text/template#New)

```go
func New(name string) *Template
```

***

## The init function

### [The init function](https://golang.org/doc/effective_go.html#init)

```go
func init()
```

## Passing Data To Templates

You get to pass in one value - that's it!

Fortunately, we have many different types which that value can be including composite types which compose together values. (These are also known as aggregate data types - they aggregate together many different values).

### Slice

Use this for passing in a bunch of values of the same type. We could have a []int or a []string or a slice of any type.

### Map

Use this for passing in key-value data.

### Struct

This is probably the most commonly used data type when passing data to templates. A struct allows you to compose together values of different types.

## Template variables

### [template variables](https://godoc.org/text/template#hdr-Variables)

#### ASSIGN

```go
{{$wisdom := .}}
```

#### USE

```go
{{$wisdom}}
```

A pipeline inside an action may initialize a variable to capture the result. The initialization has syntax

`$variable := pipeline`

where $variable is the name of the variable. An action that declares a variable produces no output.

If a "range" action initializes a variable, the variable is set to the successive elements of the iteration. Also, a "range" may declare two variables, separated by a comma:

`range $index, $element := pipeline`

in which case $index and $element are set to the successive values of the array/slice index or map key and element, respectively. Note that if there is only one variable, it is assigned the element; this is opposite to the convention in Go range clauses.

A variable's scope extends to the "end" action of the control structure ("if", "with", or "range") in which it is declared, or to the end of the template if there is no such control structure. A template invocation does not inherit variables from the point of its invocation.

When execution begins, `$` is set to the data argument passed to Execute, that is, to the starting value of dot.

## Using functions in templates

### [template function documentation](https://godoc.org/text/template#hdr-Functions)

***

### [template.FuncMap](type FuncMap map[string]interface{})

FuncMap is the type of the map defining the mapping from names to functions. Each function must have either a single return value, or two return values of which the second has type error. In that case, if the second (error) return value evaluates to non-nil during execution, execution terminates and Execute returns that error.

### [template.Funcs](https://godoc.org/text/template#Template.Funcs)

```go
func (t *Template) Funcs(funcMap FuncMap) *Template
```

***

During execution functions are found in two function maps:

- first in the template,
- then in the global function map.

By default, no functions are defined in the template but the Funcs method can be used to add them.

Predefined global functions are defined in text/template.

## Formatting with time

Since MST is GMT-0700, the reference time can be thought of as:

01/02 03:04:05PM '06 -0700

(January 02, 2006)

[godoc.org/time](https://godoc.org/time#pkg-constants)

***

### func (Time) Format

```go
func (t Time) Format(layout string) string
```

Format returns a textual representation of the time value formatted according to layout, which defines the format by showing how the reference time, defined to be

```go
Mon Jan 2 15:04:05 -0700 MST 2006
```

would be displayed if it were the value; it serves as an example of the desired output. The same display rules will then be applied to the time value.

A fractional second is represented by adding a period and zeros to the end of the seconds section of layout string, as in "15:04:05.000" to format a time stamp with millisecond precision.

Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard and convenient representations of the reference time. For more information about the formats and the definition of the reference time, see the documentation for ANSIC and the other constants defined by this package.

### MST

The Mountain Time Zone of North America keeps time by subtracting seven hours from Coordinated Universal Time (UTC), during the shortest days of autumn and winter (UTC−7), and by subtracting six hours during daylight saving time in the spring, summer, and early autumn (UTC−6). The clock time in this zone is based on the mean solar time at the 105th meridian west of the Greenwich Observatory. In the United States, the exact specification for the location of time zones and the dividing lines between zones is set forth in the Code of Federal Regulations at 49 CFR 71.[a]

In the United States and Canada, this time zone is generically called Mountain Time (MT). Specifically, it is Mountain Standard Time (MST) when observing standard time (fall and winter), and Mountain Daylight Time (MDT) when observing daylight saving time (spring and summer). The term refers to the fact that the Rocky Mountains, which range from northwestern Canada to the US state of New Mexico, are located almost entirely in the time zone.

### UTC

Coordinated Universal Time, abbreviated as UTC, is the primary time standard by which the world regulates clocks and time. It is within about 1 second of mean solar time at 0° longitude; it does not observe daylight saving time. It is one of several closely related successors to Greenwich Mean Time (GMT). For most purposes, UTC is considered interchangeable with GMT, but GMT is no longer precisely defined by the scientific community.

### GMT

Greenwich Mean Time (GMT) is the mean solar time at the Royal Observatory in Greenwich, London.

GMT was formerly used as the international civil time standard, now superseded in that function by Coordinated Universal Time (UTC).

Today GMT is considered equivalent to UTC for UK civil purposes (but this is not formalised) and for navigation is considered equivalent to UT1 (the modern form of mean solar time at 0° longitude); these two meanings can differ by up to 0.9 s. Consequently, the term GMT should not be used for precise purposes.

Because of Earth's uneven speed in its elliptical orbit and its axial tilt, noon (12:00:00) GMT is rarely the exact moment the sun crosses the Greenwich meridian and reaches its highest point in the sky there. This event may occur up to 16 minutes before or after noon GMT, a discrepancy calculated by the equation of time. Noon GMT is the annual average (i.e., "mean") moment of this event, which accounts for the word "mean" in "Greenwich Mean Time".

## Global Functions

There are "predefined global functions" which you can use.

[You can read about these functions here](https://godoc.org/text/template#hdr-Functions)

The following code samples will demonstrate some of these "predefined global functions":

- index
- and
- comparison

 ***

## [Template Comments](https://godoc.org/text/template#hdr-Actions)

A comment; discarded. May contain newlines. Comments do not nest and must start and end at the delimiters, as shown here.

```go
{{/* a comment */}}
```

## Nested templates

[nested templates documentation](https://godoc.org/text/template#hdr-Nested_template_definitions)

### define

```go
{{define "TemplateName"}}
insert content here
{{end}}
```

### use

```go
{{template "TemplateName"}}
```

## Passing data to templates

These files provide you with more examples of passing data to templates.

These files use the [composition](https://en.wikipedia.org/wiki/Composition_over_inheritance) design pattern. You should favor this design pattern.

Read more about [composition with Go here](https://www.goinggo.net/2015/09/composition-with-go.html).

## Take-away

One of the main take-aways is to use a composite data type. Often this data type will be a struct. Build a struct to hold the different pieces of data you'd like to pass to your template, then pass that to your template.

## Cross-site scripting (XSS)

Cross-site scripting (XSS) is a type of computer security vulnerability typically found in web applications.

XSS enables attackers to inject client-side scripts into web pages viewed by other users.

A cross-site scripting vulnerability may be used by attackers to bypass access controls such as the [same-origin policy](https://en.wikipedia.org/wiki/Same-origin_policy): you have a script on one site that makes a request to another site. For example: you come to my cool website about kittens, and a script runs to transfer money from UnionBank to my foreign account. If it wasn't for the "same-origin policy" implemented in browsers, and if you had a cookie on your machine that said you were logged into Union Bank, then the money would transfer.

Cross-site scripting carried out on websites accounted for roughly 84% of all security vulnerabilities documented by Symantec as of 2007. Their effect may range from a petty nuisance to a significant security risk, depending on the sensitivity of the data handled by the vulnerable site and the nature of any security mitigation implemented by the site's owner.

***

### Same-origin policy

In computing, the same-origin policy is an important concept in the web application security model.

Under the policy, a web browser permits scripts contained in a first web page to access data in a second web page, but only if both web pages have the same origin.

An origin is defined as a combination of URI scheme, hostname, and port number.

This policy prevents a malicious script on one site from obtaining access to sensitive data on another site.

### Example

Assume a user is visiting a banking website and doesn't log out.

Then he goes to another site and that site has some malicious JavaScript code running in the background that requests data from the banking site.

Because the user is still logged in on the banking site, without the "same-origin policy" implemented in browsers, that malicious code could do anything on the banking site.

For example, get a list of your last transactions, create a new transaction, etc. This is because the browser can send and receive session cookies to the banking website based on the domain of the banking website. A user visiting that malicious site would expect that the site he is visiting has no access to the banking session cookie. While this is true, the JavaScript has no direct access to the banking session cookie, but it could still send and receive requests to the banking site with the banking site's session cookie, essentially acting as a normal user of the banking site!

Regarding the sending of new transactions, even CSRF (cross-site request forgery) protections by the banking site have no effect, because the script can simply do the same as the user would do.

So this is a concern for all sites where you use sessions and/or need to be logged in.

### All modern browsers implement some form of the Same-Origin Policy as it is an important security cornerstone

This mechanism bears a particular significance for modern web applications that extensively depend on HTTP cookies to maintain authenticated user sessions, as servers act based on the HTTP cookie information to reveal sensitive information or take state-changing actions.

A strict separation between content provided by unrelated sites must be maintained on the client-side to prevent the loss of data confidentiality or integrity.
