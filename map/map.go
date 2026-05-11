
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

	// Convert to lowercase
	text = strings.ToLower(text)

	// Split into characters (A, T, G, C)
	for _, char := range text {

		// Ignore spaces/newlines
		if char == ' ' || char == '\n' || char == '\r' {
			continue
		}

		result[string(char)]++
	}

	fmt.Println("Mapper", id, "finished")

	out <- result
}

// Reducer
func reducer(
	in chan map[string]int,
	finalResult map[string]int,
	done chan bool,
) {

	for partial := range in {

		for word, count := range partial {
			finalResult[word] += count
		}
	}

	done <- true
}

func main() {

	filePath := `D:\university\fourth\2nd\DDB\Section\Practice-DDB-main\Section 5\Scripts\Data\genome.fa`

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Increase scanner buffer size
	buf := make([]byte, 1024)
	scanner.Buffer(buf, 1024*1024)

	mapperChannel := make(chan map[string]int)

	done := make(chan bool)

	finalResult := make(map[string]int)

	var wg sync.WaitGroup

	// Start reducer
	go reducer(mapperChannel, finalResult, done)

	mapperID := 1

	// Read file line by line
	for scanner.Scan() {

		line := scanner.Text()

		// Skip FASTA headers
		if strings.HasPrefix(line, ">") {
			continue
		}

		wg.Add(1)

		go mapper(mapperID, line, mapperChannel, &wg)

		mapperID++
	}

	// Check scanner error
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
		return
	}

	// Wait all mappers
	wg.Wait()

	// Close channel
	close(mapperChannel)

	// Wait reducer
	<-done

	// Calculate total characters
	total := 0

	for _, count := range finalResult {
		total += count
	}

	fmt.Println("\n========================")
	fmt.Println("Total Characters:", total)
	fmt.Println("========================")

	// Print result
	fmt.Println("\nGenome Character Count:")

	for word, count := range finalResult {
		fmt.Printf("%s : %d\n", word, count)
	}
}