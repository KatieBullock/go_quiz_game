package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func quiz(timeLimit int, records [][]string) int {
	timeout := time.After(time.Duration(timeLimit) * time.Second)
	correct := 0

	for i, record := range records {
		fmt.Printf("Problem #%d: %s = ", i+1, record[0])

		inputChan := make(chan string)

		go func() {
			var answer string
			fmt.Scanln(&answer)
			inputChan <- strings.TrimSpace(answer)
		}()

		select {
		case answer := <-inputChan:
			if answer == record[1] {
				correct++
			}
		case <-timeout:
			return correct
		}
	}

	return correct
}

func main() {
	csvPtr := flag.String("csv", "problems.csv", "a csv in the format of 'question,answer'")
	timeLimitPtr := flag.Int("limit", 30, "a time limit in seconds")
	shufflePtr := flag.Bool("shuffle", false, "a 'true' or 'false' to shuffle questions")

	flag.Parse()

	file, err := os.Open(*csvPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if *shufflePtr {
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	fmt.Println("Press [Enter] to Begin!")
	fmt.Scanln()

	correct := quiz(*timeLimitPtr, records)

	fmt.Printf("\nYou scored %d out of %d\n", correct, len(records))
}
