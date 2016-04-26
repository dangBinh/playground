// Repeated parameters
def echo(args: String*) =
    for (arg <- args) println(arg)
// String* = Array[String]
val arr = Array("What", "the", "f")

echo(arr: _*)

// Named arguments
def speed(distance: Int, time: Float) =
    distance / time

speed(time = 10, distance = 100)
// Named arguments allow you to pass arguments to a function in a differenct order
// The syntax is preceded by a paramter name and an equals sign

// Default parameter values
def printTime(out: java.io.PrintStream = Console.out) = 
    out.println("time = " + System.currentTimeMillis())
