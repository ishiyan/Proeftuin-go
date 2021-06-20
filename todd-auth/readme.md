# Web Authentication, Encryption, JWT, HMAC, & OAuth With Go

Todd McLeod & Daniel Hoffmann

- Udemy [course](https://saxobank.udemy.com/course-dashboard-redirect/?course_id=3472712)

- Outline on [google doc](https://docs.google.com/document/d/1iUem-Yt4eihj-WmNQ-xER8gspANHc9YDrSWZyWB2H94/edit#)

## JSON encoding and decoding

### Git version tagging

```bash
git add -S
git commit -m "message"
git push

git tag 0.1.0
git push --tags
```

### Git stash

```bash
git stash
git stash drop
```

### Using Curl in JSON decoding examples

Use [curlbuild.com](https://curlbuilder.com/) to build a command line.

```bash
curl -XGET -H "Content-type: application/json" -d '[{"First":"Jenny"}]' localhost:8080/decode
```

```bash
curl -XGET -H "Content-type: application/json" -d '[{"First":"Jenny"},{"First":"James"}]' localhost:8080/decode
```

## Authentication basics

### Difference between authentication and authorization

- Authentication
  - Determines who you are
  - Verifies that no-one is impersonating you
  - Three ways to authenticate
    - who you are (biometrics)
    - what you have (eg, atm card; key; phone)
    - what you know (username; password, ….)
  - Two-factor authentication

- Authorization
  - What permissions you have on a system
  - Says what you are allowed to do
  - The name of the http header used for authentication

[Good article](https://www.cyberciti.biz/faq/authentication-vs-authorization)

#### Authentication

Authentication verifies **who you are**.

For example, you can login into your Unix server using the ssh client, or access your email server using the `POP3` and `SMTP` client. Usually, `PAM` (`Pluggable Authentication Modules`) are used as low-level authentication schemes into a high-level application programming interface (API), which allows programs that rely on authentication to be written independently of the underlying authentication scheme.

#### Authorization

Authorization verifies **what you are authorized to do**.

For example, you are allowed to login into your Unix server via ssh client, but you are not authorized to browser `/data2` or any other file system. Authorization occurs after successful authentication. Authorization can be controlled at file system level or using various application level configuration options such as `chroot(2)`.

Usually, the connection attempt must be both authenticated and authorized by the system. You can easily find out why connection attempts are either accepted or denied with the help of these two factors.

#### Example: Authentication And Authorization

A user called `jbond` is allowed to login to `www.cyberciti.biz` server securely using the OpenSSH ssh client/server module. In this example authentication is the mechanism whereby system running at `wwwcyberciti.biz` may securely identify user `jbond`. The authentication systems provide an answers to the questions.

- Who is the user `jbond`?
- Is the user `jbond` really who he represents himself to be?

The server running at `www.cyberciti.biz` depend on some unique bit of information known only to the vivek user. It may be as simple as a password, public key authentication, or as complicated as `Kerberos` based system. In all cases user `jbond` needs some sort of secret to login into `www.cyberciti.biz` server via the ssh client. In order to verify the identity of a user called `jbond`, the authenticating system running at `www.cyberciti.biz` will challenges the `jbond` to provide his unique information (his password, or fingerprint, etc.) — if the authenticating system can verify that the shared secret was presented correctly, the user `jbond` is considered authenticated.

`jbond` is Authenticated? What Next? Authorization.

The Unix server running at `www.cyberciti.biz` determines what level of access a particular authenticated user called `jbond` should have. For example, `jbond` can compile programs using GNU gcc compilers but not allowed to upload or download files. So

- Is user `jbond` authorized to access resource called ABC?
- Is user `jbond` authorized to perform operation XYZ?
- Is user `jbond` authorized to perform operation P on resource R?
- Is user `jbond` authorized to download or upload files?
- Is user `jbond` authorized to apply patches to the Unix systems?
- Is user `jbond` authorized to make backups?

In this example Unix server used the combination of authentication and authorization to secure the system. The system ensures that user claiming to be `jbond` is the really user `jbond` and thus prevent unauthorized users from gaining access to secured resources running on the Unix server at `www.cyberciti.biz`.

#### Dealing With Large Linux / UNIX Setups

Large Linux / UNIX installation equipped with central `LDAP` directory servers to authenticate users. A user must provide username and password against all services such as `Squid proxy`, `Wi-Fi`, `SMTP`, `POP3 email server` etc. `LDAP` directory allows you to obtain required information such as employee number, email address, department code, and much more. The directory provides additional data lookup and search capabilities. `OpenLDAP` and the `Fedora Directory Server` (`FDS`) is an `LDAP` (`Lightweight Directory Access Protocol`) servers for Linux and Unix like operating systems. `Kerberos` is a network authentication protocol. It is designed to provide strong authentication for client/server applications by using secret-key cryptography. A free implementation of this protocol is available from the [Massachusetts Institute of Technology](http://web.mit.edu/kerberos/).

[Red Hat Directory Server](https://www.cyberciti.biz/faq/authentication-vs-authorization/) is an LDAP-compliant server that centralizes user identity and application information. It provides an operating system-independent, network-based registry for storing application settings, user profiles, group data, policies, and access control information.

### Http Basic Authentication

- Basic authentication is part of the specification of http
  - send username / password with every request
  - uses authorization header & keyword “basic”
    - put “username:password” together
    - converts them to base64
      - puts generic binary data into printable form
      - base64 is reversible
        - never use with http; only https
    - use basic authentication to login

## Hashing passwords

- never store passwords
- instead, store one-way encryption “hash” values of the password
- for added security
  - hash on the client
  - hash THAT again on the server
- hashing algorithms
  - bcrypt - current choice [https://godoc.org/golang.org/x/crypto/bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)
  - scrypt - new choice [https://godoc.org/golang.org/x/crypto/scrypt](https://godoc.org/golang.org/x/crypto/scrypt)

```bash
# getting packages
go mod init todd-auth
go mod tidy
```

## Bearer Tokens & Hmac

- bearer tokens
  - added to http spec with `OAUTH2`
  - uses authorization header & keyword “bearer”
- to prevent faked bearer tokens, use cryptographic “signing”
  - cryptographic signing is a way to prove that the value was created/validated by a certain person
  - `HMAC` is a cryptographic signing function, [https://godoc.org/crypto/hmac](https://godoc.org/crypto/hmac)

[hmac example](https://github.com/GoesToEleven/SummerBootCamp/tree/a40ab4ac3f7a497d49e73f336ccae6dc29107a5b/05_golang/02/03/11_sessions/11_03_caleb_sessions_HMAC)

## JWT

- stands for [JSON Web Token](http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html)
- common Go library: [jwt-go](https://github.com/dgrijalva/jwt-go)

```text
{Standard fields}.{Your fields}.{Signature}
 --------------- . ----------- . ---------
     base64      .   base64    .  base64

base64 has no periods, so we use them as delimiters
```

`jwt-go` has 3 signing methods, see [wikipedia](https://en.wikipedia.org/wiki/Digital_signature)

- [ECDSA](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm), `Elliptic Curve Digital Signature Algorithm`
- [RSA](https://en.wikipedia.org/wiki/RSA_(cryptosystem)), `Rivest–Shamir–Adleman`
- [HMAC](https://en.wikipedia.org/wiki/HMAC), `Hash-based Message Authentication Code`

`ECDSA` and `RSA` are assymmetric, they have two keys: private to sign and public to verify signature.
`HMAC` is symmetric, it has a single key to sign and to verify signature.

## OAuth2 overview

### What is OAuth2

Oauth2 allows

- we also use Oauth to login
  - example: login with facebook
- a user authorizes one website to do something at another website
  - example: give one website access to your dropbox account to store files there, or retrieve files from there

Two most common ways to do Oauth

- authorization code
  - more secure, but requires a server
  - “three legged flow”
- implicit
  - less secure

Here are the four flows, FYI

- **Authorization Code**
  - for apps running on a web server, browser-based and mobile apps
- **Password**
  - for logging in with a username and password (only for first-party apps)
- **Client credentials**
  - for application access without a user present
- **Implicit**
  - was previously recommended for clients without a secret, but has been superseded by using the Authorization Code grant with PKCE.
- [https://aaronparecki.com/oauth-2-simplified/#authorization](https://aaronparecki.com/oauth-2-simplified/#authorization)

### Overview of the OAuth2 process

- user is as spacex.com (for example)
  - logs in with Oauth2 using google Oauth2
  - Redirects user to Google Oauth login page
    - user is asked to grant permissions
    - what to share from google account
  - google Redirects back to spacex.com with a code
  - Spacex.com exchanges code and secret for access token to google
  - Spacex.com uses token to get who the user is on google, including user id on google
