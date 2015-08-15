#!/bin/bash
# filename run
if [ -f "$1/$2.go" ];then
	go run "$1/$2.go"
else
	echo "not found files, dir: $1, file: $2"
fi