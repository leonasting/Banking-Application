package main

import (
	"fmt"

	"github.com/leonasting/Banking-Application/migrations"
)

//"4.BANKING_APP/api"
//"github.com/leonasting/Banking-Application//database"

func main() {
	fmt.Println("Starting migrations")
	migrations.Migrate()
	// Do migration
	// migrations.MigrateTransactions()
	// migrations.MigrateUsers()
	// Init database
	//database.InitDatabase()
	//api.StartApi()
}
