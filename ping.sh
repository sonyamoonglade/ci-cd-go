#!/bin/bash

timeout 5 ping -n 1 $VM_HOST > ./result

LOST=`cat ./result | grep -ow 'Lost = [0-9]' | grep -ow '[0-9]'`

if test -z $LOST;
  then
    LOST=1
fi

rm ./result

foo=$VM_HOST

for ((i=0;i<${#foo};i++));do
  echo "${foo:$i:1}"

echo $LOST

export LOST_PKG=$LOST