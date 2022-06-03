package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 //float64
	y := 2   //int

	fmt.Println(x / float64(y)) //Precisamos converter o um dos valores, pois so conseguimos fazer as operações com duas variaveis do mesmo tipo

	nota := 6.9
	notaFinal := int(nota) //Simplismente despreza o ponto flutuante
	fmt.Println(notaFinal)

	//Cuidado...
	fmt.Println("Teste " + string(97)) // O valor "97", que está sendo usado está apontando para a tabela ASC

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123)) //Biblioteca strconv que precisa ser importada, Itoa = Int to (para) a (string)

	// String pata int
	num, _ := strconv.Atoi("123") // Atoi = A (string) to int (essa função retorna um int e um erro, usando a _ estamos desconsiderando o retorno do erro)
	fmt.Println(num - 122)

}
