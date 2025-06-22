package main

import "fmt"

func oplogicos() {
	x := -10
	y := 20

	if x > 0 && y > 0 {
		fmt.Println("x e y são positivos")
	}

	if x > 0 || y > 0 {
		fmt.Println("x ou y são positivos")
	}

	if !(x > 0) {
		fmt.Println("x não é positivo")
	}
}
