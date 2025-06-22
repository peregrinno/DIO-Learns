package main

import "fmt"

func fatias() {
	arr := make([]int, 3, 9)
	arr2 := make([]int, 5)
	copy(arr2, arr)
	fmt.Println(arr)
	fmt.Println(arr2)
}
