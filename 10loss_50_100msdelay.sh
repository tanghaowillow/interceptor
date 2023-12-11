#!/bin/bash

if [ $# -eq 0 ]; then
    echo "add 'add' or 'del'"
    exit
fi

if [ $1 == "add" ]; then
    tc qdisc add dev lo root netem loss 10%
    # tc qdisc add dev lo root netem delay 100ms 50ms
elif [ $1 == "del" ];then
    tc qdisc del dev lo root
else
    echo "add 'add' or 'del'"
fi