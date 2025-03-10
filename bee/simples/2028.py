
def seq(num: int):
    cont = 1
    lista = [0]
    for i in range(num+1):
        for _ in range(i):
          cont += 1
          lista.append(i)

    print(f"Caso {i}: {len(lista)}", "numero" if len(lista) == 1 else "numeros")
    # for k in lista:
    #     print(k, end=" ")
    print("\n")


while True:
  try:
    o = int(input())
    seq(o)
  except EOFError:
    break
  