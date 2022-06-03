package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p := pessoa{"Afonso", "Leal"}

	fmt.Println(p.getNomeCompleto())

	p.setNomeCompleto("Olívia Leal")
	fmt.Println(p.getNomeCompleto())
}
