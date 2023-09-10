package http

import "net/http"

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", MainPageHandler)
	return mux
}
