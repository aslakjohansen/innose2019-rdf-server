package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/gorilla/websocket"
)

var (
    model_dir    string = "../var/model"
    ontology_dir string = "../var/ontologies"
    upgrader            = websocket.Upgrader{}
)

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////// handlers

func time_handler (rw http.ResponseWriter, request *http.Request) {
    var result float64
    var success bool
    result, success = Time()
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    \"success\": %t,\n", success)))
    rw.Write([]byte(fmt.Sprintf("    \"time\": %f\n", result)))
    rw.Write([]byte("}\n"))
}

func store_handler (rw http.ResponseWriter, request *http.Request) {
    var success bool
    var result string
    result, success = Store(model_dir)
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    \"success\": %t,\n", success)))
    rw.Write([]byte(fmt.Sprintf("    \"filename\": %s\n", result)))
    rw.Write([]byte("}\n"))
}

func namespace_handler (rw http.ResponseWriter, request *http.Request) {
    var success bool
    var result map[string]string
    result, success = Namespaces()
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    \"success\": %t,\n", success)))
    rw.Write([]byte("    \"namespaces\": {\n"))
    var i    int = 0
    var last int = len(result)-1
    for key, value := range result {
        if i==last {
            rw.Write([]byte(fmt.Sprintf("        \"%s\": \"%s\"\n", key, value)))
        } else {
            rw.Write([]byte(fmt.Sprintf("        \"%s\": \"%s\",\n", key, value)))
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
        rw.Write([]byte("    \"success\": false,\n"))
        rw.Write([]byte("    \"error\": \"unable to read query\"\n"))
        rw.Write([]byte("}\n"))
        return
    }
    
    var query_str string
    err = json.Unmarshal(query, &query_str)
    if err!=nil {
        rw.Write([]byte("{\n"))
        rw.Write([]byte("    \"success\": false,\n"))
        rw.Write([]byte("    \"error\": \"malformed query\"\n"))
        rw.Write([]byte("}\n"))
        return
    }
    
    var success bool
    var result [][]string
    result, success = Query(query_str)
    
    if success==false {
        rw.Write([]byte("{\n"))
        rw.Write([]byte("    \"success\": false,\n"))
        rw.Write([]byte("    \"error\": \"unable to evaluate query\"\n"))
        rw.Write([]byte("}\n"))
        return
    }
    
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    \"success\": %t,\n", success)))
    rw.Write([]byte("    \"resultset\": [\n"))
    for r:=0 ; r<len(result) ; r++ {
        rw.Write([]byte("        [\n"))
        
        row := result[r]
        for c:=0 ; c<len(row) ; c++ {
            cell := row[c]
            if c==len(row)-1 {
                rw.Write([]byte(fmt.Sprintf("            \"%s\"\n", cell)))
            } else {
                rw.Write([]byte(fmt.Sprintf("            \"%s\",\n", cell)))
            }
        }
        
        if r==len(result)-1 {
            rw.Write([]byte("        ]\n"))
        } else {
            rw.Write([]byte("        ],\n"))
        }
    }
    rw.Write([]byte("    ]\n"))
    rw.Write([]byte("}\n"))
}

func update_handler (rw http.ResponseWriter, request *http.Request) {
    query, err := ioutil.ReadAll(request.Body)
    if err != nil {
        rw.Write([]byte("{\n"))
        rw.Write([]byte("    \"success\": false,\n"))
        rw.Write([]byte("    \"error\": \"unable to read query\"\n"))
        rw.Write([]byte("}\n"))
        return
    }
    
    var query_str string
    err = json.Unmarshal(query, &query_str)
    if err!=nil {
        rw.Write([]byte("{\n"))
        rw.Write([]byte("    \"success\": false,\n"))
        rw.Write([]byte("    \"error\": \"malformed query\"\n"))
        rw.Write([]byte("}\n"))
        return
    }
    
    var success bool
    _, success = Update(query_str)
    
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    \"success\": %t\n", success)))
    rw.Write([]byte("}\n"))
}

func websocket_handler (rw http.ResponseWriter, request *http.Request) {
    // upgrade connection
    ws, err := upgrader.Upgrade(rw, request, nil)
    if err!=nil {
        fmt.Println("Warn: Unable to upgrade to websocket:", err)
        return
    }
    defer ws.Close()
    
    // enter service loop
    for {
        // receive
        mt, message, err := ws.ReadMessage()
        if err!=nil {
            if err.Error()!="websocket: close 1000 (normal)" {
                fmt.Println("Warn: Unable to receive through websocket:", err)
            }
            return
        }
        
        // print to screen
        fmt.Println("Received: '", message, "'")
        
        // echo back for now
        err = ws.WriteMessage(mt, message)
        if err!=nil {
            fmt.Println("Warn: Unable to send through websocket:", err)
            return
        }
    }
}

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////// main

func main () {
//    var port int16 = 8001
    
    // guard: command line arguments
    if (len(os.Args) != 5) {
        fmt.Println("Syntax: "+os.Args[0]+" INTERFACE PORT MODEL_DIR ONTOLOGY_DIR")
        fmt.Println("        "+os.Args[0]+" 0.0.0.0 8001 ../var/model ../var/ontologies")
        os.Exit(1)
    }
    var iface string = os.Args[1]
    var port  string = os.Args[2]
        model_dir    = os.Args[3]
        ontology_dir = os.Args[4]
    
    Init(model_dir, ontology_dir)
    
    go func () {
        http.HandleFunc("/time"      , time_handler)
        http.HandleFunc("/store"     , store_handler)
        http.HandleFunc("/namespaces", namespace_handler)
        http.HandleFunc("/query"     , query_handler)
        http.HandleFunc("/update"    , update_handler)
        http.HandleFunc("/websocket" , websocket_handler)
        
        // start listening
        var endpoint string = fmt.Sprintf("%s:%s", iface, port)
        fmt.Println(fmt.Sprintf("Listening to %s", endpoint))
        err := http.ListenAndServe(endpoint, nil)
        if err != nil {
            fmt.Println(err)
        }
    }()
    
    select{} // block forever
    Finalize()
}
