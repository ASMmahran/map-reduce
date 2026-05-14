package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Result struct {
	Counts map[string]int `json:"counts"`
}

var filePath string

func countGenome() map[string]int {

	result := make(map[string]int)

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error:", err)
		return result
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		// Ignore FASTA headers
		if strings.HasPrefix(line, ">") {
			continue
		}

		line = strings.ToLower(line)

		for _, char := range line {

			if char == 'a' ||
				char == 't' ||
				char == 'g' ||
				char == 'c' {

				result[string(char)]++
			}
		}
	}

	return result
}

func handleCount(w http.ResponseWriter, r *http.Request) {

	result := Result{
		Counts: countGenome(),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}

func main() {

	// CHANGE THIS FOR EACH DEVICE
	port := "8081"

	// CHANGE FILE PATH
	filePath = "data/genome1.fa"

	http.HandleFunc("/count", handleCount)

	fmt.Println("Worker running on port", port)

	http.ListenAndServe(":"+port, nil)
}
