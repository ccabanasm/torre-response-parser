package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
)

var Db *sql.DB

func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}

// InitDb function initializes connection with database
func InitDb() {
	db, err := sql.Open("postgres", "postgres://torre:torre1234@localhost:5432/torredb?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	err = MakeMigration(db)
	if err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}
	Db = db
}
