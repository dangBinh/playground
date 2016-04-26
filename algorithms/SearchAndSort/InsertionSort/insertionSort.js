function insertionSort(sortList) {
    for (var i = 1; i < sortList.length; i++) {
        var currentValue = sortList[i]
        var position = i
        while (position > 0 && sortList[position - 1] > currentValue) {
            sortList[position] = sortList[position - 1]
            position--
        }
        sortList[position] = currentValue
    }
    return sortList
}

var myList = [54,26,93,17,77,31,44,55,20]
console.log(insertionSort(myList))