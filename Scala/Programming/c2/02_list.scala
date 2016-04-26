/*
    :: cons prepends a new element to the beginning of an existing list and return the resulting list
 */
val oneTwo = List(1, 2)
val threeFour = List(3, 4)
val oneTwoThreeFour = oneTwo ::: threeFour
println("1" + oneTwo + threeFour)
println("2" + oneTwoThreeFour)
val twoThree = List(2, 3)
val oneTwoThree = 1 :: twoThree
println(oneTwoThree)
val test = 1 :: Nil
println(test)