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

function infixToPostfix(infixexpr) {
   let precedence = {
        '*': 3,
        '/': 3,
        '+': 2,
        '-': 2,
        '(': 1
   }
   let opStack = new Stack()
   let postfix = []
   let tokenList = infixexpr.split(" ")

   tokenList.forEach(function(token, index) {
        if ("ABCDEFGHIJKLMNOPQRSTUVWXYZ".indexOf(token) != -1 || "0123456789".indexOf(token) != -1) {
            postfix.push(token)
        } else if (token == '(') {
            opStack.push(token)
        } else if (token == ')') {
            let topToken = opStack.pop()
            while(topToken != "(") {
                postfix.push(topToken)
                topToken = opStack.pop()
            }
        } else {
            while(!opStack.isEmpty() && (precedence[opStack.peek()] >= precedence[token])) {
                postfix.push(opStack.pop())
            }
            opStack.push(token)
        }
   });

   while(!opStack.isEmpty()) {
    postfix.push(opStack.pop())
   }

   return postfix.join(" ")
}

console.log(infixToPostfix("A + ( B * C )"));