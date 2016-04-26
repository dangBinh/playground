"use strict"
class Deque {
    constructor() {
        this.deque = new Array()
    }

    isEmpty() {
        return this.deque.length == 0
    }

    addFront(item) {
        this.deque.push(item)
    }

    addRear(item) {
        this.deque.unshift(item)
    }

    removeFront() {
        return this.deque.pop()
    }

    removeRear() {
        return this.deque.shift()
    }

    size() {
        return this.deque.length
    }
}

function palChecker(aString) {
    var deque = new Deque()

    for (let i = 0; i < 1; i++) {
        deque.addRear(aString[i])
    }

    var stillEqual = true

    while (deque.size() > 1 && stillEqual) {
        first = deque.removeFront()
        last = deque.removeRear()
        if (first != last)
            stillEqual = false
    }

    return stillEqual
}

console.log(palChecker("radar"))