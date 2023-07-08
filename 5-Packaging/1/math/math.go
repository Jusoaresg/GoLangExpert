package math

// Quando a variavel começa com uma letra maiuscula ela fica PUBLICA
// Com a letra minuscula ela fica privada
type Math struct {
	A int
	B int
}

func (m Math) Add() int {
	return m.A + m.B
}
