all: push

# 0.0 shouldn't clobber any released builds
TAG =0.14
PREFIX = remotejob/huoneisto

binary: server.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o server

container: binary
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

set: 
	ssh root@159.203.107.223 kubectl set image deployment/huoneisto huoneisto=$(PREFIX):$(TAG) -n huoneisto

clean:
	docker rmi -f $(PREFIX):$(TAG) || true
