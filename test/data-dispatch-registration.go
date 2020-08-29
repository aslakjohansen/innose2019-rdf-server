package main

import (
    "fmt"
    
    "innose2019-rdf-server/data/reading"
    "innose2019-rdf-server/data/dispatch"
)

func main () {
    var d *dispatch.Dispatcher = dispatch.NewDispatcher()
    
    fmt.Println("Adding:")
    fmt.Println("~~~~~~")
    d.Register("a", make(chan reading.Reading))
    d.Register("a", make(chan reading.Reading))
    d.Register("a", make(chan reading.Reading))
    d.Register("b", make(chan reading.Reading))
    d.Register("b", make(chan reading.Reading))
    d.Register("c", make(chan reading.Reading))
    d.Register("a", make(chan reading.Reading))
    d.Register("b", make(chan reading.Reading))
    d.Register("c", make(chan reading.Reading))
    d.Register("d", make(chan reading.Reading))
    d.Print()
    fmt.Println("")
}
