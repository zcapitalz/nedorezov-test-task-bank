package storages

import (
	"bank/internal/domain"
	"sync"

	"github.com/segmentio/ksuid"
)

type AccountStorageInMemory struct {
	lock     sync.RWMutex
	accounts map[ksuid.KSUID]*accountInMemory
}

func NewAccountStorageInMemory() *AccountStorageInMemory {
	return &AccountStorageInMemory{
		accounts: make(map[ksuid.KSUID]*accountInMemory),
	}
}

func (s *AccountStorageInMemory) CreateAccount() (ksuid.KSUID, error) {
	account := accountInMemory{
		id: ksuid.New(),
	}
	s.lock.Lock()
	s.accounts[account.id] = &account
	s.lock.Unlock()

	return account.id, nil
}

func (s *AccountStorageInMemory) GetAccountByID(accountID ksuid.KSUID) (domain.Account, error) {
	s.lock.RLock()
	account, accountExists := s.accounts[accountID]
	s.lock.RUnlock()
	if !accountExists {
		return nil, domain.AccountNotFound{}
	}

	return account, nil
}
