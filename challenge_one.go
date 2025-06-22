package main

import (
	"fmt"
)

func main() {
	// Conversão do valor dor ponto de ebulição da água para Celsius
	ebulicaoK := 212.0
	ebulicaoC := (ebulicaoK - 32) * 5 / 9
	fmt.Printf("O valor do ponto de ebulição da água em Celsius é: %g", ebulicaoC)
}
