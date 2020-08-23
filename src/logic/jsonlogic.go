package logic

import (
    "fmt"
    "encoding/json"
    
    "innose2019-rdf-server/sparql"
)

func JsonTime (indent string) string {
    var value float64
    var success bool
    value, success = Time()
    
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("%s    \"success\": %t,\n", indent, success)
    response += fmt.Sprintf("%s    \"time\": %f\n", indent, value)
    response += fmt.Sprintf("%s}", indent)
    
    return response
}

func JsonStore (indent string, model_dir *string) string {
    var success bool
    var result string
    result, success = Store(*model_dir)
    
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("%s    \"success\": %t,\n", indent, success)
    response += fmt.Sprintf("%s    \"filename\": %s\n", indent, result)
    response += fmt.Sprintf("%s}", indent)
    
    return response
}

func JsonNamespaces (indent string) string {
    var success bool
    var result map[string]string
    result, success = Namespaces()
    
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("%s    \"success\": %t,\n", indent, success)
    response += fmt.Sprintf("%s    \"namespaces\": {\n", indent)
    var i    int = 0
    var last int = len(result)-1
    for key, value := range result {
        if i==last {
            response += fmt.Sprintf("%s        \"%s\": \"%s\"\n", indent, key, value)
        } else {
            response += fmt.Sprintf("%s        \"%s\": \"%s\",\n", indent, key, value)
        }
        i++
    }
    response += fmt.Sprintf("%s    }\n", indent)
    response += fmt.Sprintf("%s}", indent)
    
    return response
}

func JsonQuery (indent string, query string) string {
    var success bool
    var result [][]string
    result, success = Query(query)
    
    var response string = ""
    
    if success==false {
        response += fmt.Sprintf("{\n")
        response += fmt.Sprintf("%s    \"success\": false,\n", indent)
        response += fmt.Sprintf("%s    \"error\": \"unable to evaluate query\"\n", indent)
        response += fmt.Sprintf("%s}", indent)
        return response
    }
    
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("%s    \"success\": %t,\n", indent, success)
    response += fmt.Sprintf("%s    \"resultset\": [\n", indent)
    for r:=0 ; r<len(result) ; r++ {
        response += fmt.Sprintf("%s        [\n", indent)
        
        row := result[r]
        for c:=0 ; c<len(row) ; c++ {
            cell := row[c]
            if c==len(row)-1 {
                response += fmt.Sprintf("%s            \"%s\"\n", indent, cell)
            } else {
                response += fmt.Sprintf("%s            \"%s\",\n", indent, cell)
            }
        }
        
        if r==len(result)-1 {
            response += fmt.Sprintf("%s        ]\n", indent)
        } else {
            response += fmt.Sprintf("%s        ],\n", indent)
        }
    }
    response += fmt.Sprintf("%s    ]\n", indent)
    response += fmt.Sprintf("%s}", indent)
    
    return response
}

func JsonUpdate (indent string, query string) string {
    var success bool
    _, success = Update(query)
    
    var response string = ""
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("%s    \"success\": %t\n", indent, success)
    response += fmt.Sprintf("%s}", indent)
    
    return response
}

func JsonInspect (indent string, query string) string {
    var response string = ""
    
    // tokenize
    tokens, err := sparql.Tokens(sparql.NewLexer(true), []byte(query))
    if err!=nil {
        s, _ := json.Marshal(fmt.Sprint(err))
        response += fmt.Sprintf("{\n")
        response += fmt.Sprintf("%s    \"success\": false,\n", indent)
        response += fmt.Sprintf("%s    \"error\": {\n", indent)
        response += fmt.Sprintf("%s        \"type\": \"lex\",\n", indent)
        response += fmt.Sprintf("%s        \"data\": %s\n", indent, s)
        response += fmt.Sprintf("%s    }\n", indent)
        response += fmt.Sprintf("%s}", indent)
        return response
    }
    
    // parse
    parse_data, err := sparql.Parse(sparql.NewLexer(true), query)
    if err!=nil {
        s, _ := json.Marshal(fmt.Sprint(err))
        response += fmt.Sprintf("{\n")
        response += fmt.Sprintf("%s    \"success\": false,\n", indent)
        response += fmt.Sprintf("%s    \"error\": {\n", indent)
        response += fmt.Sprintf("%s        \"type\": \"parse\",\n", indent)
        response += fmt.Sprintf("%s        \"data\": %s\n", indent, s)
        response += fmt.Sprintf("%s    }\n", indent)
        response += fmt.Sprintf("%s}", indent)
        return response
    }
    
    // normalize
    norm_line, err := parse_data.Normalize("")
    if err!=nil {
        s, _ := json.Marshal(fmt.Sprint(err))
        response += fmt.Sprintf("{\n")
        response += fmt.Sprintf("%s    \"success\": false,\n", indent)
        response += fmt.Sprintf("%s    \"error\": {\n", indent)
        response += fmt.Sprintf("%s        \"type\": \"norm\",\n", indent)
        response += fmt.Sprintf("%s        \"data\": %s\n", indent, s)
        response += fmt.Sprintf("%s    }\n", indent)
        response += fmt.Sprintf("%s}", indent)
        return response
    }
    
    // resparql
    resparql_line, err := parse_data.Resparql("")
    if err!=nil {
        s, _ := json.Marshal(fmt.Sprint(err))
        response += fmt.Sprintf("{\n")
        response += fmt.Sprintf("%s    \"success\": false,\n", indent)
        response += fmt.Sprintf("%s    \"error\": {\n", indent)
        response += fmt.Sprintf("%s        \"type\": \"resparql\",\n", indent)
        response += fmt.Sprintf("%s        \"data\": %s\n", indent, s)
        response += fmt.Sprintf("%s    }\n", indent)
        response += fmt.Sprintf("%s}", indent)
        return response
    }
    
    // format result
    s_parse   , _ := json.Marshal(parse_data.String())
    s_norm    , _ := json.Marshal(norm_line)
    s_resparql, _ := json.Marshal(resparql_line)
    response += fmt.Sprintf("{\n")
    response += fmt.Sprintf("%s    \"success\": true,\n", indent)
    response += fmt.Sprintf("%s    \"tokens\": [\n", indent)
    for i, token := range tokens {
        response += fmt.Sprintf("%s        \"%s\"", indent, token)
        if i!=len(tokens)-1 {
            response += fmt.Sprintf(",")
        }
        response += fmt.Sprintf("\n")
    }
    response += fmt.Sprintf("%s    ],\n", indent)
    response += fmt.Sprintf("%s    \"sexp\": %s,\n", indent, s_parse)
    response += fmt.Sprintf("%s    \"norm\": %s,\n", indent, s_norm)
    response += fmt.Sprintf("%s    \"resparql\": %s\n", indent, s_resparql)
    response += fmt.Sprintf("%s}", indent)
    
    return response
}
