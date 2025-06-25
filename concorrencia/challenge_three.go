package main

import (
	"fmt"
	"time"
)

func challengeThree() {
	ping := make(chan struct{})
	pong := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			<-ping
			fmt.Printf("#%d: ping\n", i*2)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			<-pong
			fmt.Printf("#%d: pong\n", i*2+1)
			if i < 4 {
				ping <- struct{}{}
			}
		}
	}()

	ping <- struct{}{}

	time.Sleep(time.Second)
}

func main() {
	challengeThree()
}
