function sum(numList) {
    if (numList.length == 1)
        return numList[0]
    else {
        var num = numList[0]
        numList.splice(0, 1)
        return num + sum(numList)
    }
}

console.log(sum([1, 2, 3]))