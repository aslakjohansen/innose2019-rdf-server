package main

import (
    "fmt"
    
    "innose2019-rdf-server/sparql"
)

func main () {
    var lexer = sparql.NewLexer(true)
    
    var inputs []string = []string{
      "garbage",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2) ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2/?var2 ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2|?var3 ?var4 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2|?var3) ?var4 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2+ ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2++ ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2* ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2|?var3)/?var4 ?var5 . }",
    }
    
    for _, input := range inputs {
        fmt.Println("Case:", input)
        line, err := sparql.Parse(lexer, input)
        if err != nil {
            fmt.Println(">> Error parsing:", err)
        } else {
            fmt.Println(">>", line)
        }
    }
}
