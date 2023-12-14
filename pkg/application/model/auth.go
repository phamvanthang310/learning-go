package model

type RegisterInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type LoginCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
