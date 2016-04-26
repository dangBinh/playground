// Two-dimensional layout library

// Layout elements are objects can be constructed from simple parts with the aid of composing operators
// Composing operator == combinators because they combine elements of some domain into new elements


// Abstract classes with abstract members must itself be declared abstract
// Abstract modifier signifies that the class may have abstract members that do not have an implementation 
// Method that do have an implementation are called concrete
// Class elemment delcares the abstract method contents, but defines no concrete methods


/*********************/
// Defining parameterless methods
// def width(): int empty parentheness method
// Should not define a method that has side-effect without parentheless


/********************/
// Extending classes
// ArrayElement call subclass
// Element call superclass
// Inheritance means that all members of the superclass are also members of the subclass
// with two exceptions. First, private members of the superclass are not inherited in a subclass.
// Second a member of superclass is not inherited if a member with the same name and parameters 
// is already implemented in the subclass -> overrides
// If the member in the subclass is concrete and the member of the superclass is abstract we also
// say that the concrete member implements the abstract one
// Subtyping means that a value of the subclass can be used wherever a value of the superclass
// is required
// The composition relationship that exists between ArrayElement and Array[String] because 
// ArrayElement is composed out of class Array[String]
 
/********************/
// Overidding methods and fields
// Field contents is a implementaion of the parentheless method contents

/********************/
// Parametric fields
// class ArrayElement (val contents: Array[String]) extends Element

/********************/
// Superclass constructors


/********************/
// Polymorphism and dynamic binding
// 

/*******************/
// Factory object
// Factory object contains methods that construct other object
abstract class Element {
    def contents: Array[String] // is declared as a method that has no implementation
    val height: Int = contents.height // parameterless method
    val width: Int = if (height == 0) 0 else contents(0).length
    def above(that: Element): Element = 
        new ArrayElement(this.contents ++ that.contents) // ++ concatenates two array

    def beside(that: Element): Element = {
        val contents = new Array[String](this.contents.length)
        for (i <- 0 until this.contents.length)
            contents(i) = this.contents(i) + that.contents(i)
        new ArrayElement(contents)
    }
    override def toString = contents mkString "\n"
}

class UniformElement (
        ch: Char,
        override val width: Int,
        override val height: Int
    ) extends Element {
        private val line = ch.toString * width
        def contents = Array.fill(height)(line)
    }

class ArrayElement(conts: Array[String]) extends Element {
    def contents: Array[String] = conts
} 

class LineElement(s: String) extends ArrayElement(Array(s)) { // Invode superclass constructor
    override def width = s.length
    override def height = 1
}

object Element {
    def elem(contents: Array[String]): Element = 
        new ArrayElement(contents)

    def elem(chr: Char, width: Int, height: Int): Element =
        new UniformElement(chr, width, height)

    def elem(line: String): Element =
        new LineElement(line)
}





