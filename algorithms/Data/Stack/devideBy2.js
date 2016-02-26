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

var devideBy2 = function(decimalNumber) {
    let s = new Stack()
    let remain = 0
    while (decimalNumber > 0) {
        remain = decimalNumber % 2
        s.push(remain)
        decimalNumber = ~~((decimalNumber) / 2)
    }
    var binary = ""
    while(!s.isEmpty()) {
        binary = binary + s.pop().toString()
    }

    return binary
}

console.log(devideBy2(42))


