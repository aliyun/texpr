grammar TExpr;

parse
 : expression EOF
 ;

expression
 : LPAREN expression RPAREN                       #parenExpression
 | NOT expression                                 #notExpression
 | expression IN array                            #inExpression
 | expression NOT IN array                        #notInExpression
 | expression IS kind                             #isTypeExpression
 | expression IS NOT kind                         #isNotTypeExpression
 | expression comparator expression               #comparatorExpression
 | expression binary expression                   #binaryExpression
 | expression MATCH regex                         #matchExpression
 | variable                                       #variableExpression
 | calc                                           #calcExpression
 ;

variable
 : VARIABLE
 | literal
 | array
 ;

comparator
 : GT | GE | LT | LE | EQ | NE
 ;

binary
 : AND | OR
 ;

boolean
 : TRUE | FALSE
 ;

literal
   : Varchar
   | Integer
   | Float
   | Boolean
   ;

array
   : '[' integers? ']'
   | '[' strings? ']'
   | '[' floats? ']'
   | '[' booleans? ']'
   ;

calc
   : plus
   ;

plus
   : plus PLUS multiplying
   | plus MINUS multiplying
   | multiplying
   ;

multiplying
   : multiplying MUL pow
   | multiplying DIV pow
   | multiplying MOD pow
   | pow
   ;

pow
   : atom POW atom
   | atom
   ;

atom
   : variable
   | scientific
   | LPAREN plus RPAREN
   | function
   ;

scientific
   : number E number
   | number
   ;

function
   : funcname LPAREN parameters RPAREN
   ;

funcname
   : IDENTIFIER
   ;

parameters
   : expression (',' expression)*
   ;

number
   : MINUS? DIGIT + (POINT DIGIT +)?
   ;

regex
   : Regex
   ;

kind
   : IDENTIFIER
   ;

strings
    : Varchar (',' Varchar)* ','?
    ;

integers
    : Integer (',' Integer)* ','?
    ;

floats
    : Float (',' Float)* ','?
    ;

booleans
    : Boolean (',' Boolean)* ','?
    ;

Integer
   : '-' ? INT
   ;

Float
   : '-' ? INT '.' INT
   ;

Boolean
   : TRUE
   | FALSE
   ;

Varchar
   : STRING
   ;

Regex
   : REGEX
   ;

fragment REGEX
   : '/' ( STRING_ESCAPE_SEQ | ~[\\\r\n'] )* '/'
   ;

fragment STRING
   : '\'' ( STRING_ESCAPE_SEQ | ~[\\\r\n'] )* '\''
   | '"' ( STRING_ESCAPE_SEQ | ~[\\\r\n"] )* '"'
   ;

fragment STRING_ESCAPE_SEQ
   : '\\' .
   ;

fragment ESC
   : '\\' (["\\/bfnrt])
   ;

fragment NAME
   : [a-zA-Z][a-zA-Z0-9_\\.]+
   ;


AND        : '&&' ;
OR         : '||' ;
NOT        : 'not';
IS         : 'is';
IN         : 'in';
TRUE       : 'true' ;
FALSE      : 'false' ;
GT         : '>' ;
GE         : '>=' ;
LT         : '<' ;
LE         : '<=' ;
EQ         : '==' ;
NE         : '!=' ;
MATCH      : '=~';
LPAREN     : '(' ;
RPAREN     : ')' ;
INT        : '-'? [0-9]+ ;
FLOAT      : '-'? [0-9]+ ( '.' [0-9]+ )? ;
IDENTIFIER : [a-zA-Z_] [a-zA-Z_0-9]* ;
VARIABLE
    : '$'[a-zA-Z_] [.a-zA-Z_0-9]*
    | '@'[a-zA-Z_] [.a-zA-Z_0-9]*
    ;

COS  : 'cos';
SIN  : 'sin';
TAN  : 'tan';
ACOS : 'acos';
ASIN : 'asin';
ATAN : 'atan';
LN   : 'ln';
LOG  : 'log';
PLUS : '+';
MINUS: '-';
MUL  : '*';
DIV  : '/';
MOD  : '%';
POINT: '.';
E    : 'e' | 'E';
POW  : '^';
NL   : '\n';
DIGIT: ('0' .. '9');

WS         : [ \r\t\u000C\n]+ -> skip;
