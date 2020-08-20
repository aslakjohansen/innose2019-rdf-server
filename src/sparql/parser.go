package sparql

import (
    "fmt"
    "github.com/timtadh/lexmachine"
)

func Parse (lexer *lexmachine.Lexer, input string) (line *Node, err error) {
    defer func() {
        if e := recover(); e != nil {
            fmt.Println("e =", e)
            switch e.(type) {
            case error:
                err = e.(error)
                line = nil
            default:
                panic(e)
            }
        }
    }()
    
    scanner, err := newGoLex(lexer, []byte(input))
    if err != nil {
        return nil, err
    }
    
    fmt.Println("scanner =", scanner)
    fmt.Println("scanner.line =", scanner.line)
    yyParse(scanner)
    fmt.Println("scanner.line =", scanner.line)
    return scanner.line, nil
}

