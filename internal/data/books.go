package data

import "time"

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	Pages     int       `json:"pages"`
	Genres    []string  `json:"genres"`
	Version   int       `json:"version"`
}
