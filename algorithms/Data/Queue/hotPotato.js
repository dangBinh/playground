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
        return this.queue.pop()
    }

    size() {
        return this.queue.length
    }
}

var hotPotato = function(nameList, number) {
    let queue = new Queue()
    nameList.forEach(function(value, index) {
        queue.enqueue(value)
    });

    while (queue.size() > 1) {
        for (let i = 0; i < number; i++) {
            queue.enqueue(queue.dequeue())
        }
        queue.dequeue()
    }

    return queue.dequeue()
}

console.log(hotPotato(["A", "B", "C", "D", "E", "F"], 7))