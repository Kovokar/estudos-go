import math

voltas = 9999
menor = consumoTotal = pessoasTotais = 0
num1 = num2 = cont = 0
slilce = []

def mostrar(arr):
    global cont
    cont += 1
    print(f"Cidade# {cont}:")
    for j, sub in enumerate(arr):
        for i in range(len(sub)-1):
            if j == len(arr)-1:
                print(f"{sub[i]}-{sub[i+1]}", end="")
            else:
                print(f"{sub[i]}-{sub[i+1]}", end=" ")
    consMedio = consumoTotal / pessoasTotais
    truncated = math.floor(consMedio * 100) / 100
    print(f"\nConsumo medio: {truncated:.2f} m3.\n")

def recebeValores(voltas):
    global consumoTotal, pessoasTotais, slilce
    consumoTotal = 0
    pessoasTotais = 0
    for _ in range(voltas):
        num1, num2 = map(int, input().split())
        consumoTotal += num2
        pessoasTotais += num1
        slilce.append([num1, num2 // num1])

def bubbleSort(arr):
    n = len(arr)
    for i in range(n-1):
        for j in range(n-i-1):
            if arr[j][1] > arr[j+1][1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
    return arr

def isEqual(arr):
    n = len(arr)
    i = 0
    while i < n - 1:
        j = i + 1
        while j < n:
            if arr[i][1] == arr[j][1]:
                arr[i][0] += arr[j][0]
                arr.pop(j)
                n = len(arr)
            else:
                j += 1
        i += 1
    return arr

def esvaziaSlice(arr):
    return arr[len(arr):]

def exec():
    global voltas
    voltas = int(input())
    while voltas != 0:
        recebeValores(voltas)
        slilce = bubbleSort(slilce)
        slilce = isEqual(slilce)
        slilce = isEqual(slilce)
        slilce = isEqual(slilce)
        slilce = isEqual(slilce)
        mostrar(slilce)
        slilce = esvaziaSlice(slilce)
        voltas = int(input())
       

exec()
