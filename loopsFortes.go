package main

import "fmt"

func loopsFortes() {
	i := 0
	for i < 3 {
		fmt.Println("Hello World")
		i++
	}
}

func loopHieraquia() {
	for i := 1; i <= 12; i++ {
		for j := 1; j <= 60; j++ {
			fmt.Printf("%d:%d\n", i, j)
		}
	}
}
