#!/bin/bash

timeout 5 ping -n 1 $VM_HOST > ./result

LOST=`cat ./result | grep -ow 'Lost = [0-9]' | grep -ow '[0-9]'`

if test -z $LOST;
  then
    LOST=1
fi

rm ./result

echo $LOST

export LOST_PKG=$LOST