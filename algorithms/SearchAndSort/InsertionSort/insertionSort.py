def insertionSort(sortList):
    for index in range(1, len(sortList)):
        currentValue = sortList[index]
        position = index

        while position > 0 and sortList[position - 1] > currentValue:
            sortList[position] = sortList[position - 1]
            position = position - 1

        sortList[position] = currentValue

myList = [54,26,93,17,77,31,44,55,20]
insertionSort(myList)
print(myList)