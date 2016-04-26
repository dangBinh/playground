function makeChange(coinValueList, change, minCoins, coinUsed) {
    for(var cents = 0; cents <= change; cents++) {
        var coinCount = cents
        var newCoin = 1
        coinValueList.filter(function(value, index) {
            return value <= cents
        }).forEach(function(value, index) {
            if(minCoins[cents - value] + 1 < coinCount) {
                coinCount = minCoins[cents - value] + 1
                newCoin = value
            }
        })
        minCoins[cents] = coinCount
        coinUsed[cents] = newCoin
    }
    return minCoins[change]
}

function printCoin(coinUsed, change) {
    coin = change
    while (coin > 0) {
        var thisCoin = coinUsed[coin]
        coin = coin - thisCoin
    }
}

var minCoins = new Array()
var coinUsed = new Array()
var coinValueList = [1, 5, 10, 21, 25]
var change = 63

console.log(makeChange(coinValueList, change, minCoins, coinUsed))
