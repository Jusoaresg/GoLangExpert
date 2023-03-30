package main

//Ponteiros e Structs

type Conta struct {
	saldo int
}

/*
Coisa muito comum (Apenas explicação)

Nessa função "NewConta"
Estou criando uma função com o endereço da memoria da conta
para retornar um novo endereço de uma nova conta
*/
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

//Quando tem o asterisco eu estou mexendo no local da memoria
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{
		saldo: 100,
	}
	conta.simular(200)
	println(conta.saldo)
}
