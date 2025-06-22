package main

import "fmt"

func estruturas() {
	for i := 1; i <= 100; i++ {
		switch i % 2 {
		case 0:
			fmt.Println("Par", i)
		case 1:
			fmt.Println("Impar", i)
		}
	}
}

/*
Forma 1
	switch i % 2 {
	case 0:
		fmt.Println("Par", i)
	case 1:
		fmt.Println("Impar", i)
	}
*/

/*
Forma 2

	switch {
	case i%2 == 0:
		fmt.Println("Par", i)
	default:
		fmt.Println("Impar", i)
	}
*/
