package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

var waitGroup sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}
	
func main() {
	waitGroup.Add(2)
	court := make(chan int)

	go player("kilu", court)
	go player("clinton", court)

	court <- 1
	waitGroup.Wait()
}

func player(name string, court chan int) {
	defer waitGroup.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s won the game\n", name)
			return 
		}
		n := rand.Intn(100)
		if n %13 == 0 {
			fmt.Printf("player %s lost the game\n", name)
			close(court)
			break
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball
	}
}