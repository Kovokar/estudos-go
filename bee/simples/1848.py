

def get_bin():
    aux = 0
    while True:
        num = ""
        caw = input()
        if caw[0] == "c":
            print(aux)
            aux = 0 
            break
        for i in caw:
            if i == "-" or i == "": num += "0"
            if i == "*": num += "1"
        
        num = int(num, 2) #if len(num) != 0 else 0
        aux += num



    
for i in range(3):
    get_bin()