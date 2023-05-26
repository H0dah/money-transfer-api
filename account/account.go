package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var accounts map[string]*Account
var m sync.Mutex

type Account struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type TransferRequest struct {
	IdFrom string  `json:"id_from"`
	IdTo   string  `json:"id_to"`
	Amount float64 `json:"amount"`
}

type accountFromSource struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

// getData gets data from given url and returned it as an object
func getData(url string) ([]accountFromSource, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Endpoint Returned Not 200 status code")
	}

	res := []accountFromSource{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// constructAccountsMap convert data to account.Account type
func constructAccountsMap(data []accountFromSource) (map[string]*Account, error) {
	accounts := make(map[string]*Account)

	for _, v := range data {
		balance, err := strconv.ParseFloat(v.Balance, 64)
		if err != nil {
			return nil, fmt.Errorf("Error while parsing balance of %s", v.Name)
		}
		accounts[v.Id] = &Account{
			Name:    v.Name,
			Balance: balance,
		}
	}
	return accounts, nil
}

// InitializeAccounts initialize accounts by get them from source and constructing them
func InitializeAccounts(url string) error {

	data, err := getData(url)
	if err != nil {
		return errors.New("Error in getData function")
	}
	accounts, err = constructAccountsMap(data)
	if err != nil {
		return errors.New("Error in constructAccountsMap function")
	}
	return nil
}

func ListAccounts() map[string]*Account {
	return accounts
}

// AccountTransfer transfer balance from account to another
func AccountTransfer(transferInfo TransferRequest) error {
	m.Lock()
	defer m.Unlock()
	accountFrom := accounts[transferInfo.IdFrom]
	accountTo := accounts[transferInfo.IdTo]
	amount := transferInfo.Amount

	if amount > accountFrom.Balance {
		return errors.New("The account didn't have enough balance!")
	}

	accountFrom.Balance -= amount
	accountTo.Balance += amount
	return nil
}

// return account balance by account balance by ID
func GetAccount(accountId string) (*Account, bool) {
	account, exists := accounts[accountId]
	return account, exists
}
