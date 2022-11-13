package main

import "fmt"

func main() {
	//Maps - HashTables - Chave de Key
	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000
	fmt.Println(salarios["Wes"])

	/*
		Outras formas de criar um map (criando vazio):

		sal := make(map[string]int)
		sal1 := map[string]int{}
		sal1["Wesley"] = 1000

	*/

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d \n", nome, salario)
	}

	//_ é um blank identifier
	//Caso não queira usar o nome
	for _, salario := range salarios {
		fmt.Printf("O salario é %d \n", salario)
	}
}
