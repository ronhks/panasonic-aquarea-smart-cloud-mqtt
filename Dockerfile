FROM alpine

COPY bin/linux /app
WORKDIR /app
CMD [ "./panasonic-aquarea-smart-cloud-mqtt-linux"]
