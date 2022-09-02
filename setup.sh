
str=$DB_PWD

for((i=0; i<${#str}; i++)); do
  ch=${str:$i:1}
  echo $ch
done

sudo docker-compose -f ./deployment/docker-compose.yaml down && \
sudo docker-compose -f ./deployment/docker-compose.yaml pull sonyamoonglade/ci_cd_app && \
sudo docker-compose -f ./deployment/docker-compose.yaml up -d
