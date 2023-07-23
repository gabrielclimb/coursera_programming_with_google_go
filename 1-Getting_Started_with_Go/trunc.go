package main

import "fmt"

func main() {
	var floatNumber float32
	fmt.Print("Enter a floating number:\n")
	fmt.Scan(&floatNumber)
	fmt.Print(int(floatNumber))
}
