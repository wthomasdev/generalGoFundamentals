package main

// if statement

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//
	// if firstRank, secondRank := 39, 614; firstRank < secondRank {
	// 	fmt.Println("\nFirst Course is doing better" + " than second course")
	// } else if firstRank > secondRank {
	// 	fmt.Println("\nOh dear... your first " + "course must be doing abysmally!")
	// } else {
	// 	fmt.Println("\nBoth courses are either" + " the same or something weird is going on")
	// }

	//Switch

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("we got an even number -", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got an odd number -", tmpNum)
	default:
		fmt.Println("Go is kinda screwy")
	}

}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
