# ddlite

A golang library for building SQLite DDL queries (CREATE TABLE, ALTER TABLE, etc). It focuses only on queries modifying tables and columns but not data and supports only SQLite, which makes it very safe, reliable, small, and easy to use.

The primary application is writing database migrations in Go for migration tools like [goose](https://github.com/pressly/goose).

Features:

* **Only SQLite**.
* **Only DDL**.
* **Type safe**.
* **Supports 99% of SQLite syntax**.
* **Human-readable output**.
* **Well-documented**.

## QnA

* **Why not raw SQL?**
* **Why only one database?**
* **Why SQLite?**
* **Is it maintained?**

## Installation

```bash
go get github.com/orsinium-labs/ddlite
```

## Usage

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
