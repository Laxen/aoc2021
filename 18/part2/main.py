import ast
import copy

def parseInput(filename):
    pairs = []

    with open(filename, "r") as f:
        for line in f:
            pair = ast.literal_eval(line)
            pairs.append(pair)

    return pairs

def addPair(p1, p2):
    return [p1, p2]

def addLeft(p, val):
    if isinstance(p[0], list):
        addLeft(p[0], val)
    else:
        p[0] += val

def addRight(p, val):
    if isinstance(p[1], list):
        addRight(p[1], val)
    else:
        p[1] += val

def reduce(p, depth):
    if depth >= 4:
        return 0, p[0], p[1], True
    
    if isinstance(p[0], list):
        newp, a, b, wasReduced = reduce(p[0], depth+1)
        p[0] = newp
        if wasReduced:
            if b > 0:
                if isinstance(p[1], list):
                    addLeft(p[1], b)
                else:
                    p[1] += b
            return p, a, 0, True
    if isinstance(p[1], list):
        newp, a, b, wasReduced = reduce(p[1], depth+1)
        p[1] = newp
        if wasReduced:
            if a > 0:
                if isinstance(p[0], list):
                    addRight(p[0], a)
                else:
                    p[0] += a
            return p, 0, b, True

    return p, -1, -1, False

def split(p):
    if isinstance(p, int):
        if p >= 10:
            a = int(p/2)
            b = int(p/2+0.5)
            return [a,b], True
        return p, False

    newp, wasSplit = split(p[0])
    p[0] = newp
    if wasSplit:
        return p, True

    newp, wasSplit = split(p[1])
    p[1] = newp
    if wasSplit:
        return p, True
    
    return p, False

def magnitude(p):
    if isinstance(p, int):
        return p
    
    a = magnitude(p[0])
    b = magnitude(p[1])

    return 3*a + 2*b

pairs = parseInput("input.txt")
maxMag = 0
for i, origPair in enumerate(pairs):
    for origAdd in pairs[0:i]:
        pair = copy.deepcopy(origPair)
        add = copy.deepcopy(origAdd)

        pair = addPair(pair, add)

        wasReduced = True
        while True:
            while wasReduced:
                pair, _, _, wasReduced = reduce(pair, 0)

            pair, wasReduced = split(pair)

            if wasReduced:
                continue
            break
        maxMag = max(maxMag, magnitude(pair))

    for origAdd in pairs[i+1:len(pairs)]:
        pair = copy.deepcopy(origPair)
        add = copy.deepcopy(origAdd)

        pair = addPair(pair, add)

        wasReduced = True
        while True:
            while wasReduced:
                pair, _, _, wasReduced = reduce(pair, 0)

            pair, wasReduced = split(pair)

            if wasReduced:
                continue
            break
        maxMag = max(maxMag, magnitude(pair))

print(maxMag)