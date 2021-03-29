#!/bin/sh

#gomake
go build

./mysite -rpc=true -rpc_port=10001 -http=false&
master_pid=$!

sleep 1

./mysite -master=localhost:10001 -http=:8082 &
slave_pid1=$!

./mysite -http=:8083 &
slave_pid2=$!

echo "Visit: http://localhost:8082/v1/ping"
echo "Visit: http://localhost:8083/v1/ping"
echo "Press enter to shut down"

# shellcheck disable=SC2162
read
kill $master_pid
kill $slave_pid1
kill $slave_pid2