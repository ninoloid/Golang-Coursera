package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Name struct
type Name struct {
	fname string
	lname string
}

func stringGenerator(s string) string {
	if len(s) <= 20 {
		return s
	}

	runes := []rune(s)
	return string(runes[0:20])
}

func main() {
	sliceOfName := make([]Name, 0)

	filenameScanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the file name: ")
	filenameScanner.Scan()
	filename := filenameScanner.Text()

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		var tempObj Name
		splittedString := strings.Split(fileScanner.Text(), " ")
		tempObj.fname, tempObj.lname = stringGenerator(splittedString[0]), stringGenerator(splittedString[1])
		sliceOfName = append(sliceOfName, tempObj)
	}

	for _, name := range sliceOfName {
		fmt.Println("\nFirst Name: ", name.fname)
		fmt.Println("Last Name: ", name.lname)
	}
}
