function toStr(number, base) {
    var convertString = "0123456789ABCDEF"
    if (number < base) {
        return convertString[number]
    } else {
        return toStr(Math.floor(number / base), base) + convertString[number % base]
    }
}

console.log(toStr(1435, 16))