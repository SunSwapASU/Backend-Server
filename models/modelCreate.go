package models

type CategoryCreate struct {
	Name string `json:"name"`
}

type ItemCreate struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryName string `json:"categoryName"`
	Condition    string `json:"condition"`
	Owner        string `json:"owner"`
}
