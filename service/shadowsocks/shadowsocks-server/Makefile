VERSION=v1.2.0
SHADOWSOCKS_VERSION=v1.14.3

docker: build-image
	docker push hiwenzi/shadowsocks-server:${VERSION}

build-image:
	docker build -t hiwenzi/shadowsocks-server:${VERSION} --build-arg ss_version=${SHADOWSOCKS_VERSION} --no-cache .
	docker image prune -f