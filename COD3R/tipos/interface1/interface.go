package main

import "fmt"

type imprimivel interface {
	imprimir() string
}

type livro struct {
	titulo string
	autor  string
}

func (l livro) imprimir() string {
	return fmt.Sprintf("O livro '%s' foi escrito por %s\n", l.titulo, l.autor)
}

type autor struct {
	nome  string
	idade int
}

func (a autor) imprimir() string {
	return fmt.Sprintf("O autor %s tem %d anos\n", a.nome, a.idade)
}

func main() {
	l := livro{"Harry Potter", "JK Roling"}
	a := autor{"Afonso leal", 30}
	Imprimir(a)
	Imprimir(l)

}

func Imprimir(i imprimivel) {
	fmt.Printf(i.imprimir())
}
