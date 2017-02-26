package habitat

/*
Habitat - the habitat in which the culture lives.

A more useful description: this is the hub in which the mobiles,
objects, and players all get moved around in a whirlwind of channel
dispatching and processing.
*/

import (
    "time"
    "gondolin/loader"
    "gondolin/culture"
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

type Pool struct {
    Mobiles map[string]culture.Mobile
    Locations map[string]culture.Location
    Objects map[string]culture.Object
}

func NewPool() Pool {
    pool := Pool{}
    pool.Mobiles = map[string]culture.Mobile{}
    pool.Locations = map[string]culture.Location{}
    pool.Objects = map[string]culture.Object{}

    return pool
}

func populate(m loader.Message) Pool {
    pool := NewPool()
    for _, l := range m.Loc {

        cl := culture.Location{}
        cl.ID = l.ID
        cl.Title = l.Title
        cl.Description = l.Description
        cl.Exits = map[string]culture.Exit{}
        for _, ex := range l.Exits {
            exit := culture.Exit{ex.Direction, ex.Entity, ex.Type}
            cl.Exits[ex.Direction] = exit
        }

        pool.Locations[l.ID] = cl
    }

    return pool
}

func Run() {
    //pool := populate(loader.Load())
      populate(loader.Load())

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
