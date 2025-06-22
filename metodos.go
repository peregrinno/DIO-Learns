package main

import "fmt"

type CarroTop struct {
	Marca  string
	Modelo string
	Ano    int
}

func (c CarroTop) Andar() {
	fmt.Println("O carro", c.Marca, "esta andando")
}

func (c CarroTop) Abastecer() {
	fmt.Println("O carro", c.Marca, "esta abastecendo")
}

func (c CarroTop) Parar() {
	fmt.Println("O carro", c.Marca, "esta parado")
}

type Veiculo interface {
	Andar()
	Abastecer()
	Parar()
}

func m(v Veiculo) {
	v.Andar()
	v.Abastecer()
	v.Parar()
}

func metodos() {
	c := CarroTop{
		Marca:  "Fiat",
		Modelo: "Uno",
		Ano:    2022,
	}
	m(c)
}
