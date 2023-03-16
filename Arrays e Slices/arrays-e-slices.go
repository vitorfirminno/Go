package main

import (
	"fmt"
	"reflect"


)

func main(){
	fmt.Println("Arrays e Slices")

	var array1[5] string
	fmt.Println(array1)

	array2 := [5]string{"Posição 1" , "Posição 2","Posicão 3" ,"Posicão 4", "Posicão 5"}
	fmt.Println(array2)

	var qtd int

	qtd = 79

	array3 := [...]int{qtd, 123, 45,54,87,87}
	fmt.Println(array3)

	//-------------------------------------------------------------------------------------

	slice := []int{10, 11, 12, 165, 1965, 165, 65}
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18) // adicionar 
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	fmt.Println(array2)

	//ARRAYS INTERNOS
	fmt.Println("-----------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade
	fmt.Println(slice3)

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacidade


}
