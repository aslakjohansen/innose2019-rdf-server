package transport

import (
    "fmt"
    "encoding/json"
    
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
///////////////////////////////////////////////////////////////// command types

type TimeEntry struct {
    Entry
}
func (e *TimeEntry) Handle (s *session.Session) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonTime("    "))
    response += fmt.Sprintf("}\n")
    
    s.ResponseChannel <- []byte(response)
}

type StoreEntry struct {
    Entry
}
func (e *StoreEntry) Handle (s *session.Session) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonStore("    ", model_dir))
    response += fmt.Sprintf("}\n")
    
    s.ResponseChannel <- []byte(response)
}

type NamespacesEntry struct {
    Entry
}
func (e *NamespacesEntry) Handle (s *session.Session) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonNamespaces("    "))
    response += fmt.Sprintf("}\n")
    
    s.ResponseChannel <- []byte(response)
}

type QueryEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *QueryEntry) Handle (s *session.Session) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonQuery("    ", e.Query))
    response += fmt.Sprintf("}\n")
    
    s.ResponseChannel <- []byte(response)
}

type UpdateEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *UpdateEntry) Handle (s *session.Session) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonUpdate("    ", e.Query))
    response += fmt.Sprintf("}\n")
    
    s.ResponseChannel <- []byte(response)
}

type InspectEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *InspectEntry) Handle (s *session.Session) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonInspect("    ", e.Query))
    response += fmt.Sprintf("}\n")
    
    s.ResponseChannel <- []byte(response)
}

type SubscribeEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *SubscribeEntry) Handle (s *session.Session) {
    var sub *subscription.Subscription = subscription.NewSubscription(s.ResponseChannel)
    s.AddSubscription(e.Identifier, sub)
    
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": \"%s\"\n", "subscribed")
    response += fmt.Sprintf("}\n")
    
    s.ResponseChannel <- []byte(response)
}

type UnsubscribeEntry struct {
    Entry
    Subscription string `json:"subscription"`
}
func (e *UnsubscribeEntry) Handle (s *session.Session) {
    s.RemoveSubscription(e.Subscription)
    
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": \"%s\"\n", "unsubscribed")
    response += fmt.Sprintf("}\n")
    
    s.ResponseChannel <- []byte(response)
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
