function binarySearch(myList, item) {
    var first = 0
    var last = myList.length - 1
    var found = false
    while(first <= last && !found) {
        middle = Math.floor((first + last) / 2)
        if (myList[middle] == item) {
            found = true
        } else {
            if(item < myList[middle]) {
                last = middle - 1
            } else {
                first = middle + 1
            }
        }
    }
    return found
}

console.log(binarySearch([1, 2, 3, 4, 5, 6], 7))