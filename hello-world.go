package main

import (
	"fmt"
	"os"
)

func main() {

	name := os.Getenv("USER")
	// course := "Intro to Go"

	// 	fmt.Println("\nHi", name, "you're currently watching", course)
	//
	// 	changeCourse(&course)
	//
	// 	fmt.Println("\n You are now watching course", course)

	for _, env := range os.Environ() {
		fmt.Println(env)
		fmt.Println(name)
	}

}

// func changeCourse(course *string) string {
//
// 	*course = "First Look: Muthatruckin Go Bro"
//
// 	fmt.Println("\nTrying to change your course to", *course)
//
// 	return *course
// }
