n = int(input())
carneiros = list(map(int, input().split()))

# Inicializa as variáveis
estrelas_atacadas = set()
i = 0

while 0 <= i < n:
    # Marca a estrela como atacada
    estrelas_atacadas.add(i)
    
    if carneiros[i] > 0:
        # Rouba um carneiro
        carneiros[i] -= 1
        # Decide para onde ir com base na paridade
        if carneiros[i] % 2 == 0:
            i -= 1  # Se par, volta para a estrela anterior
        else:
            i += 1  # Se ímpar, avança para a próxima estrela
    else:
        break

# Calcula o número total de carneiros não roubados
total_carneiros_restantes = sum(carneiros)

# Exibe o resultado
print(len(estrelas_atacadas), total_carneiros_restantes)
