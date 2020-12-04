package models

import (
)

type User struct {
	ID int64
	Name string
	Surname string
	Age int64
	Gender string
	Login string
	Password string
	Role string
	Remove bool
}

//func AddNewUser(db *sql.DB, user User) (err error) {
//	_, err = db.Exec("INSERT INTO users(name, surname, age, gender, login, password, role, remove) values(($1), ($2), ($3), ($4), ($5), ($6), ($7), ($8))", user.Name, user.Surname, user.Age, user.Gender, user.Login, user.Password, user.role, user.Remove)
//
//	if err != nil {
//		log.Fatal("Can't insert to DB", err)
//	}
//	return err
//}
