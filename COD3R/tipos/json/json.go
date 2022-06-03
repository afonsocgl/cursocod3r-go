package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//Structy para Json
	p1 := produto{1, "notebook", 5000.00, []string{"promoção", "Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//Json para struct
	var p2 produto
	jsonString := `{"id":2, "nome":"Caderno", "preco":20.00, "tags":["Papelaria", "Promo"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
