package handlers

import (
	"bytes"
	"encoding/json"
	"money-transfer-api/account"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestListAccounts tests that list endpoint is working correctly
func TestListAccounts(t *testing.T) {
	account.InitializeAccountsFromDisk()

	req := httptest.NewRequest("GET", "/list", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListAccounts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var accounts_returned map[string]*account.Account
	err := json.NewDecoder(rr.Body).Decode(&accounts_returned)
	if err != nil {
		t.Errorf("Error while decoding response body")
	}

	// check body isn't empty
	if accounts_returned == nil {
		t.Errorf("handler returned empty body")
	}
	// check body have some records, 10 is just a random number
	if len(accounts_returned) < 10 {
		t.Errorf("accounts returned body is small, it might be broken")
	}

}

// TestMakeTransfer tests transfer endpoint is working correctly
func TestMakeTransfer(t *testing.T) {
	body := account.TransferRequest{
		IdFrom: "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Amount: 1,
		IdTo:   "17f904c1-806f-4252-9103-74e7a5d3e340",
	}
	body_marshalled, err := json.Marshal(body)
	if err != nil {
		t.Errorf("impossible to marshall body: %s", err)
	}

	account_before_transfer, _ := account.GetAccount("17f904c1-806f-4252-9103-74e7a5d3e340")
	balance_before_transfer := account_before_transfer.Balance

	req := httptest.NewRequest("POST", "/trans", bytes.NewReader(body_marshalled))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MakeTransfer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Test Failed")
	}

	account_after_transfer, _ := account.GetAccount("17f904c1-806f-4252-9103-74e7a5d3e340")
	balance_after_transfer := account_after_transfer.Balance

	if !(balance_before_transfer+1 == balance_after_transfer) {
		t.Fatalf("Test failed here, makeTransfer doesn't work correctly!")
	}

}

// TestMakeTransferWithWrongRequestMethod tests that wrong request method(GET) return an error
func TestMakeTransferWithWrongRequestMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/transfer", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MakeTransfer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("Validate Request Method failed")
	}
}

// TestMakeTransferWithNotExistID tests that not exist id returns an error
func TestMakeTransferWithNotExistID(t *testing.T) {
	body_with_wrong_id := account.TransferRequest{
		IdFrom: "1",
		Amount: 0.0,
		IdTo:   "1",
	}
	body_with_wrong_id_marshalled, err := json.Marshal(body_with_wrong_id)
	if err != nil {
		t.Errorf("impossible to marshall body: %s", err)
	}

	req := httptest.NewRequest("POST", "/trans", bytes.NewReader(body_with_wrong_id_marshalled))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MakeTransfer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("Endpoint should return an error here")
	}
}

// TestMakeTransferWithNotAvailableBalance tests make transfer with balance bigger than available returns an error
func TestMakeTransferWithNotAvailableBalance(t *testing.T) {
	body_with_huge_balance := account.TransferRequest{
		IdFrom: "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Amount: 9999999999999999,
		IdTo:   "17f904c1-806f-4252-9103-74e7a5d3e340",
	}
	body_with_huge_balance_marshalled, err := json.Marshal(body_with_huge_balance)
	if err != nil {
		t.Errorf("impossible to marshall body: %s", err)
	}

	req := httptest.NewRequest("POST", "/trans", bytes.NewReader(body_with_huge_balance_marshalled))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MakeTransfer)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("Endpoint should return an error here")
	}

}
