FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="aerospace-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="Aerospace Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="caeba412-84e1-4e7f-b07b-9b6efa4dc117"

WORKDIR /app

COPY bin/aerospace-on-golang .

RUN chmod +x aerospace-on-golang

EXPOSE 8080

ENTRYPOINT ["./aerospace-on-golang"]