# 🪶 ddlite

[ [🐙 github](https://github.com/orsinium-labs/ddlite) ] [ [📚 docs](https://pkg.go.dev/github.com/orsinium-labs/ddlite/ddl) ]

A golang library for building SQLite DDL queries (CREATE TABLE, ALTER TABLE, etc). It focuses only on queries modifying tables and columns but not data and supports only SQLite, which makes it very safe, reliable, small, and easy to use.

The primary application is writing database migrations in Go for migration tools like [goose](https://github.com/pressly/goose).

Features:

* **Only SQLite**. No less, no more. No bloat, no options and types that do nothing, no runtime errors.
* **Only DDL**. How you manage tables and how you manage data is very different and should have a very different API. And this package has the best API you can get for DDL in Go.
* **Type safe**. Through careful use of types and generics, we make it impossible to make many of the common mistakes people make in their SQL queries.
* **Supports 99% of SQLite syntax**. The package is a careful translation of [SQL As Understood By SQLite](https://www.sqlite.org/lang.html), the full tech spec that fully covers SQL queries that SQLite can understand.
* **Human-readable output**. An internal tokenizer ensures that there are all needed spaces and parenthesis (or lack thereof) to make the generated SQL as pretty as possible.
* **Well-documented**. Every function and type has a short documentation, link to the SQLite spec, and a code example. Our goal is to make it possible to write queries without leaving IDE even if you don't know SQLite or sloppy with SQL.

## 📥 Installation

```bash
go get github.com/orsinium-labs/ddlite
```

## 🔧 Usage

```go
import "github.com/orsinium-labs/ddlite/ddl"

stmt := ddl.CreateTable("users",
    ddl.Column("name", ddl.Text, ddl.NotNull, ddl.PrimaryKey()),
    ddl.Column("age", ddl.Integer, ddl.Null),
)
sql := ddl.Must(ddl.SQL(stmt))
fmt.Println(sql)
//Output: CREATE TABLE users (name TEXT NOT NULL PRIMARY KEY, age INTEGER)
```

## 🤔 QnA

* **Why not raw SQL?** There is no THE language called SQL. There are only a bunch of dialects. Each database supports lots of dialects and aliases and usually very different subset of those. SQLite is especially quirky: some early day mistakes are now part of the SQLite spec because it's too late to change them now without breaking old apps. If you want to write correct SQL for SQLite you need to read the whole documentation or get some assistance. And we're here to save you time and provide that assistance.
* **Why only one RDBMS?** Each RDBMS (relational database management system) is very different. Focusing on one is the only way to have an API that isn't "the lowest common denominator".
* **Why SQLite?** That's a great database that many people underestimate. It's fast, powerful, reliable, and embedded.
* **Why only DDL?** How you manage tables and how you manage data is very different and should have a very different API. The idea to keep these separated is inspired by [Ecto](https://hexdocs.pm/ecto/Ecto.html), one of the best ORMs in the world, which has a separate module and separate DSL for migrations, schemas, queries, and validation.
* **Is it maintained?** SQLite doesn't get syntax updates that often. Their focus is on reliability, not features. Hence this project is pretty much feature-complete and doesn't require daily maintenance. Still, if you found a bug, contributions are welcome. Fix it, open a PR, and I'll review and merge it usually within a day.
