package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/leonasting/Banking-Application/helpers"
	// 	"github.com/leonasting/Banking-Application/transactions"
	// 	"github.com/leonasting/Banking-Application/useraccounts"
	"github.com/leonasting/Banking-Application/users"

	"github.com/gorilla/mux"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

type Register struct {
	Username string
	Email    string
	Password string
}

func login(w http.ResponseWriter, r *http.Request) {
	// Read body form the request
	body, err := io.ReadAll(r.Body)
	helpers.HandleErr(err)
	// Handle Login
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)
	// Prepare response
	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
		// Handle error in else
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := io.ReadAll(r.Body)
	helpers.HandleErr(err)
	// Handle registration
	var formattedBody Register
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
	// Prepare response
	if register["message"] == "all is fine" {
		resp := register
		json.NewEncoder(w).Encode(resp)
		// Handle error in else
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))

}

// type TransactionBody struct {
// 	UserId uint
// 	From   uint
// 	To     uint
// 	Amount int
// }

// // Create readBody function
// func readBody(r *http.Request) []byte {
// 	body, err := ioutil.ReadAll(r.Body)
// 	helpers.HandleErr(err)

// 	return body
// }

// // Refactor apiResponse
// func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
// 	if call["message"] == "all is fine" {
// 		resp := call
// 		json.NewEncoder(w).Encode(resp)
// 		// Handle error in else
// 	} else {
// 		resp := call
// 		json.NewEncoder(w).Encode(resp)
// 	}
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	// Refactor login to use readBody
// 	body := readBody(r)

// 	var formattedBody Login
// 	err := json.Unmarshal(body, &formattedBody)
// 	helpers.HandleErr(err)

// 	login := users.Login(formattedBody.Username, formattedBody.Password)
// 	// Refactor login to use apiResponse function
// 	apiResponse(login, w)
// }

// func register(w http.ResponseWriter, r *http.Request) {
// 	body := readBody(r)

// 	var formattedBody Register
// 	err := json.Unmarshal(body, &formattedBody)
// 	helpers.HandleErr(err)

// 	register := users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password)
// 	// Refactor register to use apiResponse function
// 	apiResponse(register, w)
// }

// func getUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userId := vars["id"]
// 	auth := r.Header.Get("Authorization")

// 	user := users.GetUser(userId, auth)
// 	apiResponse(user, w)
// }

// func getMyTransactions(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userId := vars["userID"]
// 	auth := r.Header.Get("Authorization")

// 	transactions := transactions.GetMyTransactions(userId, auth)
// 	apiResponse(transactions, w)
// }

// // Create function transaction in api
// func transaction(w http.ResponseWriter, r *http.Request) {
// 	body := readBody(r)
// 	auth := r.Header.Get("Authorization")
// 	var formattedBody TransactionBody
// 	err := json.Unmarshal(body, &formattedBody)
// 	helpers.HandleErr(err)

// 	transaction := useraccounts.Transaction(formattedBody.UserId, formattedBody.From, formattedBody.To, formattedBody.Amount, auth)
// 	apiResponse(transaction, w)
// }

// func StartApi() {
// 	router := mux.NewRouter()
// 	// Add panic handler middleware
// 	router.Use(helpers.PanicHandler)
// 	router.HandleFunc("/login", login).Methods("POST")
// 	router.HandleFunc("/register", register).Methods("POST")
// 	router.HandleFunc("/transaction", transaction).Methods("POST")
// 	router.HandleFunc("/transactions/{userID}", getMyTransactions).Methods("GET")
// 	router.HandleFunc("/user/{id}", getUser).Methods("GET")
// 	fmt.Println("App is working on port :8888")
// 	log.Fatal(http.ListenAndServe(":8888", router))
// }
