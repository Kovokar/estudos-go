def nota_valida(nota: float) -> bool:
    return 0 <= nota <= 10

def get_nota() -> float:
    while True:
        nota = float(input())
        if nota_valida(nota):
            return nota
        print("nota invalida")

def novo_calculo() -> bool:
    while True:
        print("novo calculo (1-sim 2-nao)")
        opcao = float(input())
        if opcao == 1:
            return True
        elif opcao == 2:
            return False

def main():
    while True:
        nota1 = get_nota()
        nota2 = get_nota()
        media = (nota1 + nota2) / 2
        print(f"media = {media:.2f}")
        
        if not novo_calculo():
            break

main()
