# Go Webservice Demo

A REST-ish Webservice written in Go as seen in the talk.

## Run

With go installed, just hit `go run .` from within the directory.

The Server will listen on http://localhost:1111

### Routes

|  Route | Verb | Description | Example |
| --- | --- | --- | --- |
| `/todos` | GET | List all entries |  `curl -i http://localhost:1111/todos/` |
| `/todo/:id` | GET | Get entry with id `:id` | `curl -i http://localhost:1111/todo/1` |
| `/todo/:id` | DELETE | Get entry with id `:id` | `curl -i -X DELETE http://localhost:1111/todo/1` |
| `/todo/` | POST, PUSH | Get entry with id `:id` | `curl -i -X POST -H "Content-Type: application/json" -d "{\"id\":2, \"message\": \"something\"}" http://localhost:1111/todo/` |
