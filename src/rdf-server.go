package main

import (
    "fmt"
    "io/ioutil"
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
    input, err := ioutil.ReadAll(request.Body)
    if err != nil {
        fmt.Println(err)
    }
    var inputString string = string(input)
    fmt.Println("received '"+inputString+"'")
    rw.Write([]byte(inputString))
}

func update_handler (rw http.ResponseWriter, request *http.Request) {
    input, err := ioutil.ReadAll(request.Body)
    if err != nil {
        fmt.Println(err)
    }
    var inputString string = string(input)
    
    rw.Write([]byte(inputString))
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
