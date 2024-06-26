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

func ExampleText() {
	col := ddl.Column("bio", ddl.Text(), ddl.Null)
	stmt := ddl.AddColumn("users", col)
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users ADD COLUMN bio TEXT
}

func ExampleEnum() {
	col := ddl.Column("pet", ddl.Enum("cat", "dog", "other", "none"), ddl.Null)
	stmt := ddl.AddColumn("users", col)
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users ADD COLUMN pet ENUM(cat, dog, other, none)
}

func ExampleFixedChar() {
	col := ddl.Column("iso_code", ddl.FixedChar(2), ddl.Null)
	stmt := ddl.AddColumn("countries", col)
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE countries ADD COLUMN iso_code CHAR(2)
}

func ExampleVarChar() {
	col := ddl.Column("status", ddl.VarChar(240), ddl.Null)
	stmt := ddl.AddColumn("users", col)
	sql := ddl.Must(ddl.SQL(dialects.PostgreSQL, stmt))
	fmt.Println(sql)
	//Output: ALTER TABLE users ADD COLUMN status VARCHAR(240)
}
