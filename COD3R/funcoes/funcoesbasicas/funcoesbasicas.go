package main

import "fmt"

func f1() {
	fmt.Println("Func basica")
}

func f2(p1, p2 string) {
	fmt.Printf("%s %s", p1, p2)
}

func f3() string {
	return "Func 3"
}

func main() {
	f1()
	f2("Afonso", "Leal")
	fmt.Println(f3())
}
