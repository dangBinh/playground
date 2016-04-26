function selectionSort(sortList) {
    for (var i = sortList.length - 1; i >=0; i--) {
        var temp;
        var positionOfMax = 0
        for (var j = 1; j < i + 1; j++) {
            if (sortList[j] > sortList[positionOfMax]) {
                positionOfMax = j
            }
        }
        temp = sortList[i]
        sortList[i] = sortList[positionOfMax]
        sortList[positionOfMax] = temp
    }
    return sortList
}

var myList = [54,26,93,17,77,31,44,55,20]
console.log(selectionSort(myList))