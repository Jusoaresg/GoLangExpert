package main

import "github.com/Jusoaresg/GoLangExpert/5-Packaging/3/math"

/**
Comando usado para o replace
go mod edit -replace github.com/Jusoaresg/GoLangExpert/5-Packaging/3/math=../math

Funciona bem localmente mas quando for dar uploud vai dar problema
*/
func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}
