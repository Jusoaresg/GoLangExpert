package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	//Retorno no command line
	err := t.Execute(os.Stdout, Cursos{
		{"GO", 40},
		{"Java", 80},
		{"Python", 20},
		{"C#", 50},
	})
	if err != nil {
		panic(err)
	}
}
