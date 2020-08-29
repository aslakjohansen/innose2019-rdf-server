package main

import (
    "fmt"
    
    "innose2019-rdf-server/data/reading"
    "innose2019-rdf-server/data/dispatch"
)

func main () {
    var d *dispatch.Dispatcher = dispatch.NewDispatcher()
    var as [] chan reading.Reading = [] chan reading.Reading {
        make(chan reading.Reading),
        make(chan reading.Reading),
        make(chan reading.Reading),
        make(chan reading.Reading),
    }
    var bs [] chan reading.Reading = [] chan reading.Reading {
        make(chan reading.Reading),
        make(chan reading.Reading),
        make(chan reading.Reading),
    }
    var cs [] chan reading.Reading = [] chan reading.Reading {
        make(chan reading.Reading),
        make(chan reading.Reading),
    }
    var ds [] chan reading.Reading = [] chan reading.Reading {
        make(chan reading.Reading),
    }
    
    fmt.Println("Adding:")
    fmt.Println("~~~~~~")
    d.Register("a", as[0])
    d.Register("a", as[1])
    d.Register("a", as[2])
    d.Register("b", bs[0])
    d.Register("b", bs[1])
    d.Register("c", cs[0])
    d.Register("a", as[3])
    d.Register("b", bs[2])
    d.Register("c", cs[1])
    d.Register("d", ds[0])
    d.Print()
    fmt.Println("")
}
