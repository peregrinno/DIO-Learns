package main

import "fmt"

func challengeTwo() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("PinPan")
		case i%3 == 0:
			fmt.Println("Pin")
		case i%5 == 0:
			fmt.Println("Pan")
		default:
			fmt.Println(i)
		}
	}
}
