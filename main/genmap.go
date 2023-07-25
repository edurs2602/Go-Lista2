package main

import (
	"fmt"
	"generics"
	"os"
)

func main() {
	//Instanciando mapa. (Tipos de chave e valor ser definidos na declaração?)
	meuMapa := generics.Map[int, string]{Entries: []generics.Entry[int, string]{}, Length: 0}

	var function, value string
	var key int

	for {
		fmt.Println("Insira o comando e os parâmetros(Se necessário):")
		fmt.Scanln(&function, &key, &value)

		switch function {

		case "add":
			meuMapa.Add(key, value)
			fmt.Println("Valor adicionado!")

		case "get":
			meuMapa.Get(key)

		case "size":
			meuMapa.Size()

		case "print":
			meuMapa.Print()

		case "exit":
			fmt.Println("Encerrando aplicação!")
			os.Exit(0)

		default:
			fmt.Println("Por favor, insira um comando/valor válido.")
		}
	}
}
