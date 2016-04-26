function BubbleSort(sortList) {
    for(var i = sortList.length - 1; i >= 0; i--) {
        for (var j = 0; j < i; j++) {
            if(sortList[j] > sortList[j+1]) {
                temp = sortList[j]
                sortList[j] = sortList[j+1]
                sortList[j+1] = temp
            }
        }
    }
    return sortList
}

var myList = [1, 10, 2, 5, 9]
console.log(BubbleSort(myList))

function ShortBubbleSort(sortList) {
    var exchange = true
    var sortListNum = sortList.length - 1
    while (sortListNum > 0 && exchange) {
        exchange = false
        for (var i = 0; i < sortListNum; i++) {
            if(sortList[i] > sortList[i+1]) {
                exchange = true
                temp = sortList[i]
                sortList[i] = sortList[i+1]
                sortList[i+1] = temp
            }
        }
        sortListNum--
    }
    return sortList
}

var myList = [1, 10, 9, 2]
console.log(ShortBubbleSort(myList))