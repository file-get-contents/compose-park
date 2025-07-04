package main

import (
        "fmt"
        "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World\n")
}

func main() {
        http.HandleFunc("/", Handler)
        err := http.ListenAndServe(":3000", nil)
        fmt.Println(err)
}
