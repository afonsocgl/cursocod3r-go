package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 5 && n < 9 {
		return "B"
	} else {
		return "C"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.7))
	fmt.Println(notaParaConceito(5.0))
	fmt.Println(notaParaConceito(4.9))
}
