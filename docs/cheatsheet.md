# Cheatsheet

## Code execution

| function          | description                       |
| ----------------- | --------------------------------- |
| `sequel.SQL`      |  |
| `sequel.Exec`     |  |
| `sequel.FetchOne` |  |

## Query builders

| function          | sql               | description                       |
| ----------------- | ----------------- | --------------------------------- |
| `qb.Delete`       | `DELETE`          | delete records from a table       |
| `qb.Insert`       | `INSERT`          | insert new records into a table   |
| `qb.Select`       | `SELECT`          | fetch records  matching filters   |
| `qb.Update`       | `UPDATE`          | update records matching filters   |
| `qb.CreateTable`  | `CREATE TABLE`    | create a new table                |
| `qb.DropTable`    | `DROP TABLE`      | delete a table                    |

## Basic expressions

| function  | description                           |
| --------- | ------------------------------------- |
| `qb.F`    | function                              |
| `qb.F2`   | function with 2 arguments             |
| `qb.C`    | column                                |
| `qb.M`    | optional column (field is a monad)    |
| `qb.V`    | value                                 |

## Boolean expressions

| function  | sql   | description           |
| --------- | ----- | --------------------- |
| `qb.Eq`   | `=`   | equal to              |
| `qb.E`    | `=`   | equal to alias to compare field to value  |
| `qb.Neq`  | `<>`  | not equal to          |
| `qb.Gt`   | `>`   | greater than          |
| `qb.Gte`  | `>=`  | greater than or equal |
| `qb.Lt`   | `<`   | less than             |
| `qb.Lte`  | `<=`  | less than or equal    |
| `qb.And`  | `AND` | both are true         |
| `qb.Or`   | `OR`  | any is true           |

## Column types

| function          | sql           | description                   |
| ----------------- | ------------- | ----------------------------- |
| `qb.SmallInt`     | `SMALLINT`    | small-range integer           |
| `qb.Integer`      | `INTEGER`     | typical choice for integer    |
| `qb.BigInt`       | `BIGINT`      | large-range integer           |
| `qb.Real`         | `REAL`        | variable-precision, inexact   |
| `qb.SmallSerial`  | `SMALLSERIAL` | small autoincrementing integer |
| `qb.Serial`       | `SERIAL`      | autoincrementing integer      |
| `qb.BigSerial`    | `BIGSERIAL`   | large autoincrementing integer |
| `qb.Text`         | `TEXT`        | variable unlimited length     |
| `qb.Character`    | `CHARACTER`   | fixed-length, blank padded    |
| `qb.VarChar`      | `VARCHAR`     | variable-length with limit    |
| `qb.Boolean`      | `BOOLEAN`     | state of true or false        |
| `qb.Date`         | `DATE`        | date (no time of day)         |
| `qb.TimeStamp`    | `TIMESTAMP`   | both date and time            |
