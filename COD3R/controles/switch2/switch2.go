package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Tipo desconhecido"
	}
}

func main() {
	fmt.Println(tipo(1234))
	fmt.Println(tipo(12.34))
	fmt.Println(tipo("aHKGJahk"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
