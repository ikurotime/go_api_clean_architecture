# Go Clean Arquitecture Example 🎯

This is a simple example of a Go project using Clean Arquitecture.

## Project Structure 📁

```sh
.
├── src/
│   ├── domain/
│       └── user.go
│   ├── infrastructure/
│   │   ├── router.go
│   │   └── sqlHandler.go
│   ├── interfaces/
│   │   ├── api/
│   │   │   ├── context.go
│   │   │   └── user_controller.go
│   │   └── database/
│   │   │   ├── sql_handler.go
│   │   │   └── user_repository.go
│   ├── usecase/
│   │   ├── user_interactor.go
│   │   └── user_repository.go
│   └── server.go
├── go.mod
├── go.sum
└── README.md
```

## How to run 🚀

```sh
go run src/server.go
```
