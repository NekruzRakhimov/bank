package services

import (
	"OnlineBank/models"
	"database/sql"
	"fmt"
)

const Welcome = `Добро пожаловать!
1.Авторизоваться
2.Список банкоматов
3.Выйти`

func LoginOperation ()(login, password string) {
	fmt.Println(`Введите логин и пароль`)
	fmt.Println("Login:")
	fmt.Scan(&login)
	fmt.Println("Password:")
	fmt.Scan(&password)
	return login, password
}

func PrintingATMs (database *sql.DB) {
	var i int64
	rows, err := database.Query("select * from atms")
	if err != nil {
		fmt.Println("Can't select the list of ATMs. Error is:", err)
	}
	fmt.Println("№", "Адрес данного банкомата")
	for rows.Next(){
		i++
		atm := models.ATM{}
		err = rows.Scan(&atm.ID, &atm.Address, &atm.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(i, " ", atm.Address)
	}
}

func Login (database *sql.DB, login, password string) {
	var User models.User
	_ = database.QueryRow(`Select * from users where login = ($1) and password = ($2)`, login, password).Scan(
		&User.ID,
		&User.Name,
		&User.Surname,
		&User.Age,
		&User.Gender,
		&User.Login,
		&User.Password,
		&User.Role,
		&User.Remove,
	)
	UserServices(database, User)
}