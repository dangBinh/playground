/*
    Function defines start with def
    Function's name is followed by a comma-sperated list of parameters in parentheses
    A type annotation must follow every function parameter preceded by a colon
    Function's result type
    If the function is recursive you may specify the function's result type
    You may leave the result type and the compiler will infer it
    If function consists of just one statement you can leave off curly brace
    A result type Unit indicates the function return no interesting value (Unit type == void type in Java)
 */

def max(x: Int, y: Int): Int = {
    if (x > y)
        x
    else
        y
}

def max2(x: Int, y: Int) = if (x > y) x  else y

max(3, 5)

def greet() = print("Hello") // function take no paramter
