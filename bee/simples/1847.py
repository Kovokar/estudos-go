def get() -> int:
    a,b,c = map(int, input().split())
    return a,b,c


def cases(a, b, c) -> str:
    match (a, b, c):
        #desceu e subiu ou igual
        case(x,y,z) if x > y and (y <= z):
            print(":)")
        # subiu e desceu ou igual
        case(x,y,z) if x < y and (y >= z):
            print(":(")

        case(x,y,z) if x < y and y < z and (y-x > z-y):
            print(":(")

        case(x,y,z) if x < y and y < z and (y-x <= z-y):
            print(":)")

        case(x,y,z) if x > y and y > z and (x-y > y-z):
            print(":)")

        case(x,y,z) if x > y and y > z and (x-y <= y-z):
            print(":(")

        case(x,y,z) if x == y:
            if y < z:
                print(":)")
            else:
                print(":(")




a,b,c = get()
cases(a,b,c)