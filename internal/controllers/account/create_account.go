package accountcontroller

import (
	"bank/internal/controllers/httputils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountResponseBody struct {
	AccountID string `json:"accountID"`
}

// @Summary Create an account
// @Description Create a new account
// @Produce json
// @Success 200 {object} createAccountResponseBody
// @Failure 500 {object} httputils.HTTPError "Internal server error"
// @Router /accounts [post]
func (c *AccountController) createAccount(ctx *gin.Context) {
	accountID, err := c.accountStorage.CreateAccount()
	if err != nil {
		httputils.InternalError(ctx)
		return
	}

	respBody := createAccountResponseBody{
		AccountID: accountID.String(),
	}
	ctx.JSON(http.StatusOK, respBody)
}
