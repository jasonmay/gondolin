package culture

import (
    "fmt"
    "gondolin/loader"
)

type Pool struct {
    mobiles map[string]Mobile
    locations map[string]Location
    objects map[string]Object
}

func NewPool() Pool {
    pool := Pool{}
    pool.mobiles = map[string]Mobile{}
    pool.locations = map[string]Location{}
    pool.objects = map[string]Object{}

    return pool
}

func Populate(m loader.Message) Pool {
    println("yeah something should happen here")
    pool := NewPool()
    for _, l := range m.Loc {

        cl := Location{}
        cl.Title = l.Title

        pool.locations[l.ID] = cl
    }

    fmt.Printf("%v", pool)
    return pool
}
