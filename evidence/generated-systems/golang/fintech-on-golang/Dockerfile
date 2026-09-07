FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="fintech-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="FinTech Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="dc96afef-ed3f-42d5-a52e-22438d46afdc"

WORKDIR /app

COPY bin/fintech-on-golang .

RUN chmod +x fintech-on-golang

EXPOSE 8080

ENTRYPOINT ["./fintech-on-golang"]