package main

import (
	"OnlineBank/db"
	"OnlineBank/pkg/core/services"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Start(database *sql.DB) {
	for{
		Authorization(database)
	}
}

func Authorization(database *sql.DB)(login, password string) {
	fmt.Println(services.Welcome)
	var choice int64
	fmt.Scan(&choice)
	switch choice {
	case 1:
		login, password = services.LoginOperation()
		services.Login(database, login, password)
	case 2:
		services.PrintingATMs(database)
	case 3:
		fmt.Println("Good bye")
	default:
		fmt.Println("Некоректно введены данные, попробуйте ещё раз")
	}
	return login, password
}

func main() {
	database, err := sql.Open("sqlite3", "OnlineBank")
	if err != nil {
		log.Fatal("Can't open DB. Error is", err)
	}

	db.Init(database)
	Start(database)
}
