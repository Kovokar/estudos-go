def get() -> tuple:
    return map(float, input().split())

def main():
    vitoria1, vitoria2, empate = 0, 0, 0

    while True:
        v1, v2 = get()
        if v1 > v2:
            vitoria1 += 1
        elif v2 > v1:
            vitoria2 += 1
        else:
            empate += 1

        print("Novo grenal (1-sim 2-nao)")
        if int(input()) == 2:
            break

    print(f"{vitoria1 + vitoria2 + empate} grenais")
    print(f"Inter:{vitoria1}")
    print(f"Gremio:{vitoria2}")
    print(f"Empates:{empate}")

    if vitoria1 > vitoria2:
        print("Inter venceu mais")
    elif vitoria2 > vitoria1:
        print("Gremio venceu mais")
    else:
        print("Nao houve vencedor")

main()
