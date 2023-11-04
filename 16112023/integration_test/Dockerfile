FROM bitnami/minideb:latest
RUN apt-get update && apt-get -y install netcat-traditional tzdata ca-certificates && apt-get clean
ARG APPLICATION_NAME=kibitalk
WORKDIR /opt/$APPLICATION_NAME
COPY bin/$APPLICATION_NAME /opt/$APPLICATION_NAME/$APPLICATION_NAME
ENV TZ=Asia/Jakarta
