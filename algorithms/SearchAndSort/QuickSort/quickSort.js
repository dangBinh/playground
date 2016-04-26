function quickSort(sortList) {
    quickSortHelper(sortList, 0, sortList.length - 1)
    return sortList
}

function quickSortHelper(sortList, first, last) {
    if (first < last) {
        var splitPoint = partition(sortList, first, last)
        
        quickSortHelper(sortList, first, splitPoint - 1)
        quickSortHelper(sortList, splitPoint + 1, last)
    }
}

function partition(sortList, first, last) {
    var pivotValue = sortList[first]
    var leftMark = first + 1
    var rightMark = last

    var found = false
    while (!found) {
        while (leftMark <= rightMark && sortList[leftMark] <= pivotValue) {
            leftMark++
        }

        while (sortList[rightMark] >= pivotValue && rightMark >= leftMark) {
            rightMark--
        }
        if (rightMark < leftMark) {
            found = true
        } else {
            temp = sortList[leftMark]
            sortList[leftMark] = sortList[rightMark]
            sortList[rightMark] = temp
        }
    }

    temp = sortList[first]
    sortList[first] = sortList[rightMark]
    sortList[rightMark] = temp

    return rightMark
}

var myList = [54,26,93,17,77,31,44,55,20]
console.log(quickSort(myList))