package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	userData := map[string]string{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please input your name: ")
	scanner.Scan()
	userData["name"] = scanner.Text()

	fmt.Print("Please input your address: ")
	scanner.Scan()
	userData["address"] = scanner.Text()

	userDataJSON, err := json.MarshalIndent(userData, "", "  ")

	if err == nil {
		fmt.Println("\nHere's the result:")
		fmt.Println(string(userDataJSON))
	} else {
		fmt.Println("Something wrong")
	}
}
