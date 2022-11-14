package main

import (
	"fmt"
	"net/http"
)

func (app *application) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world, this time I will for sure work")
}
