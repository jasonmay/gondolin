package habitat

/*
Habitat - the habitat in which the culture lives.

A more useful description: this is the hub in which the mobiles,
objects, and players all get moved around in a whirlwind of channel
dispatching and processing.
*/

import (
    "time"
    "gondolin/culture"
    "gondolin/loader"
)

func forward(f chan int, n int) {
    f <- n
    time.Sleep(time.Second * time.Duration(n))
    next := n % 5 + 1
    go forward(f, next)
}

func backward(b chan int, n int) {
    b <- n
    time.Sleep(time.Second * time.Duration(n))
    next := ((n + 3) % 5) + 1
    go backward(b, next)
}

func Run() {
    culture.Populate(loader.Load())

    f := make(chan int)
    b := make(chan int)

    go forward(f, 1)
    go backward(b, 1)

    for {
        select {
        case n := <-f:
            println("forward to ", n)
        case n := <-b:
            println("backward to ", n)
        }
    }
}
