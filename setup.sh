docker-compose -f ./compose/docker-compose.yaml down && \
docker-compose -f ./compose/docker-compose.yaml pull sonyamoonglade/ci_cd_app:latest && \
docker-compose -f ./compose/docker-compose.yaml up -d