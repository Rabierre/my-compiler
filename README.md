BNF description for LL(>=1) grammars

```
Program ::= DeclList ?
DeclList ::= ( VarDecl | FunctionDecl ) DeclList ?
FunctionDecl ::= Type identifier "(" VarDeclList ? ")" CompoundStmt
VarDeclList ::= VarDecl VarDeclList ?
VarDecl ::= Type IdentList
IdentList ::= identifier ( "," IdentList ) ?
       | identifier ( "=" Expr ) ?
Type ::= "int"
       | "double"
Stmt ::= ForStmt
       | Expr
       | IfStmt
       | CompoundStmt
       | ReturnStmt
ForStmt ::= "for" "(" OptExpr ";" OptExpr ";" OptExpr ")" CompoundStmt
OptExpr ::= Expr ?
IfStmt ::= "if" "(" Expr ")" CompoundStmt ElsePart
ElsePart ::= ( "else" CompoundStmt ) ?
CompoundStmt ::= "{" VarDeclList ? StmtList ? "}"
ReturnStmt ::= "return" Expr ?
StmtList ::= Stmt StmtList ?
Expr ::= identifier "=" Expr
       | Term
       | Factor
Term ::= Factor ( Op Term )?
Factor ::= "(" Expr ")"
         | AddSub Factor
         | identifier "(" ExprList ? ")"
         | identifier
         | number
         | string
ExprList ::= Expr ( "," ExprList ) ?
Comment ::= "//" string ?
Op ::= "=="
     | "<"
     | ">"
     | "<="
     | ">="
     | "!="
     | "*"
     | "/"
     | "+="
     | "-="
     | "*="
     | "/="
     | "&="
     | "|="
     | "&&"
     | "||"
     | "&"
     | "|"
     | "^"
AddSub ::= "+"
         | "-"
```