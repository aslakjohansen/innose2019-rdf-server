package command

import (
    "fmt"
    
    "innose2019-rdf-server/logic"
)

func Time (result_channel chan []byte) {
    var value float64
    var success bool
    fmt.Println(" 1")
    value, success = logic.Time()
    fmt.Println(" 2")
    
    var result string = ""
    result += "{\n"
    result += fmt.Sprintf("    \"success\": %t,\n", success)
    result += fmt.Sprintf("    \"time\": %f\n", value)
    result += "}\n"
    fmt.Println(" 3", result)
    
    result_channel <- []byte(result)
    fmt.Println(" 4")
}

