FROM eclipse-temurin:11-jre-alpine

LABEL org.opencontainers.image.vendor="Harbormaster"
LABEL org.opencontainers.image.title="hrOnAxon"
LABEL org.opencontainers.image.version="0.0.1"
LABEL com.harbormaster.blueprint="Axon4-Framework-Server"
LABEL com.harbormaster.model="Human Resources Industry Domain Model"
LABEL com.harbormaster.generated="2026-09-09"
#LABEL com.harbormaster.certification="2aa7109b-6184-4b51-8aa4-7e0f6a26aca6"

RUN addgroup -S spring && adduser -S -G spring spring
USER spring:spring

ARG JAR_FILE_RELATIVE_LOCATION=.
ARG JAR_FILE=${JAR_FILE_RELATIVE_LOCATION}/*.jar

COPY ${JAR_FILE} app.jar

EXPOSE 8080

ENTRYPOINT ["java","-jar","/app.jar"]