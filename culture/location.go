package culture

type Location struct {
    Title string
    Exits []Exit
}

type Exit struct {
    Direction string
    Entity string
    Type string
}
