package main

import (
	"encoding/json"
	"os"
)

/*
TAGS para marcar que
caso tenha uma n no JSON é um Numero
e caso tem um s no JSON é um Saldo
*/
type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}

	//Marshal é transformar em JSON, serializar em JSON
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}

	//O JSON sempre retorna em BYTES
	println(string(res))

	/*
		Caso eu não queira criar uma variavel

		Eu posso criar um encoder para gravar
		em um arquivo ou no console
	*/

	// "os.Stdout" é a saida padrão, o propio console
	//Lembrando que ele retorna um error
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	/*
		Tambem pode fazer assim, caso queira economizar uma linha

		encoder := json.NewEncoder(os.Stdout).encoder.Encode(conta)

		ou

		json.NewEncoder(os.Stdout).encoder.Encode(conta)
	*/

	//Decodificar agora (Usando as TAGS para marcar qual é o numero e saldo)
	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	//Ele pega a referencia na memoria, para mudar mesmo a contaX
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)

}
