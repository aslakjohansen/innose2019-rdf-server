package command

import (
    "fmt"
    
    "innose2019-rdf-server/logic"
)

func Time (indent string) string {
    var value float64
    var success bool
    value, success = logic.Time()
    
    var response string = ""
    response += fmt.Sprintf("%s{\n", indent)
    response += fmt.Sprintf("%s    \"success\": %t,\n", indent, success)
    response += fmt.Sprintf("%s    \"time\": %f\n", indent, value)
    response += fmt.Sprintf("%s}\n", indent)
    
    return response
}

func Store (indent string, model_dir *string) string {
    var success bool
    var result string
    result, success = logic.Store(*model_dir)
    
    var response string = ""
    response += fmt.Sprintf("%s{\n", indent)
    response += fmt.Sprintf("%s    \"success\": %t,\n", indent, success)
    response += fmt.Sprintf("%s    \"filename\": %s\n", indent, result)
    response += fmt.Sprintf("%s}\n", indent)
    
    return response
}

func Namespaces (indent string) string {
    var success bool
    var result map[string]string
    result, success = logic.Namespaces()
    
    var response string = ""
    response += fmt.Sprintf("%s{\n", indent)
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
    response += fmt.Sprintf("%s}\n", indent)
    
    return response
}
