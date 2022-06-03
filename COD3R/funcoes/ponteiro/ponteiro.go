package main

import "fmt"

func incremento(n *int) {
	*n++
}

func main() {
	x := 1

	incremento(&x)
	fmt.Println(x)
}
