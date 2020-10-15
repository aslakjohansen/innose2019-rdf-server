package transport

import (
    "fmt"
    "encoding/json"
    
    "innose2019-rdf-server/sparql"
    "innose2019-rdf-server/logic"
    "innose2019-rdf-server/session"
    "innose2019-rdf-server/subscription"
    . "innose2019-rdf-server/message"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////// base type

type Entry struct {
    Command string    `json:"command"`
    Identifier string `json:"id"`
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////// helpers

func send_response_error (channel chan interface{}, id string, details string) {
    var m MessageError = MessageError{Message{id, "error", false}, details}
    channel <- &m
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////// command types

type TimeEntry struct {
    Entry
}
func (e *TimeEntry) Handle (s *session.Session) {
    var value float64
    var success bool
    fmt.Println("dispatch:TimeEntry.Handle")
    value, success = logic.Time()
    var response MessageTime = MessageTime{Message{e.Identifier, "time", success}, value}
    s.ResponseConduit.Channel <-&response
}

type StoreEntry struct {
    Entry
}
func (e *StoreEntry) Handle (s *session.Session) {
    var success bool
    var filename string
    filename, success = logic.Store(*model_dir)
    var response MessageStore = MessageStore{Message{e.Identifier, "store", success}, filename}
    s.ResponseConduit.Channel <-&response
}

type NamespacesEntry struct {
    Entry
}
func (e *NamespacesEntry) Handle (s *session.Session) {
    var success bool
    var result map[string]string
    result, success = logic.Namespaces()
    var response MessageNamespaces = MessageNamespaces{Message{e.Identifier, "namespaces", success}, result}
    s.ResponseConduit.Channel <-&response
}

type QueryEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *QueryEntry) Handle (s *session.Session) {
    var success bool
    var result [][]string
    
    result, success = logic.Query(e.Query)
    var response MessageQuery = MessageQuery{Message{e.Identifier, "query", success}, result}
    s.ResponseConduit.Channel <-&response
}

type UpdateEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *UpdateEntry) Handle (s *session.Session) {
    var success bool
    _, success = logic.Update(e.Query)
    subscription.Update()
    var response MessageUpdate = MessageUpdate{Message{e.Identifier, "update", success}}
    s.ResponseConduit.Channel <-&response
}

type InspectEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *InspectEntry) Handle (s *session.Session) {
    var m MessageInspect
    m.Id   = e.Identifier
    m.Type = "inspect"
    
    query := e.Query
    
    // tokenize
    tokens, err1 := sparql.Tokens(sparql.NewLexer(true), []byte(query))
    m.Tokens.Success = err1==nil
    m.Tokens.Value   = tokens
    
    // parse
    parse_data, err2 := sparql.Parse(sparql.NewLexer(true), query)
    m.Parse.Success = err2==nil
    m.Parse.Value   = parse_data.String()
    
    // normalize
    norm_line, err3 := parse_data.Normalize("")
    m.Normalize.Success = err3==nil
    m.Normalize.Value   = norm_line
    
    // resparql
    resparql_line, err4 := parse_data.Resparql("")
    m.Resparql.Success = err4==nil
    m.Resparql.Value   = resparql_line
    
    m.Success = err1==nil && err2==nil && err3==nil && err4==nil
    
    s.ResponseConduit.Channel <- &m
}

type SubscribeEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *SubscribeEntry) Handle (s *session.Session) {
    var lexer = sparql.NewLexer(true)
    var node *sparql.Node
    var err   error
    var q     string
    
    // parse tree
    fmt.Println(e.Query)
    node, err = sparql.Parse(lexer, e.Query)
    if err != nil {
        send_response_error(s.ResponseConduit.Channel, e.Identifier, "[RESPARQL] Parse error:"+err.Error())
        return
    }
    
    // extract pure sparql
    q, err = node.Resparql("")
    if err != nil {
        send_response_error(s.ResponseConduit.Channel, e.Identifier, "[RESPARQL] Sparqlification error:"+err.Error())
        return
    }
    
    var sub *subscription.Subscription = subscription.NewSubscription(e.Identifier, q, s.ResponseConduit)
    s.AddSubscription(e.Identifier, sub)
    
    var response MessageSubscribe = MessageSubscribe{Message{e.Identifier, "subscribe", true}}
    s.ResponseConduit.Channel <- &response
    
    sub.Push()
}

type UnsubscribeEntry struct {
    Entry
    Subscription string `json:"subscription"`
}
func (e *UnsubscribeEntry) Handle (s *session.Session) {
    subscription.Unsubscribe(e.Identifier)
    var response MessageUnsubscribe = MessageUnsubscribe{Message{e.Identifier, "unsubscribe", true}}
    s.ResponseConduit.Channel <-&response
}

type SubscriptionsEntry struct {
    Entry
}
func (e *SubscriptionsEntry) Handle (s *session.Session) {
    var ids []string = s.GetSubscriptionIdentifiers()
    var response MessageSubscriptions = MessageSubscriptions{Message{e.Identifier, "subscriptions", true}, ids}
    s.ResponseConduit.Channel <-&response
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////// main dispatcher

func Dispatch (input []byte, s *session.Session) {
    var entry Entry
    err := json.Unmarshal(input, &entry)
    if err!=nil {
        fmt.Println("Error: Unable to do initial parsing of ws message:", string(input))
        return
    }
    
    switch entry.Command {
    case "time":
        var typed_entry TimeEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(s)
    case "store":
        var typed_entry StoreEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(s)
    case "namespaces":
        var typed_entry NamespacesEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(s)
    case "query":
        var typed_entry QueryEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(s)
    case "update":
        var typed_entry UpdateEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(s)
    case "inspect":
        var typed_entry InspectEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(s)
    case "subscribe":
        var typed_entry SubscribeEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(s)
    case "unsubscribe":
        var typed_entry UnsubscribeEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(s)
    case "subscriptions":
        var typed_entry SubscriptionsEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(s)
    }
}
