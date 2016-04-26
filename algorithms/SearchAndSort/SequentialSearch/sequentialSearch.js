function sequentialSearch(list, item) {
    var found = false
    var pos = 0
    while (pos < list.length && !found) {
        if (list[pos] == item) {
            found = true
        } else {
            pos++
        }
    }
    return found
}

console.log(sequentialSearch([1, 2, 3, 4], 4))