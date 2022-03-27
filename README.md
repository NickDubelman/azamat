# azamat

A lightweight sidekick for accessing SQL databases in Go.

<img src="https://upload.wikimedia.org/wikipedia/commons/e/e5/Ken_Davitian_2010.jpg" alt="Azamat Bagatov" width="200"/>

## Introduction

- _Declare_ your db tables as structs
- _Associate_ each table with a type that represents an entry in the table (using generics)
- Use the table structs to _build_ and _execute_ queries

### What isn't azamat?

Before we talk about what azamat is, let's talk about what it _isn't_.

Azamat is NOT an ORM. With azamat, you are still going to be writing SQL.

If you are looking for an ORM-like experience in Go, I highly recommend [Ent](https://entgo.io/).

### What is azamat?

If you don't want to (or can't) use an ORM for some reason, azamat might be able to help you.

Azamat is a thin collection of utilities glued together to help you structure your Go code to access SQL databases.

If it helps to see an equation:

> azamat = [sqlx](http://jmoiron.github.io/sqlx/) + [squirrel](https://github.com/Masterminds/squirrel) 🐿 + [generics](https://go.dev/blog/intro-generics)

If you aren't familiar with these tools, I encourage you to click the links to find out what they're about.

As a brief summary:

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

	// Insert an entry to table
	todoTitle := "buy food for bear"
	insert := TodoTable.
		Insert().
		Columns("title", "completed").
		Values(todoTitle, false)

	todoID, err := insert.Run(db)

	// We don't have to implement the GetByID boilerplate
	todo, err := TodoTable.GetByID(db, todoID)

	// We don't have to implement the GetAll boilerplate
	todos, err := TodoTable.GetAll(db)

	// We can build and execute queries
	query := TodoTable.Select().Where("completed = ?", false)
	todos, err = query.All(db) // 'All' gets all rows returned by the query

	query = TodoTable.Select().Where("title = ?", todoTitle)
	todo, err = query.Only(db) // 'Only' expects a single row to be returned

	// Update entry in table
	update := TodoTable.Update().
		Set("title", todoTitle+"🐻").
		Set("completed", true).
		Where("id = ?", todoID)

	_, err = update.Run(db)

	// Delete entry from table
	delete := TodoTable.Delete().Where("id = ?", todoID)
	_, err = delete.Run(db)

    // All of the above code is able to be run on a DB or as part of a Tx
}
```
