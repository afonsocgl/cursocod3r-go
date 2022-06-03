package main

import "fmt"

func limiteCredito(salario int) string {
	defer fmt.Println("Analisando...")

	if salario > 5000 {
		fmt.Printf("Estamos analisando o salario de R$%d\n", salario)
		return "Limite de cédito aprovado"
	} else {
		fmt.Printf("Estamos analisando o salário de R$%d\n", salario)
		return "Limite de crédito reprovado"
	}
}

func main() {
	fmt.Println(limiteCredito(5001))
	fmt.Println(limiteCredito(4999))
}
