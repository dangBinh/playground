/*
    While loop: imperative style
    Functional language is that functions are first class constructs
 */

args.foreach(arg => println(arg))

args.foreach((arg: String) => println(arg))

args.foreach(println)

/*
    Function literal: (x: Int, y: Int) => x + y
    for : a functional style in scala
 */

for (arg <- args) // arg is a val because you cannot be reassigned in for expression 
    println(arg)