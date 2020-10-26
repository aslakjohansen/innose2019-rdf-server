package main

import (
    "fmt"
    "sync"
    
    "github.com/gorilla/websocket"
)

const (
    url string = "ws://127.0.0.1:8001/websocket"
)

func consumer (wg *sync.WaitGroup, wgr *sync.WaitGroup, name string, id string) {
    var command   string
    var err       error
    var p       []byte
    
    // make connection
    con, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
        fmt.Println("Error: Unable to connect:", err)
        return
    }
    defer con.Close()
    
    // print
    fmt.Println("consumer "+name+": init")
    
    // subscribe
    command = fmt.Sprintf("{\"command\": \"subscribe\", \"id\": \"%s\", \"query\": \"SELECT ?rel WHERE { brick:missing1 ?rel brick:missing2 . }\"}", id)
    err = con.WriteMessage(websocket.TextMessage, []byte(command))
    if err != nil {
        fmt.Println("Error: Unable to write message", err)
        return
    }
    
    // print
    fmt.Println("consumer "+name+": subscribed")
    
    for {
        // read
        _, p, _ = con.ReadMessage()
        
        // print
        fmt.Println("consumer "+name+": "+string(p))
    }
}

func producer (wg *sync.WaitGroup, wgr *sync.WaitGroup, finish_channel chan bool) {
    var command   string
    var err       error
    var p       []byte
    
    // make connection
    con, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
        fmt.Println("Error: Unable to connect:", err)
        return
    }
    defer con.Close()
    
    // print
    fmt.Println("producer: init")
    
    // query (should be empty)
    command = fmt.Sprintf("{\"command\": \"query\", \"id\": \"%s\", \"query\": \"SELECT ?rel WHERE { brick:missing1 ?rel brick:missing2 . }\"}", "producer-id")
    err = con.WriteMessage(websocket.TextMessage, []byte(command))
    if err != nil {
        fmt.Println("Error: Producer unable to write message", err)
        return
    }
    _, p, _ = con.ReadMessage()
    
    // print
    fmt.Println("producer: starting point: "+string(p))
    
    // print
    fmt.Println("producer: insert")
    
    // insert
    command = fmt.Sprintf("{\"command\": \"update\", \"id\": \"%s\", \"query\": \"INSERT { brick:missing1 brick:rel1 brick:missing2 } WHERE {}\"}", "prod1")
    err = con.WriteMessage(websocket.TextMessage, []byte(command))
    if err != nil {
        fmt.Println("Error: Unable to write message", err)
        return
    }
    _, p, _ = con.ReadMessage()
    
    // query (should not be empty)
    command = fmt.Sprintf("{\"command\": \"query\", \"id\": \"%s\", \"query\": \"SELECT ?rel WHERE { brick:missing1 ?rel brick:missing2 . }\"}", "producer-id")
    err = con.WriteMessage(websocket.TextMessage, []byte(command))
    if err != nil {
        fmt.Println("Error: Producer unable to write message", err)
        return
    }
    _, p, _ = con.ReadMessage()
    fmt.Println("producer: next point: "+string(p))
    
    // print
    fmt.Println("producer: insert")
    
    // insert
    command = fmt.Sprintf("{\"command\": \"update\", \"id\": \"%s\", \"query\": \"INSERT { brick:missing1 brick:rel2 brick:missing2 } WHERE {}\" }", "prod1")
    err = con.WriteMessage(websocket.TextMessage, []byte(command))
    if err != nil {
        fmt.Println("Error: Unable to write message", err)
        return
    }
    _, p, _ = con.ReadMessage()
    
    // print
    fmt.Println("producer: delete")
    
    // delete
    command = fmt.Sprintf("{\"command\": \"update\", \"id\": \"%s\", \"query\": \"DELETE { brick:missing1 brick:rel1 brick:missing2 } WHERE {}\" }", "prod1")
    err = con.WriteMessage(websocket.TextMessage, []byte(command))
    if err != nil {
        fmt.Println("Error: Unable to write message", err)
        return
    }
    _, p, _ = con.ReadMessage()
    
    // print
    fmt.Println("producer: delete")
    
    // delete
    command = fmt.Sprintf("{\"command\": \"update\", \"id\": \"%s\", \"query\": \"DELETE { brick:missing1 brick:rel2 brick:missing2 } WHERE {}\" }", "prod1")
    err = con.WriteMessage(websocket.TextMessage, []byte(command))
    if err != nil {
        fmt.Println("Error: Unable to write message", err)
        return
    }
    _, p, _ = con.ReadMessage()
    
    // print
    fmt.Println("producer: finishing")
    
//    // signal finish
//    finish_channel <- true
}

func main () {
    var finish_channel chan bool = make(chan bool)
    var wg sync.WaitGroup
    var wgr sync.WaitGroup
//    wg.Add(3)
//    wgr.Add(1)
    
    // start 3 consumers with two different IDs
    go consumer(&wg, &wgr, "consumer1", "id1")
    go consumer(&wg, &wgr, "consumer2", "id2")
    go consumer(&wg, &wgr, "consumer3", "id2")
    
    // start producer
    go producer(&wg, &wgr, finish_channel)
    
    // wait for finish
    <- finish_channel
}

