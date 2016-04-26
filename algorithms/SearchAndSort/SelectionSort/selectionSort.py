def selectionSort(sortList):
    for slot in range(len(sortList) - 1, 0, -1):
        positionOfMax = 0
        for location in range(1, slot + 1):
            if (sortList[location] > sortList[positionOfMax]):
                positionOfMax = location

        temp = sortList[slot]
        sortList[slot] = sortList[positionOfMax]
        sortList[positionOfMax] = temp

myList = [54,26,93,17,77,31,44,55,20]
selectionSort(myList)
print(myList)