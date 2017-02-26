package habitat

import (
    "time"
    "gondolin/loader"
)

func Run() {
    // fmt.Printf("%v\n", loader.Load())
    loader.Load()
    //m := loader.Load()
    // fmt.Printf("%v\n", m.Mob[0].Properties[1].Name)
    // fmt.Printf("%v\n", m.Mob[0].Properties[1].Value)
    println("Loaded culture from JSON")

    timer := time.NewTimer(time.Second * 2)
    <- timer.C
    println("Timer expired")
}
