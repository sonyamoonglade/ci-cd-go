docker-compose -f ./docker-compose.yaml down && \
docker-compose -f ./docker-compose.yaml pull sonyamoonglade/ci_cd_app:latest && \
docker-compose -f ./docker-compose.yaml up -d