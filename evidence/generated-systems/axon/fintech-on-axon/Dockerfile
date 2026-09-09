FROM eclipse-temurin:11-jre-alpine

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="fintechOnAxon"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Axon4-Framework-Server"
LABEL com.harbormaster.model="FinTech Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-09"
#LABEL com.harbormaster.certification="ceef30d7-3406-40f8-ab65-e08c38668c4f"

RUN addgroup -S spring && adduser -S -G spring spring
USER spring:spring

ARG JAR_FILE_RELATIVE_LOCATION=.
ARG JAR_FILE=${JAR_FILE_RELATIVE_LOCATION}/*.jar

COPY ${JAR_FILE} app.jar

EXPOSE 8080

ENTRYPOINT ["java","-jar","/app.jar"]