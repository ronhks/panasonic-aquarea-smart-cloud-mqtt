# Panasonic Aquaera Smart Cloud integration woth MQTT
At home I have a Home Assistant and I would like control my Panasonic Aquarea Heatpump.

This project a transformation layer from [Panasonic Smart Cloud](https://aquarea-smart.panasonic.com/) with [MQTT](https://en.wikipedia.org/wiki/MQTT). 

If you satisfied, thanks for buying a coffee for me :)

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/zsoltdenes)


## How to start
### Install
`docker pull`
### Configure
Create a config file copy as `config` from [config.example](./config.example)
Change the parameters
### Run
`docker run`
### Usage
The MQTT topic are under the `MqttTopicRoot`.
#### Read values
* `/outdoor/temp/now` - Outdoor temp
* `/heat/temp/max` - Heating max. temp 
* `/heat/temp/min` - Heating min. temp
* `/heat/operation` - Heating operation status
  * `0` - OFF
  * `1` - ON
* `/water/temp/now` - Actual Hot Water temp
* `/water/temp/max` - Hot water max. temp.
* `/water/temp/min` - Hot water min. temp.
* `/water/operation` - Hot water operation status
  * `0` - OFF
  * `1` - ON
* `/operation` - Device operation status
  * `0` - OFF
  * `1` - ON

#### Change values
* `/water/temp/set` - set the HOT WATER temp. 
  * sample REQUEST JSON:
    ```json
    {
      "newTemp" : 42
    }
    ```
* `/water/operation/on` - Turn ON the Hot Water
* `/water/operation/off` - Turn OFF the Hot Water
* `/heat/operation/on` - Turn ON the heating
* `/heat/operation/off` - Turn OFF the heating
* `/operation/on` - Turn ON the device
* `/operation/off` - Turn OFF the device

## Contributing
Welcome all type of contributing! :)

### Requirements
* docker daemon
* MQTT server
* Panasonic Smart Cloud registration
* [Panasonic Aquarea T-CAP Wifi Module - Panasonic CZ-TAW1](https://www.panasonicproclub.com/uploads/PL/catalogues/CZ-TAW1_quick%20guide.pdf)
#### 3rd party libs
**Thanks for**
* [BurntSushi/toml](https://github.com/BurntSushi/toml) for easier config handling
* [eclipse/paho.mqtt.golan](https://github.com/eclipse/paho.mqtt.golan) for MQTT handling
* [sirupsen/logrus](https://github.com/sirupsen/logrus) for easier logging