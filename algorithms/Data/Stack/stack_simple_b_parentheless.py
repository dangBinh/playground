from stack import Stack

def parChecker(symbols):
    s = Stack()
    balanced = True
    index = 0
    while index < len(symbols) and balanced:
        symbol = symbols[index]
        if symbol == "(":
            s.push(symbol)
        else:
            if s.isEmpty():
                balanced = False
            else:
                s.pop()

        index = index + 1

    if balanced and s.isEmpty():
        return True
    else:
        return False

print(parChecker('((()))'))
print(parChecker('(()'))