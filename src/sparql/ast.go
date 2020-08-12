package sparql

import (
    "fmt"
    "strings"
    
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

