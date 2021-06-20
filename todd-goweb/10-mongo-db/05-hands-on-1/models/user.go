package models

// changed Id type to string

// User is a model
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name" `
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}
