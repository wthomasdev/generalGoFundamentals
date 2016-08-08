package main

import "fmt"

func main() {
	// leagueTitles := make(map[string]int)
	// leagueTitles["Sunderland"] = 6
	// leagueTitles["Newcastle"] = 4
	// leagueTitles["Man City"] = 3
	// recentHead2Head := map[string]int{
	// 	"Sunderland": 5,
	// 	"Newcastle":  0,
	// }
	// fmt.Println(recentHead2Head)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	testMap["B"] = 99999
	testMap["F"] = 1987

	for key, value := range testMap {
		fmt.Printf("\nKey is: %v Value is: %v", key, value)
	}
}
