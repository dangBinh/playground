var string1 = "abcd"
var string2 = "bcda"

function anagramSoluiton3(s1, s2) {
    var c1 = new Array(26+1).join('0').split('').map(parseFloat)
    var c2 = new Array(26+1).join('0').split('').map(parseFloat)

    for (var i = 0; i < s1.length; i++) {
        var pos = s1[i].charCodeAt(0) - "a".charCodeAt(0)
        c1[pos] = c1[pos] + 1
    }

    for (var i = 0; i < s2.length; i++) {
        var pos = s2[i].charCodeAt(0) - "a".charCodeAt(0)
        c2[pos] = c2[pos] + 1
    }

    var j = 0
    var ok = true

    while (j < 26 && ok) {
        if (c1[j] == c2[j]) {
            j++
        } else {
            ok = false
        }
    }

    return ok
}

console.log(anagramSoluiton3(string1, string2))