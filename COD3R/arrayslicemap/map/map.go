package main

import "fmt"

func main() {
	//map`s devem ser sempre iniciados
	aprovados := make(map[int]string)

	aprovados[123] = "Afonso"
	aprovados[456] = "Juliana"
	aprovados[789] = "Olívia"

	fmt.Println(aprovados)

	for chave, nome := range aprovados {
		fmt.Printf("%s - CPF %d\n", nome, chave)
	}

	delete(aprovados, 123)
	fmt.Println(aprovados)
}
