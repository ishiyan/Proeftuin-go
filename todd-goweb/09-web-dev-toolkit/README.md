# Misc topics

## Hash message authentication code (HMAC)

In cryptography, a keyed-hash message authentication code (HMAC) is a specific type of message authentication code (MAC) involving a cryptographic hash function (hence the 'H') in combination with a secret cryptographic key. As with any MAC, it may be used to simultaneously verify both the data integrity and the authentication of a message. Any cryptographic hash function, such as MD5 or SHA-1, may be used in the calculation of an HMAC; the resulting MAC algorithm is termed HMAC-MD5 or HMAC-SHA1 accordingly. The cryptographic strength of the HMAC depends upon the cryptographic strength of the underlying hash function, the size of its hash output, and on the size and quality of the key.

## Base64 encoding

Base encoding of data is used in many situations to store or transfer data in environments that, perhaps for legacy reasons, are restricted to US-ASCII data. Base encoding can also be used in new applications that do not have legacy restrictions, simply because it makes it possible to manipulate objects with text editors. In the past, different applications have had different requirements and thus sometimes implemented base encodings in slightly different ways. Today, protocol specifications sometimes use base encodings in general, and "base64" in particular, without a precise description or reference. Base64 is a group of similar binary-to-text encoding schemes that represent binary data in an ASCII string format. `https://tools.ietf.org/html/rfc4648`

In terms of actual standards, there have been a few attempts to codify cookie behaviour but none thus far actually reflect the real world.

RFC 2109 was an attempt to codify and fix the original Netscape cookie spec. In this standard many more special characters are disallowed, as it uses RFC 2616 tokens (a - is still allowed there), and only the value may be specified in a quoted-string with other characters. No browser ever implemented the limitations, the special handling of quoted strings and escaping, or the new features in this spec.
RFC 2965 was another go at it, tidying up 2109 and adding more features under a `version 2 cookies` scheme. Nobody ever implemented any of that either. This spec has the same token-and-quoted-string limitations as the earlier version and it's just as much a load of nonsense.
RFC 6265 is an HTML5-era attempt to clear up the historical mess. It still doesn't match reality exactly but it's much better then the earlier attempts. It is at least a proper subset of what browsers support, not introducing any syntax that is supposed to work but doesn't (like the previous quoted-string).
In 6265 the cookie name is still specified as an RFC 2616 token, which means you can pick from the alphanums plus:

```text
!#$%&'*+-.^_`|~
```

In the cookie value it formally bans the (filtered by browsers) control characters and (inconsistently-implemented) non-ASCII characters. It retains cookie spec's prohibition on space, comma and semicolon, plus for compatibility with any poor idiots who actually implemented the earlier RFCs it also banned backslash and quotes, other than quotes wrapping the whole value (but in that case the quotes are still considered part of the value, not an encoding scheme). So that leaves you with the alphanums plus:

```text
!#$%&'()*+-./:<=>?@[]^_`{|}~
```

In the real world we are still using the original-and-worst Netscape cookie spec, so code that consumes cookies should be prepared to encounter pretty much anything, but for code that produces cookies it is advisable to stick with the subset in RFC 6265.

`http://stackoverflow.com/questions/1969232/allowed-characters-in-cookies`

## Web storage

Web storage offers two different storage areas�local storage and session storage�which differ in scope and lifetime. Data placed in local storage is per origin (the combination of protocol, hostname, and port number as defined in the same-origin policy) (the data is available to all scripts loaded from pages from the same origin that previously stored the data) and persists after the browser is closed. Session storage is per-origin-per-window-or-tab and is limited to the lifetime of the window. Session storage is intended to allow separate instances of the same web application to run in different windows without interfering with each other, a use case that's not well supported by cookies.

Use cookies for secure storage.

Cookies have been around longer and have been built for secure storage.

Web storage is the relative new comer. There are some articles that talk about it being compromised.

### Session storage

Available only during the current session

### Local storage

Available until explicitly deleted

## Context

Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes. Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled. The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent) and return a derived Context (the child) and a CancelFunc. Calling the CancelFunc cancels the child and its children, removes the parent's reference to the child, and stops any associated timers. Failing to call the CancelFunc leaks the child and its children until the parent is canceled or the timer fires. The go vet tool checks that CancelFuncs are used on all control-flow paths. Programs that use Contexts should follow these rules to keep interfaces consistent across packages and enable static analysis tools to check context propagation: Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx:

```go
func DoSomething(ctx context.Context, arg Arg) error {
    // ... use ctx ...
}
```

Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use. Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions. The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.

### Call chain

Context makes it possible to manage a chain of calls within the same call path by signaling context’s Done channel.

source: `https://rakyll.org/leakingctx/`

## TLS & HTTPS

Transport Layer Security (TLS) and its predecessor, Secure Sockets Layer (SSL), both frequently referred to as "SSL", are cryptographic protocols that provide communications security over a computer network. Several versions of the protocols find widespread use in applications such as web browsing, email, Internet faxing, instant messaging, and voice-over-IP (VoIP). Websites use TLS to secure all communications between their servers and web browsers. HTTPS (also called HTTP over TLS, HTTP over SSL, and HTTP Secure) is a protocol for secure communication over a computer network which is widely used on the Internet. HTTPS consists of communication over Hypertext Transfer Protocol (HTTP) within a connection encrypted by Transport Layer Security or its predecessor, Secure Sockets Layer. The main motivation for HTTPS is authentication of the visited website and protection of the privacy and integrity of the exchanged data.

## JSON - JavaScript Object Notation

In computing, JSON (JavaScript Object Notation) is an open-standard format that uses human-readable text to transmit data objects consisting of attribute–value pairs.

It is the most common data format used for asynchronous browser/server communication, largely replacing XML.

JSON is a language-independent data format.

It derives from JavaScript, but as of 2016 many programming languages include code to generate and parse JSON-format data.

The official Internet media type for JSON is application/json.

JSON filenames use the extension .json.

Douglas Crockford originally specified the JSON format in the early 2000s; two competing standards, RFC 7159 and ECMA-404, defined it in 2013. The ECMA standard describes only the allowed syntax, whereas the RFC covers some security and interoperability considerations. JSON grew out of a need for stateful, real-time server-to-browser communication without using browser plugins such as Flash or Java applets, which were the dominant methods in the early 2000s. Douglas Crockford was the first to specify and popularize the JSON format.

The JSON.org Web site was launched in 2002. In December 2005, Yahoo! began offering some of its Web services in JSON. Google started offering JSON feeds for its GData web protocol in December 2006.

### Go & JSON - Marshal & Encode

Now we will see how to use Go with JSON. The most important thing to understand is that you can marshal *OR* encode Go code to JSON. Regardless of whether or not you use “marshal” or “encode”, your Go data structures will be turned into JSON. So what’s the difference? Marshal is for turning Go data structures into JSON and then assigning the JSON to a variable. Encode is used to turn Go data structures into JSON and then send it over the wire. Both “marshal” and “encode” have their counterparts: “unmarshal” and “decode”.

You can learn about Go & JSON at `https://godoc.org/encoding/json` - Package json implements encoding and decoding of JSON as defined in RFC 4627. The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.

You can also read about Go & JSON at this Go official blogpost: `https://blog.golang.org/json-and-go`

## AJAX introduction

Ajax (also AJAX) is short for asynchronous JavaScript and XML.

Ajax is a set of web development techniques using many web technologies on the client-side to create asynchronous Web applications.

With Ajax, web applications can send data to and retrieve from a server asynchronously (in the background) without interfering with the display and behavior of the existing page.

By decoupling the data interchange layer from the presentation layer, Ajax allows for web pages, and by extension web applications, to change content dynamically without the need to reload the entire page.

In practice, modern implementations commonly substitute JSON for XML due to the advantages of being native to JavaScript.

JavaScript and the XMLHttpRequest object provide a method for exchanging data asynchronously between browser and server to avoid full page reloads.

### AJAX server side

Now we will use Go to program our server’s response to an AJAX request. Remember, all AJAX is doing is making a request. A request includes some HTTP method and some route. For instance, a request might be “GET /user/score”. To handle an AJAX request on the server, we create a func to handle for that specific request; we program what we want the server to send back.
