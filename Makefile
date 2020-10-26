
# 镜像标签
IMAGE_TAG:=go-web:v1.0.0
# 容器名字
CONTAINER_NAME:=go-web


default: http

# 运行http服务
http:
	go run main.go http


# docker部署
dd:
	# 删除旧容器
	docker stop $(CONTAINER_NAME) 2>/dev/null;true
	docker rm $(CONTAINER_NAME) 2>/dev/null;true
	# 删除旧镜像
	docker rmi $(IMAGE_TAG) 2>/dev/null;true
	# 构建新镜像
	docker build -f ./Dockerfile -t $(IMAGE_TAG) .
	# 删除中间镜像
	docker image prune -f
	# 创建新容器
	docker run -d --name $(CONTAINER_NAME) -p 80:8080 $(IMAGE_TAG)

	@echo "Congratulations on your successful deployment of go-web service!"
	@echo "Now,open http://127.0.0.1/ping to use it!"

# k8s 部署
kd:





