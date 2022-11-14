package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/hello", app.helloHandler)
	router.HandlerFunc(http.MethodPost, "/books", app.addBookHandler)
	router.HandlerFunc(http.MethodGet, "/books/:id", app.getBookHandler)

	return router
}
