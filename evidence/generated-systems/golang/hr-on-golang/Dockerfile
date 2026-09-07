FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="hr-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Human Resources Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="dfd96a93-7c9f-4d50-9e29-a9cd73f56320"

WORKDIR /app

COPY bin/hr-on-golang .

RUN chmod +x hr-on-golang

EXPOSE 8080

ENTRYPOINT ["./hr-on-golang"]