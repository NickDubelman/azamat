# azamat

A lightweight sidekick for accessing SQL databases in Go.

<img src="https://upload.wikimedia.org/wikipedia/commons/e/e5/Ken_Davitian_2010.jpg" alt="Azamat Bagatov" width="200"/>

## Introduction

### What isn't azamat?

Before we talk about what azamat is, let's talk about what it _isn't_.

Azamat is NOT an ORM. With azamat, you are still going to be writing SQL.

If you are looking for an ORM-like dev experience in Go, I highly recommend [Ent](https://entgo.io/).

### What is azamat?

If you don't want to or can't use an ORM for some reason, azamat might be able to help you.

Azamat is a thin collection of utilities glued together to help you structure your Go code to access SQL databases.

If it helps to see an equation:

> azamat = [sqlx](https://github.com/jmoiron/sqlx) + [squirrel](https://github.com/Masterminds/squirrel) 🐿 + [generics](https://go.dev/doc/tutorial/generics)

If you aren't familiar with any of these tools, I encourage you to click the links to find out what they're about. As a brief summary:

- [sqlx](https://github.com/jmoiron/sqlx) extends Go's standard library [database/sql](https://pkg.go.dev/database/sql) with some useful utilities
- [squirrel](https://github.com/Masterminds/squirrel) 🐿 is a SQL query builder. It allows us to programmatically build SQL queries. This is easier to read, less error prone, and more ergonomic than building query strings manually.
- [generics](https://go.dev/doc/tutorial/generics) are a new feature added to Go in version 1.18 (March 2022). Azamat leverages generics to save you from having to write boilerplate code

## Guided Tutorial

Pick a directory to create a project. I am going to use my `Desktop/Dev/Playground` directory:

```sh
cd Desktop/Dev/Playground
mkdir azamat-todos
cd azamat-todos
```

Initialize a new Go module:

```sh
go mod init azamat-todos
touch main.go
```

Open `main.go` in your editor of choice. To start, we are just going to establish a connection to a database so we can start playing around with azamat:

```go
package main

import (
    "fmt"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to db
	db, err := azamat.Connect("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
    fmt.Println(db)
}
```

Run the above code and make sure it works:

```sh
go mod tidy # this will detect and add our go-sqlite3 dep to go.mod
go run main.go
```

To keep things simple for this tutorial, we are connecting to an in-memory sqlite database. If you want, you can use a MySQL, MariaDB, Postgres or other SQL database instead. You just have to import the corresponding Go driver. You may also have to adapt some of the SQL syntax, but it should be straightforward.
