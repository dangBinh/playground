"use strict"
class Stack {
    constructor() {
        this.items = new Array()
    }

    isEmpty() {
        return this.items.length == 0
    }

    push(item) {
        return this.items.push(item)
    }

    pop() {
        return this.items.pop()
    }

    peek() {
        return this.items[this.items.length - 1]
    }

    size() {
        return this.items.length
    }
}

function postfixEval(postfixexpr) {
    let operandStack = new Stack();
    let tokenList = postfixexpr.split(" ")

    tokenList.forEach(function(token, index) {
        if ("0123456789".indexOf(token) != -1) {
            operandStack.push(parseInt(token))
        } else {
            let operand1 = operandStack.pop()
            let operand2 = operandStack.pop()
            let result = doMath(token, operand1, operand2)
            operandStack.push(result)
        }
    })

    return operandStack.pop()
}

function doMath(token, operand1, operand2) {
    switch (token) {
        case "*":
            return operand1 * operand2
        case "/":
            return operand1 / operand2
        case "-":
            return operand1 - operand2
        case "+":
            return operand1 + operand2
    }
}

console.log(postfixEval("8 3 2 + *"));