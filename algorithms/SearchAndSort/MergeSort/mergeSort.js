function mergeSort(sortList) {
    if (sortList.length > 1) {
        console.log(sortList)
        var mid = Math.round(sortList.length / 2)
        var leftHalf = sortList.slice(0, mid)
        var rightHalf = sortList.slice(mid, sortList.length)

        leftHalf = mergeSort(leftHalf)
        rightHalf = mergeSort(rightHalf)

        var i = 0, j = 0, k = 0;
        
        while (i < leftHalf.length && j < rightHalf.length) {
            if (leftHalf[i] < rightHalf[j]) {
                sortList[k] = leftHalf[i]
                i++
            } else {
                sortList[k] = rightHalf[j]
                j++
            }
            k++
        }

        while (i < leftHalf.length) {
            sortList[k] = leftHalf[i]
            i++
            k++
        }

        while (j < rightHalf.length) {
            sortList[k] = rightHalf[j]
            j++
            k++
        }
    }
    return sortList
}

var myList = [54,26,93,17,77,31,44,55,20]
console.log(mergeSort(myList))