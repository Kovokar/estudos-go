while True:
    try:
        a, b = input().split()  
        if (int(a) == int(b) or a == b) and a != "0" and b != 0:
            print(0)
            pass
        elif b == a * (len(b) // len(a))  and a != "0" and b != 0:
            print(0)
        else:
            try:
                b = int(b.replace(a, ""))
                print(b)
            except ValueError:
                pass  # Se não for possível converter o valor de b para int, ignora
    except EOFError:
        break  # Sai do loop quando ocorrer EOF (fim da entrada)
