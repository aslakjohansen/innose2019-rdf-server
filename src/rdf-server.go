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
    query, err := ioutil.ReadAll(request.Body)
    if err != nil {
        rw.Write([]byte("{\n"))
        rw.Write([]byte("    'success': false,\n"))
        rw.Write([]byte("    'error': 'unable to read query'\n"))
        rw.Write([]byte("}\n"))
        return
    }
    var query_str string = string(query)
    var query_len int    = len(query_str)
    
    // guard: proper string
    if len(query_str)<2 || query_str[0]!='"' || query_str[query_len-1]!='"' {
        rw.Write([]byte("{\n"))
        rw.Write([]byte("    'success': false,\n"))
        rw.Write([]byte("    'error': 'malformed query'\n"))
        rw.Write([]byte("}\n"))
        return
    }
    
    fmt.Println("rdf-server.go:query_handler", query_str[1:query_len-1])
    
    var success bool
    var result [][]string
    result, success = Query(query_str[1:query_len-1])
    
    if success==false {
        rw.Write([]byte("{\n"))
        rw.Write([]byte("    'success': false,\n"))
        rw.Write([]byte("    'error': 'unable to evaluate query'\n"))
        rw.Write([]byte("}\n"))
        return
    }
    
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    'success': %t,\n", success)))
    rw.Write([]byte("    'resultset': [\n"))
    for r:=0 ; r<len(result) ; r++ {
        rw.Write([]byte("        [\n"))
        
        row := result[r]
        for c:=0 ; c<len(row) ; c++ {
            cell := row[c]
            if c==len(row)-1 {
                rw.Write([]byte(fmt.Sprintf("            '%s'\n", cell)))
            } else {
                rw.Write([]byte(fmt.Sprintf("            '%s',\n", cell)))
            }
        }
        
        if r==len(result)-1 {
            rw.Write([]byte("        ]\n"))
        } else {
            rw.Write([]byte("        ],\n"))
        }
    }
    rw.Write([]byte("]\n"))
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
