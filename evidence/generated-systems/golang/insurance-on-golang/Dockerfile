FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="insurance-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Insurance Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="34d614a5-0db9-4850-8df7-ac0e6abf8afa"

WORKDIR /app

COPY bin/insurance-on-golang .

RUN chmod +x insurance-on-golang

EXPOSE 8080

ENTRYPOINT ["./insurance-on-golang"]