package main

import "fmt"

func main() {

	// for timer := 10; timer >= 0; timer-- {
	// 	if timer == 0 {
	// 		fmt.Println("Boom!")
	// 		break
	// 	}
	// 	fmt.Println(timer)
	// 	time.Sleep(1 * time.Second)
	// }

	coursesInProg := []string{"Docker Deep Dive", "Will's Wonder Course", "Docker and Kubernetes"}

	coursesCompleted := []string{"Docker Deep Dive",
		"Go Fundamentals", "Puppet Fundamentals"}

	for _, i := range coursesInProg {
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\n Hey we found a clash.", i, "is in both lists")
			}
		}

	}

}
