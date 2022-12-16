package main

//Closures ou Funções anonimas
//Basicamente uma função dentro de outra função

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(1, 3, 45, 6, 34) * 2
	}()

	fmt.Println(sum(1, 3, 45, 6, 34))
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
