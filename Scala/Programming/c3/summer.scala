import ChecksumAccumulator.calculate

object Summer {
    def main(args: Array[String]) {
        for (arg <- args) {
            println(arg + ": " + calculate(arg))
        }
    }
}

// Any standalone object with main method of the proper signature can be used as the entry point into 
// application