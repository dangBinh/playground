// Implicit conversion
implicit def function2ActionListener(f: ActionEvent => Unit) = 
    new ActionListener {
        def actionPerformed(event: ActionEvent) = f(event)
    }

button.addActionListener(
    (_: ActionEvent) => println("pressed!")
)

// Marking Rule: Only definitions marked implicit are available
// Scope Rule: An inserted implicit conversion must be in scope as a single 
// identifier or be associated with the source or target type of the converion
// One-at-a-time Rule: Only one implicit is tried
// Explicits-First Rule: whenerver code type checks as it is written, no implicits are attempted


// Implicit conversion to an expected type
// is the first place the compiler will use implicits whenever the compiler sees an X, but needs
// a Y it will look for an implicit function that convert X to Y    

// Converting the receiver ap dung goi method 

// Implicit parameter

// View bounds
// T <% Ordered[T]: I can use any T so long T can be treated as an Ordered[T]