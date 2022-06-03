package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func conceitoNota(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(8.0, 8.99):
		return "B"
	case n.entre(5.0, 7.99):
		return "C"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(conceitoNota(2.0))
}
