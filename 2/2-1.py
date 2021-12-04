with open("input.txt", "r") as f:
    lines = f.read().splitlines()

horiz = 0
depth = 0
for line in lines:
    var, num = line.split()
    if var == "forward":
        horiz = horiz + int(num)
    elif var == "down":
        depth = depth + int(num)
    elif var == "up":
        depth = depth - int(num)

print(horiz, depth)
print(horiz*depth)
