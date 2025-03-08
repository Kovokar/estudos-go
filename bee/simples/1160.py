def get():
    v1, v2, v3, v4 = map(float, input().split())
    
    for i in range(101): 
        if v1 > v2:  
            return f"{i} anos."
        
        v1 = int(v1 * (1 + (v3 / 100))) if v3 != 0 else v1
        v2 = int(v2 * (1 + (v4 / 100))) if v4 != 0 else v2
    
    return "Mais de 1 seculo."

def main():
    voltas = int(input())  
    for i in range(voltas):
        print(get())

main()
