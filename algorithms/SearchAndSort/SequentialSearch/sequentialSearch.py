def sequentialSearch(myList, item):
    pos = 0
    found = False

    while pos < len(myList) and not found:
        if myList[pos] == item:
            found = True
        else:
            pos = pos + 1

    return found

testList = [1, 2, 3, 4, 6, 7, 10]
print(sequentialSearch(testList, 3))
print(sequentialSearch(testList, 12))