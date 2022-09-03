#!/bin/bash

mv ./migrations ./deployment/migrations && \
sudo docker-compose -f ./deployment/docker-compose.yaml stop && \
sudo docker pull sonyamoonglade/ci_cd_app && \
sudo docker-compose -f ./deployment/docker-compose.yaml --env-file ./.env up -d
