package utils

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/IKHINtech/erp_microservice_go_grpc/auth-microservice/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string, cfg config.Config) (string, error) {
	_, err := strconv.Atoi(cfg.JWTExpireHours)
	if err != nil {
		return "", errors.New("invalid JWT expiration time")
	}
	exp := time.Now().Unix()
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET tidak boleh kosong")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     exp,
	})
	return token.SignedString([]byte(cfg.JWTSecret))
}
