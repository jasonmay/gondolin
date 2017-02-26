package loader
import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
)

type Loc struct {
    ID string `json:"id"`
    Description string `json:"description"`
    Title string `json:"title"`
    Exits []Exit `json:"exits"`
    Properties []Property `json:"properties"`
}

type Exit struct {
    Name string `json:"name"`
    Entity string `json:"entity"`
    Type string `json:"type"`
}

type Obj struct {
    Properties []Property `json:"properties"`
}

type Mob struct {
    ID string `json:"id"`
    Properties []Property `json:"properties"`
}

type Property struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

type Message struct {
    Loc []Loc `json:"loc"`
    Obj []Obj `json:"obj"`
    Mob []Mob `json:"mob"`
}

func Load() Message {
    f, e := ioutil.ReadFile("zones.json")
    if e != nil {
        fmt.Printf("eyyy\n")
        os.Exit(1)
    }
    var m Message
    json.Unmarshal(f, &m)

    return m
}
