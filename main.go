package main

import "github.com/leonasting/Banking-Application/api"

//"4.BANKING_APP/api"
//"github.com/leonasting/Banking-Application//database"

func main() {
	//fmt.Println("Starting migrations")
	//migrations.Migrate()
	// Do migration
	// migrations.MigrateTransactions()
	// migrations.MigrateUsers()
	// Init database
	//database.InitDatabase()
	api.StartApi()
}
