package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
    "strconv"
)

var x = 0

func inc(){
    x += 1
}

func tracer(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(64)
    
    trace := strings.TrimSpace(r.FormValue("trace"))
    time := strings.TrimSpace(r.FormValue("name"))
    
    inc()
    
    fmt.Println(strconv.Itoa(x) + " - "+ time)
    fmt.Println("---------------------")
    fmt.Println(trace)
    fmt.Println("---------------------")
    fmt.Println("\n")
    fmt.Println("\n")
}

func main() {
    fmt.Println("Tracer")
    fmt.Println("\n")
    http.HandleFunc("/", tracer) // set router
    err := http.ListenAndServe(":5003", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
