package main

import (
    "fmt"
    
    "innose2019-rdf-server/session"
    "innose2019-rdf-server/subscription"
)

func main () {
    var s *session.Session = session.NewSession()
    
    for i := 0; i < 10; i++ {
        var identifier string = fmt.Sprintf("identifier%d", i)
        var response_channel chan []byte = make(chan []byte)
        var sub = subscription.NewSubscription(response_channel)
        s.AddSubscription(identifier, sub)
        
        fmt.Println(s);
        fmt.Println("")
    }
    
    
}

