package transport

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    
    "github.com/gorilla/websocket"
    
    "innose2019-rdf-server/logic"
    "innose2019-rdf-server/command"
)

var (
    model_dir *string
    upgrader          = websocket.Upgrader{}
)

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////// lifecycle management

func Init (iface string, port string, model_dir_arg *string) {
    model_dir    = model_dir_arg
    
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
}

func Finalize () {
}

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////// handlers

func time_handler (rw http.ResponseWriter, request *http.Request) {
        var result_channel chan string = make(chan string)
        go command.Time(result_channel)
        var result string = <- result_channel
        rw.Write([]byte(result))
//    var result float64
//    var success bool
//    result, success = logic.Time()
//    rw.Write([]byte("{\n"))
//    rw.Write([]byte(fmt.Sprintf("    \"success\": %t,\n", success)))
//    rw.Write([]byte(fmt.Sprintf("    \"time\": %f\n", result)))
//    rw.Write([]byte("}\n"))
}

func store_handler (rw http.ResponseWriter, request *http.Request) {
    var success bool
    var result string
    result, success = logic.Store(*model_dir)
    rw.Write([]byte("{\n"))
    rw.Write([]byte(fmt.Sprintf("    \"success\": %t,\n", success)))
    rw.Write([]byte(fmt.Sprintf("    \"filename\": %s\n", result)))
    rw.Write([]byte("}\n"))
}

func namespace_handler (rw http.ResponseWriter, request *http.Request) {
    var success bool
    var result map[string]string
    result, success = logic.Namespaces()
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
    result, success = logic.Query(query_str)
    
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
    _, success = logic.Update(query_str)
    
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

