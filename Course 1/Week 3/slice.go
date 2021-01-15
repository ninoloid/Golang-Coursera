package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()

	mySlice := make([]int, 3)

	index := 0
	for scanner.Scan() {
		userInput := strings.ToLower(scanner.Text())
		if userInput == "x" {
			break
		}
		if convertedInput, error := strconv.Atoi(userInput); error == nil {
			if index < 3 {
				mySlice[index] = convertedInput
				index++
			} else {
				mySlice = append(mySlice, convertedInput)
				index++
			}
		}
	}

	sort.Ints(mySlice)
	fmt.Println(mySlice)
}
