package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Mapper
func mapper(id int, text string, out chan map[string]int, wg *sync.WaitGroup) {

	defer wg.Done()

	result := make(map[string]int)

	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		result[word]++
	}

	fmt.Println("Mapper", id, "finished")

	out <- result
}

// Reducer
func reducer(in chan map[string]int, finalResult map[string]int, done chan bool) {

	for partial := range in {

		for word, count := range partial {

			finalResult[word] += count
		}
	}

	done <- true
}

func main() {

	filePath := "D:\\university\\fourth\\2nd\\DDB\\Section\\Practice-DDB-main\\Section 5\\Scripts\\Data\\genome.fa"

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	mapperChannel := make(chan map[string]int)

	done := make(chan bool)

	finalResult := make(map[string]int)

	var wg sync.WaitGroup

	// Start reducer
	go reducer(mapperChannel, finalResult, done)

	mapperID := 10

	// Read file line by line
	for scanner.Scan() {

		line := scanner.Text()

		wg.Add(1)

		go mapper(mapperID, line, mapperChannel, &wg)

		mapperID++
	}

	// Wait all mappers
	wg.Wait()

	// Close mapper channel
	close(mapperChannel)

	// Wait reducer
	<-done

	// Total words
	totalWords := 0

	for _, count := range finalResult {
		totalWords += count
	}

	fmt.Println("\n========================")
	fmt.Println("Total Words:", totalWords)
	fmt.Println("========================")

}
