parser grammar TExprParser;


options { tokenVocab=TExprLexer; }


@header {
    package com.aliyun.tauris.expression.ast;
}

parse
 : expression EOF
 ;

expression
 : LPAREN expression RPAREN                       #parenExpression
 | NOT expression                                 #notExpression
 | expression IN container                        #inExpression
 | expression NOT IN container                    #notInExpression
 | expression MATCH regex                         #matchExpression
 | expression IS type                             #isTypeExpression
 | expression IS NOT type                         #isNotTypeExpression
 | expression comparator expression               #comparatorExpression
 | expression binary expression                   #binaryExpression
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

bool
 : TRUE | FALSE
 ;

literal
   : String
   | Integer
   | Float
   | Boolean
   ;

container
    : array
    | variable
    | String
    ;

array
   : LBRACKET integers? RBRACKET
   | LBRACKET strings? RBRACKET
   | LBRACKET floats? RBRACKET
   | LBRACKET booleans? RBRACKET
   ;

calc
   : bit
   ;

bit
   : bit BAND shift
   | bit BEOR shift
   | bit BIOR shift
   | shift
   ;

shift
   : shift LSHIFT plus
   | shift RSHIFT plus
   | shift RSHIFT3 plus
   | plus
   ;

plus
   : plus PLUS multiplying
   | plus MINUS multiplying
   | multiplying
   ;

multiplying
   : multiplying MUL atom
   | multiplying DIV atom
   | multiplying MOD atom
   | atom
   ;

atom
   : variable
   | scientific
   | LPAREN bit RPAREN
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
   : REGEX
   ;

type
   : IDENTIFIER
   ;

strings
    : String (COMMA String)* COMMA?
    ;

integers
    : Integer (COMMA Integer)* COMMA?
    ;

floats
    : Float (COMMA Float)* COMMA?
    ;

booleans
    : Boolean (COMMA Boolean)* COMMA?
    ;
