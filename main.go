package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	filepath := flag.Arg(0)
	if filepath == "" {
		fmt.Println("Please specify a file to convert ")
		os.Exit(1)
	}

	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		fmt.Println(fmt.Errorf("Could not open '%s': %w", filepath, err))
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	var records [][]string
	records, err = reader.ReadAll()
	if err != nil {
		fmt.Println(fmt.Errorf("Failed to read contents of '%s': %w", filepath, err))
		os.Exit(1)
	}

	for _, record := range records {
		fmt.Println(fmt.Printf("%s", record))
	}
}
