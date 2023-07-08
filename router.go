package main

import (
	"net/http"

	"visoff.ru/main/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/static/", handlers.Static)
	mux.Handle("/components/", http.StripPrefix("/components/", http.FileServer(http.Dir("./components/"))))
	return mux
}
