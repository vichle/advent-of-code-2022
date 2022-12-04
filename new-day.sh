#!/bin/sh

mkdir $1 && cd $1
cp ../template_main.go main.go
go mod init github.com/vichle/advent-of-code-2022/$1
cd ..
go work use ./$1
