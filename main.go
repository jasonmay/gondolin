package main

import (
    "gondolin/habitat"
    "net/http"
    "fmt"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>hey</h1>")
}

func main() {
    go habitat.Run()

    http.HandleFunc("/", testHandler)
    http.ListenAndServe(":8080", nil)
}
