package accountcontroller

import (
	"bank/internal/controllers/httputils"
	"bank/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

type withdrawTransactionRequestBody struct {
	Transaction
	Ammount float64 `json:"ammount" binding:"number,gt=0"`
}

// @Summary Withdraw transaction
// @Description Perform a withdrawal transaction from the account
// @Tag WithdrawTransaction
// @Accept json
// @Param accountID path string true "Account ID"
// @Param body body withdrawTransactionRequestBody true "Withdraw transaction request body"
// @Success 200 {object} nil "Success response"
// @Failure 400 {object} httputils.HTTPError "Bad request"
// @Failure 500 {object} httputils.HTTPError "Internal server error"
// @Router /accounts/{accountID}/transactions [post]
func (c *AccountController) handleWithdrawTransaction(ctx *gin.Context) {
	accountID := ctx.MustGet("accountID").(ksuid.KSUID)

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

	var reqBody withdrawTransactionRequestBody
	err = ctx.ShouldBindBodyWithJSON(&reqBody)
	if err != nil {
		httputils.BindJSONError(ctx, err)
		return
	}

	err = account.Withdraw(reqBody.Ammount)
	if err != nil {
		httputils.InternalError(ctx)
		return
	}
}
