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
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 #\n?var2 ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 #comment\n?var2 ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2) ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2/?var2 ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2|?var3 ?var4 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2|?var3) ?var4 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2+ ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2++ ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2* ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { ?var1 (?var2|?var3)/?var4 ?var5 . }",
      "SELECT ?var1 ?var2 WHERE { { ?var1 ?var1 ?var2 } UNION { ?var1 ?var2 ?var2 } . }",
      "SELECT ?var1 ?var2 WHERE { ?var1 ?var2 < http://www.google.com#test > . }",
      "SELECT ?var1 ?var2 WHERE { ?var1 ?var2 <http://www.google.com#test> . }",
      "SELECT ?var1 ?var2 WHERE { ?var1 ?var2 \"\" . }",
      "SELECT ?var1 ?var2 WHERE { ?var1 ?var2 \"a b\" . }",
      "SELECT ?var1 ?var2 WHERE { ?var1 ?var2 \"a \\\" b\" . }",
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
