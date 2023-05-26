package main

import (
	"fmt"
	"log"
	"money-transfer-api/account"
	"money-transfer-api/handlers"
	"money-transfer-api/helpers"
	"net/http"
)

func main() {
	url := "https://git.io/Jm76h"
	data, err := helpers.GetData(url)
	if err != nil {
		log.Panic(err)
	}

	account.Accounts, err = helpers.ConstructAccountsMap(data)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("The system is Ready to make a Transfer")

	// Api 1 to list accounts
	http.HandleFunc("/list", handlers.ListAccounts)

	// Api 2 to make a transfer
	http.HandleFunc("/transfer", handlers.MakeTransfer)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln(err)
	}
}
