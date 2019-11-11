package main

import (
	"sync"
	"fmt"
)

func main() {
	var (
		waitGroup sync.WaitGroup
		mutex sync.Mutex
	)
	waitGroup.Add(2)

	go func () {
		defer waitGroup.Done()
		for i := 0; i < 3; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				//mutex is used to lock the shared resource, in this case
				// the standard output, so that only one goroutine can 
				// use it at a time
				// If the resource used is an integer or pointer 
				//atomic functions can also be used.
				mutex.Lock()
				fmt.Printf("%c", char)
				mutex.Unlock()
			}
			fmt.Printf("\t")
		}
		fmt.Println(" ")
	}()

	go func () {
		defer waitGroup.Done()
		for i := 0; i < 3; i++ {
			for char := 'A'; char < 'A'+26; char++ {
				mutex.Lock()
				fmt.Printf("%c", char)
				mutex.Unlock()
			}
			fmt.Printf("\t")
		}
		fmt.Println(" ")
	}()

	fmt.Println("about to print the alphabets")
	waitGroup.Wait()
}