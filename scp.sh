cat ./ssh_key && \
echo $VM_USER && \
scp -i ./ssh_key ./docker-compose.yaml $VM_USER@$VM_HOST:. && \
scp -i ./ssh_key ./migrations $VM_USER@$VM_HOST:./migrations && \
scp -i ./ssh_key ./setup.sh $VM_USER@$VM_HOST:./setup.sh
