from stack import Stack

def infixToPostfix(infixexpr):
    precedence = {}
    precedence['*'] = 3
    precedence['/'] = 3
    precedence['+'] = 2
    precedence['-'] = 2
    precedence['('] = 1
    opstack = Stack()
    postfixLIst = [] # output list
    tokenList = infixexpr.split()

    for token in tokenList:
        if token in "ABCDEFGHIJKLMNOPQRSTUVWXYZ" or token in "0123456789":
            postfixLIst.append(token)
        elif token == '(':
            opstack.push(token)
        elif token == ')':
            topToken = opstack.pop()
            while topToken != '(':
                postfixLIst.append(topToken)
                topToken = opstack.pop()
        else:
            while (not opstack.isEmpty()) and (precedence[opstack.peek()] >= precedence[token]):
                postfixLIst.append(opstack.pop())
            opstack.push(token)

    while not opstack.isEmpty():
        postfixLIst.append(opstack.pop())

    return " ".join(postfixLIst)

print(infixToPostfix("A + ( B * C )"))

