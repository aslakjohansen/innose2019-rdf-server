package command

import (
    "fmt"
    
    "innose2019-rdf-server/logic"
)

func Time (result_channel chan string) {
    var value float64
    var success bool
    value, success = logic.Time()
    
    var result string = ""
    result += "{\n"
    result += fmt.Sprintf("    \"success\": %t,\n", success)
    result += fmt.Sprintf("    \"time\": %f\n", value)
    result += "}\n"
    
    result_channel <- result
}

