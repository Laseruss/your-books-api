package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Laseruss/your-books-api/internal/data"
)

func (app *application) addBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Add a new book")
}

func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	book := data.Book{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Shuggie Bain",
		Author:    "Douglas Stuart",
		Year:      2020,
		Pages:     463,
		Genres:    []string{"Fiction", "LGBT", "Britain"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, book, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
