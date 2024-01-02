FROM kitabisa/debian-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/kitabisa/villager-db"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/villager-db

WORKDIR /opt/villager-db

COPY ./bin/villager-db /opt/villager-db/
COPY ./migrations/sql /opt/villager-db/migrations/sql/

RUN chmod +x /opt/villager-db/villager-db

# Create appuser
RUN adduser --disabled-password --gecos '' villager-db
USER villager-db

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/villager-db/bin/villager-db"]
