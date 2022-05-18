package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// initDB Create mysql connection.
func initDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	return db
}

// QueryName Query name by params id.
func QueryName(id string) (string, error) {
	db := initDB()
	var name string
	sqlStatement := `SELECT id, email FROM users WHERE id=$1;` // Query sql.

	err := db.QueryRow(sqlStatement, id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows { // find no rows
			return "", errors.Wrapf(err, "sql: %s,  params: %v", sqlStatement, id)
		} else {
			return "", err
		}
	}
	return name, nil
}

func main() {
	QueryName("00173439code")
}
