package models

import (
	"database/sql"
	"fmt"
)

type ATM struct {
	ID int64
	Address string
	Status bool
}

func AddNewATM(db *sql.DB, address string) {
	_, err := db.Exec("INSERT INTO atms(address) VALUES (($1))", address)
	if err != nil {
		fmt.Println("Couldn't insert new ATM. Error is:", err)
	}
}