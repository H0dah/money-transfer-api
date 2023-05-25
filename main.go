package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var accounts map[string]*Account
var m sync.Mutex

func main() {

	data, err := getData("https://git.io/Jm76h")
	if err != nil {
		log.Panic(err)
	}

	accounts, err = constructAccountsMap(data)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("The system is Ready to make a Transfer")

	// Api 1 to list accounts
	http.HandleFunc("/list", listAccounts)

	// Api 2 to make a transfer
	http.HandleFunc("/transfer", makeTransfer)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln(err)
	}
}
