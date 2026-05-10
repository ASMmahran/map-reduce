package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func saveFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("File") // مهم
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	os.MkdirAll("uploads", os.ModePerm)

	dst, err := os.Create("uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File %s uploaded successfully", handler.Filename)
}

func main() {
	servar := http.Server{
		Addr: "10.65.103.159:9080",
	}
	fmt.Println("Servar runining on ", servar.Addr)
	http.HandleFunc("/save", saveFile)
	servar.ListenAndServe()

}
