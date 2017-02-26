package culture

import (
    "fmt"
    "gondolin/loader"
)

type Pool struct {
    Mobiles map[string]Mobile
    Locations map[string]Location
    Objects map[string]Object
}

func NewPool() Pool {
    pool := Pool{}
    pool.Mobiles = map[string]Mobile{}
    pool.Locations = map[string]Location{}
    pool.Objects = map[string]Object{}

    return pool
}

func Populate(m loader.Message) Pool {
    println("yeah something should happen here")
    pool := NewPool()
    for _, l := range m.Loc {

        cl := Location{}
        cl.Title = l.Title

        pool.Locations[l.ID] = cl
    }

    fmt.Printf("%v", pool)
    return pool
}
