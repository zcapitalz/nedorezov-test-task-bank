package accountcontroller

import (
	"bank/internal/controllers/httputils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
)

func ParseAccountIDMiddleware(c *gin.Context) {
	accountID, err := ksuid.Parse(c.Param("accountID"))
	if err != nil {
		httputils.BindURIError(c, errors.Wrap(err, "parse accountID"))
		return
	}

	c.Set("accountID", accountID)
	c.Next()
}
