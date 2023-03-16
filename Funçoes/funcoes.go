package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(nu1, nu2 int8) (int8, int8){
	soma := nu1 + nu2
	subtração := nu1 - nu2
	return soma ,subtração
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("função F")
		return txt
	}


	resultado :=  f("texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, _ := calculosMatematicos(10, 15)

	_, resultadoSubtracao := calculosMatematicos(10, 15)

	var1 , var2 := calculosMatematicos(10, 15)

	fmt.Println(var1, var2)

	fmt.Println(resultadoSoma)

	fmt.Println(resultadoSubtracao)
	
}