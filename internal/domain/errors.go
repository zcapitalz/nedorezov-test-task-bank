package domain

type AccountNotFound struct {
}

func (err AccountNotFound) Error() string {
	return "account not found"
}

type InvalidAmmountError struct {
	Message string
}

func (err InvalidAmmountError) Error() string {
	return err.Message
}

type InsufficientFundsError struct {
	Message string
}

func (err InsufficientFundsError) Error() string {
	return err.Message
}
