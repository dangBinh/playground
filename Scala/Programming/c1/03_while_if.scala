/*
    Block code between curly braces
    i++, ++i don't work in scala
 */
var i = 0
while (i < args.length) {
    if (i != 0) {
        println(args(i))
    }
    i += 1
}