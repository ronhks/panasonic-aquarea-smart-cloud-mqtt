module github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt

go 1.17

require (
	github.com/BurntSushi/toml v1.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect
)

//replace  github.com/ronhks/panasonic-aquarea-smart-cloud-mqtt => ./src
