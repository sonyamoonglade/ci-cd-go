#!/bin/bash
echo ("$VM_SSH") | scp -i /dev/stdin ./docker-compose.yaml $VM_USER@$VM_HOST:. && \
echo ("$VM_SSH") | scp -i /dev/stdin ./migrations $VM_USER@$VM_HOST:./migrations && \
echo ("$VM_SSH") | scp -i /dev/stdin ./setup.sh $VM_USER@$VM_HOST:./setup.sh
