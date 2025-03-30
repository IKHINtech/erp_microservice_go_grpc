package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		md := metadata.New(map[string]string{"authorization": token})
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		c.Set("grpc_ctx", ctx) // Simpan context dalam gin.Context
		c.Next()
	}
}
