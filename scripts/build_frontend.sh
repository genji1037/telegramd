#!/usr/bin/env bash

cd ..
telegramd=`pwd`

echo "build frontend ..."
cd ${telegramd}/server/access/frontend
go build
ps -ef | grep frontend |grep -v build_frontend|grep -v 'grep'|awk '{print $2}'|xargs kill -9
nohup ./frontend > frontend.log &

