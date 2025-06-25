package main

import "fmt"

func media(alunos []float64) (float64, error) {
	total := 0.0
	for _, aluno := range alunos {
		if aluno < 0 || aluno > 10 {
			return 0, fmt.Errorf("Nota invalida: %f", aluno)
		}
		total += aluno
	}
	return total / float64(len(alunos)), nil
}

func learnOne() {
	alunos := []float64{7.0, 7.5, 5.0, 7.5, 10.4}
	media, err := media(alunos)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Media: ", media)
}

/* =============== */

// Closure: Chamar variaveis externas para dentro de um função
func learnTwo() {
	cont := 0
	incrementar := func() int { cont++; return cont }
	fmt.Println(incrementar())
	fmt.Println(incrementar())
	fmt.Println(incrementar())
}

/* =============== */

// Recursividade: Chamar uma função dentro de outra função
func learnThree(n int) int {
	if n == 0 {
		return 1
	}
	return n * learnThree(n-1)
}

/* =============== */

// Defer: Atrasar a execução de uma função
func learnFour() {
	fmt.Println("Segunda")
}

func learnFive() {
	fmt.Println("Terca")
}

/* =============== */

// Panic: Gerar um erro e parar a execução
func learnSix() {
	defer fmt.Println("Tudo bom")
	fmt.Println("Opa")
	panic("Erro")
}

/* =============== */

// Recover: Capturar um erro e continuar a execução
func learnSeven() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recuperado:", r)
		}
	}()
	fmt.Println("Opa")
	panic("Erro")
}

/* =============== */

// Ponteiro e Operador
func pointer(xPtr *int) *int {
	*xPtr = 20
	return xPtr
}

func learnEight() {
	x := 10
	pointer(&x)
	fmt.Println("Valor de x:", x)
}

/* =============== */

func main() {
	//learnOne()
	//learnTwo()
	//fmt.Println(learnThree(9))
	//defer learnFive()
	//learnFour()
	//learnSix()
	//learnSeven()
	learnEight()
}
