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
      "PREFIX a : <http://b> SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ?var3 . }",
      "PREFIX a:<http://b> SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ?var3 . }",
      "PREFIX a:<http://b> PREFIX c:<http://d> SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ?var3 . }",
      "SELECT ?var1 ?var2 ?var3 WHERE { a:b ?var2 ?var3 . }",
      "DATA ?var1 ?var3 SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }",
      "PREFIX a:<http://b> DATA ?var1 ?var3 SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }",
    }
    
    for _, input := range inputs {
        fmt.Println("Case:", input)
        
        tokens, _ := sparql.Tokens(lexer, []byte(input))
        fmt.Println("[TOKENS]")
        for _, token := range tokens {
            fmt.Println(" -", token)
        }
        
        line, err := sparql.Parse(lexer, input)
        if err != nil {
            fmt.Println("[PARSE] Error parsing:", err)
        } else {
            fmt.Println("[PARSE]", line)
        }
        
        fmt.Println("")
    }
}
