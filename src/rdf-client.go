package main

import (
    "os"
    "fmt"
    
    "github.com/gorilla/websocket"
)

func main () {
    if (len(os.Args) != 3) {
        fmt.Println("Syntax: "+os.Args[0]+" INTERFACE PORT")
        fmt.Println("        "+os.Args[0]+" 127.0.0.1 8001")
        os.Exit(1)
    }
    var iface string = os.Args[1]
    var port  string = os.Args[2]
    
    fmt.Println(fmt.Sprintf("Connecting to %s:%s", iface, port))
}

