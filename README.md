# go-API
A simple REST API built to get more familiar with Go. It exposes an authenticated endpoint that returns a user's coin balance, backed by a mock database. 

## What it does
GET /account/coins returns the coin balance for a given user. Requests are authenticated with a username and an auth token before the handler runs. 
```bash
curl "http://localhost:8000/account/coins?username=alex" \
  -H "Authorization: 123ABC"
```
```json
{ "Code": 200, "Balance": 100 }
```
Mock users available: alex/123ABC, jason/456DEF, marie/789DEF. 
## Concepts covered 
* **Routing** **with** **chi**,

    Grouped routes under /account with route-scoped middleware. 
* **Middleware**
  
    A global StripSlashes middleware plus a custom *Authorization* middleware that validates the username and token before the request reaches the handler.
* **Interface-based** **Design**

    The database is defined as a *DatabaseInterface*, with a *mockDB* implementation. Swapping in a real database means writing one new implementation, no handler changes.
* **Structured** **error** **handling**

    Centralized *RequestErrorHandler* and *InternalErrorHandler* return consistent JSON error responses with status codes.

* **Request** **parsing**

    Query params decoded into structs with *gorilla/schema*.
* **Structured logging**

    Logrus with caller reporting enabled.

## Project structure 
```
api/              Shared types and error handlers
cmd/api/          Application entry point (main.go)
internal/         
  handlers/       Route definitions and request handlers 
  middleware/     Authorization middleware
  tools/          Database interface and mock implementation 
```
## Running the API 
```
go run ./cmd/api
```
The server starts on localhost:8000 

## Testing
Endpoints were tested and verified with Postman. 
