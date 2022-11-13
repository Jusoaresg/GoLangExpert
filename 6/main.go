package main

import "fmt"

func main() {
	//Slices
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)

	//: é basicamente um ponto de corte
	//Pega tudo antes do 0 ":0" ou seja nada
	fmt.Printf("len=%d cap=%d %v \n", len(s[:0]), cap(s[:0]), s[:0])

	//Pega tudo antes do 4 ":4"
	fmt.Printf("len=%d cap=%d %v \n", len(s[:4]), cap(s[:4]), s[:4])

	//Pega tudo depois do 2 ":2"
	fmt.Printf("len=%d cap=%d %v \n", len(s[2:]), cap(s[2:]), s[2:])

	//append adiciona um valor no slice
	s = append(s, 110)
	/*
		Os slices sempre apontam para um array
		Sempre que você adiciona um valor ele tem que criar um
		novo array dobrando a capacidade do anterior
		Por isso a capacidade do slice duplicou
	*/
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
