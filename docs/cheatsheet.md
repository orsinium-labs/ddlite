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
| `dml.Delete`       | `DELETE`          | delete records from a table       |
| `dml.Insert`       | `INSERT`          | insert new records into a table   |
| `dml.Select`       | `SELECT`          | fetch records  matching filters   |
| `dml.Update`       | `UPDATE`          | update records matching filters   |
| `dml.CreateTable`  | `CREATE TABLE`    | create a new table                |
| `dml.DropTable`    | `DROP TABLE`      | delete a table                    |

## Basic expressions

| function  | description                           |
| --------- | ------------------------------------- |
| `dml.F`    | function                              |
| `dml.F2`   | function with 2 arguments             |
| `dml.C`    | column                                |
| `dml.M`    | optional column (field is a monad)    |
| `dml.V`    | value                                 |

## Boolean expressions

| function  | sql   | description           |
| --------- | ----- | --------------------- |
| `dml.Eq`   | `=`   | equal to              |
| `dml.E`    | `=`   | equal to alias to compare field to value  |
| `dml.Neq`  | `<>`  | not equal to          |
| `dml.Gt`   | `>`   | greater than          |
| `dml.Gte`  | `>=`  | greater than or equal |
| `dml.Lt`   | `<`   | less than             |
| `dml.Lte`  | `<=`  | less than or equal    |
| `dml.And`  | `AND` | both are true         |
| `dml.Or`   | `OR`  | any is true           |

## Column types

| function          | sql           | description                   |
| ----------------- | ------------- | ----------------------------- |
| `dml.SmallInt`     | `SMALLINT`    | small-range integer           |
| `dml.Integer`      | `INTEGER`     | typical choice for integer    |
| `dml.BigInt`       | `BIGINT`      | large-range integer           |
| `dml.Real`         | `REAL`        | variable-precision, inexact   |
| `dml.SmallSerial`  | `SMALLSERIAL` | small autoincrementing integer |
| `dml.Serial`       | `SERIAL`      | autoincrementing integer      |
| `dml.BigSerial`    | `BIGSERIAL`   | large autoincrementing integer |
| `dml.Text`         | `TEXT`        | variable unlimited length     |
| `dml.Character`    | `CHARACTER`   | fixed-length, blank padded    |
| `dml.VarChar`      | `VARCHAR`     | variable-length with limit    |
| `dml.Boolean`      | `BOOLEAN`     | state of true or false        |
| `dml.Date`         | `DATE`        | date (no time of day)         |
| `dml.TimeStamp`    | `TIMESTAMP`   | both date and time            |
