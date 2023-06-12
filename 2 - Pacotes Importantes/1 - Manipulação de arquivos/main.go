package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Criação de um arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//Gravando coisas no arquivo e pegando o tamanho dele

	/*
		Caso queira gravar string
	*/

	//tamanho, err := f.WriteString("Hello, World!")

	/*
		Caso queira gravar bytes
	*/

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho)

	//Para fechar o arquivo
	f.Close()

	//Para ler os arquivos
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	/*Lembrando de sempre conventer o tipo, por que no arquivo são bytes*/
	fmt.Println(string(arquivo))

	//Para fazer uma leitura de pouco em pouco
	//abrir o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//Pacote do Golang, usado para podermos usar buffers
	reader := bufio.NewReader(arquivo2)

	//Fazer o buffer ler de 10 em 10 bytes por vez
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		//Os dois pontos significa a posição
		fmt.Println(string(buffer[:n]))
	}

	//Remover o arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
