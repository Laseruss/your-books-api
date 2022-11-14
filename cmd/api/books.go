package main

import (
	"fmt"
	"net/http"
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

	fmt.Fprintf(w, "show the details of the book with id: %d\n", id)
}
