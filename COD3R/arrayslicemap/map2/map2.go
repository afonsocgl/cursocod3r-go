package main

import "fmt"

func main() {
	funcionarioESalarios := map[string]float64{
		"Afonso Leal":  12321.3,
		"Juliana Leal": 19872987.98,
		"Olívia Leal":  271864.98,
	}

	fmt.Println(funcionarioESalarios)

	funcionarioESalarios["Ian Leal"] = 1872123.24
	delete(funcionarioESalarios, "Olívia Leal")
	for nome, salario := range funcionarioESalarios {
		fmt.Printf("%s tem salário R$%.2f\n", nome, salario)
	}
}
