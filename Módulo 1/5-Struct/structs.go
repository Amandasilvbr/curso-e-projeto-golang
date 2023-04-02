package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

type endereco struct{
	usuario
	rua string
	numero int
	bairro string
}

func main() {
	pessoa := usuario{"Pessoa", 12}
	fmt.Println(pessoa.nome)
	fmt.Println(pessoa.nome)
	fmt.Println(pessoa.idade)

}