v1 = int(input())
v2 = int(input())

if v1 > v2:
    v1, v2 = v2, v1

sum = sum([i for i in range(v1, v2 + 1) if i % 13 != 0])

print(sum)
