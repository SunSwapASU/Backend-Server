package models

type UserSearch struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ItemSearch struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	OwnerId  string `json:"ownerId"`
}

type CategorySearch struct {
	Name string `json:"name"`
}
