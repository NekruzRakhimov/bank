package services

import (
	"OnlineBank/models"
	"database/sql"
	"fmt"
)

const Usersvc = `1.Мои счета
2.Перевод денег
3.Оплата ком.услуг
4.История моих транзакций
5.Список банкоматов
6.Выход из личного кабинета`

func UserServices(database *sql.DB, user models.User)  {
	fmt.Println("							>", user.Name, user.Name)
	fmt.Println("Выберите нужую вам услугу")
	fmt.Println(Usersvc)
	var choice int64
	fmt.Scan(&choice)
	switch choice {
	case 1:
		PrintingUsersAccounts(database, user.ID)
	}
}

func PrintingUsersAccounts(database *sql.DB, user_id int64) {
	var i int64
	rows, err := database.Query("select * from accounts where user_id = ($1)", user_id)
	if err != nil {
		fmt.Println("Can't select the list of ATMs. Error is:", err)
	}
	for rows.Next(){
		i++
		account := models.Account{}
		err = rows.Scan(
			&account.ID,
			&account.User_id,
			&account.Amount,
			&account.Number,
			&account.Currency,
			&account.Remove,
			)
		if err != nil {
			fmt.Println(i, err)
			continue
		}
		fmt.Println(i, " ", account.Number)
	}
}
