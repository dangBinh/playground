"use strict"

class Node {
    constructor(initData) {
        this.data = initData
        this.next = null
    }

    getData () {
        return this.data
    }

    getNext() {
        return this.next
    }

    setData(dat) {
        this.data = data
    }

    setNext(next) {
        this.next = next
    }
}

class UnorderedList {
    constructor() {
        this.head = null
    }

    isEmpty() {
        return this.head == null
    }

    add(item) {
        var temp = new Node(item);
        temp.setNext(this.head);
        this.head = temp;
    }

    search(item) {
        var current = this.head
        var found = false
        while (current != null && !found) {
            if (current.getData() == item) {
                found = true
            } else {
                current = current.getNext()
            }
        }
        return found
    }

    remove(item) {
        var current = this.head
        var previous = null
        var found = false
        while (!found) {
            if(current.getData() == item) {
                found = true
            } else {
                previous = current
                current = current.getNext()
            }
        }
        if (previous == null) {
            this.head = current.getNext()
        } else {
            previous.getNext(current.getNext())
        }
    }

    index(item) {
        var current = this.head
        var found = false;
        var pos = 0
        while (current != null && !found) {
            if (current.getData() == item) {
                found = true
            } else {
                pos++
                current = current.getNext()
            }
        }
        return pos
    }

    insert(pos, item) {
        var current = this.head
        var previous = null
        var found = false
        var count = 0;
        while (!found) {
            if(pos == count) {
                found = true
            } else {
                count++
                previous = current
                current = current.getNext()
            }
        }
        if (previous == null) {
            // Head
            this.add(item)
        } else if(current == null) {
            // Tail
            var temp = new Node(item)
            temp.setNext(null)
            current.setNext(temp)
        } else {
            var temp = new Node(item)
            temp.setNext(current)
            previous.setNext(temp)
        }
    }

    pop(pos) {
        var current = this.head
        var previous = null
        var found = false
        var count = 0
        while (current != null && !found) {
            if (count == pos) {
                found == true
            } else {
                count++
                previous = current
                current = current.getNext()
            }
        }
        if (previous == null) {
            this.head = current.getNext()
        } else {
            previous.getNext(current.getNext())
        }
    }
}

var ul = new UnorderedList()
ul.add(1)
ul.add(2)
console.log(ul.search(2));
ul.remove(2)
console.log(ul.search(2))
ul.add(2)
console.log(ul.index(1))
ul.insert(1, 3)
console.log(ul.index(3))
