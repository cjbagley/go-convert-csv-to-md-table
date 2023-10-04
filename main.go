package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	path "path/filepath"
)

func main() {
	noHeader := flag.Bool("no-header", false, "Pass if no header row (otherwise defaults to assuming one is present)")
	flag.Parse()

	readFile := flag.Arg(0)
	if readFile == "" {
		exit("Please specify a file to convert")
	}

	dir := path.Dir(readFile)
	readFilename := path.Base(readFile)

	file, err := os.Open(readFile)
	defer file.Close()

	if err != nil {
		exit(fmt.Errorf("Could not open '%s': %v", readFile, err).Error())
	}

	writeFile, err := os.Create(dir + "/updated-" + readFilename)
	defer writeFile.Close()

	reader := csv.NewReader(file)
	writer := csv.NewWriter(writeFile)
	defer writer.Flush()

	var record []string
	if *noHeader == false {
		record, err = reader.Read()
        writer.Write(processRow(record))
	}

	for {
		record, err = reader.Read()
		if err == io.EOF {
			fmt.Println("Successfully created")
			break
		}

		if err != nil {
			exit(fmt.Errorf("Failed to read contents of '%s': %w", readFile, err).Error())
		}

        writer.Write(processRow(record))
	}
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func processRow(row []string) []string {
    firstColumn := true 
    for _, r := range row {
        if firstColumn {
            
        }
		fmt.Println(r)
	}
	return row
}
