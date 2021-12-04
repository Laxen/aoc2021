with open("input.txt", "r") as f:
    lines = f.read().splitlines()

nums = []
for line in lines:
    l = list(line)
    nums.append([int(i) for i in l])

gamma = 0
eps = 0
for i in range(len(nums[0])):
    s = sum([num[i] for num in nums])
    if s > len(lines) / 2:
        gamma = gamma + 2**(len(nums[0])-i-1)
    else:
        eps = eps + 2**(len(nums[0])-i-1)

print(gamma*eps)
