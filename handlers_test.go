package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// test list endpoint is working correctly
func TestListAccounts(t *testing.T) {
	data, _ := getData(url)
	accounts, _ = constructAccountsMap(data)
	req := httptest.NewRequest("GET", "/list", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(listAccounts)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var accounts_returned map[string]*Account
	err := json.NewDecoder(rr.Body).Decode(&accounts_returned)
	if err != nil {
		t.Errorf("Error while decoding response body")
	}

	// check body isn't empty
	if accounts_returned == nil {
		t.Errorf("handler returned empty body")
	}
	// check body is correct
	eq := reflect.DeepEqual(accounts_returned, accounts)
	if !eq {
		t.Errorf("handler returned wrong body")
	}
}

// test transfer endpoint is working correctly
func TestMakeTransfer(t *testing.T) {
	body := MakeTransfer{
		IdFrom: "3d253e29-8785-464f-8fa0-9e4b57699db9",
		Amount: 1,
		IdTo:   "17f904c1-806f-4252-9103-74e7a5d3e340",
	}
	body_marshalled, err := json.Marshal(body)
	if err != nil {
		t.Errorf("impossible to marshall body: %s", err)
	}

	balance_before_transfer := accounts["17f904c1-806f-4252-9103-74e7a5d3e340"].Balance

	req := httptest.NewRequest("POST", "/trans", bytes.NewReader(body_marshalled))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeTransfer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Test Failed")
	}

	balance_after_transfer := accounts["17f904c1-806f-4252-9103-74e7a5d3e340"].Balance

	if !(balance_before_transfer+1 == balance_after_transfer) {
		t.Fatalf("Test failed here, makeTransfer doesn't work correctly!")
	}

}

// test that wrong request method(GET) return an error
func TestMakeTransferWithWrongRequestMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/transfer", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(makeTransfer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("Validate Request Method failed")
	}
}

// test that not exist id return an error
func TestMakeTransferWithNotExistID(t *testing.T) {
	body_with_wrong_id := MakeTransfer{
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
	handler := http.HandlerFunc(makeTransfer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("Endpoint should return an error here")
	}
}

// test make transfer with balance bigger than available returns an error
func TestMakeTransferWithNotAvailableBalance(t *testing.T) {
	body_with_huge_balance := MakeTransfer{
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
	handler := http.HandlerFunc(makeTransfer)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("Endpoint should return an error here")
	}

}
