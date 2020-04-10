lexer grammar TExprLexer;

@header {
    package com.aliyun.tauris.expression.ast;
}

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

String
   : STRING
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
MATCH      : '=~' -> pushMode(REG);
LPAREN     : '(' ;
RPAREN     : ')' ;
INT        : '-'? [0-9]+ ;
FLOAT      : '-'? [0-9]+ ( '.' [0-9]+ )? ;
IDENTIFIER : [a-zA-Z_] [a-zA-Z_0-9]* ;
VARIABLE
    : '$'[a-zA-Z_] [.a-zA-Z_0-9]*
    | '@'[a-zA-Z_] [.a-zA-Z_0-9]*
    ;

PLUS : '+';
MINUS: '-';
MUL  : '*';
DIV  : '/';
MOD  : '%';
POINT: '.';
E    : 'e' | 'E';
LSHIFT : '<<' ;
RSHIFT : '>>' ;
RSHIFT3 : '>>>' ;
BAND  : '&'; //位与
BEOR  : '^'; //异或
BIOR  : '|'; //位或
NL   : '\n';
DIGIT: ('0' .. '9');
COMMA: ',';
LBRACKET: '[';
RBRACKET: ']';
WS: [ \r\t\u000C\n]+ -> skip;

// $var =~ /../
mode REG;

SLASH: '/';

REGEX
   :  [ ]* SLASH ( STRING_ESCAPE_SEQ | ~[\\\r\n/'] )+ SLASH -> popMode
   ;

