package main

import (
	"io"
	"net/http"
)

func main() {

	//Criando uma chamada(request)
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	//Pegando o body
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	//Fechar o body para não vazar
	req.Body.Close()

}
