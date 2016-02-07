var string1 = "abcd"
var string2 = "bcda"

function anagramSolution2(s1, s2) {
    s1 = s1.split("").sort().join("")
    s2 = s2.split("").sort().join("")

    var pos = 0;
    var ok = true;

    while (pos < s1.length && ok) {
        if (s1[pos] == s2[pos]) {
            pos += 1
        } else {
            ok = false
        }
    }

    return ok
}

console.log(anagramSolution2(string1, string2))