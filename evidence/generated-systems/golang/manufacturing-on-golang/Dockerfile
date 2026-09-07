FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="manufacturing-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Manufacturing Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="f83987b6-0201-4b8a-98f7-55841588f155"

WORKDIR /app

COPY bin/manufacturing-on-golang .

RUN chmod +x manufacturing-on-golang

EXPOSE 8080

ENTRYPOINT ["./manufacturing-on-golang"]