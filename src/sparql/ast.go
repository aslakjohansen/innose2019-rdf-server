package sparql

import (
    "fmt"
    "strings"
    "errors"
    
    "github.com/timtadh/lexmachine"
)

type Node struct {
    Name     string
    Token    *lexmachine.Token
    Children []*Node
}

func NewNode (name string, token *lexmachine.Token) *Node {
    return &Node{
        Name:  name,
        Token: token,
    }
}

func (n *Node) AddChild (child *Node) *Node {
    n.Children = append(n.Children, child)
    return n
}

func (n *Node) PrependChild (child *Node) *Node {
    children := append(make([]*Node, 0, cap(n.Children)+1), child)
    n.Children = append(children, n.Children...)
    return n
}

func (n *Node) CollapseChildList () *Node {
    var current *Node
    
    // find length
    var length int  = 0
    current = n
    for current!=nil {
        length++
        if len(current.Children)>1 {
            current = current.Children[1]
        } else {
            current = nil
        }
    }
    
    // allocate
    var result []*Node = make([]*Node, length)
    
    // copy
    var i int = 0
    current = n
    for current!=nil {
        if current.Name=="list" {
            result[i] = current.Children[0]
        } else {
            result[i] = current
        }
        i++
        if len(current.Children)>1 {
            current = current.Children[1]
        } else {
            current = nil
        }
    }
    
    // store
    n.Children = result
    
    return n
}

func (n *Node) Resparql (indent string) (string, error) {
    var result  string = ""
    var err     error  = nil
    
    switch n.Name {
    case "query":
        clone := *n // TODO: Is that really a clone?
        clone.Children[1].Children = make([]*Node, 0) // data
        clone.Children[2].Children = make([]*Node, 0) // units
        result, err = clone.Normalize(indent)
    case "select":
        result, err = n.Normalize(indent)
    default:
        err = errors.New(fmt.Sprintf("No case handler defined for sparqlifying a node with name \"%s\".", n.Name))
    }
    return result, err
}

func (n *Node) GetDataIndices () []int {
    switch n.Name {
    case "query":
        var result = make([]int, len(n.Children[1].Children))
        for i, dataitem := range n.Children[1].Children {
            for j, selectitem := range n.Children[3].Children[0].Children {
                if dataitem.String()==selectitem.String() {
                    result[i] = j
                    break
                }
            }
        }
        return result
    case "select":
    default:
        fmt.Println(fmt.Sprintf("No case handler defined for getting the indices of a node with name \"%s\".", n.Name))
    }
    return make([]int, 0)
}

func (n *Node) Normalize (indent string) (string, error) {
    var result  string = ""
    var cresult string
    var err     error  = nil
    
    switch n.Name {
    case "query":
        // prefix
        if len(n.Children[0].Children)>0 {
            for _, child := range n.Children[0].Children {
                cresult, err = child.Normalize("")
                if err!=nil {
                    break
                }
                result += cresult
            }
            result += fmt.Sprintf("%s\n", indent)
        }
        
        // data
        if len(n.Children[1].Children)>0 {
            result += fmt.Sprintf("%sDATA\n", indent)
            for _, child := range n.Children[1].Children {
                cresult, err = child.Normalize("")
                if err!=nil {
                    break
                }
                result += fmt.Sprintf("%s    %s\n", indent, cresult)
            }
        }
        
        // units
        if len(n.Children[2].Children)>0 {
            result += fmt.Sprintf("%sUNITS\n", indent)
            for _, child := range n.Children[2].Children {
                cresult, err = child.Normalize("")
                if err!=nil {
                    break
                }
                result += fmt.Sprintf("%s    %s", indent, cresult)
            }
        }
        
        // select
        cresult, err = n.Children[3].Normalize("")
        if err!=nil {
            break
        }
        result += cresult
    case "select":
        // select
        result += fmt.Sprintf("%sSELECT", indent)
        for _, child := range n.Children[0].Children {
            cresult, err = child.Normalize("")
            if err!=nil {
                break
            }
            result += " "+cresult
        }
        result += fmt.Sprintf("\n")
        
        // where
        result += fmt.Sprintf("%sWHERE {\n", indent)
        for _, child := range n.Children[1].Children {
            cresult, err = child.Normalize("    ")
            if err!=nil {
                break
            }
            result += cresult
        }
        result += fmt.Sprintf("%s}\n", indent)
    case "list":
        result = fmt.Sprintf("'%s contents missing'\n", n.Name)
    case "prefix":
        var id string
        id, err = n.Children[0].Normalize("")
        if err!=nil {
            break
        }
        
        var uri string
        uri, err = n.Children[1].Normalize("")
        if err!=nil {
            break
        }
        
        result = fmt.Sprintf("%sPREFIX %s: <%s>\n", indent, id, uri)
    case "mapping":
        var key string
        key, err = n.Children[0].Normalize("")
        if err!=nil {
            break
        }
        
        var value string
        value, err = n.Children[1].Normalize("")
        if err!=nil {
            break
        }
        
        result = fmt.Sprintf("%s%s -> %s\n", indent, key, value)
    case "var":
        result = string(n.Token.Lexeme)
    case "restriction":
        var sub string
        sub, err = n.Children[0].Normalize("")
        if err!=nil {
            break
        }
        
        var pred string
        pred, err = n.Children[1].Normalize("")
        if err!=nil {
            break
        }
        
        var objz string = ":-("
        objz, err = n.Children[2].Normalize("")
        if err!=nil {
            break
        }
        
        result = fmt.Sprintf("%s%s %s %s .\n", indent, sub, pred, objz)
    case "union":
        var first string
        first, err = n.Children[0].Normalize(indent+"    ")
        if err!=nil {
            break
        }
        
        var second string
        second, err = n.Children[1].Normalize(indent+"    ")
        if err!=nil {
            break
        }
        
        result += fmt.Sprintf("%s{\n", indent)
        result += first
        result += fmt.Sprintf("%s} UNION {\n", indent)
        result += second
        result += fmt.Sprintf("%s} .\n", indent)
    case "uri":
        result = string(n.Token.Lexeme)
    case "id":
        result = string(n.Token.Lexeme)
    case "prefixed":
        var namespace string
        namespace, err = n.Children[0].Normalize("")
        if err!=nil {
            break
        }
        
        var name string
        name, err = n.Children[1].Normalize("")
        if err!=nil {
            break
        }
        
        result = fmt.Sprintf("%s:%s", namespace, name)
    case "string":
        result = string(n.Token.Lexeme)
    case "sequence":
        var first string
        first, err = n.Children[0].Normalize("")
        if err!=nil {
            break
        }
        
        var second string
        second, err = n.Children[1].Normalize("")
        if err!=nil {
            break
        }
        
        result = fmt.Sprintf("%s/%s", first, second)
    case "choice":
        var first string
        first, err = n.Children[0].Normalize("")
        if err!=nil {
            break
        }
        
        var second string
        second, err = n.Children[1].Normalize("")
        if err!=nil {
            break
        }
        
        result = fmt.Sprintf("(%s|%s)", first, second)
    case "one-or-more":
        var subpath string
        subpath, err = n.Children[0].Normalize("")
        if err!=nil {
            break
        }
        result = fmt.Sprintf("%s+", subpath)
    case "zero-or-more":
        var subpath string
        subpath, err = n.Children[0].Normalize("")
        if err!=nil {
            break
        }
        result = fmt.Sprintf("%s*", subpath)
    default:
        err = errors.New(fmt.Sprintf("No case handler defined for normalizing a node with name \"%s\".", n.Name))
    }
    
    return result, err
}

func (n *Node) String () string {
    parts := make([]string, 0, len(n.Children))
    parts = append(parts, n.Name)
    if n.Token != nil && string(n.Token.Lexeme) != n.Name {
        parts = append(parts, fmt.Sprintf("%q", string(n.Token.Lexeme)))
    }
    for _, k := range n.Children {
        parts = append(parts, k.String())
    }
    if len(parts) > 1 {
        return fmt.Sprintf("(%v)", strings.Join(parts, " "))
    }
    return strings.Join(parts, " ")
}

