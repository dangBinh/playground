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

    size() {
        return this.items.length
    }
}

var parChecker = function (symbols) {
    var s = new Stack()
    var balanced = true
    var i = 0
    while (i < symbols.length) {
        var symbol = symbols[i]
        if (symbol == "(") {
            s.push(symbol)
        } else {
            if (s.isEmpty()) {
                balanced = false
            } else {
                s.pop()
            }
        }
        i++
    }
    if (balanced && s.isEmpty()) {
        return true
    } else {
        return false
    }
}

console.log(parChecker("(())"))
console.log(parChecker("(()"))