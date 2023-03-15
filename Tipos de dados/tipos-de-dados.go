package main

import (
	"errors"
	"fmt"
)

func main() {

	var inteiro int16 = 57
	fmt.Println(inteiro)

	var inteiro2 int = 409403690
	//quando é um tipo vazio é baseado nos bits da arquitetura do  PC
	fmt.Println(inteiro2)

	inteiro3 := 59548493843
	fmt.Println(inteiro3)

	 var inteiro0 int
	 fmt.Println(inteiro0)

	 var inteiro6 uint32
	 fmt.Println(inteiro6)

	//fim dos inteiros 


	var numeroReal float32 = 3.14
	fmt.Println(numeroReal)

	var numeroReal2 float32
	fmt.Println(numeroReal2)





	//tipos de numeros reais são obrigados a conter o numero de bits

	//fim dos numeros reais
	
	var boleano bool = true
	fmt.Println(boleano)

	var boleano2 bool
	fmt.Println(boleano2)

	boleano3 := true
	fmt.Println(boleano3)

	//fim  do boleano


	var texto string = "teste"
	fmt.Println(texto)

	char := 'B'
	//o char é declarado com aspas simples e ele retorna um numero da tabela ASCII
	fmt.Println(char)

	//Fim Strings

	var erro error 
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}