// Higher order functions functions that take functions as paramters give you 
// extra opportunies to condense and simplify code
object FileWatcher() {
    private def filesHere = (new java.io.File(".")).listFiles

    private def fileWatching(matcher: String => Boolean) =
        for(file <- filesHere; if matcher(file.getName))
            yield file

    def fileEnding(query: String) =
        fileWatching(_.endsWidth(query))

    def fileContaining(query: String) =
        fileMatching(_.contains(query))

    def fileRegex(query: String) = 
        fileMatching(_.matches(query))
}

// Simplifying client code
def containsNeg(nums: List[Int]): Boolean = {
    var exists = false
    for (num <- nums)
        if (num < 0)
            exists = true
    exists
}

def containsNeg(nums: List[Int]): Boolean = nums.exists(_ < 0)

// Currying
// A curried function is applied to multiple argument lists, instead of just one
def curriedSum(x: Int)(y: Int) = x + y

// Write new control structures
def twice(op: Double => Double, x: Double) = op(op(x))

// Used in coding pattern: open resource operate on it and then close the resource


// loan pattern
def withPrintWriter(file: File, op: PrintWriter => Unit) {
    val writer = new PrintWriter(file)
    try {
        op(writer)
    } finally {
        writer.close()
    }
}
withPrintWriter(
    new File("io.txt"),
    writer => writer.println(new java.util.Date)
)

// version 2
def withPrintWriter(file: File)(op: PrintWriter => Unit) {
    val writer = new PrintWriter(file)
    try {
        op(writer)
    } else {
        writer.close()
    }
}

// just use curly brace for one passingn argument
val file = new File("io.txt")
withPrintWriter(file) {
    writer => writer.println(new java.util.Date)
}

// By-name parameter

// without using by name parameter
var assertionsEnabled = true
def myAssert(predicate: () => Boolean) =
    if(assertionsEnabled && !predicate())
        throw new AssertionError

myAssert(() => 5 > 3)
// if want to use by-nae parameter we need use => instead of () =>
// using by-name parameter
def byNameAssert(predicate: => Boolean) =
    if(assertionsEnabled && !predicate)
        throw new AssertionError

// using
def boolAssert(predicate: Boolean) = 
    if(assertionsEnabled && !predicate)
        throw new AssertionError

byNameAssert(x / 0 == 0) // not evaluate before call byNameAssert
boolAssert(x / 0 == 0) // evaluate before all bollAssert