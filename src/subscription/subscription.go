package subscription

import (
    "fmt"
    "sync"
    
    // "innose2019-rdf-server/sparql"
    "innose2019-rdf-server/logic"
)

var (
    dispatch_mux sync.Mutex
    dispatch map[string](*DispatchEntry) = make(map[string](*DispatchEntry))
)

// type ResultSet ([][]string)

type DispatchEntry struct {
    Cache            *([][]string)
    Subscriptions *[]*Subscription
}

type ResultDiff struct {
    // Plus  *([][]string) `json:"+"`
    // Minus *([][]string) `json:"-"`
    Plus  [][]string `json:"+"`
    Minus [][]string `json:"-"`
}

type Subscription struct {
    id              string
    Query           string
    ResponseChannel chan []byte
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////// resultset

func NewResultDiff () *ResultDiff {
    var result ResultDiff
    result.Plus  = nil
    result.Minus = nil
    return &result
}

func (r *ResultDiff) Transmit (channel chan []byte, id string) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("  \"type\": \"resultset\",\n")
    response += fmt.Sprintf("  \"id\": \"%s\",\n", id)
    response += fmt.Sprintf("  \"+\": [\n")
    for i, row := range r.Plus {
        response += fmt.Sprintf("    [\n")
        for j, cell := range row {
            response += fmt.Sprintf("      \"%s\"", cell)
            if j!=len(row)-1 {
                response += fmt.Sprintf(",")
            }
            response += fmt.Sprintf("\n")
        }
        response += fmt.Sprintf("    ]")
        if i!=len(r.Plus)-1 {
            response += fmt.Sprintf(",")
        }
        response += fmt.Sprintf("\n")
    }
    response += fmt.Sprintf("  ],\n")
    response += fmt.Sprintf("  \"-\": [\n")
    for i, row := range r.Minus {
        response += fmt.Sprintf("    [\n")
        for j, cell := range row {
            response += fmt.Sprintf("      \"%s\"", cell)
            if j!=len(row)-1 {
                response += fmt.Sprintf(",")
            }
            response += fmt.Sprintf("\n")
        }
        response += fmt.Sprintf("    ]")
        if i!=len(r.Minus)-1 {
            response += fmt.Sprintf(",")
        }
        response += fmt.Sprintf("\n")
    }
    response += fmt.Sprintf("  ]\n")
    response += fmt.Sprintf("}")
    
    channel <- []byte(response)
}

func resultset_diff (a *([][]string), b *([][]string)) *ResultDiff {
    var result ResultDiff
    var found_row bool
    resultset_print(a)
    fmt.Println("")
    resultset_print(b)
    fmt.Println("subscription.resultset_diff");
    
    // TODO: will this work?
    result.Plus  = nil
    result.Minus = nil
    
    // fill out.Transmit minus
    for _, rowa := range *a {
        found_row = false
        for _, rowb := range *b {
            var match bool = true
            for i, elema := range rowa {
                elemb := rowb[i]
                if elema != elemb {
                    match = false
                    break
                }
            }
            if match {
                found_row = true
                break
            }
        }
        if !found_row {
            fmt.Println("  minus added")
            result.Minus = append(result.Minus, rowa)
        }
    }
    
    // fill out plus
    for _, rowb := range *b {
        found_row = false
        for _, rowa := range *a {
            var match bool = true
            for i, elemb := range rowb {
                elema := rowa[i]
                if elemb != elema {
                    match = false
                    break
                }
            }
            if match {
                found_row = true
                break
            }
        }
        if !found_row {
            fmt.Println("  plus added")
            result.Plus = append(result.Plus, rowb)
        }
    }
    
    return &result
}

func resultset_print (data *[][]string) {
    for _, row := range *data {
        for _, cell := range row {
            fmt.Print("["+cell+"]")
        }
        fmt.Println("")
    }
}

///////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////// dispatch entry

func NewDispatchEntry () *DispatchEntry {
    var de DispatchEntry
    
    de.Cache = nil
    tmp := make([]*Subscription, 0)
    de.Subscriptions = &tmp
    
    return &de
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////// interface functions

// TODO: mutex around?
func Update () {
    for query, de := range dispatch {
        var result  [][]string
        var success bool
        result, success = logic.Query(query)
        
        // guard: no success
        if !success {
            fmt.Println("Unable to perform query during update.")
            continue
        }
        
        var diff *ResultDiff = resultset_diff(de.Cache, &result)
        
        for _, subscription := range *de.Subscriptions {
            diff.Transmit(subscription.ResponseChannel, subscription.id)
        }
        
        de.Cache = &result
    }
}

func Unsubscribe (id string) bool {
    // TODO: This id is only unique through the context of a session. The session should have an id->subscription dictionary and call destroy on unsubscription
    return true
}

func NewSubscription (id string, query string, response_channel chan []byte) *Subscription {
    var s Subscription
    
    s.id              = id
    s.Query           = query
    s.ResponseChannel = response_channel
    
    // add to dispatch data structure
    dispatch_mux.Lock()
    defer dispatch_mux.Unlock()
    de, ok := dispatch[query]
    if !ok {
        result, success := logic.Query(query)
        if !success {
            fmt.Println("TODO: send back an error message from subscription.go:NewSubscription")
            return nil // TODO: Check this on the caller side
        }
        de = NewDispatchEntry()
        de.Cache = &result
        dispatch[query] = de
    }
    tmp := append(*de.Subscriptions, &s)
    de.Subscriptions = &tmp
    
    return &s
}

func (s *Subscription) Destroy () {
    // remove from dispatch data structure
    dispatch_mux.Lock()
    subscriptions := dispatch[s.Query].Subscriptions
    for i, elem := range *subscriptions {
        if elem==s {
            tmp := append((*subscriptions)[:i], (*subscriptions)[i+1:]...)
            dispatch[s.Query].Subscriptions = &tmp
            break
        }
    }
    dispatch_mux.Unlock()
}

func (s *Subscription) String () string {
    var result string = ""
    
    result += fmt.Sprintf("subscription")
    
    return result
}

func (s *Subscription) Push () {
    dispatch_mux.Lock()
    defer dispatch_mux.Unlock()
    
    de, success := dispatch[s.Query]
    if !success {
        fmt.Println("Internal error detected in Subscription.Push.")
        return
    }
    
    var result *ResultDiff = NewResultDiff()
    result.Plus = *de.Cache
    result.Transmit(s.ResponseChannel, s.id)
}

