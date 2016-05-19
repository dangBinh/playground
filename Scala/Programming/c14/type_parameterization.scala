// A function queue is a data structure with three options:
// head return the first element of the queue
// tail return a queue without its first element
// enqueue return a new queue with a given element appended at the end
// Purely functional queues called fully persistent data structures where old versions remain available
// even after extensions or modifications
 
 class Queue[T](
    private val leading: List[T],
    private val trailing: List[T]
) {
    private def mirror = 
        if (leading.isEmpty)
            new Queue(trailing.reserve, Nil)
        else
            this
    def head = mirror.leading.head
    def tail = {
        val q = mirror
        new Queue(q.leading.tail, q.trailing)
    }
    def enqueue(x: T) = 
         new Queue(leading, x :: trailing)
 }

 // Information hiding
 // Private constructor and factory methods
 // Hide the primary constructor by adding a private modfier in front of the class paramater list
 
 class Queue[T] private {
    private val leading: List[T],
    private val trailing: List[T]
 }

// Queue is a trait and Queue[String] is a type
// Queue is also called a type contructor because with it you can construct a type by specifying a type parameter
// The term generic means that you are defining many specific types with one generically written  class or trait
// trait Queue is covariant example S is a subtype of T then should Queue[S] be a subtype of Queue[T]
// Generic types have by default nonvairant subtyping
// Covariant subtyping
trait Queue[+T] {
    
}

// Contravariant subtyping
trait  Queue[-T] {
    
}
// if T is a subtype of type S this would imply that Queue[S] is a subtype of Queue[T]
// A type parameter is covariant, nonvariant, contravariant is called parameter's variance
// The + and - symbols you can place next to type parameters are called variance annotations
// A positions is any location in class body where where a type paramter may be used

abstract class Cat[-T, +U] {
    def meow[W−](volume: T−, listener: Cat[U+, T−]−)
    : Cat[Cat[U+, T−]−, U+]+
}

class Queue[+T] (
    private val leading: List[T],
    private val trailing: List[T]
) {
    def enqueue(U >: T)(x : U) =
        new Queue[U](leading, x :: trailing)
}
// U >: T defines T as the lower bound for U, U is required to be a supertype of T


// Contravariance
// Upper bound