//Habitat - the habitat in which the culture lives.
//A more useful description: this is the hub in which the mobiles,
//objects, and players all get moved around in a whirlwind of channel
//dispatching and processing.
package habitat

import (
    "time"
    "fmt"
    "math/rand"
    "strconv"
    "gondolin/loader"
    "gondolin/culture"
)

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

// moveMobileTick waits N seconds and then calls itself again in a goroutine,
// where N is a gaussian distributed number of "speed" mean and "speed * 0.1"
// unit stddev.
// The legacy approach to this was "every 2 seconds, loop through all mobiles
// and do a computational thing all in one tight loop" which seems a little
// choppy for a smooth flowing game. If there were two or more mobiles in the same
// location, and multiple returned positive for the random movement test, any player
// in that location would see all the mobiles leave at the exact same time. This
// variance-adorned tick approach results in a smoother flow of movement among
// the map.
func moveMobileTick(mobile culture.Mobile, moveTick chan culture.Mobile) {
    moveTime := 200.0 / float64(mobile.Speed)
    moveTimeWithVariance := rand.NormFloat64() * (moveTime * 0.1) + moveTime
    time.Sleep(time.Millisecond * time.Duration(moveTimeWithVariance * 1000.0))
    moveTick <- mobile
    go moveMobileTick(mobile, moveTick)
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
    fmt.Printf("loaded %d locations\n", len(pool.Locations))

    for _, m := range m.Mob {

        cm := culture.Mobile{
            MoveChan: make(chan float64),
        }
        cm.ID = m.ID

        cm.Speed = 0
        for _, p := range m.Properties {
            switch p.Name {
            case "speed":
                speed, err := strconv.Atoi(p.Value)
                cm.Speed = speed
                if err != nil {
                    println(err)
                    continue
                }
                break
            case "location":
                cm.Location = pool.Locations[p.Value]
                break
            }
        }
        pool.Mobiles[m.ID] = cm
    }
    fmt.Printf("loaded %d mobiles\n", len(pool.Mobiles))

    return pool
}

func Run() {
    pool := populate(loader.Load())

    moveTick := make(chan culture.Mobile)

    for _, m := range pool.Mobiles {
        if m.Speed > 0 {
            go moveMobileTick(m, moveTick)
        }
    }

    for {
        select {
        case m := <-moveTick:
            println("tick!", m.ID)
            pool.Travel(m)
        }
    }
}
