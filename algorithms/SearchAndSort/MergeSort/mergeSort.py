def mergeSort(sortList):
    if len(sortList) > 1:
        mid = len(sortList)//2
        lefthalf = sortList[:mid]
        righthalf = sortList[mid:]

        mergeSort(lefthalf)
        mergeSort(righthalf)

        i = 0
        j = 0
        k = 0
        print(i < len(lefthalf) and j < len(righthalf))
        while i < len(lefthalf) and j < len(righthalf):
            if lefthalf[i] < righthalf[j]:
                sortList[k] = lefthalf[i]
                i = i + 1
            else:
                sortList[k] = righthalf[j]
                j = j + 1
            k = k + 1

        while i < len(lefthalf):
            sortList[k] = lefthalf[i]
            i = i + 1
            k = k + 1

        while j < len(righthalf):
            sortList[k] = righthalf[j]
            j = j + 1
            k = k + 1

myList = [54,26,93,17,77,31,44,55,20]
mergeSort(myList)
print(myList)
