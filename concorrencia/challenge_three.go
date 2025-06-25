package main

import (
	"fmt"
)

func challengeThree() {
	c1 := make(chan string, 10)
	c2 := make(chan string, 10)

	go func() {
		for i := 0; i < 5; i++ {
			c1 <- "ping"
		}
		close(c1)
	}()
	go func() {
		for i := 0; i < 5; i++ {
			c2 <- "pong"
		}
		close(c2)
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg1, ok := <-c1:
			if ok {
				fmt.Printf("#%d: %s\n", i, msg1)
			}
		case msg2, ok := <-c2:
			if ok {
				fmt.Printf("#%d: %s\n", i, msg2)
			}
		}
	}
}

func main() {
	challengeThree()
}
