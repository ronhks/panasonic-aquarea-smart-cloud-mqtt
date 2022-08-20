version = 1.1.4

all: clean compile-all docker

help:
	echo "Available params:"
	echo "  * all: clean compile-linux"
	echo "  * docker: build image"
	echo "  * docker-push: push the image"

clean:
	rm -R bin
	mkdir bin

compile-linux-all: linux-386 linux-amd64 linux-arm linux-arm64

linux-386:
	echo "Compiling bin for linux-386"
	GOOS=linux GOARCH=386 go build -o bin/linux/panasonic-aquarea-smart-cloud-mqtt-linux-386

linux-amd64:
	echo "Compiling bin for linux-amd64"
	GOOS=linux GOARCH=amd64 go build -o bin/linux/panasonic-aquarea-smart-cloud-mqtt-linux-amd64

linux-arm:
	echo "Compiling bin for linux-arm"
	GOOS=linux GOARCH=arm go build -o bin/linux/panasonic-aquarea-smart-cloud-mqtt-linux-arm

linux-arm64:
	echo "Compiling bin for linux-arm64"
	GOOS=linux GOARCH=arm64 go build -o bin/linux/panasonic-aquarea-smart-cloud-mqtt-linux-arm64

compile-osx-all: osx-amd64 osx-arm64

osx-arm64:
	echo "Compiling bin for osx-arm64"
	GOOS=darwin GOARCH=arm64 go build -o bin/osx/panasonic-aquarea-smart-cloud-mqtt-osx-arm64

osx-amd64:
	echo "Compiling bin for osx-amd64"
	GOOS=darwin GOARCH=amd64 go build -o bin/osx/panasonic-aquarea-smart-cloud-mqtt-osx-amd64

compile-win-all: win-386 win-amd64 win-arm

win-arm:
	echo "Compiling bin for windows-arm64"
	GOOS=windows GOARCH=arm go build -o bin/win/panasonic-aquarea-smart-cloud-mqtt-windows-arm.exe

win-386:
	echo "Compiling bin for windows-386"
	GOOS=windows GOARCH=386 go build -o bin/win/panasonic-aquarea-smart-cloud-mqtt-windows-386.exe

win-amd64:
	echo "Compiling bin for windows-arm64"
	GOOS=windows GOARCH=amd64 go build -o bin/win/panasonic-aquarea-smart-cloud-mqtt-windows-amd64.exe

docker-linux-amd64: clean
	docker build -f ./Dockerfile-linux-amd64 . --tag ronhks/panasonic-aquarea-smart-cloud-mqtt:$(version)-amd64
	docker tag ronhks/panasonic-aquarea-smart-cloud-mqtt:$(version)-amd64 ronhks/panasonic-aquarea-smart-cloud-mqtt:latest-amd64

docker-linux-armv7: clean
	docker build -f ./Dockerfile-linux-arm . --tag ronhks/panasonic-aquarea-smart-cloud-mqtt:$(version)-arm
	docker tag ronhks/panasonic-aquarea-smart-cloud-mqtt:$(version)-arm ronhks/panasonic-aquarea-smart-cloud-mqtt:latest-arm


docker-linux-arm64: clean
	docker build -f ./Dockerfile-linux-arm64 . --tag ronhks/panasonic-aquarea-smart-cloud-mqtt:$(version)-arm64
	docker tag ronhks/panasonic-aquarea-smart-cloud-mqtt:$(version)-arm64 ronhks/panasonic-aquarea-smart-cloud-mqtt:latest-arm64

docker-clean:
	docker images 'ronhks/panasonic-aquarea-smart-cloud-mqtt' -a -q | xargs -r docker rmi -f $(docker images | grep 'ronhks/panasonic-aquarea-smart-cloud-mqtt')

docker-build: docker-linux-amd64 docker-linux-armv7 docker-linux-arm64

docker-push:
	docker push ronhks/panasonic-aquarea-smart-cloud-mqtt:$(version)-arm
	docker push ronhks/panasonic-aquarea-smart-cloud-mqtt:latest-arm
	docker push ronhks/panasonic-aquarea-smart-cloud-mqtt:$(version)-arm64
	docker push ronhks/panasonic-aquarea-smart-cloud-mqtt:latest-arm64
	docker push ronhks/panasonic-aquarea-smart-cloud-mqtt:$(version)-amd64
	docker push ronhks/panasonic-aquarea-smart-cloud-mqtt:latest-amd64

git-tag:
	git tag $(version)
	git push

docker-build-all-and-push-all: docker-clean docker-build docker-push

compile-all: compile-linux-all compile-osx-all compile-win-all
release: clean compile-all docker-build-all-and-push-all