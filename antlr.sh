#!/bin/sh

export PATH=/usr/local/bin:$PATH

antlr -o ./ast -no-listener -visitor -Dlanguage=Go -package ast TExprLexer.g4 TExprParser.g4
