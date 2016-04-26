function recDC(coinValueList, change, knownResult) {
    var minCoin = change
    if (coinValueList.indexOf(change) != -1) {
        return 1;
    } else if(knownResult[change] > 0) {
        return knownResult[change]
    } else {
        coinValueList.filter(function(value, index) {
            if(value <= change) {
                return value
            }    
        })
        .forEach(function(value, index) {
            numConis = 1 + recDC(coinValueList, change - value, knownResult)
            if (numConis < minCoin) {
                minCoin = numConis
                knownResult[change] = numConis
            }
        })
    }
    return minCoin
}

console.log(recDC([1, 5, 10, 25], 63, new Array()))