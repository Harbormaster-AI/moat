FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="inventory-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="HCIDL Inventory Capabilities Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="79caa871-f298-436a-86d6-b46b898d903a"

WORKDIR /app

COPY bin/inventory-on-golang .

RUN chmod +x inventory-on-golang

EXPOSE 8080

ENTRYPOINT ["./inventory-on-golang"]