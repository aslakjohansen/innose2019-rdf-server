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
%token SELECT
%token WHERE
%token UNION
%token DATA
%token UNITS
%token PREFIX

%token VAR
%token URI
%token ID
%token STRING
%token ARROW

%token LBRACE
%token RBRACE
%token LT
%token GT
%token PERIOD
%token SLASH
%token VBAR
%token PLUS
%token ASTERISK
%token COLON
%token LPAR
%token RPAR

%left LPAR SLASH VBAR

%% /* The grammar follows.  */

Query
    : SelectStatement {
        node := NewNode("query", $1.token)
        node.AddChild(NewNode("list", $1.token)) // missing prefix
        node.AddChild(NewNode("list", $1.token)) // missing data
        node.AddChild(NewNode("list", $1.token)) // missing units
        node.AddChild($1.ast) // select
        yylex.(*golex).line = node
      }
    | PrefixList SelectStatement {
        prefixlist := $1.ast
        prefixlist.CollapseChildList()
        node := NewNode("query", $1.token)
        node.AddChild(prefixlist) // prefix
        node.AddChild(NewNode("list", $1.token)) // missing data
        node.AddChild(NewNode("list", $1.token)) // missing units
        node.AddChild($2.ast)     // select
        yylex.(*golex).line = node
      }
    | DATA VarList SelectStatement {
        datalist := $2.ast
        datalist.CollapseChildList()
        node := NewNode("query", $1.token)
        node.AddChild(NewNode("list", $1.token)) // missing prefix
        node.AddChild(datalist) // data
        node.AddChild(NewNode("list", $1.token)) // missing units
        node.AddChild($3.ast)   // select
        yylex.(*golex).line = node
      }
    | PrefixList DATA VarList SelectStatement {
        prefixlist := $1.ast
        prefixlist.CollapseChildList()
        datalist := $3.ast
        datalist.CollapseChildList()
        node := NewNode("query", $1.token)
        node.AddChild(prefixlist) // prefix
        node.AddChild(datalist)   // data
        node.AddChild(NewNode("list", $1.token)) // missing units
        node.AddChild($4.ast)     // select
        yylex.(*golex).line = node
      }
    | DATA VarList UNITS UnitList SelectStatement {
        datalist := $2.ast
        datalist.CollapseChildList()
        unitlist := $4.ast
        unitlist.CollapseChildList()
        node := NewNode("query", $1.token)
        node.AddChild(NewNode("list", $1.token)) // missing prefix
        node.AddChild(datalist) // data
        node.AddChild(unitlist) // units
        node.AddChild($5.ast)   // select
        yylex.(*golex).line = node
      }
    | PrefixList DATA VarList UNITS UnitList SelectStatement {
        prefixlist := $1.ast
        prefixlist.CollapseChildList()
        datalist := $3.ast
        datalist.CollapseChildList()
        unitlist := $5.ast
        unitlist.CollapseChildList()
        node := NewNode("query", $1.token)
        node.AddChild(prefixlist) // prefix
        node.AddChild(datalist)   // data
        node.AddChild(unitlist)   // units
        node.AddChild($6.ast)     // select
        yylex.(*golex).line = node
      }
    ;

SelectStatement
    : SELECT VarList WHERE LBRACE RestrictionList RBRACE {
        varlist := $2.ast
        varlist.CollapseChildList()
        reslist := $5.ast
        reslist.CollapseChildList()
        node := NewNode("select", $1.token)
        node.AddChild(varlist)
        node.AddChild(reslist)
        $$.ast = node
      }
    ;

PrefixList
    : Prefix PrefixList {
        node := NewNode("list", $1.token)
        node.AddChild($1.ast)
        node.AddChild($2.ast)
        $$.ast = node
        
      }
    | Prefix {
        node := NewNode("list", $1.token)
        node.AddChild($1.ast)
        $$.ast = node
      }
    ;

Prefix
    : PREFIX Id COLON LT Uri GT {
        node := NewNode("prefix", $1.token)
        node.AddChild($2.ast)
        node.AddChild($5.ast)
        $$.ast = node
      }
    ;

UnitList
    : UnitPair UnitList {
        node := NewNode("list", $1.token)
        node.AddChild($1.ast)
        node.AddChild($2.ast)
        $$.ast = node
        
      }
    | UnitPair {
        node := NewNode("list", $1.token)
        node.AddChild($1.ast)
        $$.ast = node
      }
    ;

UnitPair
    : ConcreteEntity ARROW ConcreteEntity {
        node := NewNode("mapping", $2.token)
        node.AddChild($1.ast)
        node.AddChild($3.ast)
        $$.ast = node
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
    | LBRACE Restriction RBRACE UNION LBRACE Restriction RBRACE {
        node := NewNode("union", $4.token)
        node.AddChild($2.ast)
        node.AddChild($6.ast)
        $$.ast = node
      }
    ;

Entity
    : Var {
        $$.ast = NewNode("var", $1.token)
      }
    | Literal {
        $$.ast = $1.ast
      }
    | ConcreteEntity {
        $$.ast = $1.ast
      }
    ;

ConcreteEntity
    : LT URI GT {
        $$.ast = $2.ast
      }
    | PrefixedEntity {
        $$.ast = $1.ast
      }
    ;

Uri
    : URI {
        $$.ast = NewNode("uri", $1.token)
      }
    ;

Id
    : ID {
        $$.ast = NewNode("id", $1.token)
      }
    ;

PrefixedEntity
    : Id COLON Id {
        node := NewNode("prefixed", $2.token)
        node.AddChild($1.ast)
        node.AddChild($3.ast)
        $$.ast = node
      }
    ;

Literal
    : STRING {
        node := NewNode("string", $1.token)
        $$.ast = node
      }
    ;

Path
    : Path SLASH Path {
        node := NewNode("sequence", $2.token)
        node.AddChild($1.ast)
        node.AddChild($3.ast)
        $$.ast = node
      }
    | Path VBAR Path {
        node := NewNode("choice", $2.token)
        node.AddChild($1.ast)
        node.AddChild($3.ast)
        $$.ast = node
      }
    | Step PLUS {
        node := NewNode("one-or-more", $2.token)
        node.AddChild($1.ast)
        $$.ast = node
      }
    | Step ASTERISK {
        node := NewNode("zero-or-more", $2.token)
        node.AddChild($1.ast)
        $$.ast = node
      }
    | LPAR Path RPAR {
        $$.ast = $2.ast
      }
    | Step {
        $$.ast = $1.ast
      }
    ;

Step
    : Var {
        $$.ast = $1.ast
      }
    ;

%%
