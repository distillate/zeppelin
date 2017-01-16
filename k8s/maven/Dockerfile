FROM maven:3.3.9-jdk-8

MAINTAINER Alexander Andryashin <aandryashin@gmail.com>

RUN groupadd -r maven && useradd -m -r -g maven maven

# TODO: Add minio client to image with the following
#    curl -o /usr/local/bin/mc https://dl.minio.io/client/mc/release/linux-amd64/mc && \
#    chmod +x /usr/local/bin/mc

COPY entrypoint.sh /

RUN chmod +x /entrypoint.sh

USER maven

ENTRYPOINT ["/entrypoint.sh"]

CMD ["mvn", "clean", "test"]
