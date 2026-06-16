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
* *Routing* *with* *chi*,
.... Grouped routes under /account with route-scoped middleware. 
* *Middleware*
.... A global StripSlashes middleware plus a custom Authorization middleware that validates the username and token before the request reaches the handler.  
## To run the API: 
```
go run cmd/api/main.go  
```
