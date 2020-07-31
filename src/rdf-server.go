package main

import (
    "fmt"
    "net/http"
)

func time_handler (rw http.ResponseWriter, request *http.Request) {
    var result float64 = 42.56
    rw.Write([]byte(fmt.Sprintf("%f",result)))
}

func main () {
    var port int16 = 8001
    
    go func () {
        http.HandleFunc("/time", time_handler)
        
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
