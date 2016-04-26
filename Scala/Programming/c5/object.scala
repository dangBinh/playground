// The indentifiers in parentheness after the class name are called class parameters
// Cosntructor other than the primary constructor call auxiliary constructor
// An alphanumeric identifier starts with a letter or underscore which can be followed by further letters, digits, underscores
// An operator identifier consists of one or more operator character
// Mixed identifier unary_+, myvar_= used as method name defines an assignement operator
// Literal identifier is an arbitrary string enclosed in back ticks ` ` use for reserved in Scala
// Implicit conversation automatically converts integer to rational number

implicit def intToRational(x: Int) = new Rational(x)

class Rational(n: Int, d: Int) {
    require(d != 0)
    private val g = gcd(n.abs, d.abs)
    val number: Int = n / g
    val denom: Int = d / g
    def this(n: Int) = this(n, 1)
    override def toString = number + "/" + denom
    def + (that: Rational): Rational = 
        new Rational(number * that.denom + that.number * denom, denom * that.denom)
    def + (i :Int): Rational = 
        new Rational(number + i * denom, denom)
    def - (that: Rational): Rational = 
        new Rational(number * that.denom - that.number * denom, denom * that.denom)
    def - (i :Int): Rational = 
        new Rational(number - i * denom, denom)
    def * (that: Rational): Rational = 
        new Rational(number * that.number, denom * that.denom)
    def * (i: Int): Rational = 
        new Rational(number * i, denom)
    def / (that: Rational): Rational = 
        new Rational(number * that.denom, denom * that.number )
    def / (i: Int): Rational = 
        new Rational(number, denom * i)

    private def gcd(a: Int, b: Int): Int = {
        if (b == 0) a else gcd(b, a % b)
    }
}

val x = new Rational(1, 2)
println(new Rational(5))
println(new Rational(66, 42))

val m = new Rational(2, 3)
val l = new Rational(4, 5)
println(m + l)
println(l + 1)
println(1 + l)