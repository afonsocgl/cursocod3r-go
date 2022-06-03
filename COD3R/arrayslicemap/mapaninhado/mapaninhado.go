package main

import "fmt"

func main() {
	funcionarioPorLetra := map[string]map[string]float64{
		"A": {
			"Afonso Leal":   1762.87,
			"Antonio Jorge": 18762.09,
		},
		"B": {
			"Barbara Galvao": 198719.09,
			"Breno Gama":     98712.09,
		},
		"C": {
			"Carlos Galvao": 98712.09,
		},
	}

	for chave, valor := range funcionarioPorLetra {
		fmt.Printf("Funcionarios da letra %s sao %v\n", chave, valor)
	}
}
