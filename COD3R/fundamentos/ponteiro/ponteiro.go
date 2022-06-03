package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	//O * representa que voce esta trabalhando com um espaco de memoria, declarando um ponteiro
	//nil (nulo) resprresnta que o enderco de emoria ainda nao existe

	p = &i // O & pega o endereco de memoria de uma variavel qualquer, logo estamos atribuindo ao ponteiro p o endereco de memoria da variavel i
	*p++   // Estamos pegando o valor que esta no espaco de memoria apontado pelo ponteiro p (valor da variavel i) e fazendo um incremento
	i++    // Estamos fazendo um incremento diretamente na variavel

	fmt.Println(p, *p, i, &i)

}
