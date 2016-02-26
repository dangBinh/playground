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

var s = new Stack()
s.push(1)
s.push(2)
console.log(s.size())