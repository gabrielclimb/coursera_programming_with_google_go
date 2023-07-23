package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter a input:\n")
	fmt.Scan(&input)

	lowerCaseInput := strings.ToLower(input)
	if strings.HasPrefix(lowerCaseInput, "i") && strings.HasSuffix(lowerCaseInput, "n") && strings.Contains(lowerCaseInput, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
