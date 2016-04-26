import ChecksumAccumulator.calculate

object FallWilterSpringSummer extends App {
     for (session <- List("fall", "spring", "witer")) {
        println(session + ": " + calculate(session))
     }
}

// The code between curly braces is collected into primary constructor of singleton object and it
// is executed when the class is initialized
