package transport

import (
    "fmt"
    "encoding/json"
    
    "innose2019-rdf-server/logic"
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
func (e *TimeEntry) Handle (response_channel chan []byte) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonTime("    "))
    response += fmt.Sprintf("}\n")
    
    response_channel <- []byte(response)
}

type StoreEntry struct {
    Entry
}
func (e *StoreEntry) Handle (response_channel chan []byte) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonStore("    ", model_dir))
    response += fmt.Sprintf("}\n")
    
    response_channel <- []byte(response)
}

type NamespacesEntry struct {
    Entry
}
func (e *NamespacesEntry) Handle (response_channel chan []byte) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonNamespaces("    "))
    response += fmt.Sprintf("}\n")
    
    response_channel <- []byte(response)
}

type QueryEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *QueryEntry) Handle (response_channel chan []byte) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonQuery("    ", e.Query))
    response += fmt.Sprintf("}\n")
    
    response_channel <- []byte(response)
}

type UpdateEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *UpdateEntry) Handle (response_channel chan []byte) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonUpdate("    ", e.Query))
    response += fmt.Sprintf("}\n")
    
    response_channel <- []byte(response)
}

type InspectEntry struct {
    Entry
    Query string `json:"query"`
}
func (e *InspectEntry) Handle (response_channel chan []byte) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\": %s\n", logic.JsonInspect("    ", e.Query))
    response += fmt.Sprintf("}\n")
    
    response_channel <- []byte(response)
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////// main dispatcher

func Dispatch (input []byte, response_channel chan []byte) {
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
        typed_entry.Handle(response_channel)
    case "store":
        var typed_entry StoreEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(response_channel)
    case "namespaces":
        var typed_entry NamespacesEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(response_channel)
    case "query":
        var typed_entry QueryEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(response_channel)
    case "update":
        var typed_entry UpdateEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(response_channel)
    case "inspect":
        var typed_entry InspectEntry
        err := json.Unmarshal(input, &typed_entry)
        if err!=nil {
            fmt.Println("Error: Unable to do second parsing of ws message:", string(input))
            return
        }
        typed_entry.Handle(response_channel)
    }
}
