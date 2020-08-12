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

%% /* The grammar follows.  */

SelectStatement
    : SELECT VarList WHERE LBRACE RBRACE {
        varlist := $2.ast
        varlist.CollapseChildList()
        node := NewNode("select", $1.token)
        node.AddChild(varlist)
        yylex.(*golex).line = node
      }
    ;

VarList
    : Var VarList {
        $$.ast = NewNode("list", $1.token)
        $$.ast.AddChild($1.ast)
        $$.ast.AddChild($2.ast)
      }
    | Var {
        $$.ast = NewNode("list", $1.token)
        $$.ast.AddChild(NewNode("var", $1.token))
      }
    ;

Var
    : VAR {
        $$.ast = NewNode("var", $1.token)
      }
    ;

;
%%
