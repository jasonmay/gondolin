package main

import (
    "fmt"
    "gondolin/loader"
)

func main() {
    // fmt.Printf("%v\n", loader.Load())
    m := loader.Load()
    fmt.Printf("%v\n", m.Mob[0])
}
