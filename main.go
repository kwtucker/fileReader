package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func main() {
	// the filename or path that is passed in. Then reads the file into memory.
	filename := os.Args[1]
	_, err := ioutil.ReadFile(filename)
	checkError(err)

	// opens file with read and write permissions
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	checkError(err)

	read, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// File content
	newContents := string(read)

	// It’s idiomatic to defer a Close immediately after opening a file.
	defer file.Close()

	// Setting a scanner and writer.
	scanner := bufio.NewScanner(file)

	// Initializing a array.
	var slice = []string{}

	// loop through the file.
	for scanner.Scan() {
		line := scanner.Text()

		// Look through the line and check for the start and end delimiter.
		// Then slice the start to the end.
		if strings.Contains(line, "(:") && strings.Contains(line, ":)") {
			lineDelimiter_Start := strings.Index(line, "(:")
			lineDelimiter_End := strings.Index(line, ":)")

			lineSlice := line[lineDelimiter_Start+2 : lineDelimiter_End]
			slice = append(slice, lineSlice)

			comment := "(:" + lineSlice + ":)"

			// Replace in the file string the comment with a white space
			newContents = strings.Replace(newContents, comment, " ", -1)
		}
	}

	err = ioutil.WriteFile(filename, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("slice", slice)

	// save changes
	file.Sync()

}
