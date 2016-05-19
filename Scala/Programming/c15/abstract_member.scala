trait Abstract {
    type T // abstract type
    def transform(x: T): T
    val initial: T // abstract val
    var current: T // abstract var
}

class Concrete extends Abstract {
    type T = String
    def transform(x: String) = x + x
    val initial = "hi"
    var current = initial
}

// Term Abstract Type means a type declared to be a member of a class or trait 
// without specifying a definition
// Abstract val 
// Abstract var
// Pre-initialized fields
new {
    val numberArg = 1 * x
    val denomArg = 2 * x
} with RationalTrait

// Lazy vals
// lazy val never evaluate more than once
object Demo {
    lazy val x = { println("initializing x"); "done"}
}

// Abstract type
class Food 
abstract class Animal {
    type SuitableFood <: Food
    def eat(food: SuitableFood)
}

class Grass extends Food
class Cow extends Animal {
    type SuitableFood = Grass
    override def eat(food: Grass) {}
}

// path-dependent type this type consists of an object reference ??
// structural subtyping ??
// enumerations ??