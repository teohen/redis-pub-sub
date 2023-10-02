#! /bin/bash
image="teohen/rinha-de-backend"
tag="local"

docker build --no-cache=true -t "${image}:${tag}" .

sudo docker-compose rm -f 
sudo docker-compose down
sudo docker-compose up --build

