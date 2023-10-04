package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	path "path/filepath"
	"strings"
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

	writeFile, err := os.Create(dir + "/" + strings.Replace(readFilename, ".csv", ".org", 1))
	defer writeFile.Close()

	reader := csv.NewReader(file)

	var record []string
	if *noHeader == false {
		record, err = reader.Read()
        writeFile.WriteString("||\n| ")
        writeFile.WriteString(strings.Join(record, " | ")) 
        writeFile.WriteString(" |\n||\n")
	}

	for {
		record, err = reader.Read()
		if err == io.EOF {
            writeFile.WriteString("||")
			fmt.Println("Successfully created")
			break
		}

		if err != nil {
			exit(fmt.Errorf("Failed to read contents of '%s': %w", readFile, err).Error())
		}

        writeFile.WriteString("| ")
        writeFile.WriteString(strings.Join(record, " | ")) 
        writeFile.WriteString(" |\n")
	}
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
