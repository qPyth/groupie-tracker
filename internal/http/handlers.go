package http

import (
	"fmt"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ASD")
}
