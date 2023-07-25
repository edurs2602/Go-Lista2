package generics

import "fmt"

type Entry[C int | string, V any] struct {
	Key   C
	Value V
}

type Map[C int | string, V any] struct {
	Entries []Entry[C, V]
	Length  int
}

func (mapa *Map[C, V]) Add(key C, value V) {
	//Percorrendo mapa
	for i := 0; i < len(mapa.Entries); i++ {
		if mapa.Entries[i].Key == key {
			//Caso exista valor para a chave informada, ele é sobrescrito
			mapa.Entries[i].Value = value
			return
		}
	}
	//Criando nova entrada no mapa
	mapa.Entries = append(mapa.Entries, Entry[C, V]{key, value})
	mapa.Length += 1
}

func (mapa *Map[C, V]) Get(keyToSearch C) {
	for i := 0; i < len(mapa.Entries); i++ {
		if mapa.Entries[i].Key == keyToSearch {
			fmt.Println("Chave da entrada:", mapa.Entries[i].Key, "-- Valor da entrada:", mapa.Entries[i].Value)
			return
		}
	}
	fmt.Println("O mapa não possui nenhuma entrada com a chave informada!")
}

func (mapa *Map[C, V]) Size() {
	fmt.Println(mapa.Length)
}

func (mapa *Map[C, V]) Print() {
	fmt.Println("Imprimindo mapa:")
	for i := 0; i < len(mapa.Entries); i++ {
		fmt.Println("Chave da entrada:", mapa.Entries[i].Key, "\nValor da entrada:", mapa.Entries[i].Value)
	}
}
