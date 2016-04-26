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