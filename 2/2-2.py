with open("input.txt", "r") as f:
    lines = f.read().splitlines()

horiz = 0
depth = 0
aim = 0
for line in lines:
    var, num = line.split()
    if var == "forward":
        horiz = horiz + int(num)
        depth = depth + aim * int(num)
    elif var == "down":
        aim = aim + int(num)
    elif var == "up":
        aim = aim - int(num)

print(horiz*depth)
