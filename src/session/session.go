package session

import (
    "fmt"
    
    "innose2019-rdf-server/subscription"
)

type Session struct {
    ResponseChannel chan []byte
    Subscriptions   map[string](*subscription.Subscription)
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////// interface functions

func NewSession (response_channel chan []byte) *Session {
    var s Session
    
    s.ResponseChannel = response_channel
    s.Subscriptions   = make(map[string](*subscription.Subscription))
    
    return &s
}

func (s *Session) Destroy () {
    close(s.ResponseChannel)
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

func (s *Session) String () string {
    var result string = ""
    
    result += fmt.Sprintf("session [\n")
    for i, sub := range s.Subscriptions {
        result += fmt.Sprintf("  %s: %s\n", i, sub)
    }
    result += fmt.Sprintf("]")
    
    return result
}
