package main

import "fmt"

type pessoa struct{
	nome string
	idade int
	altura int
	sobrenome string
}

type estudante struct{
	pessoa
	curso string
	ano int
}

func main(){
	pessoa := pessoa{"Pessoa x", 21, 174, "Sobrenome x"}
	estudante := estudante{pessoa, "Curso x", 2}
	fmt.Println(estudante.nome)
}