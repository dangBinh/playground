"use strict"
class Queue {
    constructor() {
        this.queue = new Array()
    }

    isEmpty() {
        this.queue.length <= 0
    }

    enqueue(item) {
        this.queue.unshift(item)
    }

    dequeue() {
        this.queue.pop(item)
    }

    size() {
        return this.queue.length
    }
}

var q = new Queue()
q.enqueue(1)
console.log(q.size())