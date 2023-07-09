package main

import (
	"github.com/Jusoaresg/GoLangExpert/5-Packaging/3/math"
	"github.com/google/uuid"
)

/*
Inicando o go workspace
go work init "modulos ou bibliotecas que quero usar"

Para ignorar os pacotes que não foram achados
go mod tidy -e
*/
func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
