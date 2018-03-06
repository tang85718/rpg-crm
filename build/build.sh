#!/usr/bin/env bash
#go build -o p-vip-gateway ../main.go
echo "构建 go 应用..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o account ../main.go
# 登录 build hub
docker login --username=tangxuyao@163.com registry.cn-shenzhen.aliyuncs.com

echo "构建 build 本地镜像..."
# 构建 build 镜像并上传到 build hub
export VERSION=1.0

docker build -t crm:${VERSION} .
docker tag crm:${VERSION} registry.cn-shenzhen.aliyuncs.com/txy/local/crm:${VERSION}
docker push registry.cn-shenzhen.aliyuncs.com/txy/local/crm:${VERSION}

#docker tag crm:${VERSION} registry.cn-shenzhen.aliyuncs.com/txy/local/crm:latest
#docker push registry.cn-shenzhen.aliyuncs.com/txy/local:latest

echo "构建完成"