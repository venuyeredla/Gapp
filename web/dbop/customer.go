package dbop

import (
	"Gapp/web/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // _ refers to for side effects.
)

func Create(custoemr *models.EcomUser) {
	ctx := context.Background()
	sqlConn, con_err := db_con_pool.Conn(ctx)
	if con_err != nil {
		log.Fatal("Exceptin happned in getting connection")
	}
	defer sqlConn.Close()
	tx, tx_error := sqlConn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if tx_error != nil {
		log.Fatal("Error in initialzing transaction")
	}
	insert_query := "insert into ecom_user(first_name, last_name, email,pwd) values ('venu','gopal','venugopal@ecom.com','ecom#24'"
	_, exec_erro := sqlConn.ExecContext(ctx, insert_query)

	if exec_erro != nil {
		log.Fatal(exec_erro.Error())
	}
	tx.Commit()
}

func GetUserInfo(userid int) *models.EcomUser {
	ctx := context.Background()
	sqlConn, con_err := db_con_pool.Conn(ctx)
	if con_err != nil {
		log.Fatal("Exceptin happned in getting connection")
	}
	defer sqlConn.Close()
	tx, tx_error := sqlConn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if tx_error != nil {
		log.Fatal("Error in initialzing transaction")
	}
	row := sqlConn.QueryRowContext(ctx, "select * from ecom_user")
	var euser models.EcomUser
	e1 := row.Scan(&euser.Id, &euser.Firstname, &euser.Lastname, &euser.Email, &euser.Pwd)
	if e1 != nil {
		log.Fatal(e1.Error())
	}
	tx.Commit()
	return &euser
}

func GetUserInfos() []*models.EcomUser {
	//db, err := sql.Open("sqlite3", ":memory:")
	ctx := context.Background()
	sqlConn, con_err := db_con_pool.Conn(ctx)
	if con_err != nil {
		log.Fatal("Exceptin happned in getting connection")
	}
	defer sqlConn.Close()
	rows, err := sqlConn.QueryContext(ctx, "select * from ecom_user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	custmers := make([]*models.EcomUser, 0)

	for rows.Next() {
		var customer models.EcomUser
		err = rows.Scan(&customer.Firstname, &customer.Lastname, &customer.Firstname)
		if err != nil {
			log.Fatal(err)
		}
		custmers = append(custmers, &customer)
	}
	return custmers
}

func Update(custoemr *models.EcomUser) {
	//db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open(DATA_BASE_NAME, DB_CONNECTION_STRING)
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
