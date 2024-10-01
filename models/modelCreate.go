package models

type CategoryCreate struct {
	Name string `json:"name"`
}

type ItemCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Categories  string `json:"categories"`
	Condition   string `json:"condition"`
	Owner       string `json:"owner"`
}
