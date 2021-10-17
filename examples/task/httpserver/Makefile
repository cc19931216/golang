export tag=v1.0

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t sjc1993/httpserver:${tag} .

push: release
	echo "pushing cc/httpserver"
	docker push sjc1993/httpserver:${tag}
