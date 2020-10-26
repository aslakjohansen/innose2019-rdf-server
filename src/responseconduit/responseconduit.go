package responseconduit

import (
    // "fmt"
    "sync"
)

type ResponseConduit struct {
    Channel chan interface{} // channel we want to be able to close cleanly
    cancels []chan int       // channels for reporting that the channel has reached its end of life
    done    sync.WaitGroup   // waitgroup for making sure that no future values are sent to channel
    mutex   sync.Mutex       // ensure consistency between done and cancels
}

func NewResponseConduit () *ResponseConduit {
    var rc ResponseConduit
    
    rc.Channel =  make(chan interface{})
    
    rc.Hello(nil)
    
    return &rc
}

// declare an interest
func (rc *ResponseConduit) Hello (channel chan int) { // TODO: Add bool return to resolve potential race condition by indicating that conduit is being closed
    
    rc.mutex.Lock()
    rc.done.Add(1)
    if channel!=nil {
        rc.cancels = append(rc.cancels, channel)
    }
    rc.mutex.Unlock()
}

// signal that you are producing no more data
func (rc *ResponseConduit) Goodbye () {
    rc.done.Done()
}

// request goodbyes
func (rc *ResponseConduit) Finalize () {
    // drain channel
    go func () {
        for range rc.Channel {}
    }()
    
    // close channel when ready
    go func () {
        rc.done.Wait()
        close(rc.Channel)
    }()
    
    rc.mutex.Lock()
    rc.Goodbye()
    
    for _, channel := range rc.cancels {
        channel <- 0
    }
    rc.mutex.Unlock()
}

