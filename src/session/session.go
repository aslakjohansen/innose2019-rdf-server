package session

import (
    "fmt"
    
    "innose2019-rdf-server/subscription"
    . "innose2019-rdf-server/responseconduit"
)

type Session struct {
    // ResponseChannel chan interface{}
    ResponseConduit *ResponseConduit
    Subscriptions    map[string](*subscription.Subscription)
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////// interface functions

func NewSession (response_conduit *ResponseConduit) *Session {
// func NewSession (response_channel chan interface{}) *Session {
    var s Session
    
    // s.ResponseChannel = response_channel
    s.ResponseConduit = response_conduit
    s.Subscriptions   = make(map[string](*subscription.Subscription))
    
    var goodbye_channel chan int = make(chan int)
    response_conduit.Hello(goodbye_channel)
    go func () {
        <- goodbye_channel
        s.Destroy()
        response_conduit.Goodbye() // TODO: is this triggering a race condition
    }()
    
    return &s
}

func (s *Session) Destroy () {
    // close(s.ResponseChannel)
    for id, _ := range s.Subscriptions {
        s.RemoveSubscription(id)
    }
}

func (s *Session) AddSubscription (identifier string, sub *subscription.Subscription) {
    s.Subscriptions[identifier] = sub
}

func (s *Session) RemoveSubscription (identifier string) {
    sub, ok := s.Subscriptions[identifier]
    if ok {
        sub.Destroy()
        delete(s.Subscriptions, identifier)
    }
}

func (s *Session) GetSubscriptionIdentifiers () []string {
    var result []string = make([]string, 0)
    
    for key, _ := range s.Subscriptions {
        result = append(result, key)
    }
    
    return result;
}

func (s *Session) String () string {
    var result string = ""
    
    result += fmt.Sprintf("session [\n")
    for i, sub := range s.Subscriptions {
        result += fmt.Sprintf("  %s: %s\n", i, sub)
    }
    result += fmt.Sprintf("]")
    
    return result
}
