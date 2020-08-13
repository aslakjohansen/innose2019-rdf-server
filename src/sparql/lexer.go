package sparql

import (
    "fmt"
    
    "github.com/timtadh/lexmachine"
    "github.com/timtadh/lexmachine/machines"
)


//var (
//    regex_uri string = "([a-z][a-z0-9+.-]*):(?:\/\/((?:(?=((?:[a-z0-9-._~!$&'()*+,;=:]|%[0-9A-F]{2})*))(\3)@)?(?=(\[[0-9A-F:.]{2,}\]|(?:[a-z0-9-._~!$&'()*+,;=]|%[0-9A-F]{2})*))\5(?::(?=(\d*))\6)?)(\/(?=((?:[a-z0-9-._~!$&'()*+,;=:@\/]|%[0-9A-F]{2})*))\8)?|(\/?(?!\/)(?=((?:[a-z0-9-._~!$&'()*+,;=:@\/]|%[0-9A-F]{2})*))\10)?)(?:\?(?=((?:[a-z0-9-._~!$&'()*+,;=:@\/?]|%[0-9A-F]{2})*))\11)?(?:#(?=((?:[a-z0-9-._~!$&'()*+,;=:@\/?]|%[0-9A-F]{2})*))\12)?" // source: https://snipplr.com/view/6889/regular-expressions-for-uri-validationparsing
//)

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
    lexer.Add([]byte(`[a-zA-Z]+\:\/\/[^ \t\n\r\<\>]+`), token(URI))
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
