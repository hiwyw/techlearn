VERSION=v1.0.0

docker: build-image
	docker push hiwenzi/kcptun-server:${VERSION}

build-image:
	docker build -t hiwenzi/kcptun-server:${VERSION} --no-cache .
	docker image prune -f