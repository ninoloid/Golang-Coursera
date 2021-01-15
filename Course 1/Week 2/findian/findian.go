package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findian(x string) bool {
	return strings.HasPrefix(x, "i") && strings.Contains(x, "a") && strings.HasSuffix(x, "n")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var isFound bool
	userInput := strings.ToLower(scanner.Text())
	isFound = findian(userInput)

	if isFound {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
