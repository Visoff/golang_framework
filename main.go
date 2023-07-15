package main

import (
	"net/http"

	"visoff.ru/main/db"
)

func main() {
	db.Connect()
	mux := Router()
	http.ListenAndServe(":8080", mux)
}
