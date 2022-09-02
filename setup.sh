docker-compose -f ./deployment/docker-compose.yaml down && \
docker-compose -f ./deployment/docker-compose.yaml pull sonyamoonglade/ci_cd_app:latest && \
docker-compose -f ./deployment/docker-compose.yaml up -d
