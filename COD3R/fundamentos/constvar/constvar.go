package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Println(area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 4
		d = 6
	)

	e, f, g := "lkjdslkj", false, 200

	fmt.Println(a, b, c, d, e, f, g)
}
