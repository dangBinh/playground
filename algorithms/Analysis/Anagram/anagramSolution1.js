var str1 = "abcd"
var str2 = "bcda"

function anagramSolution1(str1, str2) {
    var ok = true
    var pos1 = 0
    while (pos1 < str1.length && ok) {
        var found = false
        var pos2 = 0
        while (pos2 < str2.length && !found) {
            if (str1[pos1] == str2[pos2]) {
                found = true
            } else {
                pos2 += 1
            }
        }
        if (found) {
            delete str2[pos2]
        } else {
            ok = false
        }
        pos1 += 1
    }
    return ok  
}

var check = anagramSolution1(str1, str2)
console.log(check)