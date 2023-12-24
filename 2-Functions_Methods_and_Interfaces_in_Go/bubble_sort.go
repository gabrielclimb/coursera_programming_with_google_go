package main

import "fmt"

func BubbleSort(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		for j := 0; j < len(*slice)-1; j++ {
			Swap(slice, j)
		}
	}
}

func Swap(slice *[]int, i int) {
	if (*slice)[i] > (*slice)[i+1] {
		(*slice)[i], (*slice)[i+1] = (*slice)[i+1], (*slice)[i]
	}
}

func ReadSlicFromUser() []int {
	var n int
	fmt.Println("Enter the number of elements you want to enter: ")
	fmt.Scan(&n)
	slice := make([]int, n)
	fmt.Println("Enter the elements: ")
	for i := 0; i < n; i++ {
		fmt.Print("Enter element ", i+1, ": ")
		fmt.Scan(&slice[i])
	}
	return slice
}

func main() {
	slice := ReadSlicFromUser()
	BubbleSort(&slice)
	fmt.Println(slice)

}
