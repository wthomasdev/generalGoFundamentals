package main

import (
	"fmt"
)

func main() {

	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	WillCourseMega := courseMeta{
		Author: "Will Thomas",
		Level:  "Intermediate",
		Rating: 5,
	}

	WillCourseMega.Author = "Jeremy Irons"

	fmt.Println("\n Will's Course Author is:", WillCourseMega.Author)

}
