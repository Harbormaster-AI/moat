FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="healthcare-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Healthcare Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="22215a5f-86fb-4c60-af3b-a7d26ecc56b9"

WORKDIR /app

COPY bin/healthcare-on-golang .

RUN chmod +x healthcare-on-golang

EXPOSE 8080

ENTRYPOINT ["./healthcare-on-golang"]