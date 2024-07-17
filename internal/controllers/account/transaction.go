package accountcontroller

import (
	"bank/internal/controllers/httputils"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Type TransactionType `json:"type"`
}

type TransactionType string

const (
	DepositTransactionType  TransactionType = "deposit"
	WithdrawTransactionType TransactionType = "withdraw"
)

func (c *AccountController) handleTransaction(ctx *gin.Context) {
	var transaction Transaction
	err := ctx.ShouldBindBodyWithJSON(&transaction)
	if err != nil {
		httputils.BindJSONError(ctx, err)
		return
	}

	switch transaction.Type {
	case DepositTransactionType:
		c.handleDepositTransaction(ctx)
	case WithdrawTransactionType:
		c.handleWithdrawTransaction(ctx)
	default:
		slog.Error("unknown transaction type", "transaction_type", transaction.Type)
		httputils.InternalError(ctx)
		return
	}
}
