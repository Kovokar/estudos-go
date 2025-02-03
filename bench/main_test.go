package main

import (
	"sort"
	"testing"
)

// Função que queremos benchmarkar
func benchmarkCidade(b *testing.B, N int, dados [][]int) {
	for i := 0; i < b.N; i++ {
		// Inicializa as variáveis
		var cidadeNumero int
		var totalPessoas, consumoTotal int
		var consumoPorPessoa map[int]int

		cidadeNumero = 1
		totalPessoas = 0
		consumoTotal = 0
		consumoPorPessoa = make(map[int]int) // Inicializa o mapa

		// Processa os dados fornecidos
		for _, d := range dados {
			X, Y := d[0], d[1]
			consumoPorPessoaPorPessoa := Y / X
			consumoPorPessoa[consumoPorPessoaPorPessoa] += X // Soma o número de habitantes
			totalPessoas += X
			consumoTotal += Y
		}

		// Ordena as chaves (consumo por pessoa)
		keys := []int{}
		for k := range consumoPorPessoa {
			keys = append(keys, k)
		}
		// Ordenação
		sort.Ints(keys)

		// Cálculo do consumo médio
		consumoMedio := float64(consumoTotal) / float64(totalPessoas)
		_ = consumoMedio // Não precisamos usar o resultado aqui para o benchmark
	}
}

// Função de benchmark principal
func BenchmarkCidade(b *testing.B) {
	// Dados de entrada para o benchmark (exemplo de caso de teste)
	dados := [][]int{
		{3, 22},
		{2, 11},
		{3, 39},
	}

	// Realiza o benchmark para o caso com 3 imóveis
	benchmarkCidade(b, 3, dados)
}
