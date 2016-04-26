class Rational(n: Int, d: Int) {
    private def gcd(x: Int, y: Int): Int = {
        if (x == 0) y
        else if (x < 0) gcd(-x, y)
        else if (y < 0) gcd(x, y)
        else gcd(y % x, x)
    }

    private val g = gcd(n, d)

    val number: Int = n/g
    val denom: Int = d/g

    def +(that: Rational) = 
        new Rational(number * that.denom + that.number * denom, denom * that.denom)

    def -(that: Rational) = 
        new Rational(number * that.denom - that.number * denom, denom * that.denom)

    def *(that: Rational) =
        new Rational(number * that.number, denom * that.denom)
    def /(that: Rational) = 
        new Rational(number * that.denom, denom * that.number)
}

var i = 1
var x = new Rational(0, 1)
while(i <= 10) {
    x += new Rational(1, i)
    i += 1
}
println("" + x.number + "/" + x.denom)