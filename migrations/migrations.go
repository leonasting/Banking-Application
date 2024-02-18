package migrations

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/leonasting/Banking-Application/helpers"
	"github.com/leonasting/Banking-Application/interfaces"
	// "github.com/leonasting/Banking-Application/database"
)

func createAccounts() {
	db := helpers.ConnectDB()

	users := &[2]interfaces.User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}

	for i := 0; i < len(users); i++ {
		// Generate password is Username only
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	db := helpers.ConnectDB()
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccounts()
}

// // Refactor createAccounts to use database package
// func createAccounts() {
// 	users := &[2]interfaces.User{
// 		{Username: "Martin", Email: "martin@martin.com"},
// 		{Username: "Michael", Email: "michael@michael.com"},
// 	}
// 	for i := 0; i < len(users); i++ {
// 		// Correct one way
// 		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
// 		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
// 		database.DB.Create(&user)

// 		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
// 		database.DB.Create(&account)
// 	}
// }

// // Refactor Migrate
// func Migrate() {
// 	User := &interfaces.User{}
// 	Account := &interfaces.Account{}
// 	Transactions := &interfaces.Transaction{}
// 	database.DB.AutoMigrate(&User, &Account, &Transactions)

// 	createAccounts()
// }

// // Delete Migrate transactions
