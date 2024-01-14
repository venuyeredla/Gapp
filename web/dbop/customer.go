package dbop

import (
	"Gapp/web/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Create(custoemr *models.Customer) {
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

	fmt.Println(version)
}

func Update(custoemr *models.Customer) {
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

	fmt.Println(version)
}

func GetCustomers() []*models.Customer {
	//db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM cars")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	custmers := make([]*models.Customer, 0)

	for rows.Next() {
		var customer models.Customer
		err = rows.Scan(&customer.Firstname, &customer.Lastname, &customer.Firstname)
		if err != nil {
			log.Fatal(err)
		}
		custmers = append(custmers, &customer)
	}
	return custmers
}
