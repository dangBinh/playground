class Node:
    def __init__(self, initdata):
        self.data = initdata
        self.next = None

    def getData(self):
        return self.data

    def getNext(self):
        return self.next

    def setData(self, newData):
        self.data = newData

    def setNext(self, newnext):
        self.next = newnext

class UnorderedList:
    def __init__(self):
        self.head = None

    def isEmpty(self):
        return self.head == None

    def add(self, item):
        temp = Node(item)
        temp.setNext(self.head)
        self.head = temp

    def size(self):
        count = 0
        current = self.head
        while current != None:
            count = count + 1
            current = current.getNext()

        return count

    def search(self, item):
        found = False
        current = self.head

        while current != None and not found:
            if current.getData() == item:
                found = True
            else:
                current = current.getNext()

        return found

    def remove(self, item):
        current = self.head
        previous = None
        found = False

        while not found:
            if current.getData() == item:
                found = True
            else:
                previous = current
                current = current.getNext()

        if previous == None:
            self.head = current.getNext()
        else:
            previous.setNext(current.getNext())

    def append(self, item):
        current = self.head
        previous = None

        while current != None:
            previous = current
            current = current.getNext()

        if previous != None:
            temp = Node(item)
            previous.setNext(temp)
        else:
            self.add(item)

    def index(self, item):
        current = self.head
        previous = None
        count = 0
        found = False

        while current != None and not found:
            if current.getData() == item:
                found = True
            count = count + 1

        return count


    def insert(self, pos, item):
        current = self.head
        found = False
        value = None
        k = 0

        while not found:
            if k == pos:
                found = True
            else:
                previous = current
                current = current.getNext()
            k = k + 1

        if previous == None:
            self.add(item)
        elif current == None:
            self.append(item)
        else:
            temp = Node(item)
            previous.setNext(temp)
            temp.setNext(current)

    def pop(self, pos):
        current = self.head
        previous = None
        found = False
        count = 0

        while current != None and not found:
            if count == pos:
                found = True
            else:
                count = count + 1
                previous = current
                current = current.getNext()
        
        if previous == None:
            self.head = current.getNext()
        elif current == None:
            previous.setNext(None)
        else:
            previous.setNext(current.getNext())



myList = UnorderedList()
myList.add(1)
myList.add(2)
myList.add(3)
myList.pop(2)
print(myList.search(1))
print(myList.index(3))



