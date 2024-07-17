package accountcontroller

import (
	"bank/internal/controllers/httputils"
	"bank/internal/domain"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

type depositTransactionRequestBody struct {
	Transaction
	Ammount float64 `json:"ammount" binding:"number,gt=0"`
}

// @Summary Deposit transaction
// @Description Deposit funds into an account
// @Tag DepositTransaction
// @Accept json
// @Param accountID path string true "Account ID"
// @Param request body depositTransactionRequestBody true "Deposit transaction request body"
// @Success 200 {object} nil "Success response"
// @Failure 400 {object} httputils.HTTPError "Bad request"
// @Failure 500 {object} httputils.HTTPError "Internal server error"
// @Router /accounts/{accountID}/transactions [post]
func (c *AccountController) handleDepositTransaction(ctx *gin.Context) {
	accountID := ctx.MustGet("accountID").(ksuid.KSUID)
	slog.Info("Depositing")
	var reqBody depositTransactionRequestBody
	err := ctx.ShouldBindBodyWithJSON(&reqBody)
	if err != nil {
		httputils.BadRequest(ctx, err)
		return
	}
	slog.Info("reqBody", "reqBody", reqBody)

	account, err := c.accountStorage.GetAccountByID(accountID)
	switch err.(type) {
	case nil:
	case domain.AccountNotFound:
		httputils.BadRequest(ctx, err)
		return
	default:
		httputils.InternalError(ctx)
		return
	}

	err = account.Deposit(reqBody.Ammount)
	if err != nil {
		httputils.InternalError(ctx)
		return
	}
}
