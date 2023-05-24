package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type AccountFromSource struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

type Account struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func getData(url string) ([]AccountFromSource, error) {
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

func constructAccountsMap(data []AccountFromSource) (map[string]*Account, error) {
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
