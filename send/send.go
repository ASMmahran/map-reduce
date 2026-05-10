package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func readFile(filePath string, ch chan []byte) {

	file, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		ch <- nil
		return
	}

	ch <- file
}

func SendFiles(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"D:\\university\\fourth\\2nd\\DDB\\Section\\Practice-DDB-main\\Section 5\\Scripts\\Data\\gene.fna",
		"D:\\university\\fourth\\2nd\\DDB\\Section\\Practice-DDB-main\\Section 5\\Scripts\\Data\\gene2.fna",
		"D:\\university\\fourth\\2nd\\DDB\\Section\\Practice-DDB-main\\Section 5\\Scripts\\Data\\fgenome.fa",
	}

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	for _, filePath := range files {

		ch := make(chan []byte)

		go readFile(filePath, ch)

		data := <-ch

		if data == nil {
			fmt.Println("Failed to read:", filePath)
			continue
		}

		fileName := filepath.Base(filePath)

		part, err := writer.CreateFormFile("Files", fileName)

		if err != nil {
			fmt.Println("Error creating form file:", err)
			continue
		}

		_, err = io.Copy(part, bytes.NewReader(data))

		if err != nil {
			fmt.Println("Copy error:", err)
			continue
		}
	}

	writer.Close()

	req, err := http.NewRequest(
		"POST",
		"http://192.168.8.8:9080/save",
		&body,
	)

	if err != nil {
		fmt.Println("Error creating request")
		return
	}

	req.Header.Set(
		"Content-Type",
		writer.FormDataContentType(),
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request")
		return
	}

	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	fmt.Println("Response:", string(respBody))

	w.Write([]byte("Files sent successfully"))
}

func main() {

	servar := http.Server{
		Addr: "192.168.8.15:9080",
	}

	fmt.Println("Servar runining on ", servar.Addr)

	http.HandleFunc("/SendFiles", SendFiles)

	err := servar.ListenAndServe()

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
