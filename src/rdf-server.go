package main

import (
    "os"
    "fmt"
    
    "innose2019-rdf-server/logic"
    "innose2019-rdf-server/transport"
)

var (
    model_dir    string = "../var/model"
    ontology_dir string = "../var/ontologies"
)

///////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////// main

func main () {
//    var port int16 = 8001
    
    // guard: command line arguments
    if (len(os.Args) != 5) {
        fmt.Println("Syntax: "+os.Args[0]+" INTERFACE PORT MODEL_DIR ONTOLOGY_DIR")
        fmt.Println("        "+os.Args[0]+" 0.0.0.0 8001 ../var/model ../var/ontologies")
        os.Exit(1)
    }
    var iface string = os.Args[1]
    var port  string = os.Args[2]
        model_dir    = os.Args[3]
        ontology_dir = os.Args[4]
    
    logic.Init(model_dir, ontology_dir)
    transport.Init(iface, port, &model_dir)
    
    select{} // block forever
    
    transport.Finalize()
    logic.Finalize()
}

