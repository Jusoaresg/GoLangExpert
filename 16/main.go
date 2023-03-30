package main

//Quando usar ponteiros

func soma(a, b int) int {
	a = 50
	return a + b
}

func somaPonteiro(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	//Ele sempre pega uma copia desses valores
	minhaVar1 := 10
	minhaVar2 := 20
	//Ou seja aqui é uma copia das duas variaveis
	println(soma(minhaVar1, minhaVar2))
	println(minhaVar1)

	println("-------Usando Ponteiro-------")

	//Aqui como é um ponteiro ele pega o valor na memoria
	println(somaPonteiro(&minhaVar1, &minhaVar2))
	println(minhaVar1)

}
