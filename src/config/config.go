package config

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Config struct {
    Modules []json.RawMessage `json:"modules"`
}

type ModuleConfig struct {
    Type string `json:"type"`
}
type ConfigHander func(x *json.RawMessage)

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////// interface functions

func Load (lut map[string]ConfigHander, filename string) *Config {
    var config Config
    
    data, err := ioutil.ReadFile(filename)
    if err!=nil {
        fmt.Println("Unable to load config file:", err)
    }
    
    err = json.Unmarshal(data, &config)
    if err!=nil {
        fmt.Println("Unable to unmarshal config file:", err)
    }
    
    for i, rawmoduleconf := range config.Modules {
        var moduleconfig ModuleConfig
        
        err = json.Unmarshal(rawmoduleconf, &moduleconfig)
        if err!=nil {
            fmt.Println("Unable to unmarshal module config", i, ":", err)
        }
        
        fmt.Println("Found config for module", moduleconfig.Type)
        handler, ok := lut[moduleconfig.Type]
        if ok {
            handler(&rawmoduleconf)
        } else {
            fmt.Println("Unable to look up handler for module", moduleconfig.Type)
        }
        
    }
    
    return &config
}

