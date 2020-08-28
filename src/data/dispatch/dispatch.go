package dispatch

import (
    "fmt"
    "sync"
    
    "innose2019-rdf-server/data/reading"
)

type Dispatcher struct {
    mutex sync.Mutex
    lut map[string](*[] chan reading.Reading)
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////// interface

func NewDispatcher () *Dispatcher {
    var d Dispatcher
    
    d.lut = make(map[string](*[]chan reading.Reading))
    
    return &d
}

func (d Dispatcher) Register (id string, channel chan reading.Reading) chan reading.Reading {
    d.mutex.Lock()
    entries, ok := d.lut[id]
    if !ok {
        fmt.Println("notok a")
        newalloc := make([]chan reading.Reading, 0)
        entries = &newalloc
        d.lut[id] = entries
        fmt.Println(d.lut[id])
        fmt.Println(entries)
        fmt.Println("notok b")
    }
    fmt.Println(entries)
    if !contains(entries, channel) {
        fmt.Println("Appending")
        fmt.Println(entries)
        *entries = append(*entries, channel)
        fmt.Println(entries)
    }
//    d.lut[id] = channel
    d.mutex.Unlock()
    return channel
}

func (d Dispatcher) Dispatch (id string, r *reading.Reading) bool {
    channels, ok := d.lut[id]
    d.Print()
    if ok {
        for _, channel := range *channels {
            channel <- *r
        }
    }
    return ok
}

func (d Dispatcher) Print () {
    fmt.Println("Dispatcher")
    for id, channels := range d.lut {
        fmt.Println(fmt.Sprintf(" - %s: %d entries", id, len(*channels)))
    }
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////// helpers

func contains (array *[]chan reading.Reading, value chan reading.Reading) bool {
    fmt.Println(array)
    for _, current := range *array {
        if current==value {
            return true
        }
    }
    return false
}

