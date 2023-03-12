package main

import (
	"errors"
	"fmt"
)

func main() {

	// 4 tipos de números inteiros em go: int8, int32, int64 (nímero de bits suportados). Float segue o mesmo padrão
	var numero int8 = 100
	fmt.Println(numero)

	// Uint é int sem + ou -

	// Stringe booleano seguem o padrão das outras linguagens

	// Em go, error é um tipo definido inicialmente como nil (nulo)
	erro := errors.New("Erro interno")
	fmt.Println(erro)
}
