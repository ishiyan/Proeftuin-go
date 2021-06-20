# Hands-on Exercises

## Hands-on exercise #1

For this hands-on exercise:

- Choose an OAuth2 provider
- Find the documentation for the authentication process
  - Find the **AuthURL** and **TokenURL** in the documentation
    - find these in the documentation from the oauth2 provider
    - hints
      - tokenURL is often a POST
      - AuthURL is often (maybe always) a GET
  - If there is a sub-package in the oauth2 package ...
    - ([https://godoc.org/golang.org/x/oauth2](https://godoc.org/golang.org/x/oauth2))  
    - … make sure the endpoints match what is in the documentation
- Find how to get a user ID from the OAuth2 provider
  - Make sure you know how to make the call yourself
  - See if you need any special scope permissions to get the information
- Sign up to use the OAuth2 provider
  - If your provider wants a redirect url, set it to **[http://localhost:8080/oauth/your-provider/receive](http://localhost:8080/oauth/your-provider/receive)**
  - Make sure you have your ClientID and ClientSecret

## Hands-on exercise #2

For this hands-on exercise:

- Modify the server from Ninja Level 2
  - Create an oauth2.Config for your provider
    - Fill it in with information from hand-on-exercise 1
      - ClientID
      - ClientSecret
      - Endpoints
        - AuthURL
        - TokenURL
      - RedirectURL
        - **[http://localhost:8080/oauth/your-provider/receive](http://localhost:8080/oauth/your-provider/receive)**
      - Scopes
        - IF NEEDED
- Create an endpoint **/oauth/your-provider/login**
  - This should be a POST endpoint
  - This should generate a uuid to use as the state value
  - A new map should be used to save the state and the expiration time for this login attempt
    - Key is a **string** state
    - Value is a **time.Time** expiration time
      - One hour in the future is a reasonable time
  - Redirect the user to the oauth2 **AuthCodeURL** value
- Modify your / index page to include an oauth login form
  - The form should post to **/oauth/your-provider/login**
  - The login form should only have a submit button
- LET’S TEST OUR CODE!
  - Verify that attempting to login sends you to the oauth provider’s login page and approving sends you back to **/oauth/your-provider/receive**
    - You do not have this endpoint yet, if you are using the http.ServeMux, your index page will serve this endpoint

## Hands-on exercise #3

For this hands-on exercise:

- Modify the server from the previous exercise
- Create an endpoint **/oauth/your-provider/receive**
  - Should get the query parameter values **state** and **code**
  - State should be used to see if this state has expired
    - **time.Now().After()**
  - Code should be sent through the **config.Exchange** function to get a token
  - Create a **TokenSource** from the token
  - Create an **http.Client** with the **TokenSource**
  - Make your call you found in exercise 01 to get a user ID
  - Print out the result from your call

## Hands-on exercise #4

For this hands-on exercise:

- Modify the server from the previous exercise
- Create a new map **oauthConnections**
  - Key should be user IDs from the oauth provider
  - Value should be user IDs in your own system
- In endpoint **/oauth/your-provider/receive**
  - Extract just the user ID from the result of your call in the previous exercise
    - This will usually require creating a struct and unmarshalling/decoding json data
  - Get the local user ID from the **oauthConnections** map
  - If there was no value in the map, set the user ID to a random user ID from your users map
  - Create a session for your user just like how **/login** does it
    - Good candidate for pulling out into a separate function
  - Redirect the user back to **/**
- **Good luck!**
- **NOTES:**
  - to establish oauth2
    - interact with oauth2 api at some site offering oauth2 connections
    - sign up
    - get **clientsecret** and **clientid** from them
    - provide them with a **REDIRECTURL**
  - mysuperawesomesite.com
    - user can choose to oauth2 with form button
  - user gets sent to oauth2 website
    - our code calls
      - **AUTHCODEURL** which sends over
        - state
        - scope
        - clientid
      - user is asked to do this at oauth2 website provider
        - sign in
        - grant permissions
  - oauth2 website sends back to my **REDIRECTURL** (in config) and adds query parameters to the URL
    - sends back
      - state
      - code
  - we now call **EXCHANGE**
    - takes in
      - **context**
      - **code**
      - **client secret**
    - makes a call to oauth2 site sending this stuff over
      - calls to **TOKENURL**
    - we get back a token
  - we don’t know if we have an **accesstoken** that has expired or a **refreshtoken** so we call **TOKENSOURCE** which is a method that will give us back a value of type **TOKENSOURCE** which has a method called **TOKEN()** that can be called to return a valid **TOKEN**
    - calls to **TOKENURL**
  - call **NEWCLIENT** to get a client
  - now, using the client, make a get request to oauth2 provider
    - this will be some endpoint from oauth2 provider not already established in code
  - now we have the USERID of the user on the oauth2 provider’s site

```json
{"email":"tuddleymc@gmail.com","name":"Todd McLeod","user_id":"amzn1.account.AFKG2XU7BJYC27TS7BBNSVXYISUQ"}
```

## Hands-on exercise #4 - continued

- Modify the server from the previous exercise
- Create a new map **oauthConnections**
  - Key should be user IDs from the oauth provider
  - Value should be user IDs in your own system
- In endpoint **/oauth/your-provider/receive**
  - Extract just the user ID from the result of your call in the previous exercise
    - This will usually require creating a struct and unmarshalling/decoding json data
  - maybe somebody has previously registered with this oauth2 provider with our site
    - so get the local user ID from the **oauthConnections** map
      - If there was no value in the map, set the user ID to a random user ID from your users map
  - Create a session for your user just like how **/login** does it
    - Good candidate for pulling out into a separate function
  - Redirect the user back to /

## Hands-on exercise #5

For this hands-on exercise:

- Modify the server from the previous exercise
- In endpoint **/oauth/your-provider/receive**
  - When there is no value in the **oauthConnections** map
    - Sign the oauth provider’s user ID
      - Your **createToken** function for creating a JWT token should work
    - Redirect the user to **/partial-register**
      - Include the signed user ID in a query parameter
      - Also include any extra information you may get from the oauth provider (name, email, etc.)
        - Make sure to query escape the values
- Create endpoint **/partial-register**
  - Send users an html page
  - Page should have a form
    - Form should let the users fill in any information needed for their account
      - Pre-fill the values of any fields with the values from the query parameters, so the user may edit them if they wish
      - Example extra information:
        - Name
        - Age
        - Agree to the Terms of Service
        - Email
      - Use an input type **hidden** to include the signed user ID
    - Form should post to **/oauth/your-provider/register**

## Hands-on exercise #6

For this hands-on exercise:

- Modify the server from the previous exercise
- Create endpoint **/oauth/your-provider/register**
  - Extract the oauth provider’s user ID from its token
    - Your **parseToken** function should work
    - Send a user back to **/** if there is a problem
  - Create an entry in your user map
    - Fill in all information using the submitted form
    - Leave the bcrypted password field blank
  - Create an entry in your **oauthConnections** map
    - Key will be your provider’s user ID
    - Value will be the new user’s ID
  - Create a session for your user just like how **/login** and **/oauth/your-provider/receive** does it
  - Redirect the user back to **/**
- Make sure your **/login** endpoint will not log in anyone if they have no password
