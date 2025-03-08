def get_vals(v1: str, v2: str) -> str:
    if v1 == v2:
        return "De novo!"
    elif \
    (v1 == "tesoura" and v2 in ["papel", "lagarto"]) or\
    (v1 == "papel" and v2 in ["pedra", "Spock"]) or\
    (v1 == "pedra" and v2 in ["lagarto", "tesoura"]) or\
    (v1 == "lagarto" and v2 in ["Spock", "papel"]) or\
    (v1 == "Spock" and v2 in ["tesoura", "pedra"]):
        return "Bazinga!"
    else:
        return "Raj trapaceou!"
    
def main():
    o = int(input())
    for i in range(o):
        v1, v2 = map(str, input().split())
        print(f"Caso #{i+1}: {get_vals(v1, v2)}")


main()