// Closure
(x: Int) => x + more
// more is a free variable because function literal does not itself give a meaning to it
// the x variable is a bound vairable because it does have a meaning in the context of the function

var more = 1
val addMore = (x: Int) => x + more

// The function value that's created at runtime form funciton literal is called a closure
// meaning: act of "closing" function literal by capturing the bindings its free variable
// A function literal with no free variables is called a close term
// Any function leteral with free variables is an open term

def makeIncreaser(more: Int) = (x: Int) => x + more // closure