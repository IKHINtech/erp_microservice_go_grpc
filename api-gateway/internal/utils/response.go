package utils

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Meta    any  `json:"meta,omitempty"` // Untuk pagination/dll
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Code    string `json:"code"` // Kode error spesifik (e.g., "AUTH_401")
	Details any    `json:"details,omitempty"`
}

func RespondSuccess(c *gin.Context, statusCode int, data any, meta any) {
	c.JSON(statusCode, SuccessResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func RespondError(c *gin.Context, statusCode int, errorCode string, err error, details any) {
	c.JSON(statusCode, ErrorResponse{
		Success: false,
		Error:   err.Error(),
		Code:    errorCode,
		Details: details,
	})
	c.Abort()
}
