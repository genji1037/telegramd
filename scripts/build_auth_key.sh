#!/usr/bin/env bash

cd ..
telegramd=`pwd`

echo "build auth_key ..."
cd ${telegramd}/server/access/auth_key
go build
ps -ef | grep auth_key |grep -v build_auth_key|grep -v 'grep'|awk '{print $2}'|xargs kill -9
nohup ./auth_key > auth_key.log &