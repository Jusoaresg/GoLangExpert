package main

const a = "Hello, World1"

// Escopo global da main
var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
)

func main() {
	//Escopo local e um Short Hand
	a := "X" //inferencia de uma string
	a = "fd"
	println(a)
}
