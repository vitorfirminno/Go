package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura float32
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main(){
	fmt.Println("Herança")

	p1 :=  pessoa{"Vitor", "Firmino", 19, 1.81}
	fmt.Println(p1)

	e1 := estudante{p1, "Analise e desenvolvimento de sistemas", "uniBROXA"}
	fmt.Println(e1)
	
	fmt.Println(e1.altura)



}