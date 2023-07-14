module github.com/orsinium-labs/qb/tests

go 1.18

replace github.com/orsinium-labs/qb => ../

require (
	github.com/jmoiron/sqlx v1.3.5
	github.com/matryer/is v1.4.1
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/orsinium-labs/qb v0.0.0-00010101000000-000000000000
)

require (
	github.com/Masterminds/squirrel v1.5.4 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
)
