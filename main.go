package main

import (
	"net/http"
)

func main() {
	mux := Router()
	http.ListenAndServe(":8080", mux)
}
