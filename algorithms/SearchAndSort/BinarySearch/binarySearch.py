def binarySearch(myList, item):
    first = 0
    last = len(myList) - 1
    found = False
    while not found && first <= last:
        middle = (first + last) // 2
        if myList[middle] == item:
            found = True
        else:
            if item < myList[middle]:
                last = middle - 1
            else:
                first = middle + 1
    return found

print(binarySearch([1, 2, 3, 4, 5, 6], 6))

# Divide and conquire

