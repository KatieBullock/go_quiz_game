package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvPtr := flag.String("csv", "problems.csv", "a csv in the format of 'question,answer'")
	limitPtr := flag.Int64("limit", 30, "time limit in seconds to answer questions")

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

	fmt.Println("Press [Enter] to Begin!")
	fmt.Scanln()

	timeout := time.After(time.Duration(*limitPtr) * time.Second)
	correct := 0

quizLoop:
	for i, record := range records {
		fmt.Printf("Problem # %d: %s = ", i+1, record[0])

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
			break quizLoop
		}
	}

	fmt.Printf("\nYou scored %d out of %d\n", correct, len(records))
}
