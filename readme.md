# gopractice

This is a collection of toy coding challenges to learn Go.

To get started, initialize a new Go module:

As a first assignment, what would be the ideal name for this module?
Second, how would you create a new module with this name?

Create a PR with your solution to the questions above. Once you are done, you can move on to the list of challenges below.

Happy coding!


## Toy coding challenges to learn Go


1. Write a function that takes a string and returns a slice of its words.

2. Write a function that takes a slice of strings and returns a new slice with all the strings that are longer than 5 characters.

3. Write a fibonacci function that takes an integer n and returns the nth fibonacci number.

4. Improve the fibonacci function using memoization.

5. Write a web server that listens on port 8080 and responds with "Hello, World!" for any request.

6. Create an api that exposes a router for the following endpoints:

   - `GET /users`: returns a list of all users
   - `GET /users/:id`: returns the user with the given id
   - `POST /users`: creates a new user
   - `PUT /users/:id`: updates the user with the given id
   - `DELETE /users/:id`: deletes the user with the given id

   > Note: You can keep all the data in memory, no need to persist it to a database.
   > You can use `net/http` to create a web server, or any other HTTP library of your choice (gin is a very good one).


### Hints

> Hint: You can use the `net/http` package to create a web server.
> You can use the `encoding/json` package to encode and decode JSON.
> You can use the `sync` package to implement memoization.
> You can use the `time` package to get the current time.
> You can use the `fmt` package to print to the console.

