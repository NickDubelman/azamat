# azamat

A lightweight sidekick for accessing SQL databases in Go.

<img alt="Azamat Bagatov" width="281" src=".github/assets/azamat.png" />

## Overview

- _Declare_ your db tables as structs
- _Associate_ each table with a type that represents an entry in the table (using generics)
- Use the table structs to _build_ and _execute_ queries

### What _isn't_ azamat?

Azamat is NOT an ORM. With azamat, you are still going to be writing SQL.

If you are looking for an ORM-like experience in Go, I highly recommend [Ent](https://entgo.io/).

### What _is_ azamat?

If you don't want to (or can't) use an ORM for some reason, azamat might be able to help you.

Azamat is a _thin_ collection of utilities glued together to help you access SQL databases in Go:

- [sqlx](http://jmoiron.github.io/sqlx/) extends Go's standard library [database/sql](https://pkg.go.dev/database/sql) with some useful utilities.
- [squirrel](https://github.com/Masterminds/squirrel) 🐿 is a SQL query builder. It allows us to programmatically build SQL queries. This is easier to read, less error prone, and more ergonomic than building query strings manually.
- [generics](https://go.dev/doc/tutorial/generics) are a new feature added to Go in version 1.18 (March 2022). Azamat leverages generics to save you from having to write boilerplate code.

## Getting Started

For a detailed introduction to azamat, check out the [guided tutorial](tutorial.md).

```go
// Imagine we have the following entity...
type Todo struct {
    ID        int
    Title     string
    Completed bool
}

// With azamat, we define our tables as /generic/ structs

// These structs accept a type parameter that tell azamat how to unmarshal rows in
// the given table

var TodoTable = azamat.Table[Todo]{
    Name:    "todos",
    Columns: []string{"id", "title", "completed"},
}

func main() {
    db, err := azamat.Connect(...)

    // We can GetByID and GetAll
    todo, err := TodoTable.GetByID(db, 420)
    todos, err := TodoTable.GetAll(db)

    // We can build queries:
    query := TodoTable.Select().Where("completed = ?", false)

    // To execute queries, we have All() and Only()

    todos, err = query.All(db) // 'All' gets all rows returned by the query

    todo, err = query.Only(db) // 'Only' expects a single row to be returned

     // We can build insert/updates/deletes:
    todoTitle := "buy food for bear"
    insert := TodoTable.
        Insert().
        Columns("title", "completed").
        Values(todoTitle, false)

    // To execute an insert/update/delete, we have Run()
    result, err := insert.Run(db)
    todoID, err := result.LastInsertId()

    update := TodoTable.Update().
        Set("title", todoTitle+"🐻").
        Set("completed", true).
        Where("id = ?", todoID)

    _, err = update.Run(db)

    delete := TodoTable.Delete().Where("id = ?", todoID)
    _, err = delete.Run(db)

    // All of the above code can be run on a DB or as part of a Tx
}
```

## Documentation

See [docs.md](docs.md)

## Similar Projects

- [genorm](https://github.com/mazrean/genorm)
- [dbr](https://github.com/gocraft/dbr)
- [go-structured-query](https://github.com/bokwoon95/go-structured-query)
