# Métodos em Go

Em Go, métodos são funções associadas a tipos específicos. Eles são definidos usando um **receiver** (receptor), que é um parâmetro especial que permite que a função seja chamada como um método de um determinado tipo.

## Definição de Métodos

A sintaxe básica de um método em Go é:

```go
package main

import "fmt"

// Definição de um tipo struct
type Pessoa struct {
    Nome string
    Idade int
}

// Método associado à struct Pessoa
func (p Pessoa) Saudacao() {
    fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
    pessoa := Pessoa{"Carlos", 30}
    pessoa.Saudacao() // Chamando o método
}
```

## Explicação
- `Pessoa` é uma struct com dois campos: `Nome` e `Idade`.
- `Saudacao()` é um método associado ao tipo `Pessoa`.
- O método é chamado usando `pessoa.Saudacao()`.

## Métodos com Ponteiros
Para modificar os valores do receptor dentro do método, é necessário usar um ponteiro:

```go
func (p *Pessoa) Envelhecer() {
    p.Idade++
}

func main() {
    pessoa := Pessoa{"Carlos", 30}
    pessoa.Envelhecer()
    fmt.Println(pessoa.Idade) // 31
}
```

Aqui, `p *Pessoa` permite que o método modifique o valor original da struct.

## Conclusão
Os métodos em Go são úteis para organizar código e associar comportamento a tipos específicos. O uso de ponteiros permite modificar os valores dentro dos métodos, tornando-os mais poderosos e flexíveis.
