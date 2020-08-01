package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

var (
    model_dir    string = "../var/model"
    ontology_dir string = "../var/ontologies"
)

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////// handlers

func time_handler (rw http.ResponseWriter, request *http.Request) {
    var result float64
    var success bool
    result, success = Time()
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    'success': %t,\n", success)))
    rw.Write([]byte(fmt.Sprintf("    'time': %f\n", result)))
    rw.Write([]byte("}\n"))
}

func store_handler (rw http.ResponseWriter, request *http.Request) {
    var success bool
    var result string
    result, success = Store(model_dir)
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    'success': %t,\n", success)))
    rw.Write([]byte(fmt.Sprintf("    'filename': %s\n", result)))
    rw.Write([]byte("}\n"))
}

func namespace_handler (rw http.ResponseWriter, request *http.Request) {
    var success bool
    var result map[string]string
    result, success = Namespaces()
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    'success': %t,\n", success)))
    rw.Write([]byte("    'namespaces': {\n"))
    var i    int = 0
    var last int = len(result)-1
    for key, value := range result {
        if i==last {
            rw.Write([]byte(fmt.Sprintf("        '%s': '%s'\n", key, value)))
        } else {
            rw.Write([]byte(fmt.Sprintf("        '%s': '%s',\n", key, value)))
        }
        i++
    }
    rw.Write([]byte("    }\n"))
    rw.Write([]byte("}\n"))
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

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////// main

func main () {
    var port int16 = 8001
    
    Init(model_dir, ontology_dir)
    
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
    Finalize()
}
