package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.7)
}
