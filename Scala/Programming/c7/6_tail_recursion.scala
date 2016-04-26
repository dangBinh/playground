// Tail recursive
// Tracing tail-recursive functions
// A tail recursive all calls will execute in a single frame
def bang(x: Int): Int = 
    if (x == 0) throw new Exception("bang!")
    else bang(x - 1)

// Limits of tail recursion
// Scala only optimizes directly recursive call back to the same function making the call

// Not for this situation
def isEven(x: Int): Boolean = 
    if (x == 0) true else isOdd(x - 1)
def isOdd(x: Int): Boolean = 
    if (x == 0) false else isEven(x - 1)