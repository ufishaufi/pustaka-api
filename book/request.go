package book

import "encoding/json"

type BookRequest struct {
	Title string `json:"title" binding:"required"`
	// Price interface{} `json:"price" binding:"required,number"`
	Price json.Number `json:"price" binding:"required, number"`
}
