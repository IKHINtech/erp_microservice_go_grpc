package middleware

import (
	"net/http"

	"github.com/IKHINtech/erp_api_gateway/internal/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			utils.RespondError(c, http.StatusInternalServerError, "INTERNAL_ERROR", err, nil)
		}
	}
}
