sudo docker-compose -f ./deployment/docker-compose.yaml down && \
sudo docker-compose -f ./deployment/docker-compose.yaml pull sonyamoonglade/ci_cd_app && \
sudo docker-compose -f ./deployment/docker-compose.yaml up -d
