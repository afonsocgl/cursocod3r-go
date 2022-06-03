package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: Função com reciver (receptor)
//Receptor funciona como o "dono" da funcao
func (p produto) calcularDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := produto{
		nome:     "caderno",
		preco:    100.00,
		desconto: 0.1,
	}

	fmt.Println(produto1.calcularDesconto())
}
