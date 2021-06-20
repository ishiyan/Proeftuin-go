# Go & postgres

```bash
docker pull postgres
docker images

mkdir -p /postgresdata
mkdir c:\postgresdata

docker volume create postgresdata
docker volume inspect postgresdata
docker volume rm postgresdata

# Start the Docker container with the run command using the postgres image.
# The /var/lib/postgresql/data directory in the container is mounted as /postgresdata on the host.
# Additionally, this command changes the name of the container to postgres:
# -it – Provides an interactive shell to the Docker container.
# -v – Use this option to attach the /postgresdata host volume to the /var/lib/postgresql/data container volume.
# -d – Starts the container as a background process.
# --name – Name of the container.
# default user: postgres, set POSTGRES_USER to override he username
docker run -it -v postgresdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=password -p 5432:5432 --name postgresdb -d postgres

# Once the postgresdb server starts running in a container, check the status
docker ps

# Always check the Docker log to see the chain of events after making changes
docker logs postgresdb

# The container is currently running in detached mode.
# Connect to the container using the interactive terminal instead
docker exec -it postgresdb bash

# Start the pql shell by typing in the interactive terminal
psql -U postgres
# or
docker exec -it postgresdb psql -U postgres

# Type \q to leave the psql shell and then exit once again to leave the Interactive shell
\q
exit

# Stop
docker stop postgresdb

# Start again
docker start postgresdb
```

## 01 - Get driver

```bash
go get github.com/lib/pq
```

## 01 - Create a db, create  a user, and grant privileges

```sql
CREATE DATABASE bookstore;
CREATE USER bond WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE bookstore to bond;
```

## 02 - Our pathway

We will be following the example created in [this article by Alex Edwards](http://www.alexedwards.net/blog/practical-persistence-sql) and licensed under a [MIT license](https://opensource.org/licenses/MIT)

In order to successfully pull records from a table in a database as our user `bond`, we will need to [ALTER](https://www.postgresql.org/docs/9.6/static/sql-alteruser.html) bond to have a different role.

### 02 - Alter bond's role

```sql
alter user bond with superuser;
```

### 02 - Switch to your bookstore database

You should already have a `bookstore` database:

list databases

```bash
\l
```

switch into that database

```bash
\c bookstore
```

directory of tables, if any

```bash
\d
```

### 02 - Create table

```sql
CREATE TABLE books (
  isbn    char(14)     PRIMARY KEY NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);
```

directory of tables

```bash
\d
```

details of table `books`

```bash
\d books
```

### 02 - Insert records

```sql
INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccol� Machiavelli', 6.99);
```

view records

```sql
SELECT * FROM books;
```

## 02 - main.go

### 02 - Importing the driver

make sure you import the driver

```go
_ "github.com/lib/pq"
```

"We don't use anything in the pq package directly, which means that the Go compiler will raise an error if we try to import it normally. But **we need the pq package's init() function to run so that our driver can register itself with database/sql.** We get around this by aliasing the package name to the blank identifier. This means pq.init() still gets executed, but the alias is harmlessly discarded (and our code runs error-free). This approach is standard for most of Go's SQL drivers." - Alex Edwards

### 02 - Define a book type struct

```go
type Book struct {
    isbn   string
    title  string
    author string
    price  float32
}
```

"Next we define a Book type. **The struct fields and their types must align to our books table.** For completeness I should point out that we've only been able to use the string and float32 types safely because we set NOT NULL constraints on the columns in our table. If the table contained nullable fields we would need to use the sql.NullString and sql.NullFloat64 types instead � see [this Gist](https://gist.github.com/alexedwards/dc3145c8e2e6d2fd6cd9) for a working example. Generally it's easiest to avoid nullable fields altogether if you can, which is what we've done here." - Alex Edwards

### 02 - Initialize a new sql.DB

```go
db, err := sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()
```

"In the main() function we initialise a new sql.DB by calling sql.Open(). We pass in the name of our driver (in this case "postgres") and the connection string (you'll need to check your driver documentation for the correct format). It's worth emphasizing that **sql.DB is not a database connection � it's an abstraction representing a pool of underlying connections.** You can change the maximum number of open and idle connections in the pool with the db.SetMaxOpenConns() and db.SetMaxIdleConns() methods respectively. A final thing to note is that **sql.DB is safe for concurrent access,** which is very convenient if you're using it in a web application (like we will shortly)." - Alex Edwards

### 02 - Ping the db

```go
    if err = db.Ping(); err != nil {
        panic(err)
    }
```

"Because  sql.Open() doesn't actually check a connection, we also call DB.Ping() to make sure that everything works OK on startup." - Alex Edwards

### 02 - Query the db

```go
rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        panic(err)
    }
    defer rows.Close()
```

"We will fetch a resultset from the books table using the DB.Query() method and assign it to a  rows variable. Then we defer rows.Close() to ensure the resultset is properly closed before the parent function returns. **Closing a resultset properly is really important. As long as a resultset is open it will keep the underlying database connection open � which in turn means the connection is not available to the pool.** So if something goes wrong and the resultset isn't closed it can rapidly lead to all the connections in your pool being used up. Another gotcha (which caught me out when I first began) is that the defer statement should come after you check for an error from DB.Query. Otherwise, if DB.Query() returns an error, you'll get a panic trying to close a nil resultset." - Alex Edwards

### 02 - Iterate through results

```go
    for rows.Next() {
        bk := Book{}
        err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
        if err != nil {
            panic(err)
        }
        bks = append(bks, bk)
    }
```

"We then use rows.Next() to iterate through the rows in the resultset. This preps the first (and then each subsequent) row to be acted on by the rows.Scan() method. Note that if iteration over all of the rows completes then the resultset automatically closes itself and frees-up the connection. We use the rows.Scan() method to copy the values from each field in the row to a new Book object that we created. We then check for any errors that occurred during Scan, and add the new Book to a slice of books." - Alex Edwards

### 02 - Make sure everything ran well

```go
    if err = rows.Err(); err != nil {
        panic(err)
    }
```

"When our rows.Next() loop has finished we call rows.Err(). This returns any error that was encountered during the interation. It's important to call this � don't just assume that we completed a successful iteration over the whole resultset." - Alex Edwards Err returns the error, if any, that was encountered during iteration. Err may be called after an explicit or implicit Close.

## 03 - Webify

Add a package level scope variable

```go
var db *sql.DB
```

### 03 - Initialize your database

Note: `defer.db.Close()` has been removed

```go
func init() {
    var err error
    db, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
    if err != nil {
        panic(err)
    }

    if err = db.Ping(); err != nil {
        panic(err)
    }
    fmt.Println("You connected to your database.")
}
```

### 03 - Add routes & server

```go
func main() {
    http.HandleFunc("/", index)
    http.ListenAndServe(":8080", nil)
}
```

### 03 - Add a index func

```go
func index(w http.ResponseWriter, r *http.Request){
    if r.Method != "GET" {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }

    rows, err := db.Query("SELECT * FROM books")
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
    defer rows.Close()

    bks := make([]Book, 0)
    for rows.Next() {
        bk := Book{}
        err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
        if err != nil {
            http.Error(w, http.StatusText(500), 500)
            return
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }

    for _, bk := range bks {
        fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
    }
}
```

### 03 - Run the application and make a request

```bash
curl -i localhost:8080/books
```

## 04 - Query a single book

add these lines of code to `func main()`

```go
  http.HandleFunc("/books/show", booksShow)
```

### 04 - func bookShow added

Arguments to the SQL function are referenced in the function body using the syntax $n: $1 refers to the first argument, $2 to the second, and so on. If an argument is of a composite type, then the dot notation, e.g., $1.name, can be used to access attributes of the argument. The arguments can only be used as data values, not as identifiers.[source: postgres docs](https://www.postgresql.org/docs/9.1/static/xfunc-sql.html)

"Behind the scenes, **db.QueryRow (and also db.Query() and db.Exec()) work by creating a new prepared statement** on the database, and subsequently execute that prepared statement using the placeholder parameters provided. This means that **all three methods are safe from SQL injection** when used correctly . From Wikipedia:

Prepared statements are resilient against SQL injection, because parameter values, which are transmitted later using a different protocol, need not be correctly escaped. If the original statement template is not derived from external input, injection cannot occur.

The placeholder parameter syntax differs depending on your database. Postgres uses the $N notation, but MySQL, SQL Server and others use the ? character as a placeholder." - Alex Edwards

### 04 - Run the application and make a request

```bash
curl -i localhost:8080/books/show?isbn=978-1505255607
```

## 05 - Run the application and make a request

```bash
curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:8080/books/create/process
```

## 08 - 09 - 10 - Code organization

[Alex Edwards Code Organization Article](http://www.alexedwards.net/blog/organising-database-access)

How you organize your code depends upon the project and your abilities.

Do not sacrifice simplicity & readability for brevity & cleverness.

Your goal should be to write code which is maintainable. This means that an intermediate developer should be able to sit down, read your code, understand it, and work with it.

### 08 - 09 - 10 - Different Approaches

FYI: there is no global scope in Go
![Scope in Go (credit: Caleb Doxsey)](scope.png)

#### 08 - One package

When needed we can put variables in the package scope

This code is the same code from the previous folder `07-delete`

#### 09 - Two packages

We will import the code from the second package when needed

**Remember:** You need to capitalize to have a variable, type, or func exported from a package

Run the application and make a request

```bash
curl -i localhost:8080/books
curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:8080/books/create/process
```

#### 10 - Three+ packages with package for config variables

We will import the code from the config package when needed. The config package will hold variables like DB and TPL

**Remember:** You need to capitalize to have a variable, type, or func exported from a package

Run the application and make a request

```bash
curl -i localhost:8080/books
curl -i -X POST -d "isbn=978-1470184841&title=Metamorphosis&author=Franz Kafka&price=5.90" localhost:8080/books/create/process
```
