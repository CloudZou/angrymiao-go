#!/usr/bin/env bash
sudo docker build -t angrymiao/discovery .

sudo docker run -d -p 7171:80 angrymiao/discovery

sudo docker run -p 7171:80 angrymiao/discovery

sudo docker login --username=angrymiao registry.cn-shenzhen.aliyuncs.com
sudo docker tag 587c17f0ab08 registry.cn-shenzhen.aliyuncs.com/angrymiao/rep:discovery-backend-1.0
sudo docker push registry.cn-shenzhen.aliyuncs.com/angrymiao/rep:discovery-backend-1.0

kubectl create secret tls discovery-backend-tls --key cert.key --cert cert.pem