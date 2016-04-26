def BubbleSort(sortList):
    for nums in range(len(sortList) - 1, 0, -1):
        for i in range(nums):
            if sortList[i] > sortList[i+1]:
                temp = sortList[i]
                sortList[i] = sortList[i+1]
                sortList[i+1] = temp
    return sortList

myList = [1, 5, 3, 9, 10, 20, 11]
BubbleSort(myList)
print(myList)

def ShortBubbleSort(sortList):
    exchange = True
    nums = len(sortList) - 1
    while nums > 0 and exchange:
        exchange = False
        for i in range(nums):
            if sortList[i] > sortList[i+1]:
                exchange = True
                temp = sortList[i]
                sortList[i] = sortList[i+1]
                sortList[i+1] = temp
        nums = nums - 1

myList = [1, 3, 5, 10]
ShortBubbleSort(myList)
print(myList)