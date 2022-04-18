all: clean compile-linux

clean:
	rm -R bin/

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

linux-all: clean linux-386 linux-amd64 linux-arm linux-arm64
linux-amd64: clean linux-amd64
osx-arm64: clean osx-arm64

install-docker:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/panasonic-aquarea-smart-cloud-mqtt-linux
	docker build
	docker push
