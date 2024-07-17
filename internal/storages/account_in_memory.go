package storages

import (
	"bank/internal/domain"
	"sync"

	"github.com/segmentio/ksuid"
)

type accountInMemory struct {
	sync.RWMutex
	id      ksuid.KSUID
	balance float64
}

func (a *accountInMemory) Deposit(ammount float64) error {
	a.RLock()
	if ammount <= 0 {
		return domain.InvalidAmmountError{Message: "ammount < 0"}
	}
	domain.LogDepositTransaction(a.id, ammount)
	a.RUnlock()

	a.Lock()
	a.balance += ammount
	a.Unlock()

	return nil
}

func (a *accountInMemory) Withdraw(ammount float64) error {
	a.RLock()
	if ammount <= 0 {
		return domain.InvalidAmmountError{Message: "ammount < 0"}
	}
	if a.balance < ammount {
		return domain.InsufficientFundsError{}
	}
	domain.LogWithdrawTransaction(a.id, ammount)
	a.RUnlock()

	a.Lock()
	a.balance -= ammount
	a.Unlock()

	return nil
}

func (a *accountInMemory) GetBalance() float64 {
	a.RLock()
	defer a.RUnlock()
	domain.LogGetAccountBalance(a.id)
	return a.balance
}
