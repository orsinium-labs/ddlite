package ddl_test

import (
	"fmt"

	"github.com/orsinium-labs/sequel-ddl/ddl"
	"github.com/orsinium-labs/sequel-ddl/dialects"
)

func ExampleAddColumn() {
	stmt := ddl.AddColumn("users", ddl.Column("name", ddl.VarChar(128), ddl.NotNull))
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users ADD COLUMN name VARCHAR(128) NOT NULL
}

func ExampleAlterColumn() {
	stmt := ddl.AlterColumn("users", "name").Set(ddl.Null)
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users ALTER COLUMN name DROP NOT NULL
}

func ExampleCreateTable() {
	stmt := ddl.CreateTable("users",
		ddl.Column("name", ddl.VarChar(128), ddl.NotNull, ddl.PrimaryKey()),
		ddl.Column("age", ddl.Int(16), ddl.Null),
	)
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: CREATE TABLE users (name VARCHAR(128) NOT NULL PRIMARY KEY, age SMALLINT)
}

func ExampleDropColumn() {
	stmt := ddl.DropColumn("users", "age")
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users DROP COLUMN age
}

func ExampleDropTable() {
	stmt := ddl.DropTable("users")
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: DROP TABLE users
}

func ExampleRenameColumn() {
	stmt := ddl.RenameColumn("users", "created", "created_at")
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users RENAME COLUMN created TO created_at
}

func ExampleRenameTable() {
	stmt := ddl.RenameTable("users", "user")
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users RENAME TO user
}

func ExampleTruncateTable() {
	stmt := ddl.TruncateTable("users")
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: TRUNCATE TABLE users
}
