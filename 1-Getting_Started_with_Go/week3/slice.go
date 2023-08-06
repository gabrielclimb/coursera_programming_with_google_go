package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const sliceSize = 3

func main() {
	var slice = make([]int, sliceSize)

	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for {
		fmt.Printf("Type any int number. Type 'X' to exit:\n")
		scanner.Scan()
		var answer = scanner.Text()
		if answer == "X" {
			break
		}

		answerConverted, err := strconv.Atoi(answer)
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, answerConverted)
		sort.Sort(sort.IntSlice(slice))
		var sliceSorted = slice[sliceSize:len(slice)]
		fmt.Println(sliceSorted)

	}

}
