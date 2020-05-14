#!/bin/sh

export PATH=/usr/local/bin:$PATH

antlr -o ./ -no-listener -visitor -Dlanguage=Go -package ast TExpr.g4
