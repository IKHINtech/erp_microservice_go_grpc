package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	client *AuthClient
}

func NewHandler(client *AuthClient) *Handler {
	return &Handler{client: client}
}

// @Summary      User Login
// @Description  Authenticate user with email/password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Credentials"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      401  {object}  ErrorResponse
// @Router       /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request format",
			Details: err.Error(),
		})
		return
	}

	token, err := h.client.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "Authentication failed",
		})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

func (h *Handler) ValidateToken(c *gin.Context) {
	var req ValidateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request format",
			Details: err.Error(),
		})
		return
	}

	token, err := h.client.ValidateToken(c.Request.Context(), req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "Authentication failed",
		})
		return
	}

	c.JSON(http.StatusOK, token)
}
