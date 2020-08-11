package transport

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    
    "github.com/gorilla/websocket"
    
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
    var result string = command.Time("")
    rw.Write([]byte(result))
}

func store_handler (rw http.ResponseWriter, request *http.Request) {
    var result string = command.Store("", model_dir)
    rw.Write([]byte(result))
}

func namespace_handler (rw http.ResponseWriter, request *http.Request) {
    var result string = command.Namespaces("")
    rw.Write([]byte(result))
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
    
    var result string = command.Query("", query_str)
    rw.Write([]byte(result))
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
    
    var result string = command.Update("", query_str)
    rw.Write([]byte(result))
}

func websocket_handler (rw http.ResponseWriter, request *http.Request) {
    // upgrade connection
    ws, err := upgrader.Upgrade(rw, request, nil)
    if err!=nil {
        fmt.Println("Warn: Unable to upgrade to websocket:", err)
        return
    }
    defer ws.Close()
    
    // response handling
    var response_channel chan []byte = make(chan []byte)
    go func () {
        for response := range response_channel {
            err = ws.WriteMessage(websocket.TextMessage, response)
            if err!=nil {
                fmt.Println("Warn: Unable to send through websocket:", err)
                return
            }
        }
    }()
    
    // enter service loop
    for {
        // receive
        _, message, err := ws.ReadMessage()
        if err!=nil {
            if err.Error()!="websocket: close 1000 (normal)" {
                fmt.Println("Warn: Unable to receive through websocket:", err)
            }
            return
        }
        
        // print to screen
        fmt.Println("Received: '", message, "'")
        
        // send off to processing
        go Dispatch(message, response_channel)
    }
}

