function shellSort(sortList) {
    var countSublist = Math.floor(sortList.length / 2)
    while(countSublist > 0) {
        for (var i = 0; i < countSublist; i++) {
            sortList = gapInsertionSort(sortList, i, countSublist)
            countSublist = Math.floor(countSublist / 2)
        }
    }
    return sortList
}

function gapInsertionSort(sortList, startPosition, gap) {
    for(var i = startPosition + gap; i < sortList.length; i = i + gap) {
        var currentValue = sortList[i]
        var position = i

        while(position >= gap && sortList[position - gap] > currentValue) {
            sortList[position] = sortList[position - gap]
            position = position - gap
        }
        sortList[position] = currentValue 
    }
    return sortList
}

var myList = [54,26,93,17,77,31,44,55,20]
console.log(shellSort(myList))