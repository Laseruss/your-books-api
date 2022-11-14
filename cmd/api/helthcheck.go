package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "environment": %q}`
	js = fmt.Sprintf(js, app.config.env)

	indentedJs := bytes.Buffer{}

	json.Indent(&indentedJs, []byte(js), "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(indentedJs.Bytes())
}
