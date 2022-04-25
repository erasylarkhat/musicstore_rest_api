package models

type Track struct {
	ID unit `json:"id" gorm:"primary_key"`
	Artist string `json:"artist"`
	Title string `json:"title"`
}