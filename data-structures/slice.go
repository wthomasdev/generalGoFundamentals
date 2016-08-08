package main

import (
	"fmt"
)

func main() {

	// mySlice := make([]int, 1, 4)

	// for i := 1; i < 17; i++ {
	// 	mySlice = append(mySlice, i)
	// 	fmt.Println("Printing a num", i)
	// 	fmt.Printf("\nCapacity is: %d", cap(mySlice))
	// }

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice)
	newSlice := []int{99, 33, 22}

	for _, i := range mySlice {
		fmt.Println("For range loop:", i)
	}

	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

}
