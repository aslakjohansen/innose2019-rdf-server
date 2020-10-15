package mqtt

import (
    "fmt"
    "time"
    "os"
    "log"
    "encoding/json"
    
    "github.com/eclipse/paho.mqtt.golang"
    
    "innose2019-rdf-server/data/reading"
    "innose2019-rdf-server/data/dispatch"
    "innose2019-rdf-server/config"
)

type LoggingConfig struct {
    Debug bool `json:"debug"`
    Error bool `json:"error"`
}

type MqttModuleConfig struct {
    config.ModuleConfig
    ClientID              string        `json:"client-d"`
    Brokers             []string        `json:"brokers"`
    SubscriptionPattern   string        `json:"subscription"`
    Logging               LoggingConfig `json:"logging"`
}

var (
    cfg         MqttModuleConfig;
    c           mqtt.Client
    dispatcher *dispatch.Dispatcher
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    var reading *reading.Reading = reading.NewFromJSON(msg.Payload())
    if reading==nil {
        fmt.Println(fmt.Sprintf("Error: Unable to unmarchal reading received over MQTT topic '%s': '%s'", msg.Topic(), msg.Payload()))
        return
    }
    
    dispatcher.Dispatch(msg.Topic(), reading)
}

func Init (configraw *json.RawMessage) {
    // parse config
    err := json.Unmarshal(*configraw, &cfg)
    if err!=nil {
        fmt.Println("Unable to unmarshal config for module 'data/live/mqtt':", err)
    }
    
    if cfg.Logging.Debug { mqtt.DEBUG = log.New(os.Stdout, "", 0) }
    if cfg.Logging.Error { mqtt.ERROR = log.New(os.Stdout, "", 0) }
    
    // configure options
    opts := mqtt.NewClientOptions()
    for _, broker := range cfg.Brokers {
        opts.AddBroker(broker)
    }
    opts.SetClientID(cfg.ClientID)
    opts.SetKeepAlive(60 * time.Second)
    opts.SetPingTimeout(1 * time.Second)
    
    // construct client
    c = mqtt.NewClient(opts)
    if token := c.Connect(); token.Wait() && token.Error()!=nil {
        fmt.Println("Error: MQTT connection failed:", token.Error())
    }
    
    // subscribe to everything
    if token := c.Subscribe(cfg.SubscriptionPattern, 1, f); token.Wait() && token.Error()!=nil {
        fmt.Println("Error: MQTT subscription failed:", token.Error())
    }
    
    dispatcher = dispatch.GetDispatcher()
    channel := dispatcher.Register("test", make(chan reading.Reading))
    
    go func (channel chan reading.Reading) {
        for {
            var r reading.Reading = <- channel
            fmt.Println(r)
        }
    }(channel)
}

func Finalize () {
    c.Disconnect(250)
}
