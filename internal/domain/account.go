package domain

type Account interface {
	Deposit(ammount float64) error
	Withdraw(ammount float64) error
	GetBalance() float64
}
