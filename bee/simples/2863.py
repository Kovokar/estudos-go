def menor():
    menor = 999999999999999999999999999999999999999999999

    voltas = int(input())

    for i in range(voltas):
        num = float(input())
        if num < menor:
            menor = num

    print(f"{menor:.2f}")

while True:
    try:
        menor()
    except EOFError:
        break
