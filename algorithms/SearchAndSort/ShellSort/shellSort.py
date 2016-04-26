def shellSort(sortList):
    sublistcount = len(sortList) // 2
    while sublistcount > 0:
        for startposition in range(sublistcount):
            gapInsertionSort(sortList, startposition, sublistcount)

        sublistcount = sublistcount // 2

def gapInsertionSort(sortList, start, gap):
    for i in range(start + gap, len(sortList), gap):
        currentValue = sortList[i]
        position = i

        while position >= gap and sortList[position - gap] > currentValue:
            sortList[position] = sortList[position - gap]
            position = position - gap

        sortList[position] = currentValue

myList = [54,26,93,17,77,31,44,55,20]
shellSort(myList)
print(myList)