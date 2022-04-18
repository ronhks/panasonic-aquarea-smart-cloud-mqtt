all: clean compile-linux

clean:
	rm -R bin
	mkdir bin

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

osx-arm64:
	echo "Compiling bin for osx-arm64"
	GOOS=darwin GOARCH=arm64 go build -o bin/osx/panasonic-aquarea-smart-cloud-mqtt-osx-arm64

osx-amd64:
	echo "Compiling bin for osx-amd64"
	GOOS=darwin GOARCH=amd64 go build -o bin/osx/panasonic-aquarea-smart-cloud-mqtt-osx-amd64

win-arm:
	echo "Compiling bin for windows-arm64"
	GOOS=windows GOARCH=arm go build -o bin/win/panasonic-aquarea-smart-cloud-mqtt-windows-arm.exe

win-386:
	echo "Compiling bin for windows-386"
	GOOS=windows GOARCH=386 go build -o bin/win/panasonic-aquarea-smart-cloud-mqtt-windows-386.exe

win-amd64:
	echo "Compiling bin for windows-arm64"
	GOOS=windows GOARCH=amd64 go build -o bin/win/panasonic-aquarea-smart-cloud-mqtt-windows-amd64.exe

linux-all: clean linux-386 linux-amd64 linux-arm linux-arm64
linux-amd64: clean linux-amd64

osx-arm64: clean osx-arm64

win-all: clean win-arm win-amd64 win-386

release-for-github: clean linux-386 linux-amd64 linux-arm linux-arm64 osx-amd64 osx-arm64 win-386 win-amd64 win-arm64
release: release-for-github docker

docker:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/panasonic-aquarea-smart-cloud-mqtt-linux
	docker build . --tag ronhks/panasonic-aquarea-smart-cloud-mqtt:1.1.2
	docker tag ronhks/panasonic-aquarea-smart-cloud-mqtt:1.1.2 ronhks/panasonic-aquarea-smart-cloud-mqtt:latest

docker-push:
	docker push ronhks/panasonic-aquarea-smart-cloud-mqtt:1.1.2
	docker push ronhks/panasonic-aquarea-smart-cloud-mqtt:latest
