class Stack:
    def __init__(self):
        self.items = []

    def isEmpty(self):
        return self.items == []

    def push(self, item):
        return self.items.append(item)
        
    def peek(self):
        return self.items[len(self.items)-1]

    def pop(self):
        return self.items.pop()

    def size(self):
        return len(self.items)

s = Stack()
s.push('hello')
s.push('true')
print(s.pop())
