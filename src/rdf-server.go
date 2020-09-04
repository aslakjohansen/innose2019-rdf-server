package main

import (
    "os"
    "fmt"
    
    "innose2019-rdf-server/config"
    "innose2019-rdf-server/logic"
    "innose2019-rdf-server/transport"
    "innose2019-rdf-server/data/live/mqtt"
)

var (
    model_dir    string = "../var/model"
    ontology_dir string = "../var/ontologies"
)

var config_lut map[string]config.ConfigHander = map[string]config.ConfigHander {
    "logic":     logic.Init,
    "transport": transport.Init,
}

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////// main

func main () {
    // guard: command line arguments
    if (len(os.Args) != 5) {
        fmt.Println("Syntax: "+os.Args[0]+" INTERFACE PORT MODEL_DIR ONTOLOGY_DIR")
        fmt.Println("        "+os.Args[0]+" 0.0.0.0 8001 ../var/model ../var/ontologies")
        os.Exit(1)
    }
        model_dir    = os.Args[3]
        ontology_dir = os.Args[4]
    
    config.Load(config_lut, "../etc/default_config.json")
    
    mqtt.Init()
    
    select{} // block forever
    
    transport.Finalize()
    logic.Finalize()
}

