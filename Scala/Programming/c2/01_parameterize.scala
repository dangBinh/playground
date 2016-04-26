/*
    Parameterization means configuring an instance when you create it
    You parameterize instance with value by passing objects to a constructor in parenthese
    You paremeterize instance with type by specifying one or more types in square brackets
 */

val big = new java.math.BigInteger("12345")

val greetingStrings = new Array[String](3)
greetingStrings(0) = "Hello"
greetingStrings(1) = "1"

for (i <- 0 to 1)
    print(greetingStrings(i))

val greetingStrings2 = new Array[String](3)
greetingStrings2.update(0, "Hello")
greetingStrings2.update(1, "Hi")

for (i <- 0.to(1))
    print(greetingStrings2.apply(i))