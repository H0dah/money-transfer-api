package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"money-transfer-api/account"
	"net/http"
	"strconv"
)

type AccountFromSource struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

// GetData gets data from given url and returned it as an object
func GetData(url string) ([]AccountFromSource, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Endpoint Returned Not 200 status code")
	}

	res := []AccountFromSource{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ConstructAccountsMap convert data to account.Account type
func ConstructAccountsMap(data []AccountFromSource) (map[string]*account.Account, error) {
	accounts := make(map[string]*account.Account)

	for _, v := range data {
		balance, err := strconv.ParseFloat(v.Balance, 64)
		if err != nil {
			return nil, fmt.Errorf("Error while parsing balance of %s", v.Name)
		}
		accounts[v.Id] = &account.Account{
			Name:    v.Name,
			Balance: balance,
		}
	}
	return accounts, nil
}
