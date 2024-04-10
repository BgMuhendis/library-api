package model

type Book struct {
	Id     int    `gorm:"primaryKey; not null" json:"id"`
	Name   string `gorm:"not null" json:"name"`
	Page   int    `gorm:"not null" json:"page"`
	Author string `gorm:"not null" json:"author"`
	Status bool   `gorm:"not null" json:"status"`
}
