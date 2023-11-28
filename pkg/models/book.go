package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key" `
	Title  string `json:"title"`
	Author string `json:"author"`
}

type AddBook struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
