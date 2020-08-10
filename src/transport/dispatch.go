package transport

import (
    "fmt"
    "encoding/json"
    
    "innose2019-rdf-server/command"
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
    fmt.Println("dispatch:time")
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\":\n")
    response += fmt.Sprintf("%s", command.Time("        "))
    response += fmt.Sprintf("}\n")
    
    fmt.Println("dispatch:time:before")
    response_channel <- []byte(response)
    fmt.Println("dispatch:time:after")
}

type StoreEntry struct {
    Entry
}
func (e *StoreEntry) Handle (response_channel chan []byte) {
    fmt.Println("dispatch:store")
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", e.Identifier)
    response += fmt.Sprintf("    \"response\":\n")
    response += fmt.Sprintf("%s", command.Store("        ", model_dir))
    response += fmt.Sprintf("}\n")
    
    fmt.Println("dispatch:store:before")
    response_channel <- []byte(response)
    fmt.Println("dispatch:store:before")
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////// main dispatcher

func Dispatch (input []byte, response_channel chan []byte) {
    fmt.Println(string(input))
    
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
    }
}
