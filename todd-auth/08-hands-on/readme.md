# Hands-on Exercises

## Hands-on exercise #1

For this hands-on exercise:

- Create a server with two endpoints
  - Endpoint / should display a webpage with a register form on it
    - This form should take in at least a username and password
    - It should post to /register
  - Endpoint /register should save the username and password in a map
    - The password should be securely stored using bcrypt
    - Redirect the user back to endpoint / afterwards

## Hands-on exercise #2

For this hands-on exercise:

- Modify the server from the previous exercise
- Add a login form to the webpage
  - The form should take in a username and password
  - The form should post to a **/login** endpoint
- Add a new endpoint **/login**
  - The endpoint should compare the given credentials with stored credentials in the user map
    - Make sure to use the bcrypt function to compare the password
  - If the credentials match, display a webpage saying login successful
  - If the credentials do not match, display a webpage saying login failed

## Hands-on exercise #3

For this hands-on exercise:

- Modify the server from the previous exercise
- Create two functions
  - **createToken**
    - Should take in a session id
    - Should use HMAC to create a signature of the session id
    - Should combine the signature with the session id and return the signed string
      - Note, you will need to convert the hmac to a printable form
        - Hex
        - Base64
  - **parseToken**
    - Should take in a signed string from createToken
    - Should separate the signature from the session id
    - Should verify the signature matches the session id
    - Should return the session id

HINT:

- cryptography
  - large field
  - You don't need to understand it fully to use it
- Hashing
  - MD5 - don't use
  - **SHA**
  - Bcrypt
  - Scrypt
- Signing
  - Symmetric Key
    - **HMAC**
    - same key to sign (encrypt) / verify (decrypt)
  - Asymmetric Key
    - RSA
    - ECDSA - better than RSA; faster; smaller keys
    - private key to sign (encrypt) / public key to verify (decrypt)
  - **JWT**

NOT DISCUSSED IN VIDEO

- Encryption
  - Symmetric key
    - AES
  - Asymmetric Key
    - RSA

## Hands-on exercise #4

For this hands-on exercise:

- Modify the server from the previous exercise
- On **/login**
  - Generate a session id
    - Use a map **sessions** between session id and username
  - Use **createToken** with the generated session id
  - Set the token as the value of a session cookie
- Change login endpoint to redirect the user back to /
- On **/**
  - Use **parseToken** to get the session id
  - Use the sessions map to get the username
  - Display the username in the page if one has been found
    - No username will be displayed if
      - No session cookie set
      - Session cookie validation failed
      - No session in the sessions map to a username

## Hands-on exercise #5

For this hands-on exercise:

- Modify the server from the previous exercise
- have the db map store a struct of user fields, including bcrypted password
- display a user field (not bcrypted password) when someone’s session is active

## Hands-on exercise #6

For this hands-on exercise:

- Modify the server from the previous exercise
- Modify **createToken** and **parseToken** to use JWT
  - Create a custom claims type that embeds **jwt.StandardClaims**
  - The custom claims type should have the **session id** in it
  - Make sure to set the **ExpiresAt** field with a time limit
    - time.Now()
  - Use an HMAC signing method
  - Make sure to check if the token is valid in the **parseToken** endpoint

Question

Will we still need our sessions table / database?
YES!

## Hands-on exercise #7

For this hands-on exercise:

- Modify the server from the previous exercise
- Add a new form to **/**
  - It should just have a submit button labeled logout
  - It should post to **/logout**
- Add a new endpoint **/logout**
  - Delete the session of the current user from the **sessions** map
  - Set the cookie of the user to be deleted
