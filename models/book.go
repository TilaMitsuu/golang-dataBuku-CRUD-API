package models

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
}	