package main

import (
    "fmt"
    "net/http"
)

func time_handler (rw http.ResponseWriter, request *http.Request) {
    var result float64 = 42.56
    rw.Write([]byte(fmt.Sprintf("%f",result)))
}

func store_handler (rw http.ResponseWriter, request *http.Request) {
    var result float64 = 42.56
    rw.Write([]byte(fmt.Sprintf("%f",result)))
}

func namespace_handler (rw http.ResponseWriter, request *http.Request) {
    var result float64 = 42.56
    rw.Write([]byte(fmt.Sprintf("%f",result)))
}

func query_handler (rw http.ResponseWriter, request *http.Request) {
    var result float64 = 42.56
    rw.Write([]byte(fmt.Sprintf("%f",result)))
}

func update_handler (rw http.ResponseWriter, request *http.Request) {
    var result float64 = 42.56
    rw.Write([]byte(fmt.Sprintf("%f",result)))
}

func main () {
    var port int16 = 8001
    
    go func () {
        http.HandleFunc("/time"      , time_handler)
        http.HandleFunc("/store"     , store_handler)
        http.HandleFunc("/namespaces", namespace_handler)
        http.HandleFunc("/query"     , query_handler)
        http.HandleFunc("/update"    , update_handler)
        
        fmt.Println(fmt.Sprintf("Listening to localhost:%d", port))
        
        // start listening
        var endpoint string = fmt.Sprintf(":%d", port)
        err := http.ListenAndServe(endpoint, nil)
        if err != nil {
            fmt.Println(err)
        }
    }()
    
    select{} // block forever
}
