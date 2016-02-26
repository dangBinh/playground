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
    let pattern = "({["
    while (i < symbols.length) {
        var symbol = symbols[i]
        if (pattern.indexOf(symbol) != -1) {
            s.push(symbol)
        } else {
            if (s.isEmpty()) {
                balanced = false
            } else {
                let top = s.pop()
                if (!matches(top, symbol)) {
                    balanced = false
                }
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

var matches = function(open, close) {
    let pattern_open = "({["
    let pattern_close = ")}]"
    return pattern_open.indexOf(open) == pattern_close.indexOf(close)
}

console.log(parChecker("(({{}}))"))
console.log(parChecker("(()"))