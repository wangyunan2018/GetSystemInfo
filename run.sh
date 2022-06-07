#!/bin/bash

if go version && echo $? -eq 0; then
	go run main.go
else
	wget https://golang.google.cn/dl/go1.17.linux-amd64.tar.gz
	tar -C /usr/local/ -xzvf go1.17.linux-amd64.tar.gz
	
	cp /etc/profile /etc/profile.bak
	echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
	
	source /etc/profile
	go version
    go env -w GOPROXY=https://goproxy.cn
	go run main.go
fi