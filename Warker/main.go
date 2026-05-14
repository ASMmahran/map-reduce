package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Result struct {
	Counts map[string]int `json:"counts"`
}

var filePath string

// Count genome characters
func countGenome() map[string]int {

	fmt.Println("[WORKER] Starting Genome Processing...")

	result := make(map[string]int)

	file, err := os.Open(filePath)

	if err != nil {

		fmt.Println("[ERROR] Cannot Open File:", err)

		return result
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := 0

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

		lines++
	}

	fmt.Println("[WORKER] Finished Processing")
	fmt.Println("[WORKER] Total Lines Processed:", lines)

	return result
}

// Handle request from master
func handleCount(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n===================================")
	fmt.Println("[WORKER] Request Received From Master")
	fmt.Println("[TIME]", time.Now().Format("15:04:05"))
	fmt.Println("[CLIENT]", r.RemoteAddr)
	fmt.Println("===================================")

	start := time.Now()

	result := Result{
		Counts: countGenome(),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)

	fmt.Println("[WORKER] Result Sent Successfully")

	fmt.Println("[WORKER] Processing Time:",
		time.Since(start))
}

func main() {

	// CHANGE FOR EACH DEVICE
	port := "8081"

	// CHANGE FILE FOR EACH DEVICE
	filePath = "data/genome1.fa"

	fmt.Println("===================================")
	fmt.Println(" Distributed MapReduce WORKER ")
	fmt.Println("===================================")

	fmt.Println("[STATUS] Worker Started Successfully")
	fmt.Println("[PORT] Running On Port:", port)
	fmt.Println("[FILE] Reading File:", filePath)
	fmt.Println("[TIME]", time.Now().Format("15:04:05"))
	fmt.Println("===================================")

	http.HandleFunc("/count", handleCount)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {

		fmt.Println("[ERROR] Server Failed:", err)
	}
}
