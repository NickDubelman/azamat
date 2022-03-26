# Guided Tutorial

## Initializing our project

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

    "github.com/NickDubelman/azamat"
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
go mod tidy # this will detect our deps and add them to go.mod
go run main.go
```

To keep things simple for this tutorial, we are connecting to an in-memory sqlite database. If you want, you can use a MySQL, MariaDB, Postgres or other SQL database instead. You just have to import the corresponding Go driver. You may also have to adapt some of the SQL syntax, but it should be straightforward.

## Accessing our database

We are ready to actually do stuff with our sidekick! With azamat, our development workflow primarily revolves around defining our database tables as Go structs. We use these tables to build and execute queries.

Not only do these table definitions facilitate our actual database access, but they also act as useful _documentation_. Since they are defined with our Go code, they provide a place for developers to check the shape of any given database entity. Ofcourse, there's no guaranteee this documentation will be in-sync with the actual table definitions, but we would have that problem either way (ie: even if we don't document our tables as Go structs, our code has to be in-sync with the actual table definitions).

For this tutorial, we'll pretend we're building a todo list, a timeless classic. Let's define our first table. In azamat, the `Table` struct is generic. Using generics, we associate the table with the corresponding type that represents entries in the table.

Define the following above your `main` function:

```go
type Todo struct {
    ID       int
    Title    string
    Author   string
}

var TodoTable = azamat.Table[Todo]{
    Name: "todos",
    Columns: []string{"id", "title", "author"},
    RawSchema: `
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        author TEXT NOT NULL
    `,
}
```

The `TodoTable` struct declares the name of the table, as well as its columns. Observe that it also associates the table with the `Todo` struct, which represents rows in the table. We will use the `Todo` struct throughout our application to represent this entity.

The `RawSchema` field is _optional_. If we specify it, we can ask azamat to create the table for us (which we'll do in a second). As discussed above, colocating the schema definition with our code also gives us useful documentation.

Now that we have declared the table, let's create it and start doing things with it.

Change your `main` function to:

```go
func main() {
    // Connect to db
    db, err := azamat.Connect("sqlite3", ":memory:")
    if err != nil {
        panic(err)
    }

    // Create table
    if err := TodoTable.CreateIfNotExists(db); err != nil {
        panic(err)
    }

    // Create an entry in the table
    insert := TodoTable.
        Insert().
        Columns("title", "author").
        Values("look after Borat", "Azamat")

    todoID, err := insert.Run(db)
    if err != nil {
        panic(err)
    }

    // Get all entries in the table
    todos, err := TodoTable.GetAll(db)
    if err != nil {
        panic(err)
    }
    fmt.Println("all todos:", todos)

    // Get an entry from the table by its ID
    todo, err := TodoTable.GetByID(db, todoID)
    if err != nil {
        panic(err)
    }
    fmt.Println("todo by id:", todo)

    // Write a custom query
    me := "Azamat"
    query := TodoTable.Select().Where("author = ?", me)

    // We use All() when we want all of the rows returned by the query
    todos, err = query.All(db)
    if err != nil {
        panic(err)
    }
    fmt.Println("custom query (all):", todos)

    // We use Only() when we expect the query to only return a single row
    todo, err = query.Only(db)
    if err != nil {
        panic(err)
    }
    fmt.Println("custom query (only):", todo)
}
```
