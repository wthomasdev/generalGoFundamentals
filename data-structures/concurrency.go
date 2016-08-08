package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Hello Will")
	}()

	go func() {
		defer waitGrp.Done()

		fmt.Println("Will is learning Go")
	}()

	waitGrp.Wait()
}
