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
        yylex.(*golex).line = NewNode("select", $1.token).AddKid($2.ast)
      }
    ;

VarList
    : Var VarList {
        $$.ast = NewNode("varlist", $1.token).AddKid($1.ast).AddKid($2.ast)
      }
    | Var {
        $$.ast = NewNode("var", $1.token)
      }
    ;

Var
    : VAR {
        $$.ast = NewNode("var", $1.token)
      }
    ;

;
%%
