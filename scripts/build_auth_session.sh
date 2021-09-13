#!/usr/bin/env bash

cd ..
telegramd=`pwd`

echo "build auth_session ..."
cd ${telegramd}/service/auth_session
go build
ps -ef | grep auth_session |grep -v build_auth_session|grep -v 'grep'|awk '{print $2}'|xargs kill -9
nohup ./auth_session > auth_session.log &