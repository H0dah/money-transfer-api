package account

import (
	"errors"
	"sync"
)

var Accounts map[string]*Account
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

// AccountTransfer transfer balance from account to another
func AccountTransfer(transferInfo TransferRequest) error {
	m.Lock()
	defer m.Unlock()
	accountFrom := Accounts[transferInfo.IdFrom]
	accountTo := Accounts[transferInfo.IdTo]
	amount := transferInfo.Amount

	if amount > accountFrom.Balance {
		return errors.New("The account didn't have enough balance!")
	}

	accountFrom.Balance -= amount
	accountTo.Balance += amount
	return nil
}
