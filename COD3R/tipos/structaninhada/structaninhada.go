package main

import "fmt"

type item struct {
	itemID int
	qtd    int
	preco  float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorPedido() float64 { //apos o termo func o que estar dentro dos () é chamado de reciver
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{itemID: 1, qtd: 2, preco: 10.00},
			item{itemID: 2, qtd: 2, preco: 10.00},
			item{itemID: 3, qtd: 2, preco: 10.00},
		},
	}

	fmt.Printf("O valor total do pedido é R$%.2f", pedido.valorPedido())
}
