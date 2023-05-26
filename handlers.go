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

// listAccounts returns list of accounts http response
func listAccounts(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(accounts)
	if err != nil {
		err_response := ErrorResponse{Error: "couldn't encode Accounts to json"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_response)
		return
	}
}

// validateTransferRequest returns MakeTransfer struct and error
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

// makeTransfer make transfer between accounts
func makeTransfer(w http.ResponseWriter, r *http.Request) {

	requestBody, err := validateTransferRequest(r)
	if err != nil {
		err_response := ErrorResponse{Error: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err_response)
		return
	}

	err = accountTransfer(requestBody)
	if err != nil {
		err_response := ErrorResponse{Error: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err_response)
		return
	}

	message := "The Transfer is done"
	json.NewEncoder(w).Encode(message)
	return

}
