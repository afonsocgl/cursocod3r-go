package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo, como já definimos as variaveis de retorno nao precisamos declaralas novamente no return
}

func main() {
	x, y := trocar(1, 2)

	fmt.Println(x, y)
}
