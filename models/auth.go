package models

type RegisterCreds struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginCreds struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
