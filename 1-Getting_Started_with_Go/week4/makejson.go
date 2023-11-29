package main

import (
	"encoding/json"
	"fmt"
)

/*
Escreva um programa que solicite ao usuário que primeiro digite um nome e depois um endereço.
Seu programa deve criar um mapa e adicionar o nome e o endereço ao mapa usando as chaves "name" e "address",
respectivamente. Seu programa deve usar Marshal() para criar um objeto JSON a partir do mapa e, em seguida,
deve imprimir o objeto JSON.

Envie seu código-fonte para o programa,
"makejson.go".
*/

type Person struct {
	Name    string
	Address string
}

func main2() {
	var name string
	var address string

	fmt.Print("Write a name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Print("error in name")
		panic(err)
	}

	fmt.Print("Write an address: ")
	_, err = fmt.Scan(&address)
	if err != nil {
		fmt.Println("error in address")
		panic(err)
	}

	home := Person{
		Name:    name,
		Address: address,
	}
	fmt.Println(home)
	result, err := json.Marshal(&home)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))

}
