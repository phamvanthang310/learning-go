package model

type RegisterInfo struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginCredential struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthClaims struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
