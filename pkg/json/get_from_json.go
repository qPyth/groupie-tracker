package json

import (
	"encoding/json"
	"log"
	"net/http"
)

type DataJSON interface {
	FillFromJson([]byte) error
}

func FetchAndFill(url string, data DataJSON) DataJSON {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		log.Fatal(err)
	}
	return data
}
