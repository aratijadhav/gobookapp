package rest

import (
	"fmt"
	"net/http"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home page")
}
