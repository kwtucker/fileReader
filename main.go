package fileReader

import (
	"bufio"
	"fmt"
	"github.com/deckarep/gosx-notifier"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// ReadFile reades files and writes
func ReadFile(filename string, notify bool) []string {

	// Set notification values
	notifyTitle := "Save Files"
	notifyMessage := "Forgit Event In 15 seconds"

	var (
		slice        = []string{}
		concatSwitch = false
		temp         string
	)
	// Check which OS the user is using then notify them.
	if notify {
		switch runtime.GOOS {
		case "darwin":
			// osx notification
			note := gosxnotifier.NewNotification(notifyMessage)
			note.Title = notifyTitle
			note.AppIcon = "logo_icon.png"
			err := note.Push()
			if err != nil {
				fmt.Println("Uh oh!")
			}
		case "linux":
			exec.Command("notify-send", "-i", "./logo_icon.png", notifyTitle, notifyMessage, "-u", "critical").Run()
		}

		// Delay the app from reading and writing for
		time.Sleep(time.Second * 15)
	}

	// opens file with read and write permissions
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	// Will close the file after main function is finished
	defer file.Close()

	// Grab contents of file
	read, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// File content to string
	newContents := string(read)

	// Setting a scanner buffer layer
	scanner := bufio.NewScanner(file)

	// loop through the file.
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "(" && i+1 < len(line) && string(line[i+1]) == ":" && concatSwitch == false {
				concatSwitch = true
				lineDelimiterStart := strings.Index(line, "(:")
				lineSliceStart := line[lineDelimiterStart : lineDelimiterStart+2]
				newContents = strings.Replace(newContents, lineSliceStart, " ", -1)
				i++
			} else {
				if concatSwitch == true {
					if string(line[i]) == "(" && i+1 < len(line) && string(line[i+1]) == ":" && concatSwitch == true {
						temp = ""
						concatSwitch = true
						lineDelimiterStart := strings.Index(line, "(:")
						lineSliceStart := line[lineDelimiterStart : lineDelimiterStart+2]
						newContents = strings.Replace(newContents, lineSliceStart, lineSliceStart, -1)
						i++
					} else if string(line[i]) == ":" && i+1 < len(line) && string(line[i+1]) == ")" {
						concatSwitch = false
						// Removes the end delimiter by finding the start index,
						// making a slice of the delimiter
						// and replacing it will empty strings
						lineDelimiterEnd := strings.Index(line, ":)")
						lineSliceEnd := line[lineDelimiterEnd : lineDelimiterEnd+2]
						newContents = strings.Replace(newContents, lineSliceEnd, " ", -1)
						slice = append(slice, temp)
						temp = ""
						i++
					} else {
						temp += string(line[i])
					}
				}
			}
		}
	}

	err = ioutil.WriteFile(filename, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}

	// save changes
	file.Sync()
	return slice
}
