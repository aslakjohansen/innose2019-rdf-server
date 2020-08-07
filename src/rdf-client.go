package main

import (
    "os"
    "fmt"
    "sync"
    "time"
    
    "github.com/gorilla/websocket"
    "github.com/peterh/liner"
)

var (
    history_filename   string
    buffer           []string = make([]string, 0)
    buffer_mutex       sync.Mutex
)

func buffer_add (entry string) {
    buffer_mutex.Lock()
    buffer = append(buffer, string(entry))
    buffer_mutex.Unlock()
}

func buffer_remove () *string {
    var line *string = nil
    buffer_mutex.Lock()
    if len(buffer)>0 {
        line = &buffer[0]
        buffer = buffer[1:]
    }
    buffer_mutex.Unlock()
    return line
}

func receiver (con *websocket.Conn, receiver_closed chan struct{}, sender_closed *bool) {
    defer close(receiver_closed)
    
    for {
        // read a single message
        _, message, err := con.ReadMessage()
        if err != nil {
            if !*sender_closed {
                fmt.Println("Error: Unable to read", err)
            }
            return
        }
        
        // add to input buffer
        buffer_add(string(message))
    }
}

func command_reader (command_channel chan string, interrupt_channel chan os.Signal) {
    line_handler := liner.NewLiner()
    line_handler.SetCtrlCAborts(true)
    defer line_handler.Close()
    
    // load history
    if _, err := os.Stat(history_filename); !os.IsNotExist(err) {
        if fo, err := os.Open(history_filename); err == nil {
            line_handler.ReadHistory(fo)
            fo.Close()
        }
    }
    
    // main loop
    for {
        line, err := line_handler.Prompt(">> ")
        if err!=nil {
            if err!=liner.ErrPromptAborted {
                fmt.Println("Error: Unable to read input:", err)
            }
            interrupt_channel <- os.Interrupt
            return
        }
        
        line_handler.AppendHistory(line)
        
        // store history
        fo, err := os.Create(history_filename)
        if err!=nil {
            fmt.Println("Error: Unable to create history file:", err)
            interrupt_channel <- os.Interrupt
            return
        }
        line_handler.WriteHistory(fo)
        fo.Close()
        
        command_channel <- line
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
    
    // build history filename
    home, err := os.UserHomeDir()
    if err!=nil {
        fmt.Println("Error: Unable to determine home directory of user:", err)
        return
    }
    history_filename = home+"/.rdf-client_history"
    
    // setup signal handler for ^C
    var interrupt       chan os.Signal = make(chan os.Signal, 2)
    var receiver_closed chan struct{}  = make(chan struct{})
    var command_channel chan string    = make(chan string)
    var sender_closed   bool           = false
    
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
    
    go receiver(con, receiver_closed, &sender_closed)
    go command_reader(command_channel, interrupt)
    
    // main loop
    for {
        select {
            case command := <-command_channel:
                if command=="\n" || command=="" {
                    for {
                        entry := buffer_remove()
                        if entry==nil {
                            break
                        }
                        fmt.Println(*entry)
                    }
                } else {
                    err := con.WriteMessage(websocket.TextMessage, []byte(command))
                    if err != nil {
                        fmt.Println("Error: Unable to write message", err)
                        return
                    }
                }
            case <-receiver_closed:
                return
            case <-interrupt:
                sender_closed = true
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

