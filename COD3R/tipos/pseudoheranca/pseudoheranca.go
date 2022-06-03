package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

type onix struct {
	carro //Criando campos anonimos (não precisamos passar nenhum informacao adicional para ele), funciona como uma heranca
	turbo bool
}

func main() {
	c := onix{}
	c.nome = "Onix Plus"
	c.velocidade = 150
	c.turbo = true

	fmt.Printf("O carro %s tem turbo? %v\n Qual sua velocidade? %dkm/h", c.nome, c.turbo, c.velocidade)
}
