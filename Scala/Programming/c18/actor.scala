import scala.actors._

object SillyActor extends Actor {
    def act() {
        for(i <- to 5) {
            println("I'm acting!")
            Thread.sleep(1000)
        }
    }
}

SillyActor.start()

import scala.actors.Actor._

val seriousActor2 = actor {
    for (i <- 1 to 5)
        println("That is the question")
    Thread.sleep(1000)
}

// Actors communicate by sending each other messages
// Send a message by using the ! method
