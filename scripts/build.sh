#!/usr/bin/env bash

# todo(yumcoder) change dc ip
# sed -i '/ipAddress = /c\ipAddress = 127.0.0.1' a.txt
# todo(yumcoder) change folder path for nbfs

#docker start mysql-docker redis-docker etcd-docker

telegramd="/Users/cyy/go/src/strogo/telegramd"

echo "build frontend ..."
cd ${telegramd}/server/access/frontend
go get
go build
./frontend &

echo "build auth_key ..."
cd ${telegramd}/server/access/auth_key
go get
go build
./auth_key &

echo "build auth_session ..."
cd ${telegramd}/server/access/auth_session
go get
go build
./auth_session &

echo "build sync ..."
cd ${telegramd}/server/sync
go get
go build
./sync &

echo "build upload ..."
cd ${telegramd}/server/upload
go get
go build
./upload &

echo "build document ..."
cd ${telegramd}/service/document
go get
go build
./document &

echo "build biz_server ..."
cd ${telegramd}/server/biz_server
go get
go build
./biz_server &

echo "build session ..."
cd ${telegramd}/server/access/session
go get
go build
./session &

echo "***** wait *****"
wait
