package main

import "fmt"

func soma(i ...int) int {
	total := 0
	for _, valor := range i {
		total += valor
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, valor := range i {
		total *= valor
	}
	return total
}

func subtrai(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, valor := range i[1:] {
		total -= valor
	}
	return total
}

func divide(i ...int) int {
	if len(i) == 0 {
		return 0
	}
	total := i[0]
	for _, valor := range i[1:] {
		if valor == 0 {
			return 0
		}
		total /= valor
	}
	return total
}

func main() {
	x := soma(1, 2, 3)
	y := multiplica(10, 10)
	z := subtrai(10, 10)
	w := divide(10, 10)
	fmt.Println(x, y, z, w)
}
