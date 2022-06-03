package main

import "fmt"

func notaConceito(n int) string {
	switch {
	case n >= 9 && n < 10:
		return "Conceito A"
	case n >= 7 && n < 9:
		return "Conceito B"
	case n >= 5 && n < 7:
		return "Conceito C"
	default:
		return "Reprovado"
	}
}

func main() {
	fmt.Println(notaConceito(9))
	fmt.Println(notaConceito(7))
	fmt.Println(notaConceito(5))
	fmt.Println(notaConceito(4))
}
