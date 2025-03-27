package auth

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	ValidateTokenRequest struct {
		Token string `json:"token" binding:"required"`
	}

	ValidateTokenResponse struct {
		IsValid bool     `json:"is_valid,omitempty"`
		UserId  string   `json:"user_id,omitempty"`
		Roles   []string `json:"roles,omitempty"`
	}

	RegisterUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	RefreshTokenResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	ErrorResponse struct {
		Error   string `json:"error"`
		Details string `json:"details,omitempty"`
	}
)
