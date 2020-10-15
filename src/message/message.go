package transport

type Message struct {
    Id      string `json:"id"`
    Type    string `json:"type"`
    Success bool   `json:"success"`
}

type MessageError struct {
    Message
    Error string `json:"error"`
}

type MessageTime struct {
    Message
    Value float64 `json:"time"`
}

type MessageStore struct {
    Message
    Filename string `json:"filename"`
}

type MessageNamespaces struct {
    Message
    Namespaces map[string]string `json:"namespaces"`
}

type MessageQuery struct {
    Message
    ResultSet [][]string `json:"resultset"`
}

type MessageUpdate struct {
    Message
}

type MessageInspect struct {
    Message
    Tokens struct {
        Success   bool   `json:"success"`
        Value   []string `json:"value"`
    } `json:"tokens"`
    Parse struct {
        Success bool   `json:"success"`
        Value   string `json:"value"`
    } `json:"parse"`
    Normalize struct {
        Success bool   `json:"success"`
        Value   string `json:"value"`
    } `json:"normalize"`
    Resparql struct {
        Success bool   `json:"success"`
        Value   string `json:"value"`
    } `json:"resparql"`
}

type MessageSubscribe struct {
    Message
}

type MessageUnsubscribe struct {
    Message
}

type MessageSubscriptions struct {
    Message
    Value []string `json:"value"`
}

type MessageResultSet struct {
    Message
    Plus  [][]string `json:"+"`
    Minus [][]string `json:"-"`
}

