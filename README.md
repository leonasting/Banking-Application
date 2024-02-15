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
- You need to setup your connection string in the two files, vulnerable-db.go, and migrations.go
- Start migration by commenting migration in main.go and commenting API
- Type  (that will migrate to your db):
```
go run main.go
```
- Comment migration and uncomment api
- Type:
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