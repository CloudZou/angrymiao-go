#!/usr/bin/env bash
sudo docker build -t angrymiao/discovery-backend .
sudo docker login --username=angrymiao registry.cn-shenzhen.aliyuncs.com
sudo docker tag a8c027308e9e registry.cn-shenzhen.aliyuncs.com/angrymiao/rep:discovery-backend-1.0
sudo docker push registry.cn-shenzhen.aliyuncs.com/angrymiao/rep:discovery-backend-1.0