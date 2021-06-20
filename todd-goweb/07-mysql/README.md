# Using MySQL

- install workbench
- create mysql db on aws
- connect workbench to rds mysql db
- `https://www.youtube.com/watch?v=k68Y-XYapEI`

## Install MySQL

- [Download MySQL Community Server](http://dev.mysql.com/downloads/)

## We will need a MySQL driver

```bash
go get github.com/go-sql-driver/mysql
```

- [read the documentation](https://github.com/go-sql-driver/mysql#installation)
- [see all SQL drivers](https://github.com/golang/go/wiki/SQLDrivers)
- [Astaxie's book](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.2.html)

## Include the driver in your imports

```go
import (
    _ "github.com/go-sql-driver/mysql"
)
```

- [Read the documentation](https://github.com/go-sql-driver/mysql#usage)

## Determine the Data Source Name

- `user:password@tcp(localhost:5555)/dbname?charset=utf8`
- [Read the documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name)

## Open a connection

```go
db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")
```

[package sql](https://godoc.org/database/sql)