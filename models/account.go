package models

import (
	"database/sql"
	"fmt"
)

type Account struct {
	ID int64
	User_id int64
	Amount int64
	Number string
	Currency string
	Remove bool
}

func AddNewAccount(db *sql.DB, account Account) {
	_, err := db.Exec("INSERT INTO accounts(user_id, amount, number, currency) VALUES (($1), ($2), ($3), ($4))", account.User_id, account.Amount, account.Number, account.Currency)
	if err != nil {
		fmt.Println("Can't insert new account. Error is:", err)
	}
}