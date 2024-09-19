package model

type Book struct {
	Id          string `json:"_id,omitempty"`
	Title       string `json:"title" validate:"required"`
	Author      string `json:"author" validate:"required"`
	PublishData string `json:"publish_data" validate:"required"`
	Status      string `json:"status" validate:"required"`
}
