package command

import (
    "fmt"
    
    "innose2019-rdf-server/logic"
)

//func Time (result_channel chan string, indent string) string {
func Time (indent string) string {
    var value float64
    var success bool
    value, success = logic.Time()
    
    var result string = ""
    result += fmt.Sprintf("%s{\n", indent)
    result += fmt.Sprintf("%s    \"success\": %t,\n", indent, success)
    result += fmt.Sprintf("%s    \"time\": %f\n", indent, value)
    result += fmt.Sprintf("%s}\n", indent)
    
    return result
//    result_channel <- result
}

