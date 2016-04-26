import scala.collection.mutable.Map

class ChecksumAccumulator {
    private var sum = 0

    def add(b: Byte) { sum += b }

    def checkSum(): Int = ~(sum & 0xFF) + 1
}

// Static method in Java == Singleton Object in Scala
object ChecksumAccumulator {
    private val cache = Map[String, Int]()

    def calculate(s: String): Int = 
        if (cache.contains(s))
            cache(s)
        else {
            val acc = new ChecksumAccumulator
            for (c <- s) 
                acc.add(c.toByte)
            val cs = acc.checkSum()
            cache += (s -> cs)
            cs
        }

}

// Singleton object shares the same name with a class and it's called that class's companion object
// The class is called the companion class of the singleton object 
// A class and its companion object can access each other's private members
// Singleton is first-class object
// Singleton object can take parameters whereas class companion
// Can not initial singleton object with new keyword
// A singleton object does not share the same name with a companion class is called a standalone object
