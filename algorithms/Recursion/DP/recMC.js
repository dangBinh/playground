function recMC(coinValueList, change) {
    var minCoins = change
    if (coinValueList.indexOf(change) != -1) {
        return 1
    } else {
        var coinValueList = coinValueList.filter(function(coinValue) {
                return coinValue <= change
            })
        coinValueList.forEach(function(coinValue) {
            var numCoins = 1 + recMC(coinValueList, change - coinValue)
            if (numCoins < minCoins) {
                minCoins = numCoins
            }
        })
        return minCoins
    }
}

console.log(recMC([1, 5, 10, 25], 26))