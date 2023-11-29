/*
Escreva um programa que leia informações de um arquivo e as represente em uma fatia de structs.
Suponha que haja um arquivo de texto que contenha uma série de nomes.
Cada linha do arquivo de texto tem um nome e um sobrenome, nessa ordem, separados por um único
espaço na linha.

Seu programa definirá uma estrutura de nomes com dois campos, fname para
o primeiro nome e lname para o último nome.

Cada campo será uma string de tamanho 20 (caracteres). Seu programa deve solicitar ao usuário o nome do arquivo de texto.
Seu programa lerá sucessivamente cada linha do arquivo de texto e criará um struct que contém
o primeiro e o último nome encontrados no arquivo. Cada struct criada será adicionada a uma
fatia e, depois que todas as linhas tiverem sido lidas do arquivo, seu programa terá uma fatia
contendo uma struct para cada linha do arquivo. Depois de ler todas as linhas do arquivo,
seu programa deve iterar pela fatia de structs e imprimir o primeiro e o último nome encontrados em cada struct.

Envie seu código-fonte para o programa "read.go".
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	FName string
	LName string
}

func main() {
	var filePath string
	var namesList []Name

	fmt.Print("File Path: ")
	_, err := fmt.Scanf("%s", &filePath)

	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fullName := fileScanner.Text()
		nameParts := strings.Split(fullName, " ")
		if len(nameParts) == 2 {

			firstName := nameParts[0]
			lastName := nameParts[1]

			if len(firstName) > 20 {
				firstName = firstName[:19]
			}
			if len(lastName) > 20 {
				lastName = lastName[:19]
			}

			name := Name{
				FName: firstName,
				LName: lastName,
			}
			namesList = append(namesList, name)
		}
	}

	for index, name := range namesList {
		fmt.Printf("%d - First Name: %s, Last Name: %s\n", index, name.FName, name.LName)
	}
}
