package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("teste")
	auxiliar.Auxiliar()	


	erro := checkmail.ValidateFormat("vitorfds18@gmail.com")
	fmt.Println(erro)
}