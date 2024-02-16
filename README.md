# Banking-Application

This repository is based on work presented in blog by Duomly Golang Course [link](https://github.com/Duomly/go-bank-backend)


# folder structure
```shell
Customer_Churn_Analysis/
├── api
│   └── api.go   
├── database
│   └── database.go
├── helpers
│   └── helpers.go
├── interfaces
│   └── interfaces.go
├── migrations
│   └── migrations.go
├── transactions
│   └── interfaces.go
├── useraccounts
│   └── useraccounts.go
├── users
│   └── users.go
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Run

- Setup a PostgreSQL db
- Create database with Name bankingapp and update connection details in connectDb(migration.go)
```
go run main.go
```

## Main File

main.go - This file acts as the main file which calls other functions.

The main application module  is initialized using the command 
```
go mod init github.com/leonasting/Banking-Application
```


### Migration File

It contains schema interfaces for User and Account.
* User has Username, email and password as fields. This interface contains users personal details.
* Account has type, name, balance and userid.This contains users acount details.

## Flow of logic

Main Function calls migragtion function.
Migration function first connects with DB and creates the two interfces.
Migration function Again conects to DB to insert Users which are hardcoded.