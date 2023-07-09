package math

// Quando a variavel começa com uma letra maiuscula ela fica PUBLICA
// Com a letra minuscula ela fica privada
type math struct {
	a int
	b int
}

func NewMath(a int, b int) math {
	return math{
		a: a,
		b: b,
	}
}

func (m math) Add() int {
	return m.a + m.b
}
