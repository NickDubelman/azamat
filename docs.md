# Docs

## Query Builders

For building queries, azamat leans on [squirrel](https://github.com/Masterminds/squirrel) 🐿. There aren't really docs per se, but there are tests and examples that show how to build queries.

The azamat builders mostly just wrap the corresponding squirrel builders. The builder types are:

- `SelectBuilder`
- `InsertBuilder`
- `UpdateBuilder`
- `DeleteBuilder`

The azamat version of `SelectBuilder` includes `All()` and `Only()` methods for executing the query. `All` is used when you expect multiple results from the query, while `Only` is for when you expect just a single result.

The azamat version of `SelectBuilder` includes a generic type param `T` that associates the query with a Go type that represents a row in the result set. This allows `All()` and `Only()` to be typed. There's no guarantee that the type is "in-sync" with the actual database, so you should make sure to write tests for this kind of stuff.

The azamat versions of `InsertBuilder`, `UpdateBuilder`, and `DeleteBuilder` include a `Run()` method for executing those types of statements. This is just a convenience shorthand for squirrel's `builder.RunWith(db).Exec()`

Azamat allows you to build queries "from scratch", which is what you _would_ do if you were using squirrel directly. However, when using azamat, you should define your tables as structs and build the queries off of those structs.

## Table Struct

The `Table` struct allows us to define the db tables we are interacting with. There are _three_ required pieces of information to define a `Table`. The first two are:

```go
Name    string // name of the table
Columns []string // columns of the table
```

The third requirement is a generic type param, which tells azamat what an individual row in the table looks like. Here's a basic example which should make it clear:

```go
type Todo struct {
    ID int
    Title string
    Completed bool
}

var TodoTable = azamat.Table[Todo]{
    Name: "todos",
    Columns: []string{"id", "title", "completed"}
}
```

Once you've defined your table structs, you can use them to build queries via the `Select`, `Insert`, `Update`, and `Delete` methods.

## Runner Interface

You may have code that sometimes runs on its own, and other times runs as part of a transaction. To address this use case, azamat has a `Runner` interface. A `Runner` is basically a type union: `sqlx.DB | sqlx.Tx`.

`Runner` is used by azamat internally for functions that can be run on their own _or_ as part of a Tx. You can also use `azamat.Runner` in your own code, if it helps. You don't necessarily want to use `Runner` everywhere, though; sometimes you'll have functions that explicitly should _only_ run as part of a transaction.

## View Struct

Azamat includes a `View` struct that is similar to its `Table` struct. Despite its name, `azamat.View` does not refer specifically to [SQL "views"](https://www.w3schools.com/sql/sql_view.asp) but it is very similar in concept.

A SQL view is essentially a "virtual" table: it is a dynamic query where the result set can be queried like a table. For example, you might represent a commonly-used join as a view so that you can just query the view instead of having to always write the join.

Azamat's `View` is intended to fulfill the same purpose. A `View` is only for querying, so it doesn't have `Insert()`, `Update()`, or `Delete()`. Like `Table`, it can be used to build custom queries with `Select()` and also has `GetAll`, `GetByID`, and `GetByIDs`.

If a `View` is associated with a query that references multiple tables with an `id` column, you will have to specify an `IDFrom` which designates the table that should be used when running `GetByID` and `GetByIDs`.

## CommitTransaction

When you are interacting with SQL transactions, you must make sure to always call `Commit` or `Rollback`. If you ever somehow forget, you'll likely have tables locked until garbage collection.

`azamat.CommitTransaction` takes a callback that contains your transaction. The callback returns an error; if there is an error, it calls `Rollback`, otherwise it `Commit`s the changes. It will also recover from panics and call `Rollback`.

## Postgres

Postgres uses a different placeholder format than other SQL dialects.

As such, you have to tell azamat if are using Postgres. If all of your tables are Postgres, you can just do `azamat.Postgres = true` to set this globally.

If only some of your tables are Postgres, you can opt-in those specific tables by setting the `Postgres` field of the `Table` struct.

## Other Notes
