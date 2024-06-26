# nano-auth
Nano service for JWT based authentication and authorization

## Steps

1. Create git repository on github.com or other git repository
2. Clone your repository `git clone github.com/username/repo` 
3. Initialize go module with `go mod init github.com/remisb/nano-auth`
4. Initialize vuejs project


## Bearer

The bearer token is a string that's usually generated by the server.
The bearer token is generated by server after login request from the client.
The client can then use received token to access the resources.
Bearer token can be a hexadecimal character string.
Bearer token can be a structured JWT token.
Bearer token authentication is more secure:
  - username and password are used only once to get the token.
  - does not require to send username and password on every request.
  - tokens can be easily revoked if they are leaked.
  - revoked token prevents malicious access to the resource.
  - the best practice is to generate tokens that cannot be regenerated and that are not based on a pattern
Bearer token disadvantages:
  - the token is a randomly generated string that cannot be decoded.
  - it is impossible to get any further information about the owner of the token.
  - the server needs an extra query to fetch the details about the token owner.
  - having identifying information such as a username in the token can put the token at risk of being guessed and generated by an attacker.



## Terms

- Authentication (user identification) verifies the identity of  system user
    - another system trying to access the system
    - a person trying to access the system
    - helps identify your users
- Types of Authentication:
    - Basic HTTP
    - Bearer token
    - JWT
    - OIDC
    - SAML
- Authorizaton is used to check you can check user authorization before providing access to system resources.

## Ideas

- https://kinde.com
- https://www.descope.com/sign-up-google-search
- https://www.reddit.com/r/golang/comments/11l0cnk/authentication_in_go_best_practices/
- [Authentication for Go Applications: The Secure Way](https://www.jetbrains.com/guide/go/tutorials/authentication-for-go-apps/auth/)
- 




https://jasonwatmore.com/vue-3-pinia-jwt-authentication-with-refresh-tokens-example-tutorial