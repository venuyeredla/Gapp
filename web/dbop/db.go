package dbop

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_PATH  = "../../Gapp.db"
	C_CREATE = "CREATE TABLE Customer(id INTEGER PRIMARY KEY, name TEXT, price INT);"
)

func TestConnection() {
	//db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(version)
}

func ExecuteQuery(query string) (bool, error) {
	//db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, dbError := db.Exec(query)

	if dbError != nil {
		log.Fatal(err)
		return false, dbError
	}
	return true, nil
}
