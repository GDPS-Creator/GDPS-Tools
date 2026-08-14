package mysql

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func Check(host, user, pass, db string) string {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, db)
    conn, err := sql.Open("mysql", dsn)
    if err != nil {
        return "MySQL ERROR: " + err.Error()
    }
    defer conn.Close()

    if err := conn.Ping(); err != nil {
        return "MySQL ERROR: " + err.Error()
    }

    return "MySQL OK"
}
