import scala.collection.mutable.Set
import scala.collection.immutable.HashSet
import scala.collection.mutable.Map
/*
    Set
 */
var jetSet = Set("Boeing", "Airbus") // immutable set
jetSet += "Lear"
println(jetSet.contains("Cessna"))

val movieSet = Set("Hitch", "Poltergeist")
movieSet += "Shrek"
println(movieSet)

/*
    Hash Set
 */
 val hashSet = HashSet("Tomatoes", "Chillies")
 println(hashSet + "Coriander")

 /*
    Map
  */
 val treasureMap = Map[Int, String]()
 treasureMap += (1 -> "Go to 1")
 treasureMap += (2 -> "Go to 2")
 treasureMap += (3 -> "Go to 3")
 println(treasureMap(2))

/*
    Functional programming in scala is no vars and side effect
 */
