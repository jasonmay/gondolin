package culture

type Location struct {
    ID string
    Title string
    Description string
    Exits map[string]Exit
}

type Exit struct {
    Direction string
    Entity string
    Type string
}
