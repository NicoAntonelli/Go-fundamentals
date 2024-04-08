# Git Fundamentals

This project was created with the objective of collecting and showing fundamental concepts to learn Go easily

> Currently this repo is divided into 2 modules: Basics & Simple-API

## Basics

Contains different packages with fundamental Go concepts, ranging from creating a function to concurrency and error handling

### Packages:

- [01- Functions and validation](basics/main/01-funcValidation.go)
- [02- Arrays, Slices & Maps](basics/main/02-arraySliceMap.go)
- [03- For Loop](basics/main/03-forLoop.go)
- [04- Testing loop timing](basics/main/04-testLoopTime.go)
- [05- String & Runes](basics/main/05-stringRunes.go)
- [06- Structs](basics/main/06-structs.go)
- [07- Pointers](basics/main/07-pointers.go)
- [08- Go Routines with wait groups](basics/main/08-goRoutinesWaitGroups.go)
- [09- Go Routines with channels](basics/main/09-goRoutinesChannels.go)
- [10- Generics](basics/main/10-generics.go)
- [11- Files](basics/main/11-files.go)
- [12- Error Handling](basics/main/12-errorHandling.go)

### Usage - Go Commands:

First, you have to go to the Basics folder, with the command `cd Basics`. Then, you can run the entire mini-project with the command `go run ./main`. You can also edit any file and create a new `main.exe` file with the command `go run build ./main`.

## Simple-API

Demo of a simple api with a single get that simulates fetching info from virtual wallets, with:

- Error logs
- Auth middleware with token
- Using packages such as Gorilla Mux and Gorilla Schema

Usage - Go Commands:

First, you have to go to the Simple-API folder, with the command `cd Simple-API`. Then, you can run the mini-API with the command `go run main/main.go`. Then, you have to open _Postman_, _cURL_ or any tool to test apis.

By default, the server will be running on `localhost:8080`. The endpoint to be tested is `/balance`: it waits for a GET request, has a "username" parameter in the URL (string type) and also requires an Authorization header with a string that pretends to be the security token.

All accepted values for the request can be found in the mocked DB, in the file: `internals/tools/mockDB.go`

> Request example: localhost:8080/balance?username=alex
