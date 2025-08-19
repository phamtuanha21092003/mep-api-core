FROM debian:bookworm-slim

ENV TUSD_VERSION=v2.8.0

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN curl -L "https://github.com/tus/tusd/releases/download/${TUSD_VERSION}/tusd_linux_amd64.tar.gz" \
    | tar -xz --strip-components=1 -C /usr/local/bin \
    && chmod +x /usr/local/bin/tusd

COPY bin/start_tusd_server.sh bin/start_tusd_server.sh

RUN chmod +x bin/start_tusd_server.sh

EXPOSE 8000

CMD ["bin/start_tusd_server.sh"]
