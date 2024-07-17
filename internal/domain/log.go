package domain

import (
	"log/slog"
	"time"

	"github.com/segmentio/ksuid"
)

func LogDepositTransaction(accountID ksuid.KSUID, ammount float64) {
	slog.Info("Deposit transaction",
		"time", time.Now().Format(time.RFC3339),
		"accountID", accountID.String(),
		"ammount", ammount)
}

func LogWithdrawTransaction(accountID ksuid.KSUID, ammount float64) {
	slog.Info("Withdraw transaction",
		"time", time.Now().Format(time.RFC3339),
		"accountID", accountID.String(),
		"ammount", ammount)
}

func LogGetAccountBalance(accountID ksuid.KSUID) {
	slog.Info("Get account balance",
		"time", time.Now().Format(time.RFC3339),
		"accountID", accountID.String())
}
