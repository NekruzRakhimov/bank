package db

import (
	"database/sql"
	"log"
)

func Init(database *sql.DB) {
	DDLs := []string {CreateUsersTable, CreateAccountsTable, CreateATMsTable}
	for _, ddl := range DDLs {
		_, err := database.Exec(ddl)
		if err != nil {
			log.Fatalf("Can't init %s error is %d", ddl, err)
		}
	}
}
