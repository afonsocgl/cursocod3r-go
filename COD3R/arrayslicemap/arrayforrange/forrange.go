package main

import "fmt"

func main() {
	// Essa notacao serve para criar um array sem tamanho definido, no caso o compilador que vai contar a quantidade de informacoes que foram colocadas dentro das chaves
	numeros := [...]int{1, 2, 3, 4, 5}

	//for range, utilizado para pegar de um so vez o indice e o valor do indice
	for indice, valor := range numeros {
		fmt.Printf("O indice %d tem valor %d\n", indice, valor)
	}

	// Caso nao queira utilizar um dos valores (indice ou valor) basta utilizar no lugar deles _, assim deixamos de pegar este valor...

	for _, valor := range numeros {
		fmt.Println(valor)
	}
}
