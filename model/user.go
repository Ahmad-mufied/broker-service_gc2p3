package model

type User struct {
	Id       string `json:"_id,omitempty"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
