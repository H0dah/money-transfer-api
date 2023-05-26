package handlers

import (
	"encoding/json"
	"errors"
	"money-transfer-api/account"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// ListAccounts returns list of accounts http response
func ListAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := account.ListAccounts()
	err := json.NewEncoder(w).Encode(accounts)
	if err != nil {
		err_response := ErrorResponse{Error: "couldn't encode Accounts to json"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_response)
		return
	}
}

// validateTransferRequest validate that request body is correct and data given is exist
func validateTransferRequest(r *http.Request) (account.TransferRequest, error) {

	if r.Method != "POST" {
		return account.TransferRequest{}, errors.New("Invalid Request Method")
	}
	var requestBody account.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return account.TransferRequest{}, errors.New("Invalid Request Body")
	}

	_, exists := account.GetAccount(requestBody.IdFrom)
	if !exists {
		return account.TransferRequest{}, errors.New("The id_from you entered doesn't exist!")
	}
	_, exists = account.GetAccount(requestBody.IdTo)
	if !exists {
		return account.TransferRequest{}, errors.New("The id_to you entered doesn't exist!")

	}

	return requestBody, nil

}

// MakeTransfer make transfer between accounts
func MakeTransfer(w http.ResponseWriter, r *http.Request) {

	requestBody, err := validateTransferRequest(r)
	if err != nil {
		err_response := ErrorResponse{Error: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err_response)
		return
	}

	err = account.AccountTransfer(requestBody)
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
