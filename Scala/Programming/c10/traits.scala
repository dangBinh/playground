// A trait definition look like a class
trait Philosofical {
    def philosophize() {
        println("I cosumem memory")
    }
}

// A trait is defined, it can be mixed in to class using either the extends or with keywords

class Frog extends Philosofical {
    override def toString = "green"
}

class Animal
class Frog extends Animal with Philosofical {
    override def toString = "green"
}

class Point(val x: Int, val y: Int)

abstract class Component {

}

trait Rectangular {

}

abstract class Component extends Rectangular {
    
}

class Rectange(val: topLeft Point, val: bottomRight: Point) extends Rectangular {
    
}

class Rational(n: Int, d: Int) extends Ordered[Rational] {
    def compare(that: Rational) = 
        (this.number * that.denom) - (that.number * this.denom)
}

// Stackable traits
