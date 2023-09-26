# Chain of responsibility

**Chain of Responsibility** is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

Here in this project we are assuming that we need to hit to cache first, if data is not present in cache fetch the data via an HTTP request.
