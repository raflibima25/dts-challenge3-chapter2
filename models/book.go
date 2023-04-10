package models

type Book struct {
	ID       int    `json:"id"`
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
}
