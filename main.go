package main

import (
	h "groupie-tracker/internal/http"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", h.Routes()))
}
