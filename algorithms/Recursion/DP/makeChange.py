def makeChnage(coinValueList, change, minCoins, coinUsed):
    for cents in range (change + 1):
        coinCount = cents
        newCoin = 1
        for j in [c for c in coinValueList if c <= cents]:
            if minCoins[cents - j] + 1 < coinCount:
                coinCount = minCoins[cents - j] + 1
                newCoin = j
        minCoins[cents] = coinCount
        coinUsed[cents] = newCoin
    return minCoins[change]

def printCoin(coinUsed, change):
    coin = change
    while coin > 0:
        thisCoin = coinUsed[coin]
        print(thisCoin)
        coin = coin - thisCoin


change = 63
cointList = [1, 5, 10, 21, 25]
coinUsed = [0]*64
coinCount = [0]*64
print(makeChnage(cointList, change, coinCount, coinUsed))
print(printCoin(coinUsed, change))
print(coinUsed)
