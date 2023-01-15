package main

//Ponteiros
//* = Refferencing, Pega o valor guardado na memoria

func main() {
	//Memória -> Endereço -> Valor
	a := 10
	var ponteiro = &a
	*ponteiro = 20
	b := &a
	println(*b)
}
