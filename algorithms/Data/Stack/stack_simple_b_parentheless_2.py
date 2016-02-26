from stack import Stack

def parChecker(symbols):
    s = Stack()
    balanced = True
    index = 0
    while index < len(symbols) and balanced:
        symbol = symbols[index]
        pattern = "({["
        if symbol in pattern:
            s.push(symbol)
        else:
            if s.isEmpty():
                balacend = False
            else:
                top = s.pop()
                if not matches(top, symbol):
                    balacend = False
        index = index + 1

    if s.isEmpty() and balanced:
        return True
    else:
        return False

def matches(open, close):
    pattern_open = "({["
    pattern_close = ")}]"
    return pattern_open.index(open) == pattern_close.index(close)

print(parChecker('(({{}}))'))
print(parChecker("(({{}})"))



