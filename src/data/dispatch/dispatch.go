package dispatch

import (
//    "fmt"
    
    "innose2019-rdf-server/data/reading"
)

type Dispatcher struct {
    lut map[string](chan reading.Reading)
}

func NewDispatcher () *Dispatcher {
    var d Dispatcher
    
    d.lut = make(map[string](chan reading.Reading))
    
    return &d
}

func (d Dispatcher) Register (id string, channel chan reading.Reading) chan reading.Reading {
    d.lut[id] = channel
    return channel
}

func (d Dispatcher) Dispatch (id string, r *reading.Reading) bool {
    channel, ok := d.lut[id]
    if ok {
        channel <- *r
    }
    return ok
}

