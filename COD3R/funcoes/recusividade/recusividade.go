package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido %d", n)
	case n == 1:
		return 1, nil
	default:
		nFatorial, _ := fatorial(n - 1)
		return n * nFatorial, nil
	}
}

func main() {
	resultado, err := fatorial(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("O resultado do fatorial é %d", resultado)
	}
}
