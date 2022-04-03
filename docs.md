# Docs

## Query Builders

For building queries, azamat leans on [squirrel](https://github.com/Masterminds/squirrel). While 🐿 doesn't really have docs, it has tests and examples that show how to build queries. It also maps pretty directly to SQL and the source code is easy to understand.

The azamat builders wrap the corresponding squirrel builders. The builder types are:

- `SelectBuilder`
- `InsertBuilder`
- `UpdateBuilder`
- `DeleteBuilder`

```go
query := azamat.Select[User]("id", "name", "email").From("users")

insert := azamat.
    Insert("users").
    Columns("name", "email").
    Values("Borat", "borat@aol.com")

update := azamat.
    Update("users").
    Set("name", "Pamela").
    Where("id = ?", id)

delete := azamat.Delete("users").Where("id = ?", id)
```

The azamat version of `SelectBuilder` includes `All()` and `Only()` methods for executing the query. `All` is used when you expect multiple results from the query, while `Only` is for when you expect just a single result.

Under the hood, `All` and `Only` use [Select from sqlx](http://jmoiron.github.io/sqlx/#getAndSelect), which loads the entire result set into memory at the same time. If you expect a large result set, you should use `query.RunWith(db).Query()` and ensure you close the rows. For _most_ queries, you shouldn't have to worry about loading the entire result set into memory.

The azamat version of `SelectBuilder` includes a generic type param `T` that associates the query with a Go type that represents a row in the result set. This allows `All` and `Only` to be typed. There's no guarantee that the type is "in-sync" with the actual database, so you should make sure to write tests for this kind of stuff.

The azamat versions of `InsertBuilder`, `UpdateBuilder`, and `DeleteBuilder` include a `Run()` method for executing those types of statements. This is just a convenience shorthand for squirrel's `builder.RunWith(db).Exec()`

Azamat allows us to build queries "from scratch", which is what we _would_ do if we were using squirrel directly. However, the preferred approach is to _define our tables as structs and build the queries off of those structs_.

## Table Struct

The `Table` struct allows us to define the db tables we are interacting with. There are _three_ required pieces of information to define a `Table`. The first two are:

```go
Name    string   // name of the table
Columns []string // columns of the table
```

The third requirement is a generic type param, which tells azamat what an individual row in the table looks like. Here's a basic example which should make it clear:

```go
type Todo struct {
    ID        int
    Title     string
    Completed bool
}

var TodoTable = azamat.Table[Todo]{
    Name: "todos",
    Columns: []string{
        "id",
        "title",
        "completed",
    },
}
```

Once we've defined our table structs, we can use them to [build queries](#query-builders) via the `Select`, `Insert`, `Update`, and `Delete` methods.

### Table `RawSchema`

Optionally, we can specify the `RawSchema` for a `Table` definition. This serves two purposes:

1. It allows us to use the `Create` and `CreateIfNotExists` methods to create the table. This can be useful in our tests, or in general if we just want our code to create our db tables.
1. It acts as documentation. For example, you could hover over a `Table` struct in your editor and it would show you the schema. This documentation _declares_ what our code _thinks_ our db tables look like.

If you do plan to leverage the `RawSchema` as documentation or in tests, you just have to make sure you keep it in-sync with your actual db tables. In order for the documentation to be useful, it has to be accurate.

## Runner Interface

You may have code that sometimes runs on its own, and other times runs as part of a transaction. To address this use case, azamat has a `Runner` interface. A `Runner` is basically a type union: `sqlx.DB | sqlx.Tx`.

`Runner` is used by azamat internally for functions that can be run on their own _or_ as part of a Tx. You can also use `azamat.Runner` in your own code, if it helps. You don't necessarily want to use `Runner` everywhere, though; sometimes you'll have functions that explicitly should _only_ run as part of a transaction.

## View Struct

Azamat includes a `View` struct that is similar to its `Table` struct. If you are familiar with [SQL "views"](https://www.w3schools.com/sql/sql_view.asp), azamat's `View` is very similar.

A SQL view is a dynamic query where the result set can be queried like a table. For example, you might represent a commonly-used join as a view so that you can just query the view, instead of always having to write the join.

Azamat's `View` is intended to fulfill the same purpose. A `View` is only for querying, so it doesn't have `Insert()`, `Update()`, or `Delete()`. Like `Table`, it can be used to build custom queries with `Select()`, and it also has `GetAll`, `GetByID`, and `GetByIDs`.

If a `View` is associated with a query that references multiple tables that have an `id` column, you have to specify an `IDFrom` to designate the table that should be referenced when running `GetByID` and `GetByIDs`.

## CommitTransaction

When you are interacting with SQL transactions, you must make sure to always call `Commit` or `Rollback`. If you ever somehow forget, you'll likely have tables locked until garbage collection.

`azamat.CommitTransaction` takes a callback that contains your transaction. The callback returns an error; if there is an error, it calls `Rollback`, otherwise it calls `Commit`. It also recovers from panics and calls `Rollback`.

## Postgres

Postgres uses a different placeholder format than other SQL dialects.

As such, you have to tell squirrel (via azamat) if you are using Postgres. If all of your tables are Postgres, you can just do `azamat.Postgres = true` to set this globally.

If only some of your tables are Postgres, you can opt-in those specific tables by setting the `Postgres` field of the `Table` struct.

## Other Notes

### Why squirrel?

### Why sqlx?

### Why generics?
