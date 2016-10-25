all: push

# 0.0 shouldn't clobber any released builds
TAG =0.11
PREFIX = gcr.io/jntlserv0/huoneisto

binary: server.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o server

container: binary
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

set: 
	 kubectl set image deployment/huoneisto huoneisto=$(PREFIX):$(TAG)

clean:
	docker rmi -f $(PREFIX):$(TAG) || true
