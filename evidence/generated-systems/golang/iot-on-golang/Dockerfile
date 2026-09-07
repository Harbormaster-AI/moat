FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="iot-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="IoT Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="17ac5a5a-d135-4ffb-9729-37337902f2fd"

WORKDIR /app

COPY bin/iot-on-golang .

RUN chmod +x iot-on-golang

EXPOSE 8080

ENTRYPOINT ["./iot-on-golang"]