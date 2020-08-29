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
        newalloc := make([]chan reading.Reading, 0)
        entries = &newalloc
        d.lut[id] = entries
    }
    
    if !contains(entries, channel) {
        *entries = append(*entries, channel)
    }
    
    d.mutex.Unlock()
    return channel
}

func (d Dispatcher) Unregister (id string, channel chan reading.Reading) bool {
    d.mutex.Lock()
    defer d.mutex.Unlock()
    
    entries, ok := d.lut[id]
    if !ok {
        return false
    }
    
    i, ok := locate(entries, channel)
    if !ok {
        return false
    }
    
    (*entries)[i] = (*entries)[len(*entries)-1]
    *entries = (*entries)[:len(*entries)-1]
    
    return true
}

func (d Dispatcher) Dispatch (id string, r *reading.Reading) bool {
    channels, ok := d.lut[id]
    
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
    for _, current := range *array {
        if current==value {
            return true
        }
    }
    return false
}

func locate (array *[]chan reading.Reading, value chan reading.Reading) (int, bool) {
    for i:=0 ; i<len(*array) ; i++ {
        if (*array)[i] == value {
            return i, true
        }
    }
    
    return -1, false
}
