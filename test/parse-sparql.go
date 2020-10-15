package main

import (
    "fmt"
    
    "innose2019-rdf-server/sparql"
)

func main () {
    var err   error
    var node *sparql.Node
    var line  string
    
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
      "PREFIX a:<http://b> DATA ?var1 ?var3 SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }",
      "PREFIX a:<http://b> DATA ?var1 ?var3 UNITS mod:temp->unit:degc mod:dist->unit:m SELECT ?var1 ?var2 ?var3 WHERE { ?var1 ?var2 ?var3 . }",
      "SELECT ?var1 ?var2 WHERE { ?var1 ?var2 ns:name . }",
      "DATA ?temp SELECT ?temp WHERE { ?temp rdf:type/brick:subClassOf* gfb:Water_Temperature_Sensor . }",
    }
    
    for _, input := range inputs {
        fmt.Println("Case:", input)
        
        // token sequence
        tokens, _ := sparql.Tokens(lexer, []byte(input))
        fmt.Println("[TOKENS]")
        for _, token := range tokens {
            fmt.Println(" -", token)
        }
        
        // parse tree
        node, err = sparql.Parse(lexer, input)
        if err != nil {
            fmt.Println("[PARSE] Error parsing:", err)
        } else {
            fmt.Println("[PARSE]", node)
        }
        
        // normalized query
        if err==nil {
            line, err = node.Normalize("")
            if err != nil {
                fmt.Println("[NORM] Normalized error:", err)
            } else {
                fmt.Println("[NORM]")
                fmt.Print(line)
            }
        }
        
        // sparqlified query
        if err==nil {
            line, err = node.Resparql("")
            if err != nil {
                fmt.Println("[RESPARQL] Sparqlification error:", err)
            } else {
                fmt.Println("[RESPARQL]")
                fmt.Print(line)
            }
        }
        
        fmt.Println("")
    }
}
