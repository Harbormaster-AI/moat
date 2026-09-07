FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="analytics-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Analytics Capabilities Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="fc8eec88-1e86-40f3-bd66-824e03ddac7e"

WORKDIR /app

COPY bin/analytics-on-golang .

RUN chmod +x analytics-on-golang

EXPOSE 8080

ENTRYPOINT ["./analytics-on-golang"]