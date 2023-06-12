package main

import (
	"io"
	"net/http"
)

//Defer é algo que a gente atrasa alguma coisa

func main() {

	//Criando uma chamada(request)
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	/*
		O defer basicamente atrasa a execução da linha,
		ou seja, após acabar todo o comando ele ira ser chamado

		Ele apenas sera chamado por ultimo
	*/
	defer req.Body.Close()

	//Pegando o body
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))

}
