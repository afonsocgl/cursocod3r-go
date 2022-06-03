package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice, basta nao colocar o tamanho que ele passa a ser um slice

	//SLice nao é um array, apenas reflete um trecho de um array

	s2 := a1[1:2]
	fmt.Println(a1, s1, s2)

}
