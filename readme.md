# Go Fundamentals

This project was created with the objective of collecting and showing fundamental concepts to learn Go easily

> Currently this repo is divided into 2 modules: Basics & Simple-API

## Basics

A unique package `main` that contains different files with fundamental Go concepts, ranging from creating a function to concurrency and error handling

### Files:

- [01- Functions and validation](Basics/main/01-funcValidation.go)
- [02- Arrays, Slices & Maps](Basics/main/02-arraySliceMap.go)
- [03- For Loop](Basics/main/03-forLoop.go)
- [04- Testing loop timing](Basics/main/04-testLoopTime.go)
- [05- String & Runes](Basics/main/05-stringRunes.go)
- [06- Structs](Basics/main/06-structs.go)
- [07- Pointers](Basics/main/07-pointers.go)
- [08- Go Routines with wait groups](Basics/main/08-goRoutinesWaitGroups.go)
- [09- Go Routines with channels](Basics/main/09-goRoutinesChannels.go)
- [10- Generics](Basics/main/10-generics.go)
- [11- Files](Basics/main/11-files.go)
- [12- Error Handling](Basics/main/12-errorHandling.go)
- [Extra - Testing with TestTables](Basics/main/main_test.go)

### Usage - Go Commands:

First, you have to go to the Basics folder, with the command `cd Basics`. Then, you can run the entire mini-project with the command `go run ./main`. You can also edit any file and create a new `main.exe` file with the command `go run build ./main`.

You can also run some sample tests cases (that use test tables) with `go test ./main -v`.

## Simple-API

Demo of a simple api that simulates fetching info from virtual wallets and user creation, with:

- Error logs
- Auth middleware with token
- Using packages such as Gorilla Mux and Gorilla Schema

### Usage - Go Commands:

First, you have to go to the Simple-API folder, with the command `cd Simple-API`. Then, you can run the mini-API with the command `go run main/main.go`. Then, you have to open _Postman_, _cURL_ or any tool to test apis. By default, the server will be running on `localhost:8080`.

The endpoints to be tested are:
| Endpoint | Method | Description | Observation |
| :- | :-: | :- | :- |
| `/user` | GET | Provides a list of mocked users | - |
| `/balance`|GET |Shows user's wallet balance | Uses Auth Middleware: so, it needs a `username` parameter in the URL (string type) and also requires an `Authorization` header with a "security token" string |
| `/user` | POST | Creates a new user and their new security token | - |

> All accepted values for the balance request can be found in the mocked DB, in the file: `internals/tools/mockDB.go`

### Requests examples with cURL

Get all users:

```
curl -L localhost:8080/user
```

Get one user´s wallet balance:

```
curl -L localhost:8080/balance?username=alex -H "Authorization: 123ABC"
```

Create a new user:

```
curl -L 'localhost:8080/user' -H 'Content-Type: application/json' --data-raw '{
    "Username": "Nick",
    "Mail": "nick@mail.com"
}'
```
