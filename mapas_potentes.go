package main

import "fmt"

type Carro struct {
	Marca  string
	Modelo string
	Ano    int
}

type Moto struct {
	Marca  string
	Modelo string
	Ano    int
}

type Veiculos struct {
	Carros []Carro
	Motos  []Moto
}

func mapasPotentes() {
	v := Veiculos{
		Carros: []Carro{
			{Marca: "Fiat", Modelo: "Uno", Ano: 2022},
			{Marca: "Chevrolet", Modelo: "Camaro", Ano: 2022},
			{Marca: "Ford", Modelo: "Mustang", Ano: 2022},
		},
		Motos: []Moto{
			{Marca: "Honda", Modelo: "CBR", Ano: 2022},
			{Marca: "Yamaha", Modelo: "YZF", Ano: 2022},
			{Marca: "Suzuki", Modelo: "Gixxer", Ano: 2022},
		},
	}

	v.Carros = append(v.Carros, Carro{Marca: "Volkswagen", Modelo: "Golf", Ano: 2022})
	v.Motos = append(v.Motos, Moto{Marca: "Harley-Davidson", Modelo: "Street 750", Ano: 2022})

	fmt.Println("Carros:", v.Carros)
	fmt.Println("Motos:", v.Motos)

}
