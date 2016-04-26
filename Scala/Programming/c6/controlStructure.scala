// If expression
var filename = "default.txt"
if (!args.isEmpty)
    filename = args(0)

// Final variable in java
val filename = 
    if (!args.isEmpty) "default.txt"
    else args(0)

// equational reasoning: The introduced variable is equal to the epxression that computes it, assuming 
// that expression has no side effects

def gcdLoop(x: Long, y: Long): Long = {
    var a = x
    var b = y
    while (a != 0) {
        val temp = a
        a = b % a
        b = temp
    }
}

var line = ""
do {
    line = readline()
    println("Read:" + line)
} while(line != "")

def gcd(x: Long, y: Long): Long = 
    if (y == 0) x else gcd(y, x % y)

// while loop need to update vars or perform I/O


// For
// Iteration through collections
val fileHere = (new java.io.File(".")).listFiles
for (file <- filesHere) 
    println(file)
// file <- fileHere syntax called a generator a new val named file is initialized with an element value
// The for  syntax works for any kind of collection

// Filtering
// Filter with for by adding a filter: an if clause inside the for's parentheness
val filesHere = (new java.io.File(".")).listFiles

for (file <- filesHere if file.getName.endsWith(".scala"))
    println(file)
// For expression called and expression because it can result in an interesting value
// a collection whose type  is determined by the for expression's <- clauses

// Multi filter
for (
    file <- filesHere
    if file.isFile
    if file.getName.endsWith(".scala")
)
    println(file)

// Nested iteration
def fileLines(file: java.io.File) = 
    scala.io.Source.fromFile(file).getLines().toList

def grep(pattern: String) = 
    for (
        file <- filesHere
        if file.getName.endsWith(".scala")
        line <- fileLines
        if line.trim.matches(pattern)
    ) println(file + ": " + line.trim)

grep(".*gcd*.")

// Mid-stream variable bindings
def grep(pattern: String) = 
    for (
        file <- filesHere
        if file.getName.endsWith(".scala")
        line <- fileLines
        trimmed = line.trim
        if trimmed.matches(pattern)
    ) println(file + ": " + trimmed)

// Producing new collection
// for clause yield body
// Geneerate value and to remember for each iteration
val forLineLengths = 
    for {
        file <- filesHere
        if file.getName.endsWith(".scala")
        line <- fileLines(file)
        trimmed = line.trim
        if trimmed.matches(".for.")
    } yield trimmed.length

throw new RuntimeException("abc")

try {

    } catch {

    }

// The finally clause
// is a way to ensure some side effect happens, such as closing a file

// Match expression
// lets you select from a number of alternatives just like a switch
// _ underscored is a wildcard symbol frequently used in Scala as a placeholder for completely 
// unknown value

val firstArg = if (args.length > 0) args(0) else ""

firstArg match {
    case "salt" => println("pepper")
    case _ => "??"
}

// match result as value
val firstArg = if (ars.length > 0) ars(0) else ""

val friend = 
    firstArg match {
        case "salt" => "pepper"
        case _  => "fuck"
    }

println(friend)

// Living without break and continue
// replace continue by and if and break by a boolean variable

var i = 0
var foundIt = false

while (i < args.length && !foundIt) {
    if(!args(i).startsWith("-")) {
        if (args(i).endsWith(".scala"))
         foundIt = true
    }
    i = i + 1
}

// remove vars

def searchForm(i: Int) = 
    if(i >= args.length) -1
    else if(args(i).startsWith("-") searchForm(i + 1))
    else if(args(i).endsWith(".scala")) i
    else searchForm(i + 1)

val i = searchForm(0)

// break method which can used to exit the an enclosing block that's marked with breakable

// Variable scope
val a = 1 
{
    val a = 2 // an inner variable is said to shadow a like-named outer variable,
    // because the outer variable becomes invisible in the inner scope  
}
