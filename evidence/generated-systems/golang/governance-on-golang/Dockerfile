FROM golang:1.22-alpine AS builder

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="governance-on-golang"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Golang"
LABEL com.harbormaster.model="HCIDL Governance Capabilities Model"
LABEL com.harbormaster.generated="2026-09-07"
#LABEL com.harbormaster.certification="6acb267f-1867-4f37-9b3a-c8a0455c627c"

WORKDIR /app

COPY bin/governance-on-golang .

RUN chmod +x governance-on-golang

EXPOSE 8080

ENTRYPOINT ["./governance-on-golang"]