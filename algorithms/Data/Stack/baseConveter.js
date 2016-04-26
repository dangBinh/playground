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

var baseConverter = function(decimalNumber) {
    let s = new Stack()
    let remain = 0
    let digits = "0123456789ABCDEF"

    while (decimalNumber > 0) {
        remain = decimalNumber % 2
        s.push(remain)
        decimalNumber = ~~((decimalNumber) / 2)
    }
    var binary = ""
    while(!s.isEmpty()) {
        binary = binary + digits[s.pop()]
    }

    return binary
}

console.log(baseConverter(25, 2))


