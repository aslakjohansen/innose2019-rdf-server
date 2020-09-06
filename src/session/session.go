package session

type Session struct {
    ResponseChannel chan []byte
}

///////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////// interface functions

func NewSession (response_channel chan []byte) *Session {
    var s Session
    
    s.ResponseChannel = response_channel
    
    return &s
}

func (s *Session) Destroy () {
    close(s.ResponseChannel)
}
