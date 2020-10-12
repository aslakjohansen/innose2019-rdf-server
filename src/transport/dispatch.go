package transport

import (
    "fmt"
    "encoding/json"
    
    "innose2019-rdf-server/sparql"
    "innose2019-rdf-server/logic"
    "innose2019-rdf-server/session"
    "innose2019-rdf-server/subscription"
)

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////// base type

type Entry struct {
    Command string    `json:"command"`
    Identifier string `json:"id"`
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////// helpers

func send_response (channel chan []byte, id string, response string) {
    var message string = ""
    message += fmt.Sprintf("{\n")
    message += fmt.Sprintf("    \"id\": \"%s\",\n", id)
    message += fmt.Sprintf("    \"response\": %s\n", response)
    message += fmt.Sprintf("}\n")
    
    channel <- []byte(message)
}

func send_response_error (channel chan []byte, id string, details string) {
    var message string = ""
    message += fmt.Sprintf("{\n")
    message += fmt.Sprintf("    \"id\": \"%s\",\n", id)
    message += fmt.Sprintf("    \"response\": \"%s\",\n", "error")
    message += fmt.Sprintf("    \"details\": \"%s\"\n", details)
    message += fmt.Sprintf("}\n")
    
    channel <- []byte(message)
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////// command types

type TimeEntry struct {
    Entry
}
func (e *TimeEntry) Handle (s *session.Session) {
    send_response(s.ResponseChannel, e.Identifier, logic.JsonTime("    "))
    // var response string = ""
    // response += fmt.Sprintf("{\n")
    // response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    // response += fmt.Sprintf("    \"response\": %s\n", logic.JsonTime("    "))
    // response += fmt.Sprintf("}\n")
    
    // s.ResponseChannel <- []byte(response)
}

type StoreEntry struct {
    Entry
}
func (e *StoreEntry) Handle (s *session.Session) {
    send_response(s.ResponseChannel, e.Identifier, logic.JsonStore("    ", model_dir))
    // var response string = ""
    // response += fmt.Sprintf("{\n")
    // response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    // response += fmt.Sprintf("    \"response\": %s\n", logic.JsonStore("    ", model_dir))
    // response += fmt.Sprintf("}\n")
    
    // s.ResponseChannel <- []byte(response)
}

type NamespacesEntry struct {
    Entry
}
func (e *NamespacesEntry) Handle (s *session.Session) {
    send_response(s.ResponseChannel, e.Identifier, logic.JsonNamespaces("    "))
    // var response string = ""
    // response += fmt.Sprintf("{\n")
    // response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    // response += fmt.Sprintf("    \"response\": %s\n", logic.JsonNamespaces("    "))
    // response += fmt.Sprintf("}\n")
    
    // s.ResponseChannel <- []byte(response)
}

type QueryEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *QueryEntry) Handle (s *session.Session) {
    send_response(s.ResponseChannel, e.Identifier, logic.JsonQuery("    ", e.Query))
    // var response string = ""
    // response += fmt.Sprintf("{\n")
    // response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    // response += fmt.Sprintf("    \"response\": %s\n", logic.JsonQuery("    ", e.Query))
    // response += fmt.Sprintf("}\n")
    
    // s.ResponseChannel <- []byte(response)
}

type UpdateEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *UpdateEntry) Handle (s *session.Session) {
    var response string = logic.JsonUpdate("    ", e.Query)
    subscription.Update()
    send_response(s.ResponseChannel, e.Identifier, response)
    // var response string = ""
    // response += fmt.Sprintf("{\n")
    // response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    // response += fmt.Sprintf("    \"response\": %s\n", logic.JsonUpdate("    ", e.Query))
    // response += fmt.Sprintf("}\n")
    
    // subscription.Update()
    
    // s.ResponseChannel <- []byte(response)
}

type InspectEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *InspectEntry) Handle (s *session.Session) {
    send_response(s.ResponseChannel, e.Identifier, logic.JsonInspect("    ", e.Query))
    // var response string = ""
    // response += fmt.Sprintf("{\n")
    // response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    // response += fmt.Sprintf("    \"response\": %s\n", logic.JsonInspect("    ", e.Query))
    // response += fmt.Sprintf("}\n")
    
    // s.ResponseChannel <- []byte(response)
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
        send_response_error(s.ResponseChannel, e.Identifier, "[RESPARQL] Parse error:"+err.Error())
        return
    }
    
    // extract pure sparql
    q, err = node.Resparql("")
    if err != nil {
        send_response_error(s.ResponseChannel, e.Identifier, "[RESPARQL] Sparqlification error:"+err.Error())
        return
    }
    
    var sub *subscription.Subscription = subscription.NewSubscription(e.Identifier, q, s.ResponseChannel)
    s.AddSubscription(e.Identifier, sub)
    
    send_response(s.ResponseChannel, e.Identifier, "subscribed")
    
    sub.Push()
    
    // var response string = ""
    // response += fmt.Sprintf("{\n")
    // response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    // response += fmt.Sprintf("    \"response\": \"%s\"\n", "subscribed")
    // response += fmt.Sprintf("}\n")
    
    // s.ResponseChannel <- []byte(response)
}

type UnsubscribeEntry struct {
    Entry
    Subscription string `json:"subscription"`
}
func (e *UnsubscribeEntry) Handle (s *session.Session) {
    subscription.Unsubscribe(e.Identifier)
    send_response(s.ResponseChannel, e.Identifier, "unsubscribed")
    // s.RemoveSubscription(e.Subscription)
    
    // var response string = ""
    // response += fmt.Sprintf("{\n")
    // response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    // response += fmt.Sprintf("    \"response\": \"%s\"\n", "unsubscribed")
    // response += fmt.Sprintf("}\n")
    
    // s.ResponseChannel <- []byte(response)
}

type SubscriptionsEntry struct {
    Entry
}
func (e *SubscriptionsEntry) Handle (s *session.Session) {
    var ids []string = s.GetSubscriptionIdentifiers()
    
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"subscriptions\": [\n")
    for i, key := range ids {
        if i==len(ids)-1 {
            response += fmt.Sprintf("    \"%s\"\n", key)
        } else {
            response += fmt.Sprintf("    \"%s\",\n", key)
        }
    }
    response += fmt.Sprintf("    ]\n")
    response += fmt.Sprintf("}\n")
    
    s.ResponseChannel <- []byte(response)
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
