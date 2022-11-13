package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(sum(51, 2))

	//Usando erros
	valor, err := sumError(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
	//fmt.Println(sumError(51, 2))

}

func sum(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

// Utilizando erros
func sumError(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
