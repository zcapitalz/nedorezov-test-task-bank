package accountcontroller

import (
	"bank/internal/controllers/httputils"
	"bank/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

type getAccountBalanceResponseBody struct {
	Balance float64 `json:"balance"`
}

// @Summary Get account balance
// @Description Retrieve the balance of an account
// @Accept json
// @Produce json
// @Param accountID path string true "Account ID"
// @Success 200 {object} getAccountBalanceResponseBody
// @Failure 400 {object} httputils.HTTPError "Bad request"
// @Failure 500 {object} httputils.HTTPError "Internal server error"
// @Router /accounts/{accountID}/balance [get]
func (c *AccountController) getAccountBalance(ctx *gin.Context) {
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

	ctx.JSON(
		http.StatusOK,
		getAccountBalanceResponseBody{
			Balance: account.GetBalance(),
		})
}
