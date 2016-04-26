def recDC(coinValueList, change, knownResult):
    minCoins = change
    if change in coinValueList:
        return 1
    elif knownResult[change] > 0:
        return knownResult[change]
    else:
        for i in [c for c in coinValueList if c <= change]:
            numCoins = 1 + recDC(coinValueList, change - i, knownResult)

            if numCoins < minCoins:
                minCoins = numCoins
                knownResult[change] = minCoins
    return minCoins

print(recDC([1, 5, 10, 25], 63, [0]*64))
