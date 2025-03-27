package auth

import (
	"net/http"

	"github.com/IKHINtech/erp_api_gateway/internal/utils"
	authv1 "github.com/IKHINtech/erp_microservice_go_grpc/gen-proto/proto/auth/v1"
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
		utils.RespondError(
			c,
			http.StatusBadRequest,
			"AUTH_400",
			err,
			gin.H{"field": map[string]string{
				"email":    "must be valid email",
				"password": "min 8 characters",
			}},
		)
		return
	}

	token, err := h.client.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		utils.RespondError(
			c,
			http.StatusUnauthorized,
			"AUTH_401",
			err,
			nil,
		)
		return
	}

	utils.RespondSuccess(c, http.StatusOK, gin.H{"token": token}, nil)
}

// @Summary Register User
// @Tags Auth
// @Description Register User
// @Accept json
// @Produce json
// @Param request body RegisterUserRequest true "User Registration"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(
			c,
			http.StatusBadRequest,
			"AUTH_400",
			err,
			gin.H{"field": map[string]string{
				"name":     "min 2 characters",
				"email":    "must be valid email",
				"password": "min 8 characters",
			}},
		)
		return
	}

	d, err := h.client.Register(c.Request.Context(),

		&authv1.RegisterUserRequest{
			Email:    req.Email,
			Password: req.Password,
			FullName: req.FullName,
		})
	if err != nil {
		utils.RespondError(
			c,
			http.StatusUnauthorized,
			"AUTH_401",
			err,
			nil,
		)
		return
	}

	utils.RespondSuccess(c, http.StatusOK, gin.H{"user": d}, nil)
}

// @Summary      Validate token
// @Description Validate token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body ValidateTokenRequest true "Token"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /auth/validate [post]
func (h *Handler) ValidateToken(c *gin.Context) {
	var req ValidateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(
			c,
			http.StatusBadRequest,
			"400",
			err,
			nil,
		)
		return
	}

	token, err := h.client.ValidateToken(c.Request.Context(), req.Token)
	if err != nil {
		utils.RespondError(
			c,
			http.StatusUnauthorized,
			"AUTH_401",
			err,
			nil,
		)

		return

	}

	utils.RespondSuccess(c, http.StatusOK, gin.H{"token": token}, nil)
}
