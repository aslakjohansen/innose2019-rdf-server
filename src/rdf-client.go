package main

import (
    "os"
    "fmt"
    "os/signal"
    "sync"
    "time"
    
    "github.com/gorilla/websocket"
)

var (
    buffer       []string   = make([]string, 0)
    buffer_mutex sync.Mutex
)

func receiver (con *websocket.Conn, receiver_closed chan struct{}) {
    defer close(receiver_closed)
    
    for {
        // read a single message
        _, message, err := con.ReadMessage()
        if err != nil {
            fmt.Println("Error: Unable to read", err)
            return
        }
        
        // add to input buffer
        buffer_mutex.Lock()
        buffer = append(buffer, string(message))
        buffer_mutex.Unlock()
    }
}

func main () {
    // args
    if (len(os.Args) != 3) {
        fmt.Println("Syntax: "+os.Args[0]+" INTERFACE PORT")
        fmt.Println("        "+os.Args[0]+" 127.0.0.1 8001")
        os.Exit(1)
    }
    var iface string = os.Args[1]
    var port  string = os.Args[2]
    
    // setup signal handler for ^C
    var interrupt       chan os.Signal = make(chan os.Signal, 1)
    var receiver_closed chan struct{}  = make(chan struct{})
    signal.Notify(interrupt, os.Interrupt)
    
    // url
    var url string = fmt.Sprintf("ws://%s:%s/websocket", iface, port)
    fmt.Println("Connecting to", url)
    
    // connect
    con, _, err := websocket.DefaultDialer.Dial(url, nil)
    if err != nil {
        fmt.Println("Error: Unable to connect:", err)
        return
    }
    defer con.Close()
    
    go receiver(con, receiver_closed)
    
    // main loop
    for {
        select {
            case <-receiver_closed:
                return
            case <-interrupt:
                fmt.Println("")
                
                // send close message
                err := con.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
                if err != nil {
                    fmt.Println("Error: Unable to send closing message", err)
                    return
                }
                
                // wait for server to close connection
                select {
                case <-receiver_closed:
                case <-time.After(time.Second): // TODO: Can we make this a push operation?
                }
                
                return
        }
    }
}

