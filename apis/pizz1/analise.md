# Estrutura DDD em Go

Este repositório demonstra como implementar a Arquitetura DDD (Domain-Driven Design) em Go, utilizando os princípios de separação de responsabilidades e clareza de domínio. A arquitetura proposta inclui os seguintes componentes principais: **Model**, **Repository**, **Controller** e **Usecase**.

## Visão Geral

A abordagem DDD visa organizar o código de forma que as regras de negócio e o domínio sejam o centro do desenvolvimento, promovendo uma melhor separação de responsabilidades e manutenibilidade. No Go, podemos aplicar esses conceitos criando camadas distintas para diferentes responsabilidades no sistema.

## Estrutura de Diretórios

Aqui está a estrutura recomendada para um projeto Go utilizando DDD:


### Explicação dos Componentes

1. **Model (Domínio)**
   - Contém as entidades e objetos de valor que representam os conceitos principais do seu domínio. A camada de modelo é onde as regras de negócio e comportamentos essenciais são definidos.
   - Exemplo: `User`, `Order`, `Product`.

2. **Repository**
   - Esta camada é responsável por acessar os dados persistidos (banco de dados, por exemplo). O repositório implementa uma interface definida no domínio, garantindo que a lógica de acesso a dados esteja isolada.
   - Exemplo: `UserRepository`, `OrderRepository`.

3. **Controller**
   - Esta camada recebe e responde a requisições de entrada (geralmente HTTP, mas pode ser qualquer outra interface). Os controladores se comunicam com os casos de uso para processar as solicitações.
   - Exemplo: `UserController`, `OrderController`.

4. **Usecase**
   - A camada de casos de uso representa as operações que o sistema executa para cumprir os requisitos de negócio. Os casos de uso orquestram as interações entre as entidades do domínio e os repositórios.
   - Exemplo: `CreateUser`, `PlaceOrder`.

## Detalhamento de Cada Camada

### 1. Model (Entidades e Objetos de Valor)

As entidades são os objetos principais do seu domínio, com identidade própria. Os objetos de valor (Value Objects) são usados para representar atributos que não têm identidade própria, como um endereço ou uma data.

Exemplo de entidade:

```go
package domain

type User struct {
    ID    int
    Name  string
    Email string
}
