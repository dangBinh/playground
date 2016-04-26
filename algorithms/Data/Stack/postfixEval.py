from stack import Stack

def postfixVal(postfixExpr):
    operandStack = Stack()
    tokenList = postfixExpr.split(" ")
    
    for token in tokenList:
        if token in "0123456789":
            operandStack.push(int(token))
        else:
            operand2 = operandStack.pop()
            operand1 = operandStack.pop()
            operandStack.push(doMath(token, operand1, operand2))

    return operandStack.pop()

def doMath(token, operand1, operand2):
    if token == "*":
        return operand1 * operand2
    elif token == "/":
        return operand1  / operand2
    elif token == "+":
        return operand1 + operand2
    elif token == "-":
        return operand1 - operand2

print(postfixVal("8 3 2 + *"))