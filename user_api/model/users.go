package model

type User struct {
	ID       string `json:"ID"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"Role"`
}
