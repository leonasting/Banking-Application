package migrations

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/leonasting/Banking-Application/helpers"
	// "github.com/leonasting/Banking-Application/database"
	// "github.com/leonasting/Banking-Application/interfaces"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=postgres dbname=bankapp password=postgres sslmode=disable")
	helpers.HandleErr(err)
	if err != nil {
		fmt.Println("Debug", err)
	} else {
		fmt.Println("Debug", "Connected to database")
	}
	return db
}

func createAccounts() {
	db := connectDB()

	users := [2]User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}

	for i := 0; i < len(users); i++ {
		// Correct one way
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})
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
