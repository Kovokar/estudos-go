def conta_carry(a, b):
    carry = 0
    count = 0
    
    while a > 0 or b > 0:
        soma = (a % 10) + (b % 10) + carry
        if soma >= 10:
            carry = 1
            count += 1
        else:
            carry = 0
        a //= 10
        b //= 10
    
    return count

while True:
    a, b = map(int, input().split())
    if a == 0 and b == 0:
        break
    
    carry_count = conta_carry(a, b)
    
    if carry_count == 0:
        print("No carry operation.")
    elif carry_count == 1:
        print("1 carry operation.")
    else:
        print(f"{carry_count} carry operations.")
