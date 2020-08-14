package sparql

import (
    "fmt"
    
    "github.com/timtadh/lexmachine"
    "github.com/timtadh/lexmachine/machines"
)

type golex struct {
    *lexmachine.Scanner
    line *Node
}

func newGoLex (lexer *lexmachine.Lexer, text []byte) (*golex, error) {
    scan, err := lexer.Scanner(text)
    if err != nil {
        return nil, err
    }
    return &golex{Scanner: scan}, nil
}

func Tokens (lexer *lexmachine.Lexer, text []byte) []string {
    var result []string = make([]string, 0)
    
    scanner, err := lexer.Scanner(text)
    if err != nil {
        result = append(result, fmt.Sprint(err))
        return result
    }
    
    for tok, err, eos := scanner.Next(); !eos; tok, err, eos = scanner.Next() {
        if ui, is := err.(*machines.UnconsumedInput); is {
            // skip the error via:
            // scanner.TC = ui.FailTC
            //
            result = append(result, fmt.Sprint(err))
            result = append(result, fmt.Sprint(ui))
            result = append(result, fmt.Sprint(is))
            return result
        } else if err != nil {
            result = append(result, fmt.Sprint(err))
            return result
        }
//        var t *lexmachine.Token = tok.(*lexmachine.Token)
        result = append(result, fmt.Sprint(tok))
//        result = append(result, fmt.Sprintf("TOKEN %d", t.Type))
    }
    
    return result
}

func (g *golex) Lex (lval *yySymType) (tokenType int) {
    s := g.Scanner
    tok, err, eof := s.Next()
    if err != nil {
        g.Error(err.Error())
    } else if eof {
        return -1 // Note: signals EOF to yyParse
    }
    
    lval.token = tok.(*lexmachine.Token)
    return lval.token.Type
}

func token (id int) lexmachine.Action {
    return func (s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
        return s.Token(id, string(m.Bytes), m), nil
    }
}

func skip (*lexmachine.Scanner, *machines.Match) (interface{}, error) {
    return nil, nil
}

func (l *golex) Error (message string) {
    panic(fmt.Errorf(message))
}

func NewLexer (dfa bool) *lexmachine.Lexer {
    var lexer = lexmachine.NewLexer()
    lexer.Add([]byte(`\?[a-zA-Z][a-zA-Z0-9]*`), token(VAR))
    lexer.Add([]byte(`(S|s)(E|e)(L|l)(E|e)(C|c)(T|t)`), token(SELECT))
    lexer.Add([]byte(`(W|w)(H|w)(E|e)(R|r)(E|e)`), token(WHERE))
    lexer.Add([]byte(`(U|u)(N|n)(I|i)(O|o)(N|n)`), token(UNION))
    lexer.Add([]byte(`(P|p)(R|r)(E|e)(F|f)(I|i)(X|x)`), token(PREFIX))
    lexer.Add([]byte(`[a-zA-Z]+\:\/\/[^ \t\n\r\<\>]+`), token(URI))
    lexer.Add([]byte(`[a-zA-Z][a-zA-Z0-9]*`), token(ID))
    lexer.Add([]byte(`\"([^\"\\]|\\.)*\"`), token(STRING))
    lexer.Add([]byte(`\{`), token(LBRACE))
    lexer.Add([]byte(`\}`), token(RBRACE))
    lexer.Add([]byte(`\(`), token(LPAR))
    lexer.Add([]byte(`\)`), token(RPAR))
    lexer.Add([]byte(`\<`), token(LT))
    lexer.Add([]byte(`\>`), token(GT))
    lexer.Add([]byte(`\.`), token(PERIOD))
    lexer.Add([]byte(`\/`), token(SLASH))
    lexer.Add([]byte(`\|`), token(VBAR))
    lexer.Add([]byte(`\+`), token(PLUS))
    lexer.Add([]byte(`\*`), token(ASTERISK))
    lexer.Add([]byte(`\:`), token(COLON))
    lexer.Add([]byte("( |\t|\n|\r)+"), skip) // whitespace
    lexer.Add([]byte("#[^\n\r]*"), skip) // comment
    
    var err error
    if dfa {
        err = lexer.CompileDFA()
    } else {
        err = lexer.CompileNFA()
    }
    if err != nil {
        panic(err)
    }
    return lexer
}
