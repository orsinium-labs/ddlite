package ddl_test

import (
	"fmt"

	"github.com/orsinium-labs/sequel-ddl/ddl"
)

func ExampleAddColumn() {
	stmt := ddl.AddColumn("users", ddl.Column("name", ddl.Text, ddl.NotNull))
	sql := ddl.Must(ddl.SQL(stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users ADD COLUMN name TEXT NOT NULL
}

func ExampleCreateTable() {
	stmt := ddl.CreateTable("users",
		ddl.Column("name", ddl.Text, ddl.NotNull, ddl.PrimaryKey()),
		ddl.Column("age", ddl.Integer, ddl.Null),
	)
	sql := ddl.Must(ddl.SQL(stmt))
	fmt.Println(sql)
	//Output: CREATE TABLE users (name TEXT NOT NULL PRIMARY KEY, age INTEGER)
}

func ExampleDropColumn() {
	stmt := ddl.DropColumn("users", "age")
	sql := ddl.Must(ddl.SQL(stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users DROP COLUMN age
}

func ExampleDropTable() {
	stmt := ddl.DropTable("users")
	sql := ddl.Must(ddl.SQL(stmt))
	fmt.Println(sql)
	//Output: DROP TABLE users
}

func ExampleRenameColumn() {
	stmt := ddl.RenameColumn("users", "created", "created_at")
	sql := ddl.Must(ddl.SQL(stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users RENAME COLUMN created TO created_at
}

func ExampleRenameTable() {
	stmt := ddl.RenameTable("users", "user")
	sql := ddl.Must(ddl.SQL(stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users RENAME TO user
}

func ExampleTruncateTable() {
	stmt := ddl.TruncateTable("users")
	sql := ddl.Must(ddl.SQL(stmt))
	fmt.Println(sql)
	//Output: DELETE FROM users
}

func ExampleText() {
	col := ddl.Column("bio", ddl.Text, ddl.Null)
	stmt := ddl.AddColumn("users", col)
	sql := ddl.Must(ddl.SQL(stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users ADD COLUMN bio TEXT
}
