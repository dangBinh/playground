"use strict"
class HashTable {
    contructor() {
        this.slots = new Array(11)
        this.size = 11
    }

    put(key, value) {
        var hashValue = hash(key)

    }

    get(key) {
        return this.slots[key]
    }

    hash(key) {
        return key % this.size
    }

    reHash(key) {
        return (key + 1) % this.size
    }
}
// Lol Javascript Object is Hash Table