package config

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Config struct {
    Modules []ModuleConfig `json:"modules"`
}

type ModuleConfig struct {
    Type string `json:"type"`
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////// interface functions

func Load (filename string) *Config {
    var config Config
    
    data, err := ioutil.ReadFile(filename)
    if err!=nil {
        fmt.Println("Unable to load config file:", err)
    }
    
    err = json.Unmarshal(data, &config)
    if err!=nil {
        fmt.Println("Unable to unmarshal config file:", err)
    }
    
    for _, moduleconf := range config.Modules {
        fmt.Println("Found config for module", moduleconf.Type)
    }
    
    return &config
}

