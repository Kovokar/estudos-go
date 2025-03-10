def seq(num: int, cont: int):
    lista = [0]
    for i in range(num + 1):
        for _ in range(i):
            lista.append(i)

    print(f"Caso {cont}: {len(lista)}", "numero" if len(lista) == 1 else "numeros")
    for i in range(len(lista)):
        if i == len(lista) - 1:  # Verifica se é o último item
            print(lista[i])
        else:
            print(lista[i], end=" ")



cont = 1
while True:
    try:
        o = int(input())
        seq(o, cont)
        print()
        cont += 1
    except EOFError:
        break
