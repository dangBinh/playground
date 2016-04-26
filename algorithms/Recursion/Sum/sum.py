def sum(numList):
    if len(numList) == 1:
        return numList[0]
    else:
        return numList[0] + sum(numList[1:])

print(sum([1, 2, 3]))