package models

type UserCreateRequest struct {
	Name  string `json:"name"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

type User struct {
	Name string `json:"name"`
	Token string `json:"token"`
}

type Users []User