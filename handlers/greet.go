package handlers

import (
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello world")
}
