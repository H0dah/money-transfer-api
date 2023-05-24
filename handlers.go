package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type MakeTransfer struct {
	IdFrom string  `json:"id_from"`
	IdTo   string  `json:"id_to"`
	Amount float64 `json:"amount"`
}

func listAccounts(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(accounts)
	if err != nil {
		err_response := ErrorResponse{Error: "couldn't encode Accounts to json"}
		json.NewEncoder(w).Encode(err_response)
		return
	}
}

func validateTransferRequest(r *http.Request) (MakeTransfer, error) {

	if r.Method != "POST" {
		return MakeTransfer{}, errors.New("Invalid Request Method")
	}
	var requestBody MakeTransfer
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return MakeTransfer{}, errors.New("Invalid Request Body")
	}

	_, ok := accounts[requestBody.IdFrom]
	if !ok {
		return MakeTransfer{}, errors.New("The id_from you entered doesn't exist!")
	}
	_, ok = accounts[requestBody.IdTo]
	if !ok {
		return MakeTransfer{}, errors.New("The id_to you entered doesn't exist!")

	}

	return requestBody, nil

}

func makeTransfer(w http.ResponseWriter, r *http.Request) {

	requestBody, err := validateTransferRequest(r)
	if err != nil {
		err_response := ErrorResponse{Error: err.Error()}
		json.NewEncoder(w).Encode(err_response)
		return
	}
	accountFrom := accounts[requestBody.IdFrom]
	accountTo := accounts[requestBody.IdTo]

	if requestBody.Amount > accountFrom.Balance {
		err_response := ErrorResponse{Error: "The account didn't have enough balance!"}
		json.NewEncoder(w).Encode(err_response)
		return
	}

	accountFrom.Balance -= requestBody.Amount
	accountTo.Balance += requestBody.Amount

	message := ErrorResponse{Error: "The Transfer is done"}
	json.NewEncoder(w).Encode(message)
	return

}
