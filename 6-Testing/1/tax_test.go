package tax

import (
	"testing"
)

/*

go test .

FORMA VERBOSA
go test . -v
*/

// Faz o cover dos testes, e da uma porcentagem de quanto os testes cobrem o codigo
//go test -coverprofile=coverage.out

// Abre um html mostrando o lugar no codigo que os testes estão e não estão cobrindo o codigo
//go tool cover -html=coverage.out

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}
	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

// go test -bench=.
// OU
// go test -bench .
// Para rodar apenas o benchmark
// go test -bench=. -run=^#
// ^# é uma regular expression
// -count=10
// -benchtime=3s
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}
