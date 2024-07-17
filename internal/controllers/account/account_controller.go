package accountcontroller

import (
	"bank/internal/controllers"
	"bank/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

type AccountController struct {
	accountStorage AccountStorage
}

type AccountStorage interface {
	CreateAccount() (ksuid.KSUID, error)
	GetAccountByID(accountID ksuid.KSUID) (domain.Account, error)
}

func NewAccountController(accountStorage AccountStorage) controllers.Controller {
	return &AccountController{
		accountStorage: accountStorage,
	}
}

func (c *AccountController) RegisterRoutes(engine *gin.Engine) {
	accountsGroup := engine.Group("api/v1/accounts")
	accountsGroup.POST("", c.createAccount)

	accountGroup := accountsGroup.Group("/:accountID", ParseAccountIDMiddleware)
	accountGroup.GET("/balance", c.getAccountBalance)

	transactionGroup := accountGroup.Group("/transactions")
	transactionGroup.POST("", c.handleTransaction)
}
