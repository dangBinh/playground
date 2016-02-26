from stack import Stack

def devideBy2(decimalNumber):
    s = Stack()

    while decimalNumber > 0:
        remain = decimalNumber % 2
        s.push(remain)
        decimalNumber = decimalNumber // 2

    binary = ""
    while not s.isEmpty():
        binary += str(s.pop())

    return binary

print(devideBy2(42))