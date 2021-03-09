package main

import (
	"net/http"

	"github.com/userq11/image-transform/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/upload", handlers.Upload)

	http.ListenAndServe(":9090", mux)
}
