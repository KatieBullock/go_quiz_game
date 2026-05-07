package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	csvPtr := flag.String("csv", "problems.csv", "a csv in the format of 'question,answer'")

	flag.Parse()

	file, err := os.Open(*csvPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // NOTE(Katie): What is happening here? The file is being closed, by why here and why defer?

	reader := csv.NewReader(file)

	correct := 0
	count := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		count++

		fmt.Printf("Problem #" + strconv.Itoa(count) + ": " + record[0] + " = ")

		var answer string
		fmt.Scanln(&answer)

		if answer == record[1] {
			correct += 1
		}
	}

	fmt.Println("You scored " + strconv.Itoa(correct) + " out of " + strconv.Itoa(count))
}
