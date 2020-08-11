package main

import (
    "fmt"
    "os"
    
    "innose2019-rdf-server/sparql"
)

func main () {
    var lexer = sparql.NewLexer(true)
    
    var input string = "SELECT ?var WHERE {}"
    
    line, err := sparql.Parse(lexer, input)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    
    fmt.Println(line)
}
