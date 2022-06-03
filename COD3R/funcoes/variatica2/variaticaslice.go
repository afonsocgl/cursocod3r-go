package main

import "fmt"

func listaAprovados(aprovados ...string) {
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s \n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Afonso", "Juliana", "Olívia"}

	listaAprovados(aprovados...)
}
