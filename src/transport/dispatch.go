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
func (te *TimeEntry) Handle (response_channel chan []byte) {
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("    \"id\": \"%s\",\n", te.Identifier)
    response += fmt.Sprintf("    \"data\":\n")
    response += fmt.Sprintf("%s", command.Time("        "))
    response += fmt.Sprintf("}\n")
    
    response_channel <- []byte(response)
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
    }
}
