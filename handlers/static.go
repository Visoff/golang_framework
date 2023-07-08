package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Static(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs("./httpdocs/" + strings.Join(strings.Split(r.URL.Path[1:], "/")[1:], "/"))
	if err != nil {
		fmt.Fprintf(w, "404 Error")
		return
	}
	fmt.Println(path)
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(w, "500 Error")
		return
	}
	result := RenderHtml(string(file[:]))
	fmt.Println(result)
	fmt.Fprintf(w, result)
}
