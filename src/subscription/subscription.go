package subscription

import (
    "fmt"
)

type Subscription struct {
    ResponseChannel chan []byte
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////// interface functions

func NewSubscription (response_channel chan []byte) *Subscription {
    var s Subscription
    
    s.ResponseChannel = response_channel
    
    return &s
}

func (s *Subscription) Destroy () {
}

func (s *Subscription) String () string {
    var result string = ""
    
    result += fmt.Sprintf("subscription")
    
    return result
}

