from stack import Stack

def baseConverter(decimalNumber, base):
    digits = "0123456789ABCDEF"

    stack = Stack()
    while decimalNumber > 0:
        remain = decimalNumber % base
        stack.push(remain)
        decimalNumber = decimalNumber // base

    string = ""
    while not stack.isEmpty():
        string = string + digits[stack.pop()]

    return string

print(baseConverter(25,2))
