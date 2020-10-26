package reading

import (
    "fmt"
    "encoding/json"
)

type Reading struct {
    Timestamp float64 `json:"time"`
    Value     float64 `json:"value"`
}

func NewFromJSON (input []byte) *Reading {
    var reading Reading
    
    err := json.Unmarshal(input, &reading)
    if err!=nil {
        fmt.Println(fmt.Sprintf("Error: Unable to unmarchal reading from '%s': %s", input, err))
        return nil
    }
    
    return &reading
}

func (r Reading) String () string {
    return fmt.Sprintf("{\n    \"time\": %f,\n    \"value\": %f\n}", r.Timestamp, r.Value)
}

