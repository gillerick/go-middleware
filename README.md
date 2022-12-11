## Working with Middleware and RPC

- This is a learning project based on chapter 3 of Naren Yellavula's
  book [Building RESTful Web Services with Go](https://www.packtpub.com/product/building-restful-web-services-with-go/9781788294287)

### What is middleware?

- A middleware is an entity that hooks into a server's request/response processing.
- A middleware can have many components with each component having a specific function to perform.
- A middleware performs the following functions in order:
    - Processes an incoming request before hitting a specific handler
    - Processes the handler function
    - Processes the response before giving it to a client
- In applications with middleware, not all requests get to the API server. Wrong requests can be filtered and prevented
  from interacting with the **business logic**. Below are some of the most common cases:
    - Using a logger to log each request before hitting the REST API
    - Validating a user's session and keeping the communication alive
    - Authenticating a user
    - Writing custom logic to scrap the request data
    - Attaching properties to responses before serving the client
- A middleware, therefore, helps in keeping housekeeping work, like logging and authentication, in its right place.