package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"student-service/pkg/application/model"
	"student-service/pkg/config"
)

type CustomClaims struct {
	UserName string `json:"userName"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateToken(user model.Student) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		UserName: user.Username,
		Name:     user.Name,
	})
	cfg, _ := config.LoadConfig()
	jwtToken, _ := token.SignedString([]byte(cfg.SecretKey))
	return jwtToken
}

func ParseAndValidateToken(jwtToken string) (*jwt.Token, error) {
	cfg, _ := config.LoadConfig()
	return jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.SecretKey), nil
	})
}
