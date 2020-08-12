%{

package sparql

import (
    "github.com/timtadh/lexmachine"
)

%}

%union{
    token *lexmachine.Token
    ast   *Node
}

%token	VAR
%token	SELECT
%token	WHERE
%token	LBRACE
%token	RBRACE
%token	PERIOD
%token	SLASH
%token	VBAR

%% /* The grammar follows.  */

SelectStatement
    : SELECT VarList WHERE LBRACE RestrictionList RBRACE {
        varlist := $2.ast
        varlist.CollapseChildList()
        reslist := $5.ast
        reslist.CollapseChildList()
        node := NewNode("select", $1.token)
        node.AddChild(varlist)
        node.AddChild(reslist)
        yylex.(*golex).line = node
      }
    ;

VarList
    : Var VarList {
        node := NewNode("list", $1.token)
        node.AddChild($1.ast)
        node.AddChild($2.ast)
        $$.ast = node
      }
    | Var {
        node := NewNode("list", $1.token)
        node.AddChild($1.ast)
        $$.ast = node
      }
    ;

Var
    : VAR {
        $$.ast = NewNode("var", $1.token)
      }
    ;

;

RestrictionList
    : Restriction PERIOD RestrictionList {
        node := NewNode("list", $1.token)
        node.AddChild($1.ast)
        node.AddChild($2.ast)
        $$.ast = node
      }
    | Restriction PERIOD {
        node := NewNode("list", $1.token)
        node.AddChild($1.ast)
        $$.ast = node
      }
    ;

Restriction
    : Entity Path Entity {
        node := NewNode("restriction", $1.token)
        node.AddChild($1.ast)
        node.AddChild($2.ast)
        node.AddChild($3.ast)
        $$.ast = node
      }
    ;

Entity
    : Var {
        $$.ast = NewNode("var", $1.token)
      }
    ;

Path
    : Path SLASH Path {
        node := NewNode("sequence", $2.token)
        node.AddChild($1.ast)
        node.AddChild($3.ast)
        $$.ast = node
      }
    |  Path VBAR Path {
        node := NewNode("choice", $2.token)
        node.AddChild($1.ast)
        node.AddChild($3.ast)
        $$.ast = node
      }
    | Var {
        $$.ast = $1.ast
      }
    ;

%%
