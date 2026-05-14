package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Result struct {
	Counts map[string]int `json:"counts"`
}

type WorkerResponse struct {
	Data map[string]int
	Err  error
}

func requestWorker(url string, ch chan WorkerResponse, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Connecting To:", url)

	resp, err := http.Get(url)

	if err != nil {

		ch <- WorkerResponse{Err: err}

		return
	}

	fmt.Println("Connected Successfully:", url)

	defer resp.Body.Close()

	var result Result

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {

		ch <- WorkerResponse{Err: err}

		return
	}

	fmt.Println("Data Received From:", url)

	ch <- WorkerResponse{
		Data: result.Counts,
	}
}
func main() {

	fmt.Println("=================================")
	fmt.Println(" Distributed MapReduce MASTER ")
	fmt.Println("=================================")

	workers := []string{
		"http://192.168.8.8:8081/count",
		"http://localhost:8082/count",
		"http://localhost:8083/count",
	}

	fmt.Println("\nWorkers Registered:")

	for _, worker := range workers {
		fmt.Println("->", worker)
	}

	finalResult := make(map[string]int)

	ch := make(chan WorkerResponse)

	var wg sync.WaitGroup

	fmt.Println("\nSending Requests To Workers...")

	for _, worker := range workers {

		wg.Add(1)

		go requestWorker(worker, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for response := range ch {

		if response.Err != nil {

			fmt.Println("Worker Failed:", response.Err)

			continue
		}

		fmt.Println("Result Received From Worker")

		for char, count := range response.Data {
			finalResult[char] += count
		}
	}

	total := 0

	fmt.Println("\n=================================")
	fmt.Println(" FINAL RESULT ")
	fmt.Println("=================================")

	for char, count := range finalResult {

		fmt.Printf("%s : %d\n", char, count)

		total += count
	}

	fmt.Println("\nTotal:", total)

	fmt.Println("\nMASTER FINISHED SUCCESSFULLY")
}
