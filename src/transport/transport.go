package transport

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    
    "github.com/gorilla/websocket"
    
    "innose2019-rdf-server/config"
    "innose2019-rdf-server/logic"
    "innose2019-rdf-server/session"
    . "innose2019-rdf-server/responseconduit"
)

var (
    model_dir *string
    upgrader          = websocket.Upgrader{}
)

type TransportModuleConfig struct {
    config.ModuleConfig
    Interface string `json:"interface"`
    Port      int    `json:"port"`
    Modeldir  string `json:"modeldir"`
}

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////// lifecycle management

func Init (configraw *json.RawMessage) {
    var config TransportModuleConfig
    
    // parse config
    err := json.Unmarshal(*configraw, &config)
    if err!=nil {
        fmt.Println("Unable to unmarshal config for module 'transport':", err)
    }
    model_dir = &config.Modeldir
    
    go func () {
        http.HandleFunc("/time"      , time_handler)
        http.HandleFunc("/store"     , store_handler)
        http.HandleFunc("/namespaces", namespace_handler)
        http.HandleFunc("/query"     , query_handler)
        http.HandleFunc("/update"    , update_handler)
        http.HandleFunc("/inspect"   , inspect_handler)
        http.HandleFunc("/websocket" , websocket_handler)
        
        // start listening
        var endpoint string = fmt.Sprintf("%s:%d", config.Interface, config.Port)
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
    var result string = logic.JsonTime("")+"\n"
    rw.Write([]byte(result))
}

func store_handler (rw http.ResponseWriter, request *http.Request) {
    var result string = logic.JsonStore("", model_dir)+"\n"
    rw.Write([]byte(result))
}

func namespace_handler (rw http.ResponseWriter, request *http.Request) {
    var result string = logic.JsonNamespaces("")+"\n"
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
    
    var result string = logic.JsonQuery("", query_str)+"\n"
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
    
    var result string = logic.JsonUpdate("", query_str)+"\n"
    rw.Write([]byte(result))
}

func inspect_handler (rw http.ResponseWriter, request *http.Request) {
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
    
    var result string = logic.JsonInspect("", query_str)+"\n"
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
    var rc *ResponseConduit = NewResponseConduit()
    go func () {
        // for pointer := range response_channel {
        for pointer := range rc.Channel {
            message, err := json.Marshal(pointer)
            if err!=nil {
                fmt.Println("Warn: Unable to marshal message for websocket transmision:", err)
                continue
            }
            err = ws.WriteMessage(websocket.TextMessage, message)
            if err!=nil {
                fmt.Println("Warn: Unable to send through websocket:", err)
                return
            }
        }
    }()
    
    // enter service loop
    var s *session.Session = session.NewSession(rc)
    for {
        // receive
        rc.Hello(nil)
        _, message, err := ws.ReadMessage()
        if err!=nil {
            if err.Error()!="websocket: close 1000 (normal)" {
                fmt.Println("Warn: Unable to receive through websocket:", err)
            }
            rc.Goodbye()
            rc.Finalize()
            return
        }
        
        // send off to processing
        go func() {
            Dispatch(message, s)
            rc.Goodbye()
        }()
    }
}

