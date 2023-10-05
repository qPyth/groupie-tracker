package main

import (
	"fmt"
	h "groupie-tracker/internal/http"
	"log"
	"net/http"
)

func main() {
	fmt.Println("asd")
	log.Fatal(http.ListenAndServe(":3000", h.Routes()))
}
