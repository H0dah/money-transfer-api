package main

import (
	"fmt"
	"log"
	"money-transfer-api/account"
	"money-transfer-api/handlers"
	"net/http"
)

func main() {
	url := "https://git.io/Jm76h"

	err := account.InitializeAccounts(url)
	if err != nil {
		log.Panic(err)

	}
	fmt.Println("The system is Ready to make a Transfer")

	http.HandleFunc("/list", handlers.ListAccounts)
	http.HandleFunc("/transfer", handlers.MakeTransfer)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln(err)
	}
}
