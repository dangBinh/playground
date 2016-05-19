// List literal
val fruit = List("apples", "oranges", "pears")
// Lists are immutable
// Lists have a recursive structure

// List type
// lists are homogenous: the elements of a list all have the same type.
// List type is covariant if S is a subtype of T then List[S] is a subtype of List[T]

// Constructing lists
// List build from two fundamental building blocks, Nil and ::
// Nill == empty list
// :: list extension at the front
// x :: xs represents a list whose first element is x followed by list xs

// Basic operations on list
// head
// tail
// isEmpty

// List patterns
val List(a, b, c) = fruit

// Concatenating two lists
List(1, 2) ::: List(3)

// Divide and conquer principle
def append(xs: List[T], ys: List[T]): List[T] {
    xs match {
        case List() => ys
        case x : xs1 => x :: append(xs1, ys)
    }
}

// Length
List(1, 2, 3).Length

val abcde = List('a', 'b', 'c', 'd', 'e')
abcde.last // e
abcde.init // a b c d return a list consiting of all elements except the last one
abcde.reverse // e d c b a
abcde take 2 // List(a, b)
abcde drop 2 // List(c ,d, e)
abcde splitAt 2 // List(a, b), List(c, d, e) return two list
abcde apply 2 // char = c
abcde.indices // Range(0, ,1 ,2 ,3 , 4)
// flattern methods takes a list of lists and flatterns it out to a single list
abcde.indices zip abcde // (0, a), (1, b)
abcde.zipWithIndex
zipped.unzip
abcde.toString // String = List(a, b, c, d, e)
xs mkString (pre, sep, post)
xs mkString sep xs mkString ("", sep, "")
val arr = abcde.toArray
arr.toList
xs copyToArray(arr, start)
val it = abcde.iterator
it.next
it.next

// High-order meethods on class List
// Mapping over lists
// The operation xs map f takes as operands a list xs of type List[T] and a function f of type T => U
List(1, 2, 3) map (_ + 1)
// flatMap returns single list in which single list are concatenated
val words = List("the", "quick")
words flatMap 
var sum = 0
List(1, 2, 3, 4, 5) foreach(sum+_)

// Filter lists
// The operation xs filter p takes as operands a list xs of type List[T] and a predicate function p 
// of type T => Boolean
List(1, 2, 3, 4, 5) filter (_ % 2 == 0)
// partition like filter but return pair of list
List(1, 2, 3, 4, 5) partition (_ % 2 == 0)
// find return first element
List(1, 2, 3, 4, 5) find (_ % 2 == 0) // 2
// takeWhile p takes the longest prefix of list xs such that every element in the prefix satifies p
// dropWhile p removes the longest prefix from list xs such that every element in the perfix satisfies p

// span xs span p equals (xs takeWhile(p), xs dropWhile(p))
List(1, 2, 3, 4, 5) span (_ > 0)

// xs forall p takes as arguments a list xs and a predicate p. Its results is true if all elements in the
// list satisfy p
// xs exists p return true if there is an element in xs that satifies the predicate p

def hasZeroRow(m List[List[Int]]) =
    m exists (row => row forall (_ == 0))

// folding lists /: and :\
def sum(xs: List[Int]): Int = (0 /: xs) (_ + _)
def product(xs: List[Int]): Int = (1 /: xs) (_ * _)

// A fold left operatioin (z /: xs) (op) involves three object: a start value z, a list xs and a binary operation op
// A fold right (xs :\ z) (op)

// sortWith
List(1, -3, 4, ,2 ,6) sortWith (_ < _)

// List method
List(1, 2, 3)
List.apply(1, 2, 3)
List.range(1, 5) // List.range(start, until)
List.fill(5)('a')
List.tabulate(5)(n => n * n)
List.concat(List('a', 'b'), List('c'))
