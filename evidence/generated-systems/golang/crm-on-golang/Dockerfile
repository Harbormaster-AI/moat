FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="crm-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="CRM Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="f92681ac-1eab-4f06-bdd3-957185bfb1d6"

WORKDIR /app

COPY bin/crm-on-golang .

RUN chmod +x crm-on-golang

EXPOSE 8080

ENTRYPOINT ["./crm-on-golang"]