package habitat

import (
    "gondolin/culture"
    "math/rand"
)

func GetRandomExit(loc culture.Location, pool Pool) (destination culture.Location) {
    exitKeys := make([]string, 0, len(loc.Exits))
    for _, k := range loc.Exits {
        if k.Type == "loc" {
            exitKeys = append(exitKeys, k.Entity)
        }
    }

    l := len(exitKeys)

    if l > 0 {
        r := rand.Int() % l
        exitValue := exitKeys[r]
        destination = pool.Locations[exitValue]
    }
    return
}

func (pool Pool) Travel(mobile culture.Mobile) {
    loc := mobile.Location
    exit := GetRandomExit(loc, pool)
    mobile.Location = exit
    println(loc.ID, "->", exit.ID)
}
