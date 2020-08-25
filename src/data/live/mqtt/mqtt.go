package mqtt

import (
    "fmt"
    "time"
    "os"
    "log"
    "github.com/eclipse/paho.mqtt.golang"
)

var (
    client_id   string = "rdf-server"
    brokers   []string = []string{
        "tcp://localhost:1883",
    }
    sub_pattern string = "#"
    
    c mqtt.Client
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("TOPIC: %s\n", msg.Topic())
    fmt.Printf("MSG: %s\n", msg.Payload())
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
}

func Finalize () {
    c.Disconnect(250)
}
