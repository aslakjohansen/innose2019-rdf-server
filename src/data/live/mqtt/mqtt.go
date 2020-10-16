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
    "innose2019-rdf-server/subscription"
    . "innose2019-rdf-server/responseconduit"
    . "innose2019-rdf-server/message"
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
    data_rc    *ResponseConduit
    data_sub   *subscription.Subscription
    namemap     map[string]string = make(map[string]string)
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    var reading *reading.Reading = reading.NewFromJSON(msg.Payload())
    if reading==nil {
        fmt.Println(fmt.Sprintf("Error: Unable to unmarchal reading received over MQTT topic '%s': '%s'", msg.Topic(), msg.Payload()))
        return
    }
    
    entity, ok := namemap[msg.Topic()] // TODO: Allow multiple entities to share a topic
    if (ok) {
        dispatcher.Dispatch(entity, reading)
    }
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
    
    // subscribe to annotation mapping and maintain local map
    data_rc = NewResponseConduit()
    var query string = `
SELECT ?entity ?broker ?topic ?data
WHERE {
    ?entity dao:hasMqttLiveData ?data .
    ?data   rdf:type dao:MqttLiveData .
    ?data   dao:hasBroker ?broker .
    ?data   dao:hasTopic  ?topic .
}
`
    data_sub = subscription.NewSubscription("dummyid", query, data_rc, nil)
    go func () {
        // var update MessageResultSet
        for updatei := range data_rc.Channel {
            update, ok := updatei.(*MessageResultSet)
            if !ok {
                fmt.Println("Unable to pass update to topic-> entity map")
                continue
            }
            for _, row := range update.Plus {
                namemap[row[2]] = row[0]
            }
            for _, row := range update.Minus {
                delete(namemap, row[2])
            }
            PrintNamemap()
        }
    }()
    data_sub.Push()
}

func Finalize () {
    c.Disconnect(250)
}

func PrintNamemap () {
    fmt.Println("<mqtt:namemap>")
    for key, value := range namemap {
        fmt.Println("    <mapping key=\""+key+"\" value=\""+value+"\" />")
    }
    fmt.Println("</mqtt:namemap>")
}
