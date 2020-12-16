package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Db *sql.DB

// InitDb function initializes connection with database
func InitDb() {
	db, err := sql.Open("postgres", "postgres://torre:torre1234@localhost:5432/torredb?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}
	Db = db
}
