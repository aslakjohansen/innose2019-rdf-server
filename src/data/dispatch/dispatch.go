package dispatch

import (
    "fmt"
    "sync"
    
    "innose2019-rdf-server/data/reading"
)

var (
    mutex  sync.Mutex
    D     *Dispatcher = nil
)

// TODO: Determine if this struct is necessary and we need the cancel field (or if the mutex in dispatch is enough)
type DispatchEntry struct {
    stream chan reading.Reading
    cancel chan bool
}

type Dispatcher struct {
    mutex sync.Mutex
    lut map[string](*[] DispatchEntry)
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////// interface

func GetDispatcher ()  *Dispatcher {
    mutex.Lock()
    defer mutex.Unlock()
    
    if D==nil {
        D = NewDispatcher()
    }
    
    return D
}

func NewDispatcher () *Dispatcher {
    var d Dispatcher
    
    d.lut = make(map[string](*[] DispatchEntry))
    
    return &d
}

func (d Dispatcher) Register (id string, channel chan reading.Reading) chan reading.Reading {
    d.mutex.Lock()
    
    entries, ok := d.lut[id]
    if !ok {
        newalloc := make([] DispatchEntry, 0)
        entries = &newalloc
        d.lut[id] = entries
    }
    
    if !contains(entries, channel) {
        var entry DispatchEntry = DispatchEntry{channel, make(chan bool)}
        
        *entries = append(*entries, entry)
    }
    
    d.mutex.Unlock()
    return channel
}

func (d Dispatcher) Unregister (id string, channel chan reading.Reading) bool {
    fmt.Println("data/dispatch:.Unregister first", id)
    d.mutex.Lock()
    defer d.mutex.Unlock()
    
    // locate group
    entries, ok := d.lut[id]
    if !ok {
        return false
    }
    
    // locate entry
    i, ok := locate(entries, channel)
    if !ok {
        fmt.Println("data/dispatch:.Unregister entry not located")
        return false
    }
    
    // array manipulation
    var entry DispatchEntry = (*entries)[i]
    (*entries)[i] = (*entries)[len(*entries)-1]
    *entries = (*entries)[:len(*entries)-1]
    
    // cleanup
    go func (entry DispatchEntry) {
        fmt.Println("data/dispatch:.Unregister cleanup a")
//        entry.cancel <- true // signal cancelation
        close(entry.stream)
        fmt.Println("data/dispatch:.Unregister cleanup b")
        for range entry.stream {} // empty stream for remaining data
        fmt.Println("data/dispatch:.Unregister cleanup c")
    }(entry)
    
    return true
}

func (d Dispatcher) Dispatch (id string, r *reading.Reading) bool {
    d.mutex.Lock()
    defer d.mutex.Unlock()
    
    entries, ok := d.lut[id]
    
    if ok {
        for _, entry := range *entries {
            entry.stream <- *r
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

func contains (array *[] DispatchEntry, value chan reading.Reading) bool {
    for _, current := range *array {
        if current.stream==value {
            return true
        }
    }
    return false
}

func locate (array *[] DispatchEntry, value chan reading.Reading) (int, bool) {
    for i:=0 ; i<len(*array) ; i++ {
        if (*array)[i].stream == value {
            return i, true
        }
    }
    
    return -1, false
}
