# azamat

A lightweight sidekick for accessing SQL databases in Go.

<img src="https://upload.wikimedia.org/wikipedia/commons/e/e5/Ken_Davitian_2010.jpg" alt="Azamat Bagatov" width="200"/>

## What isn't azamat?

Before we talk about what azamat is, let's talk about what it _isn't_.

Azamat is NOT an ORM. With azamat, you are still going to be writing SQL.

If you are looking for an ORM-like dev experience in Go, I highly recommend [Ent](https://entgo.io/).

## What is azamat?

If you don't want to or can't use an ORM for some reason, azamat might be able to help you.

Azamat is a thin collection of utilities glued together to help you structure your Go code to access SQL databases.

If it helps to see an equation:

> azamat = [sqlx](https://github.com/jmoiron/sqlx) + [squirrel](https://github.com/Masterminds/squirrel) 🐿 + [generics](https://go.dev/doc/tutorial/generics)
