// A function literal is compiled into a class that when instantiated at runtime is a function value
// Function literal exists in source code 
// Function value is that function literals exists as object at runtime

var increase = (x: Int) => x + 9999

// Short form function 
someNumbers.filter((x) => x > 0)
// x is called target typing because targeted usage of an expression is follow 
// influence the typing of that expression
someNumbers.filter(x => x > 0)

// Placeholder syntax
// You can use underscores as placeholder for one or more parameter
someNumbers.filter(_ > 0)
val f = (_: Int) + (_: Int)
// n-Underscore represents the n-parameter


// Partially applied function
someNumbers.foreach(println _)
// when you invoke a function passing in any needed arguments, you apply that function
// to the arguments
def sum (a: Int, b: Int, c: Int) = a + b + c
val a = sum _
a(1, 2, 3)
a.apply(1, 2, 3) 
