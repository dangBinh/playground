def quickSort(sortList):
    quickSortHelper(sortList, 0, len(sortList) - 1)

def quickSortHelper(sortList, first, last):
    if first < last:
        splitpoint = partition(sortList, first, last)

        quickSortHelper(sortList, first, splitpoint - 1)
        quickSortHelper(sortList, splitpoint + 1, last)

def partition(sortList, first, last):
    pivotvalue = sortList[first]

    leftmark = first + 1
    rightmark = last

    done = False
    while not done:
        while leftmark <= rightmark and sortList[leftmark] <= pivotvalue:
            leftmark = leftmark + 1

        while sortList[rightmark] >= pivotvalue and rightmark >= leftmark:
            rightmark = rightmark - 1

        if rightmark < leftmark:
            done = True
        else:
            temp = sortList[leftmark]
            sortList[leftmark] = sortList[rightmark]
            sortList[rightmark] = temp

    temp = sortList[first]
    sortList[first] = sortList[rightmark]
    sortList[rightmark] = temp

    return rightmark

myList = [54,26,93,17,77,31,44,55,20]
quickSort(myList)
print(myList)