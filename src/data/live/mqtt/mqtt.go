package mqtt

import (
    "fmt"
    "time"
    "os"
    "log"
    
    "github.com/eclipse/paho.mqtt.golang"
    
    "innose2019-rdf-server/data/reading"
    "innose2019-rdf-server/data/dispatch"
)

var (
    client_id   string = "rdf-server"
    brokers   []string = []string{
        "tcp://localhost:1883",
    }
    sub_pattern string = "#"
    
    c mqtt.Client
    
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

func Init () {
    mqtt.DEBUG = log.New(os.Stdout, "", 0)
    mqtt.ERROR = log.New(os.Stdout, "", 0)
    
    // configure options
    opts := mqtt.NewClientOptions()
    for _, broker := range brokers {
        opts.AddBroker(broker)
    }
    opts.SetClientID(client_id)
    opts.SetKeepAlive(60 * time.Second)
    opts.SetPingTimeout(1 * time.Second)
    
    // construct client
    c = mqtt.NewClient(opts)
    if token := c.Connect(); token.Wait() && token.Error()!=nil {
        fmt.Println("Error: MQTT connection failed:", token.Error())
    }
    
    // subscribe to everything
    if token := c.Subscribe(sub_pattern, 1, f); token.Wait() && token.Error()!=nil {
        fmt.Println("Error: MQTT subscription failed:", token.Error())
    }
    
    dispatcher = dispatch.NewDispatcher()
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