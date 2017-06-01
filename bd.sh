#!/bin/bash
cp -r controller/ ~/go/src/cluster/controller/
cp -r simhash/ ~/go/src/cluster/simhash/
go build cluster.go
