package main

import "fmt"

func main() {
	//Array é homogeneo (mesmos tipo de dados) e estático (tamanho fixo)
	var notas [3]float64

	notas[0], notas[1], notas[2] = 5.0, 7.3, 9.4

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total = total + notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("O aluno somou %.2f, em %d avaliações, portanto sua média é: %.2f", total, len(notas), media)

}
