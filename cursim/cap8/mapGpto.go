package main

import "fmt"

func main() {
	// Criando um map usando make
	idade := make(map[string]int)

	// Adicionando pares chave-valor ao map
	idade["Alice"] = 17
	idade["Bob"] = 30
	idade["Charlie"] = 22

	// Acessando um valor usando a chave
	fmt.Println("Idade de Alice:", idade["Alice"])
	fmt.Println("Idade de Bob:", idade["Bob"])

	// Verificando se uma chave existe no map
	valor, ok := idade["Charlie"]
	if ok {
		fmt.Println("Idade de Charlie:", valor)
	} else {
		fmt.Println("Charlie não encontrado no map")
	}

	// Atualizando um valor existente
	idade["Alice"] = 26
	fmt.Println("Nova idade de Alice:", idade["Alice"])

	// Removendo um item do map
	delete(idade, "Bob")
	fmt.Println("Mapa após remover Bob:", idade)

	// Iterando sobre um map
	fmt.Println("Iterando sobre o map:")
	for nome, idadePessoa := range idade {
		fmt.Printf("%s tem %d anos\n", nome, idadePessoa)
	}
}
