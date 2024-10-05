package models

type User struct{
	ID string `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role string `json:"role"`
}
