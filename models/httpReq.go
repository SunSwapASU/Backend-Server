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

type ItemFields struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Categories  string `json:"categories"`
	Condition   string `json:"condition"`
	Owner       string `json:"owner"`
}
