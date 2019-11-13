package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numOfGoroutines = 4
	numOfTasks = 10
)

var waitGroup sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	tasksChannel := make(chan int, numOfTasks)
	waitGroup.Add(numOfGoroutines)

	for i := 0; i < numOfGoroutines; i++ {
		go worker(tasksChannel, i)
	}

	for i := 1; i <= numOfTasks; i++ {
		tasksChannel <- i
	}

	close(tasksChannel)
	waitGroup.Wait()
}

func worker(tasksChannel chan int, name int) {
	defer waitGroup.Done()

	for {
		task, ok := <-tasksChannel

		if !ok {
			fmt.Printf("worker %d finished and exited\n", name)
			break
		}

		fmt.Printf("worker %d is working on task %d\n", name, task)

		randNum := rand.Int63n(100)
		time.Sleep(time.Duration(randNum) * time.Millisecond)

		fmt.Printf("worker %d has finished his work\n", name)

	}
}