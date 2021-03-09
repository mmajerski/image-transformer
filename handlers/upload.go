package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/userq11/image-transform/primitive"
)

// Upload handles '/upload' endpoint
func Upload(rw http.ResponseWriter, r *http.Request) {
	shapes, err := strconv.Atoi(r.FormValue("shapes"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	mode, err := strconv.Atoi(r.FormValue("mode"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)
	if ext != ".png" && ext != ".PNG" {
		http.Error(rw, "Invalid format. Required png", http.StatusBadRequest)
	}

	// Opens or create input.png for incoming file
	f, err := os.OpenFile("input.png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	Transform(shapes, mode)
	http.ServeFile(rw, r, "out.png")
}

// Transform transforms input image
func Transform(shapes int, mode int) {
	_, err := primitive.DoPrimitive("input.png", "out.png", shapes, mode)
	if err != nil {
		log.Fatal(err)
	}
}
