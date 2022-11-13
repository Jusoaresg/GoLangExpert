package main

import "fmt"

const a = "Hello, World1"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray)
	fmt.Println(len(meuArray)) //Len pega o tamanho do array
	fmt.Println(meuArray[2])
	fmt.Println(meuArray[len(meuArray)-1]) //Ultimo indice do array

	for i, v := range meuArray {
		//%d = digito
		//%v = string
		// \n = quebra linha
		fmt.Printf("O valor do indice é %d e o valor é %d \n", i, v)

	}
}
