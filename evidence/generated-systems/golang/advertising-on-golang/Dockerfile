FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="advertising-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Advertising Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="541cd823-3f55-4f95-9dd7-2f33356715f1"

WORKDIR /app

COPY bin/advertising-on-golang .

RUN chmod +x advertising-on-golang

EXPOSE 8080

ENTRYPOINT ["./advertising-on-golang"]