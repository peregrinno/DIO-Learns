package main

import (
	"fmt"
	"time"
)

func papaLeguas(number int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Papa ta correndo!", number)
	}
}

func usainBolt(number int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Usain ta correndo!", number)
	}
}

/*
func main() {
	go usainBolt(1)
	go papaLeguas(2)
	var input string
	fmt.Scan(&input)
}
*/

// Canais

func pingar(c chan string) {
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("ping %d", i)
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

/*
func main() {
	c := make(chan string)
	go pingar(c)
	imprimir(c)

	var input string
	fmt.Scanf("%s", &input)
}
*/

// Select

/*
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "c1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Recebido", msg1)
		case msg2 := <-c2:
			fmt.Println("Recebido", msg2)
		}
	}
}
*/
