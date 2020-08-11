package sparql

import (
    "github.com/timtadh/lexmachine"
)

func Parse (lexer *lexmachine.Lexer, input string) (line *Node, err error) {
    defer func() {
        if e := recover(); e != nil {
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
    
    yyParse(scanner)
    return scanner.line, nil
}

