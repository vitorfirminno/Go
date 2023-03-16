package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct{
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Vitor"
	u.idade = 19

	fmt.Println(u)

	enderecoUsuario := endereco{"rua do sexo", 18}

	usuario2 := usuario{"Thomas Turbando", 21, enderecoUsuario}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "sexo"}
	fmt.Println(usuario3)

}