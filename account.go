package main

import (
	"errors"
)

func accountTransfer(transferInfo MakeTransfer) error {
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
