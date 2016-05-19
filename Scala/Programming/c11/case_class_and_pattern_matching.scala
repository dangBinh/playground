abstract class Expr
case class Var(name: String) extends Expr
case class Number(num: Double) extends Expr
case class UnOp(operator: String, arg: Expr) extends Expr
case class BinOp(operator: String, left: Expr, right: Expr) extends Expr

// Case classes are class with case modifier
// Var(x) to construct var object
// All argumments in the parameter list of a case class implicity get a val prefix
// The compiler adds natural implementation of methods toString, hashCode, and equals to your class
// The compiler adds a copy method to your class for making modified copies
// 
// 
// Pattern matching
def simplifyTop(expr: Expr): Expr = expr match {
    case UnOp("-", UnOp("-", e)) => e
    case BinOp("+", e, Number(0)) => e
    case BinOp("+", e, Number(1)) => e
    case _ => expr
}

// A pattern match inlucdes a sequence of alternatives each starting with the keyword case
// Each alternative includes a pattern and one or more expression which will evaluated if the pattern matches
// A match expression is evaluted by trying each of the patterns in the order they are written
// A constant patteren like "+" or 1 matches values that are equal to the constant with respect to ==
// A varialbe pattern like e matches every value
// The wildcard pattern (_) also matches every value, but it does not introduce a varialbe name to refer
// to that value
// A constructor pattern looks like UnOp("-", e). This pattern matches all values of type UnOp whose first
// argument matches "-" and whose second argument matches e

// Kinds of patterns
// The wildcard pattern matches any object
// Wild card pattern can be used to ignore parts of an object that you do not care about

expr match {
    case BinOp(_, _, _) => println(expr + "abc")
    case _ => 
}

// A constant pattern matches only itself
def describe(x: Any) = x match {
    case 5 => "five" 
}

// A variable pattern matches any object

// A constructor pattern
expr match {
    case BinOp("+", e, Number(0)) => println("Deep match")
    case _ => 
}

// Sequences pattern
expr match {
    case List(0, _, _) => println("found")
    case _ => 
}

expr match {
    case List(0, _*) => println("found")
    case _ =>  
}

// Tupple patterns
def tuppleDemo(expr: Any) = 
    expr match {
        case (a, b, c) => println(a, b, c)
        case _ =>
    }

// Typed patterns
def generalSize(x: Any) = x match {
    case s: String => s.length
    case m: Map[_, _] => m.size
    case _ => -1
}

// Variable binding
expr match {
    case UnOp("abs", e @ UnOp("abs", _)) => e
    case _ => 
}

// Pattern guard
// A pattern variable may only appear once in pattern
def simplifyAdd(e: Expr) = e match {
    case BinOp("+", x, y) if x == y => 
        BinOp("*", x, Number(2))
    case _ =>
}

// Pattern overlap

// Sealed class
sealed abstract class Expr {

}

def describe (e: Expr): String = (e @unchecked) match {
    case Number(_) => "a number"
    case Var(_) => "a var" 
}

// Option type
// Take optional value
def show(x: Option[String]) = x match {
    case Some(s) => s
    case None => "?"
}




    


