package main

import (
    "os"
    "fmt"
    
    "innose2019-rdf-server/config"
    "innose2019-rdf-server/logic"
    "innose2019-rdf-server/transport"
    "innose2019-rdf-server/data/live/mqtt"
)

var config_lut map[string]config.ConfigHander = map[string]config.ConfigHander {
    "logic":          logic.Init,
    "transport":      transport.Init,
    "data/live/mqtt": mqtt.Init,
}

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////// main

func main () {
    // guard: command line arguments
    if (len(os.Args) != 2) {
        fmt.Println("Syntax: "+os.Args[0]+" CONFIG_FILE")
        fmt.Println("        "+os.Args[0]+" ../etc/default_config.json")
        os.Exit(1)
    }
    var config_filename = os.Args[1]
    
    config.Load(config_lut, config_filename)
    
    select{} // block forever
    
    transport.Finalize()
    logic.Finalize()
}

