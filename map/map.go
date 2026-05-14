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

	resp, err := http.Get(url)

	if err != nil {
		ch <- WorkerResponse{Err: err}
		return
	}

	defer resp.Body.Close()

	var result Result

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		ch <- WorkerResponse{Err: err}
		return
	}

	ch <- WorkerResponse{
		Data: result.Counts,
	}
}

func main() {

	workers := []string{
		"http://192.168.1.10:8081/count",
		"http://192.168.1.11:8082/count",
		"http://192.168.1.12:8083/count",
	}

	finalResult := make(map[string]int)

	ch := make(chan WorkerResponse)

	var wg sync.WaitGroup

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
			fmt.Println("Worker Error:", response.Err)
			continue
		}

		for char, count := range response.Data {
			finalResult[char] += count
		}
	}

	total := 0

	fmt.Println("\nFinal Result:")

	for char, count := range finalResult {

		fmt.Printf("%s : %d\n", char, count)

		total += count
	}

	fmt.Println("\nTotal:", total)
}